// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCShareableContent] class.
var (
	sCShareableContentClass     _SCShareableContentClass
	sCShareableContentClassOnce sync.Once
)

func getSCShareableContentClass() _SCShareableContentClass {
	sCShareableContentClassOnce.Do(func() {
		sCShareableContentClass = _SCShareableContentClass{objc.GetClass("SCShareableContent")}
	})
	return sCShareableContentClass
}

type _SCShareableContentClass struct {
	class objc.Class
}

// An interface definition for the [SCShareableContent] class.
type ISCShareableContent interface {
	objectivec.IObject
}

// An instance that represents a set of displays, apps, and windows that your app can capture.
//
// Use the , , and properties to create a object that specifies what display content to capture. You apply the filter to an instance of to limit its output to only the content matching your filter.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent
type SCShareableContent struct {
	objectivec.Object
}

// SCShareableContentFrom constructs a [SCShareableContent] from an unsafe.Pointer.
//
// An instance that represents a set of displays, apps, and windows that your app can capture.
func SCShareableContentFrom(ptr unsafe.Pointer) SCShareableContent {
	return SCShareableContent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCShareableContentClass) Alloc() SCShareableContent {
	rv := objc.Send[SCShareableContent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCShareableContentClass) New() SCShareableContent {
	rv := objc.Send[SCShareableContent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCShareableContent) Init() SCShareableContent {
	rv := objc.Send[SCShareableContent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCShareableContent) Autorelease() SCShareableContent {
	rv := objc.Send[SCShareableContent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCShareableContent creates a new SCShareableContent instance.
func NewSCShareableContent() SCShareableContent {
	return getSCShareableContentClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getCurrentProcessShareableContent(completionHandler:)
func (sc _SCShareableContentClass) GetCurrentProcessShareableContentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getCurrentProcessShareableContentWithCompletionHandler:"), completionHandler)
}

// The apps available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/applications
func (s_ SCShareableContent) Applications() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("applications"))
	return rv
}
// The displays available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/displays
func (s_ SCShareableContent) Displays() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("displays"))
	return rv
}
// The windows available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/windows
func (s_ SCShareableContent) Windows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("windows"))
	return rv
}


