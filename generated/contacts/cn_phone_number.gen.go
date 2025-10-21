// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// An immutable object representing a phone number for a contact.
//
// objects are thread-safe, and you may access their properties from any thread of your app.
//
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




// Returns a new phone number object initialized with the specified phone number string.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/init(stringValue:)
func NewCNPhoneNumberWithStringValue(string_ appkit.string) CNPhoneNumber {
	instance := getCNPhoneNumberClass().Alloc()
	rv := objc.Send[CNPhoneNumber](instance.ID, objc.Sel("initWithStringValue:"), string_)
	rv.Autorelease()
	return rv
}


// Returns a new phone number object initialized with the specified phone number string.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/phoneNumberWithStringValue:
func (cc _CNPhoneNumberClass) PhoneNumberWithStringValue(stringValue appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("phoneNumberWithStringValue:"), stringValue)
	return rv
}

// The string value of the phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/stringValue
func (c_ CNPhoneNumber) StringValue() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("stringValue"))
	return rv
}

// A phone numbers of a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactphonenumberskey
func (c_ CNPhoneNumber) CNContactPhoneNumbersKey() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNContactPhoneNumbersKey"))
	return rv
}


