// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GroupTouchBarItem] class.
var (
	GroupTouchBarItemClass     _GroupTouchBarItemClass
	GroupTouchBarItemClassOnce sync.Once
)

func getGroupTouchBarItemClass() _GroupTouchBarItemClass {
	GroupTouchBarItemClassOnce.Do(func() {
		GroupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}
	})
	return GroupTouchBarItemClass
}

type _GroupTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
	// properties:
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	EffectiveCompressionOptions() IUserInterfaceCompressionOptions
	SetEffectiveCompressionOptions(value IUserInterfaceCompressionOptions)
	GroupTouchBar() ITouchBar
	SetGroupTouchBar(value ITouchBar)
	GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection /* not a class type */
	SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection /* not a class type */)
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
	PrioritizedCompressionOptions() IUserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value IUserInterfaceCompressionOptions)
	// methods:
}

// A bar item that provides a bar to contain other items.


// A bar item that provides a bar to contain other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem
type GroupTouchBarItem struct {
	TouchBarItem
}

// GroupTouchBarItemFrom constructs a [GroupTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a bar to contain other items.
func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GroupTouchBarItem) Autorelease() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGroupTouchBarItem creates a new GroupTouchBarItem instance.
func NewGroupTouchBarItem() GroupTouchBarItem {
	return getGroupTouchBarItemClass().New()
}



// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/customizationlabel
func (g_ GroupTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/customizationlabel
func (g_ GroupTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The compression options that are currently active on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/effectivecompressionoptions
func (g_ GroupTouchBarItem) EffectiveCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](g_.ID, objc.Sel("effectiveCompressionOptions"))
	return rv
}


// The compression options that are currently active on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/effectivecompressionoptions
func (g_ GroupTouchBarItem) SetEffectiveCompressionOptions(value IUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEffectiveCompressionOptions:"), value)
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (g_ GroupTouchBarItem) GroupTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](g_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value ITouchBar) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupTouchBar:"), value)
}


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/groupuserinterfacelayoutdirection
func (g_ GroupTouchBarItem) GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection /* not a class type */ {
	rv := objc.Send[UserInterfaceLayoutDirection](g_.ID, objc.Sel("groupUserInterfaceLayoutDirection"))
	return rv
}


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/groupuserinterfacelayoutdirection
func (g_ GroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/preferreditemwidth
func (g_ GroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("preferredItemWidth"))
	return rv
}


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/preferreditemwidth
func (g_ GroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredItemWidth:"), value)
}


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prefersequalwidths
func (g_ GroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("prefersEqualWidths"))
	return rv
}


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prefersequalwidths
func (g_ GroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrefersEqualWidths:"), value)
}


// The allowed compression options, in the order they should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prioritizedcompressionoptions
func (g_ GroupTouchBarItem) PrioritizedCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](g_.ID, objc.Sel("prioritizedCompressionOptions"))
	return rv
}


// The allowed compression options, in the order they should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prioritizedcompressionoptions
func (g_ GroupTouchBarItem) SetPrioritizedCompressionOptions(value IUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrioritizedCompressionOptions:"), value)
}



