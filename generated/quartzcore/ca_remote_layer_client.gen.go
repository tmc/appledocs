// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [RemoteLayerClient] class.
type IRemoteLayerClient interface {
	objectivec.IObject
	Invalidate()
	ClientId() uint32
	Layer() CALayer
	SetLayer(value ILayer)
}

// A legacy class for cross-process rendering.
//
// is a legacy class for cross-process rendering. and , available with , offer an improved way to perform cross-process rendering.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a layer client from a server port.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/init(serverPort:)
func NewRemoteLayerClientWithServerPort(port unsafe.Pointer) RemoteLayerClient {
	instance := getRemoteLayerClientClass().Alloc()
	rv := objc.Send[RemoteLayerClient](instance.ID, objc.Sel("initWithServerPort:"), port)
	rv.Autorelease()
	return rv
}


// Invalidates a remote layer client.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/invalidate()
func (r_ RemoteLayerClient) Invalidate() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidate"))
}

// The ID of the remote layer client.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/clientId
func (r_ RemoteLayerClient) ClientId() uint32 {
	rv := objc.Send[uint32](r_.ID, objc.Sel("clientId"))
	return rv
}

// The layer associated with the remote client.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/layer
func (r_ RemoteLayerClient) Layer() CALayer {
	rv := objc.Send[CALayer](r_.ID, objc.Sel("layer"))
	return rv
}


// SetLayer sets the value of the layer property.
// The layer associated with the remote client.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/layer
func (r_ RemoteLayerClient) SetLayer(value ILayer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:"), value)
}


