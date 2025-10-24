// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABAddressBook */


/* debug [class_header]: Header for ABAddressBook */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABAddressBook */
// An interface definition for the [ABAddressBook] class.
type IABAddressBook interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ABAddressBook */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABAddressBook */
	// methods:
	AddRecord(record IABRecord) bool
	AddRecordError(record IABRecord, error_ unsafe.Pointer) bool
	DefaultCountryCode() foundation.String
	DefaultNameOrdering() int
	FormattedAddressFromDictionary(address objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString
	Groups() foundation.Array
	HasUnsavedChanges() bool
	Me() IABPerson
	People() foundation.Array
	RecordForUniqueId(uniqueId objc.IObject /* cross-framework: NSString */) IABRecord
	RecordClassFromUniqueId(uniqueId objc.IObject /* cross-framework: NSString */) foundation.String
	RecordsMatchingSearchElement(search IABSearchElement) foundation.Array
	RemoveRecord(record IABRecord) bool
	RemoveRecordError(record IABRecord, error_ unsafe.Pointer) bool
	Save() bool
	SaveAndReturnError(error_ unsafe.Pointer) bool
	SetMe(moi IABPerson)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABAddressBook */
// Alloc allocates a new instance without initialization.
func (ac _ABAddressBookClass) Alloc() ABAddressBook {
	rv := objc.Send[ABAddressBook](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABAddressBook */
// The main object you use to access the Address Book database.
//
// The class provides a programming interface to the Address Book—a centralized database used by multiple applications to store contact and other personal information about people. The Address Book database also supports the notion of a “group” containing one or more persons. People may belong to multiple groups, and groups may also belong to other groups with some restrictions (for example, no circular references are allowed). The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// The main object you use to access the Address Book database.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABAddressBook *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABAddressBook */

// Returns a new instance of , or if the Address Book database can’t be initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/addressBook
func (ac _ABAddressBookClass) AddressBook() ABAddressBook {
	rv := objc.Send[ABAddressBook](objc.ID(ac.class), objc.Sel("addressBook"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddressBook) */


// Returns the unique shared instance of , or if the Address Book database can’t be initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/shared()
func (ac _ABAddressBookClass) SharedAddressBook() ABAddressBook {
	rv := objc.Send[ABAddressBook](objc.ID(ac.class), objc.Sel("sharedAddressBook"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedAddressBook) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABAddressBook */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABAddressBook */

// Adds an or record to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/add(_:)
func (a_ ABAddressBook) AddRecord(record IABRecord) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addRecord:"), record)
	return rv
}/* debug [instance_methods/method]: AddRecord */


// Adds an or record to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/add(_:error:)
func (a_ ABAddressBook) AddRecordError(record IABRecord, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addRecord:error:"), record, error_)
	return rv
}/* debug [instance_methods/method]: AddRecordError */


// Returns the default country code for records with unspecified country codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/defaultCountryCode()
func (a_ ABAddressBook) DefaultCountryCode() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("defaultCountryCode"))
	return rv
}/* debug [instance_methods/method]: DefaultCountryCode */


// Returns the default name ordering defined by the user in the Address Book application’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/defaultNameOrdering()
func (a_ ABAddressBook) DefaultNameOrdering() int {
	rv := objc.Send[int](a_.ID, objc.Sel("defaultNameOrdering"))
	return rv
}/* debug [instance_methods/method]: DefaultNameOrdering */


// Returns an attributed string containing the formatted address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/formattedAddress(from:)
func (a_ ABAddressBook) FormattedAddressFromDictionary(address objc.IObject /* cross-framework: NSDictionary */) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("formattedAddressFromDictionary:"), address)
	return rv
}/* debug [instance_methods/method]: FormattedAddressFromDictionary */


// Returns an array of all the groups in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/groups()
func (a_ ABAddressBook) Groups() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_methods/method]: Groups */


// Indicates whether an address book has changes that have not been saved to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/hasUnsavedChanges()
func (a_ ABAddressBook) HasUnsavedChanges() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasUnsavedChanges"))
	return rv
}/* debug [instance_methods/method]: HasUnsavedChanges */


// Returns the record that represents the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/me()
func (a_ ABAddressBook) Me() IABPerson {
	rv := objc.Send[ABPerson](a_.ID, objc.Sel("me"))
	return rv
}/* debug [instance_methods/method]: Me */


// Returns an array of all the people in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/people()
func (a_ ABAddressBook) People() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("people"))
	return rv
}/* debug [instance_methods/method]: People */


// Returns the person or group record that matches the given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/record(forUniqueId:)
func (a_ ABAddressBook) RecordForUniqueId(uniqueId objc.IObject /* cross-framework: NSString */) IABRecord {
	rv := objc.Send[ABRecord](a_.ID, objc.Sel("recordForUniqueId:"), uniqueId)
	return rv
}/* debug [instance_methods/method]: RecordForUniqueId */


// Returns the class name of the record that matches the given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/recordClass(fromUniqueId:)
func (a_ ABAddressBook) RecordClassFromUniqueId(uniqueId objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("recordClassFromUniqueId:"), uniqueId)
	return rv
}/* debug [instance_methods/method]: RecordClassFromUniqueId */


// Returns an array of records that match the given search element, or returns an empty array if no records match the search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/records(matching:)
func (a_ ABAddressBook) RecordsMatchingSearchElement(search IABSearchElement) foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("recordsMatchingSearchElement:"), search)
	return rv
}/* debug [instance_methods/method]: RecordsMatchingSearchElement */


// Removes an or record from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/remove(_:)
func (a_ ABAddressBook) RemoveRecord(record IABRecord) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeRecord:"), record)
	return rv
}/* debug [instance_methods/method]: RemoveRecord */


// Removes an or record from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/remove(_:error:)
func (a_ ABAddressBook) RemoveRecordError(record IABRecord, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeRecord:error:"), record, error_)
	return rv
}/* debug [instance_methods/method]: RemoveRecordError */


// Saves all the changes made since the last save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/save()
func (a_ ABAddressBook) Save() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("save"))
	return rv
}/* debug [instance_methods/method]: Save */


// Saves all the changes made since the last save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/saveAndReturnError()
func (a_ ABAddressBook) SaveAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("saveAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SaveAndReturnError */


// Sets the record that represents the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBook-swift.class/setMe(_:)
func (a_ ABAddressBook) SetMe(moi IABPerson) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMe:"), moi)
}/* debug [instance_methods/method]: SetMe */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABAddressBook */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABAddressBook */



