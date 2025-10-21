// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NIDLTDOAMeasurement] class.
var (
	NIDLTDOAMeasurementClass     _NIDLTDOAMeasurementClass
	NIDLTDOAMeasurementClassOnce sync.Once
)

func getNIDLTDOAMeasurementClass() _NIDLTDOAMeasurementClass {
	NIDLTDOAMeasurementClassOnce.Do(func() {
		NIDLTDOAMeasurementClass = _NIDLTDOAMeasurementClass{objc.GetClass("NIDLTDOAMeasurement")}
	})
	return NIDLTDOAMeasurementClass
}

type _NIDLTDOAMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [NIDLTDOAMeasurement] class.
type INIDLTDOAMeasurement interface {
	objectivec.IObject
}

// Information from a Downlink Time-Difference-of-Arrival anchor that you use to derive a range estimate.
//
// Your app runs on a receiver device that fields messages from nearby physical base stations, or . The framework processes the messages into instances of this class and provides them to your app through the callback. Your app analyzes the measurements to calculate the receiver’s position relative to the anchors in the tracked area. Only sessions that run a receive Downlink Time-Difference-of-Arrival measurements.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement
type NIDLTDOAMeasurement struct {
	objectivec.Object
}

// NIDLTDOAMeasurementFrom constructs a [NIDLTDOAMeasurement] from an unsafe.Pointer.
//
// Information from a Downlink Time-Difference-of-Arrival anchor that you use to derive a range estimate.
func NIDLTDOAMeasurementFrom(ptr unsafe.Pointer) NIDLTDOAMeasurement {
	return NIDLTDOAMeasurement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NIDLTDOAMeasurementClass) Alloc() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NIDLTDOAMeasurementClass) New() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIDLTDOAMeasurement) Init() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIDLTDOAMeasurement) Autorelease() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIDLTDOAMeasurement creates a new NIDLTDOAMeasurement instance.
func NewNIDLTDOAMeasurement() NIDLTDOAMeasurement {
	return getNIDLTDOAMeasurementClass().New()
}


// A value that uniquely identifies an anchor in a tracked area.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/address
func (n_ NIDLTDOAMeasurement) Address() int {
	rv := objc.Send[int](n_.ID, objc.Sel("address"))
	return rv
}


// SetAddress sets the value of the address property.
// A value that uniquely identifies an anchor in a tracked area.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/address
func (n_ NIDLTDOAMeasurement) SetAddress(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAddress:"), value)
}

// The drift, as a ratio, across the frequencies of the receiver and the anchor.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/carrierfrequencyoffset
func (n_ NIDLTDOAMeasurement) CarrierFrequencyOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("carrierFrequencyOffset"))
	return rv
}


// SetCarrierFrequencyOffset sets the value of the carrierFrequencyOffset property.
// The drift, as a ratio, across the frequencies of the receiver and the anchor.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/carrierfrequencyoffset
func (n_ NIDLTDOAMeasurement) SetCarrierFrequencyOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCarrierFrequencyOffset:"), value)
}

// A triplet that represents the location in 3D space of the anchor that provides the measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/coordinates
func (n_ NIDLTDOAMeasurement) Coordinates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("coordinates"))
	return rv
}


// SetCoordinates sets the value of the coordinates property.
// A triplet that represents the location in 3D space of the anchor that provides the measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/coordinates
func (n_ NIDLTDOAMeasurement) SetCoordinates(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCoordinates:"), value)
}

// The type of coordinate system that the measurement conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/coordinatestype
func (n_ NIDLTDOAMeasurement) CoordinatesType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("coordinatesType"))
	return rv
}


// SetCoordinatesType sets the value of the coordinatesType property.
// The type of coordinate system that the measurement conforms to.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/coordinatestype
func (n_ NIDLTDOAMeasurement) SetCoordinatesType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCoordinatesType:"), value)
}

// The type of anchor message that the measurement derives from.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/measurementtype
func (n_ NIDLTDOAMeasurement) MeasurementType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("measurementType"))
	return rv
}


// SetMeasurementType sets the value of the measurementType property.
// The type of anchor message that the measurement derives from.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/measurementtype
func (n_ NIDLTDOAMeasurement) SetMeasurementType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMeasurementType:"), value)
}

// A timestamp, in seconds, for the time that the device receives the measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/receivetime
func (n_ NIDLTDOAMeasurement) ReceiveTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("receiveTime"))
	return rv
}


// SetReceiveTime sets the value of the receiveTime property.
// A timestamp, in seconds, for the time that the device receives the measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/receivetime
func (n_ NIDLTDOAMeasurement) SetReceiveTime(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setReceiveTime:"), value)
}

// A value that represents the signal strength, in dBm, to the anchor that provides the measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/signalstrength
func (n_ NIDLTDOAMeasurement) SignalStrength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("signalStrength"))
	return rv
}


// SetSignalStrength sets the value of the signalStrength property.
// A value that represents the signal strength, in dBm, to the anchor that provides the measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/signalstrength
func (n_ NIDLTDOAMeasurement) SetSignalStrength(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSignalStrength:"), value)
}

// A timestamp, in seconds, for the elapsed message transmission time.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/transmittime
func (n_ NIDLTDOAMeasurement) TransmitTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("transmitTime"))
	return rv
}


// SetTransmitTime sets the value of the transmitTime property.
// A timestamp, in seconds, for the elapsed message transmission time.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidltdoameasurement/transmittime
func (n_ NIDLTDOAMeasurement) SetTransmitTime(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTransmitTime:"), value)
}



