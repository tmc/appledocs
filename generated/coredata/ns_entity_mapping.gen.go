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
	// properties:
	MappingType() EntityMappingType
	SetMappingType(value EntityMappingType)
	AttributeMappings() IPropertyMapping
	SetAttributeMappings(value IPropertyMapping)
	DestinationEntityName() objc.IObject /* cross-framework: NSString */
	SetDestinationEntityName(value objc.IObject /* cross-framework: NSString */)
	DestinationEntityVersionHash() objc.IObject /* cross-framework: Data */
	SetDestinationEntityVersionHash(value objc.IObject /* cross-framework: Data */)
	EntityMigrationPolicyClassName() objc.IObject /* cross-framework: NSString */
	SetEntityMigrationPolicyClassName(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	RelationshipMappings() IPropertyMapping
	SetRelationshipMappings(value IPropertyMapping)
	SourceEntityName() objc.IObject /* cross-framework: NSString */
	SetSourceEntityName(value objc.IObject /* cross-framework: NSString */)
	SourceEntityVersionHash() objc.IObject /* cross-framework: Data */
	SetSourceEntityVersionHash(value objc.IObject /* cross-framework: Data */)
	SourceExpression() objc.IObject /* cross-framework: Expression */
	SetSourceExpression(value objc.IObject /* cross-framework: Expression */)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
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


// The array of attribute mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/attributemappings
func (e_ EntityMapping) AttributeMappings() IPropertyMapping {
	rv := objc.Send[PropertyMapping](e_.ID, objc.Sel("attributeMappings"))
	return rv
}


// The array of attribute mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/attributemappings
func (e_ EntityMapping) SetAttributeMappings(value IPropertyMapping) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributeMappings:"), value)
}


// The destination entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/destinationentityname
func (e_ EntityMapping) DestinationEntityName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("destinationEntityName"))
	return rv
}


// The destination entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/destinationentityname
func (e_ EntityMapping) SetDestinationEntityName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityName:"), value)
}


// The version hash for the destination entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/destinationentityversionhash
func (e_ EntityMapping) DestinationEntityVersionHash() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](e_.ID, objc.Sel("destinationEntityVersionHash"))
	return rv
}


// The version hash for the destination entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/destinationentityversionhash
func (e_ EntityMapping) SetDestinationEntityVersionHash(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDestinationEntityVersionHash:"), value)
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/entitymigrationpolicyclassname
func (e_ EntityMapping) EntityMigrationPolicyClassName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("entityMigrationPolicyClassName"))
	return rv
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/entitymigrationpolicyclassname
func (e_ EntityMapping) SetEntityMigrationPolicyClassName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEntityMigrationPolicyClassName:"), value)
}


// The name of the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/name
func (e_ EntityMapping) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}


// The name of the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/name
func (e_ EntityMapping) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}


// The array of relationship mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/relationshipmappings
func (e_ EntityMapping) RelationshipMappings() IPropertyMapping {
	rv := objc.Send[PropertyMapping](e_.ID, objc.Sel("relationshipMappings"))
	return rv
}


// The array of relationship mappings for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/relationshipmappings
func (e_ EntityMapping) SetRelationshipMappings(value IPropertyMapping) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelationshipMappings:"), value)
}


// The source entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceentityname
func (e_ EntityMapping) SourceEntityName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("sourceEntityName"))
	return rv
}


// The source entity name for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceentityname
func (e_ EntityMapping) SetSourceEntityName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityName:"), value)
}


// The version hash of the source entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceentityversionhash
func (e_ EntityMapping) SourceEntityVersionHash() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](e_.ID, objc.Sel("sourceEntityVersionHash"))
	return rv
}


// The version hash of the source entity for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceentityversionhash
func (e_ EntityMapping) SetSourceEntityVersionHash(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceEntityVersionHash:"), value)
}


// The source expression for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceexpression
func (e_ EntityMapping) SourceExpression() objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](e_.ID, objc.Sel("sourceExpression"))
	return rv
}


// The source expression for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/sourceexpression
func (e_ EntityMapping) SetSourceExpression(value objc.IObject /* cross-framework: Expression */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSourceExpression:"), value)
}


// The user info dictionary for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/userinfo
func (e_ EntityMapping) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/userinfo
func (e_ EntityMapping) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}



