// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectID] class.
var (
	ManagedObjectIDClass     _ManagedObjectIDClass
	ManagedObjectIDClassOnce sync.Once
)

func getManagedObjectIDClass() _ManagedObjectIDClass {
	ManagedObjectIDClassOnce.Do(func() {
		ManagedObjectIDClass = _ManagedObjectIDClass{objc.GetClass("NSManagedObjectID")}
	})
	return ManagedObjectIDClass
}

type _ManagedObjectIDClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectID] class.
type IManagedObjectID interface {
	objectivec.IObject
	Entity() IEntityDescription
	TemporaryID() bool
	PersistentStore() IPersistentStore
	IsTemporaryID() bool
	SetIsTemporaryID(value bool)
	URIRepresentation() foundation.URL
}

// A compact, universal identifier for a managed object.
//
// This identifier forms the basis for uniquing in the Core Data Framework. A managed object ID uniquely identifies the same managed object both between managed object contexts in a single application, and in multiple applications (as in distributed systems). Identifiers contain the information needed to exactly describe an object in a persistent store (like the primary key in the database), although the detailed information is not exposed. The framework completely encapsulates the “external” information and presents a clean object oriented interface. Object IDs can be transformed into a URI representation which can be archived and recreated later to refer back to a given object (using ( ) and ( ). For example, the last selected group in an application could be stored in the user defaults through the group object’s ID. You can also use object ID URI representations to store “weak” relationships across persistent stores (where no hard join is possible).


// A compact, universal identifier for a managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID
type ManagedObjectID struct {
	objectivec.Object
}

// ManagedObjectIDFrom constructs a [ManagedObjectID] from an unsafe.Pointer.
//
// A compact, universal identifier for a managed object.
func ManagedObjectIDFrom(ptr unsafe.Pointer) ManagedObjectID {
	return ManagedObjectID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectIDClass) Alloc() ManagedObjectID {
	rv := objc.Send[ManagedObjectID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectIDClass) New() ManagedObjectID {
	rv := objc.Send[ManagedObjectID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObjectID) Init() ManagedObjectID {
	rv := objc.Send[ManagedObjectID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObjectID) Autorelease() ManagedObjectID {
	rv := objc.Send[ManagedObjectID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObjectID creates a new ManagedObjectID instance.
func NewManagedObjectID() ManagedObjectID {
	return getManagedObjectIDClass().New()
}



// Returns a URI that provides an archiveable reference to the object for the object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/uriRepresentation()
func (m_ ManagedObjectID) URIRepresentation() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("URIRepresentation"))
	return rv
}


// The entity description associated with the object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/entity
func (m_ ManagedObjectID) Entity() IEntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("entity"))
	return rv
}


// A Boolean value that indicates whether the object ID is temporary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/isTemporaryID
func (m_ ManagedObjectID) TemporaryID() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("temporaryID"))
	return rv
}


// The persistent store that fetched the object for the object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/persistentStore
func (m_ ManagedObjectID) PersistentStore() IPersistentStore {
	rv := objc.Send[PersistentStore](m_.ID, objc.Sel("persistentStore"))
	return rv
}


// A Boolean value that indicates whether the object ID is temporary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/istemporaryid
func (m_ ManagedObjectID) IsTemporaryID() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isTemporaryID"))
	return rv
}


// A Boolean value that indicates whether the object ID is temporary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/istemporaryid
func (m_ ManagedObjectID) SetIsTemporaryID(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsTemporaryID:"), value)
}



