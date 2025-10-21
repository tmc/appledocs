// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItem] class.
var (
	MenuItemClass     _MenuItemClass
	MenuItemClassOnce sync.Once
)

func getMenuItemClass() _MenuItemClass {
	MenuItemClassOnce.Do(func() {
		MenuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}
	})
	return MenuItemClass
}

type _MenuItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuItem] class.
type IMenuItem interface {
	objectivec.IObject
}

// A command item in an app menu.
//
// The class includes some private functionality needed to maintain binary compatibility with other components of Cocoa. Because of this fact, you can’t replace the class with a different class, but you can subclass it if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem
type MenuItem struct {
	objectivec.Object
}

// MenuItemFrom constructs a [MenuItem] from an unsafe.Pointer.
//
// A command item in an app menu.
func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemClass) New() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItem) Init() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItem) Autorelease() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItem creates a new MenuItem instance.
func NewMenuItem() MenuItem {
	return getMenuItemClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) Badge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) SetBadge(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}

// A Boolean value that indicates whether the menu item has a submenu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/hasSubmenu
func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSubmenu"))
	return rv
}

// The menu item’s action-method selector.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/action
func (m_ MenuItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The menu item’s action-method selector.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/action
func (m_ MenuItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}

// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowsautomatickeyequivalentlocalization
func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}


// SetAllowsAutomaticKeyEquivalentLocalization sets the value of the allowsAutomaticKeyEquivalentLocalization property.
// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowsautomatickeyequivalentlocalization
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}

// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowsautomatickeyequivalentmirroring
func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}


// SetAllowsAutomaticKeyEquivalentMirroring sets the value of the allowsAutomaticKeyEquivalentMirroring property.
// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowsautomatickeyequivalentmirroring
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowskeyequivalentwhenhidden
func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsKeyEquivalentWhenHidden"))
	return rv
}


// SetAllowsKeyEquivalentWhenHidden sets the value of the allowsKeyEquivalentWhenHidden property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/allowskeyequivalentwhenhidden
func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsKeyEquivalentWhenHidden:"), value)
}

// A custom string for a menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/attributedtitle
func (m_ MenuItem) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// A custom string for a menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/attributedtitle
func (m_ MenuItem) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedTitle:"), value)
}

// The menu item’s image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/image
func (m_ MenuItem) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The menu item’s image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/image
func (m_ MenuItem) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}

// The menu item indentation level for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/indentationlevel
func (m_ MenuItem) IndentationLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("indentationLevel"))
	return rv
}


// SetIndentationLevel sets the value of the indentationLevel property.
// The menu item indentation level for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/indentationlevel
func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndentationLevel:"), value)
}

// A Boolean value that marks the menu item as an alternate to the previous menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isalternate
func (m_ MenuItem) IsAlternate() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAlternate"))
	return rv
}


// SetIsAlternate sets the value of the isAlternate property.
// A Boolean value that marks the menu item as an alternate to the previous menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isalternate
func (m_ MenuItem) SetIsAlternate(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAlternate:"), value)
}

// A Boolean value that indicates whether the menu item is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isenabled
func (m_ MenuItem) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the menu item is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isenabled
func (m_ MenuItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value that indicates whether the menu item is hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishidden
func (m_ MenuItem) IsHidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that indicates whether the menu item is hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishidden
func (m_ MenuItem) SetIsHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHidden:"), value)
}

// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishiddenorhashiddenancestor
func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHiddenOrHasHiddenAncestor"))
	return rv
}


// SetIsHiddenOrHasHiddenAncestor sets the value of the isHiddenOrHasHiddenAncestor property.
// A Boolean value that indicates whether the menu item or any of its superitems is hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishiddenorhashiddenancestor
func (m_ MenuItem) SetIsHiddenOrHasHiddenAncestor(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHiddenOrHasHiddenAncestor:"), value)
}

// A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishighlighted
func (m_ MenuItem) IsHighlighted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHighlighted"))
	return rv
}


// SetIsHighlighted sets the value of the isHighlighted property.
// A Boolean value that indicates whether the menu item should be drawn highlighted.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishighlighted
func (m_ MenuItem) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHighlighted:"), value)
}

// A Boolean value indicating whether the menu item is a section header.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/issectionheader
func (m_ MenuItem) IsSectionHeader() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSectionHeader"))
	return rv
}


// SetIsSectionHeader sets the value of the isSectionHeader property.
// A Boolean value indicating whether the menu item is a section header.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/issectionheader
func (m_ MenuItem) SetIsSectionHeader(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSectionHeader:"), value)
}

// A Boolean value indicating whether the menu item is a separator item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isseparatoritem
func (m_ MenuItem) IsSeparatorItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSeparatorItem"))
	return rv
}


// SetIsSeparatorItem sets the value of the isSeparatorItem property.
// A Boolean value indicating whether the menu item is a separator item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isseparatoritem
func (m_ MenuItem) SetIsSeparatorItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSeparatorItem:"), value)
}

// The menu item’s unmodified key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/keyequivalent
func (m_ MenuItem) KeyEquivalent() string {
	rv := objc.Send[string](m_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// SetKeyEquivalent sets the value of the keyEquivalent property.
// The menu item’s unmodified key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/keyequivalent
func (m_ MenuItem) SetKeyEquivalent(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The menu item’s keyboard equivalent modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/keyequivalentmodifiermask
func (m_ MenuItem) KeyEquivalentModifierMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// SetKeyEquivalentModifierMask sets the value of the keyEquivalentModifierMask property.
// The menu item’s keyboard equivalent modifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/keyequivalentmodifiermask
func (m_ MenuItem) SetKeyEquivalentModifierMask(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The menu item’s menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/menu
func (m_ MenuItem) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The menu item’s menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/menu
func (m_ MenuItem) SetMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenu:"), value)
}

// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/mixedstateimage
func (m_ MenuItem) MixedStateImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mixedStateImage"))
	return rv
}


// SetMixedStateImage sets the value of the mixedStateImage property.
// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/mixedstateimage
func (m_ MenuItem) SetMixedStateImage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMixedStateImage:"), value)
}

// The image of the menu item that indicates an “off” state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/offstateimage
func (m_ MenuItem) OffStateImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("offStateImage"))
	return rv
}


// SetOffStateImage sets the value of the offStateImage property.
// The image of the menu item that indicates an “off” state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/offstateimage
func (m_ MenuItem) SetOffStateImage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffStateImage:"), value)
}

// The image of the menu item that indicates an “on” state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/onstateimage
func (m_ MenuItem) OnStateImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("onStateImage"))
	return rv
}


// SetOnStateImage sets the value of the onStateImage property.
// The image of the menu item that indicates an “on” state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/onstateimage
func (m_ MenuItem) SetOnStateImage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnStateImage:"), value)
}

// The menu item whose submenu contains the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// The menu item whose submenu contains the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) SetParent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParent:"), value)
}

// The object represented by the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/representedobject
func (m_ MenuItem) RepresentedObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("representedObject"))
	return rv
}


// SetRepresentedObject sets the value of the representedObject property.
// The object represented by the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/representedobject
func (m_ MenuItem) SetRepresentedObject(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRepresentedObject:"), value)
}

// The state of the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/state
func (m_ MenuItem) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The state of the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/state
func (m_ MenuItem) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}

// The submenu of the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/submenu
func (m_ MenuItem) Submenu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("submenu"))
	return rv
}


// SetSubmenu sets the value of the submenu property.
// The submenu of the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/submenu
func (m_ MenuItem) SetSubmenu(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubmenu:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/subtitle
func (m_ MenuItem) Subtitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/subtitle
func (m_ MenuItem) SetSubtitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The menu item’s tag.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tag
func (m_ MenuItem) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// The menu item’s tag.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tag
func (m_ MenuItem) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}

// The menu item’s target.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/target
func (m_ MenuItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The menu item’s target.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/target
func (m_ MenuItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTarget:"), value)
}

// The menu item’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/title
func (m_ MenuItem) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The menu item’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/title
func (m_ MenuItem) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A help tag for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tooltip
func (m_ MenuItem) ToolTip() string {
	rv := objc.Send[string](m_.ID, objc.Sel("toolTip"))
	return rv
}


// SetToolTip sets the value of the toolTip property.
// A help tag for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tooltip
func (m_ MenuItem) SetToolTip(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// The user-assigned key equivalent for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/userkeyequivalent
func (m_ MenuItem) UserKeyEquivalent() string {
	rv := objc.Send[string](m_.ID, objc.Sel("userKeyEquivalent"))
	return rv
}


// SetUserKeyEquivalent sets the value of the userKeyEquivalent property.
// The user-assigned key equivalent for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/userkeyequivalent
func (m_ MenuItem) SetUserKeyEquivalent(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserKeyEquivalent:"), objc.String(value))
}

// The content view for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/view
func (m_ MenuItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The content view for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/view
func (m_ MenuItem) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setView:"), value)
}



