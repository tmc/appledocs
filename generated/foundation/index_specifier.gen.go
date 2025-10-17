// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IndexSpecifier] class.
var IndexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}

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




