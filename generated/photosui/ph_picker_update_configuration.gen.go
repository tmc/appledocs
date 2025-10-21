// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPickerUpdateConfiguration] class.
var (
	PHPickerUpdateConfigurationClass     _PHPickerUpdateConfigurationClass
	PHPickerUpdateConfigurationClassOnce sync.Once
)

func getPHPickerUpdateConfigurationClass() _PHPickerUpdateConfigurationClass {
	PHPickerUpdateConfigurationClassOnce.Do(func() {
		PHPickerUpdateConfigurationClass = _PHPickerUpdateConfigurationClass{objc.GetClass("PHPickerUpdateConfiguration")}
	})
	return PHPickerUpdateConfigurationClass
}

type _PHPickerUpdateConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerUpdateConfiguration] class.
type IPHPickerUpdateConfiguration interface {
	objectivec.IObject
}

// An object that defines the aspects of a photo picker’s appearance that can change while it’s presented.
//
// While a photos picker is visible, you can use an instance of this object to change its or properties. To do that, create and configure an instance of this object and pass it to the method .
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerUpdateConfiguration
type PHPickerUpdateConfiguration struct {
	objectivec.Object
}

// PHPickerUpdateConfigurationFrom constructs a [PHPickerUpdateConfiguration] from an unsafe.Pointer.
//
// An object that defines the aspects of a photo picker’s appearance that can change while it’s presented.
func PHPickerUpdateConfigurationFrom(ptr unsafe.Pointer) PHPickerUpdateConfiguration {
	return PHPickerUpdateConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerUpdateConfigurationClass) Alloc() PHPickerUpdateConfiguration {
	rv := objc.Send[PHPickerUpdateConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerUpdateConfigurationClass) New() PHPickerUpdateConfiguration {
	rv := objc.Send[PHPickerUpdateConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerUpdateConfiguration) Init() PHPickerUpdateConfiguration {
	rv := objc.Send[PHPickerUpdateConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerUpdateConfiguration) Autorelease() PHPickerUpdateConfiguration {
	rv := objc.Send[PHPickerUpdateConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerUpdateConfiguration creates a new PHPickerUpdateConfiguration instance.
func NewPHPickerUpdateConfiguration() PHPickerUpdateConfiguration {
	return getPHPickerUpdateConfigurationClass().New()
}


// The portions of a photo picker’s permiter that are borderless.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerUpdateConfiguration/edgesWithoutContentMargins
func (p_ PHPickerUpdateConfiguration) EdgesWithoutContentMargins() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("edgesWithoutContentMargins"))
	return rv
}


// SetEdgesWithoutContentMargins sets the value of the edgesWithoutContentMargins property.
// The portions of a photo picker’s permiter that are borderless.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerUpdateConfiguration/edgesWithoutContentMargins
func (p_ PHPickerUpdateConfiguration) SetEdgesWithoutContentMargins(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEdgesWithoutContentMargins:"), value)
}
// The maximum number of selections the user can make.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerUpdateConfiguration/selectionLimit
func (p_ PHPickerUpdateConfiguration) SelectionLimit() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectionLimit"))
	return rv
}


// SetSelectionLimit sets the value of the selectionLimit property.
// The maximum number of selections the user can make.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerUpdateConfiguration/selectionLimit
func (p_ PHPickerUpdateConfiguration) SetSelectionLimit(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionLimit:"), value)
}


