// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MKLookAroundSceneRequest] class.
var (
	MKLookAroundSceneRequestClass     _MKLookAroundSceneRequestClass
	MKLookAroundSceneRequestClassOnce sync.Once
)

func getMKLookAroundSceneRequestClass() _MKLookAroundSceneRequestClass {
	MKLookAroundSceneRequestClassOnce.Do(func() {
		MKLookAroundSceneRequestClass = _MKLookAroundSceneRequestClass{objc.GetClass("MKLookAroundSceneRequest")}
	})
	return MKLookAroundSceneRequestClass
}

type _MKLookAroundSceneRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKLookAroundSceneRequest] class.
type IMKLookAroundSceneRequest interface {
	objectivec.IObject
}

// A class you use to request a LookAround scene at the location you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest
type MKLookAroundSceneRequest struct {
	objectivec.Object
}

// MKLookAroundSceneRequestFrom constructs a [MKLookAroundSceneRequest] from an unsafe.Pointer.
//
// A class you use to request a LookAround scene at the location you specify.
func MKLookAroundSceneRequestFrom(ptr unsafe.Pointer) MKLookAroundSceneRequest {
	return MKLookAroundSceneRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSceneRequestClass) Alloc() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKLookAroundSceneRequestClass) New() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSceneRequest) Init() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSceneRequest) Autorelease() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSceneRequest creates a new MKLookAroundSceneRequest instance.
func NewMKLookAroundSceneRequest() MKLookAroundSceneRequest {
	return getMKLookAroundSceneRequestClass().New()
}




