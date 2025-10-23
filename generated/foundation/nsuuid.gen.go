// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UUID] class.
var (
	UUIDClass     _UUIDClass
	UUIDClassOnce sync.Once
)

func getUUIDClass() _UUIDClass {
	UUIDClassOnce.Do(func() {
		UUIDClass = _UUIDClass{objc.GetClass("NSUUID")}
	})
	return UUIDClass
}

type _UUIDClass struct {
	class objc.Class
}

// An interface definition for the [UUID] class.
type IUUID interface {
	objectivec.IObject
	// properties:
	UUIDString() string /* primitive/slice/pointer */
	// methods:
	Compare(otherUUID IUUID) NSComparisonResult /* foo */
	GetUUIDBytes(uuid unsafe.Pointer)
}

// A universally unique value that can be used to identify types, interfaces, and other items.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. UUIDs (Universally Unique Identifiers), also known as GUIDs (Globally Unique Identifiers) or IIDs (Interface Identifiers), are 128-bit values. UUIDs created by conform to RFC 4122 version 4 and are created with random bytes. The standard format for UUIDs represented in ASCII is a string punctuated by hyphens, for example . The hex representation looks, as you might expect, like a list of numerical values preceded by 0x. For example, , , , , , , , , , , , , , , , . Because a UUID is expressed simply as an array of bytes, there are no endianness considerations for different platforms. The class is toll-free bridged with CoreFoundation’s . Use UUID strings to convert between and , if needed. Two objects are not guaranteed to be comparable by pointer value (as is); use to compare two instances.


// A universally unique value that can be used to identify types, interfaces, and other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID
type UUID struct {
	objectivec.Object
}

// UUIDFrom constructs a [UUID] from an unsafe.Pointer.
//
// A universally unique value that can be used to identify types, interfaces, and other items.
func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UUIDClass) Alloc() UUID {
	rv := objc.Send[UUID](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UUIDClass) New() UUID {
	rv := objc.Send[UUID](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UUID) Init() UUID {
	rv := objc.Send[UUID](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UUID) Autorelease() UUID {
	rv := objc.Send[UUID](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUUID creates a new UUID instance.
func NewUUID() UUID {
	return getUUIDClass().New()
}



// Initializes a new UUID with the given bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidBytes:)
func NewUUIDWithUUIDBytes(bytes unsafe.Pointer) UUID {
	instance := getUUIDClass().Alloc()
	rv := objc.Send[UUID](instance.ID, objc.Sel("initWithUUIDBytes:"), bytes)
	rv.Autorelease()
	return rv
}


// Initializes a new UUID with the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidString:)
func NewUUIDWithUUIDString(string_ string /* primitive/slice/pointer */) UUID {
	instance := getUUIDClass().Alloc()
	rv := objc.Send[UUID](instance.ID, objc.Sel("initWithUUIDString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Create and returns a new UUID with RFC 4122 version 4 random bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/UUID
func (uc _UUIDClass) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("UUID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/compare(_:)
func (u_ UUID) Compare(otherUUID IUUID) NSComparisonResult /* foo */ {
	rv := objc.Send[ComparisonResult](u_.ID, objc.Sel("compare:"), otherUUID)
	return rv
}


// Returns the UUID as bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/getBytes(_:)
func (u_ UUID) GetUUIDBytes(uuid unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getUUIDBytes:"), uuid)
}


// The UUID as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/uuidString
func (u_ UUID) UUIDString() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("UUIDString"))
	return rv
}


