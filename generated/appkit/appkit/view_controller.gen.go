// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewController] class.
var ViewControllerClass objc.Class

func init() {
	ViewControllerClass = objc.GetClass("NSViewController")
}

type ViewController struct {
	objc.ID
}

func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc ViewController) Alloc() ViewController {
	ret := objc.ID(ViewControllerClass).Send(objc.RegisterName("alloc"))
	return ViewController{ret}
}

// Init initializes the instance.
func (v_ ViewController) Init() ViewController {
	ret := v_.ID.Send(objc.RegisterName("init"))
	return ViewController{ret}
}
// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) ViewController {
	instance := ViewController{}.Alloc()
	sel := objc.RegisterName("initWithNibName:bundle:")
	ret := instance.ID.Send(sel, nibNameOrNil, nibBundleOrNil)
	instance = ViewController{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Dismisses a presented view controller, using the same animator that presented it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController unsafe.Pointer) {
	sel := objc.RegisterName("dismissViewController:")
	v_.ID.Send(sel, viewController)
}
// Presents another view controller using a specified, custom animator for presentation and dismissal. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer) {
	sel := objc.RegisterName("presentViewController:animator:")
	v_.ID.Send(sel, viewController, animator)
}
// Presents another view controller as a popover. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer) {
	sel := objc.RegisterName("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:")
	v_.ID.Send(sel, viewController, positioningRect, positioningView, preferredEdge, behavior)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge int, behavior unsafe.Pointer, hasFullSizeContent bool) {
	sel := objc.RegisterName("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:")
	v_.ID.Send(sel, viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(inWidget:)
func (v_ ViewController) PresentViewControllerInWidget(viewController unsafe.Pointer) {
	sel := objc.RegisterName("presentViewControllerInWidget:")
	v_.ID.Send(sel, viewController)
}
// Presents another view controller as a modal window, also known as an alert. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController unsafe.Pointer) {
	sel := objc.RegisterName("presentViewControllerAsModalWindow:")
	v_.ID.Send(sel, viewController)
}
// Presents another view controller as a sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v_ ViewController) PresentViewControllerAsSheet(viewController unsafe.Pointer) {
	sel := objc.RegisterName("presentViewControllerAsSheet:")
	v_.ID.Send(sel, viewController)
}
// Removes the called view controller from its parent view controller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/removeFromParent()
func (v_ ViewController) RemoveFromParentViewController() {
	sel := objc.RegisterName("removeFromParentViewController")
	v_.ID.Send(sel)
}

