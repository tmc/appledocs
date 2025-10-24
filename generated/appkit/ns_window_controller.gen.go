// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSWindowController */


/* debug [class_header]: Header for NSWindowController */
// The class instance for the [WindowController] class.
var (
	WindowControllerClass     _WindowControllerClass
	WindowControllerClassOnce sync.Once
)

func getWindowControllerClass() _WindowControllerClass {
	WindowControllerClassOnce.Do(func() {
		WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}
	})
	return WindowControllerClass
}

type _WindowControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WindowController */
// An interface definition for the [WindowController] class.
type IWindowController interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for WindowController */
	// properties:
	ContentViewController() IViewController
	SetContentViewController(value IViewController)
	Document() objc.ID
	SetDocument(value objc.ID)
	WindowLoaded() bool
	Owner() objc.ID
	PreviewRepresentableActivityItems() []objc.ID
	SetPreviewRepresentableActivityItems(value []objc.ID)
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)
	Storyboard() IStoryboard
	Window() IWindow
	SetWindow(value IWindow)
	WindowFrameAutosaveName() WindowFrameAutosaveName /* typedef */
	SetWindowFrameAutosaveName(value WindowFrameAutosaveName /* typedef */)
	WindowNibName() NibName /* typedef */
	WindowNibPath() objc.IObject /* cross-framework: NSString */
	IsWindowLoaded() bool
	SetIsWindowLoaded(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WindowController */
	// methods:
	Close()
	DismissController(sender objc.IObject)
	LoadWindow()
	SetDocumentEdited(dirtyFlag bool)
	ShowWindow(sender objc.IObject)
	SynchronizeWindowTitleWithDocumentName()
	WindowDidLoad()
	WindowTitleForDocumentDisplayName(displayName objc.IObject /* cross-framework: NSString */) foundation.String
	WindowWillLoad()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WindowController */
// Alloc allocates a new instance without initialization.
func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WindowControllerClass) New() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowController) Init() WindowController {
	rv := objc.Send[WindowController](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowController) Autorelease() WindowController {
	rv := objc.Send[WindowController](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowController creates a new WindowController instance.
func NewWindowController() WindowController {
	return getWindowControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WindowController */
// A controller that manages a window, usually a window stored in a nib file.
//
// Managing a window entails: Loading and displaying the window Closing the window when appropriate Customizing the window’s title Storing the window’s frame (size and location) in the defaults database Cascading the window in relation to other document windows of the app A window controller can manage a window by itself or as a role player in AppKit’s document-based architecture, which also includes and objects. In this architecture, a window controller is created and managed by a “document” (an instance of an subclass) and, in turn, keeps a reference to the document. The relationship between a window controller and a nib file is important. Although a window controller can manage a programmatically created window, it usually manages a window in a nib file. The nib file can contain other top-level objects, including other windows, but the window controller’s responsibility is this primary window. The window controller is usually the owner of the nib file, even when it is part of a document-based app. Regardless of who is the file’s owner, the window controller is responsible for freeing all top-level objects in the nib file it loads. For simple documents—that is, documents with only one nib file containing a window—you need to do little directly with ; AppKit creates one for you. However, if the default window controller is not sufficient, you can create a custom subclass of . For documents with multiple windows or panels, your document must create separate instances of (or of custom subclasses of ), one for each window or panel. An example is a CAD app that has different windows for side, top, and front views of drawn objects. What you do in your subclass determines whether the default or separately created and configured objects are used.


// A controller that manages a window, usually a window stored in a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController
type WindowController struct {
	Responder
}

// WindowControllerFrom constructs a [WindowController] from an unsafe.Pointer.
//
// A controller that manages a window, usually a window stored in a nib file.
func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: ResponderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WindowController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(coder:)
func NewWindowControllerWithCoder(coder foundation.Coder) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowControllerWithCoder */


// Returns a window controller initialized with a given window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(window:)
func NewWindowControllerWithWindow(window IWindow) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithWindow:"), window)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowControllerWithWindow */


// Returns a window controller initialized with a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:)
func NewWindowControllerWithWindowNibName(windowNibName NibName /* typedef */) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithWindowNibName:"), windowNibName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowControllerWithWindowNibName */


// Returns a window controller initialized with a nib file and a specified owner for that nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:owner:)
func NewWindowControllerWithWindowNibNameOwner(windowNibName NibName /* typedef */, owner objc.IObject) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithWindowNibName:owner:"), windowNibName, owner)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowControllerWithWindowNibNameOwner */


// Returns a window controller initialized with a nib file at an absolute path and a specified owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibPath:owner:)
func NewWindowControllerWithWindowNibPathOwner(windowNibPath objc.IObject /* cross-framework: NSString */, owner objc.IObject) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithWindowNibPath:owner:"), windowNibPath, owner)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowControllerWithWindowNibPathOwner */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WindowController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WindowController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WindowController */

// Closes the window if it was loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/close()
func (w_ WindowController) Close() {
	objc.Send[objc.ID](w_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Dismisses the window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/dismissController(_:)
func (w_ WindowController) DismissController(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("dismissController:"), sender)
}/* debug [instance_methods/method]: DismissController */


// Loads the receiver’s window from the nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/loadWindow()
func (w_ WindowController) LoadWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadWindow"))
}/* debug [instance_methods/method]: LoadWindow */


// Sets the document edited flag for the window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/setDocumentEdited(_:)
func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDocumentEdited:"), dirtyFlag)
}/* debug [instance_methods/method]: SetDocumentEdited */


// Displays the window associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/showWindow(_:)
func (w_ WindowController) ShowWindow(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("showWindow:"), sender)
}/* debug [instance_methods/method]: ShowWindow */


// Synchronizes the displayed window title and the represented filename with the information in the associated document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/synchronizeWindowTitleWithDocumentName()
func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.Send[objc.ID](w_.ID, objc.Sel("synchronizeWindowTitleWithDocumentName"))
}/* debug [instance_methods/method]: SynchronizeWindowTitleWithDocumentName */


// Sent after the window owned by the receiver has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowDidLoad()
func (w_ WindowController) WindowDidLoad() {
	objc.Send[objc.ID](w_.ID, objc.Sel("windowDidLoad"))
}/* debug [instance_methods/method]: WindowDidLoad */


// Returns the window title to be used for a given document display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowTitle(forDocumentDisplayName:)
func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("windowTitleForDocumentDisplayName:"), displayName)
	return rv
}/* debug [instance_methods/method]: WindowTitleForDocumentDisplayName */


// Sent before the window owned by the receiver is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowWillLoad()
func (w_ WindowController) WindowWillLoad() {
	objc.Send[objc.ID](w_.ID, objc.Sel("windowWillLoad"))
}/* debug [instance_methods/method]: WindowWillLoad */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WindowController */

// The view controller for the window’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/contentViewController
func (w_ WindowController) ContentViewController() IViewController {
	rv := objc.Send[ViewController](w_.ID, objc.Sel("contentViewController"))
	return rv
}/* debug [instance_properties/getter]: contentViewController */


// The view controller for the window’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/contentViewController
func (w_ WindowController) SetContentViewController(value IViewController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentViewController:"), value)
}/* debug [instance_properties/setter]: contentViewController */


// The document associated with the window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/document
func (w_ WindowController) Document() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// The document associated with the window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/document
func (w_ WindowController) SetDocument(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDocument:"), value)
}/* debug [instance_properties/setter]: document */


// A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/isWindowLoaded
func (w_ WindowController) WindowLoaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("windowLoaded"))
	return rv
}/* debug [instance_properties/getter]: windowLoaded */


// The owner of the nib file containing the window managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/owner
func (w_ WindowController) Owner() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("owner"))
	return rv
}/* debug [instance_properties/getter]: owner */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/previewRepresentableActivityItems
func (w_ WindowController) PreviewRepresentableActivityItems() []objc.ID {
	rv := objc.Send[[]objc.ID](w_.ID, objc.Sel("previewRepresentableActivityItems"))
	return rv
}/* debug [instance_properties/getter]: previewRepresentableActivityItems */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/previewRepresentableActivityItems
func (w_ WindowController) SetPreviewRepresentableActivityItems(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreviewRepresentableActivityItems:"), nsArray)
}/* debug [instance_properties/setter]: previewRepresentableActivityItems */


// A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCascadeWindows
func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldCascadeWindows"))
	return rv
}/* debug [instance_properties/getter]: shouldCascadeWindows */


// A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCascadeWindows
func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldCascadeWindows:"), value)
}/* debug [instance_properties/setter]: shouldCascadeWindows */


// A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCloseDocument
func (w_ WindowController) ShouldCloseDocument() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldCloseDocument"))
	return rv
}/* debug [instance_properties/getter]: shouldCloseDocument */


// A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCloseDocument
func (w_ WindowController) SetShouldCloseDocument(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldCloseDocument:"), value)
}/* debug [instance_properties/setter]: shouldCloseDocument */


// The storyboard file from which the window controller was loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/storyboard
func (w_ WindowController) Storyboard() IStoryboard {
	rv := objc.Send[Storyboard](w_.ID, objc.Sel("storyboard"))
	return rv
}/* debug [instance_properties/getter]: storyboard */


// The window owned by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/window
func (w_ WindowController) Window() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("window"))
	return rv
}/* debug [instance_properties/getter]: window */


// The window owned by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/window
func (w_ WindowController) SetWindow(value IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindow:"), value)
}/* debug [instance_properties/setter]: window */


// The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowFrameAutosaveName
func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("windowFrameAutosaveName"))
	return rv
}/* debug [instance_properties/getter]: windowFrameAutosaveName */


// The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowFrameAutosaveName
func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowFrameAutosaveName:"), value)
}/* debug [instance_properties/setter]: windowFrameAutosaveName */


// The name of the nib file that stores the window associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowNibName
func (w_ WindowController) WindowNibName() NibName /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("windowNibName"))
	return rv
}/* debug [instance_properties/getter]: windowNibName */


// The full path of the nib file that stores the window associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/windowNibPath
func (w_ WindowController) WindowNibPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("windowNibPath"))
	return rv
}/* debug [instance_properties/getter]: windowNibPath */


// A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/iswindowloaded
func (w_ WindowController) IsWindowLoaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWindowLoaded"))
	return rv
}/* debug [instance_properties/getter]: isWindowLoaded */


// A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/iswindowloaded
func (w_ WindowController) SetIsWindowLoaded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsWindowLoaded:"), value)
}/* debug [instance_properties/setter]: isWindowLoaded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWindowController */


