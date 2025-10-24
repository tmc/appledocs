// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODNode */


/* debug [class_header]: Header for ODNode */
// The class instance for the [ODNode] class.
var (
	ODNodeClass     _ODNodeClass
	ODNodeClassOnce sync.Once
)

func getODNodeClass() _ODNodeClass {
	ODNodeClassOnce.Do(func() {
		ODNodeClass = _ODNodeClass{objc.GetClass("ODNode")}
	})
	return ODNodeClass
}

type _ODNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODNode */
// An interface definition for the [ODNode] class.
type IODNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODNode */
	// properties:
	Configuration() IODConfiguration
	NodeName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODNode */
	// methods:
	AccountPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary
	AddAccountPolicyToCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool
	CreateRecordWithRecordTypeNameAttributesError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inAttributes objc.IObject /* cross-framework: NSDictionary */, outError unsafe.Pointer) IODRecord
	CustomCallSendDataError(inCustomCode int, inSendData objc.IObject /* cross-framework: NSData */, outError unsafe.Pointer) foundation.Data
	CustomFunctionPayloadError(function objc.IObject /* cross-framework: NSString */, payload objc.IObject, error_ unsafe.Pointer) objc.ID
	NodeDetailsForKeysError(inKeys objc.IObject /* cross-framework: NSArray */, outError unsafe.Pointer) foundation.Dictionary
	PasswordContentCheckForRecordNameError(password objc.IObject /* cross-framework: NSString */, recordName objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	RecordWithRecordTypeNameAttributesError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inAttributes objc.IObject, outError unsafe.Pointer) IODRecord
	RemoveAccountPolicyFromCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool
	SetAccountPoliciesError(policies objc.IObject /* cross-framework: NSDictionary */, error_ unsafe.Pointer) bool
	SetCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType ODRecordType /* typedef */, inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool
	SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	SubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array
	SupportedAttributesForRecordTypeError(inRecordType ODRecordType /* typedef */, outError unsafe.Pointer) foundation.Array
	SupportedRecordTypesAndReturnError(outError unsafe.Pointer) foundation.Array
	UnreachableSubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODNode */
// Alloc allocates a new instance without initialization.
func (oc _ODNodeClass) Alloc() ODNode {
	rv := objc.Send[ODNode](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ODNodeClass) New() ODNode {
	rv := objc.Send[ODNode](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODNode) Init() ODNode {
	rv := objc.Send[ODNode](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODNode) Autorelease() ODNode {
	rv := objc.Send[ODNode](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODNode creates a new ODNode instance.
func NewODNode() ODNode {
	return getODNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODNode */
// An object serves as a Cocoa wrapper for an Open Directory node.


// An object serves as a Cocoa wrapper for an Open Directory node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode
type ODNode struct {
	objectivec.Object
}

// ODNodeFrom constructs a [ODNode] from an unsafe.Pointer.
//
// An object serves as a Cocoa wrapper for an Open Directory node.
func ODNodeFrom(ptr unsafe.Pointer) ODNode {
	return ODNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODNode */

// Creates a node object with a specified session and name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:name:)
func NewODNodeWithSessionNameError(inSession IODSession, inName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) ODNode {
	instance := getODNodeClass().Alloc()
	rv := objc.Send[ODNode](instance.ID, objc.Sel("initWithSession:name:error:"), inSession, inName, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewODNodeWithSessionNameError */


// Creates a node object with a specified session and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:type:)
func NewODNodeWithSessionTypeError(inSession IODSession, inType ODNodeType /* typedef */, outError unsafe.Pointer) ODNode {
	instance := getODNodeClass().Alloc()
	rv := objc.Send[ODNode](instance.ID, objc.Sel("initWithSession:type:error:"), inSession, inType, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewODNodeWithSessionTypeError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODNode */

// Returns an autoreleased node object with a specified session and name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeWithSession:name:error:
func (oc _ODNodeClass) NodeWithSessionNameError(inSession IODSession, inName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("nodeWithSession:name:error:"), inSession, inName, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSessionNameError) */


// Returns an autoreleased node object with a specified session and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeWithSession:type:error:
func (oc _ODNodeClass) NodeWithSessionTypeError(inSession IODSession, inType ODNodeType /* typedef */, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("nodeWithSession:type:error:"), inSession, inType, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSessionTypeError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/accountPolicies()
func (o_ ODNode) AccountPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("accountPoliciesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: AccountPoliciesAndReturnError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/addAccountPolicy(_:toCategory:)
func (o_ ODNode) AddAccountPolicyToCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addAccountPolicy:toCategory:error:"), policy, category, error_)
	return rv
}/* debug [instance_methods/method]: AddAccountPolicyToCategoryError */


// Creates a record in a specified node with specified properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/createRecord(withRecordType:name:attributes:)
func (o_ ODNode) CreateRecordWithRecordTypeNameAttributesError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inAttributes objc.IObject /* cross-framework: NSDictionary */, outError unsafe.Pointer) IODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("createRecordWithRecordType:name:attributes:error:"), inRecordType, inRecordName, inAttributes, outError)
	return rv
}/* debug [instance_methods/method]: CreateRecordWithRecordTypeNameAttributesError */


// Returns the result of a custom call to the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/customCall(_:send:)
func (o_ ODNode) CustomCallSendDataError(inCustomCode int, inSendData objc.IObject /* cross-framework: NSData */, outError unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](o_.ID, objc.Sel("customCall:sendData:error:"), inCustomCode, inSendData, outError)
	return rv
}/* debug [instance_methods/method]: CustomCallSendDataError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/customFunction(_:payload:)
func (o_ ODNode) CustomFunctionPayloadError(function objc.IObject /* cross-framework: NSString */, payload objc.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("customFunction:payload:error:"), function, payload, error_)
	return rv
}/* debug [instance_methods/method]: CustomFunctionPayloadError */


// Returns a dictionary containing details about a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeDetails(forKeys:)
func (o_ ODNode) NodeDetailsForKeysError(inKeys objc.IObject /* cross-framework: NSArray */, outError unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("nodeDetailsForKeys:error:"), inKeys, outError)
	return rv
}/* debug [instance_methods/method]: NodeDetailsForKeysError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/passwordContentCheck(_:forRecordName:)
func (o_ ODNode) PasswordContentCheckForRecordNameError(password objc.IObject /* cross-framework: NSString */, recordName objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("passwordContentCheck:forRecordName:error:"), password, recordName, error_)
	return rv
}/* debug [instance_methods/method]: PasswordContentCheckForRecordNameError */


// Returns a record from the node with a specified type and name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/record(withRecordType:name:attributes:)
func (o_ ODNode) RecordWithRecordTypeNameAttributesError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inAttributes objc.IObject, outError unsafe.Pointer) IODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("recordWithRecordType:name:attributes:error:"), inRecordType, inRecordName, inAttributes, outError)
	return rv
}/* debug [instance_methods/method]: RecordWithRecordTypeNameAttributesError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/removeAccountPolicy(_:fromCategory:)
func (o_ ODNode) RemoveAccountPolicyFromCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeAccountPolicy:fromCategory:error:"), policy, category, error_)
	return rv
}/* debug [instance_methods/method]: RemoveAccountPolicyFromCategoryError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setAccountPolicies(_:)
func (o_ ODNode) SetAccountPoliciesError(policies objc.IObject /* cross-framework: NSDictionary */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setAccountPolicies:error:"), policies, error_)
	return rv
}/* debug [instance_methods/method]: SetAccountPoliciesError */


// Sets the credentials for interaction with the node using other types of authentication available to Open Directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsWithRecordType(_:authenticationType:authenticationItems:continueItems:context:)
func (o_ ODNode) SetCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType ODRecordType /* typedef */, inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setCredentialsWithRecordType:authenticationType:authenticationItems:continueItems:context:error:"), inRecordType, inType, inItems, outItems, outContext, outError)
	return rv
}/* debug [instance_methods/method]: SetCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError */


// Sets credentials for interacting with the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsWithRecordType(_:recordName:password:)
func (o_ ODNode) SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType ODRecordType /* typedef */, inRecordName objc.IObject /* cross-framework: NSString */, inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setCredentialsWithRecordType:recordName:password:error:"), inRecordType, inRecordName, inPassword, outError)
	return rv
}/* debug [instance_methods/method]: SetCredentialsWithRecordTypeRecordNamePasswordError */


// Returns the names of subnodes for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/subnodeNames()
func (o_ ODNode) SubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("subnodeNamesAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: SubnodeNamesAndReturnError */


// Returns an array of attribute types supported by the node’s records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedAttributes(forRecordType:)
func (o_ ODNode) SupportedAttributesForRecordTypeError(inRecordType ODRecordType /* typedef */, outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("supportedAttributesForRecordType:error:"), inRecordType, outError)
	return rv
}/* debug [instance_methods/method]: SupportedAttributesForRecordTypeError */


// Returns an array of the record types supported by the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedRecordTypes()
func (o_ ODNode) SupportedRecordTypesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("supportedRecordTypesAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: SupportedRecordTypesAndReturnError */


// Returns an array of the subnodes of a given node that are currently unreachable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/unreachableSubnodeNames()
func (o_ ODNode) UnreachableSubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("unreachableSubnodeNamesAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: UnreachableSubnodeNamesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/configuration
func (o_ ODNode) Configuration() IODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The node’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeName
func (o_ ODNode) NodeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("nodeName"))
	return rv
}/* debug [instance_properties/getter]: nodeName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODNode */


