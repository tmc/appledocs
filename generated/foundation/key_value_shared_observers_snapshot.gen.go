// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [KeyValueSharedObserversSnapshot] class.
var keyValueSharedObserversSnapshotClass = _KeyValueSharedObserversSnapshotClass{objc.GetClass("NSKeyValueSharedObserversSnapshot")}

type _KeyValueSharedObserversSnapshotClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObserversSnapshot

type KeyValueSharedObserversSnapshot struct {
	objectivec.Object
}

// KeyValueSharedObserversSnapshotFrom constructs a [KeyValueSharedObserversSnapshot] from an unsafe.Pointer.
func KeyValueSharedObserversSnapshotFrom(ptr unsafe.Pointer) KeyValueSharedObserversSnapshot {
	return KeyValueSharedObserversSnapshot{objectivec.Object{objc.ID(ptr)}}
}



