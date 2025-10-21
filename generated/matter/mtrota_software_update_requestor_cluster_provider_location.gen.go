// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterProviderLocation] class.
var (
	MTROTASoftwareUpdateRequestorClusterProviderLocationClass     _MTROTASoftwareUpdateRequestorClusterProviderLocationClass
	MTROTASoftwareUpdateRequestorClusterProviderLocationClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterProviderLocationClass() _MTROTASoftwareUpdateRequestorClusterProviderLocationClass {
	MTROTASoftwareUpdateRequestorClusterProviderLocationClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterProviderLocationClass = _MTROTASoftwareUpdateRequestorClusterProviderLocationClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterProviderLocation")}
	})
	return MTROTASoftwareUpdateRequestorClusterProviderLocationClass
}

type _MTROTASoftwareUpdateRequestorClusterProviderLocationClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterProviderLocation] class.
type IMTROTASoftwareUpdateRequestorClusterProviderLocation interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq
type MTROTASoftwareUpdateRequestorClusterProviderLocation struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterProviderLocationFrom constructs a [MTROTASoftwareUpdateRequestorClusterProviderLocation] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterProviderLocationFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterProviderLocation {
	return MTROTASoftwareUpdateRequestorClusterProviderLocation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterProviderLocationClass) Alloc() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterProviderLocationClass) New() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Init() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Autorelease() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterProviderLocation creates a new MTROTASoftwareUpdateRequestorClusterProviderLocation instance.
func NewMTROTASoftwareUpdateRequestorClusterProviderLocation() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	return getMTROTASoftwareUpdateRequestorClusterProviderLocationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/fabricindex
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/fabricindex
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/providernodeid
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) ProviderNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("providerNodeID"))
	return rv
}


// SetProviderNodeID sets the value of the providerNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterproviderlocation-76vsq/providernodeid
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetProviderNodeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}



