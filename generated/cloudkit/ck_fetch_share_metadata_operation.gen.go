// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchShareMetadataOperation */


/* debug [class_header]: Header for CKFetchShareMetadataOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchShareMetadataOperation */
// An interface definition for the [CKFetchShareMetadataOperation] class.
type ICKFetchShareMetadataOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKFetchShareMetadataOperation */
	// properties:
	FetchShareMetadataCompletionBlock() unsafe.Pointer
	SetFetchShareMetadataCompletionBlock(value unsafe.Pointer)
	PerShareMetadataBlock() unsafe.Pointer
	SetPerShareMetadataBlock(value unsafe.Pointer)
	RootRecordDesiredKeys() []string
	SetRootRecordDesiredKeys(value []string)
	ShareURLs() []foundation.URL
	SetShareURLs(value []foundation.URL)
	ShouldFetchRootRecord() bool
	SetShouldFetchRootRecord(value bool)
	FetchShareMetadataResultBlock() objectivec.IObject
	SetFetchShareMetadataResultBlock(value objectivec.IObject)
	PerShareMetadataResultBlock() objectivec.IObject
	SetPerShareMetadataResultBlock(value objectivec.IObject)
	CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */
	UserInfo() objc.IObject /* cross-framework: NSString */
	SetUserInfo(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchShareMetadataOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchShareMetadataOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchShareMetadataOperationClass) Alloc() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchShareMetadataOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchShareMetadataOperation */

// Creates an operation for fetching the metadata for the specified shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/init(shareURLs:)
func NewCKFetchShareMetadataOperationWithShareURLs(shareURLs []foundation.URL) CKFetchShareMetadataOperation {
	instance := getCKFetchShareMetadataOperationClass().Alloc()
	rv := objc.Send[CKFetchShareMetadataOperation](instance.ID, objc.Sel("initWithShareURLs:"), shareURLs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchShareMetadataOperationWithShareURLs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchShareMetadataOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchShareMetadataOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchShareMetadataOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchShareMetadataOperation */

// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/fetchShareMetadataCompletionBlock
func (c_ CKFetchShareMetadataOperation) FetchShareMetadataCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareMetadataCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchShareMetadataCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/fetchShareMetadataCompletionBlock
func (c_ CKFetchShareMetadataOperation) SetFetchShareMetadataCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareMetadataCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchShareMetadataCompletionBlock */


// The closure to execute as the operation fetches individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/perShareMetadataBlock
func (c_ CKFetchShareMetadataOperation) PerShareMetadataBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareMetadataBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareMetadataBlock */


// The closure to execute as the operation fetches individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/perShareMetadataBlock
func (c_ CKFetchShareMetadataOperation) SetPerShareMetadataBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareMetadataBlock:"), value)
}/* debug [instance_properties/setter]: perShareMetadataBlock */


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/rootRecordDesiredKeys-7gvr5
func (c_ CKFetchShareMetadataOperation) RootRecordDesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("rootRecordDesiredKeys"))
	return rv
}/* debug [instance_properties/getter]: rootRecordDesiredKeys */


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/rootRecordDesiredKeys-7gvr5
func (c_ CKFetchShareMetadataOperation) SetRootRecordDesiredKeys(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootRecordDesiredKeys:"), nsArray)
}/* debug [instance_properties/setter]: rootRecordDesiredKeys */


// The URLs of the shares to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shareURLs
func (c_ CKFetchShareMetadataOperation) ShareURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](c_.ID, objc.Sel("shareURLs"))
	return rv
}/* debug [instance_properties/getter]: shareURLs */


// The URLs of the shares to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shareURLs
func (c_ CKFetchShareMetadataOperation) SetShareURLs(value []foundation.URL) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareURLs:"), nsArray)
}/* debug [instance_properties/setter]: shareURLs */


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shouldFetchRootRecord
func (c_ CKFetchShareMetadataOperation) ShouldFetchRootRecord() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldFetchRootRecord"))
	return rv
}/* debug [instance_properties/getter]: shouldFetchRootRecord */


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shouldFetchRootRecord
func (c_ CKFetchShareMetadataOperation) SetShouldFetchRootRecord(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldFetchRootRecord:"), value)
}/* debug [instance_properties/setter]: shouldFetchRootRecord */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) FetchShareMetadataResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchShareMetadataResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchShareMetadataResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) SetFetchShareMetadataResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareMetadataResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchShareMetadataResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/persharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) PerShareMetadataResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perShareMetadataResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareMetadataResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/persharemetadataresultblock
func (c_ CKFetchShareMetadataOperation) SetPerShareMetadataResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareMetadataResultBlock:"), value)
}/* debug [instance_properties/setter]: perShareMetadataResultBlock */


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKFetchShareMetadataOperation) CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}/* debug [instance_properties/getter]: CKPartialErrorsByItemIDKey */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareMetadataOperation) UserInfo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareMetadataOperation) SetUserInfo(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchShareMetadataOperation */


