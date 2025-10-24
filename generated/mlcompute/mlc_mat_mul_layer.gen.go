// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCMatMulLayer */


/* debug [class_header]: Header for MLCMatMulLayer */
// The class instance for the [CMatMulLayer] class.
var (
	CMatMulLayerClass     _CMatMulLayerClass
	CMatMulLayerClassOnce sync.Once
)

func getCMatMulLayerClass() _CMatMulLayerClass {
	CMatMulLayerClassOnce.Do(func() {
		CMatMulLayerClass = _CMatMulLayerClass{objc.GetClass("MLCMatMulLayer")}
	})
	return CMatMulLayerClass
}

type _CMatMulLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CMatMulLayer */
// An interface definition for the [CMatMulLayer] class.
type ICMatMulLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CMatMulLayer */
	// properties:
	Descriptor() IMLCMatMulDescriptor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CMatMulLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CMatMulLayer */
// Alloc allocates a new instance without initialization.
func (cc _CMatMulLayerClass) Alloc() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CMatMulLayerClass) New() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMatMulLayer) Init() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMatMulLayer) Autorelease() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMatMulLayer creates a new CMatMulLayer instance.
func NewCMatMulLayer() CMatMulLayer {
	return getCMatMulLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CMatMulLayer */
// A layer that multiplies matrices.


// A layer that multiplies matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulLayer
type CMatMulLayer struct {
	CLayer
}

// CMatMulLayerFrom constructs a [CMatMulLayer] from an unsafe.Pointer.
//
// A layer that multiplies matrices.
func CMatMulLayerFrom(ptr unsafe.Pointer) CMatMulLayer {
	return CMatMulLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CMatMulLayer */

// Creates a matrix multiplication layer with the specified descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulLayer/init(descriptor:)
func NewCMatMulLayerWithDescriptor(descriptor IMLCMatMulDescriptor) CMatMulLayer {
	rv := objc.Send[CMatMulLayer](objc.ID(getCMatMulLayerClass().class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCMatMulLayerWithDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CMatMulLayer */

// Creates a matrix multiplication layer with the specified descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulLayer/init(descriptor:)
func (cc _CMatMulLayerClass) LayerWithDescriptor(descriptor IMLCMatMulDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CMatMulLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CMatMulLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CMatMulLayer */

// The configuration object you use to create the matrix multiplication layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulLayer/descriptor
func (c_ CMatMulLayer) Descriptor() IMLCMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCMatMulLayer */


