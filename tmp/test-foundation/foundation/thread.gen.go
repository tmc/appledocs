// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ThreadClass _ThreadClass

func init() {
	ThreadClass = _ThreadClass{objc.GetClass("NSThread")}
}

type _ThreadClass struct {
	class objc.Class
}

type Thread struct {
	objc.ID
}

func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{
		ID: objc.ID(ptr),
	}
}




