// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityMapping] class.
var (
	EntityMappingClass     _EntityMappingClass
	EntityMappingClassOnce sync.Once
)

func getEntityMappingClass() _EntityMappingClass {
	EntityMappingClassOnce.Do(func() {
		EntityMappingClass = _EntityMappingClass{objc.GetClass("NSEntityMapping")}
	})
	return EntityMappingClass
}

type _EntityMappingClass struct {
	class objc.Class
}

// An interface definition for the [EntityMapping] class.
type IEntityMapping interface {
	objectivec.IObject
	AttributeMappings() []PropertyMapping
	SetAttributeMappings(value []PropertyMapping)
	DestinationEntityName() string
	SetDestinationEntityName(value string)
	DestinationEntityVersionHash() foundation.NSData
	SetDestinationEntityVersionHash(value foundation.IData)
	EntityMigrationPolicyClassName() string
	SetEntityMigrationPolicyClassName(value string)
	MappingType() EntityMappingType
	SetMappingType(value EntityMappingType)
	Name() string
	SetName(value string)
	RelationshipMappings() []PropertyMapping
	SetRelationshipMappings(value []PropertyMapping)
	SourceEntityName() string
	SetSourceEntityName(value string)
	SourceEntityVersionHash() foundation.NSData
	SetSourceEntityVersionHash(value foundation.IData)
	SourceExpression() Expression
	SetSourceExpression(value IExpression)
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
}

// A mapping instance that specifies how to map an entity from a source to a destination managed object model.


// A mapping instance that specifies how to map an entity from a source to a destination managed object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping

type EntityMapping struct {
	objectivec.Object
}

// EntityMappingFrom constructs a [EntityMapping] from an unsafe.Pointer.
//
// A mapping instance that specifies how to map an entity from a source to a destination managed object model.
func EntityMappingFrom(ptr unsafe.Pointer) EntityMapping {
	return EntityMapping{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EntityMappingClass) Alloc() EntityMapping {
	rv := objc.Send[EntityMapping](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityMappingClass) New() EntityMapping {
	rv := objc.Send[EntityMapping](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EntityMapping) Init() EntityMapping {
	rv := objc.Send[EntityMapping](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EntityMapping) Autorelease() EntityMapping {
	rv := objc.Send[EntityMapping](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntityMapping creates a new EntityMapping instance.
func NewEntityMapping() EntityMapping {
	return getEntityMappingClass().New()
}



// The array of attribute mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/attributeMappings

func (e_ EntityMapping) AttributeMappings() []PropertyMapping {
	rv := objc.Send[[]PropertyMapping](e_.ID, objc.Sel("attributeMappings"))
	return rv
}


// The array of attribute mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/attributeMappings

func (e_ EntityMapping) SetAttributeMappings(value []PropertyMapping) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributeMappings:"), nsArray)
}


// The destination entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityName

func (e_ EntityMapping) DestinationEntityName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("destinationEntityName"))
	return rv
}


// The destination entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityName

func (e_ EntityMapping) SetDestinationEntityName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityName:"), objc.String(value))
}


// The version hash for the destination entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityVersionHash

func (e_ EntityMapping) DestinationEntityVersionHash() foundation.NSData {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("destinationEntityVersionHash"))
	return rv
}


// The version hash for the destination entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityVersionHash

func (e_ EntityMapping) SetDestinationEntityVersionHash(value foundation.IData) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityVersionHash:"), value)
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/entityMigrationPolicyClassName

func (e_ EntityMapping) EntityMigrationPolicyClassName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("entityMigrationPolicyClassName"))
	return rv
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/entityMigrationPolicyClassName

func (e_ EntityMapping) SetEntityMigrationPolicyClassName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEntityMigrationPolicyClassName:"), objc.String(value))
}


// The mapping type for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/mappingType

func (e_ EntityMapping) MappingType() EntityMappingType {
	rv := objc.Send[EntityMappingType](e_.ID, objc.Sel("mappingType"))
	return rv
}


// The mapping type for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/mappingType

func (e_ EntityMapping) SetMappingType(value EntityMappingType) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMappingType:"), value)
}


// The name of the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/name

func (e_ EntityMapping) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// The name of the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/name

func (e_ EntityMapping) SetName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), objc.String(value))
}


// The array of relationship mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/relationshipMappings

func (e_ EntityMapping) RelationshipMappings() []PropertyMapping {
	rv := objc.Send[[]PropertyMapping](e_.ID, objc.Sel("relationshipMappings"))
	return rv
}


// The array of relationship mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/relationshipMappings

func (e_ EntityMapping) SetRelationshipMappings(value []PropertyMapping) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelationshipMappings:"), nsArray)
}


// The source entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityName

func (e_ EntityMapping) SourceEntityName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("sourceEntityName"))
	return rv
}


// The source entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityName

func (e_ EntityMapping) SetSourceEntityName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityName:"), objc.String(value))
}


// The version hash of the source entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityVersionHash

func (e_ EntityMapping) SourceEntityVersionHash() foundation.NSData {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("sourceEntityVersionHash"))
	return rv
}


// The version hash of the source entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityVersionHash

func (e_ EntityMapping) SetSourceEntityVersionHash(value foundation.IData) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityVersionHash:"), value)
}


// The source expression for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceExpression

func (e_ EntityMapping) SourceExpression() Expression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("sourceExpression"))
	return rv
}


// The source expression for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceExpression

func (e_ EntityMapping) SetSourceExpression(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceExpression:"), value)
}


// The user info dictionary for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/userInfo

func (e_ EntityMapping) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/userInfo

func (e_ EntityMapping) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}



