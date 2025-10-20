// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var KeyValueSharedObserversClass _KeyValueSharedObserversClass

func init() {
	KeyValueSharedObserversClass = _KeyValueSharedObserversClass{objc.GetClass("NSKeyValueSharedObservers")}
}

type _KeyValueSharedObserversClass struct {
	class objc.Class
}

type KeyValueSharedObservers struct {
	objc.ID
}

func KeyValueSharedObserversFrom(ptr unsafe.Pointer) KeyValueSharedObservers {
	return KeyValueSharedObservers{
		ID: objc.ID(ptr),
	}
}




