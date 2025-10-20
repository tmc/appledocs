// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SequenceRequestHandler] class.
var (
	SequenceRequestHandlerClass     _SequenceRequestHandlerClass
	SequenceRequestHandlerClassOnce sync.Once
)

func getSequenceRequestHandlerClass() _SequenceRequestHandlerClass {
	SequenceRequestHandlerClassOnce.Do(func() {
		SequenceRequestHandlerClass = _SequenceRequestHandlerClass{objc.GetClass("VNSequenceRequestHandler")}
	})
	return SequenceRequestHandlerClass
}

type _SequenceRequestHandlerClass struct {
	class objc.Class
}

// An interface definition for the [SequenceRequestHandler] class.
type ISequenceRequestHandler interface {
	objectivec.IObject
}

// An object that processes image-analysis requests for each frame in a sequence.
//
// Instantiate this handler to perform Vision requests on a series of images. Unlike the , you don’t specify the image on creation. Instead, you supply each image frame one by one as you continue to call one of the methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler
type SequenceRequestHandler struct {
	objectivec.Object
}

// SequenceRequestHandlerFrom constructs a [SequenceRequestHandler] from an unsafe.Pointer.
//
// An object that processes image-analysis requests for each frame in a sequence.
func SequenceRequestHandlerFrom(ptr unsafe.Pointer) SequenceRequestHandler {
	return SequenceRequestHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SequenceRequestHandlerClass) Alloc() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SequenceRequestHandlerClass) New() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SequenceRequestHandler) Init() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SequenceRequestHandler) Autorelease() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSequenceRequestHandler creates a new SequenceRequestHandler instance.
func NewSequenceRequestHandler() SequenceRequestHandler {
	return getSequenceRequestHandlerClass().New()
}




