// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unarchiverClass _UnarchiverClass

func init() {
	unarchiverClass = _UnarchiverClass{objc.GetClass("NSUnarchiver")}
}

type _UnarchiverClass struct {
	class objc.Class
}

type Unarchiver struct {
	objc.ID
}

func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		ID: objc.ID(ptr),
	}
}




