// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var setClass _SetClass

func init() {
	setClass = _SetClass{objc.GetClass("NSSet")}
}

type _SetClass struct {
	class objc.Class
}

type Set struct {
	objc.ID
}

func SetFrom(ptr unsafe.Pointer) Set {
	return Set{
		ID: objc.ID(ptr),
	}
}




