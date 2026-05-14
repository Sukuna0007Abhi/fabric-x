//go:build !pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestMSPConfigToFactoryOpts_PKCS11_IgnoredWithoutTag verifies that a
// PKCS#11 sub-section is silently ignored when Default is not set to PKCS11
// (the SW provider remains the default).
func TestMSPConfigToFactoryOpts_PKCS11_IgnoredWithoutTag(t *testing.T) {
	t.Parallel()

	mspCfg := MSPConfig{
		ConfigPath: "/tmp/msp",
		BCCSP: BCCSPConfig{
			PKCS11: BCCSPPKCS11Config{
				Library: "/usr/local/lib/libsofthsm2.so",
				Label:   "TestLabel",
				Pin:     "1234",
			},
		},
	}

	opts := mspCfg.ToFactoryOpts()

	require.Equal(t, "SW", opts.Default)
	require.NotNil(t, opts.SW)
}
