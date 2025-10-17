// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [KeyValueSharedObservers] class.
var KeyValueSharedObserversClass objc.Class

func init() {
	KeyValueSharedObserversClass = objc.GetClass("NSKeyValueSharedObservers")
}

type KeyValueSharedObservers struct {
	objc.ID
}

func KeyValueSharedObserversFrom(ptr unsafe.Pointer) KeyValueSharedObservers {
	return KeyValueSharedObservers{
		ID: objc.ID(ptr),
	}
}




