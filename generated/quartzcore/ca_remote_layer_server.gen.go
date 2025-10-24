// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CARemoteLayerServer */


/* debug [class_header]: Header for CARemoteLayerServer */
// The class instance for the [RemoteLayerServer] class.
var (
	RemoteLayerServerClass     _RemoteLayerServerClass
	RemoteLayerServerClassOnce sync.Once
)

func getRemoteLayerServerClass() _RemoteLayerServerClass {
	RemoteLayerServerClassOnce.Do(func() {
		RemoteLayerServerClass = _RemoteLayerServerClass{objc.GetClass("CARemoteLayerServer")}
	})
	return RemoteLayerServerClass
}

type _RemoteLayerServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RemoteLayerServer */
// An interface definition for the [RemoteLayerServer] class.
type IRemoteLayerServer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RemoteLayerServer */
	// properties:
	ServerPort() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RemoteLayerServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RemoteLayerServer */
// Alloc allocates a new instance without initialization.
func (rc _RemoteLayerServerClass) Alloc() RemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RemoteLayerServerClass) New() RemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RemoteLayerServer) Init() RemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RemoteLayerServer) Autorelease() RemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemoteLayerServer creates a new RemoteLayerServer instance.
func NewRemoteLayerServer() RemoteLayerServer {
	return getRemoteLayerServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RemoteLayerServer */
// A legacy class for cross-process rendering.
//
// is a legacy class for cross-process rendering. and , available with , offer an improved way to perform cross-process rendering.


// A legacy class for cross-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer
type RemoteLayerServer struct {
	objectivec.Object
}

// RemoteLayerServerFrom constructs a [RemoteLayerServer] from an unsafe.Pointer.
//
// A legacy class for cross-process rendering.
func RemoteLayerServerFrom(ptr unsafe.Pointer) RemoteLayerServer {
	return RemoteLayerServer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RemoteLayerServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RemoteLayerServer */

// Returns the (singleton) instance of the shared remote layer server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/shared()
func (rc _RemoteLayerServerClass) SharedServer() IRemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](objc.ID(rc.class), objc.Sel("sharedServer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedServer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RemoteLayerServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RemoteLayerServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RemoteLayerServer */

// The port number of the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/serverPort
func (r_ RemoteLayerServer) ServerPort() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("serverPort"))
	return rv
}/* debug [instance_properties/getter]: serverPort */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CARemoteLayerServer */



