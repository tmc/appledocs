// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCScatterLayer */


/* debug [class_header]: Header for MLCScatterLayer */
// The class instance for the [CScatterLayer] class.
var (
	CScatterLayerClass     _CScatterLayerClass
	CScatterLayerClassOnce sync.Once
)

func getCScatterLayerClass() _CScatterLayerClass {
	CScatterLayerClassOnce.Do(func() {
		CScatterLayerClass = _CScatterLayerClass{objc.GetClass("MLCScatterLayer")}
	})
	return CScatterLayerClass
}

type _CScatterLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CScatterLayer */
// An interface definition for the [CScatterLayer] class.
type ICScatterLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CScatterLayer */
	// properties:
	Dimension() uint
	ReductionType() CReductionType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CScatterLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CScatterLayer */
// Alloc allocates a new instance without initialization.
func (cc _CScatterLayerClass) Alloc() CScatterLayer {
	rv := objc.Send[CScatterLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CScatterLayerClass) New() CScatterLayer {
	rv := objc.Send[CScatterLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CScatterLayer) Init() CScatterLayer {
	rv := objc.Send[CScatterLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CScatterLayer) Autorelease() CScatterLayer {
	rv := objc.Send[CScatterLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCScatterLayer creates a new CScatterLayer instance.
func NewCScatterLayer() CScatterLayer {
	return getCScatterLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CScatterLayer */
// A layer that updates the output at an index you specify.


// A layer that updates the output at an index you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer
type CScatterLayer struct {
	CLayer
}

// CScatterLayerFrom constructs a [CScatterLayer] from an unsafe.Pointer.
//
// A layer that updates the output at an index you specify.
func CScatterLayerFrom(ptr unsafe.Pointer) CScatterLayer {
	return CScatterLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CScatterLayer */

// Creates a scatter layer with the dimension and reduction type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer/init(dimension:reductionType:)
func NewCScatterLayerWithDimensionReductionType(dimension uint, reductionType CReductionType) CScatterLayer {
	rv := objc.Send[CScatterLayer](objc.ID(getCScatterLayerClass().class), objc.Sel("layerWithDimension:reductionType:"), dimension, reductionType)
	return rv
}/* debug [class_init_methods/constructor]: NewCScatterLayerWithDimensionReductionType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CScatterLayer */

// Creates a scatter layer with the dimension and reduction type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer/init(dimension:reductionType:)
func (cc _CScatterLayerClass) LayerWithDimensionReductionType(dimension uint, reductionType CReductionType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDimension:reductionType:"), dimension, reductionType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDimensionReductionType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CScatterLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CScatterLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CScatterLayer */

// The dimension to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer/dimension
func (c_ CScatterLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// The reduction type that applies to all values in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer/reductionType
func (c_ CScatterLayer) ReductionType() CReductionType {
	rv := objc.Send[CReductionType](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCScatterLayer */


