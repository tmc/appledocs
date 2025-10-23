// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNPhoneNumber] class.
var (
	CNPhoneNumberClass     _CNPhoneNumberClass
	CNPhoneNumberClassOnce sync.Once
)

func getCNPhoneNumberClass() _CNPhoneNumberClass {
	CNPhoneNumberClassOnce.Do(func() {
		CNPhoneNumberClass = _CNPhoneNumberClass{objc.GetClass("CNPhoneNumber")}
	})
	return CNPhoneNumberClass
}

type _CNPhoneNumberClass struct {
	class objc.Class
}

// An interface definition for the [CNPhoneNumber] class.
type ICNPhoneNumber interface {
	objectivec.IObject
	CNContactPhoneNumbersKey() string
	StringValue() string
	SetStringValue(value string)
}

// An immutable object representing a phone number for a contact.
//
// objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object representing a phone number for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber
type CNPhoneNumber struct {
	objectivec.Object
}

// CNPhoneNumberFrom constructs a [CNPhoneNumber] from an unsafe.Pointer.
//
// An immutable object representing a phone number for a contact.
func CNPhoneNumberFrom(ptr unsafe.Pointer) CNPhoneNumber {
	return CNPhoneNumber{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNPhoneNumberClass) Alloc() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNPhoneNumberClass) New() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNPhoneNumber) Init() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNPhoneNumber) Autorelease() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNPhoneNumber creates a new CNPhoneNumber instance.
func NewCNPhoneNumber() CNPhoneNumber {
	return getCNPhoneNumberClass().New()
}




// A phone numbers of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactphonenumberskey
func (c_ CNPhoneNumber) CNContactPhoneNumbersKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContactPhoneNumbersKey"))
	return rv
}


// The string value of the phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber/stringvalue
func (c_ CNPhoneNumber) StringValue() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stringValue"))
	return rv
}


// The string value of the phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber/stringvalue
func (c_ CNPhoneNumber) SetStringValue(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), objc.String(value))
}


