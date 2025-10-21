// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ViewController] class.
var (
	ViewControllerClass     _ViewControllerClass
	ViewControllerClassOnce sync.Once
)

func getViewControllerClass() _ViewControllerClass {
	ViewControllerClassOnce.Do(func() {
		ViewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}
	})
	return ViewControllerClass
}

type _ViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [ViewController] class.
type IViewController interface {
	IResponder
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.ID, didCommitSelector objc.SEL, contextInfo unsafe.Pointer)
	DismissController(sender objc.ID)
	DismissViewController(viewController unsafe.Pointer)
	LoadView()
	PreferredContentSizeDidChangeForViewController(viewController unsafe.Pointer)
	PresentViewControllerAnimator(viewController unsafe.Pointer, animator objc.ID)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer, hasFullSizeContent bool)
	PresentViewControllerInWidget(viewController unsafe.Pointer)
	PresentViewControllerAsModalWindow(viewController unsafe.Pointer)
	PresentViewControllerAsSheet(viewController unsafe.Pointer)
	RemoveFromParentViewController()
	ViewDidLoad()
	ViewWillTransitionToSize(newSize coregraphics.CGSize)
}

// A controller that manages a view, typically loaded from a nib file.
//
// View controller management includes: Memory management of top-level objects similar to that performed by the class, taking the same care to prevent reference cycles when controls are bound to the nib file’s owner. Declaring a generic property, to make it easy to establish bindings in the nib to an object that isn’t yet known at nib-loading time or readily available to the code that’s doing the nib loading. Implementing the key-value binding NSEditor informal protocol, so that apps using a view controller can easily make bound controls in the views commit or discard changes by the user. In macOS 10.10 and later, a view controller offers a full set of life cycle methods, allowing you to manage the content of a window in a way that is on a par with iOS view controller management. These methods, presented in order here to reflect a typical cycle, are: In addition, in macOS 10.10 and later, a view controller participates in the responder chain. You can implement action methods directly in the view controller. Corresponding actions that originate in the view controller’s view proceed up the responder chain and are handled by those methods. Prior to OS X v10.10, a typical usage pattern for loading a nib file was to subclass and override its method to call . But in macOS 10.10 and later, the method automatically looks for a nib file with the same name as the view controller. To take advantage of this behavior, name a nib file after its corresponding view controller and pass to both parameters of the method. A view controller employs lazy loading of its view: Immediately after a view controller is loaded into memory, the value of its property is . The value changes to after the method returns and just before the system calls the method. A view controller is meant to be highly reusable, such as for dynamically representing various objects. For example, the methods of the and classes take an instance as the argument, and set the property to the object that is to be shown to the user. This allows a developer to easily create new printing accessory views using bindings and the class’s key-value coding and key-value observing compliance. When the user dismisses a printing dialog, the and classes each send NSEditor messages to each accessory view controller to ensure that the user’s changes have been committed or discarded properly. The titles of the accessories are retrieved from the view controllers and shown to the user in menus that the user can choose from.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController
type ViewController struct {
	Responder
}

// ViewControllerFrom constructs a [ViewController] from an unsafe.Pointer.
//
// A controller that manages a view, typically loaded from a nib file.
func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewControllerClass) New() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ViewController) Init() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ViewController) Autorelease() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewViewController creates a new ViewController instance.
func NewViewController() ViewController {
	return getViewControllerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewViewControllerWithCoder(coder unsafe.Pointer) ViewController {
	instance := getViewControllerClass().Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns a view controller object initialized to the nib file in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) ViewController {
	instance := getViewControllerClass().Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	rv.Autorelease()
	return rv
}


// Attempt to commit any currently edited results of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/commitEditing(withDelegate:didCommit:contextInfo:)
func (v_ ViewController) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.ID, didCommitSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-3n76y
func (v_ ViewController) DismissController(sender objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissController:"), sender)
}

// Dismisses a presented view controller, using the same animator that presented it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissViewController:"), viewController)
}

// Instantiates a view from a nib file and sets the value of the property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/loadView()
func (v_ ViewController) LoadView() {
	objc.Send[objc.ID](v_.ID, objc.Sel("loadView"))
}

// Called when there is a change in value of the property of a child view controller or a presented view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSizeDidChange(for:)
func (v_ ViewController) PreferredContentSizeDidChangeForViewController(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("preferredContentSizeDidChangeForViewController:"), viewController)
}

// Presents another view controller using a specified, custom animator for presentation and dismissal.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController unsafe.Pointer, animator objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:animator:"), viewController, animator)
}

// Presents another view controller as a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer, hasFullSizeContent bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:"), viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(inWidget:)
func (v_ ViewController) PresentViewControllerInWidget(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerInWidget:"), viewController)
}

// Presents another view controller as a modal window, also known as an alert.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsModalWindow:"), viewController)
}

// Presents another view controller as a sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v_ ViewController) PresentViewControllerAsSheet(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsSheet:"), viewController)
}

// Removes the called view controller from its parent view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/removeFromParent()
func (v_ ViewController) RemoveFromParentViewController() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromParentViewController"))
}

// Called after the view controller’s view has been loaded into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidLoad()
func (v_ ViewController) ViewDidLoad() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidLoad"))
}

// For a view controller that is part of an app extension, called when its view is about to be resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillTransition(to:)
func (v_ ViewController) ViewWillTransitionToSize(newSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillTransitionToSize:"), newSize)
}

// A Boolean value indicating whether the view controller’s view is loaded into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/isViewLoaded
func (v_ ViewController) ViewLoaded() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("viewLoaded"))
	return rv
}

// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMinimumSize
func (v_ ViewController) PreferredMinimumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("preferredMinimumSize"))
	return rv
}

// The view controller’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/view
func (v_ ViewController) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The view controller’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/view
func (v_ ViewController) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setView:"), value)
}

// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (v_ ViewController) Children() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("children"))
	return rv
}


// SetChildren sets the value of the children property.
// An array of view controllers that are hierarchical children of the view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (v_ ViewController) SetChildren(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setChildren:"), value)
}

// For a view controller that is part of an app extension, the app extension context.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/extensioncontext
func (v_ ViewController) ExtensionContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("extensionContext"))
	return rv
}


// SetExtensionContext sets the value of the extensionContext property.
// For a view controller that is part of an app extension, the app extension context.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/extensioncontext
func (v_ ViewController) SetExtensionContext(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setExtensionContext:"), value)
}

// A Boolean value indicating whether the view controller’s view is loaded into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/isviewloaded
func (v_ ViewController) IsViewLoaded() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isViewLoaded"))
	return rv
}


// SetIsViewLoaded sets the value of the isViewLoaded property.
// A Boolean value indicating whether the view controller’s view is loaded into memory.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/isviewloaded
func (v_ ViewController) SetIsViewLoaded(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsViewLoaded:"), value)
}

// The nib bundle to be loaded to instantiate the receiver’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/nibbundle
func (v_ ViewController) NibBundle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("nibBundle"))
	return rv
}


// SetNibBundle sets the value of the nibBundle property.
// The nib bundle to be loaded to instantiate the receiver’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/nibbundle
func (v_ ViewController) SetNibBundle(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNibBundle:"), value)
}

// The name of the nib file to be loaded to instantiate the receiver’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/nibname
func (v_ ViewController) NibName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("nibName"))
	return rv
}


// SetNibName sets the value of the nibName property.
// The name of the nib file to be loaded to instantiate the receiver’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/nibname
func (v_ ViewController) SetNibName(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNibName:"), value)
}

// The immediate ancestor view controller of the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/parent
func (v_ ViewController) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// The immediate ancestor view controller of the view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/parent
func (v_ ViewController) SetParent(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setParent:"), value)
}

// The desired size of the view controller’s view, in screen units.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredcontentsize
func (v_ ViewController) PreferredContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("preferredContentSize"))
	return rv
}


// SetPreferredContentSize sets the value of the preferredContentSize property.
// The desired size of the view controller’s view, in screen units.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredcontentsize
func (v_ ViewController) SetPreferredContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredContentSize:"), value)
}

// For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredmaximumsize
func (v_ ViewController) PreferredMaximumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("preferredMaximumSize"))
	return rv
}


// SetPreferredMaximumSize sets the value of the preferredMaximumSize property.
// For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredmaximumsize
func (v_ ViewController) SetPreferredMaximumSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredMaximumSize:"), value)
}

// For a view controller that is part of an app extension, the preferred screen origin.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredscreenorigin
func (v_ ViewController) PreferredScreenOrigin() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("preferredScreenOrigin"))
	return rv
}


// SetPreferredScreenOrigin sets the value of the preferredScreenOrigin property.
// For a view controller that is part of an app extension, the preferred screen origin.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/preferredscreenorigin
func (v_ ViewController) SetPreferredScreenOrigin(value coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredScreenOrigin:"), value)
}

// The view controllers, if any, that are currently presented by the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/presentedviewcontrollers
func (v_ ViewController) PresentedViewControllers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("presentedViewControllers"))
	return rv
}


// SetPresentedViewControllers sets the value of the presentedViewControllers property.
// The view controllers, if any, that are currently presented by the view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/presentedviewcontrollers
func (v_ ViewController) SetPresentedViewControllers(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPresentedViewControllers:"), value)
}

// The view controller that presented the view controller or that presented its farthest ancestor view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/presentingviewcontroller
func (v_ ViewController) PresentingViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("presentingViewController"))
	return rv
}


// SetPresentingViewController sets the value of the presentingViewController property.
// The view controller that presented the view controller or that presented its farthest ancestor view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/presentingviewcontroller
func (v_ ViewController) SetPresentingViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPresentingViewController:"), value)
}

// The object whose value is presented in the receiver’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/representedobject
func (v_ ViewController) RepresentedObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("representedObject"))
	return rv
}


// SetRepresentedObject sets the value of the representedObject property.
// The object whose value is presented in the receiver’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/representedobject
func (v_ ViewController) SetRepresentedObject(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRepresentedObject:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/sourceitemview
func (v_ ViewController) SourceItemView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("sourceItemView"))
	return rv
}


// SetSourceItemView sets the value of the sourceItemView property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/sourceitemview
func (v_ ViewController) SetSourceItemView(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSourceItemView:"), value)
}

// The storyboard from which the view controller was loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/storyboard
func (v_ ViewController) Storyboard() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("storyboard"))
	return rv
}


// SetStoryboard sets the value of the storyboard property.
// The storyboard from which the view controller was loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/storyboard
func (v_ ViewController) SetStoryboard(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStoryboard:"), value)
}

// The localized title of the receiver’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/title
func (v_ ViewController) Title() string {
	rv := objc.Send[string](v_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The localized title of the receiver’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/title
func (v_ ViewController) SetTitle(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTitle:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/viewifloaded
func (v_ ViewController) ViewIfLoaded() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("viewIfLoaded"))
	return rv
}


// SetViewIfLoaded sets the value of the viewIfLoaded property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/viewifloaded
func (v_ ViewController) SetViewIfLoaded(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setViewIfLoaded:"), value)
}


