// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCUpsampleLayer */


/* debug [class_header]: Header for MLCUpsampleLayer */
// The class instance for the [CUpsampleLayer] class.
var (
	CUpsampleLayerClass     _CUpsampleLayerClass
	CUpsampleLayerClassOnce sync.Once
)

func getCUpsampleLayerClass() _CUpsampleLayerClass {
	CUpsampleLayerClassOnce.Do(func() {
		CUpsampleLayerClass = _CUpsampleLayerClass{objc.GetClass("MLCUpsampleLayer")}
	})
	return CUpsampleLayerClass
}

type _CUpsampleLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CUpsampleLayer */
// An interface definition for the [CUpsampleLayer] class.
type ICUpsampleLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CUpsampleLayer */
	// properties:
	AlignsCorners() bool
	SampleMode() CSampleMode
	Shape() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CUpsampleLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CUpsampleLayer */
// Alloc allocates a new instance without initialization.
func (cc _CUpsampleLayerClass) Alloc() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CUpsampleLayerClass) New() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CUpsampleLayer) Init() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CUpsampleLayer) Autorelease() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCUpsampleLayer creates a new CUpsampleLayer instance.
func NewCUpsampleLayer() CUpsampleLayer {
	return getCUpsampleLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CUpsampleLayer */
// A layer that applies upsampling with the shape you specify.


// A layer that applies upsampling with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer
type CUpsampleLayer struct {
	CLayer
}

// CUpsampleLayerFrom constructs a [CUpsampleLayer] from an unsafe.Pointer.
//
// A layer that applies upsampling with the shape you specify.
func CUpsampleLayerFrom(ptr unsafe.Pointer) CUpsampleLayer {
	return CUpsampleLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CUpsampleLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CUpsampleLayer */

// Creates an upsample layer with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer/layerWithShape:
func (cc _CUpsampleLayerClass) LayerWithShape(shape []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithShape:"), shape)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithShape) */


// Creates an upsample layer with the shape, upsampling algorithm, and corner alignement option you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer/layerWithShape:sampleMode:alignsCorners:
func (cc _CUpsampleLayerClass) LayerWithShapeSampleModeAlignsCorners(shape []foundation.Number, sampleMode CSampleMode, alignsCorners bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithShape:sampleMode:alignsCorners:"), shape, sampleMode, alignsCorners)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithShapeSampleModeAlignsCorners) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CUpsampleLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CUpsampleLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CUpsampleLayer */

// A Boolean that indicates whether the layer aligns the corner pixels of the input and output tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer/alignsCorners
func (c_ CUpsampleLayer) AlignsCorners() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alignsCorners"))
	return rv
}/* debug [instance_properties/getter]: alignsCorners */


// The upsampling algorithm type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer/sampleMode
func (c_ CUpsampleLayer) SampleMode() CSampleMode {
	rv := objc.Send[CSampleMode](c_.ID, objc.Sel("sampleMode"))
	return rv
}/* debug [instance_properties/getter]: sampleMode */


// An array that contains the dimensions of the result tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer/shape-7j6sf
func (c_ CUpsampleLayer) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCUpsampleLayer */



