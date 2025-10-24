// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class FIFinderSyncController */


/* debug [class_header]: Header for FIFinderSyncController */
// The class instance for the [FIFinderSyncController] class.
var (
	FIFinderSyncControllerClass     _FIFinderSyncControllerClass
	FIFinderSyncControllerClassOnce sync.Once
)

func getFIFinderSyncControllerClass() _FIFinderSyncControllerClass {
	FIFinderSyncControllerClassOnce.Do(func() {
		FIFinderSyncControllerClass = _FIFinderSyncControllerClass{objc.GetClass("FIFinderSyncController")}
	})
	return FIFinderSyncControllerClass
}

type _FIFinderSyncControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FIFinderSyncController */
// An interface definition for the [FIFinderSyncController] class.
type IFIFinderSyncController interface {
	foundation.IExtensionContext
	
/* debug [class_interface_properties]: Properties for FIFinderSyncController */
	// properties:
	DirectoryURLs() unsafe.Pointer
	SetDirectoryURLs(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FIFinderSyncController */
	// methods:
	LastUsedDateForItemWithURL(itemURL objc.IObject /* cross-framework: NSURL */) foundation.Date
	SelectedItemURLs() []foundation.URL
	SetBadgeIdentifierForURL(badgeID objc.IObject /* cross-framework: NSString */, url objc.IObject /* cross-framework: NSURL */)
	SetBadgeImageLabelForBadgeIdentifier(image appkit.Image, label objc.IObject /* cross-framework: NSString */, badgeID objc.IObject /* cross-framework: NSString */)
	SetLastUsedDateForItemWithURLCompletion(lastUsedDate objc.IObject /* cross-framework: NSDate */, itemURL objc.IObject /* cross-framework: NSURL */, completion unsafe.Pointer)
	SetTagDataForItemWithURLCompletion(tagData objc.IObject /* cross-framework: NSData */, itemURL objc.IObject /* cross-framework: NSURL */, completion unsafe.Pointer)
	TagDataForItemWithURL(itemURL objc.IObject /* cross-framework: NSURL */) foundation.Data
	TargetedURL() foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FIFinderSyncController */
// Alloc allocates a new instance without initialization.
func (fc _FIFinderSyncControllerClass) Alloc() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FIFinderSyncControllerClass) New() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FIFinderSyncController) Init() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FIFinderSyncController) Autorelease() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFIFinderSyncController creates a new FIFinderSyncController instance.
func NewFIFinderSyncController() FIFinderSyncController {
	return getFIFinderSyncControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FIFinderSyncController */
// A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
//
// Use the Finder Sync controller to configure your extension, to set badges on items in the Finder’s window, and to get a list of selected and targeted items.


// A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController
type FIFinderSyncController struct {
	foundation.ExtensionContext
}

// FIFinderSyncControllerFrom constructs a [FIFinderSyncController] from an unsafe.Pointer.
//
// A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
func FIFinderSyncControllerFrom(ptr unsafe.Pointer) FIFinderSyncController {
	return FIFinderSyncController{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FIFinderSyncController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FIFinderSyncController */

// Returns the shared Finder Sync controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/default()
func (fc _FIFinderSyncControllerClass) DefaultController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("defaultController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultController) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/showExtensionManagementInterface()
func (fc _FIFinderSyncControllerClass) ShowExtensionManagementInterface() {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("showExtensionManagementInterface"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShowExtensionManagementInterface) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FIFinderSyncController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/isExtensionEnabled
func (fc _FIFinderSyncControllerClass) ExtensionEnabled() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("extensionEnabled"))
	return rv
}/* debug [class_properties_class/property]: extensionEnabled */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FIFinderSyncController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/lastUsedDateForItem(with:)
func (f_ FIFinderSyncController) LastUsedDateForItemWithURL(itemURL objc.IObject /* cross-framework: NSURL */) foundation.Date {
	rv := objc.Send[foundation.Date](f_.ID, objc.Sel("lastUsedDateForItemWithURL:"), itemURL)
	return rv
}/* debug [instance_methods/method]: LastUsedDateForItemWithURL */


// Returns an array of selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/selectedItemURLs()
func (f_ FIFinderSyncController) SelectedItemURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](f_.ID, objc.Sel("selectedItemURLs"))
	return rv
}/* debug [instance_methods/method]: SelectedItemURLs */


// Sets the badge for a file or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeIdentifier(_:for:)
func (f_ FIFinderSyncController) SetBadgeIdentifierForURL(badgeID objc.IObject /* cross-framework: NSString */, url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBadgeIdentifier:forURL:"), badgeID, url)
}/* debug [instance_methods/method]: SetBadgeIdentifierForURL */


// Sets the badge image and label for the given ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeImage(_:label:forBadgeIdentifier:)
func (f_ FIFinderSyncController) SetBadgeImageLabelForBadgeIdentifier(image appkit.Image, label objc.IObject /* cross-framework: NSString */, badgeID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBadgeImage:label:forBadgeIdentifier:"), image, label, badgeID)
}/* debug [instance_methods/method]: SetBadgeImageLabelForBadgeIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setLastUsedDate(_:forItemWith:completion:)
func (f_ FIFinderSyncController) SetLastUsedDateForItemWithURLCompletion(lastUsedDate objc.IObject /* cross-framework: NSDate */, itemURL objc.IObject /* cross-framework: NSURL */, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLastUsedDate:forItemWithURL:completion:"), lastUsedDate, itemURL, completion)
}/* debug [instance_methods/method]: SetLastUsedDateForItemWithURLCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setTagData(_:forItemWith:completion:)
func (f_ FIFinderSyncController) SetTagDataForItemWithURLCompletion(tagData objc.IObject /* cross-framework: NSData */, itemURL objc.IObject /* cross-framework: NSURL */, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTagData:forItemWithURL:completion:"), tagData, itemURL, completion)
}/* debug [instance_methods/method]: SetTagDataForItemWithURLCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/tagDataForItem(with:)
func (f_ FIFinderSyncController) TagDataForItemWithURL(itemURL objc.IObject /* cross-framework: NSURL */) foundation.Data {
	rv := objc.Send[foundation.Data](f_.ID, objc.Sel("tagDataForItemWithURL:"), itemURL)
	return rv
}/* debug [instance_methods/method]: TagDataForItemWithURL */


// Returns the URL of the Finder’s current target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/targetedURL()
func (f_ FIFinderSyncController) TargetedURL() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("targetedURL"))
	return rv
}/* debug [instance_methods/method]: TargetedURL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FIFinderSyncController */

// The directories managed by this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/directoryURLs
func (f_ FIFinderSyncController) DirectoryURLs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("directoryURLs"))
	return rv
}/* debug [instance_properties/getter]: directoryURLs */


// The directories managed by this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/directoryURLs
func (f_ FIFinderSyncController) SetDirectoryURLs(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDirectoryURLs:"), value)
}/* debug [instance_properties/setter]: directoryURLs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/isExtensionEnabled
func (f_ FIFinderSyncController) ExtensionEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("extensionEnabled"))
	return rv
}/* debug [instance_properties/getter]: extensionEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FIFinderSyncController */





