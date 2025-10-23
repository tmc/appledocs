// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityMigrationPolicy] class.
var (
	EntityMigrationPolicyClass     _EntityMigrationPolicyClass
	EntityMigrationPolicyClassOnce sync.Once
)

func getEntityMigrationPolicyClass() _EntityMigrationPolicyClass {
	EntityMigrationPolicyClassOnce.Do(func() {
		EntityMigrationPolicyClass = _EntityMigrationPolicyClass{objc.GetClass("NSEntityMigrationPolicy")}
	})
	return EntityMigrationPolicyClass
}

type _EntityMigrationPolicyClass struct {
	class objc.Class
}

// An interface definition for the [EntityMigrationPolicy] class.
type IEntityMigrationPolicy interface {
	objectivec.IObject
	EntityMigrationPolicyClassName() string
	SetEntityMigrationPolicyClassName(value string)
	NSMigrationDestinationObjectKey() string
	NSMigrationEntityMappingKey() string
	NSMigrationEntityPolicyKey() string
	NSMigrationManagerKey() string
	NSMigrationPropertyMappingKey() string
	NSMigrationSourceObjectKey() string
	CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error_ unsafe.Pointer) bool
}

// A policy instance that customizes the migration process for an entity mapping.
//
// You set the policy for an entity mapping by passing the name of the migration policy class as the argument to . Typically, you specify the name in the Xcode mapping model editor.


// A policy instance that customizes the migration process for an entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy
type EntityMigrationPolicy struct {
	objectivec.Object
}

// EntityMigrationPolicyFrom constructs a [EntityMigrationPolicy] from an unsafe.Pointer.
//
// A policy instance that customizes the migration process for an entity mapping.
func EntityMigrationPolicyFrom(ptr unsafe.Pointer) EntityMigrationPolicy {
	return EntityMigrationPolicy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EntityMigrationPolicyClass) Alloc() EntityMigrationPolicy {
	rv := objc.Send[EntityMigrationPolicy](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityMigrationPolicyClass) New() EntityMigrationPolicy {
	rv := objc.Send[EntityMigrationPolicy](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EntityMigrationPolicy) Init() EntityMigrationPolicy {
	rv := objc.Send[EntityMigrationPolicy](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EntityMigrationPolicy) Autorelease() EntityMigrationPolicy {
	rv := objc.Send[EntityMigrationPolicy](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntityMigrationPolicy creates a new EntityMigrationPolicy instance.
func NewEntityMigrationPolicy() EntityMigrationPolicy {
	return getEntityMigrationPolicyClass().New()
}



// Creates the destination instance(s) for a given source instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createDestinationInstances(forSource:in:manager:)
func (e_ EntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, error_)
	return rv
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/entitymigrationpolicyclassname
func (e_ EntityMigrationPolicy) EntityMigrationPolicyClassName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("entityMigrationPolicyClassName"))
	return rv
}


// The class name of the migration policy for the entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/entitymigrationpolicyclassname
func (e_ EntityMigrationPolicy) SetEntityMigrationPolicyClassName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEntityMigrationPolicyClassName:"), objc.String(value))
}


// Key for the destination object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationdestinationobjectkey
func (e_ EntityMigrationPolicy) NSMigrationDestinationObjectKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationDestinationObjectKey"))
	return rv
}


// Key for the entity mapping object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationentitymappingkey
func (e_ EntityMigrationPolicy) NSMigrationEntityMappingKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationEntityMappingKey"))
	return rv
}


// Key for the entity migration policy object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationentitypolicykey
func (e_ EntityMigrationPolicy) NSMigrationEntityPolicyKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationEntityPolicyKey"))
	return rv
}


// Key for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanagerkey
func (e_ EntityMigrationPolicy) NSMigrationManagerKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationManagerKey"))
	return rv
}


// Key for the property mapping object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationpropertymappingkey
func (e_ EntityMigrationPolicy) NSMigrationPropertyMappingKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationPropertyMappingKey"))
	return rv
}


// Key for the source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationsourceobjectkey
func (e_ EntityMigrationPolicy) NSMigrationSourceObjectKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMigrationSourceObjectKey"))
	return rv
}



