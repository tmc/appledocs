// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCReshapeLayer */


/* debug [class_header]: Header for MLCReshapeLayer */
// The class instance for the [CReshapeLayer] class.
var (
	CReshapeLayerClass     _CReshapeLayerClass
	CReshapeLayerClassOnce sync.Once
)

func getCReshapeLayerClass() _CReshapeLayerClass {
	CReshapeLayerClassOnce.Do(func() {
		CReshapeLayerClass = _CReshapeLayerClass{objc.GetClass("MLCReshapeLayer")}
	})
	return CReshapeLayerClass
}

type _CReshapeLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CReshapeLayer */
// An interface definition for the [CReshapeLayer] class.
type ICReshapeLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CReshapeLayer */
	// properties:
	Shape() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CReshapeLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CReshapeLayer */
// Alloc allocates a new instance without initialization.
func (cc _CReshapeLayerClass) Alloc() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CReshapeLayerClass) New() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CReshapeLayer) Init() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CReshapeLayer) Autorelease() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCReshapeLayer creates a new CReshapeLayer instance.
func NewCReshapeLayer() CReshapeLayer {
	return getCReshapeLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CReshapeLayer */
// A layer that reshapes a tensor with the shape you specify.


// A layer that reshapes a tensor with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReshapeLayer
type CReshapeLayer struct {
	CLayer
}

// CReshapeLayerFrom constructs a [CReshapeLayer] from an unsafe.Pointer.
//
// A layer that reshapes a tensor with the shape you specify.
func CReshapeLayerFrom(ptr unsafe.Pointer) CReshapeLayer {
	return CReshapeLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CReshapeLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CReshapeLayer */

// Creates a reshape layer with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReshapeLayer/layerWithShape:
func (cc _CReshapeLayerClass) LayerWithShape(shape []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithShape:"), shape)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithShape) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CReshapeLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CReshapeLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CReshapeLayer */

// An array that contains the size of each dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReshapeLayer/shape-840ax
func (c_ CReshapeLayer) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCReshapeLayer */



