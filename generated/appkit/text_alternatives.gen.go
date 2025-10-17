// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAlternatives] class.
var TextAlternativesClass objc.Class

func init() {
	TextAlternativesClass = objc.GetClass("NSTextAlternatives")
}

type TextAlternatives struct {
	objc.ID
}

func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{
		ID: objc.ID(ptr),
	}
}



