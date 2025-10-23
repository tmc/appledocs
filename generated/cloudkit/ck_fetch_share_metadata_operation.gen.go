// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKFetchShareMetadataOperation] class.
var (
	CKFetchShareMetadataOperationClass     _CKFetchShareMetadataOperationClass
	CKFetchShareMetadataOperationClassOnce sync.Once
)

func getCKFetchShareMetadataOperationClass() _CKFetchShareMetadataOperationClass {
	CKFetchShareMetadataOperationClassOnce.Do(func() {
		CKFetchShareMetadataOperationClass = _CKFetchShareMetadataOperationClass{objc.GetClass("CKFetchShareMetadataOperation")}
	})
	return CKFetchShareMetadataOperationClass
}

type _CKFetchShareMetadataOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchShareMetadataOperation] class.
type ICKFetchShareMetadataOperation interface {
	ICKOperation
	PerShareMetadataBlock() unsafe.Pointer
	SetPerShareMetadataBlock(value unsafe.Pointer)
	ShouldFetchRootRecord() bool
	SetShouldFetchRootRecord(value bool)
	FetchShareMetadataCompletionBlock() unsafe.Pointer
	SetFetchShareMetadataCompletionBlock(value unsafe.Pointer)
	FetchShareMetadataResultBlock() unsafe.Pointer
	SetFetchShareMetadataResultBlock(value unsafe.Pointer)
	PerShareMetadataResultBlock() unsafe.Pointer
	SetPerShareMetadataResultBlock(value unsafe.Pointer)
	RootRecordDesiredKeys() unsafe.Pointer
	SetRootRecordDesiredKeys(value unsafe.Pointer)
	ShareURLs() foundation.URL
	SetShareURLs(value foundation.URL)
	CKPartialErrorsByItemIDKey() string
	UserInfo() string
	SetUserInfo(value string)
}

// An operation that fetches metadata for one or more shares.
//
// Use this operation to fetch the metadata for one or more shares. A share’s metadata contains the share and details about the user’s participation. Fetch metadata when you want to manually accept participation in a share using . For a shared record hierarchy, the fetched metadata includes the record ID of the share’s root record. Set to to fetch the entire root record. You can further customize this behavior using to specify which fields you want to include in your fetch. This functionality isn’t applicable for a shared record zone because, unlike a shared record hierarchy, it doesn’t have a nominated root record. To run the operation, add it to any container’s operation queue. Returned metadata includes the ID of the container that stores the share. The operation executes its callbacks on a private serial queue. The operation calls once for each URL you provide, and CloudKit returns the metadata, or an error if the fetch fails. CloudKit also batches per-URL errors. If the operation completes with errors, it returns a error. The error stores individual errors in its dictionary. Use the key to extract them. When all of the following conditions are true, CloudKit returns a error: There are pending participants that don’t have matched iCloud accounts. The current user has an active iCloud account and isn’t an existing participant (pending or otherwise). On receipt of this error, call with the share’s URL to allow CloudKit to verify the user. The following example demonstrates how to create the operation, configure it, and then execute it using the default container’s operation queue:


// An operation that fetches metadata for one or more shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation
type CKFetchShareMetadataOperation struct {
	CKOperation
}

// CKFetchShareMetadataOperationFrom constructs a [CKFetchShareMetadataOperation] from an unsafe.Pointer.
//
// An operation that fetches metadata for one or more shares.
func CKFetchShareMetadataOperationFrom(ptr unsafe.Pointer) CKFetchShareMetadataOperation {
	return CKFetchShareMetadataOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchShareMetadataOperationClass) Alloc() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchShareMetadataOperationClass) New() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchShareMetadataOperation) Init() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchShareMetadataOperation) Autorelease() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchShareMetadataOperation creates a new CKFetchShareMetadataOperation instance.
func NewCKFetchShareMetadataOperation() CKFetchShareMetadataOperation {
	return getCKFetchShareMetadataOperationClass().New()
}



// The closure to execute as the operation fetches individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/perShareMetadataBlock
func (c_ CKFetchShareMetadataOperation) PerShareMetadataBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareMetadataBlock"))
	return rv
}


// The closure to execute as the operation fetches individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/perShareMetadataBlock
func (c_ CKFetchShareMetadataOperation) SetPerShareMetadataBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareMetadataBlock:"), value)
}


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shouldFetchRootRecord
func (c_ CKFetchShareMetadataOperation) ShouldFetchRootRecord() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldFetchRootRecord"))
	return rv
}


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shouldFetchRootRecord
func (c_ CKFetchShareMetadataOperation) SetShouldFetchRootRecord(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldFetchRootRecord:"), value)
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadatacompletionblock
func (c_ CKFetchShareMetadataOperation) FetchShareMetadataCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareMetadataCompletionBlock"))
	return rv
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadatacompletionblock
func (c_ CKFetchShareMetadataOperation) SetFetchShareMetadataCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareMetadataCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) FetchShareMetadataResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareMetadataResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) SetFetchShareMetadataResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareMetadataResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/persharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) PerShareMetadataResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareMetadataResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/persharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) SetPerShareMetadataResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareMetadataResultBlock:"), value)
}


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKFetchShareMetadataOperation) RootRecordDesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rootRecordDesiredKeys"))
	return rv
}


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKFetchShareMetadataOperation) SetRootRecordDesiredKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootRecordDesiredKeys:"), value)
}


// The URLs of the shares to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shareurls
func (c_ CKFetchShareMetadataOperation) ShareURLs() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("shareURLs"))
	return rv
}


// The URLs of the shares to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shareurls
func (c_ CKFetchShareMetadataOperation) SetShareURLs(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareURLs:"), value)
}


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKFetchShareMetadataOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareMetadataOperation) UserInfo() string {
	rv := objc.Send[string](c_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareMetadataOperation) SetUserInfo(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}



