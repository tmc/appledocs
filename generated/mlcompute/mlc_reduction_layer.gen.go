// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCReductionLayer */


/* debug [class_header]: Header for MLCReductionLayer */
// The class instance for the [CReductionLayer] class.
var (
	CReductionLayerClass     _CReductionLayerClass
	CReductionLayerClassOnce sync.Once
)

func getCReductionLayerClass() _CReductionLayerClass {
	CReductionLayerClassOnce.Do(func() {
		CReductionLayerClass = _CReductionLayerClass{objc.GetClass("MLCReductionLayer")}
	})
	return CReductionLayerClass
}

type _CReductionLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CReductionLayer */
// An interface definition for the [CReductionLayer] class.
type ICReductionLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CReductionLayer */
	// properties:
	Dimension() uint
	Dimensions() []foundation.Number
	ReductionType() CReductionType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CReductionLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CReductionLayer */
// Alloc allocates a new instance without initialization.
func (cc _CReductionLayerClass) Alloc() CReductionLayer {
	rv := objc.Send[CReductionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CReductionLayerClass) New() CReductionLayer {
	rv := objc.Send[CReductionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CReductionLayer) Init() CReductionLayer {
	rv := objc.Send[CReductionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CReductionLayer) Autorelease() CReductionLayer {
	rv := objc.Send[CReductionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCReductionLayer creates a new CReductionLayer instance.
func NewCReductionLayer() CReductionLayer {
	return getCReductionLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CReductionLayer */
// A layer that reduces tensor values across a specific dimension to a scalar value.
//
// Use this layer to perform reduction operations on a given dimension. The output of this layer is a tensor of the same shape as the source tensor, except the layer sets the dimension to .


// A layer that reduces tensor values across a specific dimension to a scalar value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer
type CReductionLayer struct {
	CLayer
}

// CReductionLayerFrom constructs a [CReductionLayer] from an unsafe.Pointer.
//
// A layer that reduces tensor values across a specific dimension to a scalar value.
func CReductionLayerFrom(ptr unsafe.Pointer) CReductionLayer {
	return CReductionLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CReductionLayer */

// Creates a reduction layer using the reduction type and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/init(reductionType:dimension:)
func NewCReductionLayerWithReductionTypeDimension(reductionType CReductionType, dimension uint) CReductionLayer {
	rv := objc.Send[CReductionLayer](objc.ID(getCReductionLayerClass().class), objc.Sel("layerWithReductionType:dimension:"), reductionType, dimension)
	return rv
}/* debug [class_init_methods/constructor]: NewCReductionLayerWithReductionTypeDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CReductionLayer */

// Creates a reduction layer using the reduction type and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/init(reductionType:dimension:)
func (cc _CReductionLayerClass) LayerWithReductionTypeDimension(reductionType CReductionType, dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithReductionType:dimension:"), reductionType, dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithReductionTypeDimension) */


// Creates a reduction layer using the reduction type and dimensions you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/layerWithReductionType:dimensions:
func (cc _CReductionLayerClass) LayerWithReductionTypeDimensions(reductionType CReductionType, dimensions []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithReductionType:dimensions:"), reductionType, dimensions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithReductionTypeDimensions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CReductionLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CReductionLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CReductionLayer */

// The dimension to perform the reduction operation on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/dimension
func (c_ CReductionLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// The dimensions to perform the reduction operation on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/dimensions-4359b
func (c_ CReductionLayer) Dimensions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("dimensions"))
	return rv
}/* debug [instance_properties/getter]: dimensions */


// The function reduction type the system uses for reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer/reductionType
func (c_ CReductionLayer) ReductionType() CReductionType {
	rv := objc.Send[CReductionType](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCReductionLayer */


