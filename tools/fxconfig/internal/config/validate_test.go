/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateBCCSPDefault(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		provider    string
		expectError bool
		errMatch    string
	}{
		{name: "empty defaults to SW", provider: "", expectError: false},
		{name: "explicit SW", provider: "SW", expectError: false},
		{
			name:        "unknown provider rejected",
			provider:    "magic",
			expectError: true,
			errMatch:    "not supported",
		},
		{
			name:        "lowercase sw rejected (case-sensitive)",
			provider:    "sw",
			expectError: true,
			errMatch:    "not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateBCCSPDefault(tt.provider)
			if !tt.expectError {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			if tt.errMatch != "" {
				require.Contains(t, err.Error(), tt.errMatch)
			}
		})
	}
}

// TestValidateBCCSPDefault_PKCS11 verifies the PKCS11 branch: accepted when
// the binary is built with -tags pkcs11, rejected with a clear message
// otherwise. The pkcs11Supported flag is set by build-tagged files.
func TestValidateBCCSPDefault_PKCS11(t *testing.T) {
	t.Parallel()

	err := validateBCCSPDefault("PKCS11")
	if pkcs11Supported {
		require.NoError(t, err)
		return
	}
	require.Error(t, err)
	require.Contains(t, err.Error(), "-tags pkcs11")
}

// TestErrorIfEmpty tests the errorIfEmpty helper function.
func TestErrorIfEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "empty string",
			input:       "",
			expectError: true,
		},
		{
			name:        "non-empty string",
			input:       "test",
			expectError: false,
		},
		{
			name:        "whitespace only",
			input:       "   ",
			expectError: true,
		},
		{
			name:        "single space",
			input:       " ",
			expectError: true,
		},
		{
			name:        "tab character",
			input:       "\t",
			expectError: true,
		},
		{
			name:        "newline character",
			input:       "\n",
			expectError: true,
		},
		{
			name:        "mixed whitespace",
			input:       " \t\n ",
			expectError: true,
		},
		{
			name:        "string with leading/trailing spaces",
			input:       "  test  ",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := errorIfEmpty(tt.input, "test error message")
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
