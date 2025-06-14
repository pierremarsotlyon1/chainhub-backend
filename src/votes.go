package src

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"main/contracts/erc20"
	"main/contracts/escrow"
	"main/contracts/voter"
	"main/interfaces"
	"main/utils"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var CURVE_OWNERSHIP_VOTER = common.HexToAddress("0xE478de485ad2fe566d49342Cbd03E49ed7DB3356")
var CURVE_PARAMETER_VOTER = common.HexToAddress("0xBCfF8B0b9419b9A88c44546519b1e909cF330399")

var VOTE_CREATE_TOPIC = common.HexToHash("0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394")
var VOTE_EXECUTED_TOPIC = common.HexToHash("0xbf8e2b108bb7c980e08903a8a46527699d5e84905a082d56dacb4150725c8cab")
var CAST_VOTE_TOPIC = common.HexToHash("0xb34ee265e3d4f5ec4e8b52d59b2a9be8fceca2f274ebc080d8fba797fea9391f")

const (
	VOTES_PATH        = "./data/votes.json"
	OWNERSHIP         = "ownership"
	PARAMETER         = "parameter"
	votes_config      = "./data/configs/votes-config.json"
	BUCKET_VOTES_DIR  = "data/votes"
	BUCKET_VOTES_FILE = BUCKET_VOTES_DIR + "/votes.json"
)

const BLOCK_BEFORE_SNAPSHOT = 10

func FetchVotes(client *ethclient.Client, currentBlock uint64) {
	fmt.Println("Fetching votes")

	config := utils.ReadConfig(votes_config)

	from := config.LastBlock
	if from == 0 {
		from = 10648598
	}

	votes := readVotes()
	fmt.Println("Already", len(votes), "votes")

	err := utils.ForEachBlockRange(from, currentBlock, 499, func(start uint64, end uint64) error {
		topicList := []common.Hash{
			VOTE_CREATE_TOPIC,
			VOTE_EXECUTED_TOPIC,
			CAST_VOTE_TOPIC,
		}
		var allLogs []types.Log

		for _, topic := range topicList {
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(start)),
				ToBlock:   big.NewInt(int64(end)),
				Addresses: []common.Address{
					CURVE_OWNERSHIP_VOTER,
					CURVE_PARAMETER_VOTER,
				},
				Topics: [][]common.Hash{
					{topic},
				},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				return fmt.Errorf("FetchVotes logs error for topic %s: %w", topic.Hex(), err)
			}

			allLogs = append(allLogs, logs...)
		}

		for _, vLog := range allLogs {
			voterContract, err := voter.NewVoter(vLog.Address, client)
			if err != nil {
				panic(err)
			}

			switch vLog.Topics[0].Hex() {
			case VOTE_CREATE_TOPIC.Hex():
				event, err := voterContract.ParseStartVote(vLog)
				if err != nil {
					fmt.Println(err)
					continue
				}

				found := false
				for _, vote := range votes {
					if strings.EqualFold(vote.Voter.Hex(), vLog.Address.Hex()) && event.VoteId.Cmp(vote.Id) == 0 {
						found = true
						break
					}
				}
				if found {
					continue
				}

				v, err := voterContract.GetVote(nil, event.VoteId)
				if err != nil {
					panic(err)
				}

				voteTime, err := voterContract.VoteTime(nil)
				if err != nil {
					panic(err)
				}

				voteType := PARAMETER
				if strings.EqualFold(vLog.Address.Hex(), CURVE_OWNERSHIP_VOTER.Hex()) {
					voteType = OWNERSHIP
				}

				ipfsId := ""
				description := ""
				if strings.Contains(event.Metadata, "ipfs:") {
					ipfsId = event.Metadata[len("ipfs:"):]
					description = utils.GetIpfs(ipfsId)
				}

				calldata := "0x" + hex.EncodeToString(v.Script)

				votes = append(votes, interfaces.Vote{
					Id:              event.VoteId,
					Open:            v.Open,
					Executed:        v.Executed,
					StartDate:       v.StartDate,
					EndDate:         v.StartDate + voteTime,
					SnapshotBlock:   v.SnapshotBlock,
					SupportRequired: v.SupportRequired,
					MinAcceptQuorum: v.MinAcceptQuorum,
					Yea:             v.Yea,
					Nay:             v.Nay,
					VotingPower:     v.VotingPower,
					Script:          calldata,
					VoteType:        voteType,
					Voter:           vLog.Address,
					IpfsId:          ipfsId,
					Description:     description,
					Creator:         event.Creator,
					Voters:          make([]interfaces.Voter, 0),
				})

			case VOTE_EXECUTED_TOPIC.Hex(), CAST_VOTE_TOPIC.Hex():
				var voteId *big.Int
				var executedHash common.Hash
				var newVoter interfaces.Voter

				if vLog.Topics[0].Hex() == VOTE_EXECUTED_TOPIC.Hex() {
					event, err := voterContract.ParseExecuteVote(vLog)
					if err != nil {
						fmt.Println(err)
						continue
					}
					voteId = event.VoteId
					executedHash = vLog.TxHash

				} else {
					event, err := voterContract.ParseCastVote(vLog)
					if err != nil {
						fmt.Println(err)
						continue
					}
					voteId = event.VoteId
					newVoter = interfaces.Voter{
						Stake:    utils.Quo(event.Stake, 18),
						Supports: event.Supports,
						Voter:    event.Voter,
					}
				}

				for i := 0; i < len(votes); i++ {
					vote := votes[i]

					if strings.EqualFold(vote.Voter.Hex(), vLog.Address.Hex()) && voteId.Cmp(vote.Id) == 0 {
						v, err := voterContract.GetVote(nil, voteId)
						if err != nil {
							panic(err)
						}

						votes[i].Executed = v.Executed
						votes[i].Open = v.Open
						votes[i].Nay = v.Nay
						votes[i].Yea = v.Yea
						votes[i].ExecutedHash = executedHash

						if !utils.IsNullAddress(newVoter.Voter) {
							votes[i].Voters = append(votes[i].Voters, newVoter)
						}

						break
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("FetchVotes error:", err)
		return
	}

	start := len(votes) - 1
	for i := start; i >= start-20 && i >= 0; i-- {
		if votes[i].Id.Cmp(big.NewInt(int64(810))) == 1 && len(votes[i].TenderlySimulationUrl) == 0 {
			tenderlyUrl, err := getTenderlySimulation(votes[i])
			if err == nil {
				votes[i].TenderlySimulationUrl = tenderlyUrl
			} else {
				fmt.Println(votes[i].Id, err)
			}
		}
		if len(votes[i].Description) == 0 {
			votes[i].Description = utils.GetIpfs(votes[i].IpfsId)
		}
	}

	for i := range votes {
		votes[i].Voters = removeDuplicateVoters(votes[i].Voters)
	}

	writeVotes(votes)
	utils.WriteConfig(config, currentBlock, votes_config)
}

func removeDuplicateVoters(voters []interfaces.Voter) []interfaces.Voter {
	seen := make(map[common.Address]bool)
	uniqueVoters := []interfaces.Voter{}

	for _, v := range voters {
		if !seen[v.Voter] {
			seen[v.Voter] = true
			uniqueVoters = append(uniqueVoters, v)
		}
	}

	return uniqueVoters
}

func getTenderlySimulation(vote interfaces.Vote) (string, error) {
	// Create the fork before the snapshot block
	forkRpcUrl, forkId, _, err := createFork("1", int64(vote.SnapshotBlock)-BLOCK_BEFORE_SNAPSHOT)
	if err != nil {
		return "", err
	}

	forkClient, err := ethclient.Dial(forkRpcUrl)
	if err != nil {
		return "", err
	}

	// Approve crv for vesting contract
	abi, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		return "", err
	}

	data, err := abi.Pack("approve", utils.CURVE_ESCROW_ADDRESS, utils.Mul(10000000000, 18))
	if err != nil {
		return "", err
	}

	if err := sendTx(forkClient, utils.CRV, data); err != nil {
		return "", err
	}

	// Create the lock
	abi, err = escrow.EscrowMetaData.GetAbi()
	if err != nil {
		return "", err
	}

	now := time.Now().Unix()
	data, err = abi.Pack("create_lock", utils.Mul(10000000000, 18), big.NewInt(int64(uint64(now)+utils.WEEK_TO_SEC*50*4)))
	if err != nil {
		return "", err
	}

	if err := sendTx(forkClient, utils.CURVE_ESCROW_ADDRESS, data); err != nil {
		return "", err
	}

	// Add some blocks
	httpClient := &http.Client{}
	if err = mineBlocks(httpClient, forkRpcUrl, BLOCK_BEFORE_SNAPSHOT+1); err != nil {
		return "", err
	}

	// Create vote and vote yes but default
	abi, err = voter.VoterMetaData.GetAbi()
	if err != nil {
		return "", err
	}

	script, err := hex.DecodeString(vote.Script[2:])
	if err != nil {
		fmt.Println(vote.Script)
		return "", err
	}

	data, err = abi.Pack("newVote", script, "")
	if err != nil {
		fmt.Println("err new vote pack data")
		return "", err
	}
	if err := sendTx(forkClient, vote.Voter, data); err != nil {
		fmt.Println("err new vote")
		return "", err
	}

	// Mine blocks to execute after
	// 7200 blocks per day
	// Mine one month blocks
	if err = mineBlocks(httpClient, forkRpcUrl, 7200*30); err != nil {
		fmt.Println("err mine blocks")
		return "", err
	}

	if err = addTime(httpClient, forkRpcUrl, (int64(vote.EndDate)-now)+86400); err != nil {
		fmt.Println("err add time")
		return "", err
	}

	// Exec vote
	abi, err = voter.VoterMetaData.GetAbi()
	if err != nil {
		return "", err
	}

	data, err = abi.Pack("executeVote", vote.Id)
	if err != nil {
		return "", err
	}

	if err := sendTx(forkClient, vote.Voter, data); err != nil {
		return "", err
	}

	// Get tx list
	r, err := http.NewRequest("GET", "https://api.tenderly.co/api/v1/account/"+utils.GoDotEnvVariable("TENDERLY_SLUG")+"/project/"+utils.GoDotEnvVariable("TENDERLY_PROJECT_SLUG")+"/fork/"+forkId+"/transactions?page=1&perPage=20&exclude_internal=true", nil)
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Key", utils.GoDotEnvVariable("TENDERLY_ACCESS_KEY"))

	httpClient = &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		fmt.Println("err http get fork txs")
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err read http get fork txs", res.Body)
		return "", err
	}

	txs := new(interfaces.TenderlyForkTxListResponse)
	if err := json.Unmarshal(body, txs); err != nil {
		fmt.Println("err unmarshal http get fork txs", res.Body)
		return "", err
	}

	if len(txs.ForkTransactions) == 0 {
		return "", errors.New("txs empty")
	}

	return "https://dashboard.tenderly.co/explorer/fork/" + forkId + "/simulation/" + txs.ForkTransactions[0].Id, nil
}

func readVotes() []interfaces.Vote {

	votes := make([]interfaces.Vote, 0)

	b, err := utils.ReadBucketFile(BUCKET_VOTES_FILE)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &votes); err != nil {
			log.Fatal(err)
		}
		return votes
	}

	if !utils.FileExists(VOTES_PATH) {
		return make([]interfaces.Vote, 0)
	}

	file, err := os.ReadFile(VOTES_PATH)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &votes); err != nil {
		log.Fatal(err)
	}

	return votes
}

func writeVotes(votes []interfaces.Vote) {
	if err := utils.WriteBucketFile(BUCKET_VOTES_FILE, votes); err != nil {
		fmt.Println(err)
	}
}
