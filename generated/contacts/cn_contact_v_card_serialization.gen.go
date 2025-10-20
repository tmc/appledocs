// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNContactVCardSerialization] class.
var (
	CNContactVCardSerializationClass     _CNContactVCardSerializationClass
	CNContactVCardSerializationClassOnce sync.Once
)

func getCNContactVCardSerializationClass() _CNContactVCardSerializationClass {
	CNContactVCardSerializationClassOnce.Do(func() {
		CNContactVCardSerializationClass = _CNContactVCardSerializationClass{objc.GetClass("CNContactVCardSerialization")}
	})
	return CNContactVCardSerializationClass
}

type _CNContactVCardSerializationClass struct {
	class objc.Class
}

// An interface definition for the [CNContactVCardSerialization] class.
type ICNContactVCardSerialization interface {
	objectivec.IObject
}

// An object you use to convert to and from a vCard representation of the user’s contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization
type CNContactVCardSerialization struct {
	objectivec.Object
}

// CNContactVCardSerializationFrom constructs a [CNContactVCardSerialization] from an unsafe.Pointer.
//
// An object you use to convert to and from a vCard representation of the user’s contacts.
func CNContactVCardSerializationFrom(ptr unsafe.Pointer) CNContactVCardSerialization {
	return CNContactVCardSerialization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactVCardSerializationClass) Alloc() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactVCardSerializationClass) New() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactVCardSerialization) Init() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactVCardSerialization) Autorelease() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactVCardSerialization creates a new CNContactVCardSerialization instance.
func NewCNContactVCardSerialization() CNContactVCardSerialization {
	return getCNContactVCardSerializationClass().New()
}


// Returns the contacts from the vCard data.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/contacts(with:)
func (cc _CNContactVCardSerializationClass) ContactsWithDataError(data unsafe.Pointer, error_ unsafe.Pointer) []CNContact {
	rv := objc.Send[[]CNContact](objc.ID(cc.class), objc.Sel("contactsWithData:error:"), data, error_)
	return rv
}

// Returns the vCard representation of the specified contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/data(with:)
func (cc _CNContactVCardSerializationClass) DataWithContactsError(contacts unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("dataWithContacts:error:"), contacts, error_)
	return rv
}

// Use to fetch all contact keys required to create vCard data from a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/descriptorForRequiredKeys()
func (cc _CNContactVCardSerializationClass) DescriptorForRequiredKeys() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeys"))
	return rv
}



