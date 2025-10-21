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
}

// An object that contains information about how to configure a picker view controller.
//
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




// Creates a new configuration object for a photo library.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/initWithPhotoLibrary:
func NewPHPickerConfigurationWithPhotoLibrary(photoLibrary unsafe.Pointer) PHPickerConfiguration {
	instance := getPHPickerConfigurationClass().Alloc()
	rv := objc.Send[PHPickerConfiguration](instance.ID, objc.Sel("initWithPhotoLibrary:"), photoLibrary)
	rv.Autorelease()
	return rv
}


// The aspects of a photo picker’s default appearance that your app can disable.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/disabledCapabilities
func (p_ PHPickerConfiguration) DisabledCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("disabledCapabilities"))
	return rv
}


// SetDisabledCapabilities sets the value of the disabledCapabilities property.
// The aspects of a photo picker’s default appearance that your app can disable.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/disabledCapabilities
func (p_ PHPickerConfiguration) SetDisabledCapabilities(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisabledCapabilities:"), value)
}

// The portions of a photo picker’s perimeter that are borderless.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/edgesWithoutContentMargins
func (p_ PHPickerConfiguration) EdgesWithoutContentMargins() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("edgesWithoutContentMargins"))
	return rv
}


// SetEdgesWithoutContentMargins sets the value of the edgesWithoutContentMargins property.
// The portions of a photo picker’s perimeter that are borderless.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/edgesWithoutContentMargins
func (p_ PHPickerConfiguration) SetEdgesWithoutContentMargins(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEdgesWithoutContentMargins:"), value)
}

// The filter you apply to restrict the asset types the picker displays.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/filter
func (p_ PHPickerConfiguration) Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("filter"))
	return rv
}


// SetFilter sets the value of the filter property.
// The filter you apply to restrict the asset types the picker displays.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/filter
func (p_ PHPickerConfiguration) SetFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFilter:"), value)
}

// A layout type for the photos in the picker’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/mode
func (p_ PHPickerConfiguration) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// A layout type for the photos in the picker’s view.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/mode
func (p_ PHPickerConfiguration) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMode:"), value)
}

// A mode that determines which representation to use if an asset contains more than one.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/preferredAssetRepresentationMode
func (p_ PHPickerConfiguration) PreferredAssetRepresentationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredAssetRepresentationMode"))
	return rv
}


// SetPreferredAssetRepresentationMode sets the value of the preferredAssetRepresentationMode property.
// A mode that determines which representation to use if an asset contains more than one.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/preferredAssetRepresentationMode
func (p_ PHPickerConfiguration) SetPreferredAssetRepresentationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredAssetRepresentationMode:"), value)
}

// An array of asset identifiers to preselect in the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/preselectedAssetIdentifiers
func (p_ PHPickerConfiguration) PreselectedAssetIdentifiers() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("preselectedAssetIdentifiers"))
	return rv
}


// SetPreselectedAssetIdentifiers sets the value of the preselectedAssetIdentifiers property.
// An array of asset identifiers to preselect in the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/preselectedAssetIdentifiers
func (p_ PHPickerConfiguration) SetPreselectedAssetIdentifiers(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreselectedAssetIdentifiers:"), nsArray)
}

// The selection behavior for the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/selection
func (p_ PHPickerConfiguration) Selection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selection"))
	return rv
}


// SetSelection sets the value of the selection property.
// The selection behavior for the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/selection
func (p_ PHPickerConfiguration) SetSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelection:"), value)
}

// The maximum number of selections the user can make.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/selectionLimit
func (p_ PHPickerConfiguration) SelectionLimit() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectionLimit"))
	return rv
}


// SetSelectionLimit sets the value of the selectionLimit property.
// The maximum number of selections the user can make.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfiguration-c.class/selectionLimit
func (p_ PHPickerConfiguration) SetSelectionLimit(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionLimit:"), value)
}


