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
	

	// properties:
	AllowsAutomaticKeyEquivalentLocalization() bool
	SetAllowsAutomaticKeyEquivalentLocalization(value bool)
	AllowsAutomaticKeyEquivalentMirroring() bool
	SetAllowsAutomaticKeyEquivalentMirroring(value bool)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)
	KeyEquivalent() foundation.foundation.INSString
	SetKeyEquivalent(value foundation.foundation.INSString)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	UserKeyEquivalent() foundation.foundation.INSString
	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	AttributedTitle() foundation.foundation.INSAttributedString
	SetAttributedTitle(value foundation.foundation.INSAttributedString)
	Badge() IMenuItemBadge
	SetBadge(value IMenuItemBadge)
	HasSubmenu() bool
	SetHasSubmenu(value bool)
	Image() IImage
	SetImage(value IImage)
	IndentationLevel() int
	SetIndentationLevel(value int)
	IsAlternate() bool
	SetIsAlternate(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	SetIsHiddenOrHasHiddenAncestor(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsSectionHeader() bool
	SetIsSectionHeader(value bool)
	IsSeparatorItem() bool
	SetIsSeparatorItem(value bool)
	Menu() IMenu
	SetMenu(value IMenu)
	MixedStateImage() IImage
	SetMixedStateImage(value IImage)
	OffStateImage() IImage
	SetOffStateImage(value IImage)
	OnStateImage() IImage
	SetOnStateImage(value IImage)
	Parent() IMenuItem
	SetParent(value IMenuItem)
	State() objectivec.IObject
	SetState(value objectivec.IObject)
	Submenu() IMenu
	SetSubmenu(value IMenu)
	Subtitle() foundation.foundation.INSString
	SetSubtitle(value foundation.foundation.INSString)
	Tag() int
	SetTag(value int)
	Title() foundation.foundation.INSString
	SetTitle(value foundation.foundation.INSString)
	ToolTip() foundation.foundation.INSString
	SetToolTip(value foundation.foundation.INSString)
	View() IView
	SetView(value IView)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A command item in an app menu.
//
// The class includes some private functionality needed to maintain binary compatibility with other components of Cocoa. Because of this fact, you can’t replace the class with a different class, but you can subclass it if necessary.


// A command item in an app menu.
//
// [Full Topic]
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















// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("usesUserKeyEquivalents"))
	return rv
}











// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentLocalization
func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}


// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentLocalization
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}


// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentMirroring
func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}


// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentMirroring
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsKeyEquivalentWhenHidden
func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsKeyEquivalentWhenHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsKeyEquivalentWhenHidden
func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsKeyEquivalentWhenHidden:"), value)
}


// The menu item’s unmodified key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalent
func (m_ MenuItem) KeyEquivalent() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// The menu item’s unmodified key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalent
func (m_ MenuItem) SetKeyEquivalent(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalent:"), value)
}


// The menu item’s keyboard equivalent modifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalentModifierMask
func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](m_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// The menu item’s keyboard equivalent modifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalentModifierMask
func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}


// The user-assigned key equivalent for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/userKeyEquivalent
func (m_ MenuItem) UserKeyEquivalent() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userKeyEquivalent"))
	return rv
}


// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (m_ MenuItem) UsesUserKeyEquivalents() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesUserKeyEquivalents"))
	return rv
}


// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (m_ MenuItem) SetUsesUserKeyEquivalents(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesUserKeyEquivalents:"), value)
}


// The menu item’s action-method selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/action
func (m_ MenuItem) Action() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("action"))
	return rv
}


// The menu item’s action-method selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/action
func (m_ MenuItem) SetAction(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}


// A custom string for a menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/attributedtitle
func (m_ MenuItem) AttributedTitle() foundation.foundation.INSAttributedString {
	rv := objc.Send[foundation.NSAttributedString](m_.ID, objc.Sel("attributedTitle"))
	return rv
}


// A custom string for a menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/attributedtitle
func (m_ MenuItem) SetAttributedTitle(value foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedTitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge
func (m_ MenuItem) Badge() IMenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID, objc.Sel("badge"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge
func (m_ MenuItem) SetBadge(value IMenuItemBadge) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}


// A Boolean value that indicates whether the menu item has a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/hassubmenu
func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSubmenu"))
	return rv
}


// A Boolean value that indicates whether the menu item has a submenu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/hassubmenu
func (m_ MenuItem) SetHasSubmenu(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasSubmenu:"), value)
}


// The menu item’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/image
func (m_ MenuItem) Image() IImage {
	rv := objc.Send[Image](m_.ID, objc.Sel("image"))
	return rv
}


// The menu item’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/image
func (m_ MenuItem) SetImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}


// The menu item indentation level for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/indentationlevel
func (m_ MenuItem) IndentationLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("indentationLevel"))
	return rv
}


// The menu item indentation level for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/indentationlevel
func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndentationLevel:"), value)
}


// A Boolean value that marks the menu item as an alternate to the previous menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isalternate
func (m_ MenuItem) IsAlternate() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAlternate"))
	return rv
}


// A Boolean value that marks the menu item as an alternate to the previous menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isalternate
func (m_ MenuItem) SetIsAlternate(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAlternate:"), value)
}


// A Boolean value that indicates whether the menu item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isenabled
func (m_ MenuItem) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the menu item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isenabled
func (m_ MenuItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the menu item is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishidden
func (m_ MenuItem) IsHidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value that indicates whether the menu item is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishidden
func (m_ MenuItem) SetIsHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHidden:"), value)
}


// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishiddenorhashiddenancestor
func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHiddenOrHasHiddenAncestor"))
	return rv
}


// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishiddenorhashiddenancestor
func (m_ MenuItem) SetIsHiddenOrHasHiddenAncestor(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHiddenOrHasHiddenAncestor:"), value)
}


// A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishighlighted
func (m_ MenuItem) IsHighlighted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/ishighlighted
func (m_ MenuItem) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHighlighted:"), value)
}


// A Boolean value indicating whether the menu item is a section header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/issectionheader
func (m_ MenuItem) IsSectionHeader() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSectionHeader"))
	return rv
}


// A Boolean value indicating whether the menu item is a section header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/issectionheader
func (m_ MenuItem) SetIsSectionHeader(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSectionHeader:"), value)
}


// A Boolean value indicating whether the menu item is a separator item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isseparatoritem
func (m_ MenuItem) IsSeparatorItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSeparatorItem"))
	return rv
}


// A Boolean value indicating whether the menu item is a separator item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/isseparatoritem
func (m_ MenuItem) SetIsSeparatorItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSeparatorItem:"), value)
}


// The menu item’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/menu
func (m_ MenuItem) Menu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("menu"))
	return rv
}


// The menu item’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/menu
func (m_ MenuItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenu:"), value)
}


// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/mixedstateimage
func (m_ MenuItem) MixedStateImage() IImage {
	rv := objc.Send[Image](m_.ID, objc.Sel("mixedStateImage"))
	return rv
}


// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/mixedstateimage
func (m_ MenuItem) SetMixedStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMixedStateImage:"), value)
}


// The image of the menu item that indicates an “off” state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/offstateimage
func (m_ MenuItem) OffStateImage() IImage {
	rv := objc.Send[Image](m_.ID, objc.Sel("offStateImage"))
	return rv
}


// The image of the menu item that indicates an “off” state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/offstateimage
func (m_ MenuItem) SetOffStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffStateImage:"), value)
}


// The image of the menu item that indicates an “on” state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/onstateimage
func (m_ MenuItem) OnStateImage() IImage {
	rv := objc.Send[Image](m_.ID, objc.Sel("onStateImage"))
	return rv
}


// The image of the menu item that indicates an “on” state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/onstateimage
func (m_ MenuItem) SetOnStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnStateImage:"), value)
}


// The menu item whose submenu contains the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) Parent() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("parent"))
	return rv
}


// The menu item whose submenu contains the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) SetParent(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParent:"), value)
}


// The state of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/state
func (m_ MenuItem) State() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("state"))
	return rv
}


// The state of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/state
func (m_ MenuItem) SetState(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


// The submenu of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/submenu
func (m_ MenuItem) Submenu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("submenu"))
	return rv
}


// The submenu of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/submenu
func (m_ MenuItem) SetSubmenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubmenu:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/subtitle
func (m_ MenuItem) Subtitle() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/subtitle
func (m_ MenuItem) SetSubtitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}


// The menu item’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tag
func (m_ MenuItem) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}


// The menu item’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tag
func (m_ MenuItem) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}


// The menu item’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/title
func (m_ MenuItem) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}


// The menu item’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/title
func (m_ MenuItem) SetTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}


// A help tag for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tooltip
func (m_ MenuItem) ToolTip() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("toolTip"))
	return rv
}


// A help tag for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/tooltip
func (m_ MenuItem) SetToolTip(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToolTip:"), value)
}


// The content view for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/view
func (m_ MenuItem) View() IView {
	rv := objc.Send[View](m_.ID, objc.Sel("view"))
	return rv
}


// The content view for the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/view
func (m_ MenuItem) SetView(value IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setView:"), value)
}








