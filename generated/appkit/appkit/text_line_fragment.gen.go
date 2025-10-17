// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLineFragment] class.
var TextLineFragmentClass objc.Class

func init() {
	TextLineFragmentClass = objc.GetClass("NSTextLineFragment")
}

type TextLineFragment struct {
	objc.ID
}

func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		ID: objc.ID(ptr),
	}
}



