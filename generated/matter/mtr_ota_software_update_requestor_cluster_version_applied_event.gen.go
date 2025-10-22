// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass     _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass() _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass {
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass = _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent interface {
	IMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent
	ProductID() foundation.Number
	SetProductID(value foundation.INumber)
	SoftwareVersion() foundation.Number
	SetSoftwareVersion(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn
type MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent struct {
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent
}

// MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	return MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent{
		MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent: MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass) New() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) Init() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent creates a new MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	return getMTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterversionappliedevent-970fn/productid
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterversionappliedevent-970fn/productid
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SetProductID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterversionappliedevent-970fn/softwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterversionappliedevent-970fn/softwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}



