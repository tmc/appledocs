// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var (
	KeyedUnarchiverClass     _KeyedUnarchiverClass
	KeyedUnarchiverClassOnce sync.Once
)

func getKeyedUnarchiverClass() _KeyedUnarchiverClass {
	KeyedUnarchiverClassOnce.Do(func() {
		KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}
	})
	return KeyedUnarchiverClass
}

type _KeyedUnarchiverClass struct {
	class objc.Class
}

// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	ICoder
	DecodeBoolForKey(key string) bool
	DecodeIntForKey(key string) int
	DecodeObjectForKey(key string) objc.ID
	DecodingFailurePolicy() unsafe.Pointer
	SetDecodingFailurePolicy(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	RequiresSecureCoding() bool
	SetRequiresSecureCoding(value bool)
}

// A decoder that restores data from an archive referenced by keys.
//
// is a concrete subclass of that defines methods for decoding a set of named objects (and scalar values) from a keyed archive. The class produces archives that this class can decode. The archiver creates keyed archive as a hierarchy of objects. The archiver treats each object as a namespace into which it can encode other objects. This means that an unarchiver can only decode objects encoded within the immediate scope of their parent object. Objects encoded elsewhere in the hierarchy — whether higher than, lower than, or parallel to this particular object — aren’t accessible. In this way, the keys used by a particular object to encode its instance variables need to be unique only within the scope of that object. If you invoke one of the -prefixed methods of this class using a key that does not exist in the archive, the return value indicates failure. This value varies by decoded type. For example, if a key does not exist in an archive, returns , returns , and returns . supports limited type coercion for numeric types. You can use any of the integer decode methods to decode a value encoded as any type of integer, whether a standard or an explicit 32-bit or 64-bit integer. Likewise, you can use the - or -returning decode methods to handle value encoded as a or . If an encoded value is too large to fit within the coerced type, the decoding method throws a . Further, when trying to coerce a value to an incompatible type — for example decoding an as a — the decoding method throws an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver
type KeyedUnarchiver struct {
	Coder
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive referenced by keys.
func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedUnarchiver) Autorelease() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedUnarchiver creates a new KeyedUnarchiver instance.
func NewKeyedUnarchiver() KeyedUnarchiver {
	return getKeyedUnarchiverClass().New()
}


// Decodes a Boolean value associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeBool(forKey:)
func (k_ KeyedUnarchiver) DecodeBoolForKey(key string) bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("decodeBoolForKey:"), objc.String(key))
	return rv
}

// Decodes an integer value associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeIntForKey:
func (k_ KeyedUnarchiver) DecodeIntForKey(key string) int {
	rv := objc.Send[int](k_.ID, objc.Sel("decodeIntForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an object associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeObject(forKey:)
func (k_ KeyedUnarchiver) DecodeObjectForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](k_.ID, objc.Sel("decodeObjectForKey:"), objc.String(key))
	return rv
}

// The action to take when this unarchiver fails to decode an entry.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodingFailurePolicy
func (k_ KeyedUnarchiver) DecodingFailurePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}


// SetDecodingFailurePolicy sets the value of the decodingFailurePolicy property.
// The action to take when this unarchiver fails to decode an entry.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodingFailurePolicy
func (k_ KeyedUnarchiver) SetDecodingFailurePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDecodingFailurePolicy:"), value)
}

// The receiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/delegate
func (k_ KeyedUnarchiver) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The receiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/delegate
func (k_ KeyedUnarchiver) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDelegate:"), value)
}

// Indicates whether the receiver requires all unarchived classes to conform to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/requiressecurecoding
func (k_ KeyedUnarchiver) RequiresSecureCoding() bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}


// SetRequiresSecureCoding sets the value of the requiresSecureCoding property.
// Indicates whether the receiver requires all unarchived classes to conform to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/requiressecurecoding
func (k_ KeyedUnarchiver) SetRequiresSecureCoding(value bool) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRequiresSecureCoding:"), value)
}



