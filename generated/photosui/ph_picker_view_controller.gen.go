// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PHPickerViewController] class.
var (
	PHPickerViewControllerClass     _PHPickerViewControllerClass
	PHPickerViewControllerClassOnce sync.Once
)

func getPHPickerViewControllerClass() _PHPickerViewControllerClass {
	PHPickerViewControllerClassOnce.Do(func() {
		PHPickerViewControllerClass = _PHPickerViewControllerClass{objc.GetClass("PHPickerViewController")}
	})
	return PHPickerViewControllerClass
}

type _PHPickerViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerViewController] class.
type IPHPickerViewController interface {
	appkit.IViewController
	DeselectAssetsWithIdentifiers(identifiers []string)
	MoveAssetWithIdentifierAfterAssetWithIdentifier(identifier appkit.string, afterIdentifier appkit.string)
	ScrollToInitialPosition()
	UpdatePickerUsingConfiguration(configuration IPHPickerUpdateConfiguration)
	ZoomIn()
	ZoomOut()
}

// A view controller that provides the user interface for choosing assets from the photo library.
//
// The class is an alternative to . improves stability and reliability, and includes several benefits to developers and users, such as the following: Deferred image loading and recovery UI Reliable handling of large and complex assets, like RAW and panoramic images User-selectable assets that aren’t available for Configuration of the picker to display only Live Photos Availability of objects without library access Stricter validations against invalid inputs
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController
type PHPickerViewController struct {
	appkit.ViewController
}

// PHPickerViewControllerFrom constructs a [PHPickerViewController] from an unsafe.Pointer.
//
// A view controller that provides the user interface for choosing assets from the photo library.
func PHPickerViewControllerFrom(ptr unsafe.Pointer) PHPickerViewController {
	return PHPickerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerViewControllerClass) Alloc() PHPickerViewController {
	rv := objc.Send[PHPickerViewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerViewControllerClass) New() PHPickerViewController {
	rv := objc.Send[PHPickerViewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerViewController) Init() PHPickerViewController {
	rv := objc.Send[PHPickerViewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerViewController) Autorelease() PHPickerViewController {
	rv := objc.Send[PHPickerViewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerViewController creates a new PHPickerViewController instance.
func NewPHPickerViewController() PHPickerViewController {
	return getPHPickerViewControllerClass().New()
}




// Creates a new picker view controller with the configuration you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/initWithConfiguration:
func NewPHPickerViewControllerWithConfiguration(configuration IPHPickerConfiguration) PHPickerViewController {
	instance := getPHPickerViewControllerClass().Alloc()
	rv := objc.Send[PHPickerViewController](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}


// Deselects assets that are in a selected state.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/deselectAssets(withIdentifiers:)
func (p_ PHPickerViewController) DeselectAssetsWithIdentifiers(identifiers []string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("deselectAssetsWithIdentifiers:"), identifiers)
}

// Reorders assets that are in a selected state.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/moveAsset(withIdentifier:afterAssetWithIdentifier:)
func (p_ PHPickerViewController) MoveAssetWithIdentifierAfterAssetWithIdentifier(identifier appkit.string, afterIdentifier appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("moveAssetWithIdentifier:afterAssetWithIdentifier:"), identifier, afterIdentifier)
}

// Resets the visible photo thumbnails by scrolling the view to the picker’s initial position.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/scrollToInitialPosition()
func (p_ PHPickerViewController) ScrollToInitialPosition() {
	objc.Send[objc.ID](p_.ID, objc.Sel("scrollToInitialPosition"))
}

// Customizes your app’s photo picker according to the given configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/updatePickerUsingConfiguration:
func (p_ PHPickerViewController) UpdatePickerUsingConfiguration(configuration IPHPickerUpdateConfiguration) {
	objc.Send[objc.ID](p_.ID, objc.Sel("updatePickerUsingConfiguration:"), configuration)
}

// Changes the picker’s content scale by making the photo thumbnails larger in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/zoomIn()
func (p_ PHPickerViewController) ZoomIn() {
	objc.Send[objc.ID](p_.ID, objc.Sel("zoomIn"))
}

// Changes the picker’s content scale by making the photo thumbnails smaller in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/zoomOut()
func (p_ PHPickerViewController) ZoomOut() {
	objc.Send[objc.ID](p_.ID, objc.Sel("zoomOut"))
}

// The configuration you specify when creating the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/configuration-3vf53
func (p_ PHPickerViewController) Configuration() PHPickerConfiguration {
	rv := objc.Send[PHPickerConfiguration](p_.ID, objc.Sel("configuration"))
	return rv
}

// The picker’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/delegate-8dlnb
func (p_ PHPickerViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The picker’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerViewController/delegate-8dlnb
func (p_ PHPickerViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// The opacity of the receiver. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (p_ PHPickerViewController) Opacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("opacity"))
	return rv
}


// SetOpacity sets the value of the opacity property.
// The opacity of the receiver. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (p_ PHPickerViewController) SetOpacity(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOpacity:"), value)
}


