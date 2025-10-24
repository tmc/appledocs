// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABMutableMultiValue */


/* debug [class_header]: Header for ABMutableMultiValue */
// The class instance for the [ABMutableMultiValue] class.
var (
	ABMutableMultiValueClass     _ABMutableMultiValueClass
	ABMutableMultiValueClassOnce sync.Once
)

func getABMutableMultiValueClass() _ABMutableMultiValueClass {
	ABMutableMultiValueClassOnce.Do(func() {
		ABMutableMultiValueClass = _ABMutableMultiValueClass{objc.GetClass("ABMutableMultiValue")}
	})
	return ABMutableMultiValueClass
}

type _ABMutableMultiValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABMutableMultiValue */
// An interface definition for the [ABMutableMultiValue] class.
type IABMutableMultiValue interface {
	IABMultiValue
	
/* debug [class_interface_properties]: Properties for ABMutableMultiValue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABMutableMultiValue */
	// methods:
	AddValueWithLabel(value objc.IObject, label objc.IObject /* cross-framework: NSString */) foundation.String
	InsertValueWithLabelAtIndex(value objc.IObject, label objc.IObject /* cross-framework: NSString */, index uint) foundation.String
	RemoveValueAndLabelAtIndex(index uint) bool
	ReplaceValueAtIndexWithValue(index uint, value objc.IObject) bool
	ReplaceLabelAtIndexWithLabel(index uint, label objc.IObject /* cross-framework: NSString */) bool
	SetPrimaryIdentifier(identifier objc.IObject /* cross-framework: NSString */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABMutableMultiValue */
// Alloc allocates a new instance without initialization.
func (ac _ABMutableMultiValueClass) Alloc() ABMutableMultiValue {
	rv := objc.Send[ABMutableMultiValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ABMutableMultiValueClass) New() ABMutableMultiValue {
	rv := objc.Send[ABMutableMultiValue](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABMutableMultiValue) Init() ABMutableMultiValue {
	rv := objc.Send[ABMutableMultiValue](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABMutableMultiValue) Autorelease() ABMutableMultiValue {
	rv := objc.Send[ABMutableMultiValue](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABMutableMultiValue creates a new ABMutableMultiValue instance.
func NewABMutableMultiValue() ABMutableMultiValue {
	return getABMutableMultiValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABMutableMultiValue */
// A mutable representation of a property that might have multiple values.
//
// Each value in a multivalue list must be of the same type, and must have an associated predefined or user-defined label, and unique identifier. The labels, however, need not be unique. For example, you can have multiple Home phone numbers. Each multivalue object may have a primary identifier—used as a default value when a label is not provided. For example, a person record may have multiple addresses with the labels Home and Work, where Work is designated as the primary value. Instances of are mutable, see for additional methods that access the content of a multivalue list. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// A mutable representation of a property that might have multiple values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class
type ABMutableMultiValue struct {
	ABMultiValue
}

// ABMutableMultiValueFrom constructs a [ABMutableMultiValue] from an unsafe.Pointer.
//
// A mutable representation of a property that might have multiple values.
func ABMutableMultiValueFrom(ptr unsafe.Pointer) ABMutableMultiValue {
	return ABMutableMultiValue{
		ABMultiValue: ABMultiValueFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABMutableMultiValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABMutableMultiValue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABMutableMultiValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABMutableMultiValue */

// Adds a value and its label to a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/add(_:withLabel:)
func (a_ ABMutableMultiValue) AddValueWithLabel(value objc.IObject, label objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("addValue:withLabel:"), value, label)
	return rv
}/* debug [instance_methods/method]: AddValueWithLabel */


// Inserts a value and its label at the given index in a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/insert(_:withLabel:at:)
func (a_ ABMutableMultiValue) InsertValueWithLabelAtIndex(value objc.IObject, label objc.IObject /* cross-framework: NSString */, index uint) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("insertValue:withLabel:atIndex:"), value, label, index)
	return rv
}/* debug [instance_methods/method]: InsertValueWithLabelAtIndex */


// Removes the value and label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/removeAndLabel(at:)
func (a_ ABMutableMultiValue) RemoveValueAndLabelAtIndex(index uint) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeValueAndLabelAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RemoveValueAndLabelAtIndex */


// Replaces the value at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/replace(at:withValue:)
func (a_ ABMutableMultiValue) ReplaceValueAtIndexWithValue(index uint, value objc.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("replaceValueAtIndex:withValue:"), index, value)
	return rv
}/* debug [instance_methods/method]: ReplaceValueAtIndexWithValue */


// Replaces the label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/replaceLabel(at:withLabel:)
func (a_ ABMutableMultiValue) ReplaceLabelAtIndexWithLabel(index uint, label objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("replaceLabelAtIndex:withLabel:"), index, label)
	return rv
}/* debug [instance_methods/method]: ReplaceLabelAtIndexWithLabel */


// Sets the primary value to be the value for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue-swift.class/setPrimaryIdentifier(_:)
func (a_ ABMutableMultiValue) SetPrimaryIdentifier(identifier objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrimaryIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: SetPrimaryIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABMutableMultiValue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABMutableMultiValue */



