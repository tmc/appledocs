// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCArithmeticLayer */


/* debug [class_header]: Header for MLCArithmeticLayer */
// The class instance for the [CArithmeticLayer] class.
var (
	CArithmeticLayerClass     _CArithmeticLayerClass
	CArithmeticLayerClassOnce sync.Once
)

func getCArithmeticLayerClass() _CArithmeticLayerClass {
	CArithmeticLayerClassOnce.Do(func() {
		CArithmeticLayerClass = _CArithmeticLayerClass{objc.GetClass("MLCArithmeticLayer")}
	})
	return CArithmeticLayerClass
}

type _CArithmeticLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CArithmeticLayer */
// An interface definition for the [CArithmeticLayer] class.
type ICArithmeticLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CArithmeticLayer */
	// properties:
	Operation() CArithmeticOperation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CArithmeticLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CArithmeticLayer */
// Alloc allocates a new instance without initialization.
func (cc _CArithmeticLayerClass) Alloc() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CArithmeticLayerClass) New() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CArithmeticLayer) Init() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CArithmeticLayer) Autorelease() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCArithmeticLayer creates a new CArithmeticLayer instance.
func NewCArithmeticLayer() CArithmeticLayer {
	return getCArithmeticLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CArithmeticLayer */
// A layer that performs an arithmetic operation.


// A layer that performs an arithmetic operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticLayer
type CArithmeticLayer struct {
	CLayer
}

// CArithmeticLayerFrom constructs a [CArithmeticLayer] from an unsafe.Pointer.
//
// A layer that performs an arithmetic operation.
func CArithmeticLayerFrom(ptr unsafe.Pointer) CArithmeticLayer {
	return CArithmeticLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CArithmeticLayer */

// Creates an arithmetic layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticLayer/init(operation:)
func NewCArithmeticLayerWithOperation(operation CArithmeticOperation) CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](objc.ID(getCArithmeticLayerClass().class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_init_methods/constructor]: NewCArithmeticLayerWithOperation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CArithmeticLayer */

// Creates an arithmetic layer with the operation you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticLayer/init(operation:)
func (cc _CArithmeticLayerClass) LayerWithOperation(operation CArithmeticOperation) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithOperation:"), operation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithOperation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CArithmeticLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CArithmeticLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CArithmeticLayer */

// The arithmetic layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticLayer/operation
func (c_ CArithmeticLayer) Operation() CArithmeticOperation {
	rv := objc.Send[CArithmeticOperation](c_.ID, objc.Sel("operation"))
	return rv
}/* debug [instance_properties/getter]: operation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCArithmeticLayer */


