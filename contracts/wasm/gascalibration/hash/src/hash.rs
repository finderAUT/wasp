// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_f(_ctx: &ScFuncContext, f: &FContext) {
    let n = f.params.n().value();
    let mut bytes = hex_decode("4ce645a2ed77a165be220d31e13e1e3533488d793aa05cbfdc2a98f5d9b39dfd");

    for _ in 0..n {
        bytes = _ctx.utility().hash_blake2b(&bytes).to_bytes();
    }
}
