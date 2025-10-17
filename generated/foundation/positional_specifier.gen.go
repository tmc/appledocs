// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PositionalSpecifier] class.
var PositionalSpecifierClass objc.Class

func init() {
	PositionalSpecifierClass = objc.GetClass("NSPositionalSpecifier")
}

type PositionalSpecifier struct {
	objc.ID
}

func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{
		ID: objc.ID(ptr),
	}
}



