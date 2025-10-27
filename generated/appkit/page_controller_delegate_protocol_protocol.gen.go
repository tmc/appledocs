// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	PageControllerDidTransitionToObject(pageController PageController /* not a class type */, object objectivec.IObject)
	HasPageControllerDidTransitionToObject() bool
	PageControllerFrameForObject(pageController PageController /* not a class type */, object objectivec.IObject) corefoundation.CGRect
	HasPageControllerFrameForObject() bool
	PageControllerIdentifierForObject(pageController PageController /* not a class type */, object objectivec.IObject) PageControllerObjectIdentifier
	HasPageControllerIdentifierForObject() bool
	PageControllerPrepareViewControllerWithObject(pageController PageController /* not a class type */, viewController IViewController, object objectivec.IObject)
	HasPageControllerPrepareViewControllerWithObject() bool
	PageControllerViewControllerForIdentifier(pageController PageController /* not a class type */, identifier PageControllerObjectIdentifier) IViewController
	HasPageControllerViewControllerForIdentifier() bool
	PageControllerDidEndLiveTransition(pageController PageController /* not a class type */)
	HasPageControllerDidEndLiveTransition() bool
	PageControllerWillStartLiveTransition(pageController PageController /* not a class type */)
	HasPageControllerWillStartLiveTransition() bool
}

// PageControllerDelegate is a delegate implementation builder for the PPageControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PageControllerDelegate struct {
	_PageControllerDidTransitionToObject func(pageController PageController /* not a class type */, object objectivec.IObject)
	_PageControllerFrameForObject func(pageController PageController /* not a class type */, object objectivec.IObject) corefoundation.CGRect
	_PageControllerIdentifierForObject func(pageController PageController /* not a class type */, object objectivec.IObject) PageControllerObjectIdentifier
	_PageControllerPrepareViewControllerWithObject func(pageController PageController /* not a class type */, viewController IViewController, object objectivec.IObject)
	_PageControllerViewControllerForIdentifier func(pageController PageController /* not a class type */, identifier PageControllerObjectIdentifier) IViewController
	_PageControllerDidEndLiveTransition func(pageController PageController /* not a class type */)
	_PageControllerWillStartLiveTransition func(pageController PageController /* not a class type */)
}

// SetPageControllerDidTransitionToObject sets the handler for the PageControllerDidTransitionToObject delegate method.
//
// This message is sent when any page transition is completed.
func (d *PageControllerDelegate) SetPageControllerDidTransitionToObject(f func(pageController PageController /* not a class type */, object objectivec.IObject)) {
	d._PageControllerDidTransitionToObject = f
}

// SetPageControllerFrameForObject sets the handler for the PageControllerFrameForObject delegate method.
//
// Returns the frame appropriate for displaying the specified object.
func (d *PageControllerDelegate) SetPageControllerFrameForObject(f func(pageController PageController /* not a class type */, object objectivec.IObject) corefoundation.CGRect) {
	d._PageControllerFrameForObject = f
}

// SetPageControllerIdentifierForObject sets the handler for the PageControllerIdentifierForObject delegate method.
//
// Return the identifier of the view controller that owns a view to display the object.
func (d *PageControllerDelegate) SetPageControllerIdentifierForObject(f func(pageController PageController /* not a class type */, object objectivec.IObject) PageControllerObjectIdentifier) {
	d._PageControllerIdentifierForObject = f
}

// SetPageControllerPrepareViewControllerWithObject sets the handler for the PageControllerPrepareViewControllerWithObject delegate method.
//
// Prepare the view controller and it’s view for drawing.
func (d *PageControllerDelegate) SetPageControllerPrepareViewControllerWithObject(f func(pageController PageController /* not a class type */, viewController IViewController, object objectivec.IObject)) {
	d._PageControllerPrepareViewControllerWithObject = f
}

// SetPageControllerViewControllerForIdentifier sets the handler for the PageControllerViewControllerForIdentifier delegate method.
//
// Returns a view controller the page controller uses for managing the specified identifier.
func (d *PageControllerDelegate) SetPageControllerViewControllerForIdentifier(f func(pageController PageController /* not a class type */, identifier PageControllerObjectIdentifier) IViewController) {
	d._PageControllerViewControllerForIdentifier = f
}

// SetPageControllerDidEndLiveTransition sets the handler for the PageControllerDidEndLiveTransition delegate method.
//
// This message is sent when a transition animation completes.
func (d *PageControllerDelegate) SetPageControllerDidEndLiveTransition(f func(pageController PageController /* not a class type */)) {
	d._PageControllerDidEndLiveTransition = f
}

// SetPageControllerWillStartLiveTransition sets the handler for the PageControllerWillStartLiveTransition delegate method.
//
// This message is sent when the user begins a transition.
func (d *PageControllerDelegate) SetPageControllerWillStartLiveTransition(f func(pageController PageController /* not a class type */)) {
	d._PageControllerWillStartLiveTransition = f
}

// PageControllerDidTransitionToObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerDidTransitionToObject(pageController PageController /* not a class type */, object objectivec.IObject) {
	if d._PageControllerDidTransitionToObject != nil {
		d._PageControllerDidTransitionToObject(pageController, object)
	}
}

// HasPageControllerDidTransitionToObject returns true if a handler for PageControllerDidTransitionToObject has been set.
func (d *PageControllerDelegate) HasPageControllerDidTransitionToObject() bool {
	return d._PageControllerDidTransitionToObject != nil
}

// PageControllerFrameForObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerFrameForObject(pageController PageController /* not a class type */, object objectivec.IObject) corefoundation.CGRect {
	if d._PageControllerFrameForObject != nil {
		return d._PageControllerFrameForObject(pageController, object)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasPageControllerFrameForObject returns true if a handler for PageControllerFrameForObject has been set.
func (d *PageControllerDelegate) HasPageControllerFrameForObject() bool {
	return d._PageControllerFrameForObject != nil
}

// PageControllerIdentifierForObject implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerIdentifierForObject(pageController PageController /* not a class type */, object objectivec.IObject) PageControllerObjectIdentifier {
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
func (d *PageControllerDelegate) PageControllerPrepareViewControllerWithObject(pageController PageController /* not a class type */, viewController IViewController, object objectivec.IObject) {
	if d._PageControllerPrepareViewControllerWithObject != nil {
		d._PageControllerPrepareViewControllerWithObject(pageController, viewController, object)
	}
}

// HasPageControllerPrepareViewControllerWithObject returns true if a handler for PageControllerPrepareViewControllerWithObject has been set.
func (d *PageControllerDelegate) HasPageControllerPrepareViewControllerWithObject() bool {
	return d._PageControllerPrepareViewControllerWithObject != nil
}

// PageControllerViewControllerForIdentifier implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerViewControllerForIdentifier(pageController PageController /* not a class type */, identifier PageControllerObjectIdentifier) IViewController {
	if d._PageControllerViewControllerForIdentifier != nil {
		return d._PageControllerViewControllerForIdentifier(pageController, identifier)
	}
	var zero IViewController
	return zero
}

// HasPageControllerViewControllerForIdentifier returns true if a handler for PageControllerViewControllerForIdentifier has been set.
func (d *PageControllerDelegate) HasPageControllerViewControllerForIdentifier() bool {
	return d._PageControllerViewControllerForIdentifier != nil
}

// PageControllerDidEndLiveTransition implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerDidEndLiveTransition(pageController PageController /* not a class type */) {
	if d._PageControllerDidEndLiveTransition != nil {
		d._PageControllerDidEndLiveTransition(pageController)
	}
}

// HasPageControllerDidEndLiveTransition returns true if a handler for PageControllerDidEndLiveTransition has been set.
func (d *PageControllerDelegate) HasPageControllerDidEndLiveTransition() bool {
	return d._PageControllerDidEndLiveTransition != nil
}

// PageControllerWillStartLiveTransition implements the PPageControllerDelegate interface.
func (d *PageControllerDelegate) PageControllerWillStartLiveTransition(pageController PageController /* not a class type */) {
	if d._PageControllerWillStartLiveTransition != nil {
		d._PageControllerWillStartLiveTransition(pageController)
	}
}

// HasPageControllerWillStartLiveTransition returns true if a handler for PageControllerWillStartLiveTransition has been set.
func (d *PageControllerDelegate) HasPageControllerWillStartLiveTransition() bool {
	return d._PageControllerWillStartLiveTransition != nil
}

// PageControllerDelegateObject wraps an existing Objective-C object that conforms to the PPageControllerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type PageControllerDelegateObject struct {
	objectivec.Object
}

// NewPageControllerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSPageControllerDelegate protocol.
func NewPageControllerDelegateObject(obj objectivec.Object) *PageControllerDelegateObject {
	return &PageControllerDelegateObject{obj}
}

// Make sure PageControllerDelegateObject implements PPageControllerDelegate.
var _ PPageControllerDelegate = (*PageControllerDelegateObject)(nil)

// PageControllerDidTransitionToObject implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerDidTransitionToObject(pageController PageController /* not a class type */, object objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("pageController:didTransitionToObject:"), pageController, object)
}

// HasPageControllerDidTransitionToObject returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerDidTransitionToObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerFrameForObject implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerFrameForObject(pageController PageController /* not a class type */, object objectivec.IObject) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("pageController:frameForObject:"), pageController, object)
}

// HasPageControllerFrameForObject returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerFrameForObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerIdentifierForObject implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerIdentifierForObject(pageController PageController /* not a class type */, object objectivec.IObject) PageControllerObjectIdentifier {
	return objc.Send[PageControllerObjectIdentifier](o.ID, objc.Sel("pageController:identifierForObject:"), pageController, object)
}

// HasPageControllerIdentifierForObject returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerIdentifierForObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerPrepareViewControllerWithObject implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerPrepareViewControllerWithObject(pageController PageController /* not a class type */, viewController IViewController, object objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("pageController:prepareViewController:withObject:"), pageController, viewController, object)
}

// HasPageControllerPrepareViewControllerWithObject returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerPrepareViewControllerWithObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerViewControllerForIdentifier implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerViewControllerForIdentifier(pageController PageController /* not a class type */, identifier PageControllerObjectIdentifier) IViewController {
	return objc.Send[IViewController](o.ID, objc.Sel("pageController:viewControllerForIdentifier:"), pageController, identifier)
}

// HasPageControllerViewControllerForIdentifier returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerViewControllerForIdentifier() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerDidEndLiveTransition implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerDidEndLiveTransition(pageController PageController /* not a class type */) {
	objc.Send[objc.ID](o.ID, objc.Sel("pageControllerDidEndLiveTransition:"), pageController)
}

// HasPageControllerDidEndLiveTransition returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerDidEndLiveTransition() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// PageControllerWillStartLiveTransition implements the PPageControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PageControllerDelegateObject) PageControllerWillStartLiveTransition(pageController PageController /* not a class type */) {
	objc.Send[objc.ID](o.ID, objc.Sel("pageControllerWillStartLiveTransition:"), pageController)
}

// HasPageControllerWillStartLiveTransition returns true; this is a placeholder for optional method checks.
func (o *PageControllerDelegateObject) HasPageControllerWillStartLiveTransition() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
