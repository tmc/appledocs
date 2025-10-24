// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMutableString */


/* debug [class_header]: Header for NSMutableString */
// The class instance for the [MutableString] class.
var (
	MutableStringClass     _MutableStringClass
	MutableStringClassOnce sync.Once
)

func getMutableStringClass() _MutableStringClass {
	MutableStringClassOnce.Do(func() {
		MutableStringClass = _MutableStringClass{objc.GetClass("NSMutableString")}
	})
	return MutableStringClass
}

type _MutableStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableString */
// An interface definition for the [MutableString] class.
type IMutableString interface {
	IString
	
/* debug [class_interface_properties]: Properties for MutableString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableString */
	// methods:
	AppendString(aString IString)
	AppendFormat(format IString)
	ApplyTransformReverseRangeUpdatedRange(transform StringTransform /* typedef */, reverse bool, range_ objc.IObject /* cross-framework: Range */, resultingRange RangePointer /* typedef */) bool
	DeleteCharactersInRange(range_ objc.IObject /* cross-framework: Range */)
	InsertStringAtIndex(aString IString, loc uint)
	ReplaceCharactersInRangeWithString(range_ objc.IObject /* cross-framework: Range */, aString IString)
	ReplaceOccurrencesOfStringWithStringOptionsRange(target IString, replacement IString, options StringCompareOptions, searchRange objc.IObject /* cross-framework: Range */) uint
	SetString(aString IString)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableString */
// Alloc allocates a new instance without initialization.
func (mc _MutableStringClass) Alloc() MutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableStringClass) New() MutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableString) Init() MutableString {
	rv := objc.Send[MutableString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableString) Autorelease() MutableString {
	rv := objc.Send[MutableString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableString creates a new MutableString instance.
func NewMutableString() MutableString {
	return getMutableStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableString */
// A dynamic plain-text Unicode string object.
//
// In Swift, you can use this type instead of a in cases that require reference semantics. The class declares the programmatic interface to an object that manages a mutable string—that is, a string whose contents can be edited—that conceptually represents an array of Unicode characters. To construct and manage an immutable string—or a string that cannot be changed after it has been created—use an object of the class. The class adds one primitive method— —to the basic string-handling behavior inherited from . All other methods that modify a string work through this method. For example, simply replaces the characters in a range of length, while replaces the characters in a given range with no characters. NSMutableString is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A dynamic plain-text Unicode string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString
type MutableString struct {
	string
}

// MutableStringFrom constructs a [MutableString] from an unsafe.Pointer.
//
// A dynamic plain-text Unicode string object.
func MutableStringFrom(ptr unsafe.Pointer) MutableString {
	return MutableString{
		String: stringFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableString */

// Returns an object initialized with initial storage for a given number of characters,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/init(capacity:)
func NewMutableStringWithCapacity(capacity uint) MutableString {
	instance := getMutableStringClass().Alloc()
	rv := objc.Send[MutableString](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableStringWithCapacity */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableString */

// Returns an empty object with initial storage for a given number of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/stringWithCapacity:
func (mc _MutableStringClass) StringWithCapacity(capacity uint) IMutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("stringWithCapacity:"), capacity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StringWithCapacity) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableString */

// Adds to the end of the receiver the characters of a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/append(_:)
func (m_ MutableString) AppendString(aString IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendString:"), aString)
}/* debug [instance_methods/method]: AppendString */


// Adds a constructed string to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/appendFormat:
func (m_ MutableString) AppendFormat(format IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendFormat:"), format)
}/* debug [instance_methods/method]: AppendFormat */


// Transliterates the receiver by applying a specified ICU string transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/applyTransform(_:reverse:range:updatedRange:)
func (m_ MutableString) ApplyTransformReverseRangeUpdatedRange(transform StringTransform /* typedef */, reverse bool, range_ objc.IObject /* cross-framework: Range */, resultingRange RangePointer /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("applyTransform:reverse:range:updatedRange:"), transform, reverse, range_, resultingRange)
	return rv
}/* debug [instance_methods/method]: ApplyTransformReverseRangeUpdatedRange */


// Removes from the receiver the characters in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/deleteCharacters(in:)
func (m_ MutableString) DeleteCharactersInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteCharactersInRange:"), range_)
}/* debug [instance_methods/method]: DeleteCharactersInRange */


// Inserts into the receiver the characters of a given string at a given location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/insert(_:at:)
func (m_ MutableString) InsertStringAtIndex(aString IString, loc uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertString:atIndex:"), aString, loc)
}/* debug [instance_methods/method]: InsertStringAtIndex */


// Replaces the characters from with those in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceCharacters(in:with:)
func (m_ MutableString) ReplaceCharactersInRangeWithString(range_ objc.IObject /* cross-framework: Range */, aString IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, aString)
}/* debug [instance_methods/method]: ReplaceCharactersInRangeWithString */


// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceOccurrences(of:with:options:range:)
func (m_ MutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target IString, replacement IString, options StringCompareOptions, searchRange objc.IObject /* cross-framework: Range */) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("replaceOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}/* debug [instance_methods/method]: ReplaceOccurrencesOfStringWithStringOptionsRange */


// Replaces the characters of the receiver with those in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/setString(_:)
func (m_ MutableString) SetString(aString IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setString:"), aString)
}/* debug [instance_methods/method]: SetString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableString */


