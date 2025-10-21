// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [DDMatchPhoneNumber] class.
var (
	DDMatchPhoneNumberClass     _DDMatchPhoneNumberClass
	DDMatchPhoneNumberClassOnce sync.Once
)

func getDDMatchPhoneNumberClass() _DDMatchPhoneNumberClass {
	DDMatchPhoneNumberClassOnce.Do(func() {
		DDMatchPhoneNumberClass = _DDMatchPhoneNumberClass{objc.GetClass("DDMatchPhoneNumber")}
	})
	return DDMatchPhoneNumberClass
}

type _DDMatchPhoneNumberClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchPhoneNumber] class.
type IDDMatchPhoneNumber interface {
	IDDMatch
}

// An object that contains a phone number that the data detection system matches.
//
// The DataDetection framework returns a phone number match in a object, which contains a phone number, and optionally a label that categorizes the phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber
type DDMatchPhoneNumber struct {
	DDMatch
}

// DDMatchPhoneNumberFrom constructs a [DDMatchPhoneNumber] from an unsafe.Pointer.
//
// An object that contains a phone number that the data detection system matches.
func DDMatchPhoneNumberFrom(ptr unsafe.Pointer) DDMatchPhoneNumber {
	return DDMatchPhoneNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchPhoneNumberClass) Alloc() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchPhoneNumberClass) New() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchPhoneNumber) Init() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchPhoneNumber) Autorelease() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchPhoneNumber creates a new DDMatchPhoneNumber instance.
func NewDDMatchPhoneNumber() DDMatchPhoneNumber {
	return getDDMatchPhoneNumberClass().New()
}


// A string that categorizes a phone number, such as Home or Work.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber/label
func (d_ DDMatchPhoneNumber) Label() appkit.string {
	rv := objc.Send[appkit.string](d_.ID, objc.Sel("label"))
	return rv
}

// A string that represents a phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber/phoneNumber
func (d_ DDMatchPhoneNumber) PhoneNumber() appkit.string {
	rv := objc.Send[appkit.string](d_.ID, objc.Sel("phoneNumber"))
	return rv
}



