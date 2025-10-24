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
	// properties:
	// methods:
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




