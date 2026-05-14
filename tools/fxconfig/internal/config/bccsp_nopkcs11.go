//go:build !pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"github.com/hyperledger/fabric-lib-go/bccsp/factory"
)

// pkcs11Supported reports whether this build of fxconfig was compiled with
// the pkcs11 build tag.
const pkcs11Supported = false

// applyPKCS11Opts is a no-op when fxconfig is built without the pkcs11 tag.
// MSPConfig.Validate rejects bccsp.default=PKCS11 in non-pkcs11 builds before
// this is reached on the happy path.
func applyPKCS11Opts(_ *factory.FactoryOpts, _ BCCSPPKCS11Config) {}
