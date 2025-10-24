// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCSliceLayer */


/* debug [class_header]: Header for MLCSliceLayer */
// The class instance for the [CSliceLayer] class.
var (
	CSliceLayerClass     _CSliceLayerClass
	CSliceLayerClassOnce sync.Once
)

func getCSliceLayerClass() _CSliceLayerClass {
	CSliceLayerClassOnce.Do(func() {
		CSliceLayerClass = _CSliceLayerClass{objc.GetClass("MLCSliceLayer")}
	})
	return CSliceLayerClass
}

type _CSliceLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSliceLayer */
// An interface definition for the [CSliceLayer] class.
type ICSliceLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CSliceLayer */
	// properties:
	End() []foundation.Number
	Start() []foundation.Number
	Stride() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSliceLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSliceLayer */
// Alloc allocates a new instance without initialization.
func (cc _CSliceLayerClass) Alloc() CSliceLayer {
	rv := objc.Send[CSliceLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSliceLayerClass) New() CSliceLayer {
	rv := objc.Send[CSliceLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSliceLayer) Init() CSliceLayer {
	rv := objc.Send[CSliceLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSliceLayer) Autorelease() CSliceLayer {
	rv := objc.Send[CSliceLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSliceLayer creates a new CSliceLayer instance.
func NewCSliceLayer() CSliceLayer {
	return getCSliceLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSliceLayer */
// A layer that extracts a slice from a tensor.
//
// The framework supports positive stride. Use a slice layer to slice a given source. Slicing won’t decrease the tensor dimension. The start, end, and stride vectors must be of the same size, equal to the source tensor dimension.


// A layer that extracts a slice from a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer
type CSliceLayer struct {
	CLayer
}

// CSliceLayerFrom constructs a [CSliceLayer] from an unsafe.Pointer.
//
// A layer that extracts a slice from a tensor.
func CSliceLayerFrom(ptr unsafe.Pointer) CSliceLayer {
	return CSliceLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSliceLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSliceLayer */

// Creates a slice layer with the specified start, end, and stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer/sliceLayerWithStart:end:stride:
func (cc _CSliceLayerClass) SliceLayerWithStartEndStride(start []foundation.Number, end []foundation.Number, stride []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sliceLayerWithStart:end:stride:"), start, end, stride)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SliceLayerWithStartEndStride) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSliceLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSliceLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSliceLayer */

// The end vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer/end-8z0wi
func (c_ CSliceLayer) End() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("end"))
	return rv
}/* debug [instance_properties/getter]: end */


// The start vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer/start-6dnjn
func (c_ CSliceLayer) Start() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_properties/getter]: start */


// The stride vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer/stride-8dnpu
func (c_ CSliceLayer) Stride() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("stride"))
	return rv
}/* debug [instance_properties/getter]: stride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCSliceLayer */



