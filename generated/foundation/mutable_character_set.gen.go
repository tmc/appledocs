// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableCharacterSet] class.
var mutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}

type _MutableCharacterSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableCharacterSet] class.
type IMutableCharacterSet interface {
	ICharacterSet
}

// An object representing a mutable set of Unicode character values for use in search operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet

type MutableCharacterSet struct {
	CharacterSet
}

// MutableCharacterSetFrom constructs a [MutableCharacterSet] from an unsafe.Pointer.
//
// An object representing a mutable set of Unicode character values for use in search operations.
func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		CharacterSet: CharacterSetFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MutableCharacterSetClass) Alloc() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MutableCharacterSetClass) New() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableCharacterSet) Init() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableCharacterSet) Autorelease() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableCharacterSet creates a new MutableCharacterSet instance.
func NewMutableCharacterSet() MutableCharacterSet {
	return mutableCharacterSetClass.New()
}




