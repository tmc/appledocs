// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CAReplicatorLayer */


/* debug [class_header]: Header for CAReplicatorLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReplicatorLayer */
// An interface definition for the [ReplicatorLayer] class.
type IReplicatorLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for ReplicatorLayer */
	// properties:
	InstanceAlphaOffset() float32
	SetInstanceAlphaOffset(value float32)
	InstanceBlueOffset() float32
	SetInstanceBlueOffset(value float32)
	InstanceColor() ColorRef /* not a class type */
	SetInstanceColor(value ColorRef /* not a class type */)
	InstanceCount() int
	SetInstanceCount(value int)
	InstanceDelay() float64
	SetInstanceDelay(value float64)
	InstanceGreenOffset() float32
	SetInstanceGreenOffset(value float32)
	InstanceRedOffset() float32
	SetInstanceRedOffset(value float32)
	InstanceTransform() objc.IObject /* cross-framework: CATransform3D */
	SetInstanceTransform(value objc.IObject /* cross-framework: CATransform3D */)
	PreservesDepth() bool
	SetPreservesDepth(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReplicatorLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReplicatorLayer */
// Alloc allocates a new instance without initialization.
func (rc _ReplicatorLayerClass) Alloc() ReplicatorLayer {
	rv := objc.Send[ReplicatorLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReplicatorLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReplicatorLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReplicatorLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReplicatorLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReplicatorLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReplicatorLayer */

// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceAlphaOffset
func (r_ ReplicatorLayer) InstanceAlphaOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceAlphaOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceAlphaOffset */


// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceAlphaOffset
func (r_ ReplicatorLayer) SetInstanceAlphaOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceAlphaOffset:"), value)
}/* debug [instance_properties/setter]: instanceAlphaOffset */


// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceBlueOffset
func (r_ ReplicatorLayer) InstanceBlueOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceBlueOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceBlueOffset */


// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceBlueOffset
func (r_ ReplicatorLayer) SetInstanceBlueOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceBlueOffset:"), value)
}/* debug [instance_properties/setter]: instanceBlueOffset */


// Defines the color used to multiply the source object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceColor
func (r_ ReplicatorLayer) InstanceColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](r_.ID, objc.Sel("instanceColor"))
	return rv
}/* debug [instance_properties/getter]: instanceColor */


// Defines the color used to multiply the source object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceColor
func (r_ ReplicatorLayer) SetInstanceColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceColor:"), value)
}/* debug [instance_properties/setter]: instanceColor */


// The number of copies to create, including the source layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceCount
func (r_ ReplicatorLayer) InstanceCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("instanceCount"))
	return rv
}/* debug [instance_properties/getter]: instanceCount */


// The number of copies to create, including the source layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceCount
func (r_ ReplicatorLayer) SetInstanceCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceCount:"), value)
}/* debug [instance_properties/setter]: instanceCount */


// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceDelay
func (r_ ReplicatorLayer) InstanceDelay() float64 {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("instanceDelay"))
	return rv
}/* debug [instance_properties/getter]: instanceDelay */


// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceDelay
func (r_ ReplicatorLayer) SetInstanceDelay(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceDelay:"), value)
}/* debug [instance_properties/setter]: instanceDelay */


// Defines the offset added to the green component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceGreenOffset
func (r_ ReplicatorLayer) InstanceGreenOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceGreenOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceGreenOffset */


// Defines the offset added to the green component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceGreenOffset
func (r_ ReplicatorLayer) SetInstanceGreenOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceGreenOffset:"), value)
}/* debug [instance_properties/setter]: instanceGreenOffset */


// Defines the offset added to the red component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceRedOffset
func (r_ ReplicatorLayer) InstanceRedOffset() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("instanceRedOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceRedOffset */


// Defines the offset added to the red component of the color for each replicated instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceRedOffset
func (r_ ReplicatorLayer) SetInstanceRedOffset(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceRedOffset:"), value)
}/* debug [instance_properties/setter]: instanceRedOffset */


// The transform matrix applied to the previous instance to produce the current instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) InstanceTransform() objc.IObject /* cross-framework: CATransform3D */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("instanceTransform"))
	return rv
}/* debug [instance_properties/getter]: instanceTransform */


// The transform matrix applied to the previous instance to produce the current instance. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r_ ReplicatorLayer) SetInstanceTransform(value objc.IObject /* cross-framework: CATransform3D */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInstanceTransform:"), value)
}/* debug [instance_properties/setter]: instanceTransform */


// Defines whether this layer flattens its sublayers into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/preservesDepth
func (r_ ReplicatorLayer) PreservesDepth() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("preservesDepth"))
	return rv
}/* debug [instance_properties/getter]: preservesDepth */


// Defines whether this layer flattens its sublayers into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/preservesDepth
func (r_ ReplicatorLayer) SetPreservesDepth(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreservesDepth:"), value)
}/* debug [instance_properties/setter]: preservesDepth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAReplicatorLayer */



