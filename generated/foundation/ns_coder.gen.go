// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Coder] class.
var (
	CoderClass     _CoderClass
	CoderClassOnce sync.Once
)

func getCoderClass() _CoderClass {
	CoderClassOnce.Do(func() {
		CoderClass = _CoderClass{objc.GetClass("NSCoder")}
	})
	return CoderClass
}

type _CoderClass struct {
	class objc.Class
}

// An interface definition for the [Coder] class.
type ICoder interface {
	objectivec.IObject
	DecodingFailurePolicy() unsafe.Pointer
	AllowedClasses() unsafe.Pointer
	SetAllowedClasses(value unsafe.Pointer)
	AllowsKeyedCoding() bool
	SetAllowsKeyedCoding(value bool)
	Error() IError
	SetError(value IError)
	RequiresSecureCoding() bool
	SetRequiresSecureCoding(value bool)
	SystemVersion() unsafe.Pointer
	SetSystemVersion(value unsafe.Pointer)
	NSCoderErrorMaximum() int
	SetNSCoderErrorMaximum(value int)
	NSCoderErrorMinimum() int
	SetNSCoderErrorMinimum(value int)
	NSCoderInvalidValueError() int
	SetNSCoderInvalidValueError(value int)
	NSCoderReadCorruptError() int
	SetNSCoderReadCorruptError(value int)
	NSCoderValueNotFoundError() int
	SetNSCoderValueNotFoundError(value int)
	DecodeObject() objc.ID
	DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// declares the interface used by concrete subclasses to transfer objects and other values between memory and some other format. This capability provides the basis for archiving (storing objects and data on disk) and distribution (copying objects and data items between different processes or threads). The concrete subclasses provided by Foundation for these purposes are , , , , and . Concrete subclasses of are “coder classes”, and instances of these classes are “coder objects” (or simply “coders”). A coder that can only encode values is an “encoder”, and one that can only decode values is a “decoder”. operates on objects, scalars, C arrays, structures, strings, and on pointers to these types. It doesn’t handle types whose implementation varies across platforms, such as , , function pointers, and long chains of pointers. A coder stores object type information along with the data, so an object decoded from a stream of bytes is normally of the same class as the object that was originally encoded into the stream. An object can change its class when encoded, however; this is described in . The AVFoundation framework adds methods to the class to make it easier to create archives including Core Media time structures, and extract Core Media time structure from archives.


// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder
type Coder struct {
	objectivec.Object
}

// CoderFrom constructs a [Coder] from an unsafe.Pointer.
//
// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoderClass) Alloc() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoderClass) New() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Coder) Init() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Coder) Autorelease() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoder creates a new Coder instance.
func NewCoder() Coder {
	return getCoderClass().New()
}



// Decodes and returns an object that was previously encoded with any of the methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject()
func (c_ Coder) DecodeObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObject"))
	return rv
}


// Decodes an object for the key, restricted to the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, objc.String(key))
	return rv
}


// The action the coder should take when decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c_ Coder) DecodingFailurePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}


// The set of coded classes allowed for secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowedclasses
func (c_ Coder) AllowedClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedClasses"))
	return rv
}


// The set of coded classes allowed for secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowedclasses
func (c_ Coder) SetAllowedClasses(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedClasses:"), value)
}


// A Boolean value that indicates whether the receiver supports keyed coding of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowskeyedcoding
func (c_ Coder) AllowsKeyedCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsKeyedCoding"))
	return rv
}


// A Boolean value that indicates whether the receiver supports keyed coding of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowskeyedcoding
func (c_ Coder) SetAllowsKeyedCoding(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsKeyedCoding:"), value)
}


// An error in the top-level encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/error
func (c_ Coder) Error() IError {
	rv := objc.Send[Error](c_.ID, objc.Sel("error"))
	return rv
}


// An error in the top-level encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/error
func (c_ Coder) SetError(value IError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setError:"), value)
}


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/requiressecurecoding
func (c_ Coder) RequiresSecureCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/requiressecurecoding
func (c_ Coder) SetRequiresSecureCoding(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresSecureCoding:"), value)
}


// The system version in effect for the archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/systemversion
func (c_ Coder) SystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemVersion"))
	return rv
}


// The system version in effect for the archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/systemversion
func (c_ Coder) SetSystemVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemVersion:"), value)
}


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) NSCoderErrorMaximum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) SetNSCoderErrorMaximum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMaximum:"), value)
}


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) NSCoderErrorMinimum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) SetNSCoderErrorMinimum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMinimum:"), value)
}


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) NSCoderInvalidValueError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderInvalidValueError"))
	return rv
}


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) SetNSCoderInvalidValueError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderInvalidValueError:"), value)
}


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) NSCoderReadCorruptError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderReadCorruptError"))
	return rv
}


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) SetNSCoderReadCorruptError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderReadCorruptError:"), value)
}


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) NSCoderValueNotFoundError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderValueNotFoundError"))
	return rv
}


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) SetNSCoderValueNotFoundError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderValueNotFoundError:"), value)
}



