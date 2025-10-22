// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Mnemonic() foundation.String
	MnemonicLocation() uint
	SetMnemonicLocation(location uint)
	SetTitleWithMnemonic(stringWithAmpersand string)
	Action() objc.SEL
	SetAction(value objc.SEL)
	AllowsAutomaticKeyEquivalentLocalization() bool
	SetAllowsAutomaticKeyEquivalentLocalization(value bool)
	AllowsAutomaticKeyEquivalentMirroring() bool
	SetAllowsAutomaticKeyEquivalentMirroring(value bool)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	Badge() NSMenuItemBadge
	SetBadge(value IMenuItemBadge)
	HasSubmenu() bool
	Image() Image
	SetImage(value IImage)
	IndentationLevel() int
	SetIndentationLevel(value int)
	Alternate() bool
	SetAlternate(value bool)
	Enabled() bool
	SetEnabled(value bool)
	Hidden() bool
	SetHidden(value bool)
	HiddenOrHasHiddenAncestor() bool
	Highlighted() bool
	SectionHeader() bool
	SeparatorItem() bool
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	Menu() NSMenu
	SetMenu(value IMenu)
	MixedStateImage() Image
	SetMixedStateImage(value IImage)
	OffStateImage() Image
	SetOffStateImage(value IImage)
	OnStateImage() Image
	SetOnStateImage(value IImage)
	ParentItem() NSMenuItem
	RepresentedObject() objc.ID
	SetRepresentedObject(value objc.ID)
	State() ControlStateValue
	SetState(value IControlStateValue)
	Submenu() NSMenu
	SetSubmenu(value IMenu)
	Subtitle() string
	SetSubtitle(value string)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() string
	SetTitle(value string)
	ToolTip() string
	SetToolTip(value string)
	UserKeyEquivalent() string
	View() NSView
	SetView(value IView)
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
	Parent() NSMenuItem
	SetParent(value IMenuItem)
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(coder:)
func NewMenuItemWithCoder(coder foundation.ICoder) MenuItem {
	instance := getMenuItemClass().Alloc()
	rv := objc.Send[MenuItem](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns an initialized instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(title:action:keyEquivalent:)
func NewMenuItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) MenuItem {
	instance := getMenuItemClass().Alloc()
	rv := objc.Send[MenuItem](instance.ID, objc.Sel("initWithTitle:action:keyEquivalent:"), objc.String(string_), selector, objc.String(charCode))
	rv.Autorelease()
	return rv
}


// Returns a menu item representing a section header for a logical grouping of menu commands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/sectionHeaderWithTitle:
func (mc _MenuItemClass) SectionHeaderWithTitle(title string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("sectionHeaderWithTitle:"), objc.String(title))
	return rv
}

// Returns a menu item that is used to separate logical groups of menu commands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/separator()
func (mc _MenuItemClass) SeparatorItem() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("separatorItem"))
	return rv
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("usesUserKeyEquivalents"))
	return rv
}
// An array of standard menu items related to Writing Tools. Each call to this method returns an array of newly allocated instances of NSMenuItem.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/writingToolsItems
func (mc _MenuItemClass) WritingToolsItems() []MenuItem {
	rv := objc.Send[[]MenuItem](objc.ID(mc.class), objc.Sel("writingToolsItems"))
	return rv
}
// Returns the character in the menu item title that appears underlined for use as a mnemonic.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/mnemonic
func (m_ MenuItem) Mnemonic() foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("mnemonic"))
	return rv
}

// Returns the position of the underlined character in the menu item title used as a mnemonic.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/mnemonicLocation
func (m_ MenuItem) MnemonicLocation() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("mnemonicLocation"))
	return rv
}

// Sets the character of the menu item title at location that is to be underlined.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/setMnemonicLocation:
func (m_ MenuItem) SetMnemonicLocation(location uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMnemonicLocation:"), location)
}

// Sets the title of a menu item with a character denoting an access key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/setTitleWithMnemonic(_:)
func (m_ MenuItem) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitleWithMnemonic:"), objc.String(stringWithAmpersand))
}

// The menu item’s action-method selector.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/action
func (m_ MenuItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](m_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The menu item’s action-method selector.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/action
func (m_ MenuItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}

// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentLocalization
func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}


// SetAllowsAutomaticKeyEquivalentLocalization sets the value of the allowsAutomaticKeyEquivalentLocalization property.
// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentLocalization
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}

// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentMirroring
func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}


// SetAllowsAutomaticKeyEquivalentMirroring sets the value of the allowsAutomaticKeyEquivalentMirroring property.
// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentMirroring
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsKeyEquivalentWhenHidden
func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsKeyEquivalentWhenHidden"))
	return rv
}


// SetAllowsKeyEquivalentWhenHidden sets the value of the allowsKeyEquivalentWhenHidden property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsKeyEquivalentWhenHidden
func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsKeyEquivalentWhenHidden:"), value)
}

// A custom string for a menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/attributedTitle
func (m_ MenuItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](m_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// A custom string for a menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/attributedTitle
func (m_ MenuItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedTitle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) Badge() NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) SetBadge(value IMenuItemBadge) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}

// A Boolean value that indicates whether the menu item has a submenu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/hasSubmenu
func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSubmenu"))
	return rv
}

// The menu item’s image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/image
func (m_ MenuItem) Image() Image {
	rv := objc.Send[Image](m_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The menu item’s image.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/image
func (m_ MenuItem) SetImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}

// The menu item indentation level for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/indentationLevel
func (m_ MenuItem) IndentationLevel() int {
	rv := objc.Send[int](m_.ID, objc.Sel("indentationLevel"))
	return rv
}


// SetIndentationLevel sets the value of the indentationLevel property.
// The menu item indentation level for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/indentationLevel
func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndentationLevel:"), value)
}

// A Boolean value that marks the menu item as an alternate to the previous menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isAlternate
func (m_ MenuItem) Alternate() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("alternate"))
	return rv
}


// SetAlternate sets the value of the alternate property.
// A Boolean value that marks the menu item as an alternate to the previous menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isAlternate
func (m_ MenuItem) SetAlternate(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlternate:"), value)
}

// A Boolean value that indicates whether the menu item is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isEnabled
func (m_ MenuItem) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether the menu item is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isEnabled
func (m_ MenuItem) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}

// A Boolean value that indicates whether the menu item is hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHidden
func (m_ MenuItem) Hidden() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hidden"))
	return rv
}


// SetHidden sets the value of the hidden property.
// A Boolean value that indicates whether the menu item is hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHidden
func (m_ MenuItem) SetHidden(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHidden:"), value)
}

// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHiddenOrHasHiddenAncestor
func (m_ MenuItem) HiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hiddenOrHasHiddenAncestor"))
	return rv
}

// A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHighlighted
func (m_ MenuItem) Highlighted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("highlighted"))
	return rv
}

// A Boolean value indicating whether the menu item is a section header.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isSectionHeader
func (m_ MenuItem) SectionHeader() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sectionHeader"))
	return rv
}

// A Boolean value indicating whether the menu item is a separator item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/isSeparatorItem
func (m_ MenuItem) SeparatorItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("separatorItem"))
	return rv
}

// The menu item’s unmodified key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalent
func (m_ MenuItem) KeyEquivalent() string {
	rv := objc.Send[string](m_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// SetKeyEquivalent sets the value of the keyEquivalent property.
// The menu item’s unmodified key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalent
func (m_ MenuItem) SetKeyEquivalent(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The menu item’s keyboard equivalent modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalentModifierMask
func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](m_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// SetKeyEquivalentModifierMask sets the value of the keyEquivalentModifierMask property.
// The menu item’s keyboard equivalent modifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalentModifierMask
func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The menu item’s menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/menu
func (m_ MenuItem) Menu() NSMenu {
	rv := objc.Send[NSMenu](m_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The menu item’s menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/menu
func (m_ MenuItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenu:"), value)
}

// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/mixedStateImage
func (m_ MenuItem) MixedStateImage() Image {
	rv := objc.Send[Image](m_.ID, objc.Sel("mixedStateImage"))
	return rv
}


// SetMixedStateImage sets the value of the mixedStateImage property.
// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/mixedStateImage
func (m_ MenuItem) SetMixedStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMixedStateImage:"), value)
}

// The image of the menu item that indicates an “off” state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/offStateImage
func (m_ MenuItem) OffStateImage() Image {
	rv := objc.Send[Image](m_.ID, objc.Sel("offStateImage"))
	return rv
}


// SetOffStateImage sets the value of the offStateImage property.
// The image of the menu item that indicates an “off” state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/offStateImage
func (m_ MenuItem) SetOffStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffStateImage:"), value)
}

// The image of the menu item that indicates an “on” state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/onStateImage
func (m_ MenuItem) OnStateImage() Image {
	rv := objc.Send[Image](m_.ID, objc.Sel("onStateImage"))
	return rv
}


// SetOnStateImage sets the value of the onStateImage property.
// The image of the menu item that indicates an “on” state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/onStateImage
func (m_ MenuItem) SetOnStateImage(value IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnStateImage:"), value)
}

// The menu item whose submenu contains the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/parent
func (m_ MenuItem) ParentItem() NSMenuItem {
	rv := objc.Send[NSMenuItem](m_.ID, objc.Sel("parentItem"))
	return rv
}

// The object represented by the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/representedObject
func (m_ MenuItem) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("representedObject"))
	return rv
}


// SetRepresentedObject sets the value of the representedObject property.
// The object represented by the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/representedObject
func (m_ MenuItem) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRepresentedObject:"), value)
}

// The state of the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/state
func (m_ MenuItem) State() ControlStateValue {
	rv := objc.Send[ControlStateValue](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The state of the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/state
func (m_ MenuItem) SetState(value IControlStateValue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}

// The submenu of the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/submenu
func (m_ MenuItem) Submenu() NSMenu {
	rv := objc.Send[NSMenu](m_.ID, objc.Sel("submenu"))
	return rv
}


// SetSubmenu sets the value of the submenu property.
// The submenu of the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/submenu
func (m_ MenuItem) SetSubmenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubmenu:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/subtitle
func (m_ MenuItem) Subtitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/subtitle
func (m_ MenuItem) SetSubtitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The menu item’s tag.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/tag
func (m_ MenuItem) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// The menu item’s tag.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/tag
func (m_ MenuItem) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}

// The menu item’s target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/target
func (m_ MenuItem) Target() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The menu item’s target.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/target
func (m_ MenuItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTarget:"), value)
}

// The menu item’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/title
func (m_ MenuItem) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The menu item’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/title
func (m_ MenuItem) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A help tag for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/toolTip
func (m_ MenuItem) ToolTip() string {
	rv := objc.Send[string](m_.ID, objc.Sel("toolTip"))
	return rv
}


// SetToolTip sets the value of the toolTip property.
// A help tag for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/toolTip
func (m_ MenuItem) SetToolTip(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// The user-assigned key equivalent for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/userKeyEquivalent
func (m_ MenuItem) UserKeyEquivalent() string {
	rv := objc.Send[string](m_.ID, objc.Sel("userKeyEquivalent"))
	return rv
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (m_ MenuItem) UsesUserKeyEquivalents() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesUserKeyEquivalents"))
	return rv
}


// SetUsesUserKeyEquivalents sets the value of the usesUserKeyEquivalents property.
// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (m_ MenuItem) SetUsesUserKeyEquivalents(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesUserKeyEquivalents:"), value)
}

// The content view for the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/view
func (m_ MenuItem) View() NSView {
	rv := objc.Send[NSView](m_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The content view for the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/view
func (m_ MenuItem) SetView(value IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setView:"), value)
}

// An array of standard menu items related to Writing Tools. Each call to this method returns an array of newly allocated instances of NSMenuItem.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/writingToolsItems
func (m_ MenuItem) WritingToolsItems() []MenuItem {
	rv := objc.Send[[]MenuItem](m_.ID, objc.Sel("writingToolsItems"))
	return rv
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

// The menu item whose submenu contains the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) Parent() NSMenuItem {
	rv := objc.Send[NSMenuItem](m_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// The menu item whose submenu contains the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/parent
func (m_ MenuItem) SetParent(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParent:"), value)
}


