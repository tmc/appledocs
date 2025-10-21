// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ABAddressBook] class.
var (
	ABAddressBookClass     _ABAddressBookClass
	ABAddressBookClassOnce sync.Once
)

func getABAddressBookClass() _ABAddressBookClass {
	ABAddressBookClassOnce.Do(func() {
		ABAddressBookClass = _ABAddressBookClass{objc.GetClass("ABAddressBook")}
	})
	return ABAddressBookClass
}

type _ABAddressBookClass struct {
	class objc.Class
}

// An interface definition for the [ABAddressBook] class.
type IABAddressBook interface {
	objectivec.IObject
	AddRecord(record unsafe.Pointer) bool
	AddRecordError(record unsafe.Pointer, error_ unsafe.Pointer) bool
	DefaultCountryCode() string
	DefaultNameOrdering() int
	FormattedAddressFromDictionary(address objc.ID) unsafe.Pointer
	Groups() unsafe.Pointer
	HasUnsavedChanges() bool
	Me() unsafe.Pointer
	People() unsafe.Pointer
	RecordForUniqueId(uniqueId string) unsafe.Pointer
	RecordClassFromUniqueId(uniqueId string) string
	RecordsMatchingSearchElement(search unsafe.Pointer) unsafe.Pointer
	RemoveRecord(record unsafe.Pointer) bool
	RemoveRecordError(record unsafe.Pointer, error_ unsafe.Pointer) bool
	Save() bool
	SaveAndReturnError(error_ unsafe.Pointer) bool
	SetMe(moi unsafe.Pointer)
}

// The main object you use to access the Address Book database.
//
// The class provides a programming interface to the Address Book—a centralized database used by multiple applications to store contact and other personal information about people. The Address Book database also supports the notion of a “group” containing one or more persons. People may belong to multiple groups, and groups may also belong to other groups with some restrictions (for example, no circular references are allowed). The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class
type ABAddressBook struct {
	objectivec.Object
}

// ABAddressBookFrom constructs a [ABAddressBook] from an unsafe.Pointer.
//
// The main object you use to access the Address Book database.
func ABAddressBookFrom(ptr unsafe.Pointer) ABAddressBook {
	return ABAddressBook{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ABAddressBookClass) Alloc() ABAddressBook {
	rv := objc.Send[ABAddressBook](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABAddressBookClass) New() ABAddressBook {
	rv := objc.Send[ABAddressBook](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABAddressBook) Init() ABAddressBook {
	rv := objc.Send[ABAddressBook](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABAddressBook) Autorelease() ABAddressBook {
	rv := objc.Send[ABAddressBook](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABAddressBook creates a new ABAddressBook instance.
func NewABAddressBook() ABAddressBook {
	return getABAddressBookClass().New()
}


// Returns a new instance of , or if the Address Book database can’t be initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/addressBook
func (ac _ABAddressBookClass) AddressBook() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("addressBook"))
	return rv
}

// Returns the unique shared instance of , or if the Address Book database can’t be initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/shared()
func (ac _ABAddressBookClass) SharedAddressBook() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("sharedAddressBook"))
	return rv
}

// Adds an or record to the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/add(_:)
func (a_ ABAddressBook) AddRecord(record unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addRecord:"), record)
	return rv
}

// Adds an or record to the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/add(_:error:)
func (a_ ABAddressBook) AddRecordError(record unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addRecord:error:"), record, error_)
	return rv
}

// Returns the default country code for records with unspecified country codes.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/defaultCountryCode()
func (a_ ABAddressBook) DefaultCountryCode() string {
	rv := objc.Send[string](a_.ID, objc.Sel("defaultCountryCode"))
	return rv
}

// Returns the default name ordering defined by the user in the Address Book application’s preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/defaultNameOrdering()
func (a_ ABAddressBook) DefaultNameOrdering() int {
	rv := objc.Send[int](a_.ID, objc.Sel("defaultNameOrdering"))
	return rv
}

// Returns an attributed string containing the formatted address.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/formattedAddress(from:)
func (a_ ABAddressBook) FormattedAddressFromDictionary(address objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("formattedAddressFromDictionary:"), address)
	return rv
}

// Returns an array of all the groups in the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/groups()
func (a_ ABAddressBook) Groups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("groups"))
	return rv
}

// Indicates whether an address book has changes that have not been saved to the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/hasUnsavedChanges()
func (a_ ABAddressBook) HasUnsavedChanges() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasUnsavedChanges"))
	return rv
}

// Returns the record that represents the logged-in user.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/me()
func (a_ ABAddressBook) Me() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("me"))
	return rv
}

// Returns an array of all the people in the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/people()
func (a_ ABAddressBook) People() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("people"))
	return rv
}

// Returns the person or group record that matches the given unique ID.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/record(forUniqueId:)
func (a_ ABAddressBook) RecordForUniqueId(uniqueId string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("recordForUniqueId:"), objc.String(uniqueId))
	return rv
}

// Returns the class name of the record that matches the given unique ID.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/recordClass(fromUniqueId:)
func (a_ ABAddressBook) RecordClassFromUniqueId(uniqueId string) string {
	rv := objc.Send[string](a_.ID, objc.Sel("recordClassFromUniqueId:"), objc.String(uniqueId))
	return rv
}

// Returns an array of records that match the given search element, or returns an empty array if no records match the search element.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/records(matching:)
func (a_ ABAddressBook) RecordsMatchingSearchElement(search unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("recordsMatchingSearchElement:"), search)
	return rv
}

// Removes an or record from the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/remove(_:)
func (a_ ABAddressBook) RemoveRecord(record unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeRecord:"), record)
	return rv
}

// Removes an or record from the Address Book database.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/remove(_:error:)
func (a_ ABAddressBook) RemoveRecordError(record unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeRecord:error:"), record, error_)
	return rv
}

// Saves all the changes made since the last save.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/save()
func (a_ ABAddressBook) Save() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("save"))
	return rv
}

// Saves all the changes made since the last save.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/saveAndReturnError()
func (a_ ABAddressBook) SaveAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveAndReturnError:"), error_)
	return rv
}

// Sets the record that represents the logged-in user.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook/setMe(_:)
func (a_ ABAddressBook) SetMe(moi unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMe:"), moi)
}



