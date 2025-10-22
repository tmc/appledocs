// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationManager] class.
var (
	MigrationManagerClass     _MigrationManagerClass
	MigrationManagerClassOnce sync.Once
)

func getMigrationManagerClass() _MigrationManagerClass {
	MigrationManagerClassOnce.Do(func() {
		MigrationManagerClass = _MigrationManagerClass{objc.GetClass("NSMigrationManager")}
	})
	return MigrationManagerClass
}

type _MigrationManagerClass struct {
	class objc.Class
}

// An interface definition for the [MigrationManager] class.
type IMigrationManager interface {
	objectivec.IObject
	AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance IManagedObject, destinationInstance IManagedObject, entityMapping IEntityMapping)
	CancelMigrationWithError(error_ foundation.IError)
	DestinationEntityForEntityMapping(mEntity IEntityMapping) EntityDescription
	DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []ManagedObject) []ManagedObject
	MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError(sourceURL foundation.IURL, sStoreType string, sOptions objectivec.IObject, mappings IMappingModel, dURL foundation.IURL, dStoreType string, dOptions objectivec.IObject, error_ unsafe.Pointer) bool
	Reset()
	SourceEntityForEntityMapping(mEntity IEntityMapping) EntityDescription
	SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []ManagedObject) []ManagedObject
	CurrentEntityMapping() NSEntityMapping
	DestinationContext() NSManagedObjectContext
	DestinationModel() NSManagedObjectModel
	MappingModel() NSMappingModel
	MigrationProgress() float32
	SourceContext() NSManagedObjectContext
	SourceModel() NSManagedObjectModel
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
	UsesStoreSpecificMigrationManager() bool
	SetUsesStoreSpecificMigrationManager(value bool)
}

// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.


// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager

type MigrationManager struct {
	objectivec.Object
}

// MigrationManagerFrom constructs a [MigrationManager] from an unsafe.Pointer.
//
// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.
func MigrationManagerFrom(ptr unsafe.Pointer) MigrationManager {
	return MigrationManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MigrationManagerClass) Alloc() MigrationManager {
	rv := objc.Send[MigrationManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MigrationManagerClass) New() MigrationManager {
	rv := objc.Send[MigrationManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MigrationManager) Init() MigrationManager {
	rv := objc.Send[MigrationManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MigrationManager) Autorelease() MigrationManager {
	rv := objc.Send[MigrationManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMigrationManager creates a new MigrationManager instance.
func NewMigrationManager() MigrationManager {
	return getMigrationManagerClass().New()
}




// Initializes a migration manager instance with given source and destination models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/init(sourceModel:destinationModel:)

func NewMigrationManagerWithSourceModelDestinationModel(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MigrationManager {
	instance := getMigrationManagerClass().Alloc()
	rv := objc.Send[MigrationManager](instance.ID, objc.Sel("initWithSourceModel:destinationModel:"), sourceModel, destinationModel)
	rv.Autorelease()
	return rv
}




// Associates a given source managed object instance with an array of destination instances for a given property mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/associate(sourceInstance:withDestinationInstance:for:)

func (m_ MigrationManager) AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance IManagedObject, destinationInstance IManagedObject, entityMapping IEntityMapping) {
	objc.Send[objc.ID](m_.ID, objc.Sel("associateSourceInstance:withDestinationInstance:forEntityMapping:"), sourceInstance, destinationInstance, entityMapping)
}



// Cancels the migration with a given error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/cancelMigrationWithError(_:)

func (m_ MigrationManager) CancelMigrationWithError(error_ foundation.IError) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelMigrationWithError:"), error_)
}



// Returns the entity description for the destination entity of a given entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationEntity(for:)

func (m_ MigrationManager) DestinationEntityForEntityMapping(mEntity IEntityMapping) EntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("destinationEntityForEntityMapping:"), mEntity)
	return rv
}



// Returns the managed object instances created in the destination store for the named entity mapping for the given array of source instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationInstances(forEntityMappingName:sourceInstances:)

func (m_ MigrationManager) DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []ManagedObject) []ManagedObject {
	rv := objc.Send[[]ManagedObject](m_.ID, objc.Sel("destinationInstancesForEntityMappingNamed:sourceInstances:"), objc.String(mappingName), sourceInstances)
	return rv
}



// Migrates the store at a given source URL to the store at a given destination URL, performing all of the mappings specified in a given mapping model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/migrateStore(from:sourceType:options:with:toDestinationURL:destinationType:destinationOptions:)

func (m_ MigrationManager) MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError(sourceURL foundation.IURL, sStoreType string, sOptions objectivec.IObject, mappings IMappingModel, dURL foundation.IURL, dStoreType string, dOptions objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("migrateStoreFromURL:type:options:withMappingModel:toDestinationURL:destinationType:destinationOptions:error:"), sourceURL, objc.String(sStoreType), sOptions, mappings, dURL, objc.String(dStoreType), dOptions, error_)
	return rv
}



// Resets the association tables for the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/reset()

func (m_ MigrationManager) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}



// Returns the entity description for the source entity of a given entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceEntity(for:)

func (m_ MigrationManager) SourceEntityForEntityMapping(mEntity IEntityMapping) EntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("sourceEntityForEntityMapping:"), mEntity)
	return rv
}



// Returns the managed object instances in the source store used to create the given destination instances for the passed in property mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceInstances(forEntityMappingName:destinationInstances:)

func (m_ MigrationManager) SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []ManagedObject) []ManagedObject {
	rv := objc.Send[[]ManagedObject](m_.ID, objc.Sel("sourceInstancesForEntityMappingNamed:destinationInstances:"), objc.String(mappingName), destinationInstances)
	return rv
}


// The entity mapping currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/currentEntityMapping

func (m_ MigrationManager) CurrentEntityMapping() NSEntityMapping {
	rv := objc.Send[NSEntityMapping](m_.ID, objc.Sel("currentEntityMapping"))
	return rv
}


// The managed object context the migration manager uses for writing the destination persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationContext

func (m_ MigrationManager) DestinationContext() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m_.ID, objc.Sel("destinationContext"))
	return rv
}


// The destination model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationModel

func (m_ MigrationManager) DestinationModel() NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](m_.ID, objc.Sel("destinationModel"))
	return rv
}


// The mapping model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/mappingModel

func (m_ MigrationManager) MappingModel() NSMappingModel {
	rv := objc.Send[NSMappingModel](m_.ID, objc.Sel("mappingModel"))
	return rv
}


// A number between and that indicates the proportion of completeness of the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/migrationProgress

func (m_ MigrationManager) MigrationProgress() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("migrationProgress"))
	return rv
}


// The managed object context the migration manager uses for reading the source persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceContext

func (m_ MigrationManager) SourceContext() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m_.ID, objc.Sel("sourceContext"))
	return rv
}


// The source model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceModel

func (m_ MigrationManager) SourceModel() NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](m_.ID, objc.Sel("sourceModel"))
	return rv
}


// The user info for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/userInfo

func (m_ MigrationManager) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/userInfo

func (m_ MigrationManager) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInfo:"), value)
}


// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/usesStoreSpecificMigrationManager

func (m_ MigrationManager) UsesStoreSpecificMigrationManager() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesStoreSpecificMigrationManager"))
	return rv
}


// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/usesStoreSpecificMigrationManager

func (m_ MigrationManager) SetUsesStoreSpecificMigrationManager(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesStoreSpecificMigrationManager:"), value)
}


