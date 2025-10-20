// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var KeyValueSharedObserversSnapshotClass _KeyValueSharedObserversSnapshotClass

func init() {
	KeyValueSharedObserversSnapshotClass = _KeyValueSharedObserversSnapshotClass{objc.GetClass("NSKeyValueSharedObserversSnapshot")}
}

type _KeyValueSharedObserversSnapshotClass struct {
	class objc.Class
}

type KeyValueSharedObserversSnapshot struct {
	objc.ID
}

func KeyValueSharedObserversSnapshotFrom(ptr unsafe.Pointer) KeyValueSharedObserversSnapshot {
	return KeyValueSharedObserversSnapshot{
		ID: objc.ID(ptr),
	}
}




