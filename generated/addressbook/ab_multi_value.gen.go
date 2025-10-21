// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ABMultiValue] class.
type IABMultiValue interface {
	objectivec.IObject
	Count() uint
	IdentifierAtIndex(index uint) string
	IndexForIdentifier(identifier string) uint
	LabelAtIndex(index uint) string
	LabelForIdentifier(identifier string) objc.ID
	PrimaryIdentifier() string
	PropertyType() unsafe.Pointer
	ValueAtIndex(index uint) objc.ID
	ValueForIdentifier(identifier string) objc.ID
}

// An immutable representation of a property that might have multiple values.
//
// Each value in a multivalue list must be of the same type, and must have an associated predefined or user-defined label, and unique identifier. The labels, however, need not be unique. For example, you can have multiple Home phone numbers. Each multivalue object may have a primary identifier—used as a default value when a label is not provided. For example, a person record may have multiple addresses with the labels Home and Work, where Work is designated as the primary value. Instances of this class are immutable, see for methods that manipulate the content of a multivalue list. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.
//
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

// Alloc allocates a new instance without initialization.
func (ac _ABMultiValueClass) Alloc() ABMultiValue {
	rv := objc.Send[ABMultiValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the number of entries in a multivalue list.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/count()
func (a_ ABMultiValue) Count() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("count"))
	return rv
}

// Returns the identifier for the given index.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/identifier(at:)
func (a_ ABMultiValue) IdentifierAtIndex(index uint) string {
	rv := objc.Send[string](a_.ID, objc.Sel("identifierAtIndex:"), index)
	return rv
}

// Returns the index for the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/index(forIdentifier:)
func (a_ ABMultiValue) IndexForIdentifier(identifier string) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexForIdentifier:"), objc.String(identifier))
	return rv
}

// Returns the label for the given index.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/label(at:)
func (a_ ABMultiValue) LabelAtIndex(index uint) string {
	rv := objc.Send[string](a_.ID, objc.Sel("labelAtIndex:"), index)
	return rv
}

// Returns the label for the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/label(forIdentifier:)
func (a_ ABMultiValue) LabelForIdentifier(identifier string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("labelForIdentifier:"), objc.String(identifier))
	return rv
}

// Returns the identifier for the primary value.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/primaryIdentifier()
func (a_ ABMultiValue) PrimaryIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("primaryIdentifier"))
	return rv
}

// Returns the type for the values in a multivalue list.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/propertyType()
func (a_ ABMultiValue) PropertyType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("propertyType"))
	return rv
}

// Returns the value for the given index.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/value(at:)
func (a_ ABMultiValue) ValueAtIndex(index uint) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}

// Returns the value for the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValue/value(forIdentifier:)
func (a_ ABMultiValue) ValueForIdentifier(identifier string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForIdentifier:"), objc.String(identifier))
	return rv
}



