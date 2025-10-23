// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKDatabase] class.
var (
	CKDatabaseClass     _CKDatabaseClass
	CKDatabaseClassOnce sync.Once
)

func getCKDatabaseClass() _CKDatabaseClass {
	CKDatabaseClassOnce.Do(func() {
		CKDatabaseClass = _CKDatabaseClass{objc.GetClass("CKDatabase")}
	})
	return CKDatabaseClass
}

type _CKDatabaseClass struct {
	class objc.Class
}

// An interface definition for the [CKDatabase] class.
type ICKDatabase interface {
	objectivec.IObject
	DatabaseScope() CKDatabaseScope
	QualityOfService() unsafe.Pointer
	SetQualityOfService(value unsafe.Pointer)
	AddOperation(operation ICKDatabaseOperation)
	DeleteRecordWithIDCompletionHandler(recordID CKRecordID, completionHandler unsafe.Pointer)
	DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer)
	DeleteSubscriptionWithIDCompletionHandler(subscriptionID unsafe.Pointer, completionHandler unsafe.Pointer)
	FetchRecordWithIDCompletionHandler(recordID CKRecordID, completionHandler unsafe.Pointer)
	FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer)
	FetchAllRecordZonesWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchAllSubscriptionsWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchSubscriptionWithIDCompletionHandler(subscriptionID unsafe.Pointer, completionHandler unsafe.Pointer)
	SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler unsafe.Pointer)
	SaveRecordCompletionHandler(record ICKRecord, completionHandler unsafe.Pointer)
	SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler unsafe.Pointer)
}

// An object that represents a collection of record zones and subscriptions.
//
// A database takes requests and operations and applies them to the objects it contains, whether that’s record zones, records, or subscriptions. Each of your app’s users has access to the three separate databases: A public database that’s accessible to all users of your app. A private database that’s accessible only to the user of the current device. A shared database that’s accessible only to the user of the current device, which contains records that other iCloud users share with them. The public database is always available, even when the device doesn’t have an active iCloud account. In this scenario, your app can fetch specific records and perform searches, but it can’t create or modify records. CloudKit requires an iCloud account for writing to the public database so it can identify the authors of any changes. All access to the private and shared databases requires an iCloud account. You don’t create instances of , nor do you subclass it. Instead, you access the required database using one of your app’s containers. For more information, see . By default, CloudKit executes the methods in this class with a low-priority quality of service (QoS). To use a higher-priority QoS, perform the following: Create an instance of and set its property to the preferred value. Call the databaseʼs method and provide the configuration and a trailing closure. In the closure, use the provided database to execute the relevant methods at the preferred QoS.


// An object that represents a collection of record zones and subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase
type CKDatabase struct {
	objectivec.Object
}

// CKDatabaseFrom constructs a [CKDatabase] from an unsafe.Pointer.
//
// An object that represents a collection of record zones and subscriptions.
func CKDatabaseFrom(ptr unsafe.Pointer) CKDatabase {
	return CKDatabase{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKDatabaseClass) Alloc() CKDatabase {
	rv := objc.Send[CKDatabase](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKDatabaseClass) New() CKDatabase {
	rv := objc.Send[CKDatabase](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDatabase) Init() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDatabase) Autorelease() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabase creates a new CKDatabase instance.
func NewCKDatabase() CKDatabase {
	return getCKDatabaseClass().New()
}



// Executes the specified operation in the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/add(_:)
func (c_ CKDatabase) AddOperation(operation ICKDatabaseOperation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOperation:"), operation)
}


// Deletes a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordID:completionHandler:)
func (c_ CKDatabase) DeleteRecordWithIDCompletionHandler(recordID CKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteRecordWithID:completionHandler:"), recordID, completionHandler)
}


// Deletes a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordZoneID:completionHandler:)
func (c_ CKDatabase) DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteRecordZoneWithID:completionHandler:"), zoneID, completionHandler)
}


// Deletes a specific subscription and delivers the deleted subscription’s identifier to a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/deleteSubscriptionWithID:completionHandler:
func (c_ CKDatabase) DeleteSubscriptionWithIDCompletionHandler(subscriptionID unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}


// Fetches a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordID:completionHandler:)
func (c_ CKDatabase) FetchRecordWithIDCompletionHandler(recordID CKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchRecordWithID:completionHandler:"), recordID, completionHandler)
}


// Fetches a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordZoneID:completionHandler:)
func (c_ CKDatabase) FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchRecordZoneWithID:completionHandler:"), zoneID, completionHandler)
}


// Fetches all record zones from the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchAllRecordZones(completionHandler:)
func (c_ CKDatabase) FetchAllRecordZonesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllRecordZonesWithCompletionHandler:"), completionHandler)
}


// Fetches all subscriptions from the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchAllSubscriptions(completionHandler:)
func (c_ CKDatabase) FetchAllSubscriptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllSubscriptionsWithCompletionHandler:"), completionHandler)
}


// Fetches a specific subscription and delivers it to a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchSubscriptionWithID:completionHandler:
func (c_ CKDatabase) FetchSubscriptionWithIDCompletionHandler(subscriptionID unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}


// Saves a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-32ffr
func (c_ CKDatabase) SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveRecordZone:completionHandler:"), zone, completionHandler)
}


// Saves a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-3tatz
func (c_ CKDatabase) SaveRecordCompletionHandler(record ICKRecord, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveRecord:completionHandler:"), record, completionHandler)
}


// Saves a specific subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-9pona
func (c_ CKDatabase) SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveSubscription:completionHandler:"), subscription, completionHandler)
}


// The type of database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/databaseScope
func (c_ CKDatabase) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c_.ID, objc.Sel("databaseScope"))
	return rv
}


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/qualityofservice
func (c_ CKDatabase) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/qualityofservice
func (c_ CKDatabase) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}



