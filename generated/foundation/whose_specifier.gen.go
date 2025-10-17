// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WhoseSpecifier] class.
var WhoseSpecifierClass objc.Class

func init() {
	WhoseSpecifierClass = objc.GetClass("NSWhoseSpecifier")
}

type WhoseSpecifier struct {
	objc.ID
}

func WhoseSpecifierFrom(ptr unsafe.Pointer) WhoseSpecifier {
	return WhoseSpecifier{
		ID: objc.ID(ptr),
	}
}



