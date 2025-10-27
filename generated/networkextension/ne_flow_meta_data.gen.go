// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEFlowMetaData] class.
var (
	NEFlowMetaDataClass     _NEFlowMetaDataClass
	NEFlowMetaDataClassOnce sync.Once
)

func getNEFlowMetaDataClass() _NEFlowMetaDataClass {
	NEFlowMetaDataClassOnce.Do(func() {
		NEFlowMetaDataClass = _NEFlowMetaDataClass{objc.GetClass("NEFlowMetaData")}
	})
	return NEFlowMetaDataClass
}

type _NEFlowMetaDataClass struct {
	class objc.Class
}





// An interface definition for the [NEFlowMetaData] class.
type INEFlowMetaData interface {
	objectivec.IObject
	

	// properties:
	FilterFlowIdentifier() foundation.UUID
	SourceAppAuditToken() foundation.foundation.INSData
	SourceAppSigningIdentifier() foundation.foundation.INSString
	SourceAppUniqueIdentifier() foundation.foundation.INSData
	RoutingMethod() NETunnelProviderRoutingMethod
	SetRoutingMethod(value NETunnelProviderRoutingMethod)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFlowMetaDataClass) Alloc() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFlowMetaDataClass) New() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFlowMetaData) Init() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFlowMetaData) Autorelease() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFlowMetaData creates a new NEFlowMetaData instance.
func NewNEFlowMetaData() NEFlowMetaData {
	return getNEFlowMetaDataClass().New()
}





// Additional information about data flowing through a per-app VPN provider.
//
// This metadata is only present for data flowing through per-app VPN providers, that is, app proxy providers and packet tunnel providers in per-app VPN mode, as indicated by the property.


// Additional information about data flowing through a per-app VPN provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData
type NEFlowMetaData struct {
	objectivec.Object
}

// NEFlowMetaDataFrom constructs a [NEFlowMetaData] from an unsafe.Pointer.
//
// Additional information about data flowing through a per-app VPN provider.
func NEFlowMetaDataFrom(ptr unsafe.Pointer) NEFlowMetaData {
	return NEFlowMetaData{objectivec.Object{objc.ID(ptr)}}
}

























// The identifier of the content filter flow corresponding to this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/filterFlowIdentifier
func (n_ NEFlowMetaData) FilterFlowIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](n_.ID, objc.Sel("filterFlowIdentifier"))
	return rv
}


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppAuditToken
func (n_ NEFlowMetaData) SourceAppAuditToken() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}


// A string that contains the signing identifier of the source application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppSigningIdentifier
func (n_ NEFlowMetaData) SourceAppSigningIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppSigningIdentifier"))
	return rv
}


// A data instance that contains a unique hash value for the source application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppUniqueIdentifier
func (n_ NEFlowMetaData) SourceAppUniqueIdentifier() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return rv
}


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NEFlowMetaData) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NEFlowMetaData) SetRoutingMethod(value NETunnelProviderRoutingMethod) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoutingMethod:"), value)
}








