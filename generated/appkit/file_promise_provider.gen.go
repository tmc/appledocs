// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilePromiseProvider] class.
var filePromiseProviderClass = _FilePromiseProviderClass{objc.GetClass("NSFilePromiseProvider")}

type _FilePromiseProviderClass struct {
	class objc.Class
}

// An object that provides a promise for the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider

type FilePromiseProvider struct {
	objectivec.Object
}

// FilePromiseProviderFrom constructs a [FilePromiseProvider] from an unsafe.Pointer.
//
// An object that provides a promise for the pasteboard.
func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{objectivec.Object{objc.ID(ptr)}}
}



