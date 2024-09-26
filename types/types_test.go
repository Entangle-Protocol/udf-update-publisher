package types

import (
	"math/big"
	"testing"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

// test ECDSA signature serialization
func TestECDSASignature_ToBytes(t *testing.T) {
	r := require.New(t)

	testPkHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privKey, err := keystore.ParseKeyFromHex(testPkHex)
	r.NoError(err)

	msgToSign := []byte("test message")
	msgHash := crypto.Keccak256(msgToSign)
	expectedSigBytes := []byte{ 0xbf, 0x50, 0xb8, 0x99, 0x85, 0xbd, 0x2, 0x4b, 0xd4, 0xf2, 0x5e, 0xa2, 0x1e, 0x72, 0xe0, 0x56, 0xd4, 0x46, 0xdd, 0xe9, 0x8a, 0xac, 0x81, 0xf3, 0x10, 0x3c, 0x9e, 0x46, 0x9e, 0x23, 0x1a, 0xad, 0x51, 0x91, 0x1, 0xf0, 0x2d, 0xaa, 0xbb, 0xd4, 0xaf, 0x51, 0xdf, 0x7f, 0xa2, 0x12, 0xc1, 0x33, 0x88, 0xa9, 0x26, 0x10, 0x84, 0x2b, 0xda, 0xe8, 0x7, 0x26, 0x60, 0x99, 0x36, 0x7c, 0xc6, 0x86, 0x1b }

	// Create signature with test key
	ethSig, err := crypto.Sign(msgHash, privKey)
	r.NoError(err)
	ethSig[64] += 27
	sig := ECDSASignature{
		R: ethcommon.BytesToHash(ethSig[:32]),
		S: ethcommon.BytesToHash(ethSig[32:64]),
		V: ethSig[64],
	}

	sigBytes, err := sig.ToBytes()
	r.NoError(err)

	r.Equal(expectedSigBytes, sigBytes)
}

func TestMerkleRootUpdateMultiple_ToCalldata(t *testing.T) {
	r := require.New(t)

	testCases := []struct {
		desc     string
		merkleRoot string
		signatures []string
		updateData []struct {
			DataKey   string
			Price     string
			Timestamp *big.Int
			MerkleProof []string
		}
		expectedCalldata string
	}{
		{
			desc: "Test with 4 valid updates",
			merkleRoot: "3255dd261db8574a36cb60c3e347133859503ca5d6380819a6c94f6b4d2b581a",
			signatures: []string{
				"41852b85f1f679fab1d58bfd56d63a567589fd1ec1ccfe7d1c3686477556458b291552814a99ee169c81d0bfba017f82123dc10471889671e2576e8367af44d81b",
			},
			updateData: []struct {
				DataKey   string
				Price     string
				Timestamp *big.Int
				MerkleProof []string
			}{
				{
					DataKey:   "USDC/USD",
					Price:     "1010000000000000000",
					Timestamp: big.NewInt(1712156481),
					MerkleProof: []string{
						"2ba524d9865b095ddb19382bbe899287b32cd6a9d499f638a821ccb603bdf9f4",
						"25430d4373ac54a15f2db004d189906a37da3997afd50a1b374800e82270e338",
					},
				},
				{
					DataKey:   "ETH/USD",
					Price:     "3230000000000000000000",
					Timestamp: big.NewInt(1712156481),
					MerkleProof: []string{
						"906f01a562e1a7412219bbf14b3ee5bbd3e698a6f40d4f62906a93b95a3c6a09",
						"4f9996a39fd91a88bafb355c8b32636f0c43004993e2642024530d599667ba78",
					},
				},
				{
					DataKey:   "BTC/USD",
					Price:     "69510000000000000000000",
					Timestamp: big.NewInt(1712156481),
					MerkleProof: []string{
						"1e5b99650c2052bca7105183838063cfdcaab5d7e6ae9e09d8f8e0faa55e60bb",
						"25430d4373ac54a15f2db004d189906a37da3997afd50a1b374800e82270e338",
					},
				},
				{
					DataKey:   "NGL/USD",
					Price:     "830000000000000000",
					Timestamp: big.NewInt(1712156481),
					MerkleProof: []string{
						"62dcaa2d204c8b187d556135fff9fedf2c91be7d366755c801a4d63cd5d63249",
						"4f9996a39fd91a88bafb355c8b32636f0c43004993e2642024530d599667ba78",
					},
				},
			},
			expectedCalldata: "3255dd261db8574a36cb60c3e347133859503ca5d6380819a6c94f6b4d2b581a0141852b85f1f679fab1d58bfd56d63a567589fd1ec1ccfe7d1c3686477556458b291552814a99ee169c81d0bfba017f82123dc10471889671e2576e8367af44d81b04022ba524d9865b095ddb19382bbe899287b32cd6a9d499f638a821ccb603bdf9f425430d4373ac54a15f2db004d189906a37da3997afd50a1b374800e82270e33800000000000000000000000000000000000000000000000000000000660d6f410000000000000000000000000000000000000000000000000e043da617250000555344432f55534400000000000000000000000000000000000000000000000002906f01a562e1a7412219bbf14b3ee5bbd3e698a6f40d4f62906a93b95a3c6a094f9996a39fd91a88bafb355c8b32636f0c43004993e2642024530d599667ba7800000000000000000000000000000000000000000000000000000000660d6f410000000000000000000000000000000000000000000000af19412eb9ffb800004554482f55534400000000000000000000000000000000000000000000000000021e5b99650c2052bca7105183838063cfdcaab5d7e6ae9e09d8f8e0faa55e60bb25430d4373ac54a15f2db004d189906a37da3997afd50a1b374800e82270e33800000000000000000000000000000000000000000000000000000000660d6f41000000000000000000000000000000000000000000000eb82507d03c7a5800004254432f555344000000000000000000000000000000000000000000000000000262dcaa2d204c8b187d556135fff9fedf2c91be7d366755c801a4d63cd5d632494f9996a39fd91a88bafb355c8b32636f0c43004993e2642024530d599667ba7800000000000000000000000000000000000000000000000000000000660d6f410000000000000000000000000000000000000000000000000b84c09a3b9300004e474c2f55534400000000000000000000000000000000000000000000000000",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			// Build update bytes
			merkleRoot := ethcommon.Hex2Bytes(tC.merkleRoot)

			signatures := make([]ECDSASignature, len(tC.signatures))
			for i, sig := range tC.signatures {
				sigBytes := ethcommon.Hex2Bytes(sig)
				signatures[i] = ECDSASignature{
					R: ethcommon.BytesToHash(sigBytes[:32]),
					S: ethcommon.BytesToHash(sigBytes[32:64]),
					V: sigBytes[64],
				}
			}

			merkleUpdates := make([]*MerkleRootUpdate, len(tC.updateData))
			for i, update := range tC.updateData {
				dataKey, err := utils.AsciiToPaddedHex(update.DataKey)
				r.NoError(err)

				merkleProof := make([][32]byte, len(update.MerkleProof))
				for j, proof := range update.MerkleProof {
					merkleProof[j] = ethcommon.HexToHash(proof)
				}

				price, success := new(big.Int).SetString(update.Price, 10)
				r.True(success)

				update := &MerkleRootUpdate{
					DataKey: dataKey,
					NewMerkleRoot: ethcommon.BytesToHash(merkleRoot[:]),
					MerkleProof: merkleProof,
					Signatures: signatures,
					Price: price,
					Timestamp: update.Timestamp,
				}

				merkleUpdates[i] = update
			}

			multipleUpdate, err := NewMekrleRootUpdateMultipleFromUpdates(merkleUpdates)
			r.NoError(err)

			expectedCalldataBytes := ethcommon.Hex2Bytes(tC.expectedCalldata)

			updatesCalldata, err := multipleUpdate.ToCalldata()
			r.NoError(err)
			r.Equal(expectedCalldataBytes, updatesCalldata)
		})
	}
}
