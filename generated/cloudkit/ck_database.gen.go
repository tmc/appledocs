// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKDatabase */


/* debug [class_header]: Header for CKDatabase */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKDatabase */
// An interface definition for the [CKDatabase] class.
type ICKDatabase interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKDatabase */
	// properties:
	DatabaseScope() CKDatabaseScope
	QualityOfService() objectivec.IObject
	SetQualityOfService(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKDatabase */
	// methods:
	AddOperation(operation ICKDatabaseOperation)
	DeleteRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler unsafe.Pointer)
	DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer)
	DeleteSubscriptionWithIDCompletionHandler(subscriptionID objectivec.IObject, completionHandler unsafe.Pointer)
	FetchRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler unsafe.Pointer)
	FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer)
	FetchAllRecordZonesWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchAllSubscriptionsWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchSubscriptionWithIDCompletionHandler(subscriptionID objectivec.IObject, completionHandler unsafe.Pointer)
	SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler unsafe.Pointer)
	SaveRecordCompletionHandler(record objc.IObject /* cross-framework: CKRecord */, completionHandler unsafe.Pointer)
	SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKDatabase */
// Alloc allocates a new instance without initialization.
func (cc _CKDatabaseClass) Alloc() CKDatabase {
	rv := objc.Send[CKDatabase](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKDatabase */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKDatabase *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKDatabase */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKDatabase */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKDatabase */

// Executes the specified operation in the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/add(_:)
func (c_ CKDatabase) AddOperation(operation ICKDatabaseOperation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOperation:"), operation)
}/* debug [instance_methods/method]: AddOperation */


// Deletes a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordID:completionHandler:)
func (c_ CKDatabase) DeleteRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteRecordWithID:completionHandler:"), recordID, completionHandler)
}/* debug [instance_methods/method]: DeleteRecordWithIDCompletionHandler */


// Deletes a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordZoneID:completionHandler:)
func (c_ CKDatabase) DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteRecordZoneWithID:completionHandler:"), zoneID, completionHandler)
}/* debug [instance_methods/method]: DeleteRecordZoneWithIDCompletionHandler */


// Deletes a specific subscription and delivers the deleted subscription’s identifier to a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/deleteSubscriptionWithID:completionHandler:
func (c_ CKDatabase) DeleteSubscriptionWithIDCompletionHandler(subscriptionID objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}/* debug [instance_methods/method]: DeleteSubscriptionWithIDCompletionHandler */


// Fetches a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordID:completionHandler:)
func (c_ CKDatabase) FetchRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchRecordWithID:completionHandler:"), recordID, completionHandler)
}/* debug [instance_methods/method]: FetchRecordWithIDCompletionHandler */


// Fetches a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordZoneID:completionHandler:)
func (c_ CKDatabase) FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchRecordZoneWithID:completionHandler:"), zoneID, completionHandler)
}/* debug [instance_methods/method]: FetchRecordZoneWithIDCompletionHandler */


// Fetches all record zones from the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchAllRecordZones(completionHandler:)
func (c_ CKDatabase) FetchAllRecordZonesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllRecordZonesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchAllRecordZonesWithCompletionHandler */


// Fetches all subscriptions from the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchAllSubscriptions(completionHandler:)
func (c_ CKDatabase) FetchAllSubscriptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllSubscriptionsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchAllSubscriptionsWithCompletionHandler */


// Fetches a specific subscription and delivers it to a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetchSubscriptionWithID:completionHandler:
func (c_ CKDatabase) FetchSubscriptionWithIDCompletionHandler(subscriptionID objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}/* debug [instance_methods/method]: FetchSubscriptionWithIDCompletionHandler */


// Saves a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-32ffr
func (c_ CKDatabase) SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveRecordZone:completionHandler:"), zone, completionHandler)
}/* debug [instance_methods/method]: SaveRecordZoneCompletionHandler */


// Saves a specific record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-3tatz
func (c_ CKDatabase) SaveRecordCompletionHandler(record objc.IObject /* cross-framework: CKRecord */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveRecord:completionHandler:"), record, completionHandler)
}/* debug [instance_methods/method]: SaveRecordCompletionHandler */


// Saves a specific subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-9pona
func (c_ CKDatabase) SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveSubscription:completionHandler:"), subscription, completionHandler)
}/* debug [instance_methods/method]: SaveSubscriptionCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKDatabase */

// The type of database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/databaseScope
func (c_ CKDatabase) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c_.ID, objc.Sel("databaseScope"))
	return rv
}/* debug [instance_properties/getter]: databaseScope */


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/qualityofservice
func (c_ CKDatabase) QualityOfService() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/qualityofservice
func (c_ CKDatabase) SetQualityOfService(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKDatabase */



