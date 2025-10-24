// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ToolbarItem] class.
var (
	ToolbarItemClass     _ToolbarItemClass
	ToolbarItemClassOnce sync.Once
)

func getToolbarItemClass() _ToolbarItemClass {
	ToolbarItemClassOnce.Do(func() {
		ToolbarItemClass = _ToolbarItemClass{objc.GetClass("NSToolbarItem")}
	})
	return ToolbarItemClass
}

type _ToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolbarItem] class.
type IToolbarItem interface {
	objectivec.IObject
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	AllowsDuplicatesInToolbar() bool
	Autovalidates() bool
	SetAutovalidates(value bool)
	BackgroundTintColor() IColor
	SetBackgroundTintColor(value IColor)
	Badge() IItemBadge
	SetBadge(value IItemBadge)
	Image() IImage
	SetImage(value IImage)
	Bordered() bool
	SetBordered(value bool)
	Enabled() bool
	SetEnabled(value bool)
	Hidden() bool
	SetHidden(value bool)
	Navigational() bool
	SetNavigational(value bool)
	Visible() bool
	ItemIdentifier() objc.IObject /* cross-framework: ToolbarItemIdentifier */
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MaxSize() objc.IObject /* cross-framework: Size */
	SetMaxSize(value objc.IObject /* cross-framework: Size */)
	MenuFormRepresentation() IMenuItem
	SetMenuFormRepresentation(value IMenuItem)
	MinSize() objc.IObject /* cross-framework: Size */
	SetMinSize(value objc.IObject /* cross-framework: Size */)
	PaletteLabel() objc.IObject /* cross-framework: NSString */
	SetPaletteLabel(value objc.IObject /* cross-framework: NSString */)
	PossibleLabels() unsafe.Pointer
	SetPossibleLabels(value unsafe.Pointer)
	Style() ToolbarItemStyle
	SetStyle(value ToolbarItemStyle)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	ToolTip() objc.IObject /* cross-framework: NSString */
	SetToolTip(value objc.IObject /* cross-framework: NSString */)
	Toolbar() IToolbar
	View() IView
	SetView(value IView)
	VisibilityPriority() objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */
	SetVisibilityPriority(value objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */)
	IsBordered() bool
	SetIsBordered(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsNavigational() bool
	SetIsNavigational(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
	// methods:
	Validate()
}

// A single item that appears in a window’s toolbar.
//
// An object displays an image and text string in the toolbar area of a window. You can also create toolbar items that display custom views you provide. Toolbar items provide fast access to common commands or features in the window. For example, the Finder window uses toolbar items to help someone navigate the file system. You typically create toolbar items at the same time you create your window’s toolbar. The system provides some standard items like spacers you can include in your toolbar. It also provides items that display standard interfaces like the color panel or font panel. For any custom toolbar items you create, provide an action method to call when someone clicks the item. You can display your toolbar item’s content using a custom view if you prefer, rather than an image and text label. If you specify an object for the view, the system automatically adjusts the minimum and maximum size of the search field to the system-standard values.


// A single item that appears in a window’s toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem
type ToolbarItem struct {
	objectivec.Object
}

// ToolbarItemFrom constructs a [ToolbarItem] from an unsafe.Pointer.
//
// A single item that appears in a window’s toolbar.
func ToolbarItemFrom(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolbarItem) Init() ToolbarItem {
	rv := objc.Send[ToolbarItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolbarItem) Autorelease() ToolbarItem {
	rv := objc.Send[ToolbarItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolbarItem creates a new ToolbarItem instance.
func NewToolbarItem() ToolbarItem {
	return getToolbarItemClass().New()
}



// Creates a toolbar item with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func NewToolbarItemWithItemIdentifier(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */) ToolbarItem {
	instance := getToolbarItemClass().Alloc()
	rv := objc.Send[ToolbarItem](instance.ID, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	rv.Autorelease()
	return rv
}


// Creates a toolbar item with property values from the specified bar button item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:barButtonItem:)
func NewToolbarItemWithItemIdentifierBarButtonItem(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, barButtonItem BarButtonItem /* not a class type */) ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(getToolbarItemClass().class), objc.Sel("itemWithItemIdentifier:barButtonItem:"), itemIdentifier, barButtonItem)
	return rv
}



// Creates a toolbar item with property values from the specified bar button item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:barButtonItem:)
func (tc _ToolbarItemClass) ItemWithItemIdentifierBarButtonItem(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, barButtonItem BarButtonItem /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("itemWithItemIdentifier:barButtonItem:"), itemIdentifier, barButtonItem)
	return rv
}


// Validates the toolbar item’s menu and its ability to perfrom its action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/validate()
func (t_ ToolbarItem) Validate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("validate"))
}


// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t_ ToolbarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](t_.ID, objc.Sel("action"))
	return rv
}


// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t_ ToolbarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAction:"), value)
}


// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/allowsDuplicatesInToolbar
func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDuplicatesInToolbar"))
	return rv
}


// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autovalidates"))
	return rv
}


// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutovalidates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) BackgroundTintColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundTintColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) SetBackgroundTintColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundTintColor:"), value)
}


// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/badge-2b38p
func (t_ ToolbarItem) Badge() IItemBadge {
	rv := objc.Send[ItemBadge](t_.ID, objc.Sel("badge"))
	return rv
}


// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/badge-2b38p
func (t_ ToolbarItem) SetBadge(value IItemBadge) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBadge:"), value)
}


// The image to display for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t_ ToolbarItem) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// The image to display for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t_ ToolbarItem) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isBordered
func (t_ ToolbarItem) Bordered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bordered"))
	return rv
}


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isBordered
func (t_ ToolbarItem) SetBordered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBordered:"), value)
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t_ ToolbarItem) Enabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t_ ToolbarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) Hidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidden:"), value)
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) Navigational() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("navigational"))
	return rv
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNavigational:"), value)
}


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isVisible
func (t_ ToolbarItem) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}


// The value you use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemIdentifier
func (t_ ToolbarItem) ItemIdentifier() objc.IObject /* cross-framework: ToolbarItemIdentifier */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("itemIdentifier"))
	return rv
}


// The label that appears for this item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t_ ToolbarItem) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}


// The label that appears for this item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t_ ToolbarItem) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}


// The toolbar item’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/maxSize
func (t_ ToolbarItem) MaxSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("maxSize"))
	return rv
}


// The toolbar item’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/maxSize
func (t_ ToolbarItem) SetMaxSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}


// The menu item to use when the toolbar item is in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/menuFormRepresentation
func (t_ ToolbarItem) MenuFormRepresentation() IMenuItem {
	rv := objc.Send[MenuItem](t_.ID, objc.Sel("menuFormRepresentation"))
	return rv
}


// The menu item to use when the toolbar item is in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/menuFormRepresentation
func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMenuFormRepresentation:"), value)
}


// The toolbar item’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/minSize
func (t_ ToolbarItem) MinSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("minSize"))
	return rv
}


// The toolbar item’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/minSize
func (t_ ToolbarItem) SetMinSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}


// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t_ ToolbarItem) PaletteLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("paletteLabel"))
	return rv
}


// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t_ ToolbarItem) SetPaletteLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPaletteLabel:"), value)
}


// The set of labels that the item might display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/possibleLabels
func (t_ ToolbarItem) PossibleLabels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("possibleLabels"))
	return rv
}


// The set of labels that the item might display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/possibleLabels
func (t_ ToolbarItem) SetPossibleLabels(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPossibleLabels:"), value)
}


// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) Style() ToolbarItemStyle {
	rv := objc.Send[ToolbarItemStyle](t_.ID, objc.Sel("style"))
	return rv
}


// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) SetStyle(value ToolbarItemStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t_ ToolbarItem) Tag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("tag"))
	return rv
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t_ ToolbarItem) SetTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTag:"), value)
}


// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t_ ToolbarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("target"))
	return rv
}


// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t_ ToolbarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTarget:"), value)
}


// The title of the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/title
func (t_ ToolbarItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("title"))
	return rv
}


// The title of the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/title
func (t_ ToolbarItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), value)
}


// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t_ ToolbarItem) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("toolTip"))
	return rv
}


// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t_ ToolbarItem) SetToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), value)
}


// The toolbar that currently includes the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolbar
func (t_ ToolbarItem) Toolbar() IToolbar {
	rv := objc.Send[Toolbar](t_.ID, objc.Sel("toolbar"))
	return rv
}


// The custom view you use to draw the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/view
func (t_ ToolbarItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// The custom view you use to draw the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/view
func (t_ ToolbarItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t_ ToolbarItem) VisibilityPriority() objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t_ ToolbarItem) SetVisibilityPriority(value objc.IObject /* cross-framework: ToolbarItemVisibilityPriority */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (t_ ToolbarItem) IsBordered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isBordered"))
	return rv
}


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (t_ ToolbarItem) SetIsBordered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsBordered:"), value)
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (t_ ToolbarItem) IsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (t_ ToolbarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (t_ ToolbarItem) IsHidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (t_ ToolbarItem) SetIsHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHidden:"), value)
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (t_ ToolbarItem) IsNavigational() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isNavigational"))
	return rv
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (t_ ToolbarItem) SetIsNavigational(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsNavigational:"), value)
}


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (t_ ToolbarItem) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (t_ ToolbarItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}


