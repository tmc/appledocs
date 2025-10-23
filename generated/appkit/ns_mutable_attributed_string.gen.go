// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MutableAttributedString] class.
var (
	MutableAttributedStringClass     _MutableAttributedStringClass
	MutableAttributedStringClassOnce sync.Once
)

func getMutableAttributedStringClass() _MutableAttributedStringClass {
	MutableAttributedStringClassOnce.Do(func() {
		MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
	})
	return MutableAttributedStringClass
}

type _MutableAttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	objectivec.IObject
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type MutableAttributedString struct {
	objectivec.Object
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{objectivec.Object{objc.ID(ptr)}}
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




