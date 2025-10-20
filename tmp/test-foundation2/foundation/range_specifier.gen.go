// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var rangeSpecifierClass _RangeSpecifierClass

func init() {
	rangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}
}

type _RangeSpecifierClass struct {
	class objc.Class
}

type RangeSpecifier struct {
	objc.ID
}

func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ID: objc.ID(ptr),
	}
}




