// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReplicatorLayer] class.
var (
	replicatorLayerClass     _ReplicatorLayerClass
	replicatorLayerClassOnce sync.Once
)

func getReplicatorLayerClass() _ReplicatorLayerClass {
	replicatorLayerClassOnce.Do(func() {
		replicatorLayerClass = _ReplicatorLayerClass{objc.GetClass("CAReplicatorLayer")}
	})
	return replicatorLayerClass
}

type _ReplicatorLayerClass struct {
	class objc.Class
}

// An interface definition for the [ReplicatorLayer] class.
type IReplicatorLayer interface {
	ILayer
}

// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer
type ReplicatorLayer struct {
	Layer
}

// ReplicatorLayerFrom constructs a [ReplicatorLayer] from an unsafe.Pointer.
//
// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations.
func ReplicatorLayerFrom(ptr unsafe.Pointer) ReplicatorLayer {
	return ReplicatorLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReplicatorLayerClass) Alloc() ReplicatorLayer {
	rv := objc.Send[ReplicatorLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReplicatorLayerClass) New() ReplicatorLayer {
	rv := objc.Send[ReplicatorLayer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReplicatorLayer) Init() ReplicatorLayer {
	rv := objc.Send[ReplicatorLayer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReplicatorLayer) Autorelease() ReplicatorLayer {
	rv := objc.Send[ReplicatorLayer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReplicatorLayer creates a new ReplicatorLayer instance.
func NewReplicatorLayer() ReplicatorLayer {
	return getReplicatorLayerClass().New()
}




