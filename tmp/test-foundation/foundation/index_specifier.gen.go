// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var IndexSpecifierClass _IndexSpecifierClass

func init() {
	IndexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}
}

type _IndexSpecifierClass struct {
	class objc.Class
}

type IndexSpecifier struct {
	objc.ID
}

func IndexSpecifierFrom(ptr unsafe.Pointer) IndexSpecifier {
	return IndexSpecifier{
		ID: objc.ID(ptr),
	}
}




