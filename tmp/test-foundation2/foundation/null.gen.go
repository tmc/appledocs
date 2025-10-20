// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var nullClass _NullClass

func init() {
	nullClass = _NullClass{objc.GetClass("NSNull")}
}

type _NullClass struct {
	class objc.Class
}

type Null struct {
	objc.ID
}

func NullFrom(ptr unsafe.Pointer) Null {
	return Null{
		ID: objc.ID(ptr),
	}
}




