// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var nameSpecifierClass _NameSpecifierClass

func init() {
	nameSpecifierClass = _NameSpecifierClass{objc.GetClass("NSNameSpecifier")}
}

type _NameSpecifierClass struct {
	class objc.Class
}

type NameSpecifier struct {
	objc.ID
}

func NameSpecifierFrom(ptr unsafe.Pointer) NameSpecifier {
	return NameSpecifier{
		ID: objc.ID(ptr),
	}
}




