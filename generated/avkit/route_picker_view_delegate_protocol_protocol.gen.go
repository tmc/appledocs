// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PRoutePickerViewDelegate is the AVRoutePickerViewDelegate protocol interface.
//
// A protocol that defines the methods to adopt to respond to route picker view presentation events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.15+
//   - tvOS 11.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVRoutePickerViewDelegate
type PRoutePickerViewDelegate interface {
	// Optional methods
	RoutePickerViewDidEndPresentingRoutes(routePickerView IAVRoutePickerView)
	HasRoutePickerViewDidEndPresentingRoutes() bool
	RoutePickerViewWillBeginPresentingRoutes(routePickerView IAVRoutePickerView)
	HasRoutePickerViewWillBeginPresentingRoutes() bool
}

// RoutePickerViewDelegate is a delegate implementation builder for the PRoutePickerViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RoutePickerViewDelegate struct {
	_RoutePickerViewDidEndPresentingRoutes func(routePickerView IAVRoutePickerView)
	_RoutePickerViewWillBeginPresentingRoutes func(routePickerView IAVRoutePickerView)
}

// SetRoutePickerViewDidEndPresentingRoutes sets the handler for the RoutePickerViewDidEndPresentingRoutes delegate method.
//
// Tells the delegate when the route picker view finishes presenting routes to the user.
func (d *RoutePickerViewDelegate) SetRoutePickerViewDidEndPresentingRoutes(f func(routePickerView IAVRoutePickerView)) {
	d._RoutePickerViewDidEndPresentingRoutes = f
}

// SetRoutePickerViewWillBeginPresentingRoutes sets the handler for the RoutePickerViewWillBeginPresentingRoutes delegate method.
//
// Tells the delegate that the route picker view is about to begin presenting routes to the user.
func (d *RoutePickerViewDelegate) SetRoutePickerViewWillBeginPresentingRoutes(f func(routePickerView IAVRoutePickerView)) {
	d._RoutePickerViewWillBeginPresentingRoutes = f
}

// RoutePickerViewDidEndPresentingRoutes implements the PRoutePickerViewDelegate interface.
func (d *RoutePickerViewDelegate) RoutePickerViewDidEndPresentingRoutes(routePickerView IAVRoutePickerView) {
	if d._RoutePickerViewDidEndPresentingRoutes != nil {
		d._RoutePickerViewDidEndPresentingRoutes(routePickerView)
	}
}

// HasRoutePickerViewDidEndPresentingRoutes returns true if a handler for RoutePickerViewDidEndPresentingRoutes has been set.
func (d *RoutePickerViewDelegate) HasRoutePickerViewDidEndPresentingRoutes() bool {
	return d._RoutePickerViewDidEndPresentingRoutes != nil
}

// RoutePickerViewWillBeginPresentingRoutes implements the PRoutePickerViewDelegate interface.
func (d *RoutePickerViewDelegate) RoutePickerViewWillBeginPresentingRoutes(routePickerView IAVRoutePickerView) {
	if d._RoutePickerViewWillBeginPresentingRoutes != nil {
		d._RoutePickerViewWillBeginPresentingRoutes(routePickerView)
	}
}

// HasRoutePickerViewWillBeginPresentingRoutes returns true if a handler for RoutePickerViewWillBeginPresentingRoutes has been set.
func (d *RoutePickerViewDelegate) HasRoutePickerViewWillBeginPresentingRoutes() bool {
	return d._RoutePickerViewWillBeginPresentingRoutes != nil
}
