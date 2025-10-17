// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [KeyValueSharedObservers] class.
var keyValueSharedObserversClass = _KeyValueSharedObserversClass{objc.GetClass("NSKeyValueSharedObservers")}

type _KeyValueSharedObserversClass struct {
	class objc.Class
}

// An interface definition for the [KeyValueSharedObservers] class.
type IKeyValueSharedObservers interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers

type KeyValueSharedObservers struct {
	objectivec.Object
}

// KeyValueSharedObserversFrom constructs a [KeyValueSharedObservers] from an unsafe.Pointer.
func KeyValueSharedObserversFrom(ptr unsafe.Pointer) KeyValueSharedObservers {
	return KeyValueSharedObservers{objectivec.Object{objc.ID(ptr)}}
}



