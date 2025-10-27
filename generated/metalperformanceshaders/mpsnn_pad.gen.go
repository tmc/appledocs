// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Pad] class.
var (
	PadClass     _PadClass
	PadClassOnce sync.Once
)

func getPadClass() _PadClass {
	PadClassOnce.Do(func() {
		PadClass = _PadClass{objc.GetClass("MPSNNPad")}
	})
	return PadClass
}

type _PadClass struct {
	class objc.Class
}





// An interface definition for the [Pad] class.
type IPad interface {
	ICNNKernel
	

	// properties:
	FillValue() objectivec.IObject
	SetFillValue(value objectivec.IObject)
	PaddingSizeAfter() ImageCoordinate get set /* not a class type */
	SetPaddingSizeAfter(value ImageCoordinate get set /* not a class type */)
	PaddingSizeBefore() ImageCoordinate get set /* not a class type */
	SetPaddingSizeBefore(value ImageCoordinate get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PadClass) Alloc() Pad {
	rv := objc.Send[Pad](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PadClass) New() Pad {
	rv := objc.Send[Pad](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pad) Init() Pad {
	rv := objc.Send[Pad](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pad) Autorelease() Pad {
	rv := objc.Send[Pad](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPad creates a new Pad instance.
func NewPad() Pad {
	return getPadClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad
type Pad struct {
	CNNKernel
}

// PadFrom constructs a [Pad] from an unsafe.Pointer.
func PadFrom(ptr unsafe.Pointer) Pad {
	return Pad{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037428-initwithcoder
func NewPadWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) Pad {
	instance := getPadClass().Alloc()
	rv := objc.Send[Pad](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037429-initwithdevice
func NewPadWithDevice(device unsafe.Pointer) Pad {
	instance := getPadClass().Alloc()
	rv := objc.Send[Pad](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037430-initwithdevice
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfter(device unsafe.Pointer, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate) Pad {
	instance := getPadClass().Alloc()
	rv := objc.Send[Pad](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:"), device, paddingSizeBefore, paddingSizeAfter)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037431-initwithdevice
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device unsafe.Pointer, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, fillValueArray foundation.Data) Pad {
	instance := getPadClass().Alloc()
	rv := objc.Send[Pad](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:"), device, paddingSizeBefore, paddingSizeAfter, fillValueArray)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037427-fillvalue
func (p_ Pad) FillValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("fillValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037427-fillvalue
func (p_ Pad) SetFillValue(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFillValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037432-paddingsizeafter
func (p_ Pad) PaddingSizeAfter() ImageCoordinate get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("paddingSizeAfter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037432-paddingsizeafter
func (p_ Pad) SetPaddingSizeAfter(value ImageCoordinate get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaddingSizeAfter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037433-paddingsizebefore
func (p_ Pad) PaddingSizeBefore() ImageCoordinate get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("paddingSizeBefore"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037433-paddingsizebefore
func (p_ Pad) SetPaddingSizeBefore(value ImageCoordinate get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaddingSizeBefore:"), value)
}







