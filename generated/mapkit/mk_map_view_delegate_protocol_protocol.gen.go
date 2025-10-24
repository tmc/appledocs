// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMKMapViewDelegate is the MKMapViewDelegate protocol interface.
//
// Optional methods that you use to receive map-related update messages.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKMapViewDelegate
type PMKMapViewDelegate interface {
	// Optional methods
	MapViewAnnotationViewCalloutAccessoryControlTapped(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, control appkit.Control)
	HasMapViewAnnotationViewCalloutAccessoryControlTapped() bool
	MapViewAnnotationViewDidChangeDragStateFromOldState(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, newState MKAnnotationViewDragState, oldState MKAnnotationViewDragState)
	HasMapViewAnnotationViewDidChangeDragStateFromOldState() bool
	MapViewClusterAnnotationForMemberAnnotations(mapView IMKMapView, memberAnnotations []objc.ID) MKClusterAnnotation
	HasMapViewClusterAnnotationForMemberAnnotations() bool
	MapViewDidAddAnnotationViews(mapView IMKMapView, views []objc.IObject /* cross-framework: MKAnnotationView */)
	HasMapViewDidAddAnnotationViews() bool
	MapViewDidAddOverlayRenderers(mapView IMKMapView, renderers []MKOverlayRenderer)
	HasMapViewDidAddOverlayRenderers() bool
	MapViewDidAddOverlayViews(mapView IMKMapView, overlayViews objc.IObject /* cross-framework: NSArray */)
	HasMapViewDidAddOverlayViews() bool
	MapViewDidChangeUserTrackingModeAnimated(mapView IMKMapView, mode MKUserTrackingMode, animated bool)
	HasMapViewDidChangeUserTrackingModeAnimated() bool
	MapViewDidDeselectAnnotation(mapView IMKMapView, annotation unsafe.Pointer)
	HasMapViewDidDeselectAnnotation() bool
	MapViewDidDeselectAnnotationView(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)
	HasMapViewDidDeselectAnnotationView() bool
	MapViewDidFailToLocateUserWithError(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)
	HasMapViewDidFailToLocateUserWithError() bool
	MapViewDidSelectAnnotationView(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)
	HasMapViewDidSelectAnnotationView() bool
	MapViewDidSelectAnnotation(mapView IMKMapView, annotation unsafe.Pointer)
	HasMapViewDidSelectAnnotation() bool
	MapViewDidUpdateUserLocation(mapView IMKMapView, userLocation IMKUserLocation)
	HasMapViewDidUpdateUserLocation() bool
	MapViewRegionDidChangeAnimated(mapView IMKMapView, animated bool)
	HasMapViewRegionDidChangeAnimated() bool
	MapViewRegionWillChangeAnimated(mapView IMKMapView, animated bool)
	HasMapViewRegionWillChangeAnimated() bool
	MapViewRendererForOverlay(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayRenderer
	HasMapViewRendererForOverlay() bool
	MapViewSelectionAccessoryForAnnotation(mapView IMKMapView, annotation unsafe.Pointer) MKSelectionAccessory
	HasMapViewSelectionAccessoryForAnnotation() bool
	MapViewViewForOverlay(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayView
	HasMapViewViewForOverlay() bool
	MapViewViewForAnnotation(mapView IMKMapView, annotation unsafe.Pointer) MKAnnotationView
	HasMapViewViewForAnnotation() bool
	MapViewDidChangeVisibleRegion(mapView IMKMapView)
	HasMapViewDidChangeVisibleRegion() bool
	MapViewDidFailLoadingMapWithError(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)
	HasMapViewDidFailLoadingMapWithError() bool
	MapViewDidFinishLoadingMap(mapView IMKMapView)
	HasMapViewDidFinishLoadingMap() bool
	MapViewDidFinishRenderingMapFullyRendered(mapView IMKMapView, fullyRendered bool)
	HasMapViewDidFinishRenderingMapFullyRendered() bool
	MapViewDidStopLocatingUser(mapView IMKMapView)
	HasMapViewDidStopLocatingUser() bool
	MapViewWillStartLoadingMap(mapView IMKMapView)
	HasMapViewWillStartLoadingMap() bool
	MapViewWillStartLocatingUser(mapView IMKMapView)
	HasMapViewWillStartLocatingUser() bool
	MapViewWillStartRenderingMap(mapView IMKMapView)
	HasMapViewWillStartRenderingMap() bool
}

// MKMapViewDelegate is a delegate implementation builder for the PMKMapViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MKMapViewDelegate struct {
	_MapViewAnnotationViewCalloutAccessoryControlTapped func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, control appkit.Control)
	_MapViewAnnotationViewDidChangeDragStateFromOldState func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, newState MKAnnotationViewDragState, oldState MKAnnotationViewDragState)
	_MapViewClusterAnnotationForMemberAnnotations func(mapView IMKMapView, memberAnnotations []objc.ID) MKClusterAnnotation
	_MapViewDidAddAnnotationViews func(mapView IMKMapView, views []objc.IObject /* cross-framework: MKAnnotationView */)
	_MapViewDidAddOverlayRenderers func(mapView IMKMapView, renderers []MKOverlayRenderer)
	_MapViewDidAddOverlayViews func(mapView IMKMapView, overlayViews objc.IObject /* cross-framework: NSArray */)
	_MapViewDidChangeUserTrackingModeAnimated func(mapView IMKMapView, mode MKUserTrackingMode, animated bool)
	_MapViewDidDeselectAnnotation func(mapView IMKMapView, annotation unsafe.Pointer)
	_MapViewDidDeselectAnnotationView func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)
	_MapViewDidFailToLocateUserWithError func(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)
	_MapViewDidSelectAnnotationView func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)
	_MapViewDidSelectAnnotation func(mapView IMKMapView, annotation unsafe.Pointer)
	_MapViewDidUpdateUserLocation func(mapView IMKMapView, userLocation IMKUserLocation)
	_MapViewRegionDidChangeAnimated func(mapView IMKMapView, animated bool)
	_MapViewRegionWillChangeAnimated func(mapView IMKMapView, animated bool)
	_MapViewRendererForOverlay func(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayRenderer
	_MapViewSelectionAccessoryForAnnotation func(mapView IMKMapView, annotation unsafe.Pointer) MKSelectionAccessory
	_MapViewViewForOverlay func(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayView
	_MapViewViewForAnnotation func(mapView IMKMapView, annotation unsafe.Pointer) MKAnnotationView
	_MapViewDidChangeVisibleRegion func(mapView IMKMapView)
	_MapViewDidFailLoadingMapWithError func(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)
	_MapViewDidFinishLoadingMap func(mapView IMKMapView)
	_MapViewDidFinishRenderingMapFullyRendered func(mapView IMKMapView, fullyRendered bool)
	_MapViewDidStopLocatingUser func(mapView IMKMapView)
	_MapViewWillStartLoadingMap func(mapView IMKMapView)
	_MapViewWillStartLocatingUser func(mapView IMKMapView)
	_MapViewWillStartRenderingMap func(mapView IMKMapView)
}

// SetMapViewAnnotationViewCalloutAccessoryControlTapped sets the handler for the MapViewAnnotationViewCalloutAccessoryControlTapped delegate method.
//
// Tells the delegate when the user taps one of the annotation view’s accessory buttons.
func (d *MKMapViewDelegate) SetMapViewAnnotationViewCalloutAccessoryControlTapped(f func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, control appkit.Control)) {
	d._MapViewAnnotationViewCalloutAccessoryControlTapped = f
}

// SetMapViewAnnotationViewDidChangeDragStateFromOldState sets the handler for the MapViewAnnotationViewDidChangeDragStateFromOldState delegate method.
//
// Tells the delegate when the drag state of one of its annotation views changes.
func (d *MKMapViewDelegate) SetMapViewAnnotationViewDidChangeDragStateFromOldState(f func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, newState MKAnnotationViewDragState, oldState MKAnnotationViewDragState)) {
	d._MapViewAnnotationViewDidChangeDragStateFromOldState = f
}

// SetMapViewClusterAnnotationForMemberAnnotations sets the handler for the MapViewClusterAnnotationForMemberAnnotations delegate method.
//
// Asks the delegate to provide a cluster annotation object for the specified annotations.
func (d *MKMapViewDelegate) SetMapViewClusterAnnotationForMemberAnnotations(f func(mapView IMKMapView, memberAnnotations []objc.ID) MKClusterAnnotation) {
	d._MapViewClusterAnnotationForMemberAnnotations = f
}

// SetMapViewDidAddAnnotationViews sets the handler for the MapViewDidAddAnnotationViews delegate method.
//
// Tells the delegate when the map view adds one or more annotation views to the map.
func (d *MKMapViewDelegate) SetMapViewDidAddAnnotationViews(f func(mapView IMKMapView, views []objc.IObject /* cross-framework: MKAnnotationView */)) {
	d._MapViewDidAddAnnotationViews = f
}

// SetMapViewDidAddOverlayRenderers sets the handler for the MapViewDidAddOverlayRenderers delegate method.
//
// Tells the delegate when the map view adds one or more renderer objects to the map.
func (d *MKMapViewDelegate) SetMapViewDidAddOverlayRenderers(f func(mapView IMKMapView, renderers []MKOverlayRenderer)) {
	d._MapViewDidAddOverlayRenderers = f
}

// SetMapViewDidAddOverlayViews sets the handler for the MapViewDidAddOverlayViews delegate method.
//
// Tells the delegate when the map adds one or more overlay views to the map.
func (d *MKMapViewDelegate) SetMapViewDidAddOverlayViews(f func(mapView IMKMapView, overlayViews objc.IObject /* cross-framework: NSArray */)) {
	d._MapViewDidAddOverlayViews = f
}

// SetMapViewDidChangeUserTrackingModeAnimated sets the handler for the MapViewDidChangeUserTrackingModeAnimated delegate method.
//
// Tells the delegate when the user-tracking mode changes.
func (d *MKMapViewDelegate) SetMapViewDidChangeUserTrackingModeAnimated(f func(mapView IMKMapView, mode MKUserTrackingMode, animated bool)) {
	d._MapViewDidChangeUserTrackingModeAnimated = f
}

// SetMapViewDidDeselectAnnotation sets the handler for the MapViewDidDeselectAnnotation delegate method.
//
// Tells the delegate when the user deselects one or more annotations.
func (d *MKMapViewDelegate) SetMapViewDidDeselectAnnotation(f func(mapView IMKMapView, annotation unsafe.Pointer)) {
	d._MapViewDidDeselectAnnotation = f
}

// SetMapViewDidDeselectAnnotationView sets the handler for the MapViewDidDeselectAnnotationView delegate method.
//
// Tells the delegate when the user deselects one or more of its annotation views.
func (d *MKMapViewDelegate) SetMapViewDidDeselectAnnotationView(f func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)) {
	d._MapViewDidDeselectAnnotationView = f
}

// SetMapViewDidFailToLocateUserWithError sets the handler for the MapViewDidFailToLocateUserWithError delegate method.
//
// Tells the delegate when an attempt to locate the user’s location fails.
func (d *MKMapViewDelegate) SetMapViewDidFailToLocateUserWithError(f func(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)) {
	d._MapViewDidFailToLocateUserWithError = f
}

// SetMapViewDidSelectAnnotationView sets the handler for the MapViewDidSelectAnnotationView delegate method.
//
// Tells the delegate when the user selects one or more of its annotation views.
func (d *MKMapViewDelegate) SetMapViewDidSelectAnnotationView(f func(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */)) {
	d._MapViewDidSelectAnnotationView = f
}

// SetMapViewDidSelectAnnotation sets the handler for the MapViewDidSelectAnnotation delegate method.
//
// Tells the delegate when the user selects one or more annotations.
func (d *MKMapViewDelegate) SetMapViewDidSelectAnnotation(f func(mapView IMKMapView, annotation unsafe.Pointer)) {
	d._MapViewDidSelectAnnotation = f
}

// SetMapViewDidUpdateUserLocation sets the handler for the MapViewDidUpdateUserLocation delegate method.
//
// Tells the delegate when the map view updates the user’s location.
func (d *MKMapViewDelegate) SetMapViewDidUpdateUserLocation(f func(mapView IMKMapView, userLocation IMKUserLocation)) {
	d._MapViewDidUpdateUserLocation = f
}

// SetMapViewRegionDidChangeAnimated sets the handler for the MapViewRegionDidChangeAnimated delegate method.
//
// Tells the delegate when the region the map view is displaying changes.
func (d *MKMapViewDelegate) SetMapViewRegionDidChangeAnimated(f func(mapView IMKMapView, animated bool)) {
	d._MapViewRegionDidChangeAnimated = f
}

// SetMapViewRegionWillChangeAnimated sets the handler for the MapViewRegionWillChangeAnimated delegate method.
//
// Tells the delegate when the region the map view is displaying is about to change.
func (d *MKMapViewDelegate) SetMapViewRegionWillChangeAnimated(f func(mapView IMKMapView, animated bool)) {
	d._MapViewRegionWillChangeAnimated = f
}

// SetMapViewRendererForOverlay sets the handler for the MapViewRendererForOverlay delegate method.
//
// Asks the delegate for a renderer object to use when drawing the specified overlay.
func (d *MKMapViewDelegate) SetMapViewRendererForOverlay(f func(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayRenderer) {
	d._MapViewRendererForOverlay = f
}

// SetMapViewSelectionAccessoryForAnnotation sets the handler for the MapViewSelectionAccessoryForAnnotation delegate method.
//
// Specifies the accessory to display for a selected annotation
func (d *MKMapViewDelegate) SetMapViewSelectionAccessoryForAnnotation(f func(mapView IMKMapView, annotation unsafe.Pointer) MKSelectionAccessory) {
	d._MapViewSelectionAccessoryForAnnotation = f
}

// SetMapViewViewForOverlay sets the handler for the MapViewViewForOverlay delegate method.
//
// Asks the delegate for the overlay view to use when displaying the specified overlay object.
func (d *MKMapViewDelegate) SetMapViewViewForOverlay(f func(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayView) {
	d._MapViewViewForOverlay = f
}

// SetMapViewViewForAnnotation sets the handler for the MapViewViewForAnnotation delegate method.
//
// Returns the view associated with the specified annotation object.
func (d *MKMapViewDelegate) SetMapViewViewForAnnotation(f func(mapView IMKMapView, annotation unsafe.Pointer) MKAnnotationView) {
	d._MapViewViewForAnnotation = f
}

// SetMapViewDidChangeVisibleRegion sets the handler for the MapViewDidChangeVisibleRegion delegate method.
//
// Tells the delegate when the map view’s visible region changes.
func (d *MKMapViewDelegate) SetMapViewDidChangeVisibleRegion(f func(mapView IMKMapView)) {
	d._MapViewDidChangeVisibleRegion = f
}

// SetMapViewDidFailLoadingMapWithError sets the handler for the MapViewDidFailLoadingMapWithError delegate method.
//
// Tells the delegate that the specified view is unable to load the map data.
func (d *MKMapViewDelegate) SetMapViewDidFailLoadingMapWithError(f func(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */)) {
	d._MapViewDidFailLoadingMapWithError = f
}

// SetMapViewDidFinishLoadingMap sets the handler for the MapViewDidFinishLoadingMap delegate method.
//
// Tells the delegate when the specified map view successfully loads the needed map data.
func (d *MKMapViewDelegate) SetMapViewDidFinishLoadingMap(f func(mapView IMKMapView)) {
	d._MapViewDidFinishLoadingMap = f
}

// SetMapViewDidFinishRenderingMapFullyRendered sets the handler for the MapViewDidFinishRenderingMapFullyRendered delegate method.
//
// Tells the delegate when the map view finishes rendering all visible tiles.
func (d *MKMapViewDelegate) SetMapViewDidFinishRenderingMapFullyRendered(f func(mapView IMKMapView, fullyRendered bool)) {
	d._MapViewDidFinishRenderingMapFullyRendered = f
}

// SetMapViewDidStopLocatingUser sets the handler for the MapViewDidStopLocatingUser delegate method.
//
// Tells the delegate when the map view stops tracking the user’s location.
func (d *MKMapViewDelegate) SetMapViewDidStopLocatingUser(f func(mapView IMKMapView)) {
	d._MapViewDidStopLocatingUser = f
}

// SetMapViewWillStartLoadingMap sets the handler for the MapViewWillStartLoadingMap delegate method.
//
// Tells the delegate that the specified map view is about to retrieve some map data.
func (d *MKMapViewDelegate) SetMapViewWillStartLoadingMap(f func(mapView IMKMapView)) {
	d._MapViewWillStartLoadingMap = f
}

// SetMapViewWillStartLocatingUser sets the handler for the MapViewWillStartLocatingUser delegate method.
//
// Tells the delegate that the map view is about to start tracking the user’s location.
func (d *MKMapViewDelegate) SetMapViewWillStartLocatingUser(f func(mapView IMKMapView)) {
	d._MapViewWillStartLocatingUser = f
}

// SetMapViewWillStartRenderingMap sets the handler for the MapViewWillStartRenderingMap delegate method.
//
// Tells the delegate that the map view is about to start rendering some of its tiles.
func (d *MKMapViewDelegate) SetMapViewWillStartRenderingMap(f func(mapView IMKMapView)) {
	d._MapViewWillStartRenderingMap = f
}

// MapViewAnnotationViewCalloutAccessoryControlTapped implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewAnnotationViewCalloutAccessoryControlTapped(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, control appkit.Control) {
	if d._MapViewAnnotationViewCalloutAccessoryControlTapped != nil {
		d._MapViewAnnotationViewCalloutAccessoryControlTapped(mapView, view, control)
	}
}

// HasMapViewAnnotationViewCalloutAccessoryControlTapped returns true if a handler for MapViewAnnotationViewCalloutAccessoryControlTapped has been set.
func (d *MKMapViewDelegate) HasMapViewAnnotationViewCalloutAccessoryControlTapped() bool {
	return d._MapViewAnnotationViewCalloutAccessoryControlTapped != nil
}

// MapViewAnnotationViewDidChangeDragStateFromOldState implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewAnnotationViewDidChangeDragStateFromOldState(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */, newState MKAnnotationViewDragState, oldState MKAnnotationViewDragState) {
	if d._MapViewAnnotationViewDidChangeDragStateFromOldState != nil {
		d._MapViewAnnotationViewDidChangeDragStateFromOldState(mapView, view, newState, oldState)
	}
}

// HasMapViewAnnotationViewDidChangeDragStateFromOldState returns true if a handler for MapViewAnnotationViewDidChangeDragStateFromOldState has been set.
func (d *MKMapViewDelegate) HasMapViewAnnotationViewDidChangeDragStateFromOldState() bool {
	return d._MapViewAnnotationViewDidChangeDragStateFromOldState != nil
}

// MapViewClusterAnnotationForMemberAnnotations implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewClusterAnnotationForMemberAnnotations(mapView IMKMapView, memberAnnotations []objc.ID) MKClusterAnnotation {
	if d._MapViewClusterAnnotationForMemberAnnotations != nil {
		return d._MapViewClusterAnnotationForMemberAnnotations(mapView, memberAnnotations)
	}
	var zero MKClusterAnnotation
	return zero
}

// HasMapViewClusterAnnotationForMemberAnnotations returns true if a handler for MapViewClusterAnnotationForMemberAnnotations has been set.
func (d *MKMapViewDelegate) HasMapViewClusterAnnotationForMemberAnnotations() bool {
	return d._MapViewClusterAnnotationForMemberAnnotations != nil
}

// MapViewDidAddAnnotationViews implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidAddAnnotationViews(mapView IMKMapView, views []objc.IObject /* cross-framework: MKAnnotationView */) {
	if d._MapViewDidAddAnnotationViews != nil {
		d._MapViewDidAddAnnotationViews(mapView, views)
	}
}

// HasMapViewDidAddAnnotationViews returns true if a handler for MapViewDidAddAnnotationViews has been set.
func (d *MKMapViewDelegate) HasMapViewDidAddAnnotationViews() bool {
	return d._MapViewDidAddAnnotationViews != nil
}

// MapViewDidAddOverlayRenderers implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidAddOverlayRenderers(mapView IMKMapView, renderers []MKOverlayRenderer) {
	if d._MapViewDidAddOverlayRenderers != nil {
		d._MapViewDidAddOverlayRenderers(mapView, renderers)
	}
}

// HasMapViewDidAddOverlayRenderers returns true if a handler for MapViewDidAddOverlayRenderers has been set.
func (d *MKMapViewDelegate) HasMapViewDidAddOverlayRenderers() bool {
	return d._MapViewDidAddOverlayRenderers != nil
}

// MapViewDidAddOverlayViews implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidAddOverlayViews(mapView IMKMapView, overlayViews objc.IObject /* cross-framework: NSArray */) {
	if d._MapViewDidAddOverlayViews != nil {
		d._MapViewDidAddOverlayViews(mapView, overlayViews)
	}
}

// HasMapViewDidAddOverlayViews returns true if a handler for MapViewDidAddOverlayViews has been set.
func (d *MKMapViewDelegate) HasMapViewDidAddOverlayViews() bool {
	return d._MapViewDidAddOverlayViews != nil
}

// MapViewDidChangeUserTrackingModeAnimated implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidChangeUserTrackingModeAnimated(mapView IMKMapView, mode MKUserTrackingMode, animated bool) {
	if d._MapViewDidChangeUserTrackingModeAnimated != nil {
		d._MapViewDidChangeUserTrackingModeAnimated(mapView, mode, animated)
	}
}

// HasMapViewDidChangeUserTrackingModeAnimated returns true if a handler for MapViewDidChangeUserTrackingModeAnimated has been set.
func (d *MKMapViewDelegate) HasMapViewDidChangeUserTrackingModeAnimated() bool {
	return d._MapViewDidChangeUserTrackingModeAnimated != nil
}

// MapViewDidDeselectAnnotation implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidDeselectAnnotation(mapView IMKMapView, annotation unsafe.Pointer) {
	if d._MapViewDidDeselectAnnotation != nil {
		d._MapViewDidDeselectAnnotation(mapView, annotation)
	}
}

// HasMapViewDidDeselectAnnotation returns true if a handler for MapViewDidDeselectAnnotation has been set.
func (d *MKMapViewDelegate) HasMapViewDidDeselectAnnotation() bool {
	return d._MapViewDidDeselectAnnotation != nil
}

// MapViewDidDeselectAnnotationView implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidDeselectAnnotationView(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */) {
	if d._MapViewDidDeselectAnnotationView != nil {
		d._MapViewDidDeselectAnnotationView(mapView, view)
	}
}

// HasMapViewDidDeselectAnnotationView returns true if a handler for MapViewDidDeselectAnnotationView has been set.
func (d *MKMapViewDelegate) HasMapViewDidDeselectAnnotationView() bool {
	return d._MapViewDidDeselectAnnotationView != nil
}

// MapViewDidFailToLocateUserWithError implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidFailToLocateUserWithError(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */) {
	if d._MapViewDidFailToLocateUserWithError != nil {
		d._MapViewDidFailToLocateUserWithError(mapView, error_)
	}
}

// HasMapViewDidFailToLocateUserWithError returns true if a handler for MapViewDidFailToLocateUserWithError has been set.
func (d *MKMapViewDelegate) HasMapViewDidFailToLocateUserWithError() bool {
	return d._MapViewDidFailToLocateUserWithError != nil
}

// MapViewDidSelectAnnotationView implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidSelectAnnotationView(mapView IMKMapView, view objc.IObject /* cross-framework: MKAnnotationView */) {
	if d._MapViewDidSelectAnnotationView != nil {
		d._MapViewDidSelectAnnotationView(mapView, view)
	}
}

// HasMapViewDidSelectAnnotationView returns true if a handler for MapViewDidSelectAnnotationView has been set.
func (d *MKMapViewDelegate) HasMapViewDidSelectAnnotationView() bool {
	return d._MapViewDidSelectAnnotationView != nil
}

// MapViewDidSelectAnnotation implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidSelectAnnotation(mapView IMKMapView, annotation unsafe.Pointer) {
	if d._MapViewDidSelectAnnotation != nil {
		d._MapViewDidSelectAnnotation(mapView, annotation)
	}
}

// HasMapViewDidSelectAnnotation returns true if a handler for MapViewDidSelectAnnotation has been set.
func (d *MKMapViewDelegate) HasMapViewDidSelectAnnotation() bool {
	return d._MapViewDidSelectAnnotation != nil
}

// MapViewDidUpdateUserLocation implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidUpdateUserLocation(mapView IMKMapView, userLocation IMKUserLocation) {
	if d._MapViewDidUpdateUserLocation != nil {
		d._MapViewDidUpdateUserLocation(mapView, userLocation)
	}
}

// HasMapViewDidUpdateUserLocation returns true if a handler for MapViewDidUpdateUserLocation has been set.
func (d *MKMapViewDelegate) HasMapViewDidUpdateUserLocation() bool {
	return d._MapViewDidUpdateUserLocation != nil
}

// MapViewRegionDidChangeAnimated implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewRegionDidChangeAnimated(mapView IMKMapView, animated bool) {
	if d._MapViewRegionDidChangeAnimated != nil {
		d._MapViewRegionDidChangeAnimated(mapView, animated)
	}
}

// HasMapViewRegionDidChangeAnimated returns true if a handler for MapViewRegionDidChangeAnimated has been set.
func (d *MKMapViewDelegate) HasMapViewRegionDidChangeAnimated() bool {
	return d._MapViewRegionDidChangeAnimated != nil
}

// MapViewRegionWillChangeAnimated implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewRegionWillChangeAnimated(mapView IMKMapView, animated bool) {
	if d._MapViewRegionWillChangeAnimated != nil {
		d._MapViewRegionWillChangeAnimated(mapView, animated)
	}
}

// HasMapViewRegionWillChangeAnimated returns true if a handler for MapViewRegionWillChangeAnimated has been set.
func (d *MKMapViewDelegate) HasMapViewRegionWillChangeAnimated() bool {
	return d._MapViewRegionWillChangeAnimated != nil
}

// MapViewRendererForOverlay implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewRendererForOverlay(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayRenderer {
	if d._MapViewRendererForOverlay != nil {
		return d._MapViewRendererForOverlay(mapView, overlay)
	}
	var zero MKOverlayRenderer
	return zero
}

// HasMapViewRendererForOverlay returns true if a handler for MapViewRendererForOverlay has been set.
func (d *MKMapViewDelegate) HasMapViewRendererForOverlay() bool {
	return d._MapViewRendererForOverlay != nil
}

// MapViewSelectionAccessoryForAnnotation implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewSelectionAccessoryForAnnotation(mapView IMKMapView, annotation unsafe.Pointer) MKSelectionAccessory {
	if d._MapViewSelectionAccessoryForAnnotation != nil {
		return d._MapViewSelectionAccessoryForAnnotation(mapView, annotation)
	}
	var zero MKSelectionAccessory
	return zero
}

// HasMapViewSelectionAccessoryForAnnotation returns true if a handler for MapViewSelectionAccessoryForAnnotation has been set.
func (d *MKMapViewDelegate) HasMapViewSelectionAccessoryForAnnotation() bool {
	return d._MapViewSelectionAccessoryForAnnotation != nil
}

// MapViewViewForOverlay implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewViewForOverlay(mapView IMKMapView, overlay unsafe.Pointer) MKOverlayView {
	if d._MapViewViewForOverlay != nil {
		return d._MapViewViewForOverlay(mapView, overlay)
	}
	var zero MKOverlayView
	return zero
}

// HasMapViewViewForOverlay returns true if a handler for MapViewViewForOverlay has been set.
func (d *MKMapViewDelegate) HasMapViewViewForOverlay() bool {
	return d._MapViewViewForOverlay != nil
}

// MapViewViewForAnnotation implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewViewForAnnotation(mapView IMKMapView, annotation unsafe.Pointer) MKAnnotationView {
	if d._MapViewViewForAnnotation != nil {
		return d._MapViewViewForAnnotation(mapView, annotation)
	}
	var zero MKAnnotationView
	return zero
}

// HasMapViewViewForAnnotation returns true if a handler for MapViewViewForAnnotation has been set.
func (d *MKMapViewDelegate) HasMapViewViewForAnnotation() bool {
	return d._MapViewViewForAnnotation != nil
}

// MapViewDidChangeVisibleRegion implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidChangeVisibleRegion(mapView IMKMapView) {
	if d._MapViewDidChangeVisibleRegion != nil {
		d._MapViewDidChangeVisibleRegion(mapView)
	}
}

// HasMapViewDidChangeVisibleRegion returns true if a handler for MapViewDidChangeVisibleRegion has been set.
func (d *MKMapViewDelegate) HasMapViewDidChangeVisibleRegion() bool {
	return d._MapViewDidChangeVisibleRegion != nil
}

// MapViewDidFailLoadingMapWithError implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidFailLoadingMapWithError(mapView IMKMapView, error_ objc.IObject /* cross-framework: Error */) {
	if d._MapViewDidFailLoadingMapWithError != nil {
		d._MapViewDidFailLoadingMapWithError(mapView, error_)
	}
}

// HasMapViewDidFailLoadingMapWithError returns true if a handler for MapViewDidFailLoadingMapWithError has been set.
func (d *MKMapViewDelegate) HasMapViewDidFailLoadingMapWithError() bool {
	return d._MapViewDidFailLoadingMapWithError != nil
}

// MapViewDidFinishLoadingMap implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidFinishLoadingMap(mapView IMKMapView) {
	if d._MapViewDidFinishLoadingMap != nil {
		d._MapViewDidFinishLoadingMap(mapView)
	}
}

// HasMapViewDidFinishLoadingMap returns true if a handler for MapViewDidFinishLoadingMap has been set.
func (d *MKMapViewDelegate) HasMapViewDidFinishLoadingMap() bool {
	return d._MapViewDidFinishLoadingMap != nil
}

// MapViewDidFinishRenderingMapFullyRendered implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidFinishRenderingMapFullyRendered(mapView IMKMapView, fullyRendered bool) {
	if d._MapViewDidFinishRenderingMapFullyRendered != nil {
		d._MapViewDidFinishRenderingMapFullyRendered(mapView, fullyRendered)
	}
}

// HasMapViewDidFinishRenderingMapFullyRendered returns true if a handler for MapViewDidFinishRenderingMapFullyRendered has been set.
func (d *MKMapViewDelegate) HasMapViewDidFinishRenderingMapFullyRendered() bool {
	return d._MapViewDidFinishRenderingMapFullyRendered != nil
}

// MapViewDidStopLocatingUser implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewDidStopLocatingUser(mapView IMKMapView) {
	if d._MapViewDidStopLocatingUser != nil {
		d._MapViewDidStopLocatingUser(mapView)
	}
}

// HasMapViewDidStopLocatingUser returns true if a handler for MapViewDidStopLocatingUser has been set.
func (d *MKMapViewDelegate) HasMapViewDidStopLocatingUser() bool {
	return d._MapViewDidStopLocatingUser != nil
}

// MapViewWillStartLoadingMap implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewWillStartLoadingMap(mapView IMKMapView) {
	if d._MapViewWillStartLoadingMap != nil {
		d._MapViewWillStartLoadingMap(mapView)
	}
}

// HasMapViewWillStartLoadingMap returns true if a handler for MapViewWillStartLoadingMap has been set.
func (d *MKMapViewDelegate) HasMapViewWillStartLoadingMap() bool {
	return d._MapViewWillStartLoadingMap != nil
}

// MapViewWillStartLocatingUser implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewWillStartLocatingUser(mapView IMKMapView) {
	if d._MapViewWillStartLocatingUser != nil {
		d._MapViewWillStartLocatingUser(mapView)
	}
}

// HasMapViewWillStartLocatingUser returns true if a handler for MapViewWillStartLocatingUser has been set.
func (d *MKMapViewDelegate) HasMapViewWillStartLocatingUser() bool {
	return d._MapViewWillStartLocatingUser != nil
}

// MapViewWillStartRenderingMap implements the PMKMapViewDelegate interface.
func (d *MKMapViewDelegate) MapViewWillStartRenderingMap(mapView IMKMapView) {
	if d._MapViewWillStartRenderingMap != nil {
		d._MapViewWillStartRenderingMap(mapView)
	}
}

// HasMapViewWillStartRenderingMap returns true if a handler for MapViewWillStartRenderingMap has been set.
func (d *MKMapViewDelegate) HasMapViewWillStartRenderingMap() bool {
	return d._MapViewWillStartRenderingMap != nil
}
