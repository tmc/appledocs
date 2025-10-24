// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABRecord */


/* debug [class_header]: Header for ABRecord */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABRecord */
// An interface definition for the [ABRecord] class.
type IABRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ABRecord */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	UniqueId() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABRecord */
	// methods:
	IsReadOnly() bool
	RemoveValueForProperty(property objc.IObject /* cross-framework: NSString */) bool
	SetValueForProperty(value objc.IObject, property objc.IObject /* cross-framework: NSString */) bool
	SetValueForPropertyError(value objc.IObject, property objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	ValueForProperty(property objc.IObject /* cross-framework: NSString */) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABRecord */
// Alloc allocates a new instance without initialization.
func (ac _ABRecordClass) Alloc() ABRecord {
	rv := objc.Send[ABRecord](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABRecord */
// An abstract class that defines the common properties for all Address Book records.
//
// is an abstract superclass providing a common interface to, and defining common properties for, all Address Book records. A property is a field in the database record, such as the first or last name of a person record. ABRecord defines the types of properties supported, and basic methods for getting, setting, and removing property values. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// An abstract class that defines the common properties for all Address Book records.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABRecord */

// Initializes a record using the given address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/init(addressBook:)
func NewABRecordWithAddressBook(addressBook IABAddressBook) ABRecord {
	instance := getABRecordClass().Alloc()
	rv := objc.Send[ABRecord](instance.ID, objc.Sel("initWithAddressBook:"), addressBook)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewABRecordWithAddressBook */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABRecord */

// Returns whether a record is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/isReadOnly()
func (a_ ABRecord) IsReadOnly() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadOnly"))
	return rv
}/* debug [instance_methods/method]: IsReadOnly */


// Removes the value for a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/removeValue(forProperty:)
func (a_ ABRecord) RemoveValueForProperty(property objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeValueForProperty:"), property)
	return rv
}/* debug [instance_methods/method]: RemoveValueForProperty */


// Sets the value of a given property for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/setValue(_:forProperty:)
func (a_ ABRecord) SetValueForProperty(value objc.IObject, property objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forProperty:"), value, property)
	return rv
}/* debug [instance_methods/method]: SetValueForProperty */


// Sets the value of a given property for a record, returning error information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/setValue(_:forProperty:error:)
func (a_ ABRecord) SetValueForPropertyError(value objc.IObject, property objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forProperty:error:"), value, property, error_)
	return rv
}/* debug [instance_methods/method]: SetValueForPropertyError */


// Returns the value of a given property for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/value(forProperty:)
func (a_ ABRecord) ValueForProperty(property objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForProperty:"), property)
	return rv
}/* debug [instance_methods/method]: ValueForProperty */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABRecord */

// A user-visible string representing the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/displayName
func (a_ ABRecord) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// Returns the unique ID for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecord-swift.class/uniqueId
func (a_ ABRecord) UniqueId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("uniqueId"))
	return rv
}/* debug [instance_properties/getter]: uniqueId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABRecord */


