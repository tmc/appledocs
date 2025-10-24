// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUUID */


/* debug [class_header]: Header for NSUUID */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UUID */
// An interface definition for the [UUID] class.
type IUUID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UUID */
	// properties:
	UUIDString() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UUID */
	// methods:
	Compare(otherUUID IUUID) ComparisonResult
	GetUUIDBytes(uuid objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UUID */
// Alloc allocates a new instance without initialization.
func (uc _UUIDClass) Alloc() UUID {
	rv := objc.Send[UUID](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UUID */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UUID */

// Initializes a new UUID with the given bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidBytes:)
func NewUUIDWithUUIDBytes(bytes objectivec.IObject) UUID {
	instance := getUUIDClass().Alloc()
	rv := objc.Send[UUID](instance.ID, objc.Sel("initWithUUIDBytes:"), bytes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUUIDWithUUIDBytes */


// Initializes a new UUID with the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidString:)
func NewUUIDWithUUIDString(string_ IString) UUID {
	instance := getUUIDClass().Alloc()
	rv := objc.Send[UUID](instance.ID, objc.Sel("initWithUUIDString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUUIDWithUUIDString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UUID */

// Create and returns a new UUID with RFC 4122 version 4 random bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/UUID
func (uc _UUIDClass) UUID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("UUID"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UUID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UUID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UUID */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/compare(_:)
func (u_ UUID) Compare(otherUUID IUUID) ComparisonResult {
	rv := objc.Send[ComparisonResult](u_.ID, objc.Sel("compare:"), otherUUID)
	return rv
}/* debug [instance_methods/method]: Compare */


// Returns the UUID as bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/getBytes(_:)
func (u_ UUID) GetUUIDBytes(uuid objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getUUIDBytes:"), uuid)
}/* debug [instance_methods/method]: GetUUIDBytes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UUID */

// The UUID as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID/uuidString
func (u_ UUID) UUIDString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("UUIDString"))
	return rv
}/* debug [instance_properties/getter]: UUIDString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUUID */


