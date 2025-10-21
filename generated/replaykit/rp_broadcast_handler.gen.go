// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RPBroadcastHandler] class.
var (
	RPBroadcastHandlerClass     _RPBroadcastHandlerClass
	RPBroadcastHandlerClassOnce sync.Once
)

func getRPBroadcastHandlerClass() _RPBroadcastHandlerClass {
	RPBroadcastHandlerClassOnce.Do(func() {
		RPBroadcastHandlerClass = _RPBroadcastHandlerClass{objc.GetClass("RPBroadcastHandler")}
	})
	return RPBroadcastHandlerClass
}

type _RPBroadcastHandlerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastHandler] class.
type IRPBroadcastHandler interface {
	objectivec.IObject
	UpdateBroadcastURL(broadcastURL unsafe.Pointer)
	UpdateServiceInfo(serviceInfo unsafe.Pointer)
}

// An object that sends messages to the broadcasting app.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler
type RPBroadcastHandler struct {
	objectivec.Object
}

// RPBroadcastHandlerFrom constructs a [RPBroadcastHandler] from an unsafe.Pointer.
//
// An object that sends messages to the broadcasting app.
func RPBroadcastHandlerFrom(ptr unsafe.Pointer) RPBroadcastHandler {
	return RPBroadcastHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastHandlerClass) Alloc() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastHandlerClass) New() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastHandler) Init() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastHandler) Autorelease() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastHandler creates a new RPBroadcastHandler instance.
func NewRPBroadcastHandler() RPBroadcastHandler {
	return getRPBroadcastHandlerClass().New()
}


// Sends the current broadcast URL to the broadcast controller.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler/updateBroadcast(_:)
func (r_ RPBroadcastHandler) UpdateBroadcastURL(broadcastURL unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateBroadcastURL:"), broadcastURL)
}

// Sends information about the current broadcast to the broadcasting app.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler/updateServiceInfo(_:)
func (r_ RPBroadcastHandler) UpdateServiceInfo(serviceInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateServiceInfo:"), serviceInfo)
}



