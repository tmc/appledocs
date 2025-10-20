// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var RandomSpecifierClass _RandomSpecifierClass

func init() {
	RandomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}
}

type _RandomSpecifierClass struct {
	class objc.Class
}

type RandomSpecifier struct {
	objc.ID
}

func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ID: objc.ID(ptr),
	}
}




