// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCPoolingLayer */


/* debug [class_header]: Header for MLCPoolingLayer */
// The class instance for the [CPoolingLayer] class.
var (
	CPoolingLayerClass     _CPoolingLayerClass
	CPoolingLayerClassOnce sync.Once
)

func getCPoolingLayerClass() _CPoolingLayerClass {
	CPoolingLayerClassOnce.Do(func() {
		CPoolingLayerClass = _CPoolingLayerClass{objc.GetClass("MLCPoolingLayer")}
	})
	return CPoolingLayerClass
}

type _CPoolingLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CPoolingLayer */
// An interface definition for the [CPoolingLayer] class.
type ICPoolingLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CPoolingLayer */
	// properties:
	Descriptor() IMLCPoolingDescriptor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CPoolingLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CPoolingLayer */
// Alloc allocates a new instance without initialization.
func (cc _CPoolingLayerClass) Alloc() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CPoolingLayerClass) New() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPoolingLayer) Init() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPoolingLayer) Autorelease() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPoolingLayer creates a new CPoolingLayer instance.
func NewCPoolingLayer() CPoolingLayer {
	return getCPoolingLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CPoolingLayer */
// A layer that summarizes the average presence of a feature.


// A layer that summarizes the average presence of a feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingLayer
type CPoolingLayer struct {
	CLayer
}

// CPoolingLayerFrom constructs a [CPoolingLayer] from an unsafe.Pointer.
//
// A layer that summarizes the average presence of a feature.
func CPoolingLayerFrom(ptr unsafe.Pointer) CPoolingLayer {
	return CPoolingLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CPoolingLayer */

// Creates a pooling layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingLayer/init(descriptor:)
func NewCPoolingLayerWithDescriptor(descriptor IMLCPoolingDescriptor) CPoolingLayer {
	rv := objc.Send[CPoolingLayer](objc.ID(getCPoolingLayerClass().class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCPoolingLayerWithDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CPoolingLayer */

// Creates a pooling layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingLayer/init(descriptor:)
func (cc _CPoolingLayerClass) LayerWithDescriptor(descriptor IMLCPoolingDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CPoolingLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CPoolingLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CPoolingLayer */

// The configuration object you use to create the pooling layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingLayer/descriptor
func (c_ CPoolingLayer) Descriptor() IMLCPoolingDescriptor {
	rv := objc.Send[CPoolingDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCPoolingLayer */


