// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNPostalAddressFormatter] class.
var (
	CNPostalAddressFormatterClass     _CNPostalAddressFormatterClass
	CNPostalAddressFormatterClassOnce sync.Once
)

func getCNPostalAddressFormatterClass() _CNPostalAddressFormatterClass {
	CNPostalAddressFormatterClassOnce.Do(func() {
		CNPostalAddressFormatterClass = _CNPostalAddressFormatterClass{objc.GetClass("CNPostalAddressFormatter")}
	})
	return CNPostalAddressFormatterClass
}

type _CNPostalAddressFormatterClass struct {
	class objc.Class
}

// An interface definition for the [CNPostalAddressFormatter] class.
type ICNPostalAddressFormatter interface {
	foundation.IFormatter
	AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objectivec.IObject) foundation.AttributedString
	StringFromPostalAddress(postalAddress ICNPostalAddress) foundation.String
	Style() CNPostalAddressFormatterStyle
	SetStyle(value CNPostalAddressFormatterStyle)
	CNPostalAddressCityKey() string
	CNPostalAddressCountryKey() string
	CNPostalAddressISOCountryCodeKey() string
	CNPostalAddressLocalizedPropertyNameAttribute() string
	CNPostalAddressPostalCodeKey() string
	CNPostalAddressPropertyAttribute() string
	CNPostalAddressStateKey() string
	CNPostalAddressStreetKey() string
	CNPostalAddressSubAdministrativeAreaKey() string
	CNPostalAddressSubLocalityKey() string
}

// An object that you use to format a contact’s postal addresses.
//
// A object handles international formatting of postal addresses. It is recommended that you create an instance of this class when formatting many postal addresses, and use the instance methods; otherwise use the class methods.


// An object that you use to format a contact’s postal addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter

type CNPostalAddressFormatter struct {
	foundation.Formatter
}

// CNPostalAddressFormatterFrom constructs a [CNPostalAddressFormatter] from an unsafe.Pointer.
//
// An object that you use to format a contact’s postal addresses.
func CNPostalAddressFormatterFrom(ptr unsafe.Pointer) CNPostalAddressFormatter {
	return CNPostalAddressFormatter{
		Formatter: foundation.FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNPostalAddressFormatterClass) Alloc() CNPostalAddressFormatter {
	rv := objc.Send[CNPostalAddressFormatter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNPostalAddressFormatterClass) New() CNPostalAddressFormatter {
	rv := objc.Send[CNPostalAddressFormatter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNPostalAddressFormatter) Init() CNPostalAddressFormatter {
	rv := objc.Send[CNPostalAddressFormatter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNPostalAddressFormatter) Autorelease() CNPostalAddressFormatter {
	rv := objc.Send[CNPostalAddressFormatter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNPostalAddressFormatter creates a new CNPostalAddressFormatter instance.
func NewCNPostalAddressFormatter() CNPostalAddressFormatter {
	return getCNPostalAddressFormatterClass().New()
}



// Returns a postal address as an attributed string and formatted for the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/attributedString(from:style:withDefaultAttributes:)

func (cc _CNPostalAddressFormatterClass) AttributedStringFromPostalAddressStyleWithDefaultAttributes(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle, attributes objectivec.IObject) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](objc.ID(cc.class), objc.Sel("attributedStringFromPostalAddress:style:withDefaultAttributes:"), postalAddress, style, attributes)
	return rv
}


// Returns a postal address as a string and formatted for the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:style:)

func (cc _CNPostalAddressFormatterClass) StringFromPostalAddressStyle(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("stringFromPostalAddress:style:"), postalAddress, style)
	return rv
}



// Returns a formatted postal address as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/attributedString(from:withDefaultAttributes:)

func (c_ CNPostalAddressFormatter) AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objectivec.IObject) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringFromPostalAddress:withDefaultAttributes:"), postalAddress, attributes)
	return rv
}



// Returns a formatted postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:)

func (c_ CNPostalAddressFormatter) StringFromPostalAddress(postalAddress ICNPostalAddress) foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("stringFromPostalAddress:"), postalAddress)
	return rv
}


// The style to apply when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/style

func (c_ CNPostalAddressFormatter) Style() CNPostalAddressFormatterStyle {
	rv := objc.Send[CNPostalAddressFormatterStyle](c_.ID, objc.Sel("style"))
	return rv
}


// The style to apply when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/style

func (c_ CNPostalAddressFormatter) SetStyle(value CNPostalAddressFormatterStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}


// The city of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscitykey

func (c_ CNPostalAddressFormatter) CNPostalAddressCityKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey

func (c_ CNPostalAddressFormatter) CNPostalAddressCountryKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey

func (c_ CNPostalAddressFormatter) CNPostalAddressISOCountryCodeKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}


// An attribute that identifies the localized property of postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresslocalizedpropertynameattribute

func (c_ CNPostalAddressFormatter) CNPostalAddressLocalizedPropertyNameAttribute() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressLocalizedPropertyNameAttribute"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey

func (c_ CNPostalAddressFormatter) CNPostalAddressPostalCodeKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}


// An attribute that identifies the purpose of a range of characters in an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspropertyattribute

func (c_ CNPostalAddressFormatter) CNPostalAddressPropertyAttribute() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressPropertyAttribute"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey

func (c_ CNPostalAddressFormatter) CNPostalAddressStateKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey

func (c_ CNPostalAddressFormatter) CNPostalAddressStreetKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}


// The subadministrative area of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssubadministrativeareakey

func (c_ CNPostalAddressFormatter) CNPostalAddressSubAdministrativeAreaKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressSubAdministrativeAreaKey"))
	return rv
}


// The sublocality of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssublocalitykey

func (c_ CNPostalAddressFormatter) CNPostalAddressSubLocalityKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressSubLocalityKey"))
	return rv
}



