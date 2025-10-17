// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TokenField] class.
var TokenFieldClass objc.Class

func init() {
	TokenFieldClass = objc.GetClass("NSTokenField")
}

type TokenField struct {
	objc.ID
}

func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		ID: objc.ID(ptr),
	}
}




