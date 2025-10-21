// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNContactFormatter] class.
type ICNContactFormatter interface {
	objectivec.IObject
	AttributedStringFromContactDefaultAttributes(contact unsafe.Pointer, attributes objc.ID) unsafe.Pointer
	StringFromContact(contact unsafe.Pointer) unsafe.Pointer
}

// An object that you use to format contact information before displaying it to the user.
//
// A object handles international ordering and delimiting for the contact name components. When formatting many contacts, create an instance of this class and use the instance methods; otherwise use the class methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter
type CNContactFormatter struct {
	objectivec.Object
}

// CNContactFormatterFrom constructs a [CNContactFormatter] from an unsafe.Pointer.
//
// An object that you use to format contact information before displaying it to the user.
func CNContactFormatterFrom(ptr unsafe.Pointer) CNContactFormatter {
	return CNContactFormatter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactFormatterClass) Alloc() CNContactFormatter {
	rv := objc.Send[CNContactFormatter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Formats the contact name as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/attributedString(from:style:defaultAttributes:)
func (cc _CNContactFormatterClass) AttributedStringFromContactStyleDefaultAttributes(contact unsafe.Pointer, style unsafe.Pointer, attributes objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("attributedStringFromContact:style:defaultAttributes:"), contact, style, attributes)
	return rv
}

// Returns the delimiter to use between name components.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/delimiter(for:)
func (cc _CNContactFormatterClass) DelimiterForContact(contact unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("delimiterForContact:"), contact)
	return rv
}

// Returns the required key descriptor for the specified formatting style of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeys(for:)
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForStyle(style unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForStyle:"), style)
	return rv
}

// Returns the display name order.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/nameOrder(for:)
func (cc _CNContactFormatterClass) NameOrderForContact(contact unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("nameOrderForContact:"), contact)
	return rv
}

// Returns the contact name, formatted with the specified formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/string(from:style:)
func (cc _CNContactFormatterClass) StringFromContactStyle(contact unsafe.Pointer, style unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("stringFromContact:style:"), contact, style)
	return rv
}

// Returns the required key descriptor for the name delimiter.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForDelimiter
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForDelimiter() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForDelimiter"))
	return rv
}
// Returns the required key descriptor for the display name order.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForNameOrder
func (cc _CNContactFormatterClass) DescriptorForRequiredKeysForNameOrder() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeysForNameOrder"))
	return rv
}
// Formats the contact name as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/attributedString(from:defaultAttributes:)
func (c_ CNContactFormatter) AttributedStringFromContactDefaultAttributes(contact unsafe.Pointer, attributes objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributedStringFromContact:defaultAttributes:"), contact, attributes)
	return rv
}

// Formats the contact name.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/string(from:)
func (c_ CNContactFormatter) StringFromContact(contact unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("stringFromContact:"), contact)
	return rv
}

// Returns the required key descriptor for the name delimiter.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForDelimiter
func (c_ CNContactFormatter) DescriptorForRequiredKeysForDelimiter() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("descriptorForRequiredKeysForDelimiter"))
	return rv
}

// Returns the required key descriptor for the display name order.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/descriptorForRequiredKeysForNameOrder
func (c_ CNContactFormatter) DescriptorForRequiredKeysForNameOrder() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("descriptorForRequiredKeysForNameOrder"))
	return rv
}

// The formatting style for the contact name.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/style
func (c_ CNContactFormatter) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// The formatting style for the contact name.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatter/style
func (c_ CNContactFormatter) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}


