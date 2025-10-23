// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	Bordered() bool
	SetBordered(value bool)
	Navigational() bool
	SetNavigational(value bool)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	AllowsDuplicatesInToolbar() bool
	SetAllowsDuplicatesInToolbar(value bool)
	Autovalidates() bool
	SetAutovalidates(value bool)
	BackgroundTintColor() IColor
	SetBackgroundTintColor(value IColor)
	Badge() ItemBadge
	SetBadge(value ItemBadge)
	Image() IImage
	SetImage(value IImage)
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
	ItemIdentifier() unsafe.Pointer
	SetItemIdentifier(value unsafe.Pointer)
	ItemMenuFormRepresentation() unsafe.Pointer
	SetItemMenuFormRepresentation(value unsafe.Pointer)
	Label() string
	SetLabel(value string)
	MaxSize() coregraphics.CGSize
	SetMaxSize(value coregraphics.CGSize)
	MenuFormRepresentation() MenuItem
	SetMenuFormRepresentation(value MenuItem)
	MinSize() coregraphics.CGSize
	SetMinSize(value coregraphics.CGSize)
	PaletteLabel() string
	SetPaletteLabel(value string)
	PossibleLabels() string
	SetPossibleLabels(value string)
	Tag() int
	SetTag(value int)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
	Title() string
	SetTitle(value string)
	ToolTip() string
	SetToolTip(value string)
	Toolbar() IToolbar
	SetToolbar(value IToolbar)
	View() IView
	SetView(value IView)
	VisibilityPriority() unsafe.Pointer
	SetVisibilityPriority(value unsafe.Pointer)
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


// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("style"))
	return rv
}


// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}


// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/action
func (t_ ToolbarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("action"))
	return rv
}


// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/action
func (t_ ToolbarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAction:"), value)
}


// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/allowsduplicatesintoolbar
func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDuplicatesInToolbar"))
	return rv
}


// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/allowsduplicatesintoolbar
func (t_ ToolbarItem) SetAllowsDuplicatesInToolbar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDuplicatesInToolbar:"), value)
}


// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/autovalidates
func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autovalidates"))
	return rv
}


// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/autovalidates
func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutovalidates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/backgroundtintcolor
func (t_ ToolbarItem) BackgroundTintColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundTintColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/backgroundtintcolor
func (t_ ToolbarItem) SetBackgroundTintColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundTintColor:"), value)
}


// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/badge-17r3r
func (t_ ToolbarItem) Badge() ItemBadge {
	rv := objc.Send[ItemBadge](t_.ID, objc.Sel("badge"))
	return rv
}


// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/badge-17r3r
func (t_ ToolbarItem) SetBadge(value ItemBadge) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBadge:"), value)
}


// The image to display for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/image
func (t_ ToolbarItem) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// The image to display for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/image
func (t_ ToolbarItem) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
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


// The value you use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/itemidentifier
func (t_ ToolbarItem) ItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemIdentifier"))
	return rv
}


// The value you use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/itemidentifier
func (t_ ToolbarItem) SetItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItemIdentifier:"), value)
}


// The menu item to use for the toolbar item is in the overflow menu in a Mac app built with Mac Catalyst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/itemmenuformrepresentation
func (t_ ToolbarItem) ItemMenuFormRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemMenuFormRepresentation"))
	return rv
}


// The menu item to use for the toolbar item is in the overflow menu in a Mac app built with Mac Catalyst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/itemmenuformrepresentation
func (t_ ToolbarItem) SetItemMenuFormRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItemMenuFormRepresentation:"), value)
}


// The label that appears for this item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/label
func (t_ ToolbarItem) Label() string {
	rv := objc.Send[string](t_.ID, objc.Sel("label"))
	return rv
}


// The label that appears for this item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/label
func (t_ ToolbarItem) SetLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), objc.String(value))
}


// The toolbar item’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/maxsize
func (t_ ToolbarItem) MaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("maxSize"))
	return rv
}


// The toolbar item’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/maxsize
func (t_ ToolbarItem) SetMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}


// The menu item to use when the toolbar item is in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/menuformrepresentation
func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := objc.Send[MenuItem](t_.ID, objc.Sel("menuFormRepresentation"))
	return rv
}


// The menu item to use when the toolbar item is in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/menuformrepresentation
func (t_ ToolbarItem) SetMenuFormRepresentation(value MenuItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMenuFormRepresentation:"), value)
}


// The toolbar item’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/minsize
func (t_ ToolbarItem) MinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("minSize"))
	return rv
}


// The toolbar item’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/minsize
func (t_ ToolbarItem) SetMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}


// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/palettelabel
func (t_ ToolbarItem) PaletteLabel() string {
	rv := objc.Send[string](t_.ID, objc.Sel("paletteLabel"))
	return rv
}


// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/palettelabel
func (t_ ToolbarItem) SetPaletteLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPaletteLabel:"), objc.String(value))
}


// The set of labels that the item might display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/possiblelabels
func (t_ ToolbarItem) PossibleLabels() string {
	rv := objc.Send[string](t_.ID, objc.Sel("possibleLabels"))
	return rv
}


// The set of labels that the item might display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/possiblelabels
func (t_ ToolbarItem) SetPossibleLabels(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPossibleLabels:"), objc.String(value))
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (t_ ToolbarItem) Tag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("tag"))
	return rv
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (t_ ToolbarItem) SetTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTag:"), value)
}


// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/target
func (t_ ToolbarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("target"))
	return rv
}


// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/target
func (t_ ToolbarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTarget:"), value)
}


// The title of the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/title
func (t_ ToolbarItem) Title() string {
	rv := objc.Send[string](t_.ID, objc.Sel("title"))
	return rv
}


// The title of the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/title
func (t_ ToolbarItem) SetTitle(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tooltip
func (t_ ToolbarItem) ToolTip() string {
	rv := objc.Send[string](t_.ID, objc.Sel("toolTip"))
	return rv
}


// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tooltip
func (t_ ToolbarItem) SetToolTip(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), objc.String(value))
}


// The toolbar that currently includes the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/toolbar
func (t_ ToolbarItem) Toolbar() IToolbar {
	rv := objc.Send[Toolbar](t_.ID, objc.Sel("toolbar"))
	return rv
}


// The toolbar that currently includes the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/toolbar
func (t_ ToolbarItem) SetToolbar(value IToolbar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolbar:"), value)
}


// The custom view you use to draw the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/view
func (t_ ToolbarItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// The custom view you use to draw the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/view
func (t_ ToolbarItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (t_ ToolbarItem) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (t_ ToolbarItem) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}



