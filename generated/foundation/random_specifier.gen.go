// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RandomSpecifier] class.
var RandomSpecifierClass objc.Class

func init() {
	RandomSpecifierClass = objc.GetClass("NSRandomSpecifier")
}

type RandomSpecifier struct {
	objc.ID
}

func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ID: objc.ID(ptr),
	}
}



