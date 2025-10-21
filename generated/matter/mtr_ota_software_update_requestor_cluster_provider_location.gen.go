// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterProviderLocation] class.
var (
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClass     _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterProviderLocationClass() _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass {
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterProviderLocationClass = _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterProviderLocation")}
	})
	return MTROtaSoftwareUpdateRequestorClusterProviderLocationClass
}

type _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterProviderLocation] class.
type IMTROtaSoftwareUpdateRequestorClusterProviderLocation interface {
	IMTROTASoftwareUpdateRequestorClusterProviderLocation
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm
type MTROtaSoftwareUpdateRequestorClusterProviderLocation struct {
	MTROTASoftwareUpdateRequestorClusterProviderLocation
}

// MTROtaSoftwareUpdateRequestorClusterProviderLocationFrom constructs a [MTROtaSoftwareUpdateRequestorClusterProviderLocation] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterProviderLocationFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	return MTROtaSoftwareUpdateRequestorClusterProviderLocation{
		MTROTASoftwareUpdateRequestorClusterProviderLocation: MTROTASoftwareUpdateRequestorClusterProviderLocationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass) Alloc() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass) New() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Init() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Autorelease() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterProviderLocation creates a new MTROtaSoftwareUpdateRequestorClusterProviderLocation instance.
func NewMTROtaSoftwareUpdateRequestorClusterProviderLocation() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	return getMTROtaSoftwareUpdateRequestorClusterProviderLocationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetEndpoint(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/fabricindex
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/fabricindex
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/providernodeid
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) ProviderNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("providerNodeID"))
	return rv
}


// SetProviderNodeID sets the value of the providerNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/providernodeid
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetProviderNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}



