// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityMigrationPolicy] class.
var entityMigrationPolicyClass = _EntityMigrationPolicyClass{objc.GetClass("NSEntityMigrationPolicy")}

type _EntityMigrationPolicyClass struct {
	class objc.Class
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

// Creates the destination instance(s) for a given source instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createDestinationInstances(forSource:in:manager:)
func (e_ EntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance unsafe.Pointer, mapping unsafe.Pointer, manager unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, error)
	return rv
}


