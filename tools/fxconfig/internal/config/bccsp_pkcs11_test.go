//go:build pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMSPConfigToFactoryOpts_PKCS11(t *testing.T) {
	t.Parallel()

	mspCfg := MSPConfig{
		ConfigPath: "/tmp/msp",
		BCCSP: BCCSPConfig{
			Default: "PKCS11",
			PKCS11: BCCSPPKCS11Config{
				Library:        "/usr/local/lib/libsofthsm2.so",
				Label:          "TestLabel",
				Pin:            "1234",
				Hash:           "SHA2",
				Security:       256,
				SoftwareVerify: true,
			},
		},
	}

	opts := mspCfg.ToFactoryOpts()

	require.Equal(t, "PKCS11", opts.Default)
	require.NotNil(t, opts.PKCS11)
	require.Equal(t, "/usr/local/lib/libsofthsm2.so", opts.PKCS11.Library)
	require.Equal(t, "TestLabel", opts.PKCS11.Label)
	require.Equal(t, "1234", opts.PKCS11.Pin)
	require.Equal(t, "SHA2", opts.PKCS11.Hash)
	require.Equal(t, 256, opts.PKCS11.Security)
	require.True(t, opts.PKCS11.SoftwareVerify)
	require.Nil(t, opts.SW, "SW opts must be nil when PKCS11 provider is selected")
}

func TestMSPConfigToFactoryOpts_PKCS11_DefaultsHashAndSecurity(t *testing.T) {
	t.Parallel()

	mspCfg := MSPConfig{
		ConfigPath: "/tmp/msp",
		BCCSP: BCCSPConfig{
			Default: "PKCS11",
			PKCS11: BCCSPPKCS11Config{
				Library: "/usr/local/lib/libsofthsm2.so",
				Label:   "TestLabel",
				Pin:     "1234",
			},
		},
	}

	opts := mspCfg.ToFactoryOpts()

	require.Equal(t, "PKCS11", opts.Default)
	require.NotNil(t, opts.PKCS11)
	require.Equal(t, "SHA2", opts.PKCS11.Hash)
	require.Equal(t, 256, opts.PKCS11.Security)
}

func TestMSPConfigToFactoryOpts_PKCS11_NotActivatedWithoutDefault(t *testing.T) {
	t.Parallel()

	// Library is set but Default is not "PKCS11" => SW path wins.
	mspCfg := MSPConfig{
		ConfigPath: "/tmp/msp",
		BCCSP: BCCSPConfig{
			PKCS11: BCCSPPKCS11Config{
				Library: "/usr/local/lib/libsofthsm2.so",
				Label:   "should-be-ignored",
			},
		},
	}

	opts := mspCfg.ToFactoryOpts()

	require.Equal(t, "SW", opts.Default)
	require.NotNil(t, opts.SW)
	require.Nil(t, opts.PKCS11)
}
