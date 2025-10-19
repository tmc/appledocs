// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCStream] class.
var (
	sCStreamClass     _SCStreamClass
	sCStreamClassOnce sync.Once
)

func getSCStreamClass() _SCStreamClass {
	sCStreamClassOnce.Do(func() {
		sCStreamClass = _SCStreamClass{objc.GetClass("SCStream")}
	})
	return sCStreamClass
}

type _SCStreamClass struct {
	class objc.Class
}

// An interface definition for the [SCStream] class.
type ISCStream interface {
	objectivec.IObject
	StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateContentFilterCompletionHandler(contentFilter unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An instance that represents a stream of shareable content.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream
type SCStream struct {
	objectivec.Object
}

// SCStreamFrom constructs a [SCStream] from an unsafe.Pointer.
//
// An instance that represents a stream of shareable content.
func SCStreamFrom(ptr unsafe.Pointer) SCStream {
	return SCStream{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCStreamClass) Alloc() SCStream {
	rv := objc.Send[SCStream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCStreamClass) New() SCStream {
	rv := objc.Send[SCStream](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCStream) Init() SCStream {
	rv := objc.Send[SCStream](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCStream) Autorelease() SCStream {
	rv := objc.Send[SCStream](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCStream creates a new SCStream instance.
func NewSCStream() SCStream {
	return getSCStreamClass().New()
}


// Starts the stream with a callback to indicate whether it successfully starts.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/startCapture(completionHandler:)
func (s_ SCStream) StartCaptureWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("startCaptureWithCompletionHandler:"), completionHandler)
}
// Updates the stream by applying a new content filter.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateContentFilter(_:completionHandler:)
func (s_ SCStream) UpdateContentFilterCompletionHandler(contentFilter unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateContentFilter:completionHandler:"), contentFilter, completionHandler)
}


