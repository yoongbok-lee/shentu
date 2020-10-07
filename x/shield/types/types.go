package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Pool struct {
	PoolID           uint64        `json:"pool_id" yaml:"pool_id"`
	Active           bool          `json:"active" yaml:"active"`
	Description      string        `json:"description" yaml:"description"`
	Sponsor          string        `json:"sponsor" yaml:"sponsor"`
	Premium          MixedDecCoins `json:"premium" yaml:"premium"`
	StartBlockHeight int64         `json:"start_block_height" yaml:"start_block_height"`
	TotalCollateral  sdk.Coins     `json:"total_collateral" yaml:"total_collateral"`
	Available        sdk.Int       `json:"available" yaml:"available"`
	Shield           sdk.Coins     `json:"shield" yaml:"shield"`
	EndTime          int64         `json:"end_time" yaml:"end_time"`
	EndBlockHeight   int64         `json:"end_block_height" yaml:"end_block_height"`
}

func NewPool(
	shield sdk.Coins, deposit MixedDecCoins, sponsor string,
	endTime, startBlockHeight, endBlockHeight int64, id uint64) Pool {
	return Pool{
		Shield:           shield,
		Premium:          deposit,
		Sponsor:          sponsor,
		Active:           true,
		TotalCollateral:  shield,
		EndTime:          endTime,
		StartBlockHeight: startBlockHeight,
		EndBlockHeight:   endBlockHeight,
		PoolID:           id,
	}
}

type Collateral struct {
	PoolID            uint64             `json:"pool_id" yaml:"pool_id"`
	Provider          sdk.AccAddress     `json:"provider" yaml:"provider"`
	Amount            sdk.Coins          `json:"amount" yaml:"amount"`
	Withdrawal        sdk.Coins          `json:"withdrawal" yaml:"withdrawal"`
	LockedCollaterals []LockedCollateral `json:"locked_collaterals" yaml:"locked_collaterals"`
}

func NewCollateral(pool Pool, provider sdk.AccAddress, amount sdk.Coins) Collateral {
	return Collateral{
		PoolID:   pool.PoolID,
		Provider: provider,
		Amount:   amount,
	}
}

type PendingPayout struct {
	Amount sdk.Dec
	ToAddr string
}

type PendingPayouts []PendingPayout

func NewPendingPayouts(amount sdk.Dec, to string) PendingPayout {
	return PendingPayout{
		Amount: amount,
		ToAddr: to,
	}
}

// Provider tracks A or C's total delegation, total collateral,
// and rewards.
type Provider struct {
	Address          sdk.AccAddress `json:"address" yaml:"address"`
	DelegationBonded sdk.Coins      `json:"delegation_bonded" yaml:"delegation_bonded"`
	Collateral       sdk.Coins      `json:"collateral" yaml:"collateral"`
	TotalLocked      sdk.Coins      `json:"total_locked" yaml:"total_locked"`
	Available        sdk.Int        `json:"available" yaml:"available"`
	Rewards          MixedDecCoins  `json:"rewards" yaml:"rewards"`
}

func NewProvider(addr sdk.AccAddress) Provider {
	return Provider{
		Address:          addr,
		DelegationBonded: sdk.Coins{},
		Collateral:       sdk.Coins{},
		TotalLocked:      sdk.Coins{},
	}
}

type Purchase struct {
	TxHash             []byte         `json:"tx_hash" yaml:"tx_hash"`
	PoolID             uint64         `json:"pool_id" yaml:"pool_id"`
	Shield             sdk.Coins      `json:"shield" yaml:"shield"`
	StartBlockHeight   int64          `json:"start_block_height" yaml:"start_block_height"`
	ProtectionEndTime  time.Time      `json:"protection_end_time" yaml:"protection_end_time"`
	ClaimPeriodEndTime time.Time      `json:"claim_period_end_time" yaml:"claim_period_end_time"`
	Description        string         `json:"description" yaml:"description"`
	Purchaser          sdk.AccAddress `json:"purchaser" yaml:"purchaser"`
}

func NewPurchase(
	txhash []byte, poolID uint64, shield sdk.Coins, startBlockHeight int64, protectionEndTime, claimPeriodEndTime time.Time,
	description string, purchaser sdk.AccAddress) Purchase {
	return Purchase{
		TxHash:             txhash,
		PoolID:             poolID,
		Shield:             shield,
		StartBlockHeight:   startBlockHeight,
		ProtectionEndTime:  protectionEndTime,
		ClaimPeriodEndTime: claimPeriodEndTime,
		Description:        description,
		Purchaser:          purchaser,
	}
}

// Withdrawal stores an ongoing withdrawal of pool collateral.
type Withdrawal struct {
	PoolID  uint64         `json:"pool_id" yaml:"pool_id"`
	Address sdk.AccAddress `json:"address" yaml:"address"`
	Amount  sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewWithdrawal(poolID uint64, addr sdk.AccAddress, amount sdk.Coins) Withdrawal {
	return Withdrawal{
		PoolID:  poolID,
		Address: addr,
		Amount:  amount,
	}
}