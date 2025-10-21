// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKLookAroundSnapshotter] class.
var (
	MKLookAroundSnapshotterClass     _MKLookAroundSnapshotterClass
	MKLookAroundSnapshotterClassOnce sync.Once
)

func getMKLookAroundSnapshotterClass() _MKLookAroundSnapshotterClass {
	MKLookAroundSnapshotterClassOnce.Do(func() {
		MKLookAroundSnapshotterClass = _MKLookAroundSnapshotterClass{objc.GetClass("MKLookAroundSnapshotter")}
	})
	return MKLookAroundSnapshotterClass
}

type _MKLookAroundSnapshotterClass struct {
	class objc.Class
}

// An interface definition for the [MKLookAroundSnapshotter] class.
type IMKLookAroundSnapshotter interface {
	objectivec.IObject
}

// A utility class that you use to create a static image from a LookAround scene.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter
type MKLookAroundSnapshotter struct {
	objectivec.Object
}

// MKLookAroundSnapshotterFrom constructs a [MKLookAroundSnapshotter] from an unsafe.Pointer.
//
// A utility class that you use to create a static image from a LookAround scene.
func MKLookAroundSnapshotterFrom(ptr unsafe.Pointer) MKLookAroundSnapshotter {
	return MKLookAroundSnapshotter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSnapshotterClass) Alloc() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKLookAroundSnapshotterClass) New() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSnapshotter) Init() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSnapshotter) Autorelease() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSnapshotter creates a new MKLookAroundSnapshotter instance.
func NewMKLookAroundSnapshotter() MKLookAroundSnapshotter {
	return getMKLookAroundSnapshotterClass().New()
}




// Create a new snapshotter object with the scene and options you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/init(scene:options:)
func NewMKLookAroundSnapshotterWithSceneOptions(scene unsafe.Pointer, options unsafe.Pointer) MKLookAroundSnapshotter {
	instance := getMKLookAroundSnapshotterClass().Alloc()
	rv := objc.Send[MKLookAroundSnapshotter](instance.ID, objc.Sel("initWithScene:options:"), scene, options)
	rv.Autorelease()
	return rv
}



