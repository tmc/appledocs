// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PropertySpecifier] class.
var PropertySpecifierClass objc.Class

func init() {
	PropertySpecifierClass = objc.GetClass("NSPropertySpecifier")
}

type PropertySpecifier struct {
	objc.ID
}

func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ID: objc.ID(ptr),
	}
}



