// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCSelectionLayer */


/* debug [class_header]: Header for MLCSelectionLayer */
// The class instance for the [CSelectionLayer] class.
var (
	CSelectionLayerClass     _CSelectionLayerClass
	CSelectionLayerClassOnce sync.Once
)

func getCSelectionLayerClass() _CSelectionLayerClass {
	CSelectionLayerClassOnce.Do(func() {
		CSelectionLayerClass = _CSelectionLayerClass{objc.GetClass("MLCSelectionLayer")}
	})
	return CSelectionLayerClass
}

type _CSelectionLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSelectionLayer */
// An interface definition for the [CSelectionLayer] class.
type ICSelectionLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CSelectionLayer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSelectionLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSelectionLayer */
// Alloc allocates a new instance without initialization.
func (cc _CSelectionLayerClass) Alloc() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSelectionLayerClass) New() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSelectionLayer) Init() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSelectionLayer) Autorelease() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSelectionLayer creates a new CSelectionLayer instance.
func NewCSelectionLayer() CSelectionLayer {
	return getCSelectionLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSelectionLayer */
// A layer for selecting elements from two tensors.
//
// A selection layer takes a condition tensor that acts as a mask. It determines whether the corresponding element or row in the output comes from tensor (if the element in the condition is ) or tensor (if ).


// A layer for selecting elements from two tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSelectionLayer
type CSelectionLayer struct {
	CLayer
}

// CSelectionLayerFrom constructs a [CSelectionLayer] from an unsafe.Pointer.
//
// A layer for selecting elements from two tensors.
func CSelectionLayerFrom(ptr unsafe.Pointer) CSelectionLayer {
	return CSelectionLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSelectionLayer */

// Creates a selection layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSelectionLayer/init()
func NewCSelectionLayer() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](objc.ID(getCSelectionLayerClass().class), objc.Sel("layer"))
	return rv
}/* debug [class_init_methods/constructor]: NewCSelectionLayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSelectionLayer */

// Creates a selection layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSelectionLayer/init()
func (cc _CSelectionLayerClass) Layer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Layer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSelectionLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSelectionLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSelectionLayer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCSelectionLayer */


