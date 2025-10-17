// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Switch_] class.
var Switch_Class objc.Class

func init() {
	Switch_Class = objc.GetClass("NSSwitch")
}

type Switch_ struct {
	objc.ID
}

func Switch_From(ptr unsafe.Pointer) Switch_ {
	return Switch_{
		ID: objc.ID(ptr),
	}
}




