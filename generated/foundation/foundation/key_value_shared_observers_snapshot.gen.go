// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [KeyValueSharedObserversSnapshot] class.
var KeyValueSharedObserversSnapshotClass objc.Class

func init() {
	KeyValueSharedObserversSnapshotClass = objc.GetClass("NSKeyValueSharedObserversSnapshot")
}

type KeyValueSharedObserversSnapshot struct {
	objc.ID
}

func KeyValueSharedObserversSnapshotFrom(ptr unsafe.Pointer) KeyValueSharedObserversSnapshot {
	return KeyValueSharedObserversSnapshot{
		ID: objc.ID(ptr),
	}
}




