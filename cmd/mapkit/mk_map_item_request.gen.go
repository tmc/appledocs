// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapItemRequest] class.
var (
	MKMapItemRequestClass     _MKMapItemRequestClass
	MKMapItemRequestClassOnce sync.Once
)

func getMKMapItemRequestClass() _MKMapItemRequestClass {
	MKMapItemRequestClassOnce.Do(func() {
		MKMapItemRequestClass = _MKMapItemRequestClass{objc.GetClass("MKMapItemRequest")}
	})
	return MKMapItemRequestClass
}

type _MKMapItemRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItemRequest] class.
type IMKMapItemRequest interface {
	objectivec.IObject
}

// A utility class you use to request additional information about a map feature.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest
type MKMapItemRequest struct {
	objectivec.Object
}

// MKMapItemRequestFrom constructs a [MKMapItemRequest] from an unsafe.Pointer.
//
// A utility class you use to request additional information about a map feature.
func MKMapItemRequestFrom(ptr unsafe.Pointer) MKMapItemRequest {
	return MKMapItemRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemRequestClass) Alloc() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemRequestClass) New() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemRequest) Init() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemRequest) Autorelease() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemRequest creates a new MKMapItemRequest instance.
func NewMKMapItemRequest() MKMapItemRequest {
	return getMKMapItemRequestClass().New()
}




