// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableDictionary] class.
var (
	mutableDictionaryClass     _MutableDictionaryClass
	mutableDictionaryClassOnce sync.Once
)

func getMutableDictionaryClass() _MutableDictionaryClass {
	mutableDictionaryClassOnce.Do(func() {
		mutableDictionaryClass = _MutableDictionaryClass{objc.GetClass("NSMutableDictionary")}
	})
	return mutableDictionaryClass
}

type _MutableDictionaryClass struct {
	class objc.Class
}

// An interface definition for the [MutableDictionary] class.
type IMutableDictionary interface {
	IDictionary
}

// A dynamic collection of objects associated with unique keys.
//
// In Swift, you can use this type instead of a variable in cases that require reference semantics. The class declares the programmatic interface to objects that manage mutable associations of keys and values. It adds modification operations to the basic operations it inherits from . is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary
type MutableDictionary struct {
	Dictionary
}

// MutableDictionaryFrom constructs a [MutableDictionary] from an unsafe.Pointer.
//
// A dynamic collection of objects associated with unique keys.
func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: DictionaryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableDictionaryClass) Alloc() MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableDictionaryClass) New() MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableDictionary) Init() MutableDictionary {
	rv := objc.Send[MutableDictionary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableDictionary) Autorelease() MutableDictionary {
	rv := objc.Send[MutableDictionary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableDictionary creates a new MutableDictionary instance.
func NewMutableDictionary() MutableDictionary {
	return getMutableDictionaryClass().New()
}




