// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNContactFormatter */


/* debug [class_header]: Header for CNContactFormatter */
// The class instance for the [CNContactFormatter] class.
var (
	CNContactFormatterClass     _CNContactFormatterClass
	CNContactFormatterClassOnce sync.Once
)

func getCNContactFormatterClass() _CNContactFormatterClass {
	CNContactFormatterClassOnce.Do(func() {
		CNContactFormatterClass = _CNContactFormatterClass{objc.GetClass("CNContactFormatter")}
	})
	return CNContactFormatterClass
}

type _CNContactFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactFormatter */
// An interface definition for the [CNContactFormatter] class.
type ICNContactFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for CNContactFormatter */
	// properties:
	Style() CNContactFormatterStyle
	SetStyle(value CNContactFormatterStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactFormatter */
	// methods:
	AttributedStringFromContactDefaultAttributes(contact ICNContact, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString
	StringFromContact(contact ICNContact) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactFormatter */
// Alloc allocates a new instance without initialization.
func (cc _CNContactFormatterClass) Alloc() CNContactFormatter {
	rv := objc.Send[CNContactFormatter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactFormatterClass) New() CNContactFormatter {
	rv := objc.Send[CNContactFormatter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactFormatter) Init() CNContactFormatter {
	rv := objc.Send[CNContactFormatter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactFormatter) Autorelease() CNContactFormatter {
	rv := objc.Send[CNContactFormatter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactFormatter creates a new CNContactFormatter instance.
func NewCNContactFormatter() CNContactFormatter {
	return getCNContactFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactFormatter */
// An object that you use to format contact information before displaying it to the user.
//
// A object handles international ordering and delimiting for the contact name components. When formatting many contacts, create an instance of this class and use the instance methods; otherwise use the class methods.


// An object that you use to format contact information before displaying it to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter
type CNContactFormatter struct {
	Formatter
}

// CNContactFormatterFrom constructs a [CNContactFormatter] from an unsafe.Pointer.
//
// An object that you use to format contact information before displaying it to the user.
func CNContactFormatterFrom(ptr unsafe.Pointer) CNContactFormatter {
	return CNContactFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactFormatter */

// Formats the contact name as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/attributedString(from:style:defaultAttributes:)
func (cc _CNContactFormatterClass) AttributedStringFromContactStyleDefaultAttributes(contact ICNContact, style CNContactFormatterStyle, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](objc.ID(cc.class), objc.Sel("attributedStringFromContact:style:defaultAttributes:"), contact, style, attributes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributedStringFromContactStyleDefaultAttributes) */


// Returns the delimiter to use between name components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/delimiter(for:)
func (cc _CNContactFormatterClass) DelimiterForContact(contact ICNContact) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("delimiterForContact:"), contact)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DelimiterForContact) */


// Returns the required key descriptor for the specified formatting style of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeys(for:)
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForStyle(style CNContactFormatterStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForStyle:"), style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorForRequiredKeysForStyle) */


// Returns the display name order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/nameOrder(for:)
func (cc _CNContactFormatterClass) NameOrderForContact(contact ICNContact) CNContactDisplayNameOrder {
	rv := objc.Send[CNContactDisplayNameOrder](objc.ID(cc.class), objc.Sel("nameOrderForContact:"), contact)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameOrderForContact) */


// Returns the contact name, formatted with the specified formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/string(from:style:)
func (cc _CNContactFormatterClass) StringFromContactStyle(contact ICNContact, style CNContactFormatterStyle) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("stringFromContact:style:"), contact, style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StringFromContactStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactFormatter */

// Returns the required key descriptor for the name delimiter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForDelimiter
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForDelimiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForDelimiter"))
	return rv
}/* debug [class_properties_class/property]: descriptorForRequiredKeysForDelimiter */

// Returns the required key descriptor for the display name order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForNameOrder
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForNameOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForNameOrder"))
	return rv
}/* debug [class_properties_class/property]: descriptorForRequiredKeysForNameOrder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactFormatter */

// Formats the contact name as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/attributedString(from:defaultAttributes:)
func (c_ CNContactFormatter) AttributedStringFromContactDefaultAttributes(contact ICNContact, attributes objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringFromContact:defaultAttributes:"), contact, attributes)
	return rv
}/* debug [instance_methods/method]: AttributedStringFromContactDefaultAttributes */


// Formats the contact name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/string(from:)
func (c_ CNContactFormatter) StringFromContact(contact ICNContact) foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("stringFromContact:"), contact)
	return rv
}/* debug [instance_methods/method]: StringFromContact */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactFormatter */

// Returns the required key descriptor for the name delimiter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForDelimiter
func (c_ CNContactFormatter) DescriptorForRequiredKeysForDelimiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptorForRequiredKeysForDelimiter"))
	return rv
}/* debug [instance_properties/getter]: descriptorForRequiredKeysForDelimiter */


// Returns the required key descriptor for the display name order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForNameOrder
func (c_ CNContactFormatter) DescriptorForRequiredKeysForNameOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptorForRequiredKeysForNameOrder"))
	return rv
}/* debug [instance_properties/getter]: descriptorForRequiredKeysForNameOrder */


// The formatting style for the contact name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/style
func (c_ CNContactFormatter) Style() CNContactFormatterStyle {
	rv := objc.Send[CNContactFormatterStyle](c_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// The formatting style for the contact name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/style
func (c_ CNContactFormatter) SetStyle(value CNContactFormatterStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactFormatter */



