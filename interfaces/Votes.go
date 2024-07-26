package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Vote struct {
	Id                    *big.Int       `json:"id"`
	Open                  bool           `json:"open"`
	Executed              bool           `json:"executed"`
	StartDate             uint64         `json:"startDate"`
	EndDate               uint64         `json:"endDate"`
	SnapshotBlock         uint64         `json:"snapshotBlock"`
	SupportRequired       uint64         `json:"supportRequired"`
	MinAcceptQuorum       uint64         `json:"minAcceptQuorum"`
	Yea                   *big.Int       `json:"yea"`
	Nay                   *big.Int       `json:"nay"`
	VotingPower           *big.Int       `json:"votingPower"`
	Script                string         `json:"script"`
	VoteType              string         `json:"voteType"`
	Voter                 common.Address `json:"voter"`
	Description           string         `json:"description"`
	IpfsId                string         `json:"ipfsId"`
	Creator               common.Address `json:"creator"`
	Voters                []Voter        `json:"voters"`
	TenderlySimulationUrl string         `json:"tenderlySimulationUrl"`
}

type Voter struct {
	Voter    common.Address `json:"voter"`
	Supports bool           `json:"supports"`
	Stake    float64        `json:"stake"`
}
