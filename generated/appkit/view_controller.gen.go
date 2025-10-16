
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewController] class.
var ViewControllerClass _ViewControllerClass

func init() {
	ViewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}
}

type _ViewControllerClass struct {
	objc.Class
}

// An interface definition for the [ViewController] class.
type IViewController interface {
	ID() objc.ID
	DismissViewController(viewController unsafe.Pointer)
	InitWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) unsafe.Pointer
	PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer)
	PresentViewControllerAsModalWindow(viewController unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge unsafe.Pointer, behavior unsafe.Pointer)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge unsafe.Pointer, behavior unsafe.Pointer, hasFullSizeContent bool)
	PresentViewControllerAsSheet(viewController unsafe.Pointer)
	PresentViewControllerInWidget(viewController unsafe.Pointer)
	RemoveFromParentViewController()
}

type ViewController struct {
	id objc.ID
}

func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ ViewController) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _ViewControllerClass) New() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewViewController creates and returns a new initialized instance.
func NewViewController() ViewController {
	return ViewControllerClass.New()
}

// Init initializes the instance.
func (v_ ViewController) Init() ViewController {
	rv := objc.Send[ViewController](v_.ID(), selInit)
	return rv
}
// Dismisses a presented view controller, using the same animator that presented it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v_ ViewController) DismissViewController(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("dismissViewController:"), viewController)
}
// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/init(nibName:bundle:)
func (v_ ViewController) InitWithNibNameBundle(nibNameOrNil unsafe.Pointer, nibBundleOrNil unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	return rv
}
// Presents another view controller using a specified, custom animator for presentation and dismissal. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:animator:)
func (v_ ViewController) PresentViewControllerAnimator(viewController unsafe.Pointer, animator unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewController:animator:"), viewController, animator)
}
// Presents another view controller as a popover. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge unsafe.Pointer, behavior unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v_ ViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController unsafe.Pointer, positioningRect unsafe.Pointer, positioningView unsafe.Pointer, preferredEdge unsafe.Pointer, behavior unsafe.Pointer, hasFullSizeContent bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:"), viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/present(inWidget:)
func (v_ ViewController) PresentViewControllerInWidget(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewControllerInWidget:"), viewController)
}
// Presents another view controller as a modal window, also known as an alert. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewControllerAsModalWindow:"), viewController)
}
// Presents another view controller as a sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v_ ViewController) PresentViewControllerAsSheet(viewController unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("presentViewControllerAsSheet:"), viewController)
}
// Removes the called view controller from its parent view controller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/removeFromParent()
func (v_ ViewController) RemoveFromParentViewController() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeFromParentViewController"))
}
// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewController/preferredMinimumSize
func (v_ ViewController) PreferredMinimumSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("preferredMinimumSize"))
	return rv
}
