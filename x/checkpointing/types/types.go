package types

import (
	"bytes"
	"github.com/babylonchain/babylon/crypto/bls12381"
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/boljen/go-bitmap"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LastCommitHash []byte

type BlsSigHash []byte

type RawCkptHash []byte

func NewCheckpoint(epochNum uint64, lch LastCommitHash) *RawCheckpoint {
	return &RawCheckpoint{
		EpochNum:       epochNum,
		LastCommitHash: lch,
		Bitmap:         nil,
		BlsMultiSig:    nil,
	}
}

func NewCheckpointWithMeta(ckpt *RawCheckpoint, status CheckpointStatus) *RawCheckpointWithMeta {
	return &RawCheckpointWithMeta{
		Ckpt:   ckpt,
		Status: status,
	}
}

// Accumulate does the following things
// 1. aggregates the BLS signature
// 2. aggregates the BLS public key
// 3. updates Bitmap
// 4. accumulates voting power
// it returns True if the checkpoint is updated
func (cm *RawCheckpointWithMeta) Accumulate(
	vals epochingtypes.ValidatorSet,
	signerAddr sdk.ValAddress,
	signerBlsKey bls12381.PublicKey,
	sig bls12381.Signature,
	totalPower int64) (bool, error) {

	// the checkpoint should be accumulating
	if cm.Status != Accumulating {
		return false, ErrCkptNotAccumulating
	}

	// get validator and its index
	val, index, err := vals.FindValidatorWithIndex(signerAddr)
	if err != nil {
		return false, err
	}

	// return an error if the validator has already voted
	if bitmap.Get(cm.Ckpt.Bitmap, index) {
		return false, ErrCkptAlreadyVoted
	}

	// aggregate BLS sig
	aggSig, err := bls12381.AggrSig(*cm.Ckpt.BlsMultiSig, sig)
	if err != nil {
		return false, err
	}
	cm.Ckpt.BlsMultiSig = &aggSig

	// aggregate BLS public key
	aggPK, err := bls12381.AggrPK(*cm.BlsAggrPk, signerBlsKey)
	if err != nil {
		return false, err
	}
	cm.BlsAggrPk = &aggPK

	// update bitmap
	bitmap.Set(cm.Ckpt.Bitmap, index, true)

	// accumulate voting power and update status when the threshold is reached
	cm.PowerSum += uint64(val.Power)
	if int64(cm.PowerSum) > totalPower/3 {
		cm.Status = Sealed
	}

	return true, nil
}

func BytesToRawCkpt(cdc codec.BinaryCodec, bz []byte) (*RawCheckpoint, error) {
	ckpt := new(RawCheckpoint)
	err := cdc.Unmarshal(bz, ckpt)
	return ckpt, err
}

func CkptWithMetaToBytes(cdc codec.BinaryCodec, ckptWithMeta *RawCheckpointWithMeta) []byte {
	return cdc.MustMarshal(ckptWithMeta)
}

func BytesToCkptWithMeta(cdc codec.BinaryCodec, bz []byte) (*RawCheckpointWithMeta, error) {
	ckptWithMeta := new(RawCheckpointWithMeta)
	err := cdc.Unmarshal(bz, ckptWithMeta)
	return ckptWithMeta, err
}

func (m RawCkptHash) Equals(h RawCkptHash) bool {
	if bytes.Compare(m.Bytes(), h.Bytes()) == 0 {
		return true
	}
	return false
}
