// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMKLookAroundViewControllerDelegate is the MKLookAroundViewControllerDelegate protocol interface.
//
// Methods you implement to respond to changes in the LookAround view controller.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKLookAroundViewControllerDelegate
type PMKLookAroundViewControllerDelegate interface {
	// Optional methods
	LookAroundViewControllerDidDismissFullScreen(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerDidDismissFullScreen() bool
	LookAroundViewControllerDidPresentFullScreen(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerDidPresentFullScreen() bool
	LookAroundViewControllerDidUpdateScene(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerDidUpdateScene() bool
	LookAroundViewControllerWillDismissFullScreen(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerWillDismissFullScreen() bool
	LookAroundViewControllerWillPresentFullScreen(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerWillPresentFullScreen() bool
	LookAroundViewControllerWillUpdateScene(viewController IMKLookAroundViewController)
	HasLookAroundViewControllerWillUpdateScene() bool
}

// MKLookAroundViewControllerDelegate is a delegate implementation builder for the PMKLookAroundViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MKLookAroundViewControllerDelegate struct {
	_LookAroundViewControllerDidDismissFullScreen func(viewController IMKLookAroundViewController)
	_LookAroundViewControllerDidPresentFullScreen func(viewController IMKLookAroundViewController)
	_LookAroundViewControllerDidUpdateScene func(viewController IMKLookAroundViewController)
	_LookAroundViewControllerWillDismissFullScreen func(viewController IMKLookAroundViewController)
	_LookAroundViewControllerWillPresentFullScreen func(viewController IMKLookAroundViewController)
	_LookAroundViewControllerWillUpdateScene func(viewController IMKLookAroundViewController)
}

// SetLookAroundViewControllerDidDismissFullScreen sets the handler for the LookAroundViewControllerDidDismissFullScreen delegate method.
//
// Tells the delegate when the view controller exits full-screen mode.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerDidDismissFullScreen(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerDidDismissFullScreen = f
}

// SetLookAroundViewControllerDidPresentFullScreen sets the handler for the LookAroundViewControllerDidPresentFullScreen delegate method.
//
// Tells the delegate when the view controller enters full-screen mode.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerDidPresentFullScreen(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerDidPresentFullScreen = f
}

// SetLookAroundViewControllerDidUpdateScene sets the handler for the LookAroundViewControllerDidUpdateScene delegate method.
//
// Tells the delegate that the scene updated.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerDidUpdateScene(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerDidUpdateScene = f
}

// SetLookAroundViewControllerWillDismissFullScreen sets the handler for the LookAroundViewControllerWillDismissFullScreen delegate method.
//
// Tells the delegate when the view controller is about to exit full-screen mode.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerWillDismissFullScreen(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerWillDismissFullScreen = f
}

// SetLookAroundViewControllerWillPresentFullScreen sets the handler for the LookAroundViewControllerWillPresentFullScreen delegate method.
//
// Tells the delegate when the view controller is about to enter full-screen mode.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerWillPresentFullScreen(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerWillPresentFullScreen = f
}

// SetLookAroundViewControllerWillUpdateScene sets the handler for the LookAroundViewControllerWillUpdateScene delegate method.
//
// Tells the delegate that the scene is about to update.
func (d *MKLookAroundViewControllerDelegate) SetLookAroundViewControllerWillUpdateScene(f func(viewController IMKLookAroundViewController)) {
	d._LookAroundViewControllerWillUpdateScene = f
}

// LookAroundViewControllerDidDismissFullScreen implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerDidDismissFullScreen(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerDidDismissFullScreen != nil {
		d._LookAroundViewControllerDidDismissFullScreen(viewController)
	}
}

// HasLookAroundViewControllerDidDismissFullScreen returns true if a handler for LookAroundViewControllerDidDismissFullScreen has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerDidDismissFullScreen() bool {
	return d._LookAroundViewControllerDidDismissFullScreen != nil
}

// LookAroundViewControllerDidPresentFullScreen implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerDidPresentFullScreen(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerDidPresentFullScreen != nil {
		d._LookAroundViewControllerDidPresentFullScreen(viewController)
	}
}

// HasLookAroundViewControllerDidPresentFullScreen returns true if a handler for LookAroundViewControllerDidPresentFullScreen has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerDidPresentFullScreen() bool {
	return d._LookAroundViewControllerDidPresentFullScreen != nil
}

// LookAroundViewControllerDidUpdateScene implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerDidUpdateScene(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerDidUpdateScene != nil {
		d._LookAroundViewControllerDidUpdateScene(viewController)
	}
}

// HasLookAroundViewControllerDidUpdateScene returns true if a handler for LookAroundViewControllerDidUpdateScene has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerDidUpdateScene() bool {
	return d._LookAroundViewControllerDidUpdateScene != nil
}

// LookAroundViewControllerWillDismissFullScreen implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerWillDismissFullScreen(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerWillDismissFullScreen != nil {
		d._LookAroundViewControllerWillDismissFullScreen(viewController)
	}
}

// HasLookAroundViewControllerWillDismissFullScreen returns true if a handler for LookAroundViewControllerWillDismissFullScreen has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerWillDismissFullScreen() bool {
	return d._LookAroundViewControllerWillDismissFullScreen != nil
}

// LookAroundViewControllerWillPresentFullScreen implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerWillPresentFullScreen(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerWillPresentFullScreen != nil {
		d._LookAroundViewControllerWillPresentFullScreen(viewController)
	}
}

// HasLookAroundViewControllerWillPresentFullScreen returns true if a handler for LookAroundViewControllerWillPresentFullScreen has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerWillPresentFullScreen() bool {
	return d._LookAroundViewControllerWillPresentFullScreen != nil
}

// LookAroundViewControllerWillUpdateScene implements the PMKLookAroundViewControllerDelegate interface.
func (d *MKLookAroundViewControllerDelegate) LookAroundViewControllerWillUpdateScene(viewController IMKLookAroundViewController) {
	if d._LookAroundViewControllerWillUpdateScene != nil {
		d._LookAroundViewControllerWillUpdateScene(viewController)
	}
}

// HasLookAroundViewControllerWillUpdateScene returns true if a handler for LookAroundViewControllerWillUpdateScene has been set.
func (d *MKLookAroundViewControllerDelegate) HasLookAroundViewControllerWillUpdateScene() bool {
	return d._LookAroundViewControllerWillUpdateScene != nil
}
