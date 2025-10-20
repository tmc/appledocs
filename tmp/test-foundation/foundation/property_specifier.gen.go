// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var PropertySpecifierClass _PropertySpecifierClass

func init() {
	PropertySpecifierClass = _PropertySpecifierClass{objc.GetClass("NSPropertySpecifier")}
}

type _PropertySpecifierClass struct {
	class objc.Class
}

type PropertySpecifier struct {
	objc.ID
}

func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ID: objc.ID(ptr),
	}
}




