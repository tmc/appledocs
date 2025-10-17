// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RangeSpecifier] class.
var RangeSpecifierClass objc.Class

func init() {
	RangeSpecifierClass = objc.GetClass("NSRangeSpecifier")
}

type RangeSpecifier struct {
	objc.ID
}

func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ID: objc.ID(ptr),
	}
}




