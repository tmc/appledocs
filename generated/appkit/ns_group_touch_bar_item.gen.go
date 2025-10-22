// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	GroupTouchBar() NSTouchBar
	SetGroupTouchBar(value ITouchBar)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	EffectiveCompressionOptions() NSUserInterfaceCompressionOptions
	SetEffectiveCompressionOptions(value NSUserInterfaceCompressionOptions)
	GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
	PrioritizedCompressionOptions() NSUserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value NSUserInterfaceCompressionOptions)
}

// A bar item that provides a bar to contain other items.
//
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




// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func NewGroupTouchBarItemGroupItemWithIdentifierItems(identifier ITouchBarItemIdentifier, items []TouchBarItem) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}



// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(identifier ITouchBarItemIdentifier, items []TouchBarItem, allowedCompressionOptions NSUserInterfaceCompressionOptions) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItems(identifier ITouchBarItemIdentifier, items []TouchBarItem) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}

// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItemsAllowedCompressionOptions(identifier ITouchBarItemIdentifier, items []TouchBarItem, allowedCompressionOptions NSUserInterfaceCompressionOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}

// A bar that holds this group’s items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) GroupTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](g_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// SetGroupTouchBar sets the value of the groupTouchBar property.
// A bar that holds this group’s items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value ITouchBar) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupTouchBar:"), value)
}

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/customizationlabel
func (g_ GroupTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](g_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/customizationlabel
func (g_ GroupTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

// The compression options that are currently active on the group.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/effectivecompressionoptions
func (g_ GroupTouchBarItem) EffectiveCompressionOptions() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](g_.ID, objc.Sel("effectiveCompressionOptions"))
	return rv
}


// SetEffectiveCompressionOptions sets the value of the effectiveCompressionOptions property.
// The compression options that are currently active on the group.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/effectivecompressionoptions
func (g_ GroupTouchBarItem) SetEffectiveCompressionOptions(value NSUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEffectiveCompressionOptions:"), value)
}

// The user interface direction that controls the layout order of the items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/groupuserinterfacelayoutdirection
func (g_ GroupTouchBarItem) GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](g_.ID, objc.Sel("groupUserInterfaceLayoutDirection"))
	return rv
}


// SetGroupUserInterfaceLayoutDirection sets the value of the groupUserInterfaceLayoutDirection property.
// The user interface direction that controls the layout order of the items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/groupuserinterfacelayoutdirection
func (g_ GroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}

// The preferred width for items in the group.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/preferreditemwidth
func (g_ GroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("preferredItemWidth"))
	return rv
}


// SetPreferredItemWidth sets the value of the preferredItemWidth property.
// The preferred width for items in the group.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/preferreditemwidth
func (g_ GroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredItemWidth:"), value)
}

// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prefersequalwidths
func (g_ GroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("prefersEqualWidths"))
	return rv
}


// SetPrefersEqualWidths sets the value of the prefersEqualWidths property.
// A Boolean value that specifies that items should have equal widths when possible.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prefersequalwidths
func (g_ GroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrefersEqualWidths:"), value)
}

// The allowed compression options, in the order they should be applied.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prioritizedcompressionoptions
func (g_ GroupTouchBarItem) PrioritizedCompressionOptions() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](g_.ID, objc.Sel("prioritizedCompressionOptions"))
	return rv
}


// SetPrioritizedCompressionOptions sets the value of the prioritizedCompressionOptions property.
// The allowed compression options, in the order they should be applied.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/prioritizedcompressionoptions
func (g_ GroupTouchBarItem) SetPrioritizedCompressionOptions(value NSUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrioritizedCompressionOptions:"), value)
}


