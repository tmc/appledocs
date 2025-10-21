// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [RemoteLayerServer] class.
type IRemoteLayerServer interface {
	objectivec.IObject
}

// A legacy class for cross-process rendering.
//
// is a legacy class for cross-process rendering. and , available with , offer an improved way to perform cross-process rendering.
//
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

// Alloc allocates a new instance without initialization.
func (rc _RemoteLayerServerClass) Alloc() RemoteLayerServer {
	rv := objc.Send[RemoteLayerServer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the (singleton) instance of the shared remote layer server.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/shared()
func (rc _RemoteLayerServerClass) SharedServer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("sharedServer"))
	return rv
}

// The port number of the server.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/serverPort
func (r_ RemoteLayerServer) ServerPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("serverPort"))
	return rv
}



