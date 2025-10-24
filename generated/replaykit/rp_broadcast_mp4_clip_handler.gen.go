// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RPBroadcastMP4ClipHandler] class.
var (
	RPBroadcastMP4ClipHandlerClass     _RPBroadcastMP4ClipHandlerClass
	RPBroadcastMP4ClipHandlerClassOnce sync.Once
)

func getRPBroadcastMP4ClipHandlerClass() _RPBroadcastMP4ClipHandlerClass {
	RPBroadcastMP4ClipHandlerClassOnce.Do(func() {
		RPBroadcastMP4ClipHandlerClass = _RPBroadcastMP4ClipHandlerClass{objc.GetClass("RPBroadcastMP4ClipHandler")}
	})
	return RPBroadcastMP4ClipHandlerClass
}

type _RPBroadcastMP4ClipHandlerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastMP4ClipHandler] class.
type IRPBroadcastMP4ClipHandler interface {
	IRPBroadcastHandler
	// properties:
	// methods:
}

// An object that processes MP4 movie clips from ReplayKit.
//
// Subclass this class to handle movie clips as ReplayKit records them during the broadcast. The system calls when a movie clip is available for processing.


// An object that processes MP4 movie clips from ReplayKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastMP4ClipHandler
type RPBroadcastMP4ClipHandler struct {
	RPBroadcastHandler
}

// RPBroadcastMP4ClipHandlerFrom constructs a [RPBroadcastMP4ClipHandler] from an unsafe.Pointer.
//
// An object that processes MP4 movie clips from ReplayKit.
func RPBroadcastMP4ClipHandlerFrom(ptr unsafe.Pointer) RPBroadcastMP4ClipHandler {
	return RPBroadcastMP4ClipHandler{
		RPBroadcastHandler: RPBroadcastHandlerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastMP4ClipHandlerClass) Alloc() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastMP4ClipHandlerClass) New() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastMP4ClipHandler) Init() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastMP4ClipHandler) Autorelease() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastMP4ClipHandler creates a new RPBroadcastMP4ClipHandler instance.
func NewRPBroadcastMP4ClipHandler() RPBroadcastMP4ClipHandler {
	return getRPBroadcastMP4ClipHandlerClass().New()
}



