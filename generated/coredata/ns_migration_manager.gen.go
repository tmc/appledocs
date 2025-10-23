// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	CurrentEntityMapping() IEntityMapping
	SetCurrentEntityMapping(value IEntityMapping)
	DestinationContext() IManagedObjectContext
	SetDestinationContext(value IManagedObjectContext)
	DestinationModel() IManagedObjectModel
	SetDestinationModel(value IManagedObjectModel)
	MappingModel() IMappingModel
	SetMappingModel(value IMappingModel)
	MigrationProgress() float32 /* primitive/slice/pointer. */
	SetMigrationProgress(value float32 /* primitive/slice/pointer. */)
	SourceContext() IManagedObjectContext
	SetSourceContext(value IManagedObjectContext)
	SourceModel() IManagedObjectModel
	SetSourceModel(value IManagedObjectModel)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	UsesStoreSpecificMigrationManager() bool /* primitive/slice/pointer. */
	SetUsesStoreSpecificMigrationManager(value bool /* primitive/slice/pointer. */)
	// methods:
	DestinationEntityForEntityMapping(mEntity IEntityMapping) IEntityDescription
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



// Returns the entity description for the destination entity of a given entity mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationEntity(for:)
func (m_ MigrationManager) DestinationEntityForEntityMapping(mEntity IEntityMapping) IEntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("destinationEntityForEntityMapping:"), mEntity)
	return rv
}


// The entity mapping currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/currententitymapping
func (m_ MigrationManager) CurrentEntityMapping() IEntityMapping {
	rv := objc.Send[EntityMapping](m_.ID, objc.Sel("currentEntityMapping"))
	return rv
}


// The entity mapping currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/currententitymapping
func (m_ MigrationManager) SetCurrentEntityMapping(value IEntityMapping) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentEntityMapping:"), value)
}


// The managed object context the migration manager uses for writing the destination persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/destinationcontext
func (m_ MigrationManager) DestinationContext() IManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](m_.ID, objc.Sel("destinationContext"))
	return rv
}


// The managed object context the migration manager uses for writing the destination persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/destinationcontext
func (m_ MigrationManager) SetDestinationContext(value IManagedObjectContext) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationContext:"), value)
}


// The destination model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/destinationmodel
func (m_ MigrationManager) DestinationModel() IManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("destinationModel"))
	return rv
}


// The destination model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/destinationmodel
func (m_ MigrationManager) SetDestinationModel(value IManagedObjectModel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationModel:"), value)
}


// The mapping model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/mappingmodel
func (m_ MigrationManager) MappingModel() IMappingModel {
	rv := objc.Send[MappingModel](m_.ID, objc.Sel("mappingModel"))
	return rv
}


// The mapping model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/mappingmodel
func (m_ MigrationManager) SetMappingModel(value IMappingModel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMappingModel:"), value)
}


// A number between
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/migrationprogress
func (m_ MigrationManager) MigrationProgress() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](m_.ID, objc.Sel("migrationProgress"))
	return rv
}


// A number between
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/migrationprogress
func (m_ MigrationManager) SetMigrationProgress(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMigrationProgress:"), value)
}


// The managed object context the migration manager uses for reading the source persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/sourcecontext
func (m_ MigrationManager) SourceContext() IManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](m_.ID, objc.Sel("sourceContext"))
	return rv
}


// The managed object context the migration manager uses for reading the source persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/sourcecontext
func (m_ MigrationManager) SetSourceContext(value IManagedObjectContext) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceContext:"), value)
}


// The source model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/sourcemodel
func (m_ MigrationManager) SourceModel() IManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("sourceModel"))
	return rv
}


// The source model for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/sourcemodel
func (m_ MigrationManager) SetSourceModel(value IManagedObjectModel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceModel:"), value)
}


// The user info for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/userinfo
func (m_ MigrationManager) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info for the migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/userinfo
func (m_ MigrationManager) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInfo:"), value)
}


// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/usesstorespecificmigrationmanager
func (m_ MigrationManager) UsesStoreSpecificMigrationManager() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesStoreSpecificMigrationManager"))
	return rv
}


// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/usesstorespecificmigrationmanager
func (m_ MigrationManager) SetUsesStoreSpecificMigrationManager(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesStoreSpecificMigrationManager:"), value)
}



