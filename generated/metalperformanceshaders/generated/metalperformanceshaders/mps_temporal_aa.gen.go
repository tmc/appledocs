// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTemporalAA */


/* debug [class_header]: Header for MPSTemporalAA */
// The class instance for the [TemporalAA] class.
var (
	TemporalAAClass     _TemporalAAClass
	TemporalAAClassOnce sync.Once
)

func getTemporalAAClass() _TemporalAAClass {
	TemporalAAClassOnce.Do(func() {
		TemporalAAClass = _TemporalAAClass{objc.GetClass("MPSTemporalAA")}
	})
	return TemporalAAClass
}

type _TemporalAAClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TemporalAA */
// An interface definition for the [TemporalAA] class.
type ITemporalAA interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for TemporalAA */
	// properties:
	BlendFactor() objectivec.IObject
	SetBlendFactor(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TemporalAA */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthTexture unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TemporalAA */
// Alloc allocates a new instance without initialization.
func (tc _TemporalAAClass) Alloc() TemporalAA {
	rv := objc.Send[TemporalAA](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TemporalAAClass) New() TemporalAA {
	rv := objc.Send[TemporalAA](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporalAA) Init() TemporalAA {
	rv := objc.Send[TemporalAA](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporalAA) Autorelease() TemporalAA {
	rv := objc.Send[TemporalAA](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporalAA creates a new TemporalAA instance.
func NewTemporalAA() TemporalAA {
	return getTemporalAAClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TemporalAA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA
type TemporalAA struct {
	Kernel
}

// TemporalAAFrom constructs a [TemporalAA] from an unsafe.Pointer.
func TemporalAAFrom(ptr unsafe.Pointer) TemporalAA {
	return TemporalAA{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TemporalAA */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143586-initwithcoder
func NewTemporalAAWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) TemporalAA {
	instance := getTemporalAAClass().Alloc()
	rv := objc.Send[TemporalAA](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTemporalAAWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143587-initwithdevice
func NewTemporalAAWithDevice(device unsafe.Pointer) TemporalAA {
	instance := getTemporalAAClass().Alloc()
	rv := objc.Send[TemporalAA](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTemporalAAWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TemporalAA */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TemporalAA */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TemporalAA */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143583-copywithzone
func (t_ TemporalAA) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encode
func (t_ TemporalAA) Encode() {
	objc.Send[objc.ID](t_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encodetocommandbuffer
func (t_ TemporalAA) EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthTexture unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, motionVectorTexture, depthTexture)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143585-encodewithcoder
func (t_ TemporalAA) EncodeWithCoder(coder Coder /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("encodeWithCoder:"), coder)
}/* debug [instance_methods/method]: EncodeWithCoder */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TemporalAA */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor
func (t_ TemporalAA) BlendFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("blendFactor"))
	return rv
}/* debug [instance_properties/getter]: blendFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor
func (t_ TemporalAA) SetBlendFactor(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBlendFactor:"), value)
}/* debug [instance_properties/setter]: blendFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTemporalAA */


