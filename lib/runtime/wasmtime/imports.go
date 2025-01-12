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

package wasmtime

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go"
)

// int64ToPointerAndSize converts an int64 into a int32 pointer and a int32 length
func int64ToPointerAndSize(in int64) (ptr, length int32) {
	return int32(in), int32(in >> 32)
}

// Convert 64bit wasm span descriptor to Go memory slice
func asMemorySlice(memory []byte, span int64) []byte {
	ptr, size := int64ToPointerAndSize(span)
	return memory[ptr : ptr+size]
}

func ext_logging_log_version_1(c *wasmtime.Caller, level int32, target, msg int64) {
	logger.Trace("[ext_logging_log_version_1] executing...")
}

func ext_sandbox_instance_teardown_version_1(c *wasmtime.Caller, a int32) {
	logger.Trace("[ext_sandbox_instance_teardown_version_1] executing...")
}

func ext_sandbox_instantiate_version_1(c *wasmtime.Caller, a int32, x, y int64, z int32) int32 {
	logger.Trace("[ext_sandbox_instantiate_version_1] executing...")
	return 0
}

func ext_sandbox_invoke_version_1(c *wasmtime.Caller, a int32, x, y int64, z, d, e int32) int32 {
	logger.Trace("[ext_sandbox_invoke_version_1] executing...")
	return 0
}

func ext_sandbox_memory_get_version_1(c *wasmtime.Caller, a, z, d, e int32) int32 {
	logger.Trace("[ext_sandbox_memory_get_version_1] executing...")
	return 0
}

func ext_sandbox_memory_new_version_1(c *wasmtime.Caller, a, z int32) int32 {
	logger.Trace("[ext_sandbox_memory_new_version_1] executing...")
	return 0
}

func ext_sandbox_memory_set_version_1(c *wasmtime.Caller, a, z, d, e int32) int32 {
	logger.Trace("[ext_sandbox_memory_set_version_1] executing...")
	return 0
}

func ext_sandbox_memory_teardown_version_1(c *wasmtime.Caller, a int32) {
	logger.Trace("[ext_sandbox_memory_teardown_version_1] executing...")
}

func ext_crypto_ed25519_generate_version_1(c *wasmtime.Caller, a int32, z int64) int32 {
	logger.Trace("[ext_crypto_ed25519_generate_version_1] executing...")
	return 0
}

func ext_crypto_ed25519_verify_version_1(c *wasmtime.Caller, a int32, z int64, y int32) int32 {
	logger.Trace("[ext_crypto_ed25519_verify_version_1] executing...")
	return 0
}

func ext_crypto_finish_batch_verify_version_1(c *wasmtime.Caller) int32 {
	logger.Trace("[ext_crypto_finish_batch_verify_version_1] executing...")
	return 0
}

func ext_crypto_secp256k1_ecdsa_recover_compressed_version_1(c *wasmtime.Caller, a, z int32) int64 {
	logger.Trace("[ext_crypto_secp256k1_ecdsa_recover_compressed_version_1] executing...")
	return 0
}

func ext_crypto_sr25519_generate_version_1(c *wasmtime.Caller, a int32, z int64) int32 {
	logger.Trace("[ext_crypto_sr25519_generate_version_1] executing...")
	return 0
}

func ext_crypto_sr25519_public_keys_version_1(c *wasmtime.Caller, a int32) int64 {
	logger.Trace("[ext_crypto_sr25519_public_keys_version_1] executing...")
	return 0
}

func ext_crypto_sr25519_sign_version_1(c *wasmtime.Caller, a, z int32, y int64) int64 {
	logger.Trace("[ext_crypto_sr25519_sign_version_1] executing...")
	return 0
}

func ext_crypto_sr25519_verify_version_2(c *wasmtime.Caller, a int32, z int64, y int32) int32 {
	logger.Trace("[ext_crypto_sr25519_verify_version_2] executing...")
	return 0
}

func ext_crypto_start_batch_verify_version_1(c *wasmtime.Caller) {
	logger.Trace("[ext_crypto_start_batch_verify_version_1] executing...")
}

func ext_trie_blake2_256_ordered_root_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_trie_blake2_256_ordered_root_version_1] executing...")
	return 0
}

func ext_misc_print_hex_version_1(c *wasmtime.Caller, a int64) {
	logger.Trace("[ext_misc_print_hex_version_1] executing...")
}

func ext_misc_print_num_version_1(c *wasmtime.Caller, data int64) {
	logger.Trace("[ext_misc_print_num_version_1] executing...")
	logger.Info("[ext_print_num]", "message", fmt.Sprintf("%d", data))
}

func ext_misc_print_utf8_version_1(c *wasmtime.Caller, dataSpan int64) {
	logger.Trace("[ext_misc_print_utf8_version_1] executing...")
	m := c.GetExport("memory").Memory()
	data := asMemorySlice(m.UnsafeData(), dataSpan)
	logger.Info("[ext_print_utf8]", "message", data)
}

func ext_misc_runtime_version_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_misc_runtime_version_version_1] executing...")
	return 0
}

func ext_default_child_storage_clear_version_1(c *wasmtime.Caller, a, b int64) {
	logger.Trace("[ext_default_child_storage_clear_version_1] executing...")
}

func ext_default_child_storage_get_version_1(c *wasmtime.Caller, a, b int64) int64 {
	logger.Trace("[ext_default_child_storage_get_version_1] executing...")
	return 0
}

func ext_default_child_storage_root_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_default_child_storage_root_version_1] executing...")
	return 0
}

func ext_default_child_storage_set_version_1(c *wasmtime.Caller, a, b, z int64) {
	logger.Trace("[ext_default_child_storage_set_version_1] executing...")
}

func ext_default_child_storage_storage_kill_version_1(c *wasmtime.Caller, a int64) {
	logger.Trace("[ext_default_child_storage_storage_kill_version_1] executing...")
}

func ext_allocator_free_version_1(c *wasmtime.Caller, addr int32) {
	logger.Trace("[ext_allocator_free_version_1] executing...")
	err := ctx.Allocator.Deallocate(uint32(addr))
	if err != nil {
		logger.Error("[ext_free]", "error", err)
	}
}

func ext_allocator_malloc_version_1(c *wasmtime.Caller, size int32) int32 {
	logger.Trace("[ext_allocator_malloc_version_1] executing...")
	res, err := ctx.Allocator.Allocate(uint32(size))
	if err != nil {
		logger.Error("[ext_malloc]", "Error:", err)
	}
	return int32(res)
}

func ext_hashing_blake2_128_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_blake2_128_version_1] executing...")
	return 0
}

func ext_hashing_blake2_256_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_blake2_256_version_1] executing...")
	return 0
}

func ext_hashing_keccak_256_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_keccak_256_version_1] executing...")
	return 0
}

func ext_hashing_sha2_256_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_sha2_256_version_1] executing...")
	return 0
}

func ext_hashing_twox_128_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_twox_128_version_1] executing...")
	return 0
}

func ext_hashing_twox_64_version_1(c *wasmtime.Caller, z int64) int32 {
	logger.Trace("[ext_hashing_twox_64_version_1] executing...")
	return 0
}

func ext_offchain_is_validator_version_1(c *wasmtime.Caller) int32 {
	logger.Trace("[ext_offchain_is_validator_version_1] executing...")
	return 0
}

func ext_offchain_local_storage_compare_and_set_version_1(c *wasmtime.Caller, a int32, x, y, z int64) int32 {
	logger.Trace("[ext_offchain_local_storage_compare_and_set_version_1] executing...")
	return 0
}

func ext_offchain_local_storage_get_version_1(c *wasmtime.Caller, a int32, x int64) int64 {
	logger.Trace("[ext_offchain_local_storage_get_version_1] executing...")
	return 0
}

func ext_offchain_local_storage_set_version_1(c *wasmtime.Caller, a int32, x, y int64) {
	logger.Trace("[ext_offchain_local_storage_set_version_1] executing...")
}

func ext_offchain_network_state_version_1(c *wasmtime.Caller) int64 {
	logger.Trace("[ext_offchain_network_state_version_1] executing...")
	return 0
}

func ext_offchain_random_seed_version_1(c *wasmtime.Caller) int32 {
	logger.Trace("[ext_offchain_random_seed_version_1] executing...")
	return 0
}

func ext_offchain_submit_transaction_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_offchain_submit_transaction_version_1] executing...")
	return 0
}

func ext_storage_append_version_1(c *wasmtime.Caller, a, b int64) {
	logger.Trace("[ext_storage_append_version_1] executing...")
}

func ext_storage_changes_root_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_storage_changes_root_version_1] executing...")
	return 0
}

func ext_storage_clear_version_1(c *wasmtime.Caller, a int64) {
	logger.Trace("[ext_storage_clear_version_1] executing...")
}

func ext_storage_clear_prefix_version_1(c *wasmtime.Caller, a int64) {
	logger.Trace("[ext_storage_clear_prefix_version_1] executing...")
}

func ext_storage_commit_transaction_version_1(c *wasmtime.Caller) {
	logger.Trace("[ext_storage_commit_transaction_version_1] executing...")
}

func ext_storage_get_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_storage_get_version_1] executing...")
	return 0
}

func ext_storage_next_key_version_1(c *wasmtime.Caller, z int64) int64 {
	logger.Trace("[ext_storage_next_key_version_1] executing...")
	return 0
}

func ext_storage_read_version_1(c *wasmtime.Caller, a, b int64, x int32) int64 {
	logger.Trace("[ext_storage_read_version_1] executing...")
	return 0
}

func ext_storage_rollback_transaction_version_1(c *wasmtime.Caller) {
	logger.Trace("[ext_storage_rollback_transaction_version_1] executing...")
}

func ext_storage_root_version_1(c *wasmtime.Caller) int64 {
	logger.Trace("[ext_storage_root_version_1] executing...")
	return 0
}

func ext_storage_set_version_1(c *wasmtime.Caller, a, b int64) {
	logger.Trace("[ext_storage_set_version_1] executing...")
}

func ext_storage_start_transaction_version_1(c *wasmtime.Caller) {
	logger.Trace("[ext_storage_start_transaction_version_1] executing...")
}

func ext_offchain_index_set_version_1(c *wasmtime.Caller, a, b int64) {
	logger.Trace("[ext_offchain_index_set_version_1] executing...")
}

// ImportNodeRuntime adds the imports for the v0.8 runtime to linker
func ImportNodeRuntime(store *wasmtime.Store, memory *wasmtime.Memory) (*wasmtime.Linker, error) {
	fns := []struct {
		name string
		fn   interface{}
	}{
		{"ext_logging_log_version_1", ext_logging_log_version_1},
		{"ext_sandbox_instance_teardown_version_1", ext_sandbox_instance_teardown_version_1},
		{"ext_sandbox_instantiate_version_1", ext_sandbox_instantiate_version_1},
		{"ext_sandbox_invoke_version_1", ext_sandbox_invoke_version_1},
		{"ext_sandbox_memory_get_version_1", ext_sandbox_memory_get_version_1},
		{"ext_sandbox_memory_new_version_1", ext_sandbox_memory_new_version_1},
		{"ext_sandbox_memory_set_version_1", ext_sandbox_memory_set_version_1},
		{"ext_sandbox_memory_teardown_version_1", ext_sandbox_memory_teardown_version_1},
		{"ext_crypto_ed25519_generate_version_1", ext_crypto_ed25519_generate_version_1},
		{"ext_crypto_ed25519_verify_version_1", ext_crypto_ed25519_verify_version_1},
		{"ext_crypto_finish_batch_verify_version_1", ext_crypto_finish_batch_verify_version_1},
		{"ext_crypto_secp256k1_ecdsa_recover_compressed_version_1", ext_crypto_secp256k1_ecdsa_recover_compressed_version_1},
		{"ext_crypto_sr25519_public_keys_version_1", ext_crypto_sr25519_public_keys_version_1},
		{"ext_offchain_local_storage_set_version_1", ext_offchain_local_storage_set_version_1},
		{"ext_crypto_sr25519_sign_version_1", ext_crypto_sr25519_sign_version_1},
		{"ext_crypto_sr25519_verify_version_2", ext_crypto_sr25519_verify_version_2},
		{"ext_crypto_start_batch_verify_version_1", ext_crypto_start_batch_verify_version_1},
		{"ext_trie_blake2_256_ordered_root_version_1", ext_trie_blake2_256_ordered_root_version_1},
		{"ext_misc_print_hex_version_1", ext_misc_print_hex_version_1},
		{"ext_misc_print_num_version_1", ext_misc_print_num_version_1},
		{"ext_misc_print_utf8_version_1", ext_misc_print_utf8_version_1},
		{"ext_misc_runtime_version_version_1", ext_misc_runtime_version_version_1},
		{"ext_default_child_storage_clear_version_1", ext_default_child_storage_clear_version_1},
		{"ext_default_child_storage_get_version_1", ext_default_child_storage_get_version_1},
		{"ext_default_child_storage_root_version_1", ext_default_child_storage_root_version_1},
		{"ext_default_child_storage_set_version_1", ext_default_child_storage_set_version_1},
		{"ext_default_child_storage_storage_kill_version_1", ext_default_child_storage_storage_kill_version_1},
		{"ext_allocator_free_version_1", ext_allocator_free_version_1},
		{"ext_allocator_malloc_version_1", ext_allocator_malloc_version_1},
		{"ext_hashing_blake2_128_version_1", ext_hashing_blake2_128_version_1},
		{"ext_hashing_blake2_256_version_1", ext_hashing_blake2_256_version_1},
		{"ext_hashing_keccak_256_version_1", ext_hashing_keccak_256_version_1},
		{"ext_hashing_sha2_256_version_1", ext_hashing_sha2_256_version_1},
		{"ext_hashing_twox_128_version_1", ext_hashing_twox_128_version_1},
		{"ext_hashing_twox_64_version_1", ext_hashing_twox_64_version_1},
		{"ext_offchain_is_validator_version_1", ext_offchain_is_validator_version_1},
		{"ext_offchain_local_storage_compare_and_set_version_1", ext_offchain_local_storage_compare_and_set_version_1},
		{"ext_crypto_sr25519_generate_version_1", ext_crypto_sr25519_generate_version_1},
		{"ext_offchain_local_storage_get_version_1", ext_offchain_local_storage_get_version_1},
		{"ext_offchain_network_state_version_1", ext_offchain_network_state_version_1},
		{"ext_offchain_random_seed_version_1", ext_offchain_random_seed_version_1},
		{"ext_offchain_submit_transaction_version_1", ext_offchain_submit_transaction_version_1},
		{"ext_storage_append_version_1", ext_storage_append_version_1},
		{"ext_storage_changes_root_version_1", ext_storage_changes_root_version_1},
		{"ext_storage_clear_version_1", ext_storage_clear_version_1},
		{"ext_storage_clear_prefix_version_1", ext_storage_clear_prefix_version_1},
		{"ext_storage_commit_transaction_version_1", ext_storage_commit_transaction_version_1},
		{"ext_storage_get_version_1", ext_storage_get_version_1},
		{"ext_storage_next_key_version_1", ext_storage_next_key_version_1},
		{"ext_storage_read_version_1", ext_storage_read_version_1},
		{"ext_storage_rollback_transaction_version_1", ext_storage_rollback_transaction_version_1},
		{"ext_storage_root_version_1", ext_storage_root_version_1},
		{"ext_storage_set_version_1", ext_storage_set_version_1},
		{"ext_storage_start_transaction_version_1", ext_storage_start_transaction_version_1},
		{"ext_offchain_index_set_version_1", ext_offchain_index_set_version_1},
	}

	linker := wasmtime.NewLinker(store)
	if err := linker.Define("env", "memory", memory); err != nil {
		return nil, err
	}

	for _, f := range fns {
		if err := linker.DefineFunc("env", f.name, f.fn); err != nil {
			return nil, err
		}
	}
	return linker, nil
}
