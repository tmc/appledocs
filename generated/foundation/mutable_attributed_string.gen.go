// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableAttributedString] class.
var (
	mutableAttributedStringClass     _MutableAttributedStringClass
	mutableAttributedStringClassOnce sync.Once
)

func getMutableAttributedStringClass() _MutableAttributedStringClass {
	mutableAttributedStringClassOnce.Do(func() {
		mutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
	})
	return mutableAttributedStringClass
}

type _MutableAttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	IAttributedString
	AppendAttributedString(attrString unsafe.Pointer)
}

// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
type MutableAttributedString struct {
	AttributedString
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: AttributedStringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAttributedString) Autorelease() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAttributedString creates a new MutableAttributedString instance.
func NewMutableAttributedString() MutableAttributedString {
	return getMutableAttributedStringClass().New()
}


// Adds the characters and attributes of a given attributed string to the end of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m_ MutableAttributedString) AppendAttributedString(attrString unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}


