// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	InstanceTransform() CATransform3D /* not a class type */
	SetInstanceTransform(value CATransform3D /* not a class type */)
	InstanceAlphaOffset() float32
	SetInstanceAlphaOffset(value float32)
	InstanceBlueOffset() float32
	SetInstanceBlueOffset(value float32)
	InstanceColor() objectivec.IObject
	SetInstanceColor(value objectivec.IObject)
	InstanceCount() int
	SetInstanceCount(value int)
	InstanceDelay() float64
	SetInstanceDelay(value float64)
	InstanceGreenOffset() float32
	SetInstanceGreenOffset(value float32)
	InstanceRedOffset() float32
	SetInstanceRedOffset(value float32)
	PreservesDepth() bool
	SetPreservesDepth(value bool)
	// methods:
}

// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations.
//
// You can use a object to build complex layouts based on a single source layer that is replicated with transformation rules that can affect the position, rotation color, and time. The following shows a simple example: a red square is added to a replicator layer with an instance count of . The position of each replicated instance is offset along the axis so that it appears to the right of the previous instance. The blue and green color channels are offset so that their values reach at the final instance. The result of the code above is a row of five squares, with colors graduating from white to red. Replicator layers can be nested. The following code adds to a second replicator layer that offsets the position of each instance vertically and subtracts from the red channel. The result of adding this code is to create a grid with the value of the red channel being reduced in the vertical direction.


// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) InstanceTransform() CATransform3D /* not a class type */ {
	rv := objc.Send[Transform3D](r_.ID, objc.Sel("instanceTransform"))
	return rv
}


// The transform matrix applied to the previous instance to produce the current instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) SetInstanceTransform(value CATransform3D /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceTransform:"), value)
}


// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancealphaoffset
func (r_ ReplicatorLayer) InstanceAlphaOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceAlphaOffset"))
	return rv
}


// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancealphaoffset
func (r_ ReplicatorLayer) SetInstanceAlphaOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceAlphaOffset:"), value)
}


// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instanceblueoffset
func (r_ ReplicatorLayer) InstanceBlueOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceBlueOffset"))
	return rv
}


// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instanceblueoffset
func (r_ ReplicatorLayer) SetInstanceBlueOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceBlueOffset:"), value)
}


// Defines the color used to multiply the source object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancecolor
func (r_ ReplicatorLayer) InstanceColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("instanceColor"))
	return rv
}


// Defines the color used to multiply the source object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancecolor
func (r_ ReplicatorLayer) SetInstanceColor(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceColor:"), value)
}


// The number of copies to create, including the source layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancecount
func (r_ ReplicatorLayer) InstanceCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("instanceCount"))
	return rv
}


// The number of copies to create, including the source layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancecount
func (r_ ReplicatorLayer) SetInstanceCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceCount:"), value)
}


// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancedelay
func (r_ ReplicatorLayer) InstanceDelay() float64 {
	rv := objc.Send[TimeInterval](r_.ID, objc.Sel("instanceDelay"))
	return rv
}


// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancedelay
func (r_ ReplicatorLayer) SetInstanceDelay(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceDelay:"), value)
}


// Defines the offset added to the green component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancegreenoffset
func (r_ ReplicatorLayer) InstanceGreenOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceGreenOffset"))
	return rv
}


// Defines the offset added to the green component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instancegreenoffset
func (r_ ReplicatorLayer) SetInstanceGreenOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceGreenOffset:"), value)
}


// Defines the offset added to the red component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instanceredoffset
func (r_ ReplicatorLayer) InstanceRedOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceRedOffset"))
	return rv
}


// Defines the offset added to the red component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/instanceredoffset
func (r_ ReplicatorLayer) SetInstanceRedOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceRedOffset:"), value)
}


// Defines whether this layer flattens its sublayers into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/preservesdepth
func (r_ ReplicatorLayer) PreservesDepth() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("preservesDepth"))
	return rv
}


// Defines whether this layer flattens its sublayers into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/careplicatorlayer/preservesdepth
func (r_ ReplicatorLayer) SetPreservesDepth(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreservesDepth:"), value)
}



