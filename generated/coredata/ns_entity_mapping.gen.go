// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A mapping instance that specifies how to map an entity from a source to a destination managed object model.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/attributeMappings
func (e_ EntityMapping) AttributeMappings() []PropertyMapping {
	rv := objc.Send[[]PropertyMapping](e_.ID, objc.Sel("attributeMappings"))
	return rv
}


// SetAttributeMappings sets the value of the attributeMappings property.
// The array of attribute mappings for the entity mapping.

//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityName
func (e_ EntityMapping) DestinationEntityName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("destinationEntityName"))
	return rv
}


// SetDestinationEntityName sets the value of the destinationEntityName property.
// The destination entity name for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityName
func (e_ EntityMapping) SetDestinationEntityName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityName:"), objc.String(value))
}

// The version hash for the destination entity for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityVersionHash
func (e_ EntityMapping) DestinationEntityVersionHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("destinationEntityVersionHash"))
	return rv
}


// SetDestinationEntityVersionHash sets the value of the destinationEntityVersionHash property.
// The version hash for the destination entity for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityVersionHash
func (e_ EntityMapping) SetDestinationEntityVersionHash(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityVersionHash:"), value)
}

// The class name of the migration policy for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/entityMigrationPolicyClassName
func (e_ EntityMapping) EntityMigrationPolicyClassName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("entityMigrationPolicyClassName"))
	return rv
}


// SetEntityMigrationPolicyClassName sets the value of the entityMigrationPolicyClassName property.
// The class name of the migration policy for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/entityMigrationPolicyClassName
func (e_ EntityMapping) SetEntityMigrationPolicyClassName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEntityMigrationPolicyClassName:"), objc.String(value))
}

// The mapping type for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/mappingType
func (e_ EntityMapping) MappingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("mappingType"))
	return rv
}


// SetMappingType sets the value of the mappingType property.
// The mapping type for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/mappingType
func (e_ EntityMapping) SetMappingType(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMappingType:"), value)
}

// The name of the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/name
func (e_ EntityMapping) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/name
func (e_ EntityMapping) SetName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), objc.String(value))
}

// The array of relationship mappings for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/relationshipMappings
func (e_ EntityMapping) RelationshipMappings() []PropertyMapping {
	rv := objc.Send[[]PropertyMapping](e_.ID, objc.Sel("relationshipMappings"))
	return rv
}


// SetRelationshipMappings sets the value of the relationshipMappings property.
// The array of relationship mappings for the entity mapping.

//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityName
func (e_ EntityMapping) SourceEntityName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("sourceEntityName"))
	return rv
}


// SetSourceEntityName sets the value of the sourceEntityName property.
// The source entity name for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityName
func (e_ EntityMapping) SetSourceEntityName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityName:"), objc.String(value))
}

// The version hash of the source entity for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityVersionHash
func (e_ EntityMapping) SourceEntityVersionHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("sourceEntityVersionHash"))
	return rv
}


// SetSourceEntityVersionHash sets the value of the sourceEntityVersionHash property.
// The version hash of the source entity for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityVersionHash
func (e_ EntityMapping) SetSourceEntityVersionHash(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityVersionHash:"), value)
}

// The source expression for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceExpression
func (e_ EntityMapping) SourceExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("sourceExpression"))
	return rv
}


// SetSourceExpression sets the value of the sourceExpression property.
// The source expression for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceExpression
func (e_ EntityMapping) SetSourceExpression(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceExpression:"), value)
}

// The user info dictionary for the entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/userInfo
func (e_ EntityMapping) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// The user info dictionary for the entity mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping/userInfo
func (e_ EntityMapping) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}



