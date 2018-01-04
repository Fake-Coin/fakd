// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"encoding/hex"
	"log"
	"time"

	"github.com/NeilVallon/fakd/chaincfg/chainhash"
	"github.com/NeilVallon/fakd/wire"
)

func decode(h string) []byte {
	decoded, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 4294967295,
			},
			SignatureScript: decode(`04ffff001d0104383c3e4e6f6d61747465722077686174205368776f6f6420736179732c20686520697320666f7265766572206e6f74207261636973742e3c3e`),
			Sequence:        4294967295,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value:    5000000000,
			PkScript: decode(`414F1B353118B75E2E9E353475880ED919544C7EF59259B6AF718910211B06FB8CF6F41FA05FCF7997C337C1CF06969404C7696E582DBBBCF2396616ABA0B0DE0818AC`),
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{
	0x33, 0x7c, 0xb1, 0x3f, 0xbb, 0x8d, 0x89, 0xcf, 0x59, 0x4c, 0x1f, 0xb2, 0xf5, 0xa1, 0x78, 0xa5, 0x0d, 0x83, 0x68, 0x74, 0x47, 0x36, 0xc1, 0xdf, 0x5b, 0xb3, 0xf5, 0xfa, 0x3e, 0x20, 0xc9, 0xae,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
	0x5A, 0x5D, 0x8F, 0xF1, 0x9E, 0x91, 0x23, 0xE9, 0x03, 0x73, 0x1D, 0xE0, 0xC3, 0xC7, 0xC0, 0x4C, 0xCF, 0x0A, 0x7D, 0xE8, 0x4D, 0xCE, 0x31, 0x58, 0xF8, 0x56, 0xF8, 0x1D, 0xE5, 0xD6, 0x93, 0x4B,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},  // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot, //
		Timestamp:  time.Unix(0x5A2E238E, 0),
		Bits:       0x1E10024c,
		Nonce:      0x0F5B6AC4,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet4GenesisHash is the hash of the first block in the block chain for the
// test network (version 4).
var testNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xa0, 0x29, 0x3e, 0x4e, 0xeb, 0x3d, 0xa6, 0xe6,
	0xf5, 0x6f, 0x81, 0xed, 0x59, 0x5f, 0x57, 0x88,
	0xd, 0x1a, 0x21, 0x56, 0x9e, 0x13, 0xee, 0xfd,
	0xd9, 0x51, 0x28, 0x4b, 0x5a, 0x62, 0x66, 0x49,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet4GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 4).  It is the same as the merkle root
// for the main network.
var testNet4GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xd9, 0xce, 0xd4, 0xed, 0x11, 0x30, 0xf7, 0xb7,
	0xfa, 0xad, 0x9b, 0xe2, 0x53, 0x23, 0xff, 0xaf,
	0xa3, 0x32, 0x32, 0xa1, 0x7c, 0x3e, 0xdf, 0x6c,
	0xfd, 0x97, 0xbe, 0xe6, 0xba, 0xfb, 0xdd, 0x97,
})

// testNet4GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 4).
var testNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet4GenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1486949366, 0),
		Bits:       0x1e0ffff0,
		Nonce:      293345,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
