// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCSoftmaxLayer */


/* debug [class_header]: Header for MLCSoftmaxLayer */
// The class instance for the [CSoftmaxLayer] class.
var (
	CSoftmaxLayerClass     _CSoftmaxLayerClass
	CSoftmaxLayerClassOnce sync.Once
)

func getCSoftmaxLayerClass() _CSoftmaxLayerClass {
	CSoftmaxLayerClassOnce.Do(func() {
		CSoftmaxLayerClass = _CSoftmaxLayerClass{objc.GetClass("MLCSoftmaxLayer")}
	})
	return CSoftmaxLayerClass
}

type _CSoftmaxLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSoftmaxLayer */
// An interface definition for the [CSoftmaxLayer] class.
type ICSoftmaxLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CSoftmaxLayer */
	// properties:
	Dimension() uint
	Operation() CSoftmaxOperation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSoftmaxLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSoftmaxLayer */
// Alloc allocates a new instance without initialization.
func (cc _CSoftmaxLayerClass) Alloc() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSoftmaxLayerClass) New() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSoftmaxLayer) Init() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSoftmaxLayer) Autorelease() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSoftmaxLayer creates a new CSoftmaxLayer instance.
func NewCSoftmaxLayer() CSoftmaxLayer {
	return getCSoftmaxLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSoftmaxLayer */
// A layer that outputs a probability distribution as attention weights.


// A layer that outputs a probability distribution as attention weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer
type CSoftmaxLayer struct {
	CLayer
}

// CSoftmaxLayerFrom constructs a [CSoftmaxLayer] from an unsafe.Pointer.
//
// A layer that outputs a probability distribution as attention weights.
func CSoftmaxLayerFrom(ptr unsafe.Pointer) CSoftmaxLayer {
	return CSoftmaxLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSoftmaxLayer */

// Creates a softmax layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/init(operation:)
func NewCSoftmaxLayerWithOperation(operation CSoftmaxOperation) CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(getCSoftmaxLayerClass().class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_init_methods/constructor]: NewCSoftmaxLayerWithOperation */


// Creates a softmax layer with the operation and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/init(operation:dimension:)
func NewCSoftmaxLayerWithOperationDimension(operation CSoftmaxOperation, dimension uint) CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(getCSoftmaxLayerClass().class), objc.Sel("layerWithOperation:dimension:"), operation, dimension)
	return rv
}/* debug [class_init_methods/constructor]: NewCSoftmaxLayerWithOperationDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSoftmaxLayer */

// Creates a softmax layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/init(operation:)
func (cc _CSoftmaxLayerClass) LayerWithOperation(operation CSoftmaxOperation) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithOperation) */


// Creates a softmax layer with the operation and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/init(operation:dimension:)
func (cc _CSoftmaxLayerClass) LayerWithOperationDimension(operation CSoftmaxOperation, dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithOperation:dimension:"), operation, dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithOperationDimension) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSoftmaxLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSoftmaxLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSoftmaxLayer */

// The dimension over which you want to perform the softmax operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/dimension
func (c_ CSoftmaxLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// The softmax operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer/operation
func (c_ CSoftmaxLayer) Operation() CSoftmaxOperation {
	rv := objc.Send[CSoftmaxOperation](c_.ID, objc.Sel("operation"))
	return rv
}/* debug [instance_properties/getter]: operation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCSoftmaxLayer */


