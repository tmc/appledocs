// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCStreamConfiguration] class.
var (
	sCStreamConfigurationClass     _SCStreamConfigurationClass
	sCStreamConfigurationClassOnce sync.Once
)

func getSCStreamConfigurationClass() _SCStreamConfigurationClass {
	sCStreamConfigurationClassOnce.Do(func() {
		sCStreamConfigurationClass = _SCStreamConfigurationClass{objc.GetClass("SCStreamConfiguration")}
	})
	return sCStreamConfigurationClass
}

type _SCStreamConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [SCStreamConfiguration] class.
type ISCStreamConfiguration interface {
	objectivec.IObject
}

// An instance that provides the output configuration for a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration
type SCStreamConfiguration struct {
	objectivec.Object
}

// SCStreamConfigurationFrom constructs a [SCStreamConfiguration] from an unsafe.Pointer.
//
// An instance that provides the output configuration for a stream.
func SCStreamConfigurationFrom(ptr unsafe.Pointer) SCStreamConfiguration {
	return SCStreamConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCStreamConfigurationClass) Alloc() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCStreamConfigurationClass) New() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCStreamConfiguration) Init() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCStreamConfiguration) Autorelease() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCStreamConfiguration creates a new SCStreamConfiguration instance.
func NewSCStreamConfiguration() SCStreamConfiguration {
	return getSCStreamConfigurationClass().New()
}




