// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	ContainerIdentifier() string /* primitive/slice/pointer. */
	SetContainerIdentifier(value string /* primitive/slice/pointer. */)
	PrivateCloudDatabase() ICKDatabase
	SetPrivateCloudDatabase(value ICKDatabase)
	PublicCloudDatabase() ICKDatabase
	SetPublicCloudDatabase(value ICKDatabase)
	SharedCloudDatabase() ICKDatabase
	SetSharedCloudDatabase(value ICKDatabase)
	CKCurrentUserDefaultName() string /* primitive/slice/pointer. */
	CKOwnerDefaultName() string /* primitive/slice/pointer. */
	UserRecordID() objc.IObject /* cross-framework: CKRecordID */
	SetUserRecordID(value objc.IObject /* cross-framework: CKRecordID */)
	// methods:
	AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID objc.IObject /* cross-framework CKRecordID */, completionHandler unsafe.Pointer)
	FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer)
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func NewCKContainerWithIdentifier(containerIdentifier string /* primitive/slice/pointer. */) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(getCKContainerClass().class), objc.Sel("containerWithIdentifier:"), objc.String(containerIdentifier))
	return rv
}



// Returns the app’s default container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/default()
func (cc _CKContainerClass) DefaultContainer() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("defaultContainer"))
	return rv
}


// Creates a container for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func (cc _CKContainerClass) ContainerWithIdentifier(containerIdentifier string /* primitive/slice/pointer. */) CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("containerWithIdentifier:"), objc.String(containerIdentifier))
	return rv
}


// Determines whether the system can access the user’s iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/accountStatus(completionHandler:)
func (c_ CKContainer) AccountStatusWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("accountStatusWithCompletionHandler:"), completionHandler)
}


// Fetches the share participant with the specified user record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withUserRecordID:completionHandler:)
func (c_ CKContainer) FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID objc.IObject /* cross-framework CKRecordID */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchShareParticipantWithUserRecordID:completionHandler:"), userRecordID, completionHandler)
}


// Fetches the user record ID of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchUserRecordID(completionHandler:)
func (c_ CKContainer) FetchUserRecordIDWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchUserRecordIDWithCompletionHandler:"), completionHandler)
}


// The container’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/containeridentifier
func (c_ CKContainer) ContainerIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The container’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/containeridentifier
func (c_ CKContainer) SetContainerIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// The user’s private database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/privateclouddatabase
func (c_ CKContainer) PrivateCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("privateCloudDatabase"))
	return rv
}


// The user’s private database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/privateclouddatabase
func (c_ CKContainer) SetPrivateCloudDatabase(value ICKDatabase) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrivateCloudDatabase:"), value)
}


// The app’s public database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/publicclouddatabase
func (c_ CKContainer) PublicCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("publicCloudDatabase"))
	return rv
}


// The app’s public database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/publicclouddatabase
func (c_ CKContainer) SetPublicCloudDatabase(value ICKDatabase) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublicCloudDatabase:"), value)
}


// The database that contains shared data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/sharedclouddatabase
func (c_ CKContainer) SharedCloudDatabase() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("sharedCloudDatabase"))
	return rv
}


// The database that contains shared data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/sharedclouddatabase
func (c_ CKContainer) SetSharedCloudDatabase(value ICKDatabase) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSharedCloudDatabase:"), value)
}


// A constant that provides the current user’s default name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c_ CKContainer) CKCurrentUserDefaultName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CKCurrentUserDefaultName"))
	return rv
}


// A constant that provides the default owner’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckownerdefaultname
func (c_ CKContainer) CKOwnerDefaultName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CKOwnerDefaultName"))
	return rv
}


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) UserRecordID() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKContainer) SetUserRecordID(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserRecordID:"), value)
}


