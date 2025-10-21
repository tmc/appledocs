// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type ScreenshotOutput struct {
	objectivec.Object
}

// ScreenshotOutputFrom constructs a [ScreenshotOutput] from an unsafe.Pointer.
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


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/fileURL
func (s_ ScreenshotOutput) FileURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("fileURL"))
	return rv
}


// SetFileURL sets the value of the fileURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/fileURL
func (s_ ScreenshotOutput) SetFileURL(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) HdrImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](s_.ID, objc.Sel("hdrImage"))
	return rv
}


// SetHdrImage sets the value of the hdrImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) SetHdrImage(value coregraphics.CGImageRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHdrImage:"), value)
}



