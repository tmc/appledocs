// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DDMatchEmailAddress] class.
var (
	DDMatchEmailAddressClass     _DDMatchEmailAddressClass
	DDMatchEmailAddressClassOnce sync.Once
)

func getDDMatchEmailAddressClass() _DDMatchEmailAddressClass {
	DDMatchEmailAddressClassOnce.Do(func() {
		DDMatchEmailAddressClass = _DDMatchEmailAddressClass{objc.GetClass("DDMatchEmailAddress")}
	})
	return DDMatchEmailAddressClass
}

type _DDMatchEmailAddressClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchEmailAddress] class.
type IDDMatchEmailAddress interface {
	IDDMatch
}

// An object that contains an email address that the data detection system matches.
//
// The DataDetection framework returns an email match in a object, which includes an email address, and optionally a label that categorizes the email address.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress
type DDMatchEmailAddress struct {
	DDMatch
}

// DDMatchEmailAddressFrom constructs a [DDMatchEmailAddress] from an unsafe.Pointer.
//
// An object that contains an email address that the data detection system matches.
func DDMatchEmailAddressFrom(ptr unsafe.Pointer) DDMatchEmailAddress {
	return DDMatchEmailAddress{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchEmailAddressClass) Alloc() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchEmailAddressClass) New() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchEmailAddress) Init() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchEmailAddress) Autorelease() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchEmailAddress creates a new DDMatchEmailAddress instance.
func NewDDMatchEmailAddress() DDMatchEmailAddress {
	return getDDMatchEmailAddressClass().New()
}


// A string that represents an email address.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress/emailAddress
func (d_ DDMatchEmailAddress) EmailAddress() string {
	rv := objc.Send[string](d_.ID, objc.Sel("emailAddress"))
	return rv
}

// A string that categorizes an email address, such as Home or Work.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress/label
func (d_ DDMatchEmailAddress) Label() string {
	rv := objc.Send[string](d_.ID, objc.Sel("label"))
	return rv
}



