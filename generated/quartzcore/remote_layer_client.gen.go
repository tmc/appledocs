// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RemoteLayerClient] class.
var remoteLayerClientClass = _RemoteLayerClientClass{objc.GetClass("CARemoteLayerClient")}

type _RemoteLayerClientClass struct {
	class objc.Class
}

// A legacy class for cross-process rendering. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (rc _RemoteLayerClientClass) Alloc() RemoteLayerClient {
	rv := objc.Send[RemoteLayerClient](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return remoteLayerClientClass.New()
}
// Creates a layer client from a server port. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/init(serverPort:)
func NewRemoteLayerClientWithServerPort(port unsafe.Pointer) RemoteLayerClient {
	instance := remoteLayerClientClass.Alloc()
	rv := objc.Send[RemoteLayerClient](instance.ID, objc.Sel("initWithServerPort:"), port)
	rv.Autorelease()
	return rv
}



