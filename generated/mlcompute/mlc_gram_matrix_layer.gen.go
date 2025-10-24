// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCGramMatrixLayer */


/* debug [class_header]: Header for MLCGramMatrixLayer */
// The class instance for the [CGramMatrixLayer] class.
var (
	CGramMatrixLayerClass     _CGramMatrixLayerClass
	CGramMatrixLayerClassOnce sync.Once
)

func getCGramMatrixLayerClass() _CGramMatrixLayerClass {
	CGramMatrixLayerClassOnce.Do(func() {
		CGramMatrixLayerClass = _CGramMatrixLayerClass{objc.GetClass("MLCGramMatrixLayer")}
	})
	return CGramMatrixLayerClass
}

type _CGramMatrixLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CGramMatrixLayer */
// An interface definition for the [CGramMatrixLayer] class.
type ICGramMatrixLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CGramMatrixLayer */
	// properties:
	Scale() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CGramMatrixLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CGramMatrixLayer */
// Alloc allocates a new instance without initialization.
func (cc _CGramMatrixLayerClass) Alloc() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CGramMatrixLayerClass) New() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGramMatrixLayer) Init() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGramMatrixLayer) Autorelease() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGramMatrixLayer creates a new CGramMatrixLayer instance.
func NewCGramMatrixLayer() CGramMatrixLayer {
	return getCGramMatrixLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CGramMatrixLayer */
// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
//
// For example, if the input tensor batch function is: The computation performed by this layer is: Interpret this operation as computing all combinations of fully connected layers between the different spatial planes of the input tensor. The layer performs this operation independently for each tensor in a batch. Then the layer stores these results in the feature channel and x-coordinate indices of the output batch. Legend:


// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGramMatrixLayer
type CGramMatrixLayer struct {
	CLayer
}

// CGramMatrixLayerFrom constructs a [CGramMatrixLayer] from an unsafe.Pointer.
//
// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
func CGramMatrixLayerFrom(ptr unsafe.Pointer) CGramMatrixLayer {
	return CGramMatrixLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CGramMatrixLayer */

// Creates a gram matrix layer with the scaling factor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGramMatrixLayer/init(scale:)
func NewCGramMatrixLayerWithScale(scale float32) CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](objc.ID(getCGramMatrixLayerClass().class), objc.Sel("layerWithScale:"), scale)
	return rv
}/* debug [class_init_methods/constructor]: NewCGramMatrixLayerWithScale */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CGramMatrixLayer */

// Creates a gram matrix layer with the scaling factor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGramMatrixLayer/init(scale:)
func (cc _CGramMatrixLayerClass) LayerWithScale(scale float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithScale:"), scale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithScale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CGramMatrixLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CGramMatrixLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CGramMatrixLayer */

// The scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGramMatrixLayer/scale
func (c_ CGramMatrixLayer) Scale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scale"))
	return rv
}/* debug [instance_properties/getter]: scale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCGramMatrixLayer */


