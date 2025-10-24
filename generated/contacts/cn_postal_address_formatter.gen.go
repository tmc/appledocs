// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	IFormatter
	// properties:
	Style() CNPostalAddressFormatterStyle
	SetStyle(value CNPostalAddressFormatterStyle)
	CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressLocalizedPropertyNameAttribute() objc.IObject /* cross-framework: NSString */
	CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressPropertyAttribute() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressSubAdministrativeAreaKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressSubLocalityKey() objc.IObject /* cross-framework: NSString */
	// methods:
	AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objc.IObject /* cross-framework: NSDictionary */) objc.IObject /* cross-framework: AttributedString */
	StringFromPostalAddress(postalAddress ICNPostalAddress) objc.IObject /* cross-framework: String */
}

// An object that you use to format a contact’s postal addresses.
//
// A object handles international formatting of postal addresses. It is recommended that you create an instance of this class when formatting many postal addresses, and use the instance methods; otherwise use the class methods.


// An object that you use to format a contact’s postal addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter
type CNPostalAddressFormatter struct {
	Formatter
}

// CNPostalAddressFormatterFrom constructs a [CNPostalAddressFormatter] from an unsafe.Pointer.
//
// An object that you use to format a contact’s postal addresses.
func CNPostalAddressFormatterFrom(ptr unsafe.Pointer) CNPostalAddressFormatter {
	return CNPostalAddressFormatter{
		Formatter: FormatterFrom(ptr),
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
func (cc _CNPostalAddressFormatterClass) AttributedStringFromPostalAddressStyleWithDefaultAttributes(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle, attributes objc.IObject /* cross-framework: NSDictionary */) objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](objc.ID(cc.class), objc.Sel("attributedStringFromPostalAddress:style:withDefaultAttributes:"), postalAddress, style, attributes)
	return rv
}


// Returns a postal address as a string and formatted for the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:style:)
func (cc _CNPostalAddressFormatterClass) StringFromPostalAddressStyle(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("stringFromPostalAddress:style:"), postalAddress, style)
	return rv
}


// Returns a formatted postal address as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/attributedString(from:withDefaultAttributes:)
func (c_ CNPostalAddressFormatter) AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objc.IObject /* cross-framework: NSDictionary */) objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringFromPostalAddress:withDefaultAttributes:"), postalAddress, attributes)
	return rv
}


// Returns a formatted postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:)
func (c_ CNPostalAddressFormatter) StringFromPostalAddress(postalAddress ICNPostalAddress) objc.IObject /* cross-framework: String */ {
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
func (c_ CNPostalAddressFormatter) CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey
func (c_ CNPostalAddressFormatter) CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey
func (c_ CNPostalAddressFormatter) CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}


// An attribute that identifies the localized property of postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresslocalizedpropertynameattribute
func (c_ CNPostalAddressFormatter) CNPostalAddressLocalizedPropertyNameAttribute() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressLocalizedPropertyNameAttribute"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey
func (c_ CNPostalAddressFormatter) CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}


// An attribute that identifies the purpose of a range of characters in an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspropertyattribute
func (c_ CNPostalAddressFormatter) CNPostalAddressPropertyAttribute() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPropertyAttribute"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey
func (c_ CNPostalAddressFormatter) CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey
func (c_ CNPostalAddressFormatter) CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}


// The subadministrative area of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssubadministrativeareakey
func (c_ CNPostalAddressFormatter) CNPostalAddressSubAdministrativeAreaKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressSubAdministrativeAreaKey"))
	return rv
}


// The sublocality of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssublocalitykey
func (c_ CNPostalAddressFormatter) CNPostalAddressSubLocalityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressSubLocalityKey"))
	return rv
}



