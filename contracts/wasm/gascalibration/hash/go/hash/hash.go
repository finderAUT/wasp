// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package hash

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

func funcF(ctx wasmlib.ScFuncContext, f *FContext) {
	//todo add parameter
	bytes := wasmtypes.HexDecode("4ce645a2ed77a165be220d31e13e1e3533488d793aa05cbfdc2a98f5d9b39dfd")
	n := f.Params.N().Value()

	for i := uint32(0); i < n; i++ {
		bytes = ctx.Utility().HashBlake2b(bytes).Bytes()
	}
}
