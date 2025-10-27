// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	Enabled() bool
	SetEnabled(value bool)
	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	AttributedTitle() foundation.foundation.INSAttributedString
	SetAttributedTitle(value foundation.foundation.INSAttributedString)
	AutosaveName() objectivec.IObject
	SetAutosaveName(value objectivec.IObject)
	Behavior() objectivec.IObject
	SetBehavior(value objectivec.IObject)
	Button() IStatusBarButton
	SetButton(value IStatusBarButton)
	DoubleAction() objectivec.IObject
	SetDoubleAction(value objectivec.IObject)
	HighlightMode() bool
	SetHighlightMode(value bool)
	Image() IImage
	SetImage(value IImage)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
	Length() float64
	SetLength(value float64)
	Menu() IMenu
	SetMenu(value IMenu)
	StatusBar() IStatusBar
	SetStatusBar(value IStatusBar)
	Title() foundation.foundation.INSString
	SetTitle(value foundation.foundation.INSString)
	ToolTip() foundation.foundation.INSString
	SetToolTip(value foundation.foundation.INSString)
	View() IView
	SetView(value IView)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An individual element displayed in the system menu bar.
//
// The method creates instances of this class and automatically adds them to the menu bar. Use the property to customize the appearance and behavior of the status item.


// An individual element displayed in the system menu bar.
//
// [Full Topic]
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

























// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isEnabled
func (s_ StatusItem) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/isEnabled
func (s_ StatusItem) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}


// The selector that is sent to the status item’s target when the status item is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/action
func (s_ StatusItem) Action() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("action"))
	return rv
}


// The selector that is sent to the status item’s target when the status item is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/action
func (s_ StatusItem) SetAction(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}


// The alternate image to be displayed when a status bar item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/alternateimage
func (s_ StatusItem) AlternateImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// The alternate image to be displayed when a status bar item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/alternateimage
func (s_ StatusItem) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateImage:"), value)
}


// The attributed string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/attributedtitle
func (s_ StatusItem) AttributedTitle() foundation.foundation.INSAttributedString {
	rv := objc.Send[foundation.NSAttributedString](s_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The attributed string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/attributedtitle
func (s_ StatusItem) SetAttributedTitle(value foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttributedTitle:"), value)
}


// A unique name for saving and restoring information about a status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/autosavename-swift.property
func (s_ StatusItem) AutosaveName() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("autosaveName"))
	return rv
}


// A unique name for saving and restoring information about a status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/autosavename-swift.property
func (s_ StatusItem) SetAutosaveName(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutosaveName:"), value)
}


// The set of allowed behaviors for the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/behavior-swift.property
func (s_ StatusItem) Behavior() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("behavior"))
	return rv
}


// The set of allowed behaviors for the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/behavior-swift.property
func (s_ StatusItem) SetBehavior(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBehavior:"), value)
}


// The button displayed in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/button
func (s_ StatusItem) Button() IStatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID, objc.Sel("button"))
	return rv
}


// The button displayed in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/button
func (s_ StatusItem) SetButton(value IStatusBarButton) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButton:"), value)
}


// The selector that is sent to the status item’s target when the status item is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/doubleaction
func (s_ StatusItem) DoubleAction() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("doubleAction"))
	return rv
}


// The selector that is sent to the status item’s target when the status item is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/doubleaction
func (s_ StatusItem) SetDoubleAction(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleAction:"), value)
}


// A Boolean that indicates whether the status item is highlighted when it is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/highlightmode
func (s_ StatusItem) HighlightMode() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("highlightMode"))
	return rv
}


// A Boolean that indicates whether the status item is highlighted when it is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/highlightmode
func (s_ StatusItem) SetHighlightMode(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlightMode:"), value)
}


// The image that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/image
func (s_ StatusItem) Image() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// The image that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/image
func (s_ StatusItem) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}


// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isenabled
func (s_ StatusItem) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean that indicates whether the status item is enabled to respond to clicks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isenabled
func (s_ StatusItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value indicating if the menu bar currently displays the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isvisible
func (s_ StatusItem) IsVisible() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value indicating if the menu bar currently displays the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/isvisible
func (s_ StatusItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVisible:"), value)
}


// The amount of space in the status bar that should be allocated to the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/length
func (s_ StatusItem) Length() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("length"))
	return rv
}


// The amount of space in the status bar that should be allocated to the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/length
func (s_ StatusItem) SetLength(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLength:"), value)
}


// The pull-down menu displayed when the user clicks the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/menu
func (s_ StatusItem) Menu() IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("menu"))
	return rv
}


// The pull-down menu displayed when the user clicks the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/menu
func (s_ StatusItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenu:"), value)
}


// The status bar that displays the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/statusbar
func (s_ StatusItem) StatusBar() IStatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("statusBar"))
	return rv
}


// The status bar that displays the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/statusbar
func (s_ StatusItem) SetStatusBar(value IStatusBar) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStatusBar:"), value)
}


// The string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/title
func (s_ StatusItem) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}


// The string that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/title
func (s_ StatusItem) SetTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}


// The tool tip string that is displayed when the cursor pauses over the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/tooltip
func (s_ StatusItem) ToolTip() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("toolTip"))
	return rv
}


// The tool tip string that is displayed when the cursor pauses over the status item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/tooltip
func (s_ StatusItem) SetToolTip(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToolTip:"), value)
}


// The custom view that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/view
func (s_ StatusItem) View() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("view"))
	return rv
}


// The custom view that is displayed at the status item’s position in the status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/view
func (s_ StatusItem) SetView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setView:"), value)
}








