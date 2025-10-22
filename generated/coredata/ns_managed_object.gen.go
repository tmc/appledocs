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
	ManagedObjectClass     _ManagedObjectClass
	ManagedObjectClassOnce sync.Once
)

func getManagedObjectClass() _ManagedObjectClass {
	ManagedObjectClassOnce.Do(func() {
		ManagedObjectClass = _ManagedObjectClass{objc.GetClass("NSManagedObject")}
	})
	return ManagedObjectClass
}

type _ManagedObjectClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObject] class.
type IManagedObject interface {
	objectivec.IObject
	ChangedValuesForCurrentEvent() unsafe.Pointer
	ObjectIDsForRelationshipNamed(key string) []ManagedObjectID
	PrimitiveValueForKey(key string) objc.ID
	SetObservationInfo(inObservationInfo unsafe.Pointer)
	SetPrimitiveValueForKey(value objectivec.IObject, key string)
	WillSave()
	WillTurnIntoFault()
	Entity() NSEntityDescription
	FaultingState() uint
	HasChanges() bool
	HasPersistentChangedValues() bool
	Deleted() bool
	Fault() bool
	Inserted() bool
	Updated() bool
	ManagedObjectContext() NSManagedObjectContext
	ObjectID() NSManagedObjectID
	IsDeleted() bool
	SetIsDeleted(value bool)
	IsFault() bool
	SetIsFault(value bool)
	IsInserted() bool
	SetIsInserted(value bool)
	IsUpdated() bool
	SetIsUpdated(value bool)
	NSValidationKeyErrorKey() string
	NSValidationObjectErrorKey() string
	NSValidationPredicateErrorKey() string
	NSValidationValueErrorKey() string
	Description() string
	SetDescription(value string)
	Hash() int
	SetHash(value int)
	Superclass() unsafe.Pointer
	SetSuperclass(value unsafe.Pointer)
}

// The base class that all Core Data model objects inherit from.
//
// A managed object has an associated entity description ( ) that provides metadata about the object, including the name of the entity that the object represents and the names of its attributes and relationships. A managed object also has an associated managed object context that tracks changes to the object graph. You can’t use instances of direct subclasses of , or any other class that doesn’t inherit from , with a managed object context. You may create custom subclasses of , although this isn’t always necessary. If you don’t need custom logic, you can create a complete object graph with instances. If you instantiate a managed object directly, you must call the designated initializer .


// The base class that all Core Data model objects inherit from.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(entity:insertInto:)

func NewManagedObjectWithEntityInsertIntoManagedObjectContext(entity IEntityDescription, context IManagedObjectContext) ManagedObject {
	instance := getManagedObjectClass().Alloc()
	rv := objc.Send[ManagedObject](instance.ID, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), entity, context)
	rv.Autorelease()
	return rv
}



// Returns the entity description that is associated with this subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/entity()

func (mc _ManagedObjectClass) Entity() EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(mc.class), objc.Sel("entity"))
	return rv
}


// Returns an initialized fetch request with the entity this subclass represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/fetchRequest

func (mc _ManagedObjectClass) FetchRequest() FetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(mc.class), objc.Sel("fetchRequest"))
	return rv
}



// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/changedValuesForCurrentEvent()

func (m_ ManagedObject) ChangedValuesForCurrentEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("changedValuesForCurrentEvent"))
	return rv
}



// Provides an opportunity to respond when a value of a given property has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/didChangeValue(forKey:)

func (m_ ManagedObject) DidChangeValueForKey(key string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("didChangeValueForKey:"), objc.String(key))
}



// Returns the object IDs for all of the managed objects that are in the named relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/objectIDs(forRelationshipNamed:)

func (m_ ManagedObject) ObjectIDsForRelationshipNamed(key string) []ManagedObjectID {
	rv := objc.Send[[]ManagedObjectID](m_.ID, objc.Sel("objectIDsForRelationshipNamed:"), objc.String(key))
	return rv
}



// Returns the value for the specified property from the managed object’s private internal storage .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/primitiveValue(forKey:)

func (m_ ManagedObject) PrimitiveValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("primitiveValueForKey:"), objc.String(key))
	return rv
}



// Sets the observation info of the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/setObservationInfo(_:)

func (m_ ManagedObject) SetObservationInfo(inObservationInfo unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObservationInfo:"), inObservationInfo)
}



// Sets the value of a given property in the managed object’s private internal storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/setPrimitiveValue(_:forKey:)

func (m_ ManagedObject) SetPrimitiveValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveValue:forKey:"), value, objc.String(key))
}



// Returns the value for the property specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/value(forKey:)

func (m_ ManagedObject) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("valueForKey:"), objc.String(key))
	return rv
}



// Provides an opportunity to add code into the life cycle of the managed object before saving it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/willSave()

func (m_ ManagedObject) WillSave() {
	objc.Send[objc.ID](m_.ID, objc.Sel("willSave"))
}



// Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/willTurnIntoFault()

func (m_ ManagedObject) WillTurnIntoFault() {
	objc.Send[objc.ID](m_.ID, objc.Sel("willTurnIntoFault"))
}


// The entity description of the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/entity-swift.property

func (m_ ManagedObject) Entity() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](m_.ID, objc.Sel("entity"))
	return rv
}


// The faulting state of the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/faultingState

func (m_ ManagedObject) FaultingState() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("faultingState"))
	return rv
}


// A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/hasChanges

func (m_ ManagedObject) HasChanges() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasChanges"))
	return rv
}


// A Boolean value that indicates whether the managed object has persistent changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/hasPersistentChangedValues

func (m_ ManagedObject) HasPersistentChangedValues() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasPersistentChangedValues"))
	return rv
}


// A Boolean value that indicates whether the managed object will be deleted during the next save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/isDeleted

func (m_ ManagedObject) Deleted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deleted"))
	return rv
}


// A Boolean value that indicates whether the managed object is a fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/isFault

func (m_ ManagedObject) Fault() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("fault"))
	return rv
}


// A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/isInserted

func (m_ ManagedObject) Inserted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("inserted"))
	return rv
}


// A Boolean value that indicates whether the managed object has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/isUpdated

func (m_ ManagedObject) Updated() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("updated"))
	return rv
}


// The managed object context with which the managed object is registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/managedObjectContext

func (m_ ManagedObject) ManagedObjectContext() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m_.ID, objc.Sel("managedObjectContext"))
	return rv
}


// The object ID of the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObject/objectID

func (m_ ManagedObject) ObjectID() NSManagedObjectID {
	rv := objc.Send[NSManagedObjectID](m_.ID, objc.Sel("objectID"))
	return rv
}


// A Boolean value that indicates whether the managed object will be deleted during the next save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isdeleted

func (m_ ManagedObject) IsDeleted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDeleted"))
	return rv
}


// A Boolean value that indicates whether the managed object will be deleted during the next save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isdeleted

func (m_ ManagedObject) SetIsDeleted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDeleted:"), value)
}


// A Boolean value that indicates whether the managed object is a fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isfault

func (m_ ManagedObject) IsFault() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isFault"))
	return rv
}


// A Boolean value that indicates whether the managed object is a fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isfault

func (m_ ManagedObject) SetIsFault(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsFault:"), value)
}


// A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isinserted

func (m_ ManagedObject) IsInserted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isInserted"))
	return rv
}


// A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isinserted

func (m_ ManagedObject) SetIsInserted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsInserted:"), value)
}


// A Boolean value that indicates whether the managed object has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isupdated

func (m_ ManagedObject) IsUpdated() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdated"))
	return rv
}


// A Boolean value that indicates whether the managed object has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/isupdated

func (m_ ManagedObject) SetIsUpdated(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUpdated:"), value)
}


// The error key for the attribute that failed to validate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsvalidationkeyerrorkey

func (m_ ManagedObject) NSValidationKeyErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSValidationKeyErrorKey"))
	return rv
}


// The error key for the object that failed to validate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsvalidationobjecterrorkey

func (m_ ManagedObject) NSValidationObjectErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSValidationObjectErrorKey"))
	return rv
}


// The error key for the predicate that failed to validate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsvalidationpredicateerrorkey

func (m_ ManagedObject) NSValidationPredicateErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSValidationPredicateErrorKey"))
	return rv
}


// The error key for the value that failed to validate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsvalidationvalueerrorkey

func (m_ ManagedObject) NSValidationValueErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSValidationValueErrorKey"))
	return rv
}


// A textual representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description

func (m_ ManagedObject) Description() string {
	rv := objc.Send[string](m_.ID, objc.Sel("description"))
	return rv
}


// A textual representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description

func (m_ ManagedObject) SetDescription(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescription:"), objc.String(value))
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash

func (m_ ManagedObject) Hash() int {
	rv := objc.Send[int](m_.ID, objc.Sel("hash"))
	return rv
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash

func (m_ ManagedObject) SetHash(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHash:"), value)
}


// Returns the class object for the receiver’s superclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/superclass

func (m_ ManagedObject) Superclass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("superclass"))
	return rv
}


// Returns the class object for the receiver’s superclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/superclass

func (m_ ManagedObject) SetSuperclass(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSuperclass:"), value)
}


