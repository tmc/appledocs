// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScreenshotOutput] class.
var (
	ScreenshotOutputClass     _ScreenshotOutputClass
	ScreenshotOutputClassOnce sync.Once
)

func getScreenshotOutputClass() _ScreenshotOutputClass {
	ScreenshotOutputClassOnce.Do(func() {
		ScreenshotOutputClass = _ScreenshotOutputClass{objc.GetClass("SCScreenshotOutput")}
	})
	return ScreenshotOutputClass
}

type _ScreenshotOutputClass struct {
	class objc.Class
}

// An interface definition for the [ScreenshotOutput] class.
type IScreenshotOutput interface {
	objectivec.IObject
	// properties:
	HdrImage() ImageRef /* not a class type */
	SetHdrImage(value ImageRef /* not a class type */)
	SdrImage() ImageRef /* not a class type */
	SetSdrImage(value ImageRef /* not a class type */)
	FileURL() objc.IObject /* cross-framework: NSURL */
	SetFileURL(value objc.IObject /* cross-framework: NSURL */)
	// methods:
}

// An object that contains all images requested by the client.


// An object that contains all images requested by the client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type ScreenshotOutput struct {
	objectivec.Object
}

// ScreenshotOutputFrom constructs a [ScreenshotOutput] from an unsafe.Pointer.
//
// An object that contains all images requested by the client.
func ScreenshotOutputFrom(ptr unsafe.Pointer) ScreenshotOutput {
	return ScreenshotOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScreenshotOutputClass) Alloc() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScreenshotOutputClass) New() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenshotOutput) Init() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenshotOutput) Autorelease() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenshotOutput creates a new ScreenshotOutput instance.
func NewScreenshotOutput() ScreenshotOutput {
	return getScreenshotOutputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) HdrImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("hdrImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) SetHdrImage(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHdrImage:"), value)
}


// An output property that specifies the standard dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/sdrImage
func (s_ ScreenshotOutput) SdrImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("sdrImage"))
	return rv
}


// An output property that specifies the standard dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/sdrImage
func (s_ ScreenshotOutput) SetSdrImage(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSdrImage:"), value)
}


// A URL property that specifies the location of the saved image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotoutput/fileurl
func (s_ ScreenshotOutput) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("fileURL"))
	return rv
}


// A URL property that specifies the location of the saved image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotoutput/fileurl
func (s_ ScreenshotOutput) SetFileURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}



