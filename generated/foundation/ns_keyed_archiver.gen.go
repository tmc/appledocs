// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedArchiver] class.
var (
	KeyedArchiverClass     _KeyedArchiverClass
	KeyedArchiverClassOnce sync.Once
)

func getKeyedArchiverClass() _KeyedArchiverClass {
	KeyedArchiverClassOnce.Do(func() {
		KeyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}
	})
	return KeyedArchiverClass
}

type _KeyedArchiverClass struct {
	class objc.Class
}

// An interface definition for the [KeyedArchiver] class.
type IKeyedArchiver interface {
	ICoder
	EncodeDoubleForKey(value unsafe.Pointer, key string)
}

// An encoder that stores an object’s data to an archive referenced by keys.
//
// , a concrete subclass of , provides a way to encode objects (and scalar values) into an architecture-independent format suitable for storage in a file. When you archive a set of objects, the archiver writes the class information and instance variables for each object to the archive. The companion class decodes the data in an archive and creates a set of objects equivalent to the original set. A keyed archive differs from a non-keyed archive in that all the objects and values encoded into the archive have names, or keys. When decoding a non-keyed archive, the decoder must decode values in the same order the original encoder used. When decoding a keyed archive, the decoder requests values by name, meaning it can decode values out of sequence or not at all. Keyed archives, therefore, provide better support for forward and backward compatibility. The keys given to encoded values must be unique only within the scope of the currently-encoding object. A keyed archive is hierarchical, so the keys used by object A to encode its instance variables don’t conflict with the keys used by object B. This is true even if A and B are instances of the same class. Within a single object, however, the keys used by a subclass can conflict with keys used in its superclasses. An object can write the archive data to a file or to a mutable-data object (an instance of ) that you provide.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getKeyedArchiverClass().New()
}


// Encodes a given value and associates it with a key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeDouble:forKey:"), value, objc.String(key))
}

// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/requiressecurecoding
func (k_ KeyedArchiver) RequiresSecureCoding() bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}


// SetRequiresSecureCoding sets the value of the requiresSecureCoding property.
// Indicates whether the archiver requires all archived classes to resist object substitution attacks.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/requiressecurecoding
func (k_ KeyedArchiver) SetRequiresSecureCoding(value bool) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRequiresSecureCoding:"), value)
}

// The format in which the receiver encodes its data.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/outputformat
func (k_ KeyedArchiver) OutputFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("outputFormat"))
	return rv
}


// SetOutputFormat sets the value of the outputFormat property.
// The format in which the receiver encodes its data.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/outputformat
func (k_ KeyedArchiver) SetOutputFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setOutputFormat:"), value)
}

// The encoded data for the archiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/encodeddata
func (k_ KeyedArchiver) EncodedData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("encodedData"))
	return rv
}


// SetEncodedData sets the value of the encodedData property.
// The encoded data for the archiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/encodeddata
func (k_ KeyedArchiver) SetEncodedData(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setEncodedData:"), value)
}

// The archiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/delegate
func (k_ KeyedArchiver) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The archiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/delegate
func (k_ KeyedArchiver) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDelegate:"), value)
}



