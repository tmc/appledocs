// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableString] class.
var MutableStringClass objc.Class

func init() {
	MutableStringClass = objc.GetClass("NSMutableString")
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
func (mc MutableString) Alloc() MutableString {
	ret := objc.ID(MutableStringClass).Send(objc.RegisterName("alloc"))
	return MutableString{ret}
}

// Init initializes the instance.
func (m_ MutableString) Init() MutableString {
	ret := m_.ID.Send(objc.RegisterName("init"))
	return MutableString{ret}
}
// Returns an   object initialized with initial storage for a given number of characters, [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/init(capacity:)
func NewMutableStringWithCapacity(capacity uint) MutableString {
	instance := MutableString{}.Alloc()
	sel := objc.RegisterName("initWithCapacity:")
	ret := instance.ID.Send(sel, capacity)
	instance = MutableString{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns an empty   object with initial storage for a given number of characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/stringWithCapacity:
func (mc MutableString) StringWithCapacity(capacity uint) unsafe.Pointer {
	sel := objc.RegisterName("stringWithCapacity:")
	ret := objc.ID(MutableStringClass).Send(sel, capacity)
	return unsafe.Pointer(ret)
}
// Adds to the end of the receiver the characters of a given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/append(_:)
func (m_ MutableString) AppendString(aString string) {
	sel := objc.RegisterName("appendString:")
	m_.ID.Send(sel, aString)
}
// Adds a constructed string to the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/appendFormat:
func (m_ MutableString) AppendFormat(format string) {
	sel := objc.RegisterName("appendFormat:")
	m_.ID.Send(sel, format)
}
// Transliterates the receiver by applying a specified ICU string transform. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/applyTransform(_:reverse:range:updatedRange:)
func (m_ MutableString) ApplyTransformReverseRangeUpdatedRange(transform unsafe.Pointer, reverse bool, range_ Range, resultingRange unsafe.Pointer) bool {
	sel := objc.RegisterName("applyTransform:reverse:range:updatedRange:")
	ret := m_.ID.Send(sel, transform, reverse, range_, resultingRange)
	return ret != 0
}
// Removes from the receiver the characters in a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/deleteCharacters(in:)
func (m_ MutableString) DeleteCharactersInRange(range_ Range) {
	sel := objc.RegisterName("deleteCharactersInRange:")
	m_.ID.Send(sel, range_)
}
// Inserts into the receiver the characters of a given string at a given location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/insert(_:at:)
func (m_ MutableString) InsertStringAtIndex(aString string, loc uint) {
	sel := objc.RegisterName("insertString:atIndex:")
	m_.ID.Send(sel, aString, loc)
}
// Replaces the characters from   with those in  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/replaceCharacters(in:with:)
func (m_ MutableString) ReplaceCharactersInRangeWithString(range_ Range, aString string) {
	sel := objc.RegisterName("replaceCharactersInRange:withString:")
	m_.ID.Send(sel, range_, aString)
}
// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/replaceOccurrences(of:with:options:range:)
func (m_ MutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options unsafe.Pointer, searchRange Range) uint {
	sel := objc.RegisterName("replaceOccurrencesOfString:withString:options:range:")
	ret := m_.ID.Send(sel, target, replacement, options, searchRange)
	return uint(ret)
}
// Replaces the characters of the receiver with those in a given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableString/setString(_:)
func (m_ MutableString) SetString(aString string) {
	sel := objc.RegisterName("setString:")
	m_.ID.Send(sel, aString)
}

