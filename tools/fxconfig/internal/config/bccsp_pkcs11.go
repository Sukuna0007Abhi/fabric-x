//go:build pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"cmp"

	"github.com/hyperledger/fabric-lib-go/bccsp/factory"
	"github.com/hyperledger/fabric-lib-go/bccsp/pkcs11"
)

// pkcs11Supported reports whether this build of fxconfig was compiled with
// the pkcs11 build tag.
const pkcs11Supported = true

// applyPKCS11Opts populates the PKCS#11 sub-section of the Fabric factory
// options. Caller is responsible for setting opts.Default = "PKCS11".
func applyPKCS11Opts(opts *factory.FactoryOpts, cfg BCCSPPKCS11Config) {
	opts.PKCS11 = &pkcs11.PKCS11Opts{
		Library:        cfg.Library,
		Label:          cfg.Label,
		Pin:            cfg.Pin,
		Hash:           cmp.Or(cfg.Hash, "SHA2"),
		Security:       cmp.Or(cfg.Security, 256),
		SoftwareVerify: cfg.SoftwareVerify,
		Immutable:      cfg.Immutable,
	}
}
