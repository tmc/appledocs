// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	// methods:
	ShowPopover()
}

// A proxy for a Safari app extension toolbar item in a Safari window.
//
// Your app extension only uses this object when it wants to explicitly set the toolbar item state. Typically, other state changes occur automatically. Safari calls on your app extension handler when changes, such as navigation to a webpage, could affect the state of the toolbar item.


// A proxy for a Safari app extension toolbar item in a Safari window.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/showPopover()
func (s_ SFSafariToolbarItem) ShowPopover() {
	objc.Send[objc.ID](s_.ID, objc.Sel("showPopover"))
}



