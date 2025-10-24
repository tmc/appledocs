// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomMigrationStage] class.
var (
	CustomMigrationStageClass     _CustomMigrationStageClass
	CustomMigrationStageClassOnce sync.Once
)

func getCustomMigrationStageClass() _CustomMigrationStageClass {
	CustomMigrationStageClassOnce.Do(func() {
		CustomMigrationStageClass = _CustomMigrationStageClass{objc.GetClass("NSCustomMigrationStage")}
	})
	return CustomMigrationStageClass
}

type _CustomMigrationStageClass struct {
	class objc.Class
}

// An interface definition for the [CustomMigrationStage] class.
type ICustomMigrationStage interface {
	IMigrationStage
	// properties:
	CurrentModel() IManagedObjectModelReference
	DidMigrateHandler() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	SetDidMigrateHandler(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer)
	NextModel() IManagedObjectModelReference
	WillMigrateHandler() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	SetWillMigrateHandler(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer)
	Container() IPersistentContainer
	SetContainer(value IPersistentContainer)
	// methods:
}

// An object that enables you to participate in the migration between two versions of the same model.
//
// Use when you have two versions of a model that Core Data can’t automatically migrate. Custom migration stages enable you to participate in the migration process by assigning handlers that the stage invokes before and after it runs. The handlers provide an opportunity to prepare the persistent store’s data for the upcoming changes before the stage runs, and perform any cleanup tasks afterward. For example, to support a migration that changes an optional attribute to be nonoptional, you might assign a handler to the stage’s property that sets any instances of that attribute to a default value, thereby ensuring the migration succeeds. To access the store you’re migrating, use the property of the migration manager that Core Data provides to every handler.

// An object that enables you to participate in the migration between two versions of the same model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage
type CustomMigrationStage struct {
	MigrationStage
}

// CustomMigrationStageFrom constructs a [CustomMigrationStage] from an unsafe.Pointer.
//
// An object that enables you to participate in the migration between two versions of the same model.
func CustomMigrationStageFrom(ptr unsafe.Pointer) CustomMigrationStage {
	return CustomMigrationStage{
		MigrationStage: MigrationStageFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomMigrationStageClass) Alloc() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomMigrationStageClass) New() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomMigrationStage) Init() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomMigrationStage) Autorelease() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomMigrationStage creates a new CustomMigrationStage instance.
func NewCustomMigrationStage() CustomMigrationStage {
	return getCustomMigrationStageClass().New()
}

// Creates a custom migration stage with the specified source and destination model references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/initWithCurrentModelReference:nextModelReference:
func NewCustomMigrationStageWithCurrentModelReferenceNextModelReference(currentModel IManagedObjectModelReference, nextModel IManagedObjectModelReference) CustomMigrationStage {
	instance := getCustomMigrationStageClass().Alloc()
	rv := objc.Send[CustomMigrationStage](instance.ID, objc.Sel("initWithCurrentModelReference:nextModelReference:"), currentModel, nextModel)
	rv.Autorelease()
	return rv
}

// The reference that represents the migration’s source model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/currentModel
func (c_ CustomMigrationStage) CurrentModel() IManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](c_.ID, objc.Sel("currentModel"))
	return rv
}

// The handler to execute after the stage runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/didMigrateHandler-36uhx
func (c_ CustomMigrationStage) DidMigrateHandler() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer](c_.ID, objc.Sel("didMigrateHandler"))
	return rv
}

// The handler to execute after the stage runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/didMigrateHandler-36uhx
func (c_ CustomMigrationStage) SetDidMigrateHandler(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDidMigrateHandler:"), value)
}

// The reference that represents the migration’s destination model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/nextModel
func (c_ CustomMigrationStage) NextModel() IManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](c_.ID, objc.Sel("nextModel"))
	return rv
}

// The handler to execute before the stage runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/willMigrateHandler-72p73
func (c_ CustomMigrationStage) WillMigrateHandler() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer](c_.ID, objc.Sel("willMigrateHandler"))
	return rv
}

// The handler to execute before the stage runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/willMigrateHandler-72p73
func (c_ CustomMigrationStage) SetWillMigrateHandler(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWillMigrateHandler:"), value)
}

// The container that provides access to the migrating persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstagedmigrationmanager/container
func (c_ CustomMigrationStage) Container() IPersistentContainer {
	rv := objc.Send[PersistentContainer](c_.ID, objc.Sel("container"))
	return rv
}

// The container that provides access to the migrating persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstagedmigrationmanager/container
func (c_ CustomMigrationStage) SetContainer(value IPersistentContainer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}
