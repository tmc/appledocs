// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ODNode] class.
type IODNode interface {
	objectivec.IObject
	AccountPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary
	AddAccountPolicyToCategoryError(policy objectivec.IObject, category ODPolicyCategoryType, error_ unsafe.Pointer) bool
	CreateRecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName appkit.string, inAttributes objectivec.IObject, outError unsafe.Pointer) ODRecord
	CustomCallSendDataError(inCustomCode int, inSendData foundation.IData, outError unsafe.Pointer) foundation.Data
	CustomFunctionPayloadError(function appkit.string, payload objectivec.IObject, error_ unsafe.Pointer) objc.ID
	NodeDetailsForKeysError(inKeys objectivec.IObject, outError unsafe.Pointer) foundation.Dictionary
	PasswordContentCheckForRecordNameError(password appkit.string, recordName appkit.string, error_ unsafe.Pointer) bool
	PoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary
	RecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName appkit.string, inAttributes objectivec.IObject, outError unsafe.Pointer) ODRecord
	RemoveAccountPolicyFromCategoryError(policy objectivec.IObject, category ODPolicyCategoryType, error_ unsafe.Pointer) bool
	RemovePolicyError(policy ODPolicyType, error_ unsafe.Pointer) bool
	SetAccountPoliciesError(policies objectivec.IObject, error_ unsafe.Pointer) bool
	SetCredentialsUsingKerberosCacheError(inCacheName appkit.string, outError unsafe.Pointer) bool
	SetCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType ODAuthenticationType, inItems objectivec.IObject, outItems objectivec.IObject, outContext objectivec.IObject, outError unsafe.Pointer) bool
	SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType unsafe.Pointer, inRecordName appkit.string, inPassword appkit.string, outError unsafe.Pointer) bool
	SetPoliciesError(policies objectivec.IObject, error_ unsafe.Pointer) bool
	SetPolicyValueError(policy ODPolicyType, value objectivec.IObject, error_ unsafe.Pointer) bool
	SubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array
	SupportedAttributesForRecordTypeError(inRecordType unsafe.Pointer, outError unsafe.Pointer) foundation.Array
	SupportedPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary
	SupportedRecordTypesAndReturnError(outError unsafe.Pointer) foundation.Array
	UnreachableSubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array
}

// An object serves as a Cocoa wrapper for an Open Directory node.
//
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

// Alloc allocates a new instance without initialization.
func (oc _ODNodeClass) Alloc() ODNode {
	rv := objc.Send[ODNode](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a node object with a specified session and name.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:name:)
func NewODNodeWithSessionNameError(inSession IODSession, inName appkit.string, outError unsafe.Pointer) ODNode {
	instance := getODNodeClass().Alloc()
	rv := objc.Send[ODNode](instance.ID, objc.Sel("initWithSession:name:error:"), inSession, inName, outError)
	rv.Autorelease()
	return rv
}



// Creates a node object with a specified session and type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:type:)
func NewODNodeWithSessionTypeError(inSession IODSession, inType ODNodeType, outError unsafe.Pointer) ODNode {
	instance := getODNodeClass().Alloc()
	rv := objc.Send[ODNode](instance.ID, objc.Sel("initWithSession:type:error:"), inSession, inType, outError)
	rv.Autorelease()
	return rv
}


// Returns an autoreleased node object with a specified session and name.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeWithSession:name:error:
func (oc _ODNodeClass) NodeWithSessionNameError(inSession IODSession, inName appkit.string, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("nodeWithSession:name:error:"), inSession, inName, outError)
	return rv
}

// Returns an autoreleased node object with a specified session and type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeWithSession:type:error:
func (oc _ODNodeClass) NodeWithSessionTypeError(inSession IODSession, inType ODNodeType, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("nodeWithSession:type:error:"), inSession, inType, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/accountPolicies()
func (o_ ODNode) AccountPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("accountPoliciesAndReturnError:"), error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/addAccountPolicy(_:toCategory:)
func (o_ ODNode) AddAccountPolicyToCategoryError(policy objectivec.IObject, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addAccountPolicy:toCategory:error:"), policy, category, error_)
	return rv
}

// Creates a record in a specified node with specified properties.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/createRecord(withRecordType:name:attributes:)
func (o_ ODNode) CreateRecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName appkit.string, inAttributes objectivec.IObject, outError unsafe.Pointer) ODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("createRecordWithRecordType:name:attributes:error:"), inRecordType, inRecordName, inAttributes, outError)
	return rv
}

// Returns the result of a custom call to the node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/customCall(_:send:)
func (o_ ODNode) CustomCallSendDataError(inCustomCode int, inSendData foundation.IData, outError unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](o_.ID, objc.Sel("customCall:sendData:error:"), inCustomCode, inSendData, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/customFunction(_:payload:)
func (o_ ODNode) CustomFunctionPayloadError(function appkit.string, payload objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("customFunction:payload:error:"), function, payload, error_)
	return rv
}

// Returns a dictionary containing details about a node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeDetails(forKeys:)
func (o_ ODNode) NodeDetailsForKeysError(inKeys objectivec.IObject, outError unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("nodeDetailsForKeys:error:"), inKeys, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/passwordContentCheck(_:forRecordName:)
func (o_ ODNode) PasswordContentCheckForRecordNameError(password appkit.string, recordName appkit.string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("passwordContentCheck:forRecordName:error:"), password, recordName, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/policies()
func (o_ ODNode) PoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("policiesAndReturnError:"), error_)
	return rv
}

// Returns a record from the node with a specified type and name.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/record(withRecordType:name:attributes:)
func (o_ ODNode) RecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName appkit.string, inAttributes objectivec.IObject, outError unsafe.Pointer) ODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("recordWithRecordType:name:attributes:error:"), inRecordType, inRecordName, inAttributes, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/removeAccountPolicy(_:fromCategory:)
func (o_ ODNode) RemoveAccountPolicyFromCategoryError(policy objectivec.IObject, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeAccountPolicy:fromCategory:error:"), policy, category, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/removePolicy(_:)
func (o_ ODNode) RemovePolicyError(policy ODPolicyType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removePolicy:error:"), policy, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setAccountPolicies(_:)
func (o_ ODNode) SetAccountPoliciesError(policies objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setAccountPolicies:error:"), policies, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsUsingKerberosCache:error:
func (o_ ODNode) SetCredentialsUsingKerberosCacheError(inCacheName appkit.string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setCredentialsUsingKerberosCache:error:"), inCacheName, outError)
	return rv
}

// Sets the credentials for interaction with the node using other types of authentication available to Open Directory.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsWithRecordType(_:authenticationType:authenticationItems:continueItems:context:)
func (o_ ODNode) SetCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType ODAuthenticationType, inItems objectivec.IObject, outItems objectivec.IObject, outContext objectivec.IObject, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setCredentialsWithRecordType:authenticationType:authenticationItems:continueItems:context:error:"), inRecordType, inType, inItems, outItems, outContext, outError)
	return rv
}

// Sets credentials for interacting with the node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsWithRecordType(_:recordName:password:)
func (o_ ODNode) SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType unsafe.Pointer, inRecordName appkit.string, inPassword appkit.string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setCredentialsWithRecordType:recordName:password:error:"), inRecordType, inRecordName, inPassword, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setPolicies(_:)
func (o_ ODNode) SetPoliciesError(policies objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setPolicies:error:"), policies, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/setPolicy(_:value:)
func (o_ ODNode) SetPolicyValueError(policy ODPolicyType, value objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setPolicy:value:error:"), policy, value, error_)
	return rv
}

// Returns the names of subnodes for the node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/subnodeNames()
func (o_ ODNode) SubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("subnodeNamesAndReturnError:"), outError)
	return rv
}

// Returns an array of attribute types supported by the node’s records.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedAttributes(forRecordType:)
func (o_ ODNode) SupportedAttributesForRecordTypeError(inRecordType unsafe.Pointer, outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("supportedAttributesForRecordType:error:"), inRecordType, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedPolicies()
func (o_ ODNode) SupportedPoliciesAndReturnError(error_ unsafe.Pointer) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("supportedPoliciesAndReturnError:"), error_)
	return rv
}

// Returns an array of the record types supported by the node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedRecordTypes()
func (o_ ODNode) SupportedRecordTypesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("supportedRecordTypesAndReturnError:"), outError)
	return rv
}

// Returns an array of the subnodes of a given node that are currently unreachable.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/unreachableSubnodeNames()
func (o_ ODNode) UnreachableSubnodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("unreachableSubnodeNamesAndReturnError:"), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/configuration
func (o_ ODNode) Configuration() ODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}

// The node’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeName
func (o_ ODNode) NodeName() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("nodeName"))
	return rv
}


