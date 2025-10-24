// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNPostalAddressFormatter */


/* debug [class_header]: Header for CNPostalAddressFormatter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNPostalAddressFormatter */
// An interface definition for the [CNPostalAddressFormatter] class.
type ICNPostalAddressFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for CNPostalAddressFormatter */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNPostalAddressFormatter */
	// methods:
	AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString
	StringFromPostalAddress(postalAddress ICNPostalAddress) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNPostalAddressFormatter */
// Alloc allocates a new instance without initialization.
func (cc _CNPostalAddressFormatterClass) Alloc() CNPostalAddressFormatter {
	rv := objc.Send[CNPostalAddressFormatter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNPostalAddressFormatter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNPostalAddressFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNPostalAddressFormatter */

// Returns a postal address as an attributed string and formatted for the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/attributedString(from:style:withDefaultAttributes:)
func (cc _CNPostalAddressFormatterClass) AttributedStringFromPostalAddressStyleWithDefaultAttributes(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](objc.ID(cc.class), objc.Sel("attributedStringFromPostalAddress:style:withDefaultAttributes:"), postalAddress, style, attributes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributedStringFromPostalAddressStyleWithDefaultAttributes) */


// Returns a postal address as a string and formatted for the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:style:)
func (cc _CNPostalAddressFormatterClass) StringFromPostalAddressStyle(postalAddress ICNPostalAddress, style CNPostalAddressFormatterStyle) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("stringFromPostalAddress:style:"), postalAddress, style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StringFromPostalAddressStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNPostalAddressFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNPostalAddressFormatter */

// Returns a formatted postal address as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/attributedString(from:withDefaultAttributes:)
func (c_ CNPostalAddressFormatter) AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress ICNPostalAddress, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringFromPostalAddress:withDefaultAttributes:"), postalAddress, attributes)
	return rv
}/* debug [instance_methods/method]: AttributedStringFromPostalAddressWithDefaultAttributes */


// Returns a formatted postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/string(from:)
func (c_ CNPostalAddressFormatter) StringFromPostalAddress(postalAddress ICNPostalAddress) foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("stringFromPostalAddress:"), postalAddress)
	return rv
}/* debug [instance_methods/method]: StringFromPostalAddress */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNPostalAddressFormatter */

// The style to apply when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/style
func (c_ CNPostalAddressFormatter) Style() CNPostalAddressFormatterStyle {
	rv := objc.Send[CNPostalAddressFormatterStyle](c_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// The style to apply when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatter/style
func (c_ CNPostalAddressFormatter) SetStyle(value CNPostalAddressFormatterStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// The city of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscitykey
func (c_ CNPostalAddressFormatter) CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressCityKey */


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey
func (c_ CNPostalAddressFormatter) CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressCountryKey */


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey
func (c_ CNPostalAddressFormatter) CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressISOCountryCodeKey */


// An attribute that identifies the localized property of postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresslocalizedpropertynameattribute
func (c_ CNPostalAddressFormatter) CNPostalAddressLocalizedPropertyNameAttribute() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressLocalizedPropertyNameAttribute"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressLocalizedPropertyNameAttribute */


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey
func (c_ CNPostalAddressFormatter) CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressPostalCodeKey */


// An attribute that identifies the purpose of a range of characters in an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspropertyattribute
func (c_ CNPostalAddressFormatter) CNPostalAddressPropertyAttribute() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPropertyAttribute"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressPropertyAttribute */


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey
func (c_ CNPostalAddressFormatter) CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressStateKey */


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey
func (c_ CNPostalAddressFormatter) CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressStreetKey */


// The subadministrative area of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssubadministrativeareakey
func (c_ CNPostalAddressFormatter) CNPostalAddressSubAdministrativeAreaKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressSubAdministrativeAreaKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressSubAdministrativeAreaKey */


// The sublocality of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresssublocalitykey
func (c_ CNPostalAddressFormatter) CNPostalAddressSubLocalityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressSubLocalityKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressSubLocalityKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNPostalAddressFormatter */



