// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotHelperResponse] class.
var (
	NEHotspotHelperResponseClass     _NEHotspotHelperResponseClass
	NEHotspotHelperResponseClassOnce sync.Once
)

func getNEHotspotHelperResponseClass() _NEHotspotHelperResponseClass {
	NEHotspotHelperResponseClassOnce.Do(func() {
		NEHotspotHelperResponseClass = _NEHotspotHelperResponseClass{objc.GetClass("NEHotspotHelperResponse")}
	})
	return NEHotspotHelperResponseClass
}

type _NEHotspotHelperResponseClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotHelperResponse] class.
type INEHotspotHelperResponse interface {
	objectivec.IObject
	SetNetwork(network unsafe.Pointer)
}

// The hotspot helper’s response to a command.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResponse
type NEHotspotHelperResponse struct {
	objectivec.Object
}

// NEHotspotHelperResponseFrom constructs a [NEHotspotHelperResponse] from an unsafe.Pointer.
//
// The hotspot helper’s response to a command.
func NEHotspotHelperResponseFrom(ptr unsafe.Pointer) NEHotspotHelperResponse {
	return NEHotspotHelperResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperResponseClass) Alloc() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotHelperResponseClass) New() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelperResponse) Init() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelperResponse) Autorelease() NEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelperResponse creates a new NEHotspotHelperResponse instance.
func NewNEHotspotHelperResponse() NEHotspotHelperResponse {
	return getNEHotspotHelperResponseClass().New()
}


// Set the network that conveys the confidence level.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResponse/setNetwork(_:)
func (n_ NEHotspotHelperResponse) SetNetwork(network unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetwork:"), network)
}



