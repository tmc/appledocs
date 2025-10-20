// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var positionalSpecifierClass _PositionalSpecifierClass

func init() {
	positionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}
}

type _PositionalSpecifierClass struct {
	class objc.Class
}

type PositionalSpecifier struct {
	objc.ID
}

func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{
		ID: objc.ID(ptr),
	}
}




