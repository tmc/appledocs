// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKContainer */


/* debug [class_header]: Header for CKContainer */
// The class instance for the [CKContainer] class.
var (
	CKContainerClass     _CKContainerClass
	CKContainerClassOnce sync.Once
)

func getCKContainerClass() _CKContainerClass {
	CKContainerClassOnce.Do(func() {
		CKContainerClass = _CKContainerClass{objc.GetClass("CKContainer")}
	})
	return CKContainerClass
}

type _CKContainerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKContainer */
// An interface definition for the [CKContainer] class.
type ICKContainer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKContainer */
	// properties:
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	PrivateCloudDatabase() ICKDatabase
	PublicCloudDatabase() ICKDatabase
	SharedCloudDatabase() ICKDatabase
	CKCurrentUserDefaultName() objc.IObject /* cross-framework: NSString */
	CKOwnerDefaultName() objc.IObject /* cross-framework: NSString */
	UserRecordID() ICKRecordID
	SetUserRecordID(value ICKRecordID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKContainer */
	// methods:
	AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler unsafe.Pointer)
	AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer)
	AddOperation(operation objc.IObject /* cross-framework: CKOperation */)
	DatabaseWithDatabaseScope(databaseScope CKDatabaseScope) ICKDatabase
	FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchLongLivedOperationWithIDCompletionHandler(operationID objectivec.IObject, completionHandler unsafe.Pointer)
	FetchShareMetadataWithURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer)
	FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKContainer */
// Alloc allocates a new instance without initialization.
func (cc _CKContainerClass) Alloc() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKContainerClass) New() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKContainer) Init() CKContainer {
	rv := objc.Send[CKContainer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKContainer) Autorelease() CKContainer {
	rv := objc.Send[CKContainer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKContainer creates a new CKContainer instance.
func NewCKContainer() CKContainer {
	return getCKContainerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKContainer */
// A conduit to your app’s databases.
//
// A container manages all explicit and implicit attempts to access its contents. Every app has a default container that manages its own content. If you develop a suite of apps, you can access any containers that you have the appropriate entitlements for. Each new container distinguishes between public and private data. CloudKit always stores private data in the appropriate container directory in the user’s iCloud account.


// A conduit to your app’s databases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer
type CKContainer struct {
	objectivec.Object
}

// CKContainerFrom constructs a [CKContainer] from an unsafe.Pointer.
//
// A conduit to your app’s databases.
func CKContainerFrom(ptr unsafe.Pointer) CKContainer {
	return CKContainer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKContainer */

// Creates a container for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func NewCKContainerWithIdentifier(containerIdentifier objc.IObject /* cross-framework: NSString */) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(getCKContainerClass().class), objc.Sel("containerWithIdentifier:"), containerIdentifier)
	return rv
}/* debug [class_init_methods/constructor]: NewCKContainerWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKContainer */

// Returns the app’s default container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/default()
func (cc _CKContainerClass) DefaultContainer() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("defaultContainer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultContainer) */


// Creates a container for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func (cc _CKContainerClass) ContainerWithIdentifier(containerIdentifier objc.IObject /* cross-framework: NSString */) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("containerWithIdentifier:"), containerIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContainerWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKContainer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKContainer */

// Accepts the specified share metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/accept(_:completionHandler:)-949ea
func (c_ CKContainer) AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("acceptShareMetadata:completionHandler:"), metadata, completionHandler)
}/* debug [instance_methods/method]: AcceptShareMetadataCompletionHandler */


// Determines whether the system can access the user’s iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/accountStatus(completionHandler:)
func (c_ CKContainer) AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("accountStatusWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: AccountStatusWithCompletionHandler */


// Adds an operation to the container’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/add(_:)
func (c_ CKContainer) AddOperation(operation objc.IObject /* cross-framework: CKOperation */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOperation:"), operation)
}/* debug [instance_methods/method]: AddOperation */


// Returns the database with the specified scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/database(with:)
func (c_ CKContainer) DatabaseWithDatabaseScope(databaseScope CKDatabaseScope) ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("databaseWithDatabaseScope:"), databaseScope)
	return rv
}/* debug [instance_methods/method]: DatabaseWithDatabaseScope */


// Fetches the IDs of any long-lived operations that are running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchAllLongLivedOperationIDsWithCompletionHandler:
func (c_ CKContainer) FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllLongLivedOperationIDsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchAllLongLivedOperationIDsWithCompletionHandler */


// Fetches the long-lived operation for the specified operation ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchLongLivedOperationWithID:completionHandler:
func (c_ CKContainer) FetchLongLivedOperationWithIDCompletionHandler(operationID objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchLongLivedOperationWithID:completionHandler:"), operationID, completionHandler)
}/* debug [instance_methods/method]: FetchLongLivedOperationWithIDCompletionHandler */


// Fetches the share metadata for the specified share URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareMetadata(with:completionHandler:)
func (c_ CKContainer) FetchShareMetadataWithURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareMetadataWithURL:completionHandler:"), url, completionHandler)
}/* debug [instance_methods/method]: FetchShareMetadataWithURLCompletionHandler */


// Fetches the share participant with the specified email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withEmailAddress:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithEmailAddress:completionHandler:"), emailAddress, completionHandler)
}/* debug [instance_methods/method]: FetchShareParticipantWithEmailAddressCompletionHandler */


// Fetches the share participant with the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withPhoneNumber:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithPhoneNumber:completionHandler:"), phoneNumber, completionHandler)
}/* debug [instance_methods/method]: FetchShareParticipantWithPhoneNumberCompletionHandler */


// Fetches the share participant with the specified user record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withUserRecordID:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithUserRecordID:completionHandler:"), userRecordID, completionHandler)
}/* debug [instance_methods/method]: FetchShareParticipantWithUserRecordIDCompletionHandler */


// Fetches the user record ID of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchUserRecordID(completionHandler:)
func (c_ CKContainer) FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchUserRecordIDWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchUserRecordIDWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKContainer */

// The container’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/containerIdentifier
func (c_ CKContainer) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The user’s private database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/privateCloudDatabase
func (c_ CKContainer) PrivateCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("privateCloudDatabase"))
	return rv
}/* debug [instance_properties/getter]: privateCloudDatabase */


// The app’s public database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/publicCloudDatabase
func (c_ CKContainer) PublicCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("publicCloudDatabase"))
	return rv
}/* debug [instance_properties/getter]: publicCloudDatabase */


// The database that contains shared data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/sharedCloudDatabase
func (c_ CKContainer) SharedCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("sharedCloudDatabase"))
	return rv
}/* debug [instance_properties/getter]: sharedCloudDatabase */


// A constant that provides the current user’s default name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c_ CKContainer) CKCurrentUserDefaultName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKCurrentUserDefaultName"))
	return rv
}/* debug [instance_properties/getter]: CKCurrentUserDefaultName */


// A constant that provides the default owner’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckownerdefaultname
func (c_ CKContainer) CKOwnerDefaultName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKOwnerDefaultName"))
	return rv
}/* debug [instance_properties/getter]: CKOwnerDefaultName */


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) UserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}/* debug [instance_properties/getter]: userRecordID */


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) SetUserRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserRecordID:"), value)
}/* debug [instance_properties/setter]: userRecordID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKContainer */


