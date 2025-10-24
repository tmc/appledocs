// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PMKReverseGeocoderDelegate is the MKReverseGeocoderDelegate protocol interface.
//
// Defines the interface for receiving messages from an   object.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 3.0+ (Deprecated in 5.0)
//   - iPadOS 3.0+ (Deprecated in 5.0)
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKReverseGeocoderDelegate
type PMKReverseGeocoderDelegate interface {
	// Required methods
	ReverseGeocoderDidFailWithError(geocoder IMKReverseGeocoder, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: ReverseGeocoderDidFailWithError */
	ReverseGeocoderDidFindPlacemark(geocoder IMKReverseGeocoder, placemark IMKPlacemark)/* debug [protocol_interface/required_method]: ReverseGeocoderDidFindPlacemark */
}

// MKReverseGeocoderDelegate is a delegate implementation builder for the PMKReverseGeocoderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MKReverseGeocoderDelegate struct {
	_ReverseGeocoderDidFailWithError func(geocoder IMKReverseGeocoder, error_ objc.IObject /* cross-framework: Error */)
	_ReverseGeocoderDidFindPlacemark func(geocoder IMKReverseGeocoder, placemark IMKPlacemark)
}

// SetReverseGeocoderDidFailWithError sets the handler for the ReverseGeocoderDidFailWithError delegate method.
//
// Tells the delegate that the specified reverse geocoder failed to obtain information about its coordinate.
func (d *MKReverseGeocoderDelegate) SetReverseGeocoderDidFailWithError(f func(geocoder IMKReverseGeocoder, error_ objc.IObject /* cross-framework: Error */)) {
	d._ReverseGeocoderDidFailWithError = f
}

// SetReverseGeocoderDidFindPlacemark sets the handler for the ReverseGeocoderDidFindPlacemark delegate method.
//
// Tells the delegate that a reverse geocoder successfully obtained placemark information for its coordinate.
func (d *MKReverseGeocoderDelegate) SetReverseGeocoderDidFindPlacemark(f func(geocoder IMKReverseGeocoder, placemark IMKPlacemark)) {
	d._ReverseGeocoderDidFindPlacemark = f
}

// ReverseGeocoderDidFailWithError implements the PMKReverseGeocoderDelegate interface.
func (d *MKReverseGeocoderDelegate) ReverseGeocoderDidFailWithError(geocoder IMKReverseGeocoder, error_ objc.IObject /* cross-framework: Error */) {
	if d._ReverseGeocoderDidFailWithError != nil {
		d._ReverseGeocoderDidFailWithError(geocoder, error_)
	}
}

// HasReverseGeocoderDidFailWithError returns true if a handler for ReverseGeocoderDidFailWithError has been set.
func (d *MKReverseGeocoderDelegate) HasReverseGeocoderDidFailWithError() bool {
	return d._ReverseGeocoderDidFailWithError != nil
}

// ReverseGeocoderDidFindPlacemark implements the PMKReverseGeocoderDelegate interface.
func (d *MKReverseGeocoderDelegate) ReverseGeocoderDidFindPlacemark(geocoder IMKReverseGeocoder, placemark IMKPlacemark) {
	if d._ReverseGeocoderDidFindPlacemark != nil {
		d._ReverseGeocoderDidFindPlacemark(geocoder, placemark)
	}
}

// HasReverseGeocoderDidFindPlacemark returns true if a handler for ReverseGeocoderDidFindPlacemark has been set.
func (d *MKReverseGeocoderDelegate) HasReverseGeocoderDidFindPlacemark() bool {
	return d._ReverseGeocoderDidFindPlacemark != nil
}
