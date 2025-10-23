// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ABMutableMultiValue] class.
type IABMutableMultiValue interface {
	IABMultiValue
	// properties:
	// methods:
	AddValueWithLabel(value objectivec.IObject, label string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	InsertValueWithLabelAtIndex(value objectivec.IObject, label string /* primitive/slice/pointer. */, index uint /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	RemoveValueAndLabelAtIndex(index uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	ReplaceValueAtIndexWithValue(index uint /* primitive/slice/pointer. */, value objectivec.IObject) bool /* primitive/slice/pointer. */
	ReplaceLabelAtIndexWithLabel(index uint /* primitive/slice/pointer. */, label string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SetPrimaryIdentifier(identifier string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
}

// A mutable representation of a property that might have multiple values.
//
// Each value in a multivalue list must be of the same type, and must have an associated predefined or user-defined label, and unique identifier. The labels, however, need not be unique. For example, you can have multiple Home phone numbers. Each multivalue object may have a primary identifier—used as a default value when a label is not provided. For example, a person record may have multiple addresses with the labels Home and Work, where Work is designated as the primary value. Instances of are mutable, see for additional methods that access the content of a multivalue list. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// A mutable representation of a property that might have multiple values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue
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

// Alloc allocates a new instance without initialization.
func (ac _ABMutableMultiValueClass) Alloc() ABMutableMultiValue {
	rv := objc.Send[ABMutableMultiValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Adds a value and its label to a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/add(_:withLabel:)
func (a_ ABMutableMultiValue) AddValueWithLabel(value objectivec.IObject, label string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](a_.ID, objc.Sel("addValue:withLabel:"), value, objc.String(label))
	return rv
}


// Inserts a value and its label at the given index in a multivalue list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/insert(_:withLabel:at:)
func (a_ ABMutableMultiValue) InsertValueWithLabelAtIndex(value objectivec.IObject, label string /* primitive/slice/pointer. */, index uint /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](a_.ID, objc.Sel("insertValue:withLabel:atIndex:"), value, objc.String(label), index)
	return rv
}


// Removes the value and label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/removeAndLabel(at:)
func (a_ ABMutableMultiValue) RemoveValueAndLabelAtIndex(index uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeValueAndLabelAtIndex:"), index)
	return rv
}


// Replaces the value at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/replace(at:withValue:)
func (a_ ABMutableMultiValue) ReplaceValueAtIndexWithValue(index uint /* primitive/slice/pointer. */, value objectivec.IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("replaceValueAtIndex:withValue:"), index, value)
	return rv
}


// Replaces the label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/replaceLabel(at:withLabel:)
func (a_ ABMutableMultiValue) ReplaceLabelAtIndexWithLabel(index uint /* primitive/slice/pointer. */, label string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("replaceLabelAtIndex:withLabel:"), index, objc.String(label))
	return rv
}


// Sets the primary value to be the value for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValue/setPrimaryIdentifier(_:)
func (a_ ABMutableMultiValue) SetPrimaryIdentifier(identifier string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrimaryIdentifier:"), objc.String(identifier))
	return rv
}



