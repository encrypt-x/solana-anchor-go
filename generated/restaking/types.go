// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type BatchWithdrawal struct {
	BatchId                    uint64
	NumWithdrawalRequests      uint64
	ReceiptTokenToProcess      uint64
	ReceiptTokenBeingProcessed uint64
	ReceiptTokenProcessed      uint64
	SolReserved                uint64
	ProcessingStartedAt        *int64 `bin:"optional"`
}

func (obj BatchWithdrawal) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BatchId` param:
	err = encoder.Encode(obj.BatchId)
	if err != nil {
		return err
	}
	// Serialize `NumWithdrawalRequests` param:
	err = encoder.Encode(obj.NumWithdrawalRequests)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenToProcess` param:
	err = encoder.Encode(obj.ReceiptTokenToProcess)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenBeingProcessed` param:
	err = encoder.Encode(obj.ReceiptTokenBeingProcessed)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenProcessed` param:
	err = encoder.Encode(obj.ReceiptTokenProcessed)
	if err != nil {
		return err
	}
	// Serialize `SolReserved` param:
	err = encoder.Encode(obj.SolReserved)
	if err != nil {
		return err
	}
	// Serialize `ProcessingStartedAt` param (optional):
	{
		if obj.ProcessingStartedAt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ProcessingStartedAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *BatchWithdrawal) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BatchId`:
	err = decoder.Decode(&obj.BatchId)
	if err != nil {
		return err
	}
	// Deserialize `NumWithdrawalRequests`:
	err = decoder.Decode(&obj.NumWithdrawalRequests)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenToProcess`:
	err = decoder.Decode(&obj.ReceiptTokenToProcess)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenBeingProcessed`:
	err = decoder.Decode(&obj.ReceiptTokenBeingProcessed)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenProcessed`:
	err = decoder.Decode(&obj.ReceiptTokenProcessed)
	if err != nil {
		return err
	}
	// Deserialize `SolReserved`:
	err = decoder.Decode(&obj.SolReserved)
	if err != nil {
		return err
	}
	// Deserialize `ProcessingStartedAt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ProcessingStartedAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Fund struct {
	DataVersion      uint8
	Bump             uint8
	ReceiptTokenMint ag_solanago.PublicKey
	SupportedTokens  []TokenInfo
	SolAmountIn      uint64
	WithdrawalStatus WithdrawalStatus
}

func (obj Fund) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `SupportedTokens` param:
	err = encoder.Encode(obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Serialize `SolAmountIn` param:
	err = encoder.Encode(obj.SolAmountIn)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalStatus` param:
	err = encoder.Encode(obj.WithdrawalStatus)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Fund) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `SupportedTokens`:
	err = decoder.Decode(&obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Deserialize `SolAmountIn`:
	err = decoder.Decode(&obj.SolAmountIn)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalStatus`:
	err = decoder.Decode(&obj.WithdrawalStatus)
	if err != nil {
		return err
	}
	return nil
}

type FundInfo struct {
	LrtMint              ag_solanago.PublicKey
	SupportedTokens      []TokenInfo
	SolAmountIn          uint64
	SolReservedAmount    uint64
	SolWithdrawalFeeRate float32
	SolWithdrawalEnabled bool
}

func (obj FundInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `SupportedTokens` param:
	err = encoder.Encode(obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Serialize `SolAmountIn` param:
	err = encoder.Encode(obj.SolAmountIn)
	if err != nil {
		return err
	}
	// Serialize `SolReservedAmount` param:
	err = encoder.Encode(obj.SolReservedAmount)
	if err != nil {
		return err
	}
	// Serialize `SolWithdrawalFeeRate` param:
	err = encoder.Encode(obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	// Serialize `SolWithdrawalEnabled` param:
	err = encoder.Encode(obj.SolWithdrawalEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `SupportedTokens`:
	err = decoder.Decode(&obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Deserialize `SolAmountIn`:
	err = decoder.Decode(&obj.SolAmountIn)
	if err != nil {
		return err
	}
	// Deserialize `SolReservedAmount`:
	err = decoder.Decode(&obj.SolReservedAmount)
	if err != nil {
		return err
	}
	// Deserialize `SolWithdrawalFeeRate`:
	err = decoder.Decode(&obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `SolWithdrawalEnabled`:
	err = decoder.Decode(&obj.SolWithdrawalEnabled)
	if err != nil {
		return err
	}
	return nil
}

type FundPriceUpdated struct {
	LrtMint  ag_solanago.PublicKey
	LrtPrice uint64
	FundInfo FundInfo
}

func (obj FundPriceUpdated) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `LrtPrice` param:
	err = encoder.Encode(obj.LrtPrice)
	if err != nil {
		return err
	}
	// Serialize `FundInfo` param:
	err = encoder.Encode(obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundPriceUpdated) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `LrtPrice`:
	err = decoder.Decode(&obj.LrtPrice)
	if err != nil {
		return err
	}
	// Deserialize `FundInfo`:
	err = decoder.Decode(&obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

type FundSOLDeposited struct {
	User                        ag_solanago.PublicKey
	UserLrtAccount              ag_solanago.PublicKey
	UserReceipt                 UserReceipt
	SolDepositAmount            uint64
	SolAmountInFund             uint64
	MintedLrtMint               ag_solanago.PublicKey
	MintedLrtAmount             uint64
	LrtPrice                    uint64
	LrtAmountInUserLrtAccount   uint64
	WalletProvider              *string  `bin:"optional"`
	FpointAccrualRateMultiplier *float32 `bin:"optional"`
	FundInfo                    FundInfo
}

func (obj FundSOLDeposited) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `UserLrtAccount` param:
	err = encoder.Encode(obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `UserReceipt` param:
	err = encoder.Encode(obj.UserReceipt)
	if err != nil {
		return err
	}
	// Serialize `SolDepositAmount` param:
	err = encoder.Encode(obj.SolDepositAmount)
	if err != nil {
		return err
	}
	// Serialize `SolAmountInFund` param:
	err = encoder.Encode(obj.SolAmountInFund)
	if err != nil {
		return err
	}
	// Serialize `MintedLrtMint` param:
	err = encoder.Encode(obj.MintedLrtMint)
	if err != nil {
		return err
	}
	// Serialize `MintedLrtAmount` param:
	err = encoder.Encode(obj.MintedLrtAmount)
	if err != nil {
		return err
	}
	// Serialize `LrtPrice` param:
	err = encoder.Encode(obj.LrtPrice)
	if err != nil {
		return err
	}
	// Serialize `LrtAmountInUserLrtAccount` param:
	err = encoder.Encode(obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `WalletProvider` param (optional):
	{
		if obj.WalletProvider == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.WalletProvider)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FpointAccrualRateMultiplier` param (optional):
	{
		if obj.FpointAccrualRateMultiplier == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FpointAccrualRateMultiplier)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FundInfo` param:
	err = encoder.Encode(obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundSOLDeposited) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `UserLrtAccount`:
	err = decoder.Decode(&obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `UserReceipt`:
	err = decoder.Decode(&obj.UserReceipt)
	if err != nil {
		return err
	}
	// Deserialize `SolDepositAmount`:
	err = decoder.Decode(&obj.SolDepositAmount)
	if err != nil {
		return err
	}
	// Deserialize `SolAmountInFund`:
	err = decoder.Decode(&obj.SolAmountInFund)
	if err != nil {
		return err
	}
	// Deserialize `MintedLrtMint`:
	err = decoder.Decode(&obj.MintedLrtMint)
	if err != nil {
		return err
	}
	// Deserialize `MintedLrtAmount`:
	err = decoder.Decode(&obj.MintedLrtAmount)
	if err != nil {
		return err
	}
	// Deserialize `LrtPrice`:
	err = decoder.Decode(&obj.LrtPrice)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmountInUserLrtAccount`:
	err = decoder.Decode(&obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `WalletProvider` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.WalletProvider)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FpointAccrualRateMultiplier` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FpointAccrualRateMultiplier)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FundInfo`:
	err = decoder.Decode(&obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

type FundSOLWithdrawn struct {
	User              ag_solanago.PublicKey
	UserReceipt       UserReceipt
	RequestId         uint64
	LrtMint           ag_solanago.PublicKey
	LrtAmount         uint64
	LrtPrice          uint64
	SolWithdrawAmount uint64
	SolFeeAmount      uint64
	FundInfo          FundInfo
}

func (obj FundSOLWithdrawn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `UserReceipt` param:
	err = encoder.Encode(obj.UserReceipt)
	if err != nil {
		return err
	}
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `LrtAmount` param:
	err = encoder.Encode(obj.LrtAmount)
	if err != nil {
		return err
	}
	// Serialize `LrtPrice` param:
	err = encoder.Encode(obj.LrtPrice)
	if err != nil {
		return err
	}
	// Serialize `SolWithdrawAmount` param:
	err = encoder.Encode(obj.SolWithdrawAmount)
	if err != nil {
		return err
	}
	// Serialize `SolFeeAmount` param:
	err = encoder.Encode(obj.SolFeeAmount)
	if err != nil {
		return err
	}
	// Serialize `FundInfo` param:
	err = encoder.Encode(obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundSOLWithdrawn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `UserReceipt`:
	err = decoder.Decode(&obj.UserReceipt)
	if err != nil {
		return err
	}
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmount`:
	err = decoder.Decode(&obj.LrtAmount)
	if err != nil {
		return err
	}
	// Deserialize `LrtPrice`:
	err = decoder.Decode(&obj.LrtPrice)
	if err != nil {
		return err
	}
	// Deserialize `SolWithdrawAmount`:
	err = decoder.Decode(&obj.SolWithdrawAmount)
	if err != nil {
		return err
	}
	// Deserialize `SolFeeAmount`:
	err = decoder.Decode(&obj.SolFeeAmount)
	if err != nil {
		return err
	}
	// Deserialize `FundInfo`:
	err = decoder.Decode(&obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

type FundTokenAuthority struct {
	DataVersion      uint8
	Bump             uint8
	ReceiptTokenMint ag_solanago.PublicKey
}

func (obj FundTokenAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundTokenAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	return nil
}

type FundTokenDeposited struct {
	User                        ag_solanago.PublicKey
	UserLrtAccount              ag_solanago.PublicKey
	UserReceipt                 UserReceipt
	DepositedTokenMint          ag_solanago.PublicKey
	DepositedTokenUserAccount   ag_solanago.PublicKey
	TokenDepositAmount          uint64
	TokenAmountInFund           uint64
	MintedLrtMint               ag_solanago.PublicKey
	MintedLrtAmount             uint64
	LrtPrice                    uint64
	LrtAmountInUserLrtAccount   uint64
	WalletProvider              *string  `bin:"optional"`
	FpointAccrualRateMultiplier *float32 `bin:"optional"`
	FundInfo                    FundInfo
}

func (obj FundTokenDeposited) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `UserLrtAccount` param:
	err = encoder.Encode(obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `UserReceipt` param:
	err = encoder.Encode(obj.UserReceipt)
	if err != nil {
		return err
	}
	// Serialize `DepositedTokenMint` param:
	err = encoder.Encode(obj.DepositedTokenMint)
	if err != nil {
		return err
	}
	// Serialize `DepositedTokenUserAccount` param:
	err = encoder.Encode(obj.DepositedTokenUserAccount)
	if err != nil {
		return err
	}
	// Serialize `TokenDepositAmount` param:
	err = encoder.Encode(obj.TokenDepositAmount)
	if err != nil {
		return err
	}
	// Serialize `TokenAmountInFund` param:
	err = encoder.Encode(obj.TokenAmountInFund)
	if err != nil {
		return err
	}
	// Serialize `MintedLrtMint` param:
	err = encoder.Encode(obj.MintedLrtMint)
	if err != nil {
		return err
	}
	// Serialize `MintedLrtAmount` param:
	err = encoder.Encode(obj.MintedLrtAmount)
	if err != nil {
		return err
	}
	// Serialize `LrtPrice` param:
	err = encoder.Encode(obj.LrtPrice)
	if err != nil {
		return err
	}
	// Serialize `LrtAmountInUserLrtAccount` param:
	err = encoder.Encode(obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `WalletProvider` param (optional):
	{
		if obj.WalletProvider == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.WalletProvider)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FpointAccrualRateMultiplier` param (optional):
	{
		if obj.FpointAccrualRateMultiplier == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FpointAccrualRateMultiplier)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FundInfo` param:
	err = encoder.Encode(obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundTokenDeposited) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `UserLrtAccount`:
	err = decoder.Decode(&obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `UserReceipt`:
	err = decoder.Decode(&obj.UserReceipt)
	if err != nil {
		return err
	}
	// Deserialize `DepositedTokenMint`:
	err = decoder.Decode(&obj.DepositedTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `DepositedTokenUserAccount`:
	err = decoder.Decode(&obj.DepositedTokenUserAccount)
	if err != nil {
		return err
	}
	// Deserialize `TokenDepositAmount`:
	err = decoder.Decode(&obj.TokenDepositAmount)
	if err != nil {
		return err
	}
	// Deserialize `TokenAmountInFund`:
	err = decoder.Decode(&obj.TokenAmountInFund)
	if err != nil {
		return err
	}
	// Deserialize `MintedLrtMint`:
	err = decoder.Decode(&obj.MintedLrtMint)
	if err != nil {
		return err
	}
	// Deserialize `MintedLrtAmount`:
	err = decoder.Decode(&obj.MintedLrtAmount)
	if err != nil {
		return err
	}
	// Deserialize `LrtPrice`:
	err = decoder.Decode(&obj.LrtPrice)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmountInUserLrtAccount`:
	err = decoder.Decode(&obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `WalletProvider` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.WalletProvider)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FpointAccrualRateMultiplier` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FpointAccrualRateMultiplier)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FundInfo`:
	err = decoder.Decode(&obj.FundInfo)
	if err != nil {
		return err
	}
	return nil
}

type FundWithdrawalRequestCanceled struct {
	User                      ag_solanago.PublicKey
	UserLrtAccount            ag_solanago.PublicKey
	UserReceipt               UserReceipt
	RequestId                 uint64
	LrtMint                   ag_solanago.PublicKey
	LrtRequestedAmount        uint64
	LrtAmountInUserLrtAccount uint64
}

func (obj FundWithdrawalRequestCanceled) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `UserLrtAccount` param:
	err = encoder.Encode(obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `UserReceipt` param:
	err = encoder.Encode(obj.UserReceipt)
	if err != nil {
		return err
	}
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `LrtRequestedAmount` param:
	err = encoder.Encode(obj.LrtRequestedAmount)
	if err != nil {
		return err
	}
	// Serialize `LrtAmountInUserLrtAccount` param:
	err = encoder.Encode(obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundWithdrawalRequestCanceled) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `UserLrtAccount`:
	err = decoder.Decode(&obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `UserReceipt`:
	err = decoder.Decode(&obj.UserReceipt)
	if err != nil {
		return err
	}
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `LrtRequestedAmount`:
	err = decoder.Decode(&obj.LrtRequestedAmount)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmountInUserLrtAccount`:
	err = decoder.Decode(&obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	return nil
}

type FundWithdrawalRequested struct {
	User                      ag_solanago.PublicKey
	UserLrtAccount            ag_solanago.PublicKey
	UserReceipt               UserReceipt
	RequestId                 uint64
	LrtMint                   ag_solanago.PublicKey
	LrtRequestedAmount        uint64
	LrtAmountInUserLrtAccount uint64
}

func (obj FundWithdrawalRequested) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `UserLrtAccount` param:
	err = encoder.Encode(obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `UserReceipt` param:
	err = encoder.Encode(obj.UserReceipt)
	if err != nil {
		return err
	}
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `LrtRequestedAmount` param:
	err = encoder.Encode(obj.LrtRequestedAmount)
	if err != nil {
		return err
	}
	// Serialize `LrtAmountInUserLrtAccount` param:
	err = encoder.Encode(obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundWithdrawalRequested) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `UserLrtAccount`:
	err = decoder.Decode(&obj.UserLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `UserReceipt`:
	err = decoder.Decode(&obj.UserReceipt)
	if err != nil {
		return err
	}
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `LrtRequestedAmount`:
	err = decoder.Decode(&obj.LrtRequestedAmount)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmountInUserLrtAccount`:
	err = decoder.Decode(&obj.LrtAmountInUserLrtAccount)
	if err != nil {
		return err
	}
	return nil
}

type Metadata struct {
	WalletProvider              string
	FpointAccrualRateMultiplier float32
}

func (obj Metadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `WalletProvider` param:
	err = encoder.Encode(obj.WalletProvider)
	if err != nil {
		return err
	}
	// Serialize `FpointAccrualRateMultiplier` param:
	err = encoder.Encode(obj.FpointAccrualRateMultiplier)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Metadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `WalletProvider`:
	err = decoder.Decode(&obj.WalletProvider)
	if err != nil {
		return err
	}
	// Deserialize `FpointAccrualRateMultiplier`:
	err = decoder.Decode(&obj.FpointAccrualRateMultiplier)
	if err != nil {
		return err
	}
	return nil
}

type ReservedFund struct {
	NumCompletedWithdrawalRequests uint64
	TotalReceiptTokenProcessed     ag_binary.Uint128
	TotalSolReserved               ag_binary.Uint128
	SolRemaining                   uint64
}

func (obj ReservedFund) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NumCompletedWithdrawalRequests` param:
	err = encoder.Encode(obj.NumCompletedWithdrawalRequests)
	if err != nil {
		return err
	}
	// Serialize `TotalReceiptTokenProcessed` param:
	err = encoder.Encode(obj.TotalReceiptTokenProcessed)
	if err != nil {
		return err
	}
	// Serialize `TotalSolReserved` param:
	err = encoder.Encode(obj.TotalSolReserved)
	if err != nil {
		return err
	}
	// Serialize `SolRemaining` param:
	err = encoder.Encode(obj.SolRemaining)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservedFund) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NumCompletedWithdrawalRequests`:
	err = decoder.Decode(&obj.NumCompletedWithdrawalRequests)
	if err != nil {
		return err
	}
	// Deserialize `TotalReceiptTokenProcessed`:
	err = decoder.Decode(&obj.TotalReceiptTokenProcessed)
	if err != nil {
		return err
	}
	// Deserialize `TotalSolReserved`:
	err = decoder.Decode(&obj.TotalSolReserved)
	if err != nil {
		return err
	}
	// Deserialize `SolRemaining`:
	err = decoder.Decode(&obj.SolRemaining)
	if err != nil {
		return err
	}
	return nil
}

type TokenInfo struct {
	Address            ag_solanago.PublicKey
	TokenDecimals      uint8
	TokenCap           uint64
	TokenAmountIn      uint64
	TokenPrice         uint64
	TokenPricingSource TokenPricingSource
}

func (obj TokenInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `TokenDecimals` param:
	err = encoder.Encode(obj.TokenDecimals)
	if err != nil {
		return err
	}
	// Serialize `TokenCap` param:
	err = encoder.Encode(obj.TokenCap)
	if err != nil {
		return err
	}
	// Serialize `TokenAmountIn` param:
	err = encoder.Encode(obj.TokenAmountIn)
	if err != nil {
		return err
	}
	// Serialize `TokenPrice` param:
	err = encoder.Encode(obj.TokenPrice)
	if err != nil {
		return err
	}
	// Serialize `TokenPricingSource` param:
	{
		tmp := tokenPricingSourceContainer{}
		switch realvalue := obj.TokenPricingSource.(type) {
		case *TokenPricingSourceSPLStakePoolTuple:
			tmp.Enum = 0
			tmp.SPLStakePool = *realvalue
		case *TokenPricingSourceMarinadeStakePoolTuple:
			tmp.Enum = 1
			tmp.MarinadeStakePool = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *TokenInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `TokenDecimals`:
	err = decoder.Decode(&obj.TokenDecimals)
	if err != nil {
		return err
	}
	// Deserialize `TokenCap`:
	err = decoder.Decode(&obj.TokenCap)
	if err != nil {
		return err
	}
	// Deserialize `TokenAmountIn`:
	err = decoder.Decode(&obj.TokenAmountIn)
	if err != nil {
		return err
	}
	// Deserialize `TokenPrice`:
	err = decoder.Decode(&obj.TokenPrice)
	if err != nil {
		return err
	}
	// Deserialize `TokenPricingSource`:
	{
		tmp := new(tokenPricingSourceContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.TokenPricingSource = &tmp.SPLStakePool
		case 1:
			obj.TokenPricingSource = &tmp.MarinadeStakePool
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

type TokenLRTTransferred struct {
	LrtMint                    ag_solanago.PublicKey
	LrtAmount                  uint64
	SourceLrtAccount           ag_solanago.PublicKey
	SourceLrtAccountOwner      ag_solanago.PublicKey
	DestinationLrtAccount      ag_solanago.PublicKey
	DestinationLrtAccountOnwer ag_solanago.PublicKey
}

func (obj TokenLRTTransferred) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LrtMint` param:
	err = encoder.Encode(obj.LrtMint)
	if err != nil {
		return err
	}
	// Serialize `LrtAmount` param:
	err = encoder.Encode(obj.LrtAmount)
	if err != nil {
		return err
	}
	// Serialize `SourceLrtAccount` param:
	err = encoder.Encode(obj.SourceLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `SourceLrtAccountOwner` param:
	err = encoder.Encode(obj.SourceLrtAccountOwner)
	if err != nil {
		return err
	}
	// Serialize `DestinationLrtAccount` param:
	err = encoder.Encode(obj.DestinationLrtAccount)
	if err != nil {
		return err
	}
	// Serialize `DestinationLrtAccountOnwer` param:
	err = encoder.Encode(obj.DestinationLrtAccountOnwer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenLRTTransferred) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LrtMint`:
	err = decoder.Decode(&obj.LrtMint)
	if err != nil {
		return err
	}
	// Deserialize `LrtAmount`:
	err = decoder.Decode(&obj.LrtAmount)
	if err != nil {
		return err
	}
	// Deserialize `SourceLrtAccount`:
	err = decoder.Decode(&obj.SourceLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `SourceLrtAccountOwner`:
	err = decoder.Decode(&obj.SourceLrtAccountOwner)
	if err != nil {
		return err
	}
	// Deserialize `DestinationLrtAccount`:
	err = decoder.Decode(&obj.DestinationLrtAccount)
	if err != nil {
		return err
	}
	// Deserialize `DestinationLrtAccountOnwer`:
	err = decoder.Decode(&obj.DestinationLrtAccountOnwer)
	if err != nil {
		return err
	}
	return nil
}

type TokenPricingSource interface {
	isTokenPricingSource()
}

type tokenPricingSourceContainer struct {
	Enum              ag_binary.BorshEnum `borsh_enum:"true"`
	SPLStakePool      TokenPricingSourceSPLStakePoolTuple
	MarinadeStakePool TokenPricingSourceMarinadeStakePoolTuple
}

type TokenPricingSourceSPLStakePoolTuple struct {
	Address ag_solanago.PublicKey
}

func (obj TokenPricingSourceSPLStakePoolTuple) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenPricingSourceSPLStakePoolTuple) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	return nil
}

func (_ TokenPricingSourceSPLStakePoolTuple) isTokenPricingSource() {}

type TokenPricingSourceMarinadeStakePoolTuple struct {
	Address ag_solanago.PublicKey
}

func (obj TokenPricingSourceMarinadeStakePoolTuple) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenPricingSourceMarinadeStakePoolTuple) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	return nil
}

func (_ TokenPricingSourceMarinadeStakePoolTuple) isTokenPricingSource() {}

type UserReceipt struct {
	DataVersion        uint8
	Bump               uint8
	User               ag_solanago.PublicKey
	ReceiptTokenMint   ag_solanago.PublicKey
	WithdrawalRequests []WithdrawalRequest
}

func (obj UserReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalRequests` param:
	err = encoder.Encode(obj.WithdrawalRequests)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UserReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalRequests`:
	err = decoder.Decode(&obj.WithdrawalRequests)
	if err != nil {
		return err
	}
	return nil
}

type WithdrawalRequest struct {
	BatchId            uint64
	RequestId          uint64
	ReceiptTokenAmount uint64
	CreatedAt          int64
}

func (obj WithdrawalRequest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BatchId` param:
	err = encoder.Encode(obj.BatchId)
	if err != nil {
		return err
	}
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenAmount` param:
	err = encoder.Encode(obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WithdrawalRequest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BatchId`:
	err = decoder.Decode(&obj.BatchId)
	if err != nil {
		return err
	}
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenAmount`:
	err = decoder.Decode(&obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

type WithdrawalStatus struct {
	NextBatchId                      uint64
	NextRequestId                    uint64
	NumWithdrawalRequestsInProgress  uint64
	LastCompletedBatchId             uint64
	LastBatchProcessingStartedAt     *int64 `bin:"optional"`
	LastBatchProcessingCompletedAt   *int64 `bin:"optional"`
	SolWithdrawalFeeRate             uint16
	WithdrawalEnabledFlag            bool
	BatchProcessingThresholdAmount   uint64
	BatchProcessingThresholdDuration int64
	PendingBatchWithdrawal           BatchWithdrawal
	BatchWithdrawalsInProgress       []BatchWithdrawal
	ReservedFund                     ReservedFund
}

func (obj WithdrawalStatus) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NextBatchId` param:
	err = encoder.Encode(obj.NextBatchId)
	if err != nil {
		return err
	}
	// Serialize `NextRequestId` param:
	err = encoder.Encode(obj.NextRequestId)
	if err != nil {
		return err
	}
	// Serialize `NumWithdrawalRequestsInProgress` param:
	err = encoder.Encode(obj.NumWithdrawalRequestsInProgress)
	if err != nil {
		return err
	}
	// Serialize `LastCompletedBatchId` param:
	err = encoder.Encode(obj.LastCompletedBatchId)
	if err != nil {
		return err
	}
	// Serialize `LastBatchProcessingStartedAt` param (optional):
	{
		if obj.LastBatchProcessingStartedAt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LastBatchProcessingStartedAt)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `LastBatchProcessingCompletedAt` param (optional):
	{
		if obj.LastBatchProcessingCompletedAt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LastBatchProcessingCompletedAt)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SolWithdrawalFeeRate` param:
	err = encoder.Encode(obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalEnabledFlag` param:
	err = encoder.Encode(obj.WithdrawalEnabledFlag)
	if err != nil {
		return err
	}
	// Serialize `BatchProcessingThresholdAmount` param:
	err = encoder.Encode(obj.BatchProcessingThresholdAmount)
	if err != nil {
		return err
	}
	// Serialize `BatchProcessingThresholdDuration` param:
	err = encoder.Encode(obj.BatchProcessingThresholdDuration)
	if err != nil {
		return err
	}
	// Serialize `PendingBatchWithdrawal` param:
	err = encoder.Encode(obj.PendingBatchWithdrawal)
	if err != nil {
		return err
	}
	// Serialize `BatchWithdrawalsInProgress` param:
	err = encoder.Encode(obj.BatchWithdrawalsInProgress)
	if err != nil {
		return err
	}
	// Serialize `ReservedFund` param:
	err = encoder.Encode(obj.ReservedFund)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WithdrawalStatus) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NextBatchId`:
	err = decoder.Decode(&obj.NextBatchId)
	if err != nil {
		return err
	}
	// Deserialize `NextRequestId`:
	err = decoder.Decode(&obj.NextRequestId)
	if err != nil {
		return err
	}
	// Deserialize `NumWithdrawalRequestsInProgress`:
	err = decoder.Decode(&obj.NumWithdrawalRequestsInProgress)
	if err != nil {
		return err
	}
	// Deserialize `LastCompletedBatchId`:
	err = decoder.Decode(&obj.LastCompletedBatchId)
	if err != nil {
		return err
	}
	// Deserialize `LastBatchProcessingStartedAt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LastBatchProcessingStartedAt)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `LastBatchProcessingCompletedAt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LastBatchProcessingCompletedAt)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SolWithdrawalFeeRate`:
	err = decoder.Decode(&obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalEnabledFlag`:
	err = decoder.Decode(&obj.WithdrawalEnabledFlag)
	if err != nil {
		return err
	}
	// Deserialize `BatchProcessingThresholdAmount`:
	err = decoder.Decode(&obj.BatchProcessingThresholdAmount)
	if err != nil {
		return err
	}
	// Deserialize `BatchProcessingThresholdDuration`:
	err = decoder.Decode(&obj.BatchProcessingThresholdDuration)
	if err != nil {
		return err
	}
	// Deserialize `PendingBatchWithdrawal`:
	err = decoder.Decode(&obj.PendingBatchWithdrawal)
	if err != nil {
		return err
	}
	// Deserialize `BatchWithdrawalsInProgress`:
	err = decoder.Decode(&obj.BatchWithdrawalsInProgress)
	if err != nil {
		return err
	}
	// Deserialize `ReservedFund`:
	err = decoder.Decode(&obj.ReservedFund)
	if err != nil {
		return err
	}
	return nil
}