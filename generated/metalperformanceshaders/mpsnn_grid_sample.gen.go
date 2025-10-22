// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GridSample] class.
var (
	GridSampleClass     _GridSampleClass
	GridSampleClassOnce sync.Once
)

func getGridSampleClass() _GridSampleClass {
	GridSampleClassOnce.Do(func() {
		GridSampleClass = _GridSampleClass{objc.GetClass("MPSNNGridSample")}
	})
	return GridSampleClass
}

type _GridSampleClass struct {
	class objc.Class
}

// An interface definition for the [GridSample] class.
type IGridSample interface {
	ICNNBinaryKernel
	UseGridValueAsInputCoordinate() bool
	SetUseGridValueAsInputCoordinate(value bool)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample
type GridSample struct {
	CNNBinaryKernel
}

// GridSampleFrom constructs a [GridSample] from an unsafe.Pointer.
func GridSampleFrom(ptr unsafe.Pointer) GridSample {
	return GridSample{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GridSampleClass) Alloc() GridSample {
	rv := objc.Send[GridSample](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GridSampleClass) New() GridSample {
	rv := objc.Send[GridSample](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridSample) Init() GridSample {
	rv := objc.Send[GridSample](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridSample) Autorelease() GridSample {
	rv := objc.Send[GridSample](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridSample creates a new GridSample instance.
func NewGridSample() GridSample {
	return getGridSampleClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample/init(device:)
func NewGridSampleWithDevice(device objectivec.IObject) GridSample {
	instance := getGridSampleClass().Alloc()
	rv := objc.Send[GridSample](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/usegridvalueasinputcoordinate
func (g_ GridSample) UseGridValueAsInputCoordinate() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("useGridValueAsInputCoordinate"))
	return rv
}


// SetUseGridValueAsInputCoordinate sets the value of the useGridValueAsInputCoordinate property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/usegridvalueasinputcoordinate
func (g_ GridSample) SetUseGridValueAsInputCoordinate(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUseGridValueAsInputCoordinate:"), value)
}


