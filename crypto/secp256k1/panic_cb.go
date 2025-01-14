// Copyright 2015 Jeffrey Wilcke, Felix Lange, Gustav Simonsson. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build !gofuzz
// +build cgo

package secp256k1

import "C"
import "unsafe"

// Callbacks for converting libsecp256k1 internal faults into
// recoverable Go panics.

//export secp256k1GoPanicIllegal2
func secp256k1GoPanicIllegal2(msg *C.char, data unsafe.Pointer) {
	panic("illegal argument: " + C.GoString(msg))
}

//export secp256k1GoPanicError2
func secp256k1GoPanicError2(msg *C.char, data unsafe.Pointer) {
	panic("internal error: " + C.GoString(msg))
}
