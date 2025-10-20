// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObject] class.
var (
	managedObjectClass     _ManagedObjectClass
	managedObjectClassOnce sync.Once
)

func getManagedObjectClass() _ManagedObjectClass {
	managedObjectClassOnce.Do(func() {
		managedObjectClass = _ManagedObjectClass{objc.GetClass("NSManagedObject")}
	})
	return managedObjectClass
}

type _ManagedObjectClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObject] class.
type IManagedObject interface {
	objectivec.IObject
}

// The base class that all Core Data model objects inherit from.
//
// A managed object has an associated entity description ( ) that provides metadata about the object, including the name of the entity that the object represents and the names of its attributes and relationships. A managed object also has an associated managed object context that tracks changes to the object graph. You can’t use instances of direct subclasses of , or any other class that doesn’t inherit from , with a managed object context. You may create custom subclasses of , although this isn’t always necessary. If you don’t need custom logic, you can create a complete object graph with instances. If you instantiate a managed object directly, you must call the designated initializer .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject
type ManagedObject struct {
	objectivec.Object
}

// ManagedObjectFrom constructs a [ManagedObject] from an unsafe.Pointer.
//
// The base class that all Core Data model objects inherit from.
func ManagedObjectFrom(ptr unsafe.Pointer) ManagedObject {
	return ManagedObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectClass) Alloc() ManagedObject {
	rv := objc.Send[ManagedObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectClass) New() ManagedObject {
	rv := objc.Send[ManagedObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObject) Init() ManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObject) Autorelease() ManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObject creates a new ManagedObject instance.
func NewManagedObject() ManagedObject {
	return getManagedObjectClass().New()
}


// Initializes a managed object from an entity description and inserts it into the specified managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(entity:insertInto:)
func NewManagedObjectWithEntityInsertIntoManagedObjectContext(entity unsafe.Pointer, context unsafe.Pointer) ManagedObject {
	instance := getManagedObjectClass().Alloc()
	rv := objc.Send[ManagedObject](instance.ID, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), entity, context)
	rv.Autorelease()
	return rv
}


// The entity description of the managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/entity-swift.property
func (m_ ManagedObject) Entity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("entity"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/isInserted
func (m_ ManagedObject) Inserted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("inserted"))
	return rv
}

// The managed object context with which the managed object is registered.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/managedObjectContext
func (m_ ManagedObject) ManagedObjectContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("managedObjectContext"))
	return rv
}

// The object ID of the managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/objectID
func (m_ ManagedObject) ObjectID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectID"))
	return rv
}


