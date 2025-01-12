// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package common

var (
	// BestBlockHashKey is the db location the hash of the best (unfinalised) block header.
	BestBlockHashKey = []byte("best_hash")
	// LatestStorageHashKey is the db location of the hash of the latest storage trie.
	LatestStorageHashKey = []byte("latest_storage_hash")
	// FinalizedBlockHashKey is the db location of the hash of the latest finalised block header.
	FinalizedBlockHashKey = []byte("finalised_head")
	// GenesisDataKey is the db location of the genesis data.
	GenesisDataKey = []byte("genesis_data")
	// BlockTreeKey is the db location of the encoded block tree structure.
	BlockTreeKey = []byte("block_tree")
	// LatestFinalizedRoundKey is the key where the last finalised grandpa round is stored
	LatestFinalizedRoundKey = []byte("latest_finalised_round")
	// WorkingStorageHashKey is the storage key that the runtime uses to store the latest working state root.
	WorkingStorageHashKey = []byte("working_storage_hash")
	//NodeNameKey is the storage key to store de current node name and avoid create a new name every initialization
	NodeNameKey = []byte("node_name")
)
