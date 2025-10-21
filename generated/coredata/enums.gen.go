// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

// Enum types and constants
// NSAttributeType - The types of attribute that Core Data supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType
type AttributeType uint

const (
// URIAttributeType - An attribute that stores a uniform resource identifier.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/URIAttributeType
URIAttributeType AttributeType = 0
// UUIDAttributeType - An attribute that stores a universally unique identifier.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/UUIDAttributeType
UUIDAttributeType AttributeType = 0
// BinaryDataAttributeType - An attribute that stores binary data.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/binaryDataAttributeType
BinaryDataAttributeType AttributeType = 0
// BooleanAttributeType - An attribute that stores a Boolean value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/booleanAttributeType
BooleanAttributeType AttributeType = 0
// CompositeAttributeType - An attribute that derives its value by composing other attributes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/compositeAttributeType
CompositeAttributeType AttributeType = 0
// DateAttributeType - An attribute that stores a date.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/dateAttributeType
DateAttributeType AttributeType = 0
// DecimalAttributeType - An attribute that stores a decimal value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/decimalAttributeType
DecimalAttributeType AttributeType = 0
// DoubleAttributeType - An attribute that stores a double value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/doubleAttributeType
DoubleAttributeType AttributeType = 0
// FloatAttributeType - An attribute that stores a float value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/floatAttributeType
FloatAttributeType AttributeType = 0
// Integer16AttributeType - An attribute that stores a 16-bit signed integer value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/integer16AttributeType
Integer16AttributeType AttributeType = 0
// Integer32AttributeType - An attribute that stores a 32-bit signed integer value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/integer32AttributeType
Integer32AttributeType AttributeType = 0
// Integer64AttributeType - An attribute that stores a 64-bit signed integer value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/integer64AttributeType
Integer64AttributeType AttributeType = 0
// ObjectIDAttributeType - An attribute that stores a managed object’s ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/objectIDAttributeType
ObjectIDAttributeType AttributeType = 0
// StringAttributeType - An attribute that stores a string.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/stringAttributeType
StringAttributeType AttributeType = 0
// TransformableAttributeType - An attribute that uses a value transformer to derive its value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/transformableAttributeType
TransformableAttributeType AttributeType = 0
// UndefinedAttributeType - An attribute that doesn’t have an explicit type.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/undefinedAttributeType
UndefinedAttributeType AttributeType = 0
)

// NSBatchDeleteRequestResultType - The types of result a batch delete request can provide when it executes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType
type BatchDeleteRequestResultType uint

const (
// BatchDeleteResultTypeCount - Returns the number of managed objects the request deletes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType/resultTypeCount
BatchDeleteResultTypeCount BatchDeleteRequestResultType = 0
// BatchDeleteResultTypeObjectIDs - Returns an array of the deleted managed objects’ identifiers.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType/resultTypeObjectIDs
BatchDeleteResultTypeObjectIDs BatchDeleteRequestResultType = 0
// BatchDeleteResultTypeStatusOnly - Returns a Boolean value that indicates if the request succeeds.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType/resultTypeStatusOnly
BatchDeleteResultTypeStatusOnly BatchDeleteRequestResultType = 0
)

// NSBatchInsertRequestResultType - Result types for a batch-insertion request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType
type BatchInsertRequestResultType uint

const (
// BatchInsertRequestResultTypeStatusOnly - A value that indicates that the return type is a Boolean value representing whether the batch-insertion request succeeded.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType/statusOnly
BatchInsertRequestResultTypeStatusOnly BatchInsertRequestResultType = 0
)

// NSDeleteRule - Constants that determine what happens when you delete a relationship’s owning managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDeleteRule
type DeleteRule uint

const (
// CascadeDeleteRule - A rule that deletes the referenced managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDeleteRule/cascadeDeleteRule
CascadeDeleteRule DeleteRule = 0
// DenyDeleteRule - A rule that prevents the deletion of the owning managed object if the relationship has references to other objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDeleteRule/denyDeleteRule
DenyDeleteRule DeleteRule = 0
// NoActionDeleteRule - A rule that prevents modification of the referenced managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDeleteRule/noActionDeleteRule
NoActionDeleteRule DeleteRule = 0
// NullifyDeleteRule - A rule that nullifies the inverse relationship of the referenced managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDeleteRule/nullifyDeleteRule
NullifyDeleteRule DeleteRule = 0
)

// NSEntityMappingType - The types for mapping an entity between a source model and a destination model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType
type EntityMappingType uint

const (
// AddEntityMappingType - Specifies that this is a new entity in the destination model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/addEntityMappingType
AddEntityMappingType EntityMappingType = 0
// CopyEntityMappingType - Specifies that source instances are migrated as-is.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/copyEntityMappingType
CopyEntityMappingType EntityMappingType = 0
// CustomEntityMappingType - Specifies a custom mapping.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/customEntityMappingType
CustomEntityMappingType EntityMappingType = 0
// RemoveEntityMappingType - Specifies that this entity is not present in the destination model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/removeEntityMappingType
RemoveEntityMappingType EntityMappingType = 0
// TransformEntityMappingType - Specifies that entity exists in source and destination and is mapped.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/transformEntityMappingType
TransformEntityMappingType EntityMappingType = 0
// UndefinedEntityMappingType - Specifies that the developer handles destination instance creation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/undefinedEntityMappingType
UndefinedEntityMappingType EntityMappingType = 0
)

// NSFetchRequestResultType - Constants that specify the possible result types a fetch request can return.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType
type FetchRequestResultType uint

const (
// CountResultType - The request returns the count of the objects that match the request.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/countResultType
CountResultType FetchRequestResultType = 0
// DictionaryResultType - The request returns dictionaries.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/dictionaryResultType
DictionaryResultType FetchRequestResultType = 0
// ManagedObjectIDResultType - The request returns managed object IDs.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/managedObjectIDResultType
ManagedObjectIDResultType FetchRequestResultType = 0
// ManagedObjectResultType - The request returns managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/managedObjectResultType
ManagedObjectResultType FetchRequestResultType = 0
)

// NSFetchedResultsChangeType - Constants that specify the possible types of changes that are reported.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType
type FetchedResultsChangeType uint

const (
// FetchedResultsChangeDelete - Specifies that an object was deleted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/delete
FetchedResultsChangeDelete FetchedResultsChangeType = 0
// FetchedResultsChangeInsert - Specifies that an object was inserted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/insert
FetchedResultsChangeInsert FetchedResultsChangeType = 0
// FetchedResultsChangeMove - Specifies that an object was moved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/move
FetchedResultsChangeMove FetchedResultsChangeType = 0
// FetchedResultsChangeUpdate - Specifies that an object was changed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/update
FetchedResultsChangeUpdate FetchedResultsChangeType = 0
)

// NSManagedObjectContextConcurrencyType - The concurrency types you can use with a managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType
type ManagedObjectContextConcurrencyType uint

const (
// ConfinementConcurrencyType - Specifies that the context will use the thread confinement pattern.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType/confinementConcurrencyType
ConfinementConcurrencyType ManagedObjectContextConcurrencyType = 0
// MainQueueConcurrencyType - Specifies that the context will be associated with the main queue.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType/mainQueueConcurrencyType
MainQueueConcurrencyType ManagedObjectContextConcurrencyType = 0
// PrivateQueueConcurrencyType - Specifies that the context will be associated with a private dispatch queue.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType/privateQueueConcurrencyType
PrivateQueueConcurrencyType ManagedObjectContextConcurrencyType = 0
)

// NSMergePolicyType - Constants that define merge policy types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType
type MergePolicyType uint

const (
// ErrorMergePolicyType - The default merge policy for all managed object contexts.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType/errorMergePolicyType
ErrorMergePolicyType MergePolicyType = 0
// MergeByPropertyObjectTrumpMergePolicyType - A property-based merge policy that applies in-memory changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType/mergeByPropertyObjectTrumpMergePolicyType
MergeByPropertyObjectTrumpMergePolicyType MergePolicyType = 0
// MergeByPropertyStoreTrumpMergePolicyType - A property-based merge policy that applies external changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType/mergeByPropertyStoreTrumpMergePolicyType
MergeByPropertyStoreTrumpMergePolicyType MergePolicyType = 0
// OverwriteMergePolicyType - A merge policy type that overwrites the entire stored object.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType/overwriteMergePolicyType
OverwriteMergePolicyType MergePolicyType = 0
// RollbackMergePolicyType - A merge policy that discards unsaved changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicyType/rollbackMergePolicyType
RollbackMergePolicyType MergePolicyType = 0
)

// NSPersistentCloudKitContainerEventType - The type of event in a persistent CloudKit container, either setup, import, or export.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType
type PersistentCloudKitContainerEventType uint

const (
// PersistentCloudKitContainerEventTypeExport - An event the persistent CloudKit container generates when exporting managed objects from a store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType/export
PersistentCloudKitContainerEventTypeExport PersistentCloudKitContainerEventType = 0
// PersistentCloudKitContainerEventTypeImport - An event the persistent CloudKit container generates when importing records into a store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType/import
PersistentCloudKitContainerEventTypeImport PersistentCloudKitContainerEventType = 0
)

// NSPersistentCloudKitContainerEventResultType - The types of results from a persistent CloudKit container event fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/ResultType-swift.enum
type PersistentCloudKitContainerEventResultType uint

// NSPersistentCloudKitContainerSchemaInitializationOptions - Options that control the behavior when promoting the container’s schema to CloudKit.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions
type PersistentCloudKitContainerSchemaInitializationOptions uint

const (
// PersistentCloudKitContainerSchemaInitializationOptionsNone - Indicates there are no specified schema options.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/NSPersistentCloudKitContainerSchemaInitializationOptionsNone
PersistentCloudKitContainerSchemaInitializationOptionsNone PersistentCloudKitContainerSchemaInitializationOptions = 0
// PersistentCloudKitContainerSchemaInitializationOptionsDryRun - A flag that indicates the container validates the model and generates the records, but doesn’t upload them to CloudKit.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/dryRun
PersistentCloudKitContainerSchemaInitializationOptionsDryRun PersistentCloudKitContainerSchemaInitializationOptions = 0
// PersistentCloudKitContainerSchemaInitializationOptionsPrintSchema - Prints the generated records to the console.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/printSchema
PersistentCloudKitContainerSchemaInitializationOptionsPrintSchema PersistentCloudKitContainerSchemaInitializationOptions = 0
)

// NSPersistentHistoryChangeType - The types of changes to managed objects reflected in persistent history.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType
type PersistentHistoryChangeType uint

const (
// PersistentHistoryChangeTypeDelete - The deletion of a managed object from the persistent store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType/delete
PersistentHistoryChangeTypeDelete PersistentHistoryChangeType = 0
// PersistentHistoryChangeTypeInsert - The insertion of a managed object into the persistent store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType/insert
PersistentHistoryChangeTypeInsert PersistentHistoryChangeType = 0
// PersistentHistoryChangeTypeUpdate - An update to a managed object’s properties in the persistent store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType/update
PersistentHistoryChangeTypeUpdate PersistentHistoryChangeType = 0
)

// NSPersistentHistoryResultType - The types of results from a persistent history change request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResultType
type PersistentHistoryResultType uint

// NSPersistentStoreRequestType - Constants that specify the types of fetch requests.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType
type PersistentStoreRequestType uint

const (
// BatchDeleteRequestType - A request that deletes data for multiple managed objects from a persistent store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType/batchDeleteRequestType
BatchDeleteRequestType PersistentStoreRequestType = 0
// BatchInsertRequestType - A request that inserts data into a persistent store using a batch of managed objects or dictionaries.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType/batchInsertRequestType
BatchInsertRequestType PersistentStoreRequestType = 0
// BatchUpdateRequestType - A request that updates data for multiple managed objects in a persistent store.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType/batchUpdateRequestType
BatchUpdateRequestType PersistentStoreRequestType = 0
// FetchRequestType - Specifies that the request returns managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType/fetchRequestType
FetchRequestType PersistentStoreRequestType = 0
// SaveRequestType - Specifies that the request saves managed objects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType/saveRequestType
SaveRequestType PersistentStoreRequestType = 0
)


