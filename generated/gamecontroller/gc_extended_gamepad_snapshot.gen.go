// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GCExtendedGamepadSnapshot] class.
var (
	GCExtendedGamepadSnapshotClass     _GCExtendedGamepadSnapshotClass
	GCExtendedGamepadSnapshotClassOnce sync.Once
)

func getGCExtendedGamepadSnapshotClass() _GCExtendedGamepadSnapshotClass {
	GCExtendedGamepadSnapshotClassOnce.Do(func() {
		GCExtendedGamepadSnapshotClass = _GCExtendedGamepadSnapshotClass{objc.GetClass("GCExtendedGamepadSnapshot")}
	})
	return GCExtendedGamepadSnapshotClass
}

type _GCExtendedGamepadSnapshotClass struct {
	class objc.Class
}

// An interface definition for the [GCExtendedGamepadSnapshot] class.
type IGCExtendedGamepadSnapshot interface {
	IGCExtendedGamepad
	GCCurrentExtendedGamepadSnapshotDataVersion() unsafe.Pointer
	GCCurrentMicroGamepadSnapshotDataVersion() unsafe.Pointer
	SnapshotData() foundation.Data
	SetSnapshotData(value foundation.IData)
}

// A recording of all of the values provided by a object.
//
// To create a gamepad snapshot, call the method on a object. The class is a subclass of the class, so you use the parent class’s properties to read the individual element values. The snapshot is stored in a device independent format. To get the flattened data representation of the snapshot data, read the property.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot
type GCExtendedGamepadSnapshot struct {
	GCExtendedGamepad
}

// GCExtendedGamepadSnapshotFrom constructs a [GCExtendedGamepadSnapshot] from an unsafe.Pointer.
//
// A recording of all of the values provided by a object.
func GCExtendedGamepadSnapshotFrom(ptr unsafe.Pointer) GCExtendedGamepadSnapshot {
	return GCExtendedGamepadSnapshot{
		GCExtendedGamepad: GCExtendedGamepadFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCExtendedGamepadSnapshotClass) Alloc() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCExtendedGamepadSnapshotClass) New() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCExtendedGamepadSnapshot) Init() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCExtendedGamepadSnapshot) Autorelease() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCExtendedGamepadSnapshot creates a new GCExtendedGamepadSnapshot instance.
func NewGCExtendedGamepadSnapshot() GCExtendedGamepadSnapshot {
	return getGCExtendedGamepadSnapshotClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentextendedgamepadsnapshotdataversion
func (g_ GCExtendedGamepadSnapshot) GCCurrentExtendedGamepadSnapshotDataVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("GCCurrentExtendedGamepadSnapshotDataVersion"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentmicrogamepadsnapshotdataversion
func (g_ GCExtendedGamepadSnapshot) GCCurrentMicroGamepadSnapshotDataVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("GCCurrentMicroGamepadSnapshotDataVersion"))
	return rv
}

// Flattens a snapshot into an archivable memory representation.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepadsnapshot/snapshotdata
func (g_ GCExtendedGamepadSnapshot) SnapshotData() foundation.Data {
	rv := objc.Send[foundation.Data](g_.ID, objc.Sel("snapshotData"))
	return rv
}


// SetSnapshotData sets the value of the snapshotData property.
// Flattens a snapshot into an archivable memory representation.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepadsnapshot/snapshotdata
func (g_ GCExtendedGamepadSnapshot) SetSnapshotData(value foundation.IData) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSnapshotData:"), value)
}



