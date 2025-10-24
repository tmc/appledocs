// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCTransposeLayer */


/* debug [class_header]: Header for MLCTransposeLayer */
// The class instance for the [CTransposeLayer] class.
var (
	CTransposeLayerClass     _CTransposeLayerClass
	CTransposeLayerClassOnce sync.Once
)

func getCTransposeLayerClass() _CTransposeLayerClass {
	CTransposeLayerClassOnce.Do(func() {
		CTransposeLayerClass = _CTransposeLayerClass{objc.GetClass("MLCTransposeLayer")}
	})
	return CTransposeLayerClass
}

type _CTransposeLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTransposeLayer */
// An interface definition for the [CTransposeLayer] class.
type ICTransposeLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CTransposeLayer */
	// properties:
	Dimensions() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTransposeLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTransposeLayer */
// Alloc allocates a new instance without initialization.
func (cc _CTransposeLayerClass) Alloc() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTransposeLayerClass) New() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTransposeLayer) Init() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTransposeLayer) Autorelease() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTransposeLayer creates a new CTransposeLayer instance.
func NewCTransposeLayer() CTransposeLayer {
	return getCTransposeLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTransposeLayer */
// A layer that permutes the dimensions you specify.


// A layer that permutes the dimensions you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTransposeLayer
type CTransposeLayer struct {
	CLayer
}

// CTransposeLayerFrom constructs a [CTransposeLayer] from an unsafe.Pointer.
//
// A layer that permutes the dimensions you specify.
func CTransposeLayerFrom(ptr unsafe.Pointer) CTransposeLayer {
	return CTransposeLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTransposeLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTransposeLayer */

// Creates a transpose layer with the dimensions you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTransposeLayer/layerWithDimensions:
func (cc _CTransposeLayerClass) LayerWithDimensions(dimensions []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDimensions:"), dimensions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDimensions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTransposeLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTransposeLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTransposeLayer */

// An array that contains an input axis source for each output axis, which represents the ordering of dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTransposeLayer/dimensions-1d5nw
func (c_ CTransposeLayer) Dimensions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("dimensions"))
	return rv
}/* debug [instance_properties/getter]: dimensions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTransposeLayer */



