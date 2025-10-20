// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MiddleSpecifierClass _MiddleSpecifierClass

func init() {
	MiddleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}
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




