// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var AppleEventManagerClass _AppleEventManagerClass

func init() {
	AppleEventManagerClass = _AppleEventManagerClass{objc.GetClass("NSAppleEventManager")}
}

type _AppleEventManagerClass struct {
	class objc.Class
}

type AppleEventManager struct {
	objc.ID
}

func AppleEventManagerFrom(ptr unsafe.Pointer) AppleEventManager {
	return AppleEventManager{
		ID: objc.ID(ptr),
	}
}




