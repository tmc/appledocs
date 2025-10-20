// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var middleSpecifierClass _MiddleSpecifierClass

func init() {
	middleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}
}

type _MiddleSpecifierClass struct {
	class objc.Class
}

type MiddleSpecifier struct {
	objc.ID
}

func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ID: objc.ID(ptr),
	}
}




