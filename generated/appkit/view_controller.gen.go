// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ViewController] class.
var viewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}

type _ViewControllerClass struct {
	class objc.Class
}

// A controller that manages a view, typically loaded from a nib file. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return viewControllerClass.New()
}
// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) ViewController {
	instance := viewControllerClass.Alloc()
	rv := objc.Send[ViewController](instance.ID, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	rv.Autorelease()
	return rv
}


// Dismisses a presented view controller, using the same animator that presented it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dismissViewController:"), viewController)
}
// Presents another view controller using a specified, custom animator for presentation and dismissal. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:animator:"), viewController, animator)
}
// Presents another view controller as a popover. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer, hasFullSizeContent bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:"), viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/present(inWidget:)
func (v_ ViewController) PresentViewControllerInWidget(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerInWidget:"), viewController)
}
// Presents another view controller as a modal window, also known as an alert. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsModalWindow:"), viewController)
}
// Presents another view controller as a sheet. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v_ ViewController) PresentViewControllerAsSheet(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentViewControllerAsSheet:"), viewController)
}
// Removes the called view controller from its parent view controller. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/removeFromParent()
func (v_ ViewController) RemoveFromParentViewController() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromParentViewController"))
}

