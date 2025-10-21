// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPowerTopology] class.
var (
	MTRClusterPowerTopologyClass     _MTRClusterPowerTopologyClass
	MTRClusterPowerTopologyClassOnce sync.Once
)

func getMTRClusterPowerTopologyClass() _MTRClusterPowerTopologyClass {
	MTRClusterPowerTopologyClassOnce.Do(func() {
		MTRClusterPowerTopologyClass = _MTRClusterPowerTopologyClass{objc.GetClass("MTRClusterPowerTopology")}
	})
	return MTRClusterPowerTopologyClass
}

type _MTRClusterPowerTopologyClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPowerTopology] class.
type IMTRClusterPowerTopology interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeActiveEndpointsWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAvailableEndpointsWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
}

// Cluster Power Topology The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology
type MTRClusterPowerTopology struct {
	MTRGenericCluster
}

// MTRClusterPowerTopologyFrom constructs a [MTRClusterPowerTopology] from an unsafe.Pointer.
//
// Cluster Power Topology The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.
func MTRClusterPowerTopologyFrom(ptr unsafe.Pointer) MTRClusterPowerTopology {
	return MTRClusterPowerTopology{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPowerTopologyClass) Alloc() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPowerTopologyClass) New() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPowerTopology) Init() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPowerTopology) Autorelease() MTRClusterPowerTopology {
	rv := objc.Send[MTRClusterPowerTopology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPowerTopology creates a new MTRClusterPowerTopology instance.
func NewMTRClusterPowerTopology() MTRClusterPowerTopology {
	return getMTRClusterPowerTopologyClass().New()
}


// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/init(device:endpointID:queue:)
func NewMTRClusterPowerTopologyWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterPowerTopology {
	instance := getMTRClusterPowerTopologyClass().Alloc()
	rv := objc.Send[MTRClusterPowerTopology](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeActiveEndpoints(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeActiveEndpointsWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeActiveEndpointsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeAttributeList(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeAvailableEndpoints(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeAvailableEndpointsWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAvailableEndpointsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeClusterRevision(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeFeatureMap(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerTopology/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterPowerTopology) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


