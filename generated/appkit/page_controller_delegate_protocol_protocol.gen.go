// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPageControllerDelegate is the NSPageControllerDelegate protocol interface.
//
// The   protocol allows you to customize the behavior of instances of the NSPageController class.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPageControllerDelegate
type PPageControllerDelegate interface {
	// Optional methods
	PageControllerDidTransitionToObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject)
	HasPageControllerDidTransitionToObject() bool
	PageControllerFrameForObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) Rect
	HasPageControllerFrameForObject() bool
	PageControllerIdentifierForObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) PageControllerObjectIdentifier
	HasPageControllerIdentifierForObject() bool
	PageControllerPrepareViewControllerWithObject(pageController objc.IObject /* cross-framework: PageController */, viewController IViewController, object objc.IObject)
	HasPageControllerPrepareViewControllerWithObject() bool
	PageControllerViewControllerForIdentifier(pageController objc.IObject /* cross-framework: PageController */, identifier PageControllerObjectIdentifier /* typedef */) ViewController
	HasPageControllerViewControllerForIdentifier() bool
	PageControllerDidEndLiveTransition(pageController objc.IObject /* cross-framework: PageController */)
	HasPageControllerDidEndLiveTransition() bool
	PageControllerWillStartLiveTransition(pageController objc.IObject /* cross-framework: PageController */)
	HasPageControllerWillStartLiveTransition() bool
}

// PageControllerDelegate is a delegate implementation builder for the PPageControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PageControllerDelegate struct {
	_PageControllerDidTransitionToObject func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject)
	_PageControllerFrameForObject func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) Rect
	_PageControllerIdentifierForObject func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) PageControllerObjectIdentifier
	_PageControllerPrepareViewControllerWithObject func(pageController objc.IObject /* cross-framework: PageController */, viewController IViewController, object objc.IObject)
	_PageControllerViewControllerForIdentifier func(pageController objc.IObject /* cross-framework: PageController */, identifier PageControllerObjectIdentifier /* typedef */) ViewController
	_PageControllerDidEndLiveTransition func(pageController objc.IObject /* cross-framework: PageController */)
	_PageControllerWillStartLiveTransition func(pageController objc.IObject /* cross-framework: PageController */)
}

// SetPageControllerDidTransitionToObject sets the handler for the PageControllerDidTransitionToObject delegate method.
//
// This message is sent when any page transition is completed.
func (d *PageControllerDelegate) SetPageControllerDidTransitionToObject(f func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject)) {
	d._PageControllerDidTransitionToObject = f
}

// SetPageControllerFrameForObject sets the handler for the PageControllerFrameForObject delegate method.
//
// Returns the frame appropriate for displaying the specified object.
func (d *PageControllerDelegate) SetPageControllerFrameForObject(f func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) Rect) {
	d._PageControllerFrameForObject = f
}

// SetPageControllerIdentifierForObject sets the handler for the PageControllerIdentifierForObject delegate method.
//
// Return the identifier of the view controller that owns a view to display the object.
func (d *PageControllerDelegate) SetPageControllerIdentifierForObject(f func(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) PageControllerObjectIdentifier) {
	d._PageControllerIdentifierForObject = f
}

// SetPageControllerPrepareViewControllerWithObject sets the handler for the PageControllerPrepareViewControllerWithObject delegate method.
//
// Prepare the view controller and it’s view for drawing.
func (d *PageControllerDelegate) SetPageControllerPrepareViewControllerWithObject(f func(pageController objc.IObject /* cross-framework: PageController */, viewController IViewController, object objc.IObject)) {
	d._PageControllerPrepareViewControllerWithObject = f
}

// SetPageControllerViewControllerForIdentifier sets the handler for the PageControllerViewControllerForIdentifier delegate method.
//
// Returns a view controller the page controller uses for managing the specified identifier.
func (d *PageControllerDelegate) SetPageControllerViewControllerForIdentifier(f func(pageController objc.IObject /* cross-framework: PageController */, identifier PageControllerObjectIdentifier /* typedef */) ViewController) {
	d._PageControllerViewControllerForIdentifier = f
}

// SetPageControllerDidEndLiveTransition sets the handler for the PageControllerDidEndLiveTransition delegate method.
//
// This message is sent when a transition animation completes.
func (d *PageControllerDelegate) SetPageControllerDidEndLiveTransition(f func(pageController objc.IObject /* cross-framework: PageController */)) {
	d._PageControllerDidEndLiveTransition = f
}

// SetPageControllerWillStartLiveTransition sets the handler for the PageControllerWillStartLiveTransition delegate method.
//
// This message is sent when the user begins a transition.
func (d *PageControllerDelegate) SetPageControllerWillStartLiveTransition(f func(pageController objc.IObject /* cross-framework: PageController */)) {
	d._PageControllerWillStartLiveTransition = f
}

// PageControllerDidTransitionToObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerDidTransitionToObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) {
	if d._PageControllerDidTransitionToObject != nil {
		d._PageControllerDidTransitionToObject(pageController, object)
	}
}

// HasPageControllerDidTransitionToObject returns true if a handler for PageControllerDidTransitionToObject has been set.
func (d *PageControllerDelegate) HasPageControllerDidTransitionToObject() bool {
	return d._PageControllerDidTransitionToObject != nil
}

// PageControllerFrameForObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerFrameForObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) Rect {
	if d._PageControllerFrameForObject != nil {
		return d._PageControllerFrameForObject(pageController, object)
	}
	var zero Rect
	return zero
}

// HasPageControllerFrameForObject returns true if a handler for PageControllerFrameForObject has been set.
func (d *PageControllerDelegate) HasPageControllerFrameForObject() bool {
	return d._PageControllerFrameForObject != nil
}

// PageControllerIdentifierForObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerIdentifierForObject(pageController objc.IObject /* cross-framework: PageController */, object objc.IObject) PageControllerObjectIdentifier {
	if d._PageControllerIdentifierForObject != nil {
		return d._PageControllerIdentifierForObject(pageController, object)
	}
	var zero PageControllerObjectIdentifier
	return zero
}

// HasPageControllerIdentifierForObject returns true if a handler for PageControllerIdentifierForObject has been set.
func (d *PageControllerDelegate) HasPageControllerIdentifierForObject() bool {
	return d._PageControllerIdentifierForObject != nil
}

// PageControllerPrepareViewControllerWithObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerPrepareViewControllerWithObject(pageController objc.IObject /* cross-framework: PageController */, viewController IViewController, object objc.IObject) {
	if d._PageControllerPrepareViewControllerWithObject != nil {
		d._PageControllerPrepareViewControllerWithObject(pageController, viewController, object)
	}
}

// HasPageControllerPrepareViewControllerWithObject returns true if a handler for PageControllerPrepareViewControllerWithObject has been set.
func (d *PageControllerDelegate) HasPageControllerPrepareViewControllerWithObject() bool {
	return d._PageControllerPrepareViewControllerWithObject != nil
}

// PageControllerViewControllerForIdentifier implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerViewControllerForIdentifier(pageController objc.IObject /* cross-framework: PageController */, identifier PageControllerObjectIdentifier /* typedef */) ViewController {
	if d._PageControllerViewControllerForIdentifier != nil {
		return d._PageControllerViewControllerForIdentifier(pageController, identifier)
	}
	var zero ViewController
	return zero
}

// HasPageControllerViewControllerForIdentifier returns true if a handler for PageControllerViewControllerForIdentifier has been set.
func (d *PageControllerDelegate) HasPageControllerViewControllerForIdentifier() bool {
	return d._PageControllerViewControllerForIdentifier != nil
}

// PageControllerDidEndLiveTransition implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerDidEndLiveTransition(pageController objc.IObject /* cross-framework: PageController */) {
	if d._PageControllerDidEndLiveTransition != nil {
		d._PageControllerDidEndLiveTransition(pageController)
	}
}

// HasPageControllerDidEndLiveTransition returns true if a handler for PageControllerDidEndLiveTransition has been set.
func (d *PageControllerDelegate) HasPageControllerDidEndLiveTransition() bool {
	return d._PageControllerDidEndLiveTransition != nil
}

// PageControllerWillStartLiveTransition implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerWillStartLiveTransition(pageController objc.IObject /* cross-framework: PageController */) {
	if d._PageControllerWillStartLiveTransition != nil {
		d._PageControllerWillStartLiveTransition(pageController)
	}
}

// HasPageControllerWillStartLiveTransition returns true if a handler for PageControllerWillStartLiveTransition has been set.
func (d *PageControllerDelegate) HasPageControllerWillStartLiveTransition() bool {
	return d._PageControllerWillStartLiveTransition != nil
}
