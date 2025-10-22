// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MutableString] class.
type IMutableString interface {
	IString
	AppendString(aString string)
	AppendFormat(format string)
	ApplyTransformReverseRangeUpdatedRange(transform IStringTransform, reverse bool, range_ IRange, resultingRange IRangePointer) bool
	DeleteCharactersInRange(range_ IRange)
	InsertStringAtIndex(aString string, loc uint)
	ReplaceCharactersInRangeWithString(range_ IRange, aString string)
	ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange IRange) uint
	SetString(aString string)
}

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

// Alloc allocates a new instance without initialization.
func (mc _MutableStringClass) Alloc() MutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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





// Returns an object initialized with initial storage for a given number of characters,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/init(capacity:)

func NewMutableStringWithCapacity(capacity uint) MutableString {
	instance := getMutableStringClass().Alloc()
	rv := objc.Send[MutableString](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	rv.Autorelease()
	return rv
}



// Returns an empty object with initial storage for a given number of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/stringWithCapacity:

func (mc _MutableStringClass) StringWithCapacity(capacity uint) MutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("stringWithCapacity:"), capacity)
	return rv
}


// Adds to the end of the receiver the characters of a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/append(_:)

func (m_ MutableString) AppendString(aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendString:"), objc.String(aString))
}


// Adds a constructed string to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/appendFormat:

func (m_ MutableString) AppendFormat(format string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendFormat:"), objc.String(format))
}


// Transliterates the receiver by applying a specified ICU string transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/applyTransform(_:reverse:range:updatedRange:)

func (m_ MutableString) ApplyTransformReverseRangeUpdatedRange(transform IStringTransform, reverse bool, range_ IRange, resultingRange IRangePointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("applyTransform:reverse:range:updatedRange:"), transform, reverse, range_, resultingRange)
	return rv
}


// Removes from the receiver the characters in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/deleteCharacters(in:)

func (m_ MutableString) DeleteCharactersInRange(range_ IRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteCharactersInRange:"), range_)
}


// Inserts into the receiver the characters of a given string at a given location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/insert(_:at:)

func (m_ MutableString) InsertStringAtIndex(aString string, loc uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertString:atIndex:"), objc.String(aString), loc)
}


// Replaces the characters from with those in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceCharacters(in:with:)

func (m_ MutableString) ReplaceCharactersInRangeWithString(range_ IRange, aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(aString))
}


// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceOccurrences(of:with:options:range:)

func (m_ MutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange IRange) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("replaceOccurrencesOfString:withString:options:range:"), objc.String(target), objc.String(replacement), options, searchRange)
	return rv
}


// Replaces the characters of the receiver with those in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/setString(_:)

func (m_ MutableString) SetString(aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setString:"), objc.String(aString))
}


