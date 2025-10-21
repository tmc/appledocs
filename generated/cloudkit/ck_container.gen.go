// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKContainer] class.
type ICKContainer interface {
	objectivec.IObject
	AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler unsafe.Pointer)
	AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer)
	AddOperation(operation ICKOperation)
	DatabaseWithDatabaseScope(databaseScope ICKDatabaseScope) CKDatabase
	DiscoverAllIdentitiesWithCompletionHandler(completionHandler unsafe.Pointer)
	DiscoverUserIdentityWithEmailAddressCompletionHandler(email string, completionHandler unsafe.Pointer)
	DiscoverUserIdentityWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler unsafe.Pointer)
	DiscoverUserIdentityWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer)
	FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchLongLivedOperationWithIDCompletionHandler(operationID unsafe.Pointer, completionHandler unsafe.Pointer)
	FetchShareMetadataWithURLCompletionHandler(url foundation.IURL, completionHandler unsafe.Pointer)
	FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler unsafe.Pointer)
	FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler unsafe.Pointer)
	FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer)
	FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer)
	RequestApplicationPermissionCompletionHandler(applicationPermission ICKApplicationPermissions, completionHandler unsafe.Pointer)
	StatusForApplicationPermissionCompletionHandler(applicationPermission ICKApplicationPermissions, completionHandler unsafe.Pointer)
}

// A conduit to your app’s databases.
//
// A container manages all explicit and implicit attempts to access its contents. Every app has a default container that manages its own content. If you develop a suite of apps, you can access any containers that you have the appropriate entitlements for. Each new container distinguishes between public and private data. CloudKit always stores private data in the appropriate container directory in the user’s iCloud account.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKContainerClass) Alloc() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a container for the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func NewCKContainerWithIdentifier(containerIdentifier string) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(getCKContainerClass().class), objc.Sel("containerWithIdentifier:"), objc.String(containerIdentifier))
	return rv
}


// Returns the app’s default container.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/default()
func (cc _CKContainerClass) DefaultContainer() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("defaultContainer"))
	return rv
}

// Creates a container for the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func (cc _CKContainerClass) ContainerWithIdentifier(containerIdentifier string) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("containerWithIdentifier:"), objc.String(containerIdentifier))
	return rv
}

// Accepts the specified share metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/accept(_:completionHandler:)-949ea
func (c_ CKContainer) AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("acceptShareMetadata:completionHandler:"), metadata, completionHandler)
}

// Determines whether the system can access the user’s iCloud account.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/accountStatus(completionHandler:)
func (c_ CKContainer) AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("accountStatusWithCompletionHandler:"), completionHandler)
}

// Adds an operation to the container’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/add(_:)
func (c_ CKContainer) AddOperation(operation ICKOperation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOperation:"), operation)
}

// Returns the database with the specified scope.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/database(with:)
func (c_ CKContainer) DatabaseWithDatabaseScope(databaseScope ICKDatabaseScope) CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("databaseWithDatabaseScope:"), databaseScope)
	return rv
}

// Fetches all user identities that match entries in the user’s Contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/discoverAllIdentities(completionHandler:)
func (c_ CKContainer) DiscoverAllIdentitiesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverAllIdentitiesWithCompletionHandler:"), completionHandler)
}

// Fetches the user identity for the specified email address.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/discoverUserIdentity(withEmailAddress:completionHandler:)
func (c_ CKContainer) DiscoverUserIdentityWithEmailAddressCompletionHandler(email string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverUserIdentityWithEmailAddress:completionHandler:"), objc.String(email), completionHandler)
}

// Fetches the user identity for the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/discoverUserIdentity(withPhoneNumber:completionHandler:)
func (c_ CKContainer) DiscoverUserIdentityWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverUserIdentityWithPhoneNumber:completionHandler:"), objc.String(phoneNumber), completionHandler)
}

// Fetches the user identity for the specified user record ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/discoverUserIdentity(withUserRecordID:completionHandler:)
func (c_ CKContainer) DiscoverUserIdentityWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverUserIdentityWithUserRecordID:completionHandler:"), userRecordID, completionHandler)
}

// Fetches the IDs of any long-lived operations that are running.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchAllLongLivedOperationIDsWithCompletionHandler:
func (c_ CKContainer) FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchAllLongLivedOperationIDsWithCompletionHandler:"), completionHandler)
}

// Fetches the long-lived operation for the specified operation ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchLongLivedOperationWithID:completionHandler:
func (c_ CKContainer) FetchLongLivedOperationWithIDCompletionHandler(operationID unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchLongLivedOperationWithID:completionHandler:"), operationID, completionHandler)
}

// Fetches the share metadata for the specified share URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareMetadata(with:completionHandler:)
func (c_ CKContainer) FetchShareMetadataWithURLCompletionHandler(url foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareMetadataWithURL:completionHandler:"), url, completionHandler)
}

// Fetches the share participant with the specified email address.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withEmailAddress:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithEmailAddress:completionHandler:"), objc.String(emailAddress), completionHandler)
}

// Fetches the share participant with the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withPhoneNumber:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithPhoneNumber:completionHandler:"), objc.String(phoneNumber), completionHandler)
}

// Fetches the share participant with the specified user record ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withUserRecordID:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithUserRecordID:completionHandler:"), userRecordID, completionHandler)
}

// Fetches the user record ID of the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchUserRecordID(completionHandler:)
func (c_ CKContainer) FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchUserRecordIDWithCompletionHandler:"), completionHandler)
}

// Prompts the user to authorize the specified permission.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/requestApplicationPermission(_:completionHandler:)
func (c_ CKContainer) RequestApplicationPermissionCompletionHandler(applicationPermission ICKApplicationPermissions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestApplicationPermission:completionHandler:"), applicationPermission, completionHandler)
}

// Determines the authorization status of the specified permission.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/status(forApplicationPermission:completionHandler:)
func (c_ CKContainer) StatusForApplicationPermissionCompletionHandler(applicationPermission ICKApplicationPermissions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("statusForApplicationPermission:completionHandler:"), applicationPermission, completionHandler)
}

// The container’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/containerIdentifier
func (c_ CKContainer) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}

// The user’s private database.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/privateCloudDatabase
func (c_ CKContainer) PrivateCloudDatabase() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("privateCloudDatabase"))
	return rv
}

// The app’s public database.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/publicCloudDatabase
func (c_ CKContainer) PublicCloudDatabase() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("publicCloudDatabase"))
	return rv
}

// The database that contains shared data.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/sharedCloudDatabase
func (c_ CKContainer) SharedCloudDatabase() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("sharedCloudDatabase"))
	return rv
}

// A constant that provides the current user’s default name.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c_ CKContainer) CKCurrentUserDefaultName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKCurrentUserDefaultName"))
	return rv
}

// A constant that provides the default owner’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckownerdefaultname
func (c_ CKContainer) CKOwnerDefaultName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKOwnerDefaultName"))
	return rv
}

// The user record ID for the corresponding user record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) UserRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}


// SetUserRecordID sets the value of the userRecordID property.
// The user record ID for the corresponding user record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) SetUserRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserRecordID:"), value)
}


