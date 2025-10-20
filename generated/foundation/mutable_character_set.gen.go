// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableCharacterSet] class.
var (
	mutableCharacterSetClass     _MutableCharacterSetClass
	mutableCharacterSetClassOnce sync.Once
)

func getMutableCharacterSetClass() _MutableCharacterSetClass {
	mutableCharacterSetClassOnce.Do(func() {
		mutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}
	})
	return mutableCharacterSetClass
}

type _MutableCharacterSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableCharacterSet] class.
type IMutableCharacterSet interface {
	ICharacterSet
}

// An object representing a mutable set of Unicode character values for use in search operations.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class declares the programmatic interface to objects that manage a modifiable set of Unicode characters. You can add or remove characters from a mutable character set as numeric values in structures or as character values in strings, combine character sets by union or intersection, and invert a character set. Mutable character sets are less efficient to use than immutable character sets. If you don’t need to change a character set after creating it, create an immutable copy with and use that. defines no primitive methods. Subclasses must implement all methods declared by this class in addition to the primitives of . They must also implement . is “toll-free bridged” with its Core Foundation counterpart, . See for more information.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getMutableCharacterSetClass().New()
}




