// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NameSpecifier] class.
var NameSpecifierClass objc.Class

func init() {
	NameSpecifierClass = objc.GetClass("NSNameSpecifier")
}

type NameSpecifier struct {
	objc.ID
}

func NameSpecifierFrom(ptr unsafe.Pointer) NameSpecifier {
	return NameSpecifier{
		ID: objc.ID(ptr),
	}
}



