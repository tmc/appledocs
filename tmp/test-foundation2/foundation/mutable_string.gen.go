// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var mutableStringClass _MutableStringClass

func init() {
	mutableStringClass = _MutableStringClass{objc.GetClass("NSMutableString")}
}

type _MutableStringClass struct {
	class objc.Class
}

type MutableString struct {
	objc.ID
}

func MutableStringFrom(ptr unsafe.Pointer) MutableString {
	return MutableString{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableStringClass) Alloc() MutableString {
	rv := objc.Send[MutableString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return mutableStringClass.New()
}
// Returns an object initialized with initial storage for a given number of characters, [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/init(capacity:)
func NewMutableStringWithCapacity(capacity uint) MutableString {
	instance := mutableStringClass.Alloc()
	rv := objc.Send[MutableString](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	rv.Autorelease()
	return rv
}


// Returns an empty object with initial storage for a given number of characters. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/stringWithCapacity:
func (mc _MutableStringClass) StringWithCapacity(capacity uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("stringWithCapacity:"), capacity)
	return rv
}
// Adds to the end of the receiver the characters of a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/append(_:)
func (m_ MutableString) AppendString(aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendString:"), aString)
}
// Adds a constructed string to the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/appendFormat:
func (m_ MutableString) AppendFormat(format string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendFormat:"), format)
}
// Transliterates the receiver by applying a specified ICU string transform. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/applyTransform(_:reverse:range:updatedRange:)
func (m_ MutableString) ApplyTransformReverseRangeUpdatedRange(transform unsafe.Pointer, reverse bool, range_ unsafe.Pointer, resultingRange unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("applyTransform:reverse:range:updatedRange:"), transform, reverse, range_, resultingRange)
	return rv
}
// Removes from the receiver the characters in a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/deleteCharacters(in:)
func (m_ MutableString) DeleteCharactersInRange(range_ unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteCharactersInRange:"), range_)
}
// Inserts into the receiver the characters of a given string at a given location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/insert(_:at:)
func (m_ MutableString) InsertStringAtIndex(aString string, loc uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertString:atIndex:"), aString, loc)
}
// Replaces the characters from with those in . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceCharacters(in:with:)
func (m_ MutableString) ReplaceCharactersInRangeWithString(range_ unsafe.Pointer, aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, aString)
}
// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceOccurrences(of:with:options:range:)
func (m_ MutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange unsafe.Pointer) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("replaceOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}
// Replaces the characters of the receiver with those in a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableString/setString(_:)
func (m_ MutableString) SetString(aString string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setString:"), aString)
}


