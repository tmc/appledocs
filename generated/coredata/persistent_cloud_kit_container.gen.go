// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainer] class.
var (
	persistentCloudKitContainerClass     _PersistentCloudKitContainerClass
	persistentCloudKitContainerClassOnce sync.Once
)

func getPersistentCloudKitContainerClass() _PersistentCloudKitContainerClass {
	persistentCloudKitContainerClassOnce.Do(func() {
		persistentCloudKitContainerClass = _PersistentCloudKitContainerClass{objc.GetClass("NSPersistentCloudKitContainer")}
	})
	return persistentCloudKitContainerClass
}

type _PersistentCloudKitContainerClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainer] class.
type IPersistentCloudKitContainer interface {
	IPersistentContainer
	AcceptShareInvitations()
	FetchParticipants()
	PersistUpdatedShare()
	Share()
	AcceptShareInvitationsFromMetadataIntoPersistentStoreCompletion(metadata unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer)
	CanDeleteRecordForManagedObjectWithID(objectID unsafe.Pointer) bool
	CanModifyManagedObjectsInStore(store unsafe.Pointer) bool
	CanUpdateRecordForManagedObjectWithID(objectID unsafe.Pointer) bool
	FetchParticipantsMatchingLookupInfosIntoPersistentStoreCompletion(lookupInfos unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer)
	FetchSharesInPersistentStoreError(persistentStore unsafe.Pointer, error unsafe.Pointer) []unsafe.Pointer
	FetchSharesMatchingObjectIDsError(objectIDs unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	PersistUpdatedShareInPersistentStoreCompletion(share unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer)
	PurgeObjectsAndRecordsInZoneWithIDInPersistentStoreCompletion(zoneID unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer)
	RecordForManagedObjectID(managedObjectID unsafe.Pointer) unsafe.Pointer
	RecordIDForManagedObjectID(managedObjectID unsafe.Pointer) unsafe.Pointer
	RecordIDsForManagedObjectIDs(managedObjectIDs unsafe.Pointer) unsafe.Pointer
	RecordsForManagedObjectIDs(managedObjectIDs unsafe.Pointer) unsafe.Pointer
	ShareManagedObjectsToShareCompletion(managedObjects unsafe.Pointer, share unsafe.Pointer, completion unsafe.Pointer)
}

// A container that encapsulates the Core Data stack in your app, and mirrors select persistent stores to a CloudKit private database.
//
// is a subclass of capable of managing both CloudKit-backed and noncloud stores. By default, contains a single store description, which Core Data assigns to the first CloudKit container identifier in an app’s entitlements. Use to customize this behavior or create additional store descriptions with backing by different containers. For more information about setting up multiple stores, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer
type PersistentCloudKitContainer struct {
	PersistentContainer
}

// PersistentCloudKitContainerFrom constructs a [PersistentCloudKitContainer] from an unsafe.Pointer.
//
// A container that encapsulates the Core Data stack in your app, and mirrors select persistent stores to a CloudKit private database.
func PersistentCloudKitContainerFrom(ptr unsafe.Pointer) PersistentCloudKitContainer {
	return PersistentCloudKitContainer{
		PersistentContainer: PersistentContainerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentCloudKitContainerClass) Alloc() PersistentCloudKitContainer {
	rv := objc.Send[PersistentCloudKitContainer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentCloudKitContainerClass) New() PersistentCloudKitContainer {
	rv := objc.Send[PersistentCloudKitContainer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentCloudKitContainer) Init() PersistentCloudKitContainer {
	rv := objc.Send[PersistentCloudKitContainer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentCloudKitContainer) Autorelease() PersistentCloudKitContainer {
	rv := objc.Send[PersistentCloudKitContainer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentCloudKitContainer creates a new PersistentCloudKitContainer instance.
func NewPersistentCloudKitContainer() PersistentCloudKitContainer {
	return getPersistentCloudKitContainerClass().New()
}


// Creates the CloudKit schema for all stores in the container that manage a CloudKit database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/initializeCloudKitSchema(options:)
func NewPersistentCloudKitContainerializeCloudKitSchemaWithOptionsError(options unsafe.Pointer, error unsafe.Pointer) PersistentCloudKitContainer {
	instance := getPersistentCloudKitContainerClass().Alloc()
	rv := objc.Send[PersistentCloudKitContainer](instance.ID, objc.Sel("initializeCloudKitSchemaWithOptions:error:"), options, error)
	rv.Autorelease()
	return rv
}


// Accepts one or more invitations to participate in sharing using the specified metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746828-acceptshareinvitations
func (p_ PersistentCloudKitContainer) AcceptShareInvitations() {
	objc.Send[objc.ID](p_.ID, objc.Sel("acceptShareInvitations"))
}

// Fetches all participants that match the specified critieria.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746829-fetchparticipants
func (p_ PersistentCloudKitContainer) FetchParticipants() {
	objc.Send[objc.ID](p_.ID, objc.Sel("fetchParticipants"))
}

// Saves the share record and schedules it for export to iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746832-persistupdatedshare
func (p_ PersistentCloudKitContainer) PersistUpdatedShare() {
	objc.Send[objc.ID](p_.ID, objc.Sel("persistUpdatedShare"))
}

// Associates the specified managed objects with a new or existing share record.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746834-share
func (p_ PersistentCloudKitContainer) Share() {
	objc.Send[objc.ID](p_.ID, objc.Sel("share"))
}

// Accepts one or more invitations to participate in sharing using the specified metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/acceptShareInvitationsFromMetadata:intoPersistentStore:completion:
func (p_ PersistentCloudKitContainer) AcceptShareInvitationsFromMetadataIntoPersistentStoreCompletion(metadata unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("acceptShareInvitationsFromMetadata:intoPersistentStore:completion:"), metadata, persistentStore, completion)
}

// Returns a Boolean value that indicates whether the user can delete the managed object’s underlying CloudKit record.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canDeleteRecord(forManagedObjectWith:)
func (p_ PersistentCloudKitContainer) CanDeleteRecordForManagedObjectWithID(objectID unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canDeleteRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Returns a Boolean value that indicates whether the user can modify the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canModifyManagedObjects(in:)
func (p_ PersistentCloudKitContainer) CanModifyManagedObjectsInStore(store unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canModifyManagedObjectsInStore:"), store)
	return rv
}

// Returns a Boolean value that indicates whether the user can modify the managed object’s underlying CloudKit record.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canUpdateRecord(forManagedObjectWith:)
func (p_ PersistentCloudKitContainer) CanUpdateRecordForManagedObjectWithID(objectID unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canUpdateRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Fetches all participants that match the specified critieria.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/fetchParticipantsMatchingLookupInfos:intoPersistentStore:completion:
func (p_ PersistentCloudKitContainer) FetchParticipantsMatchingLookupInfosIntoPersistentStoreCompletion(lookupInfos unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fetchParticipantsMatchingLookupInfos:intoPersistentStore:completion:"), lookupInfos, persistentStore, completion)
}

// Returns an array that contains all share records in the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/fetchSharesInPersistentStore:error:
func (p_ PersistentCloudKitContainer) FetchSharesInPersistentStoreError(persistentStore unsafe.Pointer, error unsafe.Pointer) []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("fetchSharesInPersistentStore:error:"), persistentStore, error)
	return rv
}

// Returns a dictionary that contains the share records that CloudKit associates with specified managed object IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/fetchSharesMatchingObjectIDs:error:
func (p_ PersistentCloudKitContainer) FetchSharesMatchingObjectIDsError(objectIDs unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fetchSharesMatchingObjectIDs:error:"), objectIDs, error)
	return rv
}

// Saves the share record and schedules it for export to iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/persistUpdatedShare:inPersistentStore:completion:
func (p_ PersistentCloudKitContainer) PersistUpdatedShareInPersistentStoreCompletion(share unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("persistUpdatedShare:inPersistentStore:completion:"), share, persistentStore, completion)
}

// Deletes all CloudKit records in the specified record zone, along with their corresponding managed objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/purgeObjectsAndRecordsInZoneWithID:inPersistentStore:completion:
func (p_ PersistentCloudKitContainer) PurgeObjectsAndRecordsInZoneWithIDInPersistentStoreCompletion(zoneID unsafe.Pointer, persistentStore unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("purgeObjectsAndRecordsInZoneWithID:inPersistentStore:completion:"), zoneID, persistentStore, completion)
}

// Returns the CloudKit record for the specified managed object ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/recordForManagedObjectID:
func (p_ PersistentCloudKitContainer) RecordForManagedObjectID(managedObjectID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("recordForManagedObjectID:"), managedObjectID)
	return rv
}

// Returns the CloudKit record ID for the specified managed object ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/recordIDForManagedObjectID:
func (p_ PersistentCloudKitContainer) RecordIDForManagedObjectID(managedObjectID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("recordIDForManagedObjectID:"), managedObjectID)
	return rv
}

// Returns a dictionary that contains the CloudKit record IDs for the specified managed object IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/recordIDsForManagedObjectIDs:
func (p_ PersistentCloudKitContainer) RecordIDsForManagedObjectIDs(managedObjectIDs unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("recordIDsForManagedObjectIDs:"), managedObjectIDs)
	return rv
}

// Returns a dictionary that contains the CloudKit records for the specified managed object IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/recordsForManagedObjectIDs:
func (p_ PersistentCloudKitContainer) RecordsForManagedObjectIDs(managedObjectIDs unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("recordsForManagedObjectIDs:"), managedObjectIDs)
	return rv
}

// Associates the specified managed objects with a new or existing share record.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/shareManagedObjects:toShare:completion:
func (p_ PersistentCloudKitContainer) ShareManagedObjectsToShareCompletion(managedObjects unsafe.Pointer, share unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("shareManagedObjects:toShare:completion:"), managedObjects, share, completion)
}


