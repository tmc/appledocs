// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ScreenshotConfiguration] class.
var (
	ScreenshotConfigurationClass     _ScreenshotConfigurationClass
	ScreenshotConfigurationClassOnce sync.Once
)

func getScreenshotConfigurationClass() _ScreenshotConfigurationClass {
	ScreenshotConfigurationClassOnce.Do(func() {
		ScreenshotConfigurationClass = _ScreenshotConfigurationClass{objc.GetClass("SCScreenshotConfiguration")}
	})
	return ScreenshotConfigurationClass
}

type _ScreenshotConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ScreenshotConfiguration] class.
type IScreenshotConfiguration interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type ScreenshotConfiguration struct {
	objectivec.Object
}

// ScreenshotConfigurationFrom constructs a [ScreenshotConfiguration] from an unsafe.Pointer.
func ScreenshotConfigurationFrom(ptr unsafe.Pointer) ScreenshotConfiguration {
	return ScreenshotConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScreenshotConfigurationClass) Alloc() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScreenshotConfigurationClass) New() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenshotConfiguration) Init() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenshotConfiguration) Autorelease() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenshotConfiguration creates a new ScreenshotConfiguration instance.
func NewScreenshotConfiguration() ScreenshotConfiguration {
	return getScreenshotConfigurationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (sc _ScreenshotConfigurationClass) SupportedContentTypes() []UTType {
	rv := objc.Send[[]UTType](objc.ID(sc.class), objc.Sel("supportedContentTypes"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) DestinationRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("destinationRect"))
	return rv
}


// SetDestinationRect sets the value of the destinationRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) SetDestinationRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SourceRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("sourceRect"))
	return rv
}


// SetSourceRect sets the value of the sourceRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SetSourceRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (s_ ScreenshotConfiguration) SupportedContentTypes() []UTType {
	rv := objc.Send[[]UTType](s_.ID, objc.Sel("supportedContentTypes"))
	return rv
}



