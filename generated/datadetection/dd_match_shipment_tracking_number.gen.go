// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DDMatchShipmentTrackingNumber] class.
var (
	DDMatchShipmentTrackingNumberClass     _DDMatchShipmentTrackingNumberClass
	DDMatchShipmentTrackingNumberClassOnce sync.Once
)

func getDDMatchShipmentTrackingNumberClass() _DDMatchShipmentTrackingNumberClass {
	DDMatchShipmentTrackingNumberClassOnce.Do(func() {
		DDMatchShipmentTrackingNumberClass = _DDMatchShipmentTrackingNumberClass{objc.GetClass("DDMatchShipmentTrackingNumber")}
	})
	return DDMatchShipmentTrackingNumberClass
}

type _DDMatchShipmentTrackingNumberClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchShipmentTrackingNumber] class.
type IDDMatchShipmentTrackingNumber interface {
	IDDMatch
}

// An object that contains parcel tracking information that the data detection system matches.
//
// The DataDetection framework returns a shipment tracking number match in a object, which contains a carrier name and tracking identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber
type DDMatchShipmentTrackingNumber struct {
	DDMatch
}

// DDMatchShipmentTrackingNumberFrom constructs a [DDMatchShipmentTrackingNumber] from an unsafe.Pointer.
//
// An object that contains parcel tracking information that the data detection system matches.
func DDMatchShipmentTrackingNumberFrom(ptr unsafe.Pointer) DDMatchShipmentTrackingNumber {
	return DDMatchShipmentTrackingNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchShipmentTrackingNumberClass) Alloc() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchShipmentTrackingNumberClass) New() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchShipmentTrackingNumber) Init() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchShipmentTrackingNumber) Autorelease() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchShipmentTrackingNumber creates a new DDMatchShipmentTrackingNumber instance.
func NewDDMatchShipmentTrackingNumber() DDMatchShipmentTrackingNumber {
	return getDDMatchShipmentTrackingNumberClass().New()
}


// The name of a parcel carrier.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber/carrier
func (d_ DDMatchShipmentTrackingNumber) Carrier() string {
	rv := objc.Send[string](d_.ID, objc.Sel("carrier"))
	return rv
}

// A string that represents a carrier’s tracking identifier for a parcel.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber/trackingNumber
func (d_ DDMatchShipmentTrackingNumber) TrackingNumber() string {
	rv := objc.Send[string](d_.ID, objc.Sel("trackingNumber"))
	return rv
}


