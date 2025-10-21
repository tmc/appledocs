// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [KeyValueSharedObserversSnapshot] class.
var (
	KeyValueSharedObserversSnapshotClass     _KeyValueSharedObserversSnapshotClass
	KeyValueSharedObserversSnapshotClassOnce sync.Once
)

func getKeyValueSharedObserversSnapshotClass() _KeyValueSharedObserversSnapshotClass {
	KeyValueSharedObserversSnapshotClassOnce.Do(func() {
		KeyValueSharedObserversSnapshotClass = _KeyValueSharedObserversSnapshotClass{objc.GetClass("NSKeyValueSharedObserversSnapshot")}
	})
	return KeyValueSharedObserversSnapshotClass
}

type _KeyValueSharedObserversSnapshotClass struct {
	class objc.Class
}

// An interface definition for the [KeyValueSharedObserversSnapshot] class.
type IKeyValueSharedObserversSnapshot interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (kc _KeyValueSharedObserversSnapshotClass) Alloc() KeyValueSharedObserversSnapshot {
	rv := objc.Send[KeyValueSharedObserversSnapshot](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyValueSharedObserversSnapshotClass) New() KeyValueSharedObserversSnapshot {
	rv := objc.Send[KeyValueSharedObserversSnapshot](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyValueSharedObserversSnapshot) Init() KeyValueSharedObserversSnapshot {
	rv := objc.Send[KeyValueSharedObserversSnapshot](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyValueSharedObserversSnapshot) Autorelease() KeyValueSharedObserversSnapshot {
	rv := objc.Send[KeyValueSharedObserversSnapshot](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyValueSharedObserversSnapshot creates a new KeyValueSharedObserversSnapshot instance.
func NewKeyValueSharedObserversSnapshot() KeyValueSharedObserversSnapshot {
	return getKeyValueSharedObserversSnapshotClass().New()
}




