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
	entityMigrationPolicyClass     _EntityMigrationPolicyClass
	entityMigrationPolicyClassOnce sync.Once
)

func getEntityMigrationPolicyClass() _EntityMigrationPolicyClass {
	entityMigrationPolicyClassOnce.Do(func() {
		entityMigrationPolicyClass = _EntityMigrationPolicyClass{objc.GetClass("NSEntityMigrationPolicy")}
	})
	return entityMigrationPolicyClass
}

type _EntityMigrationPolicyClass struct {
	class objc.Class
}

// An interface definition for the [EntityMigrationPolicy] class.
type IEntityMigrationPolicy interface {
	objectivec.IObject
	CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error unsafe.Pointer) bool
}

// A policy instance that customizes the migration process for an entity mapping. [Full Topic]
//
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


// Creates the destination instance(s) for a given source instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createDestinationInstances(forSource:in:manager:)
func (e_ EntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, error)
	return rv
}


