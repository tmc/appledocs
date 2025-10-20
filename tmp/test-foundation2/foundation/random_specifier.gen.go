// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var randomSpecifierClass _RandomSpecifierClass

func init() {
	randomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}
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




