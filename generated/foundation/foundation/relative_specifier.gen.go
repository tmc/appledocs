// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RelativeSpecifier] class.
var RelativeSpecifierClass objc.Class

func init() {
	RelativeSpecifierClass = objc.GetClass("NSRelativeSpecifier")
}

type RelativeSpecifier struct {
	objc.ID
}

func RelativeSpecifierFrom(ptr unsafe.Pointer) RelativeSpecifier {
	return RelativeSpecifier{
		ID: objc.ID(ptr),
	}
}




