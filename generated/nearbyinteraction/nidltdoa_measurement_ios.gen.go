//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NIDLTDOAMeasurement


// iOS-only properties

// A value that uniquely identifies an anchor in a tracked area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/address
func (n_ NIDLTDOAMeasurement) Address() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("address"))
	return rv
}

// The drift, as a ratio, across the frequencies of the receiver and the anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/carrierFrequencyOffset
func (n_ NIDLTDOAMeasurement) CarrierFrequencyOffset() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("carrierFrequencyOffset"))
	return rv
}

// A triplet that represents the location in 3D space of the anchor that provides the measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/coordinates
func (n_ NIDLTDOAMeasurement) Coordinates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("coordinates"))
	return rv
}

// The type of coordinate system that the measurement conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/coordinatesType
func (n_ NIDLTDOAMeasurement) CoordinatesType() NIDLTDOACoordinatesType {
	rv := objc.Send[NIDLTDOACoordinatesType](n_.ID, objc.Sel("coordinatesType"))
	return rv
}

// The type of anchor message that the measurement derives from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/measurementType
func (n_ NIDLTDOAMeasurement) MeasurementType() NIDLTDOAMeasurementType {
	rv := objc.Send[NIDLTDOAMeasurementType](n_.ID, objc.Sel("measurementType"))
	return rv
}

// A timestamp, in seconds, for the time that the device receives the measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/receiveTime
func (n_ NIDLTDOAMeasurement) ReceiveTime() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("receiveTime"))
	return rv
}

// A value that represents the signal strength, in dBm, to the anchor that provides the measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/signalStrength
func (n_ NIDLTDOAMeasurement) SignalStrength() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("signalStrength"))
	return rv
}

// A timestamp, in seconds, for the elapsed message transmission time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement/transmitTime
func (n_ NIDLTDOAMeasurement) TransmitTime() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("transmitTime"))
	return rv
}







