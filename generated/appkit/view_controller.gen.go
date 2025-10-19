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
	viewControllerClass     _ViewControllerClass
	viewControllerClassOnce sync.Once
)

func getViewControllerClass() _ViewControllerClass {
	viewControllerClassOnce.Do(func() {
		viewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}
	})
	return viewControllerClass
}

type _ViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [ViewController] class.
type IViewController interface {
	IResponder
	DismissViewController(viewController unsafe.Pointer)
	PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect coregraphics.CGRect, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer, hasFullSizeContent bool)
	PresentViewControllerInWidget(viewController unsafe.Pointer)
	PresentViewControllerAsModalWindow(viewController unsafe.Pointer)
	PresentViewControllerAsSheet(viewController unsafe.Pointer)
	RemoveFromParentViewController()
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


// Returns a view controller object initialized to the nib file in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) ViewController {
	instance := getViewControllerClass().Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	rv.Autorelease()
	return rv
}


// Dismisses a presented view controller, using the same animator that presented it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissViewController:"), viewController)
}

// Presents another view controller using a specified, custom animator for presentation and dismissal.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer) {
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

// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMinimumSize
func (v_ ViewController) PreferredMinimumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("preferredMinimumSize"))
	return rv
}

