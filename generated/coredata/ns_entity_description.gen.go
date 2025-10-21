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

// The attributes of the receiver in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/attributesbyname
func (e_ EntityDescription) AttributesByName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("attributesByName"))
	return rv
}


// SetAttributesByName sets the value of the attributesByName property.
// The attributes of the receiver in a dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/attributesbyname
func (e_ EntityDescription) SetAttributesByName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributesByName:"), objc.String(value))
}

// The compound indexes for the entity as an array of arrays.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/compoundindexes
func (e_ EntityDescription) CompoundIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("compoundIndexes"))
	return rv
}


// SetCompoundIndexes sets the value of the compoundIndexes property.
// The compound indexes for the entity as an array of arrays.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/compoundindexes
func (e_ EntityDescription) SetCompoundIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompoundIndexes:"), value)
}

// The expression that computes the CoreSpotlight display name for instances of the entity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/corespotlightdisplaynameexpression
func (e_ EntityDescription) CoreSpotlightDisplayNameExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("coreSpotlightDisplayNameExpression"))
	return rv
}


// SetCoreSpotlightDisplayNameExpression sets the value of the coreSpotlightDisplayNameExpression property.
// The expression that computes the CoreSpotlight display name for instances of the entity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/corespotlightdisplaynameexpression
func (e_ EntityDescription) SetCoreSpotlightDisplayNameExpression(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCoreSpotlightDisplayNameExpression:"), value)
}

// An array of fetch index descriptions for the entity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/indexes
func (e_ EntityDescription) Indexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("indexes"))
	return rv
}


// SetIndexes sets the value of the indexes property.
// An array of fetch index descriptions for the entity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/indexes
func (e_ EntityDescription) SetIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIndexes:"), value)
}

// A Boolean value that indicates whether the receiver represents an abstract entity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/isabstract
func (e_ EntityDescription) IsAbstract() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isAbstract"))
	return rv
}


// SetIsAbstract sets the value of the isAbstract property.
// A Boolean value that indicates whether the receiver represents an abstract entity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/isabstract
func (e_ EntityDescription) SetIsAbstract(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsAbstract:"), value)
}

// The name of the class that represents the receiver’s entity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectclassname
func (e_ EntityDescription) ManagedObjectClassName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("managedObjectClassName"))
	return rv
}


// SetManagedObjectClassName sets the value of the managedObjectClassName property.
// The name of the class that represents the receiver’s entity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectclassname
func (e_ EntityDescription) SetManagedObjectClassName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setManagedObjectClassName:"), objc.String(value))
}

// The managed object model with which the receiver is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectmodel
func (e_ EntityDescription) ManagedObjectModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("managedObjectModel"))
	return rv
}


// SetManagedObjectModel sets the value of the managedObjectModel property.
// The managed object model with which the receiver is associated.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectmodel
func (e_ EntityDescription) SetManagedObjectModel(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setManagedObjectModel:"), value)
}

// An array containing the properties of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ EntityDescription) Properties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("properties"))
	return rv
}


// SetProperties sets the value of the properties property.
// An array containing the properties of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ EntityDescription) SetProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProperties:"), value)
}

// A dictionary containing the properties of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/propertiesbyname
func (e_ EntityDescription) PropertiesByName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("propertiesByName"))
	return rv
}


// SetPropertiesByName sets the value of the propertiesByName property.
// A dictionary containing the properties of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/propertiesbyname
func (e_ EntityDescription) SetPropertiesByName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPropertiesByName:"), objc.String(value))
}

// The relationships of the receiver in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/relationshipsbyname
func (e_ EntityDescription) RelationshipsByName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("relationshipsByName"))
	return rv
}


// SetRelationshipsByName sets the value of the relationshipsByName property.
// The relationships of the receiver in a dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/relationshipsbyname
func (e_ EntityDescription) SetRelationshipsByName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelationshipsByName:"), objc.String(value))
}

// The renaming identifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/renamingidentifier
func (e_ EntityDescription) RenamingIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("renamingIdentifier"))
	return rv
}


// SetRenamingIdentifier sets the value of the renamingIdentifier property.
// The renaming identifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/renamingidentifier
func (e_ EntityDescription) SetRenamingIdentifier(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRenamingIdentifier:"), objc.String(value))
}

// An array containing the sub-entities of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentities
func (e_ EntityDescription) Subentities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("subentities"))
	return rv
}


// SetSubentities sets the value of the subentities property.
// An array containing the sub-entities of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentities
func (e_ EntityDescription) SetSubentities(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSubentities:"), value)
}

// A dictionary containing the receiver’s sub-entities.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentitiesbyname
func (e_ EntityDescription) SubentitiesByName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("subentitiesByName"))
	return rv
}


// SetSubentitiesByName sets the value of the subentitiesByName property.
// A dictionary containing the receiver’s sub-entities.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentitiesbyname
func (e_ EntityDescription) SetSubentitiesByName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSubentitiesByName:"), objc.String(value))
}

// The super-entity of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/superentity
func (e_ EntityDescription) Superentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("superentity"))
	return rv
}


// SetSuperentity sets the value of the superentity property.
// The super-entity of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/superentity
func (e_ EntityDescription) SetSuperentity(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSuperentity:"), value)
}

// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/uniquenessconstraints
func (e_ EntityDescription) UniquenessConstraints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("uniquenessConstraints"))
	return rv
}


// SetUniquenessConstraints sets the value of the uniquenessConstraints property.
// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/uniquenessconstraints
func (e_ EntityDescription) SetUniquenessConstraints(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUniquenessConstraints:"), value)
}

// The user info dictionary of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/userinfo
func (e_ EntityDescription) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// The user info dictionary of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/userinfo
func (e_ EntityDescription) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}

// The version hash modifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/versionhashmodifier
func (e_ EntityDescription) VersionHashModifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("versionHashModifier"))
	return rv
}


// SetVersionHashModifier sets the value of the versionHashModifier property.
// The version hash modifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/versionhashmodifier
func (e_ EntityDescription) SetVersionHashModifier(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVersionHashModifier:"), objc.String(value))
}



