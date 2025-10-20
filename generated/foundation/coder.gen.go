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
	coderClass     _CoderClass
	coderClassOnce sync.Once
)

func getCoderClass() _CoderClass {
	coderClassOnce.Do(func() {
		coderClass = _CoderClass{objc.GetClass("NSCoder")}
	})
	return coderClass
}

type _CoderClass struct {
	class objc.Class
}

// An interface definition for the [Coder] class.
type ICoder interface {
	objectivec.IObject
	DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// declares the interface used by concrete subclasses to transfer objects and other values between memory and some other format. This capability provides the basis for archiving (storing objects and data on disk) and distribution (copying objects and data items between different processes or threads). The concrete subclasses provided by Foundation for these purposes are , , , , and . Concrete subclasses of are “coder classes”, and instances of these classes are “coder objects” (or simply “coders”). A coder that can only encode values is an “encoder”, and one that can only decode values is a “decoder”. operates on objects, scalars, C arrays, structures, strings, and on pointers to these types. It doesn’t handle types whose implementation varies across platforms, such as , , function pointers, and long chains of pointers. A coder stores object type information along with the data, so an object decoded from a stream of bytes is normally of the same class as the object that was originally encoded into the stream. An object can change its class when encoded, however; this is described in . The AVFoundation framework adds methods to the class to make it easier to create archives including Core Media time structures, and extract Core Media time structure from archives.
//
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


// Decodes an object for the key, restricted to the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, objc.String(key))
	return rv
}

// The action the coder should take when decoding fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c_ Coder) DecodingFailurePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}




