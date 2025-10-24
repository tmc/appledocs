// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// POverlayDelegate is the SKOverlayDelegate protocol interface.
//
// Methods for responding to the overlay’s appearance, dismissal, or failure to load.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKOverlayDelegate
type POverlayDelegate interface {
	// Optional methods
	StoreOverlayDidFailToLoadWithError(overlay ISKOverlay, error_ objc.IObject /* cross-framework: Error */)
	HasStoreOverlayDidFailToLoadWithError() bool
	StoreOverlayDidFinishDismissal(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	HasStoreOverlayDidFinishDismissal() bool
	StoreOverlayDidFinishPresentation(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	HasStoreOverlayDidFinishPresentation() bool
	StoreOverlayWillStartDismissal(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	HasStoreOverlayWillStartDismissal() bool
	StoreOverlayWillStartPresentation(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	HasStoreOverlayWillStartPresentation() bool
}

// OverlayDelegate is a delegate implementation builder for the POverlayDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OverlayDelegate struct {
	_StoreOverlayDidFailToLoadWithError func(overlay ISKOverlay, error_ objc.IObject /* cross-framework: Error */)
	_StoreOverlayDidFinishDismissal func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	_StoreOverlayDidFinishPresentation func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	_StoreOverlayWillStartDismissal func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
	_StoreOverlayWillStartPresentation func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)
}

// SetStoreOverlayDidFailToLoadWithError sets the handler for the StoreOverlayDidFailToLoadWithError delegate method.
//
// Indicates that an overlay failed to load.
func (d *OverlayDelegate) SetStoreOverlayDidFailToLoadWithError(f func(overlay ISKOverlay, error_ objc.IObject /* cross-framework: Error */)) {
	d._StoreOverlayDidFailToLoadWithError = f
}

// SetStoreOverlayDidFinishDismissal sets the handler for the StoreOverlayDidFinishDismissal delegate method.
//
// Indicates that platform finished dismissing an overlay.
func (d *OverlayDelegate) SetStoreOverlayDidFinishDismissal(f func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)) {
	d._StoreOverlayDidFinishDismissal = f
}

// SetStoreOverlayDidFinishPresentation sets the handler for the StoreOverlayDidFinishPresentation delegate method.
//
// Indicates that the platform finished presenting an overlay.
func (d *OverlayDelegate) SetStoreOverlayDidFinishPresentation(f func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)) {
	d._StoreOverlayDidFinishPresentation = f
}

// SetStoreOverlayWillStartDismissal sets the handler for the StoreOverlayWillStartDismissal delegate method.
//
// Indicates that the platform dismisses an overlay.
func (d *OverlayDelegate) SetStoreOverlayWillStartDismissal(f func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)) {
	d._StoreOverlayWillStartDismissal = f
}

// SetStoreOverlayWillStartPresentation sets the handler for the StoreOverlayWillStartPresentation delegate method.
//
// Indicates that the platform presents an overlay.
func (d *OverlayDelegate) SetStoreOverlayWillStartPresentation(f func(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext)) {
	d._StoreOverlayWillStartPresentation = f
}

// StoreOverlayDidFailToLoadWithError implements the POverlayDelegate interface.
func (d *OverlayDelegate) StoreOverlayDidFailToLoadWithError(overlay ISKOverlay, error_ objc.IObject /* cross-framework: Error */) {
	if d._StoreOverlayDidFailToLoadWithError != nil {
		d._StoreOverlayDidFailToLoadWithError(overlay, error_)
	}
}

// HasStoreOverlayDidFailToLoadWithError returns true if a handler for StoreOverlayDidFailToLoadWithError has been set.
func (d *OverlayDelegate) HasStoreOverlayDidFailToLoadWithError() bool {
	return d._StoreOverlayDidFailToLoadWithError != nil
}

// StoreOverlayDidFinishDismissal implements the POverlayDelegate interface.
func (d *OverlayDelegate) StoreOverlayDidFinishDismissal(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext) {
	if d._StoreOverlayDidFinishDismissal != nil {
		d._StoreOverlayDidFinishDismissal(overlay, transitionContext)
	}
}

// HasStoreOverlayDidFinishDismissal returns true if a handler for StoreOverlayDidFinishDismissal has been set.
func (d *OverlayDelegate) HasStoreOverlayDidFinishDismissal() bool {
	return d._StoreOverlayDidFinishDismissal != nil
}

// StoreOverlayDidFinishPresentation implements the POverlayDelegate interface.
func (d *OverlayDelegate) StoreOverlayDidFinishPresentation(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext) {
	if d._StoreOverlayDidFinishPresentation != nil {
		d._StoreOverlayDidFinishPresentation(overlay, transitionContext)
	}
}

// HasStoreOverlayDidFinishPresentation returns true if a handler for StoreOverlayDidFinishPresentation has been set.
func (d *OverlayDelegate) HasStoreOverlayDidFinishPresentation() bool {
	return d._StoreOverlayDidFinishPresentation != nil
}

// StoreOverlayWillStartDismissal implements the POverlayDelegate interface.
func (d *OverlayDelegate) StoreOverlayWillStartDismissal(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext) {
	if d._StoreOverlayWillStartDismissal != nil {
		d._StoreOverlayWillStartDismissal(overlay, transitionContext)
	}
}

// HasStoreOverlayWillStartDismissal returns true if a handler for StoreOverlayWillStartDismissal has been set.
func (d *OverlayDelegate) HasStoreOverlayWillStartDismissal() bool {
	return d._StoreOverlayWillStartDismissal != nil
}

// StoreOverlayWillStartPresentation implements the POverlayDelegate interface.
func (d *OverlayDelegate) StoreOverlayWillStartPresentation(overlay ISKOverlay, transitionContext ISKOverlayTransitionContext) {
	if d._StoreOverlayWillStartPresentation != nil {
		d._StoreOverlayWillStartPresentation(overlay, transitionContext)
	}
}

// HasStoreOverlayWillStartPresentation returns true if a handler for StoreOverlayWillStartPresentation has been set.
func (d *OverlayDelegate) HasStoreOverlayWillStartPresentation() bool {
	return d._StoreOverlayWillStartPresentation != nil
}
