// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MKLookAroundScene] class.
var (
	MKLookAroundSceneClass     _MKLookAroundSceneClass
	MKLookAroundSceneClassOnce sync.Once
)

func getMKLookAroundSceneClass() _MKLookAroundSceneClass {
	MKLookAroundSceneClassOnce.Do(func() {
		MKLookAroundSceneClass = _MKLookAroundSceneClass{objc.GetClass("MKLookAroundScene")}
	})
	return MKLookAroundSceneClass
}

type _MKLookAroundSceneClass struct {
	class objc.Class
}

// An interface definition for the [MKLookAroundScene] class.
type IMKLookAroundScene interface {
	objectivec.IObject
}

// A utility class that encapsulates information the framework requires to retrieve and display a specific Look Around location’s imagery.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundScene
type MKLookAroundScene struct {
	objectivec.Object
}

// MKLookAroundSceneFrom constructs a [MKLookAroundScene] from an unsafe.Pointer.
//
// A utility class that encapsulates information the framework requires to retrieve and display a specific Look Around location’s imagery.
func MKLookAroundSceneFrom(ptr unsafe.Pointer) MKLookAroundScene {
	return MKLookAroundScene{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSceneClass) Alloc() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKLookAroundSceneClass) New() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundScene) Init() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundScene) Autorelease() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundScene creates a new MKLookAroundScene instance.
func NewMKLookAroundScene() MKLookAroundScene {
	return getMKLookAroundSceneClass().New()
}




