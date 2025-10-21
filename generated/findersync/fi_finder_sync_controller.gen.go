// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [FIFinderSyncController] class.
type IFIFinderSyncController interface {
	foundation.IExtensionContext
	LastUsedDateForItemWithURL(itemURL unsafe.Pointer) unsafe.Pointer
	SelectedItemURLs() []appkit.NSURL
	SetBadgeIdentifierForURL(badgeID string, url unsafe.Pointer)
	SetBadgeImageLabelForBadgeIdentifier(image unsafe.Pointer, label string, badgeID string)
	SetLastUsedDateForItemWithURLCompletion(lastUsedDate unsafe.Pointer, itemURL unsafe.Pointer, completion unsafe.Pointer)
	SetTagDataForItemWithURLCompletion(tagData unsafe.Pointer, itemURL unsafe.Pointer, completion unsafe.Pointer)
	TagDataForItemWithURL(itemURL unsafe.Pointer) unsafe.Pointer
	TargetedURL() unsafe.Pointer
}

// A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
//
// Use the Finder Sync controller to configure your extension, to set badges on items in the Finder’s window, and to get a list of selected and targeted items.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FIFinderSyncControllerClass) Alloc() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the shared Finder Sync controller object.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/default()
func (fc _FIFinderSyncControllerClass) DefaultController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("defaultController"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/showExtensionManagementInterface()
func (fc _FIFinderSyncControllerClass) ShowExtensionManagementInterface() {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("showExtensionManagementInterface"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/isExtensionEnabled
func (fc _FIFinderSyncControllerClass) ExtensionEnabled() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("extensionEnabled"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/lastUsedDateForItem(with:)
func (f_ FIFinderSyncController) LastUsedDateForItemWithURL(itemURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("lastUsedDateForItemWithURL:"), itemURL)
	return rv
}

// Returns an array of selected items.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/selectedItemURLs()
func (f_ FIFinderSyncController) SelectedItemURLs() []appkit.NSURL {
	rv := objc.Send[[]appkit.NSURL](f_.ID, objc.Sel("selectedItemURLs"))
	return rv
}

// Sets the badge for a file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeIdentifier(_:for:)
func (f_ FIFinderSyncController) SetBadgeIdentifierForURL(badgeID string, url unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBadgeIdentifier:forURL:"), objc.String(badgeID), url)
}

// Sets the badge image and label for the given ID.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeImage(_:label:forBadgeIdentifier:)
func (f_ FIFinderSyncController) SetBadgeImageLabelForBadgeIdentifier(image unsafe.Pointer, label string, badgeID string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBadgeImage:label:forBadgeIdentifier:"), image, objc.String(label), objc.String(badgeID))
}

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setLastUsedDate(_:forItemWith:completion:)
func (f_ FIFinderSyncController) SetLastUsedDateForItemWithURLCompletion(lastUsedDate unsafe.Pointer, itemURL unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLastUsedDate:forItemWithURL:completion:"), lastUsedDate, itemURL, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setTagData(_:forItemWith:completion:)
func (f_ FIFinderSyncController) SetTagDataForItemWithURLCompletion(tagData unsafe.Pointer, itemURL unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTagData:forItemWithURL:completion:"), tagData, itemURL, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/tagDataForItem(with:)
func (f_ FIFinderSyncController) TagDataForItemWithURL(itemURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("tagDataForItemWithURL:"), itemURL)
	return rv
}

// Returns the URL of the Finder’s current target.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/targetedURL()
func (f_ FIFinderSyncController) TargetedURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("targetedURL"))
	return rv
}

// The directories managed by this extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/directoryURLs
func (f_ FIFinderSyncController) DirectoryURLs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("directoryURLs"))
	return rv
}


// SetDirectoryURLs sets the value of the directoryURLs property.
// The directories managed by this extension.

//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/directoryURLs
func (f_ FIFinderSyncController) SetDirectoryURLs(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDirectoryURLs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/isExtensionEnabled
func (f_ FIFinderSyncController) ExtensionEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("extensionEnabled"))
	return rv
}



