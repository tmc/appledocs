// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ReplicatorLayer] class.
var (
	ReplicatorLayerClass     _ReplicatorLayerClass
	ReplicatorLayerClassOnce sync.Once
)

func getReplicatorLayerClass() _ReplicatorLayerClass {
	ReplicatorLayerClassOnce.Do(func() {
		ReplicatorLayerClass = _ReplicatorLayerClass{objc.GetClass("CAReplicatorLayer")}
	})
	return ReplicatorLayerClass
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


// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceAlphaOffset
func (r_ ReplicatorLayer) InstanceAlphaOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceAlphaOffset"))
	return rv
}


// SetInstanceAlphaOffset sets the value of the instanceAlphaOffset property.
// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceAlphaOffset
func (r_ ReplicatorLayer) SetInstanceAlphaOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceAlphaOffset:"), value)
}
// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceBlueOffset
func (r_ ReplicatorLayer) InstanceBlueOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceBlueOffset"))
	return rv
}


// SetInstanceBlueOffset sets the value of the instanceBlueOffset property.
// Defines the offset added to the blue component of the color for each replicated instance. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceBlueOffset
func (r_ ReplicatorLayer) SetInstanceBlueOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceBlueOffset:"), value)
}
// Defines the color used to multiply the source object. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceColor
func (r_ ReplicatorLayer) InstanceColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](r_.ID, objc.Sel("instanceColor"))
	return rv
}


// SetInstanceColor sets the value of the instanceColor property.
// Defines the color used to multiply the source object. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceColor
func (r_ ReplicatorLayer) SetInstanceColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceColor:"), value)
}
// The number of copies to create, including the source layers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceCount
func (r_ ReplicatorLayer) InstanceCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("instanceCount"))
	return rv
}


// SetInstanceCount sets the value of the instanceCount property.
// The number of copies to create, including the source layers.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceCount
func (r_ ReplicatorLayer) SetInstanceCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceCount:"), value)
}
// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceDelay
func (r_ ReplicatorLayer) InstanceDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("instanceDelay"))
	return rv
}


// SetInstanceDelay sets the value of the instanceDelay property.
// Specifies the delay, in seconds, between replicated copies. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceDelay
func (r_ ReplicatorLayer) SetInstanceDelay(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceDelay:"), value)
}
// Defines the offset added to the green component of the color for each replicated instance. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceGreenOffset
func (r_ ReplicatorLayer) InstanceGreenOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceGreenOffset"))
	return rv
}


// SetInstanceGreenOffset sets the value of the instanceGreenOffset property.
// Defines the offset added to the green component of the color for each replicated instance. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceGreenOffset
func (r_ ReplicatorLayer) SetInstanceGreenOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceGreenOffset:"), value)
}
// Defines the offset added to the red component of the color for each replicated instance. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceRedOffset
func (r_ ReplicatorLayer) InstanceRedOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceRedOffset"))
	return rv
}


// SetInstanceRedOffset sets the value of the instanceRedOffset property.
// Defines the offset added to the red component of the color for each replicated instance. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceRedOffset
func (r_ ReplicatorLayer) SetInstanceRedOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceRedOffset:"), value)
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
// Defines whether this layer flattens its sublayers into its plane.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/preservesDepth
func (r_ ReplicatorLayer) PreservesDepth() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("preservesDepth"))
	return rv
}


// SetPreservesDepth sets the value of the preservesDepth property.
// Defines whether this layer flattens its sublayers into its plane.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/preservesDepth
func (r_ ReplicatorLayer) SetPreservesDepth(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreservesDepth:"), value)
}


