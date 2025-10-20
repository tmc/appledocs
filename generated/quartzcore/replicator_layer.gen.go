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
// You can use a object to build complex layouts based on a single source layer that is replicated with transformation rules that can affect the position, rotation color, and time. The following shows a simple example: a red square is added to a replicator layer with an instance count of . The position of each replicated instance is offset along the axis so that it appears to the right of the previous instance. The blue and green color channels are offset so that their values reach at the final instance. The result of the code above is a row of five squares, with colors graduating from white to red. Replicator layers can be nested. The following code adds to a second replicator layer that offsets the position of each instance vertically and subtracts from the red channel. The result of adding this code is to create a grid with the value of the red channel being reduced in the vertical direction.
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


// The transform matrix applied to the previous instance to produce the current instance. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) InstanceTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("instanceTransform"))
	return rv
}

// SetInstanceTransform sets the value of the instanceTransform property.
// The transform matrix applied to the previous instance to produce the current instance. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) SetInstanceTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceTransform:"), value)
}


