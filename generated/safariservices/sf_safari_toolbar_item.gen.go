// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SFSafariToolbarItem] class.
var (
	SFSafariToolbarItemClass     _SFSafariToolbarItemClass
	SFSafariToolbarItemClassOnce sync.Once
)

func getSFSafariToolbarItemClass() _SFSafariToolbarItemClass {
	SFSafariToolbarItemClassOnce.Do(func() {
		SFSafariToolbarItemClass = _SFSafariToolbarItemClass{objc.GetClass("SFSafariToolbarItem")}
	})
	return SFSafariToolbarItemClass
}

type _SFSafariToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariToolbarItem] class.
type ISFSafariToolbarItem interface {
	objectivec.IObject
	SetBadgeText(badgeText string)
	SetEnabled(enabled bool)
	SetImage(image unsafe.Pointer)
	ShowPopover()
}

// A proxy for a Safari app extension toolbar item in a Safari window.
//
// Your app extension only uses this object when it wants to explicitly set the toolbar item state. Typically, other state changes occur automatically. Safari calls on your app extension handler when changes, such as navigation to a webpage, could affect the state of the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem
type SFSafariToolbarItem struct {
	objectivec.Object
}

// SFSafariToolbarItemFrom constructs a [SFSafariToolbarItem] from an unsafe.Pointer.
//
// A proxy for a Safari app extension toolbar item in a Safari window.
func SFSafariToolbarItemFrom(ptr unsafe.Pointer) SFSafariToolbarItem {
	return SFSafariToolbarItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariToolbarItemClass) Alloc() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariToolbarItemClass) New() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariToolbarItem) Init() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariToolbarItem) Autorelease() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariToolbarItem creates a new SFSafariToolbarItem instance.
func NewSFSafariToolbarItem() SFSafariToolbarItem {
	return getSFSafariToolbarItemClass().New()
}


// Sets the badge text for the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setBadgeText(_:)
func (s_ SFSafariToolbarItem) SetBadgeText(badgeText string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBadgeText:"), objc.String(badgeText))
}

// Sets whether the toolbar item is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setEnabled(_:)
func (s_ SFSafariToolbarItem) SetEnabled(enabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), enabled)
}

// Sets the image displayed in the toolbar button.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setImage(_:)
func (s_ SFSafariToolbarItem) SetImage(image unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), image)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/showPopover()
func (s_ SFSafariToolbarItem) ShowPopover() {
	objc.Send[objc.ID](s_.ID, objc.Sel("showPopover"))
}



