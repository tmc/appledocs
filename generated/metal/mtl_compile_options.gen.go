// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLCompileOptions] class.
var mTLCompileOptionsClass = _MTLCompileOptionsClass{objc.GetClass("MTLCompileOptions")}

type _MTLCompileOptionsClass struct {
	class objc.Class
}

// Compilation settings for a Metal shader library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions

type MTLCompileOptions struct {
	objectivec.Object
}

// MTLCompileOptionsFrom constructs a [MTLCompileOptions] from an unsafe.Pointer.
//
// Compilation settings for a Metal shader library.
func MTLCompileOptionsFrom(ptr unsafe.Pointer) MTLCompileOptions {
	return MTLCompileOptions{objectivec.Object{objc.ID(ptr)}}
}



