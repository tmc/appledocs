// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DDMatchFlightNumber] class.
var (
	DDMatchFlightNumberClass     _DDMatchFlightNumberClass
	DDMatchFlightNumberClassOnce sync.Once
)

func getDDMatchFlightNumberClass() _DDMatchFlightNumberClass {
	DDMatchFlightNumberClassOnce.Do(func() {
		DDMatchFlightNumberClass = _DDMatchFlightNumberClass{objc.GetClass("DDMatchFlightNumber")}
	})
	return DDMatchFlightNumberClass
}

type _DDMatchFlightNumberClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchFlightNumber] class.
type IDDMatchFlightNumber interface {
	IDDMatch
	Airline() string
	FlightNumber() string
}

// An object that contains a flight number that the data detection system matches.
//
// The DataDetection framework returns a flight number match in a object, which contains an airline name and flight number.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber
type DDMatchFlightNumber struct {
	DDMatch
}

// DDMatchFlightNumberFrom constructs a [DDMatchFlightNumber] from an unsafe.Pointer.
//
// An object that contains a flight number that the data detection system matches.
func DDMatchFlightNumberFrom(ptr unsafe.Pointer) DDMatchFlightNumber {
	return DDMatchFlightNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchFlightNumberClass) Alloc() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchFlightNumberClass) New() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchFlightNumber) Init() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchFlightNumber) Autorelease() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchFlightNumber creates a new DDMatchFlightNumber instance.
func NewDDMatchFlightNumber() DDMatchFlightNumber {
	return getDDMatchFlightNumberClass().New()
}


// The name of an airline.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber/airline
func (d_ DDMatchFlightNumber) Airline() string {
	rv := objc.Send[string](d_.ID, objc.Sel("airline"))
	return rv
}

// A string that represents a flight number.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber/flightNumber
func (d_ DDMatchFlightNumber) FlightNumber() string {
	rv := objc.Send[string](d_.ID, objc.Sel("flightNumber"))
	return rv
}



