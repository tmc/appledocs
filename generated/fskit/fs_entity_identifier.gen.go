// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FSEntityIdentifier] class.
type IFSEntityIdentifier interface {
	objectivec.IObject
	Qualifier() foundation.Data
	SetQualifier(value foundation.Data)
	Uuid() foundation.UUID
	SetUuid(value foundation.UUID)
}

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

// Alloc allocates a new instance without initialization.
func (fc _FSEntityIdentifierClass) Alloc() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an entity identifier with the given UUID and qualifier data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)
func NewFSEntityIdentifierWithUUIDData(uuid foundation.UUID, qualifierData foundation.NSData) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[FSEntityIdentifier](instance.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	rv.Autorelease()
	return rv
}



// An optional piece of data to distinguish entities that otherwise share the same UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsentityidentifier/qualifier
func (f_ FSEntityIdentifier) Qualifier() foundation.Data {
	rv := objc.Send[foundation.Data](f_.ID, objc.Sel("qualifier"))
	return rv
}


// An optional piece of data to distinguish entities that otherwise share the same UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsentityidentifier/qualifier
func (f_ FSEntityIdentifier) SetQualifier(value foundation.Data) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setQualifier:"), value)
}


// A UUID to uniquely identify this entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsentityidentifier/uuid
func (f_ FSEntityIdentifier) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](f_.ID, objc.Sel("uuid"))
	return rv
}


// A UUID to uniquely identify this entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsentityidentifier/uuid
func (f_ FSEntityIdentifier) SetUuid(value foundation.UUID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUuid:"), value)
}


