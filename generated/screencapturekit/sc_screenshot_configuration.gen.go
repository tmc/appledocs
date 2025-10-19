// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCScreenshotConfiguration] class.
var (
	sCScreenshotConfigurationClass     _SCScreenshotConfigurationClass
	sCScreenshotConfigurationClassOnce sync.Once
)

func getSCScreenshotConfigurationClass() _SCScreenshotConfigurationClass {
	sCScreenshotConfigurationClassOnce.Do(func() {
		sCScreenshotConfigurationClass = _SCScreenshotConfigurationClass{objc.GetClass("SCScreenshotConfiguration")}
	})
	return sCScreenshotConfigurationClass
}

type _SCScreenshotConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [SCScreenshotConfiguration] class.
type ISCScreenshotConfiguration interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type SCScreenshotConfiguration struct {
	objectivec.Object
}

// SCScreenshotConfigurationFrom constructs a [SCScreenshotConfiguration] from an unsafe.Pointer.
func SCScreenshotConfigurationFrom(ptr unsafe.Pointer) SCScreenshotConfiguration {
	return SCScreenshotConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCScreenshotConfigurationClass) Alloc() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCScreenshotConfigurationClass) New() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCScreenshotConfiguration) Init() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCScreenshotConfiguration) Autorelease() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotConfiguration creates a new SCScreenshotConfiguration instance.
func NewSCScreenshotConfiguration() SCScreenshotConfiguration {
	return getSCScreenshotConfigurationClass().New()
}




