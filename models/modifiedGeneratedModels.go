package models

import "time"

type Block struct {
	ID            int64     `json:"id"`
	Height        int64     `json:"height"`
	Proposer      string    `json:"proposer"`
	Moniker       string    `json:"moniker"`
	BlockHash     string    `json:"blockHash" db:"block_hash"`
	ParentHash    string    `json:"parentHash" db:"parent_hash"`
	NumPrecommits int64     `json:"numPrecommits" db:"num_precommits"`
	NumTxs        int64     `json:"numTxs" db:"num_txs"`
	TotalTxs      int64     `json:"totalTxs" db:"total_txs"`
	Timestamp     time.Time `json:"timestamp"`
}

type PreCommit struct {
	ID               int64     `json:"id"`
	Height           int64     `json:"height"`
	Round            int64     `json:"round"`
	ValidatorAddress string    `json:"validatorAddress" db:"validator_address"`
	VotingPower      int64     `json:"votingPower" db:"voting_power"`
	ProposerPriority int64     `json:"proposerPriority" db:"proposer_priority"`
	Timestamp        time.Time `json:"timestamp"`
}

type Transaction struct {
	ID         int64     `json:"id"`
	Height     int64     `json:"height"`
	TxHash     string    `json:"txHash" db:"tx_hash"`
	Code       int64     `json:"code"`
	Messages   string    `json:"messages"`
	Signatures string    `json:"signatures"`
	Memo       *string   `json:"memo"`
	GasWanted  int64     `json:"gasWanted" db:"gas_wanted"`
	GasUsed    int64     `json:"gasUsed" db:"gas_used"`
	Timestamp  time.Time `json:"timestamp"`
}

type Validator struct {
	ID                int64     `json:"id"`
	OperatorAddress   string    `json:"operatorAddress" db:"operator_address"`
	ConsensusAddress  string    `json:"consensusAddress" db:"consensus_address"`
	ConsensusPubKey   string    `json:"consensusPubKey" db:"consensus_pub_key"`
	Moniker           string    `json:"moniker"`
	Identity          *string   `json:"identity"`
	Website           *string   `json:"website"`
	SecurityContact   *string   `json:"securityContact" db:"security_contact"`
	Details           *string   `json:"details"`
	Jailed            *bool     `json:"jailed"`
	Status            int64     `json:"status"`
	Tokens            string    `json:"tokens"`
	DelegatorShares   string    `json:"delegatorShares" db:"delegator_shares"`
	UnbondingHeight   string    `json:"unbondingHeight" db:"unbonding_height"`
	UnbondingTime     string    `json:"unbondingTime" db:"unbonding_time"`
	Rate              string    `json:"rate"`
	MaxRate           string    `json:"maxRate" db:"max_rate"`
	MaxChangeRate     string    `json:"maxChangeRate" db:"max_change_rate"`
	UpdateTime        string    `json:"updateTime" db:"update_time"`
	Timestamp         time.Time `json:"timestamp"`
	MinSelfDelegation string    `json:"minSelfDelegation" db:"min_self_delegation"`
}
