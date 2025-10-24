// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNGridSample */


/* debug [class_header]: Header for MPSNNGridSample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridSample */
// An interface definition for the [GridSample] class.
type IGridSample interface {
	ICNNBinaryKernel
	
/* debug [class_interface_properties]: Properties for GridSample */
	// properties:
	UseGridValueAsInputCoordinate() objectivec.IObject
	SetUseGridValueAsInputCoordinate(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridSample */
// Alloc allocates a new instance without initialization.
func (gc _GridSampleClass) Alloc() GridSample {
	rv := objc.Send[GridSample](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridSample */


// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridSample */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131870-initwithcoder
func NewGridSampleWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) GridSample {
	instance := getGridSampleClass().Alloc()
	rv := objc.Send[GridSample](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridSampleWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131871-initwithdevice
func NewGridSampleWithDevice(device unsafe.Pointer) GridSample {
	instance := getGridSampleClass().Alloc()
	rv := objc.Send[GridSample](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridSampleWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridSample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridSample */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131872-usegridvalueasinputcoordinate
func (g_ GridSample) UseGridValueAsInputCoordinate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("useGridValueAsInputCoordinate"))
	return rv
}/* debug [instance_properties/getter]: useGridValueAsInputCoordinate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131872-usegridvalueasinputcoordinate
func (g_ GridSample) SetUseGridValueAsInputCoordinate(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUseGridValueAsInputCoordinate:"), value)
}/* debug [instance_properties/setter]: useGridValueAsInputCoordinate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNGridSample */


