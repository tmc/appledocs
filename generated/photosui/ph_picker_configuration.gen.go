// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPickerConfiguration] class.
var (
	PHPickerConfigurationClass     _PHPickerConfigurationClass
	PHPickerConfigurationClassOnce sync.Once
)

func getPHPickerConfigurationClass() _PHPickerConfigurationClass {
	PHPickerConfigurationClassOnce.Do(func() {
		PHPickerConfigurationClass = _PHPickerConfigurationClass{objc.GetClass("PHPickerConfiguration")}
	})
	return PHPickerConfigurationClass
}

type _PHPickerConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerConfiguration] class.
type IPHPickerConfiguration interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that contains information about how to configure a picker view controller.

// An object that contains information about how to configure a picker view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class
type PHPickerConfiguration struct {
	objectivec.Object
}

// PHPickerConfigurationFrom constructs a [PHPickerConfiguration] from an unsafe.Pointer.
//
// An object that contains information about how to configure a picker view controller.
func PHPickerConfigurationFrom(ptr unsafe.Pointer) PHPickerConfiguration {
	return PHPickerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerConfigurationClass) Alloc() PHPickerConfiguration {
	rv := objc.Send[PHPickerConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerConfigurationClass) New() PHPickerConfiguration {
	rv := objc.Send[PHPickerConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerConfiguration) Init() PHPickerConfiguration {
	rv := objc.Send[PHPickerConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerConfiguration) Autorelease() PHPickerConfiguration {
	rv := objc.Send[PHPickerConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerConfiguration creates a new PHPickerConfiguration instance.
func NewPHPickerConfiguration() PHPickerConfiguration {
	return getPHPickerConfigurationClass().New()
}
