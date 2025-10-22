// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNContactsUserDefaults] class.
var (
	CNContactsUserDefaultsClass     _CNContactsUserDefaultsClass
	CNContactsUserDefaultsClassOnce sync.Once
)

func getCNContactsUserDefaultsClass() _CNContactsUserDefaultsClass {
	CNContactsUserDefaultsClassOnce.Do(func() {
		CNContactsUserDefaultsClass = _CNContactsUserDefaultsClass{objc.GetClass("CNContactsUserDefaults")}
	})
	return CNContactsUserDefaultsClass
}

type _CNContactsUserDefaultsClass struct {
	class objc.Class
}

// An interface definition for the [CNContactsUserDefaults] class.
type ICNContactsUserDefaults interface {
	objectivec.IObject
	CountryCode() string
	SortOrder() CNContactSortOrder
}

// An object that defines the default options to use when displaying contacts.


// An object that defines the default options to use when displaying contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults

type CNContactsUserDefaults struct {
	objectivec.Object
}

// CNContactsUserDefaultsFrom constructs a [CNContactsUserDefaults] from an unsafe.Pointer.
//
// An object that defines the default options to use when displaying contacts.
func CNContactsUserDefaultsFrom(ptr unsafe.Pointer) CNContactsUserDefaults {
	return CNContactsUserDefaults{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactsUserDefaultsClass) Alloc() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactsUserDefaultsClass) New() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactsUserDefaults) Init() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactsUserDefaults) Autorelease() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactsUserDefaults creates a new CNContactsUserDefaults instance.
func NewCNContactsUserDefaults() CNContactsUserDefaults {
	return getCNContactsUserDefaultsClass().New()
}



// The singleton contacts user defaults object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/shared()

func (cc _CNContactsUserDefaultsClass) SharedDefaults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sharedDefaults"))
	return rv
}


// An ISO country code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/countryCode

func (c_ CNContactsUserDefaults) CountryCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("countryCode"))
	return rv
}


// Default sorting order by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/sortOrder

func (c_ CNContactsUserDefaults) SortOrder() CNContactSortOrder {
	rv := objc.Send[CNContactSortOrder](c_.ID, objc.Sel("sortOrder"))
	return rv
}



