// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCConcatenationLayer */


/* debug [class_header]: Header for MLCConcatenationLayer */
// The class instance for the [CConcatenationLayer] class.
var (
	CConcatenationLayerClass     _CConcatenationLayerClass
	CConcatenationLayerClassOnce sync.Once
)

func getCConcatenationLayerClass() _CConcatenationLayerClass {
	CConcatenationLayerClassOnce.Do(func() {
		CConcatenationLayerClass = _CConcatenationLayerClass{objc.GetClass("MLCConcatenationLayer")}
	})
	return CConcatenationLayerClass
}

type _CConcatenationLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CConcatenationLayer */
// An interface definition for the [CConcatenationLayer] class.
type ICConcatenationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CConcatenationLayer */
	// properties:
	Dimension() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CConcatenationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CConcatenationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CConcatenationLayerClass) Alloc() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CConcatenationLayerClass) New() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConcatenationLayer) Init() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConcatenationLayer) Autorelease() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConcatenationLayer creates a new CConcatenationLayer instance.
func NewCConcatenationLayer() CConcatenationLayer {
	return getCConcatenationLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CConcatenationLayer */
// A layer that combines tensors into a single tensor.


// A layer that combines tensors into a single tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer
type CConcatenationLayer struct {
	CLayer
}

// CConcatenationLayerFrom constructs a [CConcatenationLayer] from an unsafe.Pointer.
//
// A layer that combines tensors into a single tensor.
func CConcatenationLayerFrom(ptr unsafe.Pointer) CConcatenationLayer {
	return CConcatenationLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CConcatenationLayer */

// Creates a concatenation layer with a dimension value of 1, which typically represents feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer/init()
func NewCConcatenationLayer() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(getCConcatenationLayerClass().class), objc.Sel("layer"))
	return rv
}/* debug [class_init_methods/constructor]: NewCConcatenationLayer */


// Creates a concatenation layer with the dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer/init(dimension:)
func NewCConcatenationLayerWithDimension(dimension uint) CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(getCConcatenationLayerClass().class), objc.Sel("layerWithDimension:"), dimension)
	return rv
}/* debug [class_init_methods/constructor]: NewCConcatenationLayerWithDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CConcatenationLayer */

// Creates a concatenation layer with a dimension value of 1, which typically represents feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer/init()
func (cc _CConcatenationLayerClass) Layer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Layer) */


// Creates a concatenation layer with the dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer/init(dimension:)
func (cc _CConcatenationLayerClass) LayerWithDimension(dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDimension:"), dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDimension) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CConcatenationLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CConcatenationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CConcatenationLayer */

// The dimension, or axis, along which you concatenate tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer/dimension
func (c_ CConcatenationLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCConcatenationLayer */


