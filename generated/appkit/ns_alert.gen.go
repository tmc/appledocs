// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAlert */


/* debug [class_header]: Header for NSAlert */
// The class instance for the [Alert] class.
var (
	AlertClass     _AlertClass
	AlertClassOnce sync.Once
)

func getAlertClass() _AlertClass {
	AlertClassOnce.Do(func() {
		AlertClass = _AlertClass{objc.GetClass("NSAlert")}
	})
	return AlertClass
}

type _AlertClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Alert */
// An interface definition for the [Alert] class.
type IAlert interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Alert */
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	AlertStyle() objectivec.IObject
	SetAlertStyle(value objectivec.IObject)
	Buttons() IButton
	SetButtons(value IButton)
	Delegate() objc.IObject /* cross-framework: AlertDelegate */
	SetDelegate(value objc.IObject /* cross-framework: AlertDelegate */)
	HelpAnchor() objectivec.IObject
	SetHelpAnchor(value objectivec.IObject)
	Icon() IImage
	SetIcon(value IImage)
	InformativeText() objc.IObject /* cross-framework: NSString */
	SetInformativeText(value objc.IObject /* cross-framework: NSString */)
	MessageText() objc.IObject /* cross-framework: NSString */
	SetMessageText(value objc.IObject /* cross-framework: NSString */)
	ShowsHelp() bool
	SetShowsHelp(value bool)
	ShowsSuppressionButton() bool
	SetShowsSuppressionButton(value bool)
	SuppressionButton() IButton
	SetSuppressionButton(value IButton)
	Window() IWindow
	SetWindow(value IWindow)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Alert */
	// methods:
	BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Alert */
// Alloc allocates a new instance without initialization.
func (ac _AlertClass) Alloc() Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AlertClass) New() Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Alert) Init() Alert {
	rv := objc.Send[Alert](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Alert) Autorelease() Alert {
	rv := objc.Send[Alert](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAlert creates a new Alert instance.
func NewAlert() Alert {
	return getAlertClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Alert */
// A modal dialog or sheet attached to a document window.
//
// The methods of the class allow you to specify alert level, alert text, button titles, and a custom icon should you require it. The class also lets your alerts display a help button and provides ways for apps to offer help specific to an alert. To display an alert as a sheet, call the method; to display one as an app-modal dialog, use the method. By design, an object is intended for a single alert—that is, an alert with a unique combination of title, buttons, and so on—that is displayed upon a particular condition. You should create an object for each alert dialog, creating it only when you need to display an alert, and release it when you are done. If you have a particular alert dialog that you need to show repeatedly, you can retain and reuse an instance of for this dialog. After creating an alert using one of the alert creation methods, you can customize it further prior to displaying it by customizing its attributes. See . Unless you must maintain compatibility with existing alert-processing code that uses the function-based API, you should allocate ( ) and initialize ( ) the alert object, and then set its attributes using the appropriate methods of the class.


// A modal dialog or sheet attached to a document window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert
type Alert struct {
	objectivec.Object
}

// AlertFrom constructs a [Alert] from an unsafe.Pointer.
//
// A modal dialog or sheet attached to a document window.
func AlertFrom(ptr unsafe.Pointer) Alert {
	return Alert{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Alert *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Alert */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Alert */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Alert */

// Runs the alert modally as a sheet attached to the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), sheetWindow, handler)
}/* debug [instance_methods/method]: BeginSheetModalForWindowCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Alert */

// The alert’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/accessoryview
func (a_ Alert) AccessoryView() IView {
	rv := objc.Send[View](a_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// The alert’s accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/accessoryview
func (a_ Alert) SetAccessoryView(value IView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// Indicates the alert’s severity level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/alertstyle
func (a_ Alert) AlertStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("alertStyle"))
	return rv
}/* debug [instance_properties/getter]: alertStyle */


// Indicates the alert’s severity level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/alertstyle
func (a_ Alert) SetAlertStyle(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlertStyle:"), value)
}/* debug [instance_properties/setter]: alertStyle */


// The array of response buttons for the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/buttons
func (a_ Alert) Buttons() IButton {
	rv := objc.Send[Button](a_.ID, objc.Sel("buttons"))
	return rv
}/* debug [instance_properties/getter]: buttons */


// The array of response buttons for the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/buttons
func (a_ Alert) SetButtons(value IButton) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setButtons:"), value)
}/* debug [instance_properties/setter]: buttons */


// The alert’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/delegate
func (a_ Alert) Delegate() objc.IObject /* cross-framework: AlertDelegate */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The alert’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/delegate
func (a_ Alert) SetDelegate(value objc.IObject /* cross-framework: AlertDelegate */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The alert’s HTML help anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/helpanchor
func (a_ Alert) HelpAnchor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("helpAnchor"))
	return rv
}/* debug [instance_properties/getter]: helpAnchor */


// The alert’s HTML help anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/helpanchor
func (a_ Alert) SetHelpAnchor(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHelpAnchor:"), value)
}/* debug [instance_properties/setter]: helpAnchor */


// The custom icon displayed in the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/icon
func (a_ Alert) Icon() IImage {
	rv := objc.Send[Image](a_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The custom icon displayed in the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/icon
func (a_ Alert) SetIcon(value IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIcon:"), value)
}/* debug [instance_properties/setter]: icon */


// The alert’s informative text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/informativetext
func (a_ Alert) InformativeText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("informativeText"))
	return rv
}/* debug [instance_properties/getter]: informativeText */


// The alert’s informative text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/informativetext
func (a_ Alert) SetInformativeText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInformativeText:"), value)
}/* debug [instance_properties/setter]: informativeText */


// The alert’s message text or title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/messagetext
func (a_ Alert) MessageText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("messageText"))
	return rv
}/* debug [instance_properties/getter]: messageText */


// The alert’s message text or title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/messagetext
func (a_ Alert) SetMessageText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMessageText:"), value)
}/* debug [instance_properties/setter]: messageText */


// Specifies whether the alert has a help button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showshelp
func (a_ Alert) ShowsHelp() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showsHelp"))
	return rv
}/* debug [instance_properties/getter]: showsHelp */


// Specifies whether the alert has a help button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showshelp
func (a_ Alert) SetShowsHelp(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowsHelp:"), value)
}/* debug [instance_properties/setter]: showsHelp */


// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showssuppressionbutton
func (a_ Alert) ShowsSuppressionButton() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showsSuppressionButton"))
	return rv
}/* debug [instance_properties/getter]: showsSuppressionButton */


// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showssuppressionbutton
func (a_ Alert) SetShowsSuppressionButton(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowsSuppressionButton:"), value)
}/* debug [instance_properties/setter]: showsSuppressionButton */


// The alert’s suppression checkbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/suppressionbutton
func (a_ Alert) SuppressionButton() IButton {
	rv := objc.Send[Button](a_.ID, objc.Sel("suppressionButton"))
	return rv
}/* debug [instance_properties/getter]: suppressionButton */


// The alert’s suppression checkbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/suppressionbutton
func (a_ Alert) SetSuppressionButton(value IButton) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSuppressionButton:"), value)
}/* debug [instance_properties/setter]: suppressionButton */


// The app-modal panel or document-modal sheet that corresponds to the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/window
func (a_ Alert) Window() IWindow {
	rv := objc.Send[Window](a_.ID, objc.Sel("window"))
	return rv
}/* debug [instance_properties/getter]: window */


// The app-modal panel or document-modal sheet that corresponds to the alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/window
func (a_ Alert) SetWindow(value IWindow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWindow:"), value)
}/* debug [instance_properties/setter]: window */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAlert */



