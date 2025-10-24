// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CARemoteLayerClient */


/* debug [class_header]: Header for CARemoteLayerClient */
// The class instance for the [RemoteLayerClient] class.
var (
	RemoteLayerClientClass     _RemoteLayerClientClass
	RemoteLayerClientClassOnce sync.Once
)

func getRemoteLayerClientClass() _RemoteLayerClientClass {
	RemoteLayerClientClassOnce.Do(func() {
		RemoteLayerClientClass = _RemoteLayerClientClass{objc.GetClass("CARemoteLayerClient")}
	})
	return RemoteLayerClientClass
}

type _RemoteLayerClientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RemoteLayerClient */
// An interface definition for the [RemoteLayerClient] class.
type IRemoteLayerClient interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RemoteLayerClient */
	// properties:
	ClientId() uint32 /* not a class type */
	Layer() ILayer
	SetLayer(value ILayer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RemoteLayerClient */
	// methods:
	Invalidate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RemoteLayerClient */
// Alloc allocates a new instance without initialization.
func (rc _RemoteLayerClientClass) Alloc() RemoteLayerClient {
	rv := objc.Send[RemoteLayerClient](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RemoteLayerClientClass) New() RemoteLayerClient {
	rv := objc.Send[RemoteLayerClient](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RemoteLayerClient) Init() RemoteLayerClient {
	rv := objc.Send[RemoteLayerClient](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RemoteLayerClient) Autorelease() RemoteLayerClient {
	rv := objc.Send[RemoteLayerClient](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemoteLayerClient creates a new RemoteLayerClient instance.
func NewRemoteLayerClient() RemoteLayerClient {
	return getRemoteLayerClientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RemoteLayerClient */
// A legacy class for cross-process rendering.
//
// is a legacy class for cross-process rendering. and , available with , offer an improved way to perform cross-process rendering.


// A legacy class for cross-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient
type RemoteLayerClient struct {
	objectivec.Object
}

// RemoteLayerClientFrom constructs a [RemoteLayerClient] from an unsafe.Pointer.
//
// A legacy class for cross-process rendering.
func RemoteLayerClientFrom(ptr unsafe.Pointer) RemoteLayerClient {
	return RemoteLayerClient{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RemoteLayerClient */

// Creates a layer client from a server port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/init(serverPort:)
func NewRemoteLayerClientWithServerPort(port objectivec.IObject) RemoteLayerClient {
	instance := getRemoteLayerClientClass().Alloc()
	rv := objc.Send[RemoteLayerClient](instance.ID, objc.Sel("initWithServerPort:"), port)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRemoteLayerClientWithServerPort */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RemoteLayerClient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RemoteLayerClient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RemoteLayerClient */

// Invalidates a remote layer client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/invalidate()
func (r_ RemoteLayerClient) Invalidate() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RemoteLayerClient */

// The ID of the remote layer client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/clientId
func (r_ RemoteLayerClient) ClientId() uint32 /* not a class type */ {
	rv := objc.Send[uint32](r_.ID, objc.Sel("clientId"))
	return rv
}/* debug [instance_properties/getter]: clientId */


// The layer associated with the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/layer
func (r_ RemoteLayerClient) Layer() ILayer {
	rv := objc.Send[Layer](r_.ID, objc.Sel("layer"))
	return rv
}/* debug [instance_properties/getter]: layer */


// The layer associated with the remote client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/layer
func (r_ RemoteLayerClient) SetLayer(value ILayer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:"), value)
}/* debug [instance_properties/setter]: layer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CARemoteLayerClient */


