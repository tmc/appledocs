// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StatusItem] class.
var (
	StatusItemClass     _StatusItemClass
	StatusItemClassOnce sync.Once
)

func getStatusItemClass() _StatusItemClass {
	StatusItemClassOnce.Do(func() {
		StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}
	})
	return StatusItemClass
}

type _StatusItemClass struct {
	class objc.Class
}

// An interface definition for the [StatusItem] class.
type IStatusItem interface {
	objectivec.IObject
	DrawStatusBarBackgroundInRectWithHighlight(rect coregraphics.CGRect, highlight bool)
	PopUpStatusItemMenu(menu IMenu)
	SendActionOn(mask EventMask) int
}

// An individual element displayed in the system menu bar.
//
// The method creates instances of this class and automatically adds them to the menu bar. Use the property to customize the appearance and behavior of the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem
type StatusItem struct {
	objectivec.Object
}

// StatusItemFrom constructs a [StatusItem] from an unsafe.Pointer.
//
// An individual element displayed in the system menu bar.
func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StatusItemClass) New() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusItem) Init() StatusItem {
	rv := objc.Send[StatusItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusItem) Autorelease() StatusItem {
	rv := objc.Send[StatusItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusItem creates a new StatusItem instance.
func NewStatusItem() StatusItem {
	return getStatusItemClass().New()
}


// Draws the menu background pattern for a custom status-bar item in regular or highlight pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/drawStatusBarBackground(in:withHighlight:)
func (s_ StatusItem) DrawStatusBarBackgroundInRectWithHighlight(rect coregraphics.CGRect, highlight bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawStatusBarBackgroundInRect:withHighlight:"), rect, highlight)
}

// Displays a menu under a custom status bar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/popUpMenu(_:)
func (s_ StatusItem) PopUpStatusItemMenu(menu IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("popUpStatusItemMenu:"), menu)
}

// Sets the conditions on which the status item sends action messages to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/sendAction(on:)
func (s_ StatusItem) SendActionOn(mask EventMask) int {
	rv := objc.Send[int](s_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}

// The selector that is sent to the status item’s target when the status item is clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/action
func (s_ StatusItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The selector that is sent to the status item’s target when the status item is clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/action
func (s_ StatusItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}

// The alternate image to be displayed when a status bar item is highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/alternateImage
func (s_ StatusItem) AlternateImage() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// SetAlternateImage sets the value of the alternateImage property.
// The alternate image to be displayed when a status bar item is highlighted.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/alternateImage
func (s_ StatusItem) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateImage:"), value)
}

// The attributed string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/attributedTitle
func (s_ StatusItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](s_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// The attributed string that is displayed at the status item’s position in the status bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/attributedTitle
func (s_ StatusItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttributedTitle:"), value)
}

// A unique name for saving and restoring information about a status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/autosaveName-swift.property
func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.Send[StatusItemAutosaveName](s_.ID, objc.Sel("autosaveName"))
	return rv
}


// SetAutosaveName sets the value of the autosaveName property.
// A unique name for saving and restoring information about a status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/autosaveName-swift.property
func (s_ StatusItem) SetAutosaveName(value IStatusItemAutosaveName) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutosaveName:"), value)
}

// The set of allowed behaviors for the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/behavior-swift.property
func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.Send[StatusItemBehavior](s_.ID, objc.Sel("behavior"))
	return rv
}


// SetBehavior sets the value of the behavior property.
// The set of allowed behaviors for the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/behavior-swift.property
func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBehavior:"), value)
}

// The button displayed in the status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/button
func (s_ StatusItem) Button() NSStatusBarButton {
	rv := objc.Send[NSStatusBarButton](s_.ID, objc.Sel("button"))
	return rv
}

// The selector that is sent to the status item’s target when the status item is double-clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/doubleAction
func (s_ StatusItem) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("doubleAction"))
	return rv
}


// SetDoubleAction sets the value of the doubleAction property.
// The selector that is sent to the status item’s target when the status item is double-clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/doubleAction
func (s_ StatusItem) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleAction:"), value)
}

// A Boolean that indicates whether the status item is highlighted when it is clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/highlightMode
func (s_ StatusItem) HighlightMode() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("highlightMode"))
	return rv
}


// SetHighlightMode sets the value of the highlightMode property.
// A Boolean that indicates whether the status item is highlighted when it is clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/highlightMode
func (s_ StatusItem) SetHighlightMode(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlightMode:"), value)
}

// The image that is displayed at the status item’s position in the status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/image
func (s_ StatusItem) Image() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image that is displayed at the status item’s position in the status bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/image
func (s_ StatusItem) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}

// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isEnabled
func (s_ StatusItem) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean that indicates whether the status item is enabled to respond to clicks.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isEnabled
func (s_ StatusItem) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}

// A Boolean value indicating if the menu bar currently displays the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isVisible
func (s_ StatusItem) Visible() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("visible"))
	return rv
}


// SetVisible sets the value of the visible property.
// A Boolean value indicating if the menu bar currently displays the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isVisible
func (s_ StatusItem) SetVisible(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisible:"), value)
}

// The amount of space in the status bar that should be allocated to the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/length
func (s_ StatusItem) Length() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("length"))
	return rv
}


// SetLength sets the value of the length property.
// The amount of space in the status bar that should be allocated to the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/length
func (s_ StatusItem) SetLength(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLength:"), value)
}

// The pull-down menu displayed when the user clicks the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/menu
func (s_ StatusItem) Menu() NSMenu {
	rv := objc.Send[NSMenu](s_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The pull-down menu displayed when the user clicks the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/menu
func (s_ StatusItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenu:"), value)
}

// The status bar that displays the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/statusBar
func (s_ StatusItem) StatusBar() NSStatusBar {
	rv := objc.Send[NSStatusBar](s_.ID, objc.Sel("statusBar"))
	return rv
}

// The target object to which the status item’s action message is sent when the status item is clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/target
func (s_ StatusItem) Target() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The target object to which the status item’s action message is sent when the status item is clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/target
func (s_ StatusItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}

// The string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/title
func (s_ StatusItem) Title() string {
	rv := objc.Send[string](s_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The string that is displayed at the status item’s position in the status bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/title
func (s_ StatusItem) SetTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The tool tip string that is displayed when the cursor pauses over the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/toolTip
func (s_ StatusItem) ToolTip() string {
	rv := objc.Send[string](s_.ID, objc.Sel("toolTip"))
	return rv
}


// SetToolTip sets the value of the toolTip property.
// The tool tip string that is displayed when the cursor pauses over the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/toolTip
func (s_ StatusItem) SetToolTip(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// The custom view that is displayed at the status item’s position in the status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/view
func (s_ StatusItem) View() NSView {
	rv := objc.Send[NSView](s_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The custom view that is displayed at the status item’s position in the status bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/view
func (s_ StatusItem) SetView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setView:"), value)
}

// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isenabled
func (s_ StatusItem) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean that indicates whether the status item is enabled to respond to clicks.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isenabled
func (s_ StatusItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value indicating if the menu bar currently displays the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isvisible
func (s_ StatusItem) IsVisible() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVisible"))
	return rv
}


// SetIsVisible sets the value of the isVisible property.
// A Boolean value indicating if the menu bar currently displays the status item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isvisible
func (s_ StatusItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVisible:"), value)
}



