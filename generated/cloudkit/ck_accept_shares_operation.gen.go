// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKAcceptSharesOperation] class.
var (
	CKAcceptSharesOperationClass     _CKAcceptSharesOperationClass
	CKAcceptSharesOperationClassOnce sync.Once
)

func getCKAcceptSharesOperationClass() _CKAcceptSharesOperationClass {
	CKAcceptSharesOperationClassOnce.Do(func() {
		CKAcceptSharesOperationClass = _CKAcceptSharesOperationClass{objc.GetClass("CKAcceptSharesOperation")}
	})
	return CKAcceptSharesOperationClass
}

type _CKAcceptSharesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKAcceptSharesOperation] class.
type ICKAcceptSharesOperation interface {
	ICKOperation
	AcceptSharesCompletionBlock() unsafe.Pointer
	SetAcceptSharesCompletionBlock(value unsafe.Pointer)
	AcceptSharesResultBlock() unsafe.Pointer
	SetAcceptSharesResultBlock(value unsafe.Pointer)
	PerShareCompletionBlock() unsafe.Pointer
	SetPerShareCompletionBlock(value unsafe.Pointer)
	PerShareResultBlock() unsafe.Pointer
	SetPerShareResultBlock(value unsafe.Pointer)
	ShareMetadatas() CKShareMetadata
	SetShareMetadatas(value CKShareMetadata)
	CKPartialErrorsByItemIDKey() string
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	Url() foundation.URL
	SetUrl(value foundation.URL)
	UserInfo() string
	SetUserInfo(value string)
}

// An operation that confirms a user’s participation in a share.
//
// Use this operation to accept participation in one or more shares. You create the operation with an array of share metadatas, which CloudKit provides to your app when the user taps or clicks a share’s . The method CloudKit calls varies by platform and app configuration. For more information, see . You can also fetch a share’s metadata using . If there are several metadatas, group them by their and create an operation for each container. Then add the operation to each container’s operation queue to run it. The operation executes its callbacks on a private serial queue. The operation calls once for each metadata you provide. CloudKit returns the metadata and its related share, or an error if it can’t accept the share. CloudKit also batches per-metadata errors. If the operation completes with errors, it returns a error. The error stores individual errors in its dictionary. Use the key to extract them. After CloudKit applies all record changes, the operation calls . When the closure executes, the server may continue processing residual tasks of the operation, such as creating the record zone in the user’s private database. The following example demonstrates how to accept a share that CloudKit provides to your window scene delegate. It shows how to create the operation, configure it, and execute it in the correct container:


// An operation that confirms a user’s participation in a share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
type CKAcceptSharesOperation struct {
	CKOperation
}

// CKAcceptSharesOperationFrom constructs a [CKAcceptSharesOperation] from an unsafe.Pointer.
//
// An operation that confirms a user’s participation in a share.
func CKAcceptSharesOperationFrom(ptr unsafe.Pointer) CKAcceptSharesOperation {
	return CKAcceptSharesOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKAcceptSharesOperationClass) Alloc() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKAcceptSharesOperationClass) New() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKAcceptSharesOperation) Init() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKAcceptSharesOperation) Autorelease() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAcceptSharesOperation creates a new CKAcceptSharesOperation instance.
func NewCKAcceptSharesOperation() CKAcceptSharesOperation {
	return getCKAcceptSharesOperationClass().New()
}



// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharescompletionblock
func (c_ CKAcceptSharesOperation) AcceptSharesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acceptSharesCompletionBlock"))
	return rv
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharescompletionblock
func (c_ CKAcceptSharesOperation) SetAcceptSharesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcceptSharesCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharesresultblock
func (c_ CKAcceptSharesOperation) AcceptSharesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acceptSharesResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharesresultblock
func (c_ CKAcceptSharesOperation) SetAcceptSharesResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcceptSharesResultBlock:"), value)
}


// The block to execute as CloudKit processes individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/persharecompletionblock
func (c_ CKAcceptSharesOperation) PerShareCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareCompletionBlock"))
	return rv
}


// The block to execute as CloudKit processes individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/persharecompletionblock
func (c_ CKAcceptSharesOperation) SetPerShareCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/pershareresultblock
func (c_ CKAcceptSharesOperation) PerShareResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/pershareresultblock
func (c_ CKAcceptSharesOperation) SetPerShareResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareResultBlock:"), value)
}


// The share metadatas to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/sharemetadatas
func (c_ CKAcceptSharesOperation) ShareMetadatas() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c_.ID, objc.Sel("shareMetadatas"))
	return rv
}


// The share metadatas to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/sharemetadatas
func (c_ CKAcceptSharesOperation) SetShareMetadatas(value CKShareMetadata) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareMetadatas:"), value)
}


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKAcceptSharesOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}


// The ID of the share’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/metadata/containeridentifier
func (c_ CKAcceptSharesOperation) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The ID of the share’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/metadata/containeridentifier
func (c_ CKAcceptSharesOperation) SetContainerIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKAcceptSharesOperation) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKAcceptSharesOperation) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKAcceptSharesOperation) UserInfo() string {
	rv := objc.Send[string](c_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKAcceptSharesOperation) SetUserInfo(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}



