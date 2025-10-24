// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCLayer */


/* debug [class_header]: Header for MLCLayer */
// The class instance for the [CLayer] class.
var (
	CLayerClass     _CLayerClass
	CLayerClassOnce sync.Once
)

func getCLayerClass() _CLayerClass {
	CLayerClassOnce.Do(func() {
		CLayerClass = _CLayerClass{objc.GetClass("MLCLayer")}
	})
	return CLayerClass
}

type _CLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLayer */
// An interface definition for the [CLayer] class.
type ICLayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CLayer */
	// properties:
	DeviceType() CDeviceType
	IsDebuggingEnabled() bool
	SetIsDebuggingEnabled(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LayerID() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLayer */
// Alloc allocates a new instance without initialization.
func (cc _CLayerClass) Alloc() CLayer {
	rv := objc.Send[CLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CLayerClass) New() CLayer {
	rv := objc.Send[CLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLayer) Init() CLayer {
	rv := objc.Send[CLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLayer) Autorelease() CLayer {
	rv := objc.Send[CLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLayer creates a new CLayer instance.
func NewCLayer() CLayer {
	return getCLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLayer */
// The base class for all framework layers.
//
// This class defines a polymorphic interface for subclasses. There are subclasses for each supported neural network layer type. Use the appropriate subclass initializer to create a layer object.


// The base class for all framework layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer
type CLayer struct {
	objectivec.Object
}

// CLayerFrom constructs a [CLayer] from an unsafe.Pointer.
//
// The base class for all framework layers.
func CLayerFrom(ptr unsafe.Pointer) CLayer {
	return CLayer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLayer */

// Returns a Boolean that indicates whether instances of this layer accept source tensors for the data type and device that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/supportsDataType(_:on:)
func (cc _CLayerClass) SupportsDataTypeOnDevice(dataType CDataType, device IMLCDevice) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supportsDataType:onDevice:"), dataType, device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsDataTypeOnDevice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLayer */

// A device type that indicates where the system executes the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/deviceType
func (c_ CLayer) DeviceType() CDeviceType {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("deviceType"))
	return rv
}/* debug [instance_properties/getter]: deviceType */


// A Boolean that indicates whether you choose to debug the layer when executing a graph that includes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/isDebuggingEnabled
func (c_ CLayer) IsDebuggingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDebuggingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDebuggingEnabled */


// A Boolean that indicates whether you choose to debug the layer when executing a graph that includes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/isDebuggingEnabled
func (c_ CLayer) SetIsDebuggingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDebuggingEnabled:"), value)
}/* debug [instance_properties/setter]: isDebuggingEnabled */


// A string that helps identify this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/label
func (c_ CLayer) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that helps identify this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/label
func (c_ CLayer) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// A unique number that identifies each layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer/layerID
func (c_ CLayer) LayerID() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("layerID"))
	return rv
}/* debug [instance_properties/getter]: layerID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLayer */



