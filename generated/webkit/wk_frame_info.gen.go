// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FrameInfo] class.
var (
	FrameInfoClass     _FrameInfoClass
	FrameInfoClassOnce sync.Once
)

func getFrameInfoClass() _FrameInfoClass {
	FrameInfoClassOnce.Do(func() {
		FrameInfoClass = _FrameInfoClass{objc.GetClass("WKFrameInfo")}
	})
	return FrameInfoClass
}

type _FrameInfoClass struct {
	class objc.Class
}

// An interface definition for the [FrameInfo] class.
type IFrameInfo interface {
	objectivec.IObject
}

// An object that contains information about a frame on a webpage.
//
// An instance of this class is a transient, data-only object; it does not uniquely identify a frame across multiple delegate method calls.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo
type FrameInfo struct {
	objectivec.Object
}

// FrameInfoFrom constructs a [FrameInfo] from an unsafe.Pointer.
//
// An object that contains information about a frame on a webpage.
func FrameInfoFrom(ptr unsafe.Pointer) FrameInfo {
	return FrameInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.Send[FrameInfo](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FrameInfoClass) New() FrameInfo {
	rv := objc.Send[FrameInfo](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FrameInfo) Init() FrameInfo {
	rv := objc.Send[FrameInfo](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FrameInfo) Autorelease() FrameInfo {
	rv := objc.Send[FrameInfo](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFrameInfo creates a new FrameInfo instance.
func NewFrameInfo() FrameInfo {
	return getFrameInfoClass().New()
}


// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/isMainFrame
func (f_ FrameInfo) MainFrame() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("mainFrame"))
	return rv
}

// The frame’s current request.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/request
func (f_ FrameInfo) Request() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("request"))
	return rv
}



