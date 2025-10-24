// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKAcceptSharesOperation */


/* debug [class_header]: Header for CKAcceptSharesOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKAcceptSharesOperation */
// An interface definition for the [CKAcceptSharesOperation] class.
type ICKAcceptSharesOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKAcceptSharesOperation */
	// properties:
	AcceptSharesCompletionBlock() unsafe.Pointer
	SetAcceptSharesCompletionBlock(value unsafe.Pointer)
	PerShareCompletionBlock() unsafe.Pointer
	SetPerShareCompletionBlock(value unsafe.Pointer)
	ShareMetadatas() []CKShareMetadata
	SetShareMetadatas(value []CKShareMetadata)
	AcceptSharesResultBlock() objectivec.IObject
	SetAcceptSharesResultBlock(value objectivec.IObject)
	PerShareResultBlock() objectivec.IObject
	SetPerShareResultBlock(value objectivec.IObject)
	CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */)
	Url() foundation.URL
	SetUrl(value foundation.URL)
	UserInfo() objc.IObject /* cross-framework: NSString */
	SetUserInfo(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKAcceptSharesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKAcceptSharesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKAcceptSharesOperationClass) Alloc() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKAcceptSharesOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKAcceptSharesOperation */

// Creates an operation for accepting the specified shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/init(shareMetadatas:)
func NewCKAcceptSharesOperationWithShareMetadatas(shareMetadatas []CKShareMetadata) CKAcceptSharesOperation {
	instance := getCKAcceptSharesOperationClass().Alloc()
	rv := objc.Send[CKAcceptSharesOperation](instance.ID, objc.Sel("initWithShareMetadatas:"), shareMetadatas)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKAcceptSharesOperationWithShareMetadatas */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKAcceptSharesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKAcceptSharesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKAcceptSharesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKAcceptSharesOperation */

// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
func (c_ CKAcceptSharesOperation) AcceptSharesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acceptSharesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: acceptSharesCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
func (c_ CKAcceptSharesOperation) SetAcceptSharesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcceptSharesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: acceptSharesCompletionBlock */


// The block to execute as CloudKit processes individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/perShareCompletionBlock
func (c_ CKAcceptSharesOperation) PerShareCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareCompletionBlock */


// The block to execute as CloudKit processes individual shares.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/perShareCompletionBlock
func (c_ CKAcceptSharesOperation) SetPerShareCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perShareCompletionBlock */


// The share metadatas to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/shareMetadatas
func (c_ CKAcceptSharesOperation) ShareMetadatas() []CKShareMetadata {
	rv := objc.Send[[]CKShareMetadata](c_.ID, objc.Sel("shareMetadatas"))
	return rv
}/* debug [instance_properties/getter]: shareMetadatas */


// The share metadatas to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/shareMetadatas
func (c_ CKAcceptSharesOperation) SetShareMetadatas(value []CKShareMetadata) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareMetadatas:"), nsArray)
}/* debug [instance_properties/setter]: shareMetadatas */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharesresultblock
func (c_ CKAcceptSharesOperation) AcceptSharesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("acceptSharesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: acceptSharesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharesresultblock
func (c_ CKAcceptSharesOperation) SetAcceptSharesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcceptSharesResultBlock:"), value)
}/* debug [instance_properties/setter]: acceptSharesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/pershareresultblock
func (c_ CKAcceptSharesOperation) PerShareResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perShareResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/pershareresultblock
func (c_ CKAcceptSharesOperation) SetPerShareResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareResultBlock:"), value)
}/* debug [instance_properties/setter]: perShareResultBlock */


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKAcceptSharesOperation) CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}/* debug [instance_properties/getter]: CKPartialErrorsByItemIDKey */


// The ID of the share’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/metadata/containeridentifier
func (c_ CKAcceptSharesOperation) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The ID of the share’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/metadata/containeridentifier
func (c_ CKAcceptSharesOperation) SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}/* debug [instance_properties/setter]: containerIdentifier */


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKAcceptSharesOperation) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKAcceptSharesOperation) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKAcceptSharesOperation) UserInfo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKAcceptSharesOperation) SetUserInfo(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKAcceptSharesOperation */


