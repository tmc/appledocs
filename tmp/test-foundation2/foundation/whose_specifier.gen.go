// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var whoseSpecifierClass _WhoseSpecifierClass

func init() {
	whoseSpecifierClass = _WhoseSpecifierClass{objc.GetClass("NSWhoseSpecifier")}
}

type _WhoseSpecifierClass struct {
	class objc.Class
}

type WhoseSpecifier struct {
	objc.ID
}

func WhoseSpecifierFrom(ptr unsafe.Pointer) WhoseSpecifier {
	return WhoseSpecifier{
		ID: objc.ID(ptr),
	}
}




