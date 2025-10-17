// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AppleEventManager] class.
var AppleEventManagerClass objc.Class

func init() {
	AppleEventManagerClass = objc.GetClass("NSAppleEventManager")
}

type AppleEventManager struct {
	objc.ID
}

func AppleEventManagerFrom(ptr unsafe.Pointer) AppleEventManager {
	return AppleEventManager{
		ID: objc.ID(ptr),
	}
}




