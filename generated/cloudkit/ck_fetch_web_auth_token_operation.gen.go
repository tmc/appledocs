// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchWebAuthTokenOperation */


/* debug [class_header]: Header for CKFetchWebAuthTokenOperation */
// The class instance for the [CKFetchWebAuthTokenOperation] class.
var (
	CKFetchWebAuthTokenOperationClass     _CKFetchWebAuthTokenOperationClass
	CKFetchWebAuthTokenOperationClassOnce sync.Once
)

func getCKFetchWebAuthTokenOperationClass() _CKFetchWebAuthTokenOperationClass {
	CKFetchWebAuthTokenOperationClassOnce.Do(func() {
		CKFetchWebAuthTokenOperationClass = _CKFetchWebAuthTokenOperationClass{objc.GetClass("CKFetchWebAuthTokenOperation")}
	})
	return CKFetchWebAuthTokenOperationClass
}

type _CKFetchWebAuthTokenOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchWebAuthTokenOperation */
// An interface definition for the [CKFetchWebAuthTokenOperation] class.
type ICKFetchWebAuthTokenOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchWebAuthTokenOperation */
	// properties:
	APIToken() objc.IObject /* cross-framework: NSString */
	SetAPIToken(value objc.IObject /* cross-framework: NSString */)
	FetchWebAuthTokenCompletionBlock() unsafe.Pointer
	SetFetchWebAuthTokenCompletionBlock(value unsafe.Pointer)
	FetchWebAuthTokenResultBlock() objectivec.IObject
	SetFetchWebAuthTokenResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
	QualityOfService() objectivec.IObject
	SetQualityOfService(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchWebAuthTokenOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchWebAuthTokenOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchWebAuthTokenOperationClass) Alloc() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKFetchWebAuthTokenOperationClass) New() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchWebAuthTokenOperation) Init() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchWebAuthTokenOperation) Autorelease() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchWebAuthTokenOperation creates a new CKFetchWebAuthTokenOperation instance.
func NewCKFetchWebAuthTokenOperation() CKFetchWebAuthTokenOperation {
	return getCKFetchWebAuthTokenOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchWebAuthTokenOperation */
// An operation that creates an authentication token for use with CloudKit web services.
//
// CloudKit web services provides an HTTP interface to fetch, create, update, and delete records, zones, and subscriptions. Each request you send requires an API token, which you configure in . You must create an API token for each container in each environment. If you want to send a request to an endpoint that requires an authenticated user, use this operation to fetch an authentication token. Append the authentication token, along with the API token, to the endpoint’s URL. That request then acts on behalf of the current user. Authentication tokens are short-lived and expire after a single use. For an example of using a web authentication token with a CloudKit web service, see . This operation executes the handlers you provide on an internal queue it manages. Your handlers must be capable of executing on a background queue. Tasks that need access to the main queue must redirect as appropriate. The operation calls after it executes to provide the fetched token. Use the completion handler to perform housekeeping tasks for the operation. It should also manage any failures, whether due to an error or an explicit cancellation. CloudKit operations have a default QoS of . Operations with this service level are discretionary. The system schedules their execution at an optimal time according to battery level and network conditions, among other factors. Use the property to set a more appropriate QoS for the operation. The following example shows how to create the operation, configure its callbacks, and execute it in the user’s private database:


// An operation that creates an authentication token for use with CloudKit web services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation
type CKFetchWebAuthTokenOperation struct {
	CKDatabaseOperation
}

// CKFetchWebAuthTokenOperationFrom constructs a [CKFetchWebAuthTokenOperation] from an unsafe.Pointer.
//
// An operation that creates an authentication token for use with CloudKit web services.
func CKFetchWebAuthTokenOperationFrom(ptr unsafe.Pointer) CKFetchWebAuthTokenOperation {
	return CKFetchWebAuthTokenOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchWebAuthTokenOperation */

// Creates a fetch operation for the specified API token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/init(apiToken:)
func NewCKFetchWebAuthTokenOperationWithAPIToken(APIToken objc.IObject /* cross-framework: NSString */) CKFetchWebAuthTokenOperation {
	instance := getCKFetchWebAuthTokenOperationClass().Alloc()
	rv := objc.Send[CKFetchWebAuthTokenOperation](instance.ID, objc.Sel("initWithAPIToken:"), APIToken)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchWebAuthTokenOperationWithAPIToken */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchWebAuthTokenOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchWebAuthTokenOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchWebAuthTokenOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchWebAuthTokenOperation */

// The API token that allows access to an app’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/apiToken
func (c_ CKFetchWebAuthTokenOperation) APIToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("APIToken"))
	return rv
}/* debug [instance_properties/getter]: APIToken */


// The API token that allows access to an app’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/apiToken
func (c_ CKFetchWebAuthTokenOperation) SetAPIToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAPIToken:"), value)
}/* debug [instance_properties/setter]: APIToken */


// The block to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/fetchWebAuthTokenCompletionBlock
func (c_ CKFetchWebAuthTokenOperation) FetchWebAuthTokenCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchWebAuthTokenCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchWebAuthTokenCompletionBlock */


// The block to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/fetchWebAuthTokenCompletionBlock
func (c_ CKFetchWebAuthTokenOperation) SetFetchWebAuthTokenCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchWebAuthTokenCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchWebAuthTokenCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/fetchwebauthtokenresultblock
func (c_ CKFetchWebAuthTokenOperation) FetchWebAuthTokenResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchWebAuthTokenResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchWebAuthTokenResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/fetchwebauthtokenresultblock
func (c_ CKFetchWebAuthTokenOperation) SetFetchWebAuthTokenResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchWebAuthTokenResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchWebAuthTokenResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchWebAuthTokenOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchWebAuthTokenOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKFetchWebAuthTokenOperation) QualityOfService() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKFetchWebAuthTokenOperation) SetQualityOfService(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchWebAuthTokenOperation */


