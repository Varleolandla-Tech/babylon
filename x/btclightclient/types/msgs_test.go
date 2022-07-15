package types_test

import (
	"bytes"
	"github.com/babylonchain/babylon/testutil/datagen"
	bbl "github.com/babylonchain/babylon/types"
	"github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/btcsuite/btcd/blockchain"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
	"math/rand"
	"testing"
)

func FuzzMsgInsertHeader(f *testing.F) {
	addressBytes := []byte("from________________")
	defaultHeader := bbl.GetBaseBTCHeaderBytes()
	defaultBtcdHeader := defaultHeader.ToBlockHeader()

	// Maximum btc difficulty possible
	// Use it to set the difficulty bits of blocks as well as the upper PoW limit
	// since the block hash needs to be below that
	// This is the maximum allowed given the 2^23-1 precision
	maxDifficulty, _ := new(big.Int).SetString("ffff000000000000000000000000000000000000000000000000000000000000", 16)

	f.Add(
		addressBytes,
		defaultBtcdHeader.Version,
		defaultBtcdHeader.Bits,
		defaultBtcdHeader.Nonce,
		defaultBtcdHeader.Timestamp.Unix(),
		defaultBtcdHeader.PrevBlock.String(),
		defaultBtcdHeader.MerkleRoot.String(),
		int64(17))

	f.Fuzz(func(t *testing.T, addressBytes []byte, version int32, bits uint32, nonce uint32,
		timeInt int64, prevBlockStr string, merkleRootStr string, seed int64) {

		rand.Seed(seed)
		errorKind := 0

		// Get the btcd header based on the provided data
		btcdHeader := datagen.GenRandomBtcdHeader(version, bits, nonce, timeInt, prevBlockStr, merkleRootStr)
		// If the header hex is the same as the default one, then this is the seed input
		headerBytes := bbl.NewBTCHeaderBytesFromBlockHeader(btcdHeader)
		headerHex := headerBytes.MarshalHex()
		seedInput := defaultHeader.MarshalHex() == headerHex

		// Make the address have a proper size
		if len(addressBytes) == 0 || len(addressBytes) >= 256 {
			addressBytes = datagen.GenRandomByteArray(1 + uint64(rand.Intn(255)))
		}

		// Get the signer structure
		var signer sdk.AccAddress
		signer.Unmarshal(addressBytes)

		// Perform modifications on the btcd header if it is not part of the seed input
		if !seedInput {
			errorKind = rand.Intn(3)
			switch errorKind {
			case 0:
				// Valid input
				// Set the work bits to the pow limit
				bits = blockchain.BigToCompact(maxDifficulty)
			case 1:
				// Zero PoW
				bits = blockchain.BigToCompact(big.NewInt(0))
			case 2:
				// Negative PoW
				bits = blockchain.BigToCompact(big.NewInt(-1))
			default:
				bits = blockchain.BigToCompact(maxDifficulty)
			}
		}
		// Generate a header with the provided modifications
		newBtcdHeader := datagen.GenRandomBtcdHeader(version, bits, nonce, timeInt, prevBlockStr, merkleRootStr)
		newHeader := bbl.NewBTCHeaderBytesFromBlockHeader(newBtcdHeader)
		newHeaderHex := newHeader.MarshalHex()

		// Check whether the hash is still bigger than the maximum allowed
		// This happens because even though we pass a series of "f"s as an input
		// the maximum that the bits field can contain is 2^23-1, meaning
		// that there is still space for block hashes that are less than that
		newHeaderHash := newBtcdHeader.BlockHash()
		hashNum := blockchain.HashToBig(&newHeaderHash)
		if hashNum.Cmp(maxDifficulty) > 0 {
			t.Skip()
		}

		// Check the message creation
		msgInsertHeader, err := types.NewMsgInsertHeader(signer, newHeaderHex)
		if err != nil {
			t.Errorf("Valid parameters led to error")
		}
		if msgInsertHeader == nil {
			t.Errorf("nil returned")
		}
		if msgInsertHeader.Header == nil {
			t.Errorf("nil header")
		}
		if bytes.Compare(newHeader, *(msgInsertHeader.Header)) != 0 {
			t.Errorf("Expected header bytes %s got %s", newHeader, *(msgInsertHeader.Header))
		}

		// Validate the message
		err = msgInsertHeader.ValidateHeader(maxDifficulty)
		if err != nil && errorKind == 0 {
			t.Errorf("Valid message %s failed with %s", headerHex, err)
		}
		if err == nil && errorKind != 0 {
			t.Errorf("Invalid message did not fail")
		}
	})

}
