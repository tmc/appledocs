// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterStateCacheContainer] class.
var (
	MTRClusterStateCacheContainerClass     _MTRClusterStateCacheContainerClass
	MTRClusterStateCacheContainerClassOnce sync.Once
)

func getMTRClusterStateCacheContainerClass() _MTRClusterStateCacheContainerClass {
	MTRClusterStateCacheContainerClassOnce.Do(func() {
		MTRClusterStateCacheContainerClass = _MTRClusterStateCacheContainerClass{objc.GetClass("MTRClusterStateCacheContainer")}
	})
	return MTRClusterStateCacheContainerClass
}

type _MTRClusterStateCacheContainerClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterStateCacheContainer] class.
type IMTRClusterStateCacheContainer interface {
	objectivec.IObject
	ReadAttributesWithEndpointIDClusterIDAttributeIDQueueCompletion(endpointID unsafe.Pointer, clusterID unsafe.Pointer, attributeID unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterStateCacheContainer
type MTRClusterStateCacheContainer struct {
	objectivec.Object
}

// MTRClusterStateCacheContainerFrom constructs a [MTRClusterStateCacheContainer] from an unsafe.Pointer.
func MTRClusterStateCacheContainerFrom(ptr unsafe.Pointer) MTRClusterStateCacheContainer {
	return MTRClusterStateCacheContainer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterStateCacheContainerClass) Alloc() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterStateCacheContainerClass) New() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterStateCacheContainer) Init() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterStateCacheContainer) Autorelease() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterStateCacheContainer creates a new MTRClusterStateCacheContainer instance.
func NewMTRClusterStateCacheContainer() MTRClusterStateCacheContainer {
	return getMTRClusterStateCacheContainerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterStateCacheContainer/readAttributes(withEndpointID:clusterID:attributeID:queue:completion:)
func (m_ MTRClusterStateCacheContainer) ReadAttributesWithEndpointIDClusterIDAttributeIDQueueCompletion(endpointID unsafe.Pointer, clusterID unsafe.Pointer, attributeID unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributesWithEndpointID:clusterID:attributeID:queue:completion:"), endpointID, clusterID, attributeID, queue, completion)
}



