// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	Validate()
}

// A single item that appears in a window’s toolbar.
//
// An object displays an image and text string in the toolbar area of a window. You can also create toolbar items that display custom views you provide. Toolbar items provide fast access to common commands or features in the window. For example, the Finder window uses toolbar items to help someone navigate the file system. You typically create toolbar items at the same time you create your window’s toolbar. The system provides some standard items like spacers you can include in your toolbar. It also provides items that display standard interfaces like the color panel or font panel. For any custom toolbar items you create, provide an action method to call when someone clicks the item. You can display your toolbar item’s content using a custom view if you prefer, rather than an image and text label. If you specify an object for the view, the system automatically adjusts the minimum and maximum size of the search field to the system-standard values.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func NewToolbarItemWithItemIdentifier(itemIdentifier unsafe.Pointer) ToolbarItem {
	instance := getToolbarItemClass().Alloc()
	rv := objc.Send[ToolbarItem](instance.ID, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	rv.Autorelease()
	return rv
}



// Creates a toolbar item with property values from the specified bar button item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:barButtonItem:)
func NewToolbarItemWithItemIdentifierBarButtonItem(itemIdentifier unsafe.Pointer, barButtonItem unsafe.Pointer) ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(getToolbarItemClass().class), objc.Sel("itemWithItemIdentifier:barButtonItem:"), itemIdentifier, barButtonItem)
	return rv
}


// Creates a toolbar item with property values from the specified bar button item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:barButtonItem:)
func (tc _ToolbarItemClass) ItemWithItemIdentifierBarButtonItem(itemIdentifier unsafe.Pointer, barButtonItem unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("itemWithItemIdentifier:barButtonItem:"), itemIdentifier, barButtonItem)
	return rv
}

// Validates the toolbar item’s menu and its ability to perfrom its action.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/validate()
func (t_ ToolbarItem) Validate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("validate"))
}

// The action method to call when someone clicks on the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t_ ToolbarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](t_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The action method to call when someone clicks on the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t_ ToolbarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAction:"), value)
}

// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/allowsDuplicatesInToolbar
func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDuplicatesInToolbar"))
	return rv
}

// A Boolean value that indicates whether the toolbar automatically validates the item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autovalidates"))
	return rv
}


// SetAutovalidates sets the value of the autovalidates property.
// A Boolean value that indicates whether the toolbar automatically validates the item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutovalidates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) BackgroundTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("backgroundTintColor"))
	return rv
}


// SetBackgroundTintColor sets the value of the backgroundTintColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) SetBackgroundTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundTintColor:"), value)
}

// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/badge-2b38p
func (t_ ToolbarItem) Badge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/badge-2b38p
func (t_ ToolbarItem) SetBadge(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBadge:"), value)
}

// The image to display for the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t_ ToolbarItem) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image to display for the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t_ ToolbarItem) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}

// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isBordered
func (t_ ToolbarItem) Bordered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bordered"))
	return rv
}


// SetBordered sets the value of the bordered property.
// A Boolean value that indicates whether the toolbar item has a bordered style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isBordered
func (t_ ToolbarItem) SetBordered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBordered:"), value)
}

// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t_ ToolbarItem) Enabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether the item is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t_ ToolbarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) Hidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidden"))
	return rv
}


// SetHidden sets the value of the hidden property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidden:"), value)
}

// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) Navigational() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("navigational"))
	return rv
}


// SetNavigational sets the value of the navigational property.
// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNavigational:"), value)
}

// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isVisible
func (t_ ToolbarItem) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}

// The value you use to identify the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemIdentifier
func (t_ ToolbarItem) ItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemIdentifier"))
	return rv
}

// The menu item to use for the toolbar item is in the overflow menu in a Mac app built with Mac Catalyst.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemMenuFormRepresentation
func (t_ ToolbarItem) ItemMenuFormRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemMenuFormRepresentation"))
	return rv
}


// SetItemMenuFormRepresentation sets the value of the itemMenuFormRepresentation property.
// The menu item to use for the toolbar item is in the overflow menu in a Mac app built with Mac Catalyst.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemMenuFormRepresentation
func (t_ ToolbarItem) SetItemMenuFormRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItemMenuFormRepresentation:"), value)
}

// The label that appears for this item in the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t_ ToolbarItem) Label() string {
	rv := objc.Send[string](t_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// The label that appears for this item in the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t_ ToolbarItem) SetLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The toolbar item’s maximum size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/maxSize
func (t_ ToolbarItem) MaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("maxSize"))
	return rv
}


// SetMaxSize sets the value of the maxSize property.
// The toolbar item’s maximum size.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/maxSize
func (t_ ToolbarItem) SetMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}

// The menu item to use when the toolbar item is in the overflow menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/menuFormRepresentation
func (t_ ToolbarItem) MenuFormRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("menuFormRepresentation"))
	return rv
}


// SetMenuFormRepresentation sets the value of the menuFormRepresentation property.
// The menu item to use when the toolbar item is in the overflow menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/menuFormRepresentation
func (t_ ToolbarItem) SetMenuFormRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMenuFormRepresentation:"), value)
}

// The toolbar item’s minimum size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/minSize
func (t_ ToolbarItem) MinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("minSize"))
	return rv
}


// SetMinSize sets the value of the minSize property.
// The toolbar item’s minimum size.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/minSize
func (t_ ToolbarItem) SetMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}

// The label that appears when the toolbar item is in the customization palette.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t_ ToolbarItem) PaletteLabel() string {
	rv := objc.Send[string](t_.ID, objc.Sel("paletteLabel"))
	return rv
}


// SetPaletteLabel sets the value of the paletteLabel property.
// The label that appears when the toolbar item is in the customization palette.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t_ ToolbarItem) SetPaletteLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPaletteLabel:"), objc.String(value))
}

// The set of labels that the item might display.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/possibleLabels
func (t_ ToolbarItem) PossibleLabels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("possibleLabels"))
	return rv
}


// SetPossibleLabels sets the value of the possibleLabels property.
// The set of labels that the item might display.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/possibleLabels
func (t_ ToolbarItem) SetPossibleLabels(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPossibleLabels:"), value)
}

// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t_ ToolbarItem) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}

// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t_ ToolbarItem) Tag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// An integer tag you can use to identify the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t_ ToolbarItem) SetTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTag:"), value)
}

// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t_ ToolbarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The object that defines the action method the toolbar item calls when clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t_ ToolbarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTarget:"), value)
}

// The title of the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/title
func (t_ ToolbarItem) Title() string {
	rv := objc.Send[string](t_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/title
func (t_ ToolbarItem) SetTitle(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The tooltip to display when someone hovers over the item in the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t_ ToolbarItem) ToolTip() string {
	rv := objc.Send[string](t_.ID, objc.Sel("toolTip"))
	return rv
}


// SetToolTip sets the value of the toolTip property.
// The tooltip to display when someone hovers over the item in the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t_ ToolbarItem) SetToolTip(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// The toolbar that currently includes the item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolbar
func (t_ ToolbarItem) Toolbar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("toolbar"))
	return rv
}

// The custom view you use to draw the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/view
func (t_ ToolbarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The custom view you use to draw the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/view
func (t_ ToolbarItem) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}

// The display priority associated with the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t_ ToolbarItem) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// SetVisibilityPriority sets the value of the visibilityPriority property.
// The display priority associated with the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t_ ToolbarItem) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}


