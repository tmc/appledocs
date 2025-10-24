// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCComparisonLayer */


/* debug [class_header]: Header for MLCComparisonLayer */
// The class instance for the [CComparisonLayer] class.
var (
	CComparisonLayerClass     _CComparisonLayerClass
	CComparisonLayerClassOnce sync.Once
)

func getCComparisonLayerClass() _CComparisonLayerClass {
	CComparisonLayerClassOnce.Do(func() {
		CComparisonLayerClass = _CComparisonLayerClass{objc.GetClass("MLCComparisonLayer")}
	})
	return CComparisonLayerClass
}

type _CComparisonLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CComparisonLayer */
// An interface definition for the [CComparisonLayer] class.
type ICComparisonLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CComparisonLayer */
	// properties:
	Operation() CComparisonOperation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CComparisonLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CComparisonLayer */
// Alloc allocates a new instance without initialization.
func (cc _CComparisonLayerClass) Alloc() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CComparisonLayerClass) New() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CComparisonLayer) Init() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CComparisonLayer) Autorelease() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCComparisonLayer creates a new CComparisonLayer instance.
func NewCComparisonLayer() CComparisonLayer {
	return getCComparisonLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CComparisonLayer */
// A layer that performs elementwise comparison of two tensors.
//
// The layer returns a tensor with the shape equal to the largest shape of operations. It fills with the Boolean value , where corresponds to the you specify.


// A layer that performs elementwise comparison of two tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonLayer
type CComparisonLayer struct {
	CLayer
}

// CComparisonLayerFrom constructs a [CComparisonLayer] from an unsafe.Pointer.
//
// A layer that performs elementwise comparison of two tensors.
func CComparisonLayerFrom(ptr unsafe.Pointer) CComparisonLayer {
	return CComparisonLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CComparisonLayer */

// Creates a comparison layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonLayer/init(operation:)
func NewCComparisonLayerWithOperation(operation CComparisonOperation) CComparisonLayer {
	rv := objc.Send[CComparisonLayer](objc.ID(getCComparisonLayerClass().class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_init_methods/constructor]: NewCComparisonLayerWithOperation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CComparisonLayer */

// Creates a comparison layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonLayer/init(operation:)
func (cc _CComparisonLayerClass) LayerWithOperation(operation CComparisonOperation) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithOperation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CComparisonLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CComparisonLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CComparisonLayer */

// The comparison layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonLayer/operation
func (c_ CComparisonLayer) Operation() CComparisonOperation {
	rv := objc.Send[CComparisonOperation](c_.ID, objc.Sel("operation"))
	return rv
}/* debug [instance_properties/getter]: operation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCComparisonLayer */


