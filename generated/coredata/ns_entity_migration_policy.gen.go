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
	BeginEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	EndEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	EndInstanceCreationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	EndRelationshipCreationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
	PerformCustomValidationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool
}

// A policy instance that customizes the migration process for an entity mapping.
//
// You set the policy for an entity mapping by passing the name of the migration policy class as the argument to . Typically, you specify the name in the Xcode mapping model editor.
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


// Sets up state information before the start of a given entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/begin(_:with:)
func (e_ EntityMigrationPolicy) BeginEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("beginEntityMapping:manager:error:"), mapping, manager, error_)
	return rv
}

// Creates the destination instance(s) for a given source instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createDestinationInstances(forSource:in:manager:)
func (e_ EntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, error_)
	return rv
}

// Constructs the relationships between the newly-created destination instances.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createRelationships(forDestination:in:manager:)
func (e_ EntityMigrationPolicy) CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("createRelationshipsForDestinationInstance:entityMapping:manager:error:"), dInstance, mapping, manager, error_)
	return rv
}

// Performs cleanup at the end of the migration, from any phase of the mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/end(_:manager:)
func (e_ EntityMigrationPolicy) EndEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("endEntityMapping:manager:error:"), mapping, manager, error_)
	return rv
}

// Indicates the end of the instance creation stage for the specified entity mapping, and the precursor to the next migration stage.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/endInstanceCreation(forMapping:manager:)
func (e_ EntityMigrationPolicy) EndInstanceCreationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("endInstanceCreationForEntityMapping:manager:error:"), mapping, manager, error_)
	return rv
}

// Indicates the end of the relationship creation stage for the specified entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/endRelationshipCreation(forMapping:manager:)
func (e_ EntityMigrationPolicy) EndRelationshipCreationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("endRelationshipCreationForEntityMapping:manager:error:"), mapping, manager, error_)
	return rv
}

// Provides the option to perform custom validation on migrated objects during the validation stage of the entity migration policy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/performCustomValidation(forMapping:manager:)
func (e_ EntityMigrationPolicy) PerformCustomValidationForEntityMappingManagerError(mapping unsafe.Pointer, manager unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("performCustomValidationForEntityMapping:manager:error:"), mapping, manager, error_)
	return rv
}



