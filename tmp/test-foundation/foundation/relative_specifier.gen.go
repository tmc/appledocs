// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var RelativeSpecifierClass _RelativeSpecifierClass

func init() {
	RelativeSpecifierClass = _RelativeSpecifierClass{objc.GetClass("NSRelativeSpecifier")}
}

type _RelativeSpecifierClass struct {
	class objc.Class
}

type RelativeSpecifier struct {
	objc.ID
}

func RelativeSpecifierFrom(ptr unsafe.Pointer) RelativeSpecifier {
	return RelativeSpecifier{
		ID: objc.ID(ptr),
	}
}




