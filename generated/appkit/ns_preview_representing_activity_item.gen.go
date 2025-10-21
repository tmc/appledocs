// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewRepresentingActivityItem] class.
var (
	PreviewRepresentingActivityItemClass     _PreviewRepresentingActivityItemClass
	PreviewRepresentingActivityItemClassOnce sync.Once
)

func getPreviewRepresentingActivityItemClass() _PreviewRepresentingActivityItemClass {
	PreviewRepresentingActivityItemClassOnce.Do(func() {
		PreviewRepresentingActivityItemClass = _PreviewRepresentingActivityItemClass{objc.GetClass("NSPreviewRepresentingActivityItem")}
	})
	return PreviewRepresentingActivityItemClass
}

type _PreviewRepresentingActivityItemClass struct {
	class objc.Class
}

// An interface definition for the [PreviewRepresentingActivityItem] class.
type IPreviewRepresentingActivityItem interface {
	objectivec.IObject
}

// A type that adds metadata to an item you share using the macOS share sheet.
//
// An object provides a concrete implementation of the protocol. Use it to create shareable items for common types such as strings or images, or when you don’t want to adopt the protocol directly in your app’s objects. To share the item from your app, initialize the object with this object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem
type PreviewRepresentingActivityItem struct {
	objectivec.Object
}

// PreviewRepresentingActivityItemFrom constructs a [PreviewRepresentingActivityItem] from an unsafe.Pointer.
//
// A type that adds metadata to an item you share using the macOS share sheet.
func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewRepresentingActivityItemClass) Alloc() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewRepresentingActivityItemClass) New() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewRepresentingActivityItem) Init() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewRepresentingActivityItem) Autorelease() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewRepresentingActivityItem creates a new PreviewRepresentingActivityItem instance.
func NewPreviewRepresentingActivityItem() PreviewRepresentingActivityItem {
	return getPreviewRepresentingActivityItemClass().New()
}




// Creates a metadata object with the title, image, and icon for a shareable item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:image:icon:)
func NewPreviewRepresentingActivityItemWithItemTitleImageIcon(item objectivec.IObject, title string, image IImage, icon IImage) PreviewRepresentingActivityItem {
	instance := getPreviewRepresentingActivityItemClass().Alloc()
	rv := objc.Send[PreviewRepresentingActivityItem](instance.ID, objc.Sel("initWithItem:title:image:icon:"), item, objc.String(title), image, icon)
	rv.Autorelease()
	return rv
}



// Creates a metadata object that provides a title and images for a shareable item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:imageProvider:iconProvider:)
func NewPreviewRepresentingActivityItemWithItemTitleImageProviderIconProvider(item objectivec.IObject, title string, imageProvider foundation.IItemProvider, iconProvider foundation.IItemProvider) PreviewRepresentingActivityItem {
	instance := getPreviewRepresentingActivityItemClass().Alloc()
	rv := objc.Send[PreviewRepresentingActivityItem](instance.ID, objc.Sel("initWithItem:title:imageProvider:iconProvider:"), item, objc.String(title), imageProvider, iconProvider)
	rv.Autorelease()
	return rv
}



