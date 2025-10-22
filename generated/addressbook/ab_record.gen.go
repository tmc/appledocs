// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ABRecord] class.
var (
	ABRecordClass     _ABRecordClass
	ABRecordClassOnce sync.Once
)

func getABRecordClass() _ABRecordClass {
	ABRecordClassOnce.Do(func() {
		ABRecordClass = _ABRecordClass{objc.GetClass("ABRecord")}
	})
	return ABRecordClass
}

type _ABRecordClass struct {
	class objc.Class
}

// An interface definition for the [ABRecord] class.
type IABRecord interface {
	objectivec.IObject
	IsReadOnly() bool
	RemoveValueForProperty(property string) bool
	SetValueForProperty(value objectivec.IObject, property string) bool
	SetValueForPropertyError(value objectivec.IObject, property string, error_ unsafe.Pointer) bool
	ValueForProperty(property string) objc.ID
	DisplayName() string
	UniqueId() string
}

// An abstract class that defines the common properties for all Address Book records.
//
// is an abstract superclass providing a common interface to, and defining common properties for, all Address Book records. A property is a field in the database record, such as the first or last name of a person record. ABRecord defines the types of properties supported, and basic methods for getting, setting, and removing property values. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class
type ABRecord struct {
	objectivec.Object
}

// ABRecordFrom constructs a [ABRecord] from an unsafe.Pointer.
//
// An abstract class that defines the common properties for all Address Book records.
func ABRecordFrom(ptr unsafe.Pointer) ABRecord {
	return ABRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ABRecordClass) Alloc() ABRecord {
	rv := objc.Send[ABRecord](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABRecordClass) New() ABRecord {
	rv := objc.Send[ABRecord](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABRecord) Init() ABRecord {
	rv := objc.Send[ABRecord](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABRecord) Autorelease() ABRecord {
	rv := objc.Send[ABRecord](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABRecord creates a new ABRecord instance.
func NewABRecord() ABRecord {
	return getABRecordClass().New()
}




// Initializes a record using the given address book.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/init(addressBook:)
func NewABRecordWithAddressBook(addressBook IABAddressBook) ABRecord {
	instance := getABRecordClass().Alloc()
	rv := objc.Send[ABRecord](instance.ID, objc.Sel("initWithAddressBook:"), addressBook)
	rv.Autorelease()
	return rv
}


// Returns whether a record is read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/isReadOnly()
func (a_ ABRecord) IsReadOnly() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadOnly"))
	return rv
}

// Removes the value for a given property.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/removeValue(forProperty:)
func (a_ ABRecord) RemoveValueForProperty(property string) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeValueForProperty:"), objc.String(property))
	return rv
}

// Sets the value of a given property for a record.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/setValue(_:forProperty:)
func (a_ ABRecord) SetValueForProperty(value objectivec.IObject, property string) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forProperty:"), value, objc.String(property))
	return rv
}

// Sets the value of a given property for a record, returning error information.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/setValue(_:forProperty:error:)
func (a_ ABRecord) SetValueForPropertyError(value objectivec.IObject, property string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forProperty:error:"), value, objc.String(property), error_)
	return rv
}

// Returns the value of a given property for a record.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/value(forProperty:)
func (a_ ABRecord) ValueForProperty(property string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForProperty:"), objc.String(property))
	return rv
}

// A user-visible string representing the record.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/displayName
func (a_ ABRecord) DisplayName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("displayName"))
	return rv
}

// Returns the unique ID for a record.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord/uniqueId
func (a_ ABRecord) UniqueId() string {
	rv := objc.Send[string](a_.ID, objc.Sel("uniqueId"))
	return rv
}


