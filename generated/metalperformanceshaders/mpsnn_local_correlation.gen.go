// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocalCorrelation] class.
var (
	LocalCorrelationClass     _LocalCorrelationClass
	LocalCorrelationClassOnce sync.Once
)

func getLocalCorrelationClass() _LocalCorrelationClass {
	LocalCorrelationClassOnce.Do(func() {
		LocalCorrelationClass = _LocalCorrelationClass{objc.GetClass("MPSNNLocalCorrelation")}
	})
	return LocalCorrelationClass
}

type _LocalCorrelationClass struct {
	class objc.Class
}

// An interface definition for the [LocalCorrelation] class.
type ILocalCorrelation interface {
	IReduceBinary
	// properties:
	StrideInY() uint
	SetStrideInY(value uint)
	StrideInX() int
	SetStrideInX(value int)
	WindowInX() int
	SetWindowInX(value int)
	WindowInY() int
	SetWindowInY(value int)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation
type LocalCorrelation struct {
	ReduceBinary
}

// LocalCorrelationFrom constructs a [LocalCorrelation] from an unsafe.Pointer.
func LocalCorrelationFrom(ptr unsafe.Pointer) LocalCorrelation {
	return LocalCorrelation{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LocalCorrelationClass) Alloc() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocalCorrelationClass) New() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalCorrelation) Init() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalCorrelation) Autorelease() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalCorrelation creates a new LocalCorrelation instance.
func NewLocalCorrelation() LocalCorrelation {
	return getLocalCorrelationClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/init(coder:device:)
func NewLocalCorrelationWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/strideInY
func (l_ LocalCorrelation) StrideInY() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("strideInY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/strideInY
func (l_ LocalCorrelation) SetStrideInY(value uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/strideinx
func (l_ LocalCorrelation) StrideInX() int {
	rv := objc.Send[int](l_.ID, objc.Sel("strideInX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/strideinx
func (l_ LocalCorrelation) SetStrideInX(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/windowinx
func (l_ LocalCorrelation) WindowInX() int {
	rv := objc.Send[int](l_.ID, objc.Sel("windowInX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/windowinx
func (l_ LocalCorrelation) SetWindowInX(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/windowiny
func (l_ LocalCorrelation) WindowInY() int {
	rv := objc.Send[int](l_.ID, objc.Sel("windowInY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/windowiny
func (l_ LocalCorrelation) SetWindowInY(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInY:"), value)
}


