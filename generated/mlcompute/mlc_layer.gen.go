// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CLayer] class.
type ICLayer interface {
	objectivec.IObject
	// properties:
	DeviceType() CDeviceType /* not a class type */
	SetDeviceType(value CDeviceType /* not a class type */)
	IsDebuggingEnabled() bool
	SetIsDebuggingEnabled(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LayerID() int
	SetLayerID(value int)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CLayerClass) Alloc() CLayer {
	rv := objc.Send[CLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A device type that indicates where the system executes the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/devicetype
func (c_ CLayer) DeviceType() CDeviceType /* not a class type */ {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("deviceType"))
	return rv
}


// A device type that indicates where the system executes the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/devicetype
func (c_ CLayer) SetDeviceType(value CDeviceType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeviceType:"), value)
}


// A Boolean that indicates whether you choose to debug the layer when executing a graph that includes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/isdebuggingenabled
func (c_ CLayer) IsDebuggingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDebuggingEnabled"))
	return rv
}


// A Boolean that indicates whether you choose to debug the layer when executing a graph that includes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/isdebuggingenabled
func (c_ CLayer) SetIsDebuggingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDebuggingEnabled:"), value)
}


// A string that helps identify this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/label
func (c_ CLayer) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}


// A string that helps identify this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/label
func (c_ CLayer) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}


// A unique number that identifies each layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/layerid
func (c_ CLayer) LayerID() int {
	rv := objc.Send[int](c_.ID, objc.Sel("layerID"))
	return rv
}


// A unique number that identifies each layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayer/layerid
func (c_ CLayer) SetLayerID(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayerID:"), value)
}



