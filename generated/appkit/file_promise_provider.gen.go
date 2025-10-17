// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FilePromiseProvider] class.
var FilePromiseProviderClass objc.Class

func init() {
	FilePromiseProviderClass = objc.GetClass("NSFilePromiseProvider")
}

type FilePromiseProvider struct {
	objc.ID
}

func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{
		ID: objc.ID(ptr),
	}
}



