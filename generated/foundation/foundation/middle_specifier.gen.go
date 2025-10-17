// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MiddleSpecifier] class.
var MiddleSpecifierClass objc.Class

func init() {
	MiddleSpecifierClass = objc.GetClass("NSMiddleSpecifier")
}

type MiddleSpecifier struct {
	objc.ID
}

func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ID: objc.ID(ptr),
	}
}




