// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCScreenshotOutput] class.
var (
	sCScreenshotOutputClass     _SCScreenshotOutputClass
	sCScreenshotOutputClassOnce sync.Once
)

func getSCScreenshotOutputClass() _SCScreenshotOutputClass {
	sCScreenshotOutputClassOnce.Do(func() {
		sCScreenshotOutputClass = _SCScreenshotOutputClass{objc.GetClass("SCScreenshotOutput")}
	})
	return sCScreenshotOutputClass
}

type _SCScreenshotOutputClass struct {
	class objc.Class
}

// An interface definition for the [SCScreenshotOutput] class.
type ISCScreenshotOutput interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type SCScreenshotOutput struct {
	objectivec.Object
}

// SCScreenshotOutputFrom constructs a [SCScreenshotOutput] from an unsafe.Pointer.
func SCScreenshotOutputFrom(ptr unsafe.Pointer) SCScreenshotOutput {
	return SCScreenshotOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCScreenshotOutputClass) Alloc() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCScreenshotOutputClass) New() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCScreenshotOutput) Init() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCScreenshotOutput) Autorelease() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotOutput creates a new SCScreenshotOutput instance.
func NewSCScreenshotOutput() SCScreenshotOutput {
	return getSCScreenshotOutputClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ SCScreenshotOutput) HdrImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](s_.ID, objc.Sel("hdrImage"))
	return rv
}

// SetHdrImage sets the value of the hdrImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ SCScreenshotOutput) SetHdrImage(value coregraphics.CGImageRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHdrImage:"), value)
}


