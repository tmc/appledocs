// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Alert] class.
type IAlert interface {
	objectivec.IObject
	BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer)
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(window IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModal() ModalResponse
}

// A modal dialog or sheet attached to a document window.
//
// The methods of the class allow you to specify alert level, alert text, button titles, and a custom icon should you require it. The class also lets your alerts display a help button and provides ways for apps to offer help specific to an alert. To display an alert as a sheet, call the method; to display one as an app-modal dialog, use the method. By design, an object is intended for a single alert—that is, an alert with a unique combination of title, buttons, and so on—that is displayed upon a particular condition. You should create an object for each alert dialog, creating it only when you need to display an alert, and release it when you are done. If you have a particular alert dialog that you need to show repeatedly, you can retain and reuse an instance of for this dialog. After creating an alert using one of the alert creation methods, you can customize it further prior to displaying it by customizing its attributes. See . Unless you must maintain compatibility with existing alert-processing code that uses the function-based API, you should allocate ( ) and initialize ( ) the alert object, and then set its attributes using the appropriate methods of the class.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AlertClass) Alloc() Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns an alert initialized from information in an error object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/init(error:)
func NewAlertWithError(error_ foundation.IError) Alert {
	rv := objc.Send[Alert](objc.ID(getAlertClass().class), objc.Sel("alertWithError:"), error_)
	return rv
}


// Creates an alert compatible with alerts created using the function for display as a warning-style alert.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/alertWithMessageText:defaultButton:alternateButton:otherButton:informativeTextWithFormat:
func (ac _AlertClass) AlertWithMessageTextDefaultButtonAlternateButtonOtherButtonInformativeTextWithFormat(message string, defaultButton string, alternateButton string, otherButton string, format string) Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("alertWithMessageText:defaultButton:alternateButton:otherButton:informativeTextWithFormat:"), objc.String(message), objc.String(defaultButton), objc.String(alternateButton), objc.String(otherButton), objc.String(format))
	return rv
}

// Returns an alert initialized from information in an error object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/init(error:)
func (ac _AlertClass) AlertWithError(error_ foundation.IError) Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("alertWithError:"), error_)
	return rv
}

// Runs the alert modally as a sheet attached to the specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), sheetWindow, handler)
}

// Runs the alert modally as an alert sheet attached to a specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (a_ Alert) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(window IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), window, delegate, didEndSelector, contextInfo)
}

// Runs the alert as an app-modal dialog and returns the constant that identifies the button clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/runModal()
func (a_ Alert) RunModal() ModalResponse {
	rv := objc.Send[ModalResponse](a_.ID, objc.Sel("runModal"))
	return rv
}

// The array of response buttons for the alert.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/buttons
func (a_ Alert) Buttons() []Button {
	rv := objc.Send[[]Button](a_.ID, objc.Sel("buttons"))
	return rv
}

// The custom icon displayed in the alert.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/icon
func (a_ Alert) Icon() Image {
	rv := objc.Send[Image](a_.ID, objc.Sel("icon"))
	return rv
}


// SetIcon sets the value of the icon property.
// The custom icon displayed in the alert.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/icon
func (a_ Alert) SetIcon(value IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIcon:"), value)
}

// The alert’s accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/accessoryview
func (a_ Alert) AccessoryView() NSView {
	rv := objc.Send[NSView](a_.ID, objc.Sel("accessoryView"))
	return rv
}


// SetAccessoryView sets the value of the accessoryView property.
// The alert’s accessory view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/accessoryview
func (a_ Alert) SetAccessoryView(value IView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessoryView:"), value)
}

// Indicates the alert’s severity level.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/alertstyle
func (a_ Alert) AlertStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("alertStyle"))
	return rv
}


// SetAlertStyle sets the value of the alertStyle property.
// Indicates the alert’s severity level.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/alertstyle
func (a_ Alert) SetAlertStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlertStyle:"), value)
}

// The alert’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/delegate
func (a_ Alert) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The alert’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/delegate
func (a_ Alert) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// The alert’s HTML help anchor.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/helpanchor
func (a_ Alert) HelpAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("helpAnchor"))
	return rv
}


// SetHelpAnchor sets the value of the helpAnchor property.
// The alert’s HTML help anchor.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/helpanchor
func (a_ Alert) SetHelpAnchor(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHelpAnchor:"), value)
}

// The alert’s informative text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/informativetext
func (a_ Alert) InformativeText() string {
	rv := objc.Send[string](a_.ID, objc.Sel("informativeText"))
	return rv
}


// SetInformativeText sets the value of the informativeText property.
// The alert’s informative text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/informativetext
func (a_ Alert) SetInformativeText(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInformativeText:"), objc.String(value))
}

// The alert’s message text or title.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/messagetext
func (a_ Alert) MessageText() string {
	rv := objc.Send[string](a_.ID, objc.Sel("messageText"))
	return rv
}


// SetMessageText sets the value of the messageText property.
// The alert’s message text or title.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/messagetext
func (a_ Alert) SetMessageText(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMessageText:"), objc.String(value))
}

// Specifies whether the alert has a help button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showshelp
func (a_ Alert) ShowsHelp() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showsHelp"))
	return rv
}


// SetShowsHelp sets the value of the showsHelp property.
// Specifies whether the alert has a help button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showshelp
func (a_ Alert) SetShowsHelp(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowsHelp:"), value)
}

// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showssuppressionbutton
func (a_ Alert) ShowsSuppressionButton() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showsSuppressionButton"))
	return rv
}


// SetShowsSuppressionButton sets the value of the showsSuppressionButton property.
// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/showssuppressionbutton
func (a_ Alert) SetShowsSuppressionButton(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowsSuppressionButton:"), value)
}

// The alert’s suppression checkbox.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/suppressionbutton
func (a_ Alert) SuppressionButton() NSButton {
	rv := objc.Send[NSButton](a_.ID, objc.Sel("suppressionButton"))
	return rv
}


// SetSuppressionButton sets the value of the suppressionButton property.
// The alert’s suppression checkbox.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/suppressionbutton
func (a_ Alert) SetSuppressionButton(value IButton) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSuppressionButton:"), value)
}

// The app-modal panel or document-modal sheet that corresponds to the alert.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/window
func (a_ Alert) Window() NSWindow {
	rv := objc.Send[NSWindow](a_.ID, objc.Sel("window"))
	return rv
}


// SetWindow sets the value of the window property.
// The app-modal panel or document-modal sheet that corresponds to the alert.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/window
func (a_ Alert) SetWindow(value IWindow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWindow:"), value)
}


