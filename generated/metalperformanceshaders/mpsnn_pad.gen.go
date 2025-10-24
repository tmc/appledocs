// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	PaddingSizeAfter() ImageCoordinate /* not a class type */
	SetPaddingSizeAfter(value ImageCoordinate /* not a class type */)
	PaddingSizeBefore() ImageCoordinate /* not a class type */
	SetPaddingSizeBefore(value ImageCoordinate /* not a class type */)
	FillValue() float32
	SetFillValue(value float32)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (pc _PadClass) Alloc() Pad {
	rv := objc.Send[Pad](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(coder:device:)
func NewPadWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) Pad {
	instance := getPadClass().Alloc()
	rv := objc.Send[Pad](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeAfter
func (p_ Pad) PaddingSizeAfter() ImageCoordinate /* not a class type */ {
	rv := objc.Send[ImageCoordinate](p_.ID, objc.Sel("paddingSizeAfter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeAfter
func (p_ Pad) SetPaddingSizeAfter(value ImageCoordinate /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaddingSizeAfter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeBefore
func (p_ Pad) PaddingSizeBefore() ImageCoordinate /* not a class type */ {
	rv := objc.Send[ImageCoordinate](p_.ID, objc.Sel("paddingSizeBefore"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeBefore
func (p_ Pad) SetPaddingSizeBefore(value ImageCoordinate /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaddingSizeBefore:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/fillvalue
func (p_ Pad) FillValue() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("fillValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/fillvalue
func (p_ Pad) SetFillValue(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFillValue:"), value)
}


