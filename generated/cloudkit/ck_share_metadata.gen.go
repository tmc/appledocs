// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKShareMetadata] class.
var (
	CKShareMetadataClass     _CKShareMetadataClass
	CKShareMetadataClassOnce sync.Once
)

func getCKShareMetadataClass() _CKShareMetadataClass {
	CKShareMetadataClassOnce.Do(func() {
		CKShareMetadataClass = _CKShareMetadataClass{objc.GetClass("CKShareMetadata")}
	})
	return CKShareMetadataClass
}

type _CKShareMetadataClass struct {
	class objc.Class
}

// An interface definition for the [CKShareMetadata] class.
type ICKShareMetadata interface {
	objectivec.IObject
}

// An object that describes a shared record’s metadata.
//
// A share’s metadata is an intermediary object that provides access to the share, its owner, and, for a shared record hierarchy, its root record. Metadata also includes details about the current user’s participation in the share. You don’t create metadata. CloudKit provides it to your app when the user taps or clicks a share’s , such as in an email or a message. The method CloudKit calls varies by platform and app configuration, and includes the following: For a scene-based iOS app in a running or suspended state, CloudKit calls the method on your window scene delegate. For a scene-based iOS app that’s not running, the system launches your app in response to the tap or click, and calls the method on your scene delegate. The parameter contains the metadata. Use its property to access it. For an iOS app that doesn’t use scenes, CloudKit calls your app delegate’s method. For a macOS app, CloudKit calls your app delegate’s method. For a watchOS app, CloudKit calls the method on your watch extension delegate. Respond by checking the of the provided metadata. If the status is , use to accept participation in the share. You can also fetch metadata independent of this flow using . For a shared record hierarchy, the property contains the ID of the share’s root record. When using to fetch metadata, you can include the entire root record by setting the operation’s property to . CloudKit then populates the property before it returns the metadata. You can further customize this behavior using the operation’s property to specify which fields to return. This functionality isn’t applicable for a shared record zone because, unlike a shared record hierarchy, it doesn’t have a nominated root record. The participant properties provide the current user’s acceptance status, permissions, and role. Use these values to determine what functionality to provide to the user. For example, only display editing controls for accepted participants with permissions.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata
type CKShareMetadata struct {
	objectivec.Object
}

// CKShareMetadataFrom constructs a [CKShareMetadata] from an unsafe.Pointer.
//
// An object that describes a shared record’s metadata.
func CKShareMetadataFrom(ptr unsafe.Pointer) CKShareMetadata {
	return CKShareMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKShareMetadataClass) Alloc() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKShareMetadataClass) New() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareMetadata) Init() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareMetadata) Autorelease() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareMetadata creates a new CKShareMetadata instance.
func NewCKShareMetadata() CKShareMetadata {
	return getCKShareMetadataClass().New()
}


// The ID of the share’s container.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/containerIdentifier
func (c_ CKShareMetadata) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}

// The record ID of the shared hierarchy’s root record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/hierarchicalRootRecordID
func (c_ CKShareMetadata) HierarchicalRootRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("hierarchicalRootRecordID"))
	return rv
}

// The identity of the share’s owner.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/ownerIdentity
func (c_ CKShareMetadata) OwnerIdentity() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("ownerIdentity"))
	return rv
}

// The share’s permissions for the user who retrieves the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantPermission
func (c_ CKShareMetadata) ParticipantPermission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c_.ID, objc.Sel("participantPermission"))
	return rv
}

// The share’s participant role for the user who retrieves the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantRole
func (c_ CKShareMetadata) ParticipantRole() CKShareParticipantRole {
	rv := objc.Send[CKShareParticipantRole](c_.ID, objc.Sel("participantRole"))
	return rv
}

// The share’s participation status for the user who retrieves the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantStatus
func (c_ CKShareMetadata) ParticipantStatus() CKShareParticipantAcceptanceStatus {
	rv := objc.Send[CKShareParticipantAcceptanceStatus](c_.ID, objc.Sel("participantStatus"))
	return rv
}

// The share’s participation type for the user who retrieves the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantType
func (c_ CKShareMetadata) ParticipantType() CKShareParticipantType {
	rv := objc.Send[CKShareParticipantType](c_.ID, objc.Sel("participantType"))
	return rv
}

// The share’s root record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/rootRecord
func (c_ CKShareMetadata) RootRecord() CKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("rootRecord"))
	return rv
}

// The record ID of the share’s root record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/rootRecordID
func (c_ CKShareMetadata) RootRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("rootRecordID"))
	return rv
}

// The share that owns the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/share
func (c_ CKShareMetadata) Share() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("share"))
	return rv
}

// The fields to return when fetching the root record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKShareMetadata) RootRecordDesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rootRecordDesiredKeys"))
	return rv
}


// SetRootRecordDesiredKeys sets the value of the rootRecordDesiredKeys property.
// The fields to return when fetching the root record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKShareMetadata) SetRootRecordDesiredKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootRecordDesiredKeys:"), value)
}

// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shouldfetchrootrecord
func (c_ CKShareMetadata) ShouldFetchRootRecord() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldFetchRootRecord"))
	return rv
}


// SetShouldFetchRootRecord sets the value of the shouldFetchRootRecord property.
// A Boolean value that indicates whether to retrieve the root record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shouldfetchrootrecord
func (c_ CKShareMetadata) SetShouldFetchRootRecord(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldFetchRootRecord:"), value)
}

// The URL for inviting participants to the share.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKShareMetadata) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL for inviting participants to the share.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKShareMetadata) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}



