// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMKMapItemDetailViewControllerDelegate is the MKMapItemDetailViewControllerDelegate protocol interface.
//
// The methods that you use to receive events from an associated map view controller.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - visionOS 2.0+
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKMapItemDetailViewControllerDelegate
type PMKMapItemDetailViewControllerDelegate interface {
	// Required methods
	MapItemDetailViewControllerDidFinish(detailViewController IMKMapItemDetailViewController)/* debug [protocol_interface/required_method]: MapItemDetailViewControllerDidFinish */
}

// MKMapItemDetailViewControllerDelegate is a delegate implementation builder for the PMKMapItemDetailViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MKMapItemDetailViewControllerDelegate struct {
	_MapItemDetailViewControllerDidFinish func(detailViewController IMKMapItemDetailViewController)
}

// SetMapItemDetailViewControllerDidFinish sets the handler for the MapItemDetailViewControllerDidFinish delegate method.
//
// Informs the delegate when a person dismissed the view controller.
func (d *MKMapItemDetailViewControllerDelegate) SetMapItemDetailViewControllerDidFinish(f func(detailViewController IMKMapItemDetailViewController)) {
	d._MapItemDetailViewControllerDidFinish = f
}

// MapItemDetailViewControllerDidFinish implements the PMKMapItemDetailViewControllerDelegate interface.
func (d *MKMapItemDetailViewControllerDelegate) MapItemDetailViewControllerDidFinish(detailViewController IMKMapItemDetailViewController) {
	if d._MapItemDetailViewControllerDidFinish != nil {
		d._MapItemDetailViewControllerDidFinish(detailViewController)
	}
}

// HasMapItemDetailViewControllerDidFinish returns true if a handler for MapItemDetailViewControllerDidFinish has been set.
func (d *MKMapItemDetailViewControllerDelegate) HasMapItemDetailViewControllerDidFinish() bool {
	return d._MapItemDetailViewControllerDidFinish != nil
}
