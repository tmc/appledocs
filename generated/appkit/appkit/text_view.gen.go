// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextView] class.
var TextViewClass objc.Class

func init() {
	TextViewClass = objc.GetClass("NSTextView")
}

type TextView struct {
	objc.ID
}

func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		ID: objc.ID(ptr),
	}
}




