// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABMultiValue */


/* debug [class_header]: Header for ABMultiValue */
// The class instance for the [ABMultiValue] class.
var (
	ABMultiValueClass     _ABMultiValueClass
	ABMultiValueClassOnce sync.Once
)

func getABMultiValueClass() _ABMultiValueClass {
	ABMultiValueClassOnce.Do(func() {
		ABMultiValueClass = _ABMultiValueClass{objc.GetClass("ABMultiValue")}
	})
	return ABMultiValueClass
}

type _ABMultiValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABMultiValue */
// An interface definition for the [ABMultiValue] class.
type IABMultiValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ABMultiValue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABMultiValue */
	// methods:
	Count() uint
	IdentifierAtIndex(index uint) foundation.String
	IndexForIdentifier(identifier objc.IObject /* cross-framework: NSString */) uint
	LabelAtIndex(index uint) foundation.String
	LabelForIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.ID
	PrimaryIdentifier() foundation.String
	PropertyType() ABPropertyType /* typedef */
	ValueAtIndex(index uint) objc.ID
	ValueForIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABMultiValue */
// Alloc allocates a new instance without initialization.
func (ac _ABMultiValueClass) Alloc() ABMultiValue {
	rv := objc.Send[ABMultiValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ABMultiValueClass) New() ABMultiValue {
	rv := objc.Send[ABMultiValue](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABMultiValue) Init() ABMultiValue {
	rv := objc.Send[ABMultiValue](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABMultiValue) Autorelease() ABMultiValue {
	rv := objc.Send[ABMultiValue](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABMultiValue creates a new ABMultiValue instance.
func NewABMultiValue() ABMultiValue {
	return getABMultiValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABMultiValue */
// An immutable representation of a property that might have multiple values.
//
// Each value in a multivalue list must be of the same type, and must have an associated predefined or user-defined label, and unique identifier. The labels, however, need not be unique. For example, you can have multiple Home phone numbers. Each multivalue object may have a primary identifier—used as a default value when a label is not provided. For example, a person record may have multiple addresses with the labels Home and Work, where Work is designated as the primary value. Instances of this class are immutable, see for methods that manipulate the content of a multivalue list. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// An immutable representation of a property that might have multiple values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class
type ABMultiValue struct {
	objectivec.Object
}

// ABMultiValueFrom constructs a [ABMultiValue] from an unsafe.Pointer.
//
// An immutable representation of a property that might have multiple values.
func ABMultiValueFrom(ptr unsafe.Pointer) ABMultiValue {
	return ABMultiValue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABMultiValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABMultiValue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABMultiValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABMultiValue */

// Returns the number of entries in a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/count()
func (a_ ABMultiValue) Count() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_methods/method]: Count */


// Returns the identifier for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/identifier(at:)
func (a_ ABMultiValue) IdentifierAtIndex(index uint) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("identifierAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: IdentifierAtIndex */


// Returns the index for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/index(forIdentifier:)
func (a_ ABMultiValue) IndexForIdentifier(identifier objc.IObject /* cross-framework: NSString */) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexForIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: IndexForIdentifier */


// Returns the label for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/label(at:)
func (a_ ABMultiValue) LabelAtIndex(index uint) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("labelAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: LabelAtIndex */


// Returns the label for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/label(forIdentifier:)
func (a_ ABMultiValue) LabelForIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("labelForIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: LabelForIdentifier */


// Returns the identifier for the primary value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/primaryIdentifier()
func (a_ ABMultiValue) PrimaryIdentifier() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("primaryIdentifier"))
	return rv
}/* debug [instance_methods/method]: PrimaryIdentifier */


// Returns the type for the values in a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/propertyType()
func (a_ ABMultiValue) PropertyType() ABPropertyType /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("propertyType"))
	return rv
}/* debug [instance_methods/method]: PropertyType */


// Returns the value for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/value(at:)
func (a_ ABMultiValue) ValueAtIndex(index uint) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ValueAtIndex */


// Returns the value for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue-swift.class/value(forIdentifier:)
func (a_ ABMultiValue) ValueForIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: ValueForIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABMultiValue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABMultiValue */



