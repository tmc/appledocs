// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [TemporalAA] class.
type ITemporalAA interface {
	IKernel
	

	// properties:
	BlendFactor() objectivec.IObject
	SetBlendFactor(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthTexture unsafe.Pointer)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143586-initwithcoder
func NewTemporalAAWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) TemporalAA {
	instance := getTemporalAAClass().Alloc()
	rv := objc.Send[TemporalAA](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143587-initwithdevice
func NewTemporalAAWithDevice(device unsafe.Pointer) TemporalAA {
	instance := getTemporalAAClass().Alloc()
	rv := objc.Send[TemporalAA](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143583-copywithzone
func (t_ TemporalAA) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encode
func (t_ TemporalAA) Encode() {
	objc.Send[objc.ID](t_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encodetocommandbuffer
func (t_ TemporalAA) EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthTexture unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, motionVectorTexture, depthTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143585-encodewithcoder
func (t_ TemporalAA) EncodeWithCoder(coder foundation.Coder) {
	objc.Send[objc.ID](t_.ID, objc.Sel("encodeWithCoder:"), coder)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor
func (t_ TemporalAA) BlendFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("blendFactor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor
func (t_ TemporalAA) SetBlendFactor(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBlendFactor:"), value)
}







