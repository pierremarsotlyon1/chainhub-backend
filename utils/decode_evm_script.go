package utils

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"main/interfaces"
	"math/big"
	"net/http"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const aragon_target = "0x3A93C17FC82CC33420d1809dDA9Fb715cc89dd37"

type DecodedArgs struct {
	FunctionName string                 `json:"functionName"`
	Args         map[string]interface{} `json:"args"`
}

// -----------------------------------------------------------------------------
// stringifyBigInts recursively converts *big.Int values to strings for JSON
// -----------------------------------------------------------------------------
func stringifyBigInts(v interface{}) interface{} {
	switch val := v.(type) {
	case *big.Int:
		return val.String()
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for k, vv := range val {
			newMap[k] = stringifyBigInts(vv)
		}
		return newMap
	case []interface{}:
		newArr := make([]interface{}, len(val))
		for i, vv := range val {
			newArr[i] = stringifyBigInts(vv)
		}
		return newArr
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr && !rv.IsNil() {
			return stringifyBigInts(rv.Elem().Interface())
		}
		return v
	}
}

// -----------------------------------------------------------------------------
// fetch ABI from Etherscan
// -----------------------------------------------------------------------------
func getAbi(address string) (abi.ABI, error) {
	url := fmt.Sprintf(
		"https://api.etherscan.io/v2/api?chainid=1&module=contract&action=getabi&address=%s&apikey=%s",
		address,
		GoDotEnvVariable("ETHERSCAN_API_KEY"),
	)
	resp, err := http.Get(url)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to fetch ABI: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var res struct {
		Result string `json:"result"`
	}
	if err := json.Unmarshal(body, &res); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse response: %v", err)
	}
	return abi.JSON(strings.NewReader(res.Result))
}

// -----------------------------------------------------------------------------
// decode calldata
// -----------------------------------------------------------------------------
func decodeCalldata(to, calldata string) (*DecodedArgs, error) {
	parsedAbi, err := getAbi(to)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ABI for %s: %v", to, err)
	}
	if !strings.HasPrefix(calldata, "0x") {
		calldata = "0x" + calldata
	}
	if len(calldata) < 10 {
		return nil, fmt.Errorf("invalid calldata (too short)")
	}

	method, err := parsedAbi.MethodById(common.FromHex(calldata[:10]))
	if err != nil {
		return nil, fmt.Errorf("unknown function selector: %v", err)
	}
	data := common.FromHex(calldata[10:])
	values, err := method.Inputs.UnpackValues(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode args: %v", err)
	}
	args := make(map[string]interface{})
	for i, input := range method.Inputs {
		name := input.Name
		if name == "" {
			name = fmt.Sprintf("arg%d", i)
		}
		args[name] = values[i]
	}
	return &DecodedArgs{FunctionName: method.Name, Args: args}, nil
}

// -----------------------------------------------------------------------------
// recursive EVM script parsing returning structured data
// -----------------------------------------------------------------------------
func ParseEvmScript(evmHex string) ([]interfaces.DecodedAction, error) {
	raw, err := hex.DecodeString(strings.TrimPrefix(evmHex, "0x"))
	if err != nil {
		return nil, err
	}

	var actions []interfaces.DecodedAction
	offset := 0

	version := binary.BigEndian.Uint32(raw[:4])
	if version != 1 {
		fmt.Printf("⚠️ Unsupported EVM script version: %d\n", version)
	}
	offset = 4

	i := 1
	for offset < len(raw) {
		if len(raw[offset:]) < 24 {
			break
		}

		offset += 20
		dataLen := binary.BigEndian.Uint32(raw[offset : offset+4])
		offset += 4
		if offset+int(dataLen) > len(raw) {
			break
		}

		data := raw[offset : offset+int(dataLen)]
		offset += int(dataLen)

		decoded, err := decodeCalldata(aragon_target, "0x"+hex.EncodeToString(data))
		if err != nil {
			actions = append(actions, interfaces.DecodedAction{
				Index:    i,
				Target:   aragon_target,
				Function: fmt.Sprintf("⚠️ decode error: %v", err),
			})
			i++
			continue
		}

		action := interfaces.DecodedAction{
			Index:    i,
			Target:   aragon_target,
			Function: decoded.FunctionName,
			Args:     stringifyBigInts(decoded.Args).(map[string]interface{}),
		}

		// If Agent.execute → decode internal call recursively
		if decoded.FunctionName == "execute" {
			if dataIface, ok := decoded.Args["_data"].([]byte); ok {
				dataHex := "0x" + hex.EncodeToString(dataIface)
				targetAddr := decoded.Args["_target"].(common.Address).Hex()
				subDecoded, _ := decodeCalldata(targetAddr, dataHex)
				if subDecoded != nil {
					subAction := interfaces.DecodedAction{
						Index:    1,
						Target:   targetAddr,
						Function: subDecoded.FunctionName,
						Args:     stringifyBigInts(subDecoded.Args).(map[string]interface{}),
					}
					action.SubActions = append(action.SubActions, subAction)
				}
			}
		}

		actions = append(actions, action)
		i++
	}
	return actions, nil
}
