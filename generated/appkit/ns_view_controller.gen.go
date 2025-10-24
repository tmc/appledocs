// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSViewController */


/* debug [class_header]: Header for NSViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ViewController */
// An interface definition for the [ViewController] class.
type IViewController interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for ViewController */
	// properties:
	ChildViewControllers() []ViewController
	SetChildViewControllers(value []ViewController)
	ExtensionContext() objc.IObject /* cross-framework: NSExtensionContext */
	ViewLoaded() bool
	NibBundle() foundation.Bundle
	NibName() NibName /* typedef */
	ParentViewController() IViewController
	PreferredContentSize() Size /* not a class type */
	SetPreferredContentSize(value Size /* not a class type */)
	PreferredMaximumSize() Size /* not a class type */
	PreferredMinimumSize() Size /* not a class type */
	PreferredScreenOrigin() vision.Point
	SetPreferredScreenOrigin(value vision.Point)
	PresentedViewControllers() []ViewController
	PresentingViewController() IViewController
	RepresentedObject() objc.ID
	SetRepresentedObject(value objc.ID)
	SourceItemView() IView
	SetSourceItemView(value IView)
	Storyboard() IStoryboard
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	View() IView
	SetView(value IView)
	ViewIfLoaded() IView
	Children() IViewController
	SetChildren(value IViewController)
	IsViewLoaded() bool
	SetIsViewLoaded(value bool)
	Parent() IViewController
	SetParent(value IViewController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ViewController */
	// methods:
	AddChildViewController(childViewController IViewController)
	CommitEditing() bool
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject)
	DiscardEditing()
	DismissController(sender objc.IObject)
	DismissViewController(viewController IViewController)
	InsertChildViewControllerAtIndex(childViewController IViewController, index int)
	LoadView()
	LoadViewIfNeeded()
	PreferredContentSizeDidChangeForViewController(viewController IViewController)
	PresentViewControllerAnimator(viewController IViewController, animator unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController IViewController, positioningRect Rect /* not a class type */, positioningView IView, preferredEdge RectEdge /* not a class type */, behavior PopoverBehavior /* not a class type */)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController IViewController, positioningRect Rect /* not a class type */, positioningView IView, preferredEdge RectEdge /* not a class type */, behavior PopoverBehavior /* not a class type */, hasFullSizeContent bool)
	PresentViewControllerAsModalWindow(viewController IViewController)
	PresentViewControllerAsSheet(viewController IViewController)
	RemoveChildViewControllerAtIndex(index int)
	RemoveFromParentViewController()
	TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion unsafe.Pointer)
	UpdateViewConstraints()
	ViewDidAppear()
	ViewDidDisappear()
	ViewDidLayout()
	ViewDidLoad()
	ViewWillAppear()
	ViewWillDisappear()
	ViewWillLayout()
	ViewWillTransitionToSize(newSize Size /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ViewController */
// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ViewController */
// A controller that manages a view, typically loaded from a nib file.
//
// View controller management includes: Memory management of top-level objects similar to that performed by the class, taking the same care to prevent reference cycles when controls are bound to the nib file’s owner. Declaring a generic property, to make it easy to establish bindings in the nib to an object that isn’t yet known at nib-loading time or readily available to the code that’s doing the nib loading. Implementing the key-value binding NSEditor informal protocol, so that apps using a view controller can easily make bound controls in the views commit or discard changes by the user. In macOS 10.10 and later, a view controller offers a full set of life cycle methods, allowing you to manage the content of a window in a way that is on a par with iOS view controller management. These methods, presented in order here to reflect a typical cycle, are: In addition, in macOS 10.10 and later, a view controller participates in the responder chain. You can implement action methods directly in the view controller. Corresponding actions that originate in the view controller’s view proceed up the responder chain and are handled by those methods. Prior to OS X v10.10, a typical usage pattern for loading a nib file was to subclass and override its method to call . But in macOS 10.10 and later, the method automatically looks for a nib file with the same name as the view controller. To take advantage of this behavior, name a nib file after its corresponding view controller and pass to both parameters of the method. A view controller employs lazy loading of its view: Immediately after a view controller is loaded into memory, the value of its property is . The value changes to after the method returns and just before the system calls the method. A view controller is meant to be highly reusable, such as for dynamically representing various objects. For example, the methods of the and classes take an instance as the argument, and set the property to the object that is to be shown to the user. This allows a developer to easily create new printing accessory views using bindings and the class’s key-value coding and key-value observing compliance. When the user dismisses a printing dialog, the and classes each send NSEditor messages to each accessory view controller to ensure that the user’s changes have been committed or discarded properly. The titles of the accessories are retrieved from the view controllers and shown to the user in menus that the user can choose from.


// A controller that manages a view, typically loaded from a nib file.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewViewControllerWithCoder(coder foundation.Coder) ViewController {
	instance := getViewControllerClass().Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewControllerWithCoder */


// Returns a view controller object initialized to the nib file in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil NibName /* typedef */, nibBundleOrNil foundation.Bundle) ViewController {
	instance := getViewControllerClass().Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewControllerWithNibNameBundle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ViewController */

// A convenience method for adding a child view controller at the end of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/addChild(_:)
func (v_ ViewController) AddChildViewController(childViewController IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addChildViewController:"), childViewController)
}/* debug [instance_methods/method]: AddChildViewController */


// Returns whether the receiver was able to commit any pending edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/commitEditing()
func (v_ ViewController) CommitEditing() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("commitEditing"))
	return rv
}/* debug [instance_methods/method]: CommitEditing */


// Attempt to commit any currently edited results of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/commitEditing(withDelegate:didCommit:contextInfo:)
func (v_ ViewController) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}/* debug [instance_methods/method]: CommitEditingWithDelegateDidCommitSelectorContextInfo */


// Causes the receiver to discard any changes, restoring the previous values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/discardEditing()
func (v_ ViewController) DiscardEditing() {
	objc.Send[objc.ID](v_.ID, objc.Sel("discardEditing"))
}/* debug [instance_methods/method]: DiscardEditing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-3n76y
func (v_ ViewController) DismissController(sender objc.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissController:"), sender)
}/* debug [instance_methods/method]: DismissController */


// Dismisses a presented view controller, using the same animator that presented it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissViewController:"), viewController)
}/* debug [instance_methods/method]: DismissViewController */


// Inserts a specified child view controller into the array at a specified position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/insertChild(_:at:)
func (v_ ViewController) InsertChildViewControllerAtIndex(childViewController IViewController, index int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("insertChildViewController:atIndex:"), childViewController, index)
}/* debug [instance_methods/method]: InsertChildViewControllerAtIndex */


// Instantiates a view from a nib file and sets the value of the property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/loadView()
func (v_ ViewController) LoadView() {
	objc.Send[objc.ID](v_.ID, objc.Sel("loadView"))
}/* debug [instance_methods/method]: LoadView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/loadViewIfNeeded()
func (v_ ViewController) LoadViewIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("loadViewIfNeeded"))
}/* debug [instance_methods/method]: LoadViewIfNeeded */


// Called when there is a change in value of the property of a child view controller or a presented view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSizeDidChange(for:)
func (v_ ViewController) PreferredContentSizeDidChangeForViewController(viewController IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("preferredContentSizeDidChangeForViewController:"), viewController)
}/* debug [instance_methods/method]: PreferredContentSizeDidChangeForViewController */


// Presents another view controller using a specified, custom animator for presentation and dismissal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController IViewController, animator unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:animator:"), viewController, animator)
}/* debug [instance_methods/method]: PresentViewControllerAnimator */


// Presents another view controller as a popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController IViewController, positioningRect Rect /* not a class type */, positioningView IView, preferredEdge RectEdge /* not a class type */, behavior PopoverBehavior /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}/* debug [instance_methods/method]: PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController IViewController, positioningRect Rect /* not a class type */, positioningView IView, preferredEdge RectEdge /* not a class type */, behavior PopoverBehavior /* not a class type */, hasFullSizeContent bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:"), viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}/* debug [instance_methods/method]: PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent */


// Presents another view controller as a modal window, also known as an alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsModalWindow:"), viewController)
}/* debug [instance_methods/method]: PresentViewControllerAsModalWindow */


// Presents another view controller as a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v_ ViewController) PresentViewControllerAsSheet(viewController IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsSheet:"), viewController)
}/* debug [instance_methods/method]: PresentViewControllerAsSheet */


// Removes a specified child controller from the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/removeChild(at:)
func (v_ ViewController) RemoveChildViewControllerAtIndex(index int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeChildViewControllerAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveChildViewControllerAtIndex */


// Removes the called view controller from its parent view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/removeFromParent()
func (v_ ViewController) RemoveFromParentViewController() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromParentViewController"))
}/* debug [instance_methods/method]: RemoveFromParentViewController */


// Performs a transition between two sibling child view controllers of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/transition(from:to:options:completionHandler:)
func (v_ ViewController) TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("transitionFromViewController:toViewController:options:completionHandler:"), fromViewController, toViewController, options, completion)
}/* debug [instance_methods/method]: TransitionFromViewControllerToViewControllerOptionsCompletionHandler */


// Called during Auto Layout constraint updating to enable the view controller to mediate the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/updateViewConstraints()
func (v_ ViewController) UpdateViewConstraints() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateViewConstraints"))
}/* debug [instance_methods/method]: UpdateViewConstraints */


// Called when the view controller’s view is fully transitioned onto the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidAppear()
func (v_ ViewController) ViewDidAppear() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidAppear"))
}/* debug [instance_methods/method]: ViewDidAppear */


// Called after the view controller’s view is removed from the view hierarchy in a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidDisappear()
func (v_ ViewController) ViewDidDisappear() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidDisappear"))
}/* debug [instance_methods/method]: ViewDidDisappear */


// Called immediately after the method of the view controller’s view is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidLayout()
func (v_ ViewController) ViewDidLayout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidLayout"))
}/* debug [instance_methods/method]: ViewDidLayout */


// Called after the view controller’s view has been loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidLoad()
func (v_ ViewController) ViewDidLoad() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidLoad"))
}/* debug [instance_methods/method]: ViewDidLoad */


// Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillAppear()
func (v_ ViewController) ViewWillAppear() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillAppear"))
}/* debug [instance_methods/method]: ViewWillAppear */


// Called when the view controller’s view is about to be removed from the view hierarchy in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillDisappear()
func (v_ ViewController) ViewWillDisappear() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillDisappear"))
}/* debug [instance_methods/method]: ViewWillDisappear */


// Called just before the method of the view controller’s view is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillLayout()
func (v_ ViewController) ViewWillLayout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillLayout"))
}/* debug [instance_methods/method]: ViewWillLayout */


// For a view controller that is part of an app extension, called when its view is about to be resized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillTransition(to:)
func (v_ ViewController) ViewWillTransitionToSize(newSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillTransitionToSize:"), newSize)
}/* debug [instance_methods/method]: ViewWillTransitionToSize */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ViewController */

// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/children
func (v_ ViewController) ChildViewControllers() []ViewController {
	rv := objc.Send[[]ViewController](v_.ID, objc.Sel("childViewControllers"))
	return rv
}/* debug [instance_properties/getter]: childViewControllers */


// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/children
func (v_ ViewController) SetChildViewControllers(value []ViewController) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setChildViewControllers:"), nsArray)
}/* debug [instance_properties/setter]: childViewControllers */


// For a view controller that is part of an app extension, the app extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/extensionContext
func (v_ ViewController) ExtensionContext() objc.IObject /* cross-framework: NSExtensionContext */ {
	rv := objc.Send[foundation.NSExtensionContext](v_.ID, objc.Sel("extensionContext"))
	return rv
}/* debug [instance_properties/getter]: extensionContext */


// A Boolean value indicating whether the view controller’s view is loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/isViewLoaded
func (v_ ViewController) ViewLoaded() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("viewLoaded"))
	return rv
}/* debug [instance_properties/getter]: viewLoaded */


// The nib bundle to be loaded to instantiate the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/nibBundle
func (v_ ViewController) NibBundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](v_.ID, objc.Sel("nibBundle"))
	return rv
}/* debug [instance_properties/getter]: nibBundle */


// The name of the nib file to be loaded to instantiate the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/nibName
func (v_ ViewController) NibName() NibName /* typedef */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("nibName"))
	return rv
}/* debug [instance_properties/getter]: nibName */


// The immediate ancestor view controller of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/parent
func (v_ ViewController) ParentViewController() IViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("parentViewController"))
	return rv
}/* debug [instance_properties/getter]: parentViewController */


// The desired size of the view controller’s view, in screen units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSize
func (v_ ViewController) PreferredContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("preferredContentSize"))
	return rv
}/* debug [instance_properties/getter]: preferredContentSize */


// The desired size of the view controller’s view, in screen units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSize
func (v_ ViewController) SetPreferredContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredContentSize:"), value)
}/* debug [instance_properties/setter]: preferredContentSize */


// For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMaximumSize
func (v_ ViewController) PreferredMaximumSize() Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("preferredMaximumSize"))
	return rv
}/* debug [instance_properties/getter]: preferredMaximumSize */


// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMinimumSize
func (v_ ViewController) PreferredMinimumSize() Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("preferredMinimumSize"))
	return rv
}/* debug [instance_properties/getter]: preferredMinimumSize */


// For a view controller that is part of an app extension, the preferred screen origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredScreenOrigin
func (v_ ViewController) PreferredScreenOrigin() vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("preferredScreenOrigin"))
	return rv
}/* debug [instance_properties/getter]: preferredScreenOrigin */


// For a view controller that is part of an app extension, the preferred screen origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredScreenOrigin
func (v_ ViewController) SetPreferredScreenOrigin(value vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredScreenOrigin:"), value)
}/* debug [instance_properties/setter]: preferredScreenOrigin */


// The view controllers, if any, that are currently presented by the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentedViewControllers
func (v_ ViewController) PresentedViewControllers() []ViewController {
	rv := objc.Send[[]ViewController](v_.ID, objc.Sel("presentedViewControllers"))
	return rv
}/* debug [instance_properties/getter]: presentedViewControllers */


// The view controller that presented the view controller or that presented its farthest ancestor view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentingViewController
func (v_ ViewController) PresentingViewController() IViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("presentingViewController"))
	return rv
}/* debug [instance_properties/getter]: presentingViewController */


// The object whose value is presented in the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/representedObject
func (v_ ViewController) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("representedObject"))
	return rv
}/* debug [instance_properties/getter]: representedObject */


// The object whose value is presented in the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/representedObject
func (v_ ViewController) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRepresentedObject:"), value)
}/* debug [instance_properties/setter]: representedObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/sourceItemView
func (v_ ViewController) SourceItemView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("sourceItemView"))
	return rv
}/* debug [instance_properties/getter]: sourceItemView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/sourceItemView
func (v_ ViewController) SetSourceItemView(value IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSourceItemView:"), value)
}/* debug [instance_properties/setter]: sourceItemView */


// The storyboard from which the view controller was loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/storyboard
func (v_ ViewController) Storyboard() IStoryboard {
	rv := objc.Send[Storyboard](v_.ID, objc.Sel("storyboard"))
	return rv
}/* debug [instance_properties/getter]: storyboard */


// The localized title of the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/title
func (v_ ViewController) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The localized title of the receiver’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/title
func (v_ ViewController) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/view
func (v_ ViewController) View() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/view
func (v_ ViewController) SetView(value IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setView:"), value)
}/* debug [instance_properties/setter]: view */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/viewIfLoaded
func (v_ ViewController) ViewIfLoaded() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("viewIfLoaded"))
	return rv
}/* debug [instance_properties/getter]: viewIfLoaded */


// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (v_ ViewController) Children() IViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (v_ ViewController) SetChildren(value IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setChildren:"), value)
}/* debug [instance_properties/setter]: children */


// A Boolean value indicating whether the view controller’s view is loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/isviewloaded
func (v_ ViewController) IsViewLoaded() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isViewLoaded"))
	return rv
}/* debug [instance_properties/getter]: isViewLoaded */


// A Boolean value indicating whether the view controller’s view is loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/isviewloaded
func (v_ ViewController) SetIsViewLoaded(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsViewLoaded:"), value)
}/* debug [instance_properties/setter]: isViewLoaded */


// The immediate ancestor view controller of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/parent
func (v_ ViewController) Parent() IViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// The immediate ancestor view controller of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/parent
func (v_ ViewController) SetParent(value IViewController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSViewController */


