// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectID] class.
var (
	managedObjectIDClass     _ManagedObjectIDClass
	managedObjectIDClassOnce sync.Once
)

func getManagedObjectIDClass() _ManagedObjectIDClass {
	managedObjectIDClassOnce.Do(func() {
		managedObjectIDClass = _ManagedObjectIDClass{objc.GetClass("NSManagedObjectID")}
	})
	return managedObjectIDClass
}

type _ManagedObjectIDClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectID] class.
type IManagedObjectID interface {
	objectivec.IObject
	URIRepresentation() unsafe.Pointer
}

// A compact, universal identifier for a managed object. [Full Topic]
//
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


// Returns a URI that provides an archiveable reference to the object for the object ID. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/uriRepresentation()
func (m_ ManagedObjectID) URIRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("URIRepresentation"))
	return rv
}


