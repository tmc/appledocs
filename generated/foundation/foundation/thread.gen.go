// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Thread] class.
var ThreadClass objc.Class

func init() {
	ThreadClass = objc.GetClass("NSThread")
}

type Thread struct {
	objc.ID
}

func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{
		ID: objc.ID(ptr),
	}
}




