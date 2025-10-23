// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

// Enum types and constants
// NSAttributeType - The types of attribute that Core Data supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType
type NSAttributeType uint

const (
	// NSDecimalAttributeType - An attribute that stores a decimal value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeType/decimalAttributeType
	NSDecimalAttributeType NSAttributeType = 400
)

// NSBatchDeleteRequestResultType - The types of result a batch delete request can provide when it executes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType
type NSBatchDeleteRequestResultType uint

const (
	NSBatchDeleteResultTypeStatusOnly NSBatchDeleteRequestResultType = 0
	NSBatchDeleteResultTypeObjectIDs NSBatchDeleteRequestResultType = 1
	NSBatchDeleteResultTypeCount NSBatchDeleteRequestResultType = 2
)

// NSBatchInsertRequestResultType - Result types for a batch-insertion request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType
type NSBatchInsertRequestResultType uint

const (
	NSBatchInsertRequestResultTypeStatusOnly NSBatchInsertRequestResultType = 0
	NSBatchInsertRequestResultTypeObjectIDs NSBatchInsertRequestResultType = 1
	NSBatchInsertRequestResultTypeCount NSBatchInsertRequestResultType = 2
)

// NSEntityMappingType - The types for mapping an entity between a source model and a destination model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType
type NSEntityMappingType uint

const (
	// NSTransformEntityMappingType - Specifies that entity exists in source and destination and is mapped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMappingType/transformEntityMappingType
	NSTransformEntityMappingType NSEntityMappingType = 5
)

// NSFetchRequestResultType - Constants that specify the possible result types a fetch request can return.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType
type NSFetchRequestResultType uint

const (
	// NSCountResultType - The request returns the count of the objects that match the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/countResultType
	NSCountResultType NSFetchRequestResultType = 3
	// NSDictionaryResultType - The request returns dictionaries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/dictionaryResultType
	NSDictionaryResultType NSFetchRequestResultType = 2
	// NSManagedObjectIDResultType - The request returns managed object IDs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/managedObjectIDResultType
	NSManagedObjectIDResultType NSFetchRequestResultType = 1
	// NSManagedObjectResultType - The request returns managed objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType/managedObjectResultType
	NSManagedObjectResultType NSFetchRequestResultType = 0
)

// NSFetchedResultsChangeType - Constants that specify the possible types of changes that are reported.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType
type NSFetchedResultsChangeType uint

const (
	// NSFetchedResultsChangeDelete - Specifies that an object was deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/delete
	NSFetchedResultsChangeDelete NSFetchedResultsChangeType = 2
	// NSFetchedResultsChangeInsert - Specifies that an object was inserted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/insert
	NSFetchedResultsChangeInsert NSFetchedResultsChangeType = 1
	// NSFetchedResultsChangeMove - Specifies that an object was moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/move
	NSFetchedResultsChangeMove NSFetchedResultsChangeType = 3
	// NSFetchedResultsChangeUpdate - Specifies that an object was changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/update
	NSFetchedResultsChangeUpdate NSFetchedResultsChangeType = 4
)

// NSPersistentCloudKitContainerEventType - The type of event in a persistent CloudKit container, either setup, import, or export.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType
type NSPersistentCloudKitContainerEventType uint

const (
	// NSPersistentCloudKitContainerEventTypeExport - An event the persistent CloudKit container generates when exporting managed objects from a store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType/export
	NSPersistentCloudKitContainerEventTypeExport NSPersistentCloudKitContainerEventType = 2
	// NSPersistentCloudKitContainerEventTypeImport - An event the persistent CloudKit container generates when importing records into a store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType/import
	NSPersistentCloudKitContainerEventTypeImport NSPersistentCloudKitContainerEventType = 1
)

// NSPersistentCloudKitContainerEventResultType - The types of results from a persistent CloudKit container event fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/ResultType-swift.enum
type NSPersistentCloudKitContainerEventResultType uint

const (
	NSPersistentCloudKitContainerEventResultTypeEvents NSPersistentCloudKitContainerEventResultType = 0
	NSPersistentCloudKitContainerEventResultTypeCountEvents NSPersistentCloudKitContainerEventResultType = 1
)

// NSPersistentCloudKitContainerSchemaInitializationOptions - Options that control the behavior when promoting the container’s schema to CloudKit.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions
type NSPersistentCloudKitContainerSchemaInitializationOptions uint

const (
	// NSPersistentCloudKitContainerSchemaInitializationOptionsNone - Indicates there are no specified schema options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/NSPersistentCloudKitContainerSchemaInitializationOptionsNone
	NSPersistentCloudKitContainerSchemaInitializationOptionsNone NSPersistentCloudKitContainerSchemaInitializationOptions = 0
	// NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun - A flag that indicates the container validates the model and generates the records, but doesn’t upload them to CloudKit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/dryRun
	NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun NSPersistentCloudKitContainerSchemaInitializationOptions = 2
	// NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema - Prints the generated records to the console.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions/printSchema
	NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema NSPersistentCloudKitContainerSchemaInitializationOptions = 4
)

// NSPersistentHistoryResultType - The types of results from a persistent history change request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResultType
type NSPersistentHistoryResultType uint

const (
	NSPersistentHistoryResultTypeStatusOnly NSPersistentHistoryResultType = 0
	NSPersistentHistoryResultTypeObjectIDs NSPersistentHistoryResultType = 1
	NSPersistentHistoryResultTypeCount NSPersistentHistoryResultType = 2
	NSPersistentHistoryResultTypeTransactionsOnly NSPersistentHistoryResultType = 3
	NSPersistentHistoryResultTypeChangesOnly NSPersistentHistoryResultType = 4
	NSPersistentHistoryResultTypeTransactionsAndChanges NSPersistentHistoryResultType = 5
)

// NSPersistentStoreRequestType - Constants that specify the types of fetch requests.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType
type NSPersistentStoreRequestType uint

const (
	NSFetchRequestType NSPersistentStoreRequestType = 1
	NSSaveRequestType NSPersistentStoreRequestType = 2
	NSBatchInsertRequestType NSPersistentStoreRequestType = 3
	NSBatchUpdateRequestType NSPersistentStoreRequestType = 4
	NSBatchDeleteRequestType NSPersistentStoreRequestType = 5
)


