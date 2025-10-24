// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	VersionHash() objc.IObject /* cross-framework: NSData */
	AttributesByName() IAttributeDescription
	SetAttributesByName(value IAttributeDescription)
	CompoundIndexes() unsafe.Pointer
	SetCompoundIndexes(value unsafe.Pointer)
	CoreSpotlightDisplayNameExpression() objc.IObject /* cross-framework: Expression */
	SetCoreSpotlightDisplayNameExpression(value objc.IObject /* cross-framework: Expression */)
	Indexes() FetchIndexDescription /* not a class type */
	SetIndexes(value FetchIndexDescription /* not a class type */)
	IsAbstract() bool
	SetIsAbstract(value bool)
	ManagedObjectClassName() objc.IObject /* cross-framework: NSString */
	SetManagedObjectClassName(value objc.IObject /* cross-framework: NSString */)
	ManagedObjectModel() IManagedObjectModel
	SetManagedObjectModel(value IManagedObjectModel)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Properties() IPropertyDescription
	SetProperties(value IPropertyDescription)
	PropertiesByName() IPropertyDescription
	SetPropertiesByName(value IPropertyDescription)
	RelationshipsByName() IRelationshipDescription
	SetRelationshipsByName(value IRelationshipDescription)
	RenamingIdentifier() objc.IObject /* cross-framework: NSString */
	SetRenamingIdentifier(value objc.IObject /* cross-framework: NSString */)
	Subentities() IEntityDescription
	SetSubentities(value IEntityDescription)
	SubentitiesByName() IEntityDescription
	SetSubentitiesByName(value IEntityDescription)
	Superentity() IEntityDescription
	SetSuperentity(value IEntityDescription)
	UniquenessConstraints() unsafe.Pointer
	SetUniquenessConstraints(value unsafe.Pointer)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	VersionHashModifier() objc.IObject /* cross-framework: NSString */
	SetVersionHashModifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A description of a Core Data entity.
//
// Entities are to managed objects what is to , or — to use a database analogy — what tables are to rows. An instance specifies the entity’s name, its attributes and relationships (as instances of and ) and the class that represents it. Instances of that class correspond to entries in the associated persistent store. As a minimum, an entity description requires: A name. The class name of the corresponding managed object. If you don’t specify a class name, the framework uses . You define entities in a managed object model (an instance of ) using Xcode’s data modeling tool. Core Data uses to map entries in the persistent store to managed objects in your app. It’s unlikely you’ll interact with entity descriptions directly unless you’re specifically working with models. provides a user dictionary for you to store any related, app-specific information.


// A description of a Core Data entity.
//
// [Full Topic]
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



// The version hash for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription/versionHash
func (e_ EntityDescription) VersionHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("versionHash"))
	return rv
}


// The attributes of the receiver in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/attributesbyname
func (e_ EntityDescription) AttributesByName() IAttributeDescription {
	rv := objc.Send[AttributeDescription](e_.ID, objc.Sel("attributesByName"))
	return rv
}


// The attributes of the receiver in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/attributesbyname
func (e_ EntityDescription) SetAttributesByName(value IAttributeDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributesByName:"), value)
}


// The compound indexes for the entity as an array of arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/compoundindexes
func (e_ EntityDescription) CompoundIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("compoundIndexes"))
	return rv
}


// The compound indexes for the entity as an array of arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/compoundindexes
func (e_ EntityDescription) SetCompoundIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompoundIndexes:"), value)
}


// The expression that computes the CoreSpotlight display name for instances of the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/corespotlightdisplaynameexpression
func (e_ EntityDescription) CoreSpotlightDisplayNameExpression() objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](e_.ID, objc.Sel("coreSpotlightDisplayNameExpression"))
	return rv
}


// The expression that computes the CoreSpotlight display name for instances of the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/corespotlightdisplaynameexpression
func (e_ EntityDescription) SetCoreSpotlightDisplayNameExpression(value objc.IObject /* cross-framework: Expression */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCoreSpotlightDisplayNameExpression:"), value)
}


// An array of fetch index descriptions for the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/indexes
func (e_ EntityDescription) Indexes() FetchIndexDescription /* not a class type */ {
	rv := objc.Send[FetchIndexDescription](e_.ID, objc.Sel("indexes"))
	return rv
}


// An array of fetch index descriptions for the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/indexes
func (e_ EntityDescription) SetIndexes(value FetchIndexDescription /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIndexes:"), value)
}


// A Boolean value that indicates whether the receiver represents an abstract entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/isabstract
func (e_ EntityDescription) IsAbstract() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isAbstract"))
	return rv
}


// A Boolean value that indicates whether the receiver represents an abstract entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/isabstract
func (e_ EntityDescription) SetIsAbstract(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsAbstract:"), value)
}


// The name of the class that represents the receiver’s entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectclassname
func (e_ EntityDescription) ManagedObjectClassName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("managedObjectClassName"))
	return rv
}


// The name of the class that represents the receiver’s entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectclassname
func (e_ EntityDescription) SetManagedObjectClassName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setManagedObjectClassName:"), value)
}


// The managed object model with which the receiver is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectmodel
func (e_ EntityDescription) ManagedObjectModel() IManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](e_.ID, objc.Sel("managedObjectModel"))
	return rv
}


// The managed object model with which the receiver is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/managedobjectmodel
func (e_ EntityDescription) SetManagedObjectModel(value IManagedObjectModel) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setManagedObjectModel:"), value)
}


// The entity name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/name
func (e_ EntityDescription) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}


// The entity name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/name
func (e_ EntityDescription) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}


// An array containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ EntityDescription) Properties() IPropertyDescription {
	rv := objc.Send[PropertyDescription](e_.ID, objc.Sel("properties"))
	return rv
}


// An array containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ EntityDescription) SetProperties(value IPropertyDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProperties:"), value)
}


// A dictionary containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/propertiesbyname
func (e_ EntityDescription) PropertiesByName() IPropertyDescription {
	rv := objc.Send[PropertyDescription](e_.ID, objc.Sel("propertiesByName"))
	return rv
}


// A dictionary containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/propertiesbyname
func (e_ EntityDescription) SetPropertiesByName(value IPropertyDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPropertiesByName:"), value)
}


// The relationships of the receiver in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/relationshipsbyname
func (e_ EntityDescription) RelationshipsByName() IRelationshipDescription {
	rv := objc.Send[RelationshipDescription](e_.ID, objc.Sel("relationshipsByName"))
	return rv
}


// The relationships of the receiver in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/relationshipsbyname
func (e_ EntityDescription) SetRelationshipsByName(value IRelationshipDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelationshipsByName:"), value)
}


// The renaming identifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/renamingidentifier
func (e_ EntityDescription) RenamingIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("renamingIdentifier"))
	return rv
}


// The renaming identifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/renamingidentifier
func (e_ EntityDescription) SetRenamingIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRenamingIdentifier:"), value)
}


// An array containing the sub-entities of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentities
func (e_ EntityDescription) Subentities() IEntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("subentities"))
	return rv
}


// An array containing the sub-entities of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentities
func (e_ EntityDescription) SetSubentities(value IEntityDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSubentities:"), value)
}


// A dictionary containing the receiver’s sub-entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentitiesbyname
func (e_ EntityDescription) SubentitiesByName() IEntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("subentitiesByName"))
	return rv
}


// A dictionary containing the receiver’s sub-entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/subentitiesbyname
func (e_ EntityDescription) SetSubentitiesByName(value IEntityDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSubentitiesByName:"), value)
}


// The super-entity of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/superentity
func (e_ EntityDescription) Superentity() IEntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("superentity"))
	return rv
}


// The super-entity of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/superentity
func (e_ EntityDescription) SetSuperentity(value IEntityDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSuperentity:"), value)
}


// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/uniquenessconstraints
func (e_ EntityDescription) UniquenessConstraints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("uniquenessConstraints"))
	return rv
}


// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/uniquenessconstraints
func (e_ EntityDescription) SetUniquenessConstraints(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUniquenessConstraints:"), value)
}


// The user info dictionary of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/userinfo
func (e_ EntityDescription) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/userinfo
func (e_ EntityDescription) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}


// The version hash modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/versionhashmodifier
func (e_ EntityDescription) VersionHashModifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("versionHashModifier"))
	return rv
}


// The version hash modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/versionhashmodifier
func (e_ EntityDescription) SetVersionHashModifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVersionHashModifier:"), value)
}



