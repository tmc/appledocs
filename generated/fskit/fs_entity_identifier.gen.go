// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSEntityIdentifier */


/* debug [class_header]: Header for FSEntityIdentifier */
// The class instance for the [FSEntityIdentifier] class.
var (
	FSEntityIdentifierClass     _FSEntityIdentifierClass
	FSEntityIdentifierClassOnce sync.Once
)

func getFSEntityIdentifierClass() _FSEntityIdentifierClass {
	FSEntityIdentifierClassOnce.Do(func() {
		FSEntityIdentifierClass = _FSEntityIdentifierClass{objc.GetClass("FSEntityIdentifier")}
	})
	return FSEntityIdentifierClass
}

type _FSEntityIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSEntityIdentifier */
// An interface definition for the [FSEntityIdentifier] class.
type IFSEntityIdentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSEntityIdentifier */
	// properties:
	Qualifier() objc.IObject /* cross-framework: NSData */
	SetQualifier(value objc.IObject /* cross-framework: NSData */)
	Uuid() foundation.UUID
	SetUuid(value foundation.UUID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSEntityIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSEntityIdentifier */
// Alloc allocates a new instance without initialization.
func (fc _FSEntityIdentifierClass) Alloc() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSEntityIdentifierClass) New() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSEntityIdentifier) Init() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSEntityIdentifier) Autorelease() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSEntityIdentifier creates a new FSEntityIdentifier instance.
func NewFSEntityIdentifier() FSEntityIdentifier {
	return getFSEntityIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSEntityIdentifier */
// A base type that identifies containers and volumes.
//
// An is a UUID to identify a container or volume, optionally with eight bytes of qualifying (differentiating) data. You use the qualifiers in cases in which a file server can receive multiple connections from the same client, which differ by user credentials. In this case, the identifier for each client is the server’s base UUID, and a unique qualifier that differs by client.


// A base type that identifies containers and volumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier
type FSEntityIdentifier struct {
	objectivec.Object
}

// FSEntityIdentifierFrom constructs a [FSEntityIdentifier] from an unsafe.Pointer.
//
// A base type that identifies containers and volumes.
func FSEntityIdentifierFrom(ptr unsafe.Pointer) FSEntityIdentifier {
	return FSEntityIdentifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSEntityIdentifier */

// Creates an entity identifier with the given UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:)
func NewFSEntityIdentifierWithUUID(uuid foundation.UUID) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[FSEntityIdentifier](instance.ID, objc.Sel("initWithUUID:"), uuid)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSEntityIdentifierWithUUID */


// Creates an entity identifier with the given UUID and qualifier data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)
func NewFSEntityIdentifierWithUUIDData(uuid foundation.UUID, qualifierData objc.IObject /* cross-framework: NSData */) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[FSEntityIdentifier](instance.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSEntityIdentifierWithUUIDData */


// Creates an entity identifier with the given UUID and qualifier data as a 64-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:qualifier:)
func NewFSEntityIdentifierWithUUIDQualifier(uuid foundation.UUID, qualifier uint64) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[FSEntityIdentifier](instance.ID, objc.Sel("initWithUUID:qualifier:"), uuid, qualifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSEntityIdentifierWithUUIDQualifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSEntityIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSEntityIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSEntityIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSEntityIdentifier */

// An optional piece of data to distinguish entities that otherwise share the same UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/qualifier
func (f_ FSEntityIdentifier) Qualifier() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("qualifier"))
	return rv
}/* debug [instance_properties/getter]: qualifier */


// An optional piece of data to distinguish entities that otherwise share the same UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/qualifier
func (f_ FSEntityIdentifier) SetQualifier(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setQualifier:"), value)
}/* debug [instance_properties/setter]: qualifier */


// A UUID to uniquely identify this entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/uuid
func (f_ FSEntityIdentifier) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](f_.ID, objc.Sel("uuid"))
	return rv
}/* debug [instance_properties/getter]: uuid */


// A UUID to uniquely identify this entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/uuid
func (f_ FSEntityIdentifier) SetUuid(value foundation.UUID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUuid:"), value)
}/* debug [instance_properties/setter]: uuid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSEntityIdentifier */


