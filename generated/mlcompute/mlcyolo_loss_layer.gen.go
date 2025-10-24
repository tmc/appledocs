// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCYOLOLossLayer */


/* debug [class_header]: Header for MLCYOLOLossLayer */
// The class instance for the [CYOLOLossLayer] class.
var (
	CYOLOLossLayerClass     _CYOLOLossLayerClass
	CYOLOLossLayerClassOnce sync.Once
)

func getCYOLOLossLayerClass() _CYOLOLossLayerClass {
	CYOLOLossLayerClassOnce.Do(func() {
		CYOLOLossLayerClass = _CYOLOLossLayerClass{objc.GetClass("MLCYOLOLossLayer")}
	})
	return CYOLOLossLayerClass
}

type _CYOLOLossLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CYOLOLossLayer */
// An interface definition for the [CYOLOLossLayer] class.
type ICYOLOLossLayer interface {
	ICLossLayer
	
/* debug [class_interface_properties]: Properties for CYOLOLossLayer */
	// properties:
	YoloLossDescriptor() IMLCYOLOLossDescriptor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CYOLOLossLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CYOLOLossLayer */
// Alloc allocates a new instance without initialization.
func (cc _CYOLOLossLayerClass) Alloc() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CYOLOLossLayerClass) New() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CYOLOLossLayer) Init() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CYOLOLossLayer) Autorelease() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCYOLOLossLayer creates a new CYOLOLossLayer instance.
func NewCYOLOLossLayer() CYOLOLossLayer {
	return getCYOLOLossLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CYOLOLossLayer */
// A layer that estimates loss for the YOLO algorithm.


// A layer that estimates loss for the YOLO algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossLayer
type CYOLOLossLayer struct {
	CLossLayer
}

// CYOLOLossLayerFrom constructs a [CYOLOLossLayer] from an unsafe.Pointer.
//
// A layer that estimates loss for the YOLO algorithm.
func CYOLOLossLayerFrom(ptr unsafe.Pointer) CYOLOLossLayer {
	return CYOLOLossLayer{
		CLossLayer: CLossLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CYOLOLossLayer */

// Creates a YOLO loss layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossLayer/init(descriptor:)
func NewCYOLOLossLayerWithDescriptor(lossDescriptor IMLCYOLOLossDescriptor) CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](objc.ID(getCYOLOLossLayerClass().class), objc.Sel("layerWithDescriptor:"), lossDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCYOLOLossLayerWithDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CYOLOLossLayer */

// Creates a YOLO loss layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossLayer/init(descriptor:)
func (cc _CYOLOLossLayerClass) LayerWithDescriptor(lossDescriptor IMLCYOLOLossDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:"), lossDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CYOLOLossLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CYOLOLossLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CYOLOLossLayer */

// The configuration object you use to create the YOLO loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossLayer/yoloLossDescriptor
func (c_ CYOLOLossLayer) YoloLossDescriptor() IMLCYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](c_.ID, objc.Sel("yoloLossDescriptor"))
	return rv
}/* debug [instance_properties/getter]: yoloLossDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCYOLOLossLayer */


