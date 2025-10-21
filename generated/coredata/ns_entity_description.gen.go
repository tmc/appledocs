// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityDescription] class.
var (
	EntityDescriptionClass     _EntityDescriptionClass
	EntityDescriptionClassOnce sync.Once
)

func getEntityDescriptionClass() _EntityDescriptionClass {
	EntityDescriptionClassOnce.Do(func() {
		EntityDescriptionClass = _EntityDescriptionClass{objc.GetClass("NSEntityDescription")}
	})
	return EntityDescriptionClass
}

type _EntityDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [EntityDescription] class.
type IEntityDescription interface {
	objectivec.IObject
}

// A description of a Core Data entity.
//
// Entities are to managed objects what is to , or — to use a database analogy — what tables are to rows. An instance specifies the entity’s name, its attributes and relationships (as instances of and ) and the class that represents it. Instances of that class correspond to entries in the associated persistent store. As a minimum, an entity description requires: A name. The class name of the corresponding managed object. If you don’t specify a class name, the framework uses . You define entities in a managed object model (an instance of ) using Xcode’s data modeling tool. Core Data uses to map entries in the persistent store to managed objects in your app. It’s unlikely you’ll interact with entity descriptions directly unless you’re specifically working with models. provides a user dictionary for you to store any related, app-specific information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription
type EntityDescription struct {
	objectivec.Object
}

// EntityDescriptionFrom constructs a [EntityDescription] from an unsafe.Pointer.
//
// A description of a Core Data entity.
func EntityDescriptionFrom(ptr unsafe.Pointer) EntityDescription {
	return EntityDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EntityDescriptionClass) Alloc() EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityDescriptionClass) New() EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EntityDescription) Init() EntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EntityDescription) Autorelease() EntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntityDescription creates a new EntityDescription instance.
func NewEntityDescription() EntityDescription {
	return getEntityDescriptionClass().New()
}


// Creates, configures, and returns an instance of the class for the entity with a given name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription/insertNewObject(forEntityName:into:)
func (ec _EntityDescriptionClass) InsertNewObjectForEntityForNameInManagedObjectContext(entityName string, context unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("insertNewObjectForEntityForName:inManagedObjectContext:"), objc.String(entityName), context)
	return rv
}

// The entity name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription/name
func (e_ EntityDescription) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The entity name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription/name
func (e_ EntityDescription) SetName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), objc.String(value))
}

// The version hash for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription/versionHash
func (e_ EntityDescription) VersionHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("versionHash"))
	return rv
}



