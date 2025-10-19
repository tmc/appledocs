// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedArchiver] class.
var keyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}

type _KeyedArchiverClass struct {
	class objc.Class
}

// An interface definition for the [KeyedArchiver] class.
type IKeyedArchiver interface {
	ICoder
	EncodeDoubleForKey(value float64, key string)
}

// An encoder that stores an object’s data to an archive referenced by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver

type KeyedArchiver struct {
	Coder
}

// KeyedArchiverFrom constructs a [KeyedArchiver] from an unsafe.Pointer.
//
// An encoder that stores an object’s data to an archive referenced by keys.
func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		Coder: CoderFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (kc _KeyedArchiverClass) Alloc() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (kc _KeyedArchiverClass) New() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedArchiver) Init() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedArchiver) Autorelease() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedArchiver creates a new KeyedArchiver instance.
func NewKeyedArchiver() KeyedArchiver {
	return keyedArchiverClass.New()
}


// Encodes a given value and associates it with a key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value float64, key string) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeDouble:forKey:"), value, objc.String(key))
}


