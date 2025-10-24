// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

/* debug [functions.gen.go]: Generating 76 functions for OpenDirectory */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// OpenDirectory Functions (76 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_ODContextGetTypeID func() TypeID
	_ODNodeAddAccountPolicy func(ODNodeRef, DictionaryRef, ODPolicyCategoryType, unsafe.Pointer) bool
	_ODNodeCopyAccountPolicies func(ODNodeRef, unsafe.Pointer) DictionaryRef
	_ODNodeCopyDetails func(ODNodeRef, ArrayRef, unsafe.Pointer) DictionaryRef
	_ODNodeCopyPolicies func(ODNodeRef, unsafe.Pointer) DictionaryRef
	_ODNodeCopyRecord func(ODNodeRef, ODRecordType, StringRef, TypeRef, unsafe.Pointer) ODRecordRef
	_ODNodeCopySubnodeNames func(ODNodeRef, unsafe.Pointer) ArrayRef
	_ODNodeCopySupportedAttributes func(ODNodeRef, ODRecordType, unsafe.Pointer) ArrayRef
	_ODNodeCopySupportedPolicies func(ODNodeRef, unsafe.Pointer) DictionaryRef
	_ODNodeCopySupportedRecordTypes func(ODNodeRef, unsafe.Pointer) ArrayRef
	_ODNodeCopyUnreachableSubnodeNames func(ODNodeRef, unsafe.Pointer) ArrayRef
	_ODNodeCreateCopy func(AllocatorRef, ODNodeRef, unsafe.Pointer) ODNodeRef
	_ODNodeCreateRecord func(ODNodeRef, ODRecordType, StringRef, DictionaryRef, unsafe.Pointer) ODRecordRef
	_ODNodeCreateWithName func(AllocatorRef, ODSessionRef, StringRef, unsafe.Pointer) ODNodeRef
	_ODNodeCreateWithNodeType func(AllocatorRef, ODSessionRef, ODNodeType, unsafe.Pointer) ODNodeRef
	_ODNodeCustomCall func(ODNodeRef, Index, DataRef, unsafe.Pointer) DataRef
	_ODNodeCustomFunction func(ODNodeRef, StringRef, TypeRef, unsafe.Pointer) TypeRef
	_ODNodeGetName func(ODNodeRef) StringRef
	_ODNodeGetTypeID func() TypeID
	_ODNodePasswordContentCheck func(ODNodeRef, StringRef, StringRef, unsafe.Pointer) bool
	_ODNodeRemoveAccountPolicy func(ODNodeRef, DictionaryRef, ODPolicyCategoryType, unsafe.Pointer) bool
	_ODNodeRemovePolicy func(ODNodeRef, ODPolicyType, unsafe.Pointer) bool
	_ODNodeSetAccountPolicies func(ODNodeRef, DictionaryRef, unsafe.Pointer) bool
	_ODNodeSetCredentials func(ODNodeRef, ODRecordType, StringRef, StringRef, unsafe.Pointer) bool
	_ODNodeSetCredentialsExtended func(ODNodeRef, ODRecordType, ODAuthenticationType, ArrayRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ODNodeSetCredentialsUsingKerberosCache func(ODNodeRef, StringRef, unsafe.Pointer) bool
	_ODNodeSetPolicies func(ODNodeRef, DictionaryRef, unsafe.Pointer) bool
	_ODNodeSetPolicy func(ODNodeRef, ODPolicyType, TypeRef, unsafe.Pointer) bool
	_ODQueryCopyResults func(ODQueryRef, bool, unsafe.Pointer) ArrayRef
	_ODQueryCreateWithNode func(AllocatorRef, ODNodeRef, TypeRef, ODAttributeType, ODMatchType, TypeRef, TypeRef, Index, unsafe.Pointer) ODQueryRef
	_ODQueryCreateWithNodeType func(AllocatorRef, ODNodeType, TypeRef, ODAttributeType, ODMatchType, TypeRef, TypeRef, Index, unsafe.Pointer) ODQueryRef
	_ODQueryGetTypeID func() TypeID
	_ODQueryScheduleWithRunLoop func(ODQueryRef, RunLoopRef, StringRef)
	_ODQuerySetCallback func(ODQueryRef, ODQueryCallback, unsafe.Pointer)
	_ODQuerySetDispatchQueue func(ODQueryRef, unsafe.Pointer)
	_ODQuerySynchronize func(ODQueryRef)
	_ODQueryUnscheduleFromRunLoop func(ODQueryRef, RunLoopRef, StringRef)
	_ODRecordAddAccountPolicy func(ODRecordRef, DictionaryRef, ODPolicyCategoryType, unsafe.Pointer) bool
	_ODRecordAddMember func(ODRecordRef, ODRecordRef, unsafe.Pointer) bool
	_ODRecordAddValue func(ODRecordRef, ODAttributeType, TypeRef, unsafe.Pointer) bool
	_ODRecordAuthenticationAllowed func(ODRecordRef, unsafe.Pointer) bool
	_ODRecordChangePassword func(ODRecordRef, StringRef, StringRef, unsafe.Pointer) bool
	_ODRecordContainsMember func(ODRecordRef, ODRecordRef, unsafe.Pointer) bool
	_ODRecordCopyAccountPolicies func(ODRecordRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopyDetails func(ODRecordRef, ArrayRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopyEffectivePolicies func(ODRecordRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopyPasswordPolicy func(AllocatorRef, ODRecordRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopyPolicies func(ODRecordRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopySupportedPolicies func(ODRecordRef, unsafe.Pointer) DictionaryRef
	_ODRecordCopyValues func(ODRecordRef, ODAttributeType, unsafe.Pointer) ArrayRef
	_ODRecordDelete func(ODRecordRef, unsafe.Pointer) bool
	_ODRecordGetRecordName func(ODRecordRef) StringRef
	_ODRecordGetRecordType func(ODRecordRef) StringRef
	_ODRecordGetTypeID func() TypeID
	_ODRecordPasswordChangeAllowed func(ODRecordRef, StringRef, unsafe.Pointer) bool
	_ODRecordRemoveAccountPolicy func(ODRecordRef, DictionaryRef, ODPolicyCategoryType, unsafe.Pointer) bool
	_ODRecordRemoveMember func(ODRecordRef, ODRecordRef, unsafe.Pointer) bool
	_ODRecordRemovePolicy func(ODRecordRef, ODPolicyType, unsafe.Pointer) bool
	_ODRecordRemoveValue func(ODRecordRef, ODAttributeType, TypeRef, unsafe.Pointer) bool
	_ODRecordSecondsUntilAuthenticationsExpire func(ODRecordRef) int64
	_ODRecordSecondsUntilPasswordExpires func(ODRecordRef) int64
	_ODRecordSetAccountPolicies func(ODRecordRef, DictionaryRef, unsafe.Pointer) bool
	_ODRecordSetNodeCredentials func(ODRecordRef, StringRef, StringRef, unsafe.Pointer) bool
	_ODRecordSetNodeCredentialsExtended func(ODRecordRef, ODRecordType, ODAuthenticationType, ArrayRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ODRecordSetNodeCredentialsUsingKerberosCache func(ODRecordRef, StringRef, unsafe.Pointer) bool
	_ODRecordSetPolicies func(ODRecordRef, DictionaryRef, unsafe.Pointer) bool
	_ODRecordSetPolicy func(ODRecordRef, ODPolicyType, TypeRef, unsafe.Pointer) bool
	_ODRecordSetValue func(ODRecordRef, ODAttributeType, TypeRef, unsafe.Pointer) bool
	_ODRecordSynchronize func(ODRecordRef, unsafe.Pointer) bool
	_ODRecordVerifyPassword func(ODRecordRef, StringRef, unsafe.Pointer) bool
	_ODRecordVerifyPasswordExtended func(ODRecordRef, ODAuthenticationType, ArrayRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ODRecordWillAuthenticationsExpire func(ODRecordRef, uint64) bool
	_ODRecordWillPasswordExpire func(ODRecordRef, uint64) bool
	_ODSessionCopyNodeNames func(AllocatorRef, ODSessionRef, unsafe.Pointer) ArrayRef
	_ODSessionCreate func(AllocatorRef, DictionaryRef, unsafe.Pointer) ODSessionRef
	_ODSessionGetTypeID func() TypeID
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_ODContextGetTypeID, lib, "ODContextGetTypeID")
	tryRegister(&_ODNodeAddAccountPolicy, lib, "ODNodeAddAccountPolicy")
	tryRegister(&_ODNodeCopyAccountPolicies, lib, "ODNodeCopyAccountPolicies")
	tryRegister(&_ODNodeCopyDetails, lib, "ODNodeCopyDetails")
	tryRegister(&_ODNodeCopyPolicies, lib, "ODNodeCopyPolicies")
	tryRegister(&_ODNodeCopyRecord, lib, "ODNodeCopyRecord")
	tryRegister(&_ODNodeCopySubnodeNames, lib, "ODNodeCopySubnodeNames")
	tryRegister(&_ODNodeCopySupportedAttributes, lib, "ODNodeCopySupportedAttributes")
	tryRegister(&_ODNodeCopySupportedPolicies, lib, "ODNodeCopySupportedPolicies")
	tryRegister(&_ODNodeCopySupportedRecordTypes, lib, "ODNodeCopySupportedRecordTypes")
	tryRegister(&_ODNodeCopyUnreachableSubnodeNames, lib, "ODNodeCopyUnreachableSubnodeNames")
	tryRegister(&_ODNodeCreateCopy, lib, "ODNodeCreateCopy")
	tryRegister(&_ODNodeCreateRecord, lib, "ODNodeCreateRecord")
	tryRegister(&_ODNodeCreateWithName, lib, "ODNodeCreateWithName")
	tryRegister(&_ODNodeCreateWithNodeType, lib, "ODNodeCreateWithNodeType")
	tryRegister(&_ODNodeCustomCall, lib, "ODNodeCustomCall")
	tryRegister(&_ODNodeCustomFunction, lib, "ODNodeCustomFunction")
	tryRegister(&_ODNodeGetName, lib, "ODNodeGetName")
	tryRegister(&_ODNodeGetTypeID, lib, "ODNodeGetTypeID")
	tryRegister(&_ODNodePasswordContentCheck, lib, "ODNodePasswordContentCheck")
	tryRegister(&_ODNodeRemoveAccountPolicy, lib, "ODNodeRemoveAccountPolicy")
	tryRegister(&_ODNodeRemovePolicy, lib, "ODNodeRemovePolicy")
	tryRegister(&_ODNodeSetAccountPolicies, lib, "ODNodeSetAccountPolicies")
	tryRegister(&_ODNodeSetCredentials, lib, "ODNodeSetCredentials")
	tryRegister(&_ODNodeSetCredentialsExtended, lib, "ODNodeSetCredentialsExtended")
	tryRegister(&_ODNodeSetCredentialsUsingKerberosCache, lib, "ODNodeSetCredentialsUsingKerberosCache")
	tryRegister(&_ODNodeSetPolicies, lib, "ODNodeSetPolicies")
	tryRegister(&_ODNodeSetPolicy, lib, "ODNodeSetPolicy")
	tryRegister(&_ODQueryCopyResults, lib, "ODQueryCopyResults")
	tryRegister(&_ODQueryCreateWithNode, lib, "ODQueryCreateWithNode")
	tryRegister(&_ODQueryCreateWithNodeType, lib, "ODQueryCreateWithNodeType")
	tryRegister(&_ODQueryGetTypeID, lib, "ODQueryGetTypeID")
	tryRegister(&_ODQueryScheduleWithRunLoop, lib, "ODQueryScheduleWithRunLoop")
	tryRegister(&_ODQuerySetCallback, lib, "ODQuerySetCallback")
	tryRegister(&_ODQuerySetDispatchQueue, lib, "ODQuerySetDispatchQueue")
	tryRegister(&_ODQuerySynchronize, lib, "ODQuerySynchronize")
	tryRegister(&_ODQueryUnscheduleFromRunLoop, lib, "ODQueryUnscheduleFromRunLoop")
	tryRegister(&_ODRecordAddAccountPolicy, lib, "ODRecordAddAccountPolicy")
	tryRegister(&_ODRecordAddMember, lib, "ODRecordAddMember")
	tryRegister(&_ODRecordAddValue, lib, "ODRecordAddValue")
	tryRegister(&_ODRecordAuthenticationAllowed, lib, "ODRecordAuthenticationAllowed")
	tryRegister(&_ODRecordChangePassword, lib, "ODRecordChangePassword")
	tryRegister(&_ODRecordContainsMember, lib, "ODRecordContainsMember")
	tryRegister(&_ODRecordCopyAccountPolicies, lib, "ODRecordCopyAccountPolicies")
	tryRegister(&_ODRecordCopyDetails, lib, "ODRecordCopyDetails")
	tryRegister(&_ODRecordCopyEffectivePolicies, lib, "ODRecordCopyEffectivePolicies")
	tryRegister(&_ODRecordCopyPasswordPolicy, lib, "ODRecordCopyPasswordPolicy")
	tryRegister(&_ODRecordCopyPolicies, lib, "ODRecordCopyPolicies")
	tryRegister(&_ODRecordCopySupportedPolicies, lib, "ODRecordCopySupportedPolicies")
	tryRegister(&_ODRecordCopyValues, lib, "ODRecordCopyValues")
	tryRegister(&_ODRecordDelete, lib, "ODRecordDelete")
	tryRegister(&_ODRecordGetRecordName, lib, "ODRecordGetRecordName")
	tryRegister(&_ODRecordGetRecordType, lib, "ODRecordGetRecordType")
	tryRegister(&_ODRecordGetTypeID, lib, "ODRecordGetTypeID")
	tryRegister(&_ODRecordPasswordChangeAllowed, lib, "ODRecordPasswordChangeAllowed")
	tryRegister(&_ODRecordRemoveAccountPolicy, lib, "ODRecordRemoveAccountPolicy")
	tryRegister(&_ODRecordRemoveMember, lib, "ODRecordRemoveMember")
	tryRegister(&_ODRecordRemovePolicy, lib, "ODRecordRemovePolicy")
	tryRegister(&_ODRecordRemoveValue, lib, "ODRecordRemoveValue")
	tryRegister(&_ODRecordSecondsUntilAuthenticationsExpire, lib, "ODRecordSecondsUntilAuthenticationsExpire")
	tryRegister(&_ODRecordSecondsUntilPasswordExpires, lib, "ODRecordSecondsUntilPasswordExpires")
	tryRegister(&_ODRecordSetAccountPolicies, lib, "ODRecordSetAccountPolicies")
	tryRegister(&_ODRecordSetNodeCredentials, lib, "ODRecordSetNodeCredentials")
	tryRegister(&_ODRecordSetNodeCredentialsExtended, lib, "ODRecordSetNodeCredentialsExtended")
	tryRegister(&_ODRecordSetNodeCredentialsUsingKerberosCache, lib, "ODRecordSetNodeCredentialsUsingKerberosCache")
	tryRegister(&_ODRecordSetPolicies, lib, "ODRecordSetPolicies")
	tryRegister(&_ODRecordSetPolicy, lib, "ODRecordSetPolicy")
	tryRegister(&_ODRecordSetValue, lib, "ODRecordSetValue")
	tryRegister(&_ODRecordSynchronize, lib, "ODRecordSynchronize")
	tryRegister(&_ODRecordVerifyPassword, lib, "ODRecordVerifyPassword")
	tryRegister(&_ODRecordVerifyPasswordExtended, lib, "ODRecordVerifyPasswordExtended")
	tryRegister(&_ODRecordWillAuthenticationsExpire, lib, "ODRecordWillAuthenticationsExpire")
	tryRegister(&_ODRecordWillPasswordExpire, lib, "ODRecordWillPasswordExpire")
	tryRegister(&_ODSessionCopyNodeNames, lib, "ODSessionCopyNodeNames")
	tryRegister(&_ODSessionCreate, lib, "ODSessionCreate")
	tryRegister(&_ODSessionGetTypeID, lib, "ODSessionGetTypeID")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the type ID for the Open Directory context.
//
// Added in macOS .
// Returns the type ID for the Open Directory context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODContextGetTypeID()
func ODContextGetTypeID() TypeID {
	return _ODContextGetTypeID()
}/* debug [functions.gen.go/function]: ODContextGetTypeID */

// ODNodeAddAccountPolicy is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeAddAccountPolicy(_:_:_:_:)
func ODNodeAddAccountPolicy(node ODNodeRef, policy DictionaryRef, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	return _ODNodeAddAccountPolicy(node, policy, category, error_)
}/* debug [functions.gen.go/function]: ODNodeAddAccountPolicy */

// ODNodeCopyAccountPolicies is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyAccountPolicies(_:_:)
func ODNodeCopyAccountPolicies(node ODNodeRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODNodeCopyAccountPolicies(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopyAccountPolicies */

// Returns a dictionary containing details about a node.
//
// Added in macOS 10.6.
// Returns a dictionary containing details about a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyDetails(_:_:_:)
func ODNodeCopyDetails(node ODNodeRef, keys ArrayRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODNodeCopyDetails(node, keys, error_)
}/* debug [functions.gen.go/function]: ODNodeCopyDetails */

// ODNodeCopyPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyPolicies(_:_:)
func ODNodeCopyPolicies(node ODNodeRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODNodeCopyPolicies(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopyPolicies */

// Returns a reference to a record of a node.
//
// Added in macOS 10.6.
// Returns a reference to a record of a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyRecord(_:_:_:_:_:)
func ODNodeCopyRecord(node ODNodeRef, recordType ODRecordType, recordName StringRef, attributes TypeRef, error_ unsafe.Pointer) ODRecordRef {
	return _ODNodeCopyRecord(node, recordType, recordName, attributes, error_)
}/* debug [functions.gen.go/function]: ODNodeCopyRecord */

// Returns the names of subnodes for a given node.
//
// Added in macOS 10.6.
// Returns the names of subnodes for a given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySubnodeNames(_:_:)
func ODNodeCopySubnodeNames(node ODNodeRef, error_ unsafe.Pointer) ArrayRef {
	return _ODNodeCopySubnodeNames(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopySubnodeNames */

// Returns an array of attribute types supported by a given node.
//
// Added in macOS 10.6.
// Returns an array of attribute types supported by a given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedAttributes(_:_:_:)
func ODNodeCopySupportedAttributes(node ODNodeRef, recordType ODRecordType, error_ unsafe.Pointer) ArrayRef {
	return _ODNodeCopySupportedAttributes(node, recordType, error_)
}/* debug [functions.gen.go/function]: ODNodeCopySupportedAttributes */

// ODNodeCopySupportedPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedPolicies(_:_:)
func ODNodeCopySupportedPolicies(node ODNodeRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODNodeCopySupportedPolicies(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopySupportedPolicies */

// Returns an array of the record types supported by a given node.
//
// Added in macOS 10.6.
// Returns an array of the record types supported by a given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedRecordTypes(_:_:)
func ODNodeCopySupportedRecordTypes(node ODNodeRef, error_ unsafe.Pointer) ArrayRef {
	return _ODNodeCopySupportedRecordTypes(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopySupportedRecordTypes */

// Returns an array of the subnodes of a given node that are currently unreachable.
//
// Added in macOS 10.6.
// Returns an array of the subnodes of a given node that are currently unreachable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyUnreachableSubnodeNames(_:_:)
func ODNodeCopyUnreachableSubnodeNames(node ODNodeRef, error_ unsafe.Pointer) ArrayRef {
	return _ODNodeCopyUnreachableSubnodeNames(node, error_)
}/* debug [functions.gen.go/function]: ODNodeCopyUnreachableSubnodeNames */

// Returns a copy of an existing node.
//
// Added in macOS 10.6.
// Returns a copy of an existing node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateCopy(_:_:_:)
func ODNodeCreateCopy(allocator AllocatorRef, node ODNodeRef, error_ unsafe.Pointer) ODNodeRef {
	return _ODNodeCreateCopy(allocator, node, error_)
}/* debug [functions.gen.go/function]: ODNodeCreateCopy */

// Creates a record in a specified node with specified properties.
//
// Added in macOS 10.6.
// Creates a record in a specified node with specified properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateRecord(_:_:_:_:_:)
func ODNodeCreateRecord(node ODNodeRef, recordType ODRecordType, recordName StringRef, attributeDict DictionaryRef, error_ unsafe.Pointer) ODRecordRef {
	return _ODNodeCreateRecord(node, recordType, recordName, attributeDict, error_)
}/* debug [functions.gen.go/function]: ODNodeCreateRecord */

// Returns a new node created with a specified name.
//
// Added in macOS 10.6.
// Returns a new node created with a specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateWithName(_:_:_:_:)
func ODNodeCreateWithName(allocator AllocatorRef, session ODSessionRef, nodeName StringRef, error_ unsafe.Pointer) ODNodeRef {
	return _ODNodeCreateWithName(allocator, session, nodeName, error_)
}/* debug [functions.gen.go/function]: ODNodeCreateWithName */

// Returns a new node created with a specified type.
//
// Added in macOS 10.6.
// Returns a new node created with a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateWithNodeType(_:_:_:_:)
func ODNodeCreateWithNodeType(allocator AllocatorRef, session ODSessionRef, nodeType ODNodeType, error_ unsafe.Pointer) ODNodeRef {
	return _ODNodeCreateWithNodeType(allocator, session, nodeType, error_)
}/* debug [functions.gen.go/function]: ODNodeCreateWithNodeType */

// Returns the result of a custom call to a node.
//
// Added in macOS 10.6.
// Returns the result of a custom call to a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCustomCall(_:_:_:_:)
func ODNodeCustomCall(node ODNodeRef, customCode Index, data DataRef, error_ unsafe.Pointer) DataRef {
	return _ODNodeCustomCall(node, customCode, data, error_)
}/* debug [functions.gen.go/function]: ODNodeCustomCall */

// ODNodeCustomFunction is a OpenDirectory function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeCustomFunction(_:_:_:_:)
func ODNodeCustomFunction(node ODNodeRef, function StringRef, payload TypeRef, error_ unsafe.Pointer) TypeRef {
	return _ODNodeCustomFunction(node, function, payload, error_)
}/* debug [functions.gen.go/function]: ODNodeCustomFunction */

// Returns the name of a node.
//
// Added in macOS 10.6.
// Returns the name of a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeGetName(_:)
func ODNodeGetName(node ODNodeRef) StringRef {
	return _ODNodeGetName(node)
}/* debug [functions.gen.go/function]: ODNodeGetName */

// Returns the type ID for an Open Directory node.
//
// Added in macOS 10.6.
// Returns the type ID for an Open Directory node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeGetTypeID()
func ODNodeGetTypeID() TypeID {
	return _ODNodeGetTypeID()
}/* debug [functions.gen.go/function]: ODNodeGetTypeID */

// ODNodePasswordContentCheck is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodePasswordContentCheck(_:_:_:_:)
func ODNodePasswordContentCheck(node ODNodeRef, password StringRef, recordName StringRef, error_ unsafe.Pointer) bool {
	return _ODNodePasswordContentCheck(node, password, recordName, error_)
}/* debug [functions.gen.go/function]: ODNodePasswordContentCheck */

// ODNodeRemoveAccountPolicy is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeRemoveAccountPolicy(_:_:_:_:)
func ODNodeRemoveAccountPolicy(node ODNodeRef, policy DictionaryRef, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	return _ODNodeRemoveAccountPolicy(node, policy, category, error_)
}/* debug [functions.gen.go/function]: ODNodeRemoveAccountPolicy */

// ODNodeRemovePolicy is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeRemovePolicy(_:_:_:)
func ODNodeRemovePolicy(node ODNodeRef, policyType ODPolicyType, error_ unsafe.Pointer) bool {
	return _ODNodeRemovePolicy(node, policyType, error_)
}/* debug [functions.gen.go/function]: ODNodeRemovePolicy */

// ODNodeSetAccountPolicies is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetAccountPolicies(_:_:_:)
func ODNodeSetAccountPolicies(node ODNodeRef, policies DictionaryRef, error_ unsafe.Pointer) bool {
	return _ODNodeSetAccountPolicies(node, policies, error_)
}/* debug [functions.gen.go/function]: ODNodeSetAccountPolicies */

// Sets credentials for interacting with a node.
//
// Added in macOS 10.6.
// Sets credentials for interacting with a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentials(_:_:_:_:_:)
func ODNodeSetCredentials(node ODNodeRef, recordType ODRecordType, recordName StringRef, password StringRef, error_ unsafe.Pointer) bool {
	return _ODNodeSetCredentials(node, recordType, recordName, password, error_)
}/* debug [functions.gen.go/function]: ODNodeSetCredentials */

// Sets credentials for interacting with a node using a specified authentication method.
//
// Added in macOS 10.6.
// Sets credentials for interacting with a node using a specified authentication method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentialsExtended(_:_:_:_:_:_:_:)
func ODNodeSetCredentialsExtended(node ODNodeRef, recordType ODRecordType, authType ODAuthenticationType, authItems ArrayRef, outAuthItems unsafe.Pointer, outContext unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ODNodeSetCredentialsExtended(node, recordType, authType, authItems, outAuthItems, outContext, error_)
}/* debug [functions.gen.go/function]: ODNodeSetCredentialsExtended */

// Sets credentials for interacting with a node with the Kerberos cache.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.6.
// Sets credentials for interacting with a node with the Kerberos cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentialsUsingKerberosCache
func ODNodeSetCredentialsUsingKerberosCache(node ODNodeRef, cacheName StringRef, error_ unsafe.Pointer) bool {
	return _ODNodeSetCredentialsUsingKerberosCache(node, cacheName, error_)
}/* debug [functions.gen.go/function]: ODNodeSetCredentialsUsingKerberosCache */

// ODNodeSetPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetPolicies(_:_:_:)
func ODNodeSetPolicies(node ODNodeRef, policies DictionaryRef, error_ unsafe.Pointer) bool {
	return _ODNodeSetPolicies(node, policies, error_)
}/* debug [functions.gen.go/function]: ODNodeSetPolicies */

// ODNodeSetPolicy is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetPolicy(_:_:_:_:)
func ODNodeSetPolicy(node ODNodeRef, policyType ODPolicyType, value TypeRef, error_ unsafe.Pointer) bool {
	return _ODNodeSetPolicy(node, policyType, value, error_)
}/* debug [functions.gen.go/function]: ODNodeSetPolicy */

// Returns results from a query synchronously.
//
// Added in macOS 10.6.
// Returns results from a query synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryCopyResults(_:_:_:)
func ODQueryCopyResults(query ODQueryRef, allowPartialResults bool, error_ unsafe.Pointer) ArrayRef {
	return _ODQueryCopyResults(query, allowPartialResults, error_)
}/* debug [functions.gen.go/function]: ODQueryCopyResults */

// Creates a query with a node using provided parameters.
//
// Added in macOS 10.6.
// Creates a query with a node using provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryCreateWithNode(_:_:_:_:_:_:_:_:_:)
func ODQueryCreateWithNode(allocator AllocatorRef, node ODNodeRef, recordTypeOrList TypeRef, attribute ODAttributeType, matchType ODMatchType, queryValueOrList TypeRef, returnAttributeOrList TypeRef, maxResults Index, error_ unsafe.Pointer) ODQueryRef {
	return _ODQueryCreateWithNode(allocator, node, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, error_)
}/* debug [functions.gen.go/function]: ODQueryCreateWithNode */

// Creates a query for a particular node type using provided parameters.
//
// Added in macOS 10.6.
// Creates a query for a particular node type using provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryCreateWithNodeType(_:_:_:_:_:_:_:_:_:)
func ODQueryCreateWithNodeType(allocator AllocatorRef, nodeType ODNodeType, recordTypeOrList TypeRef, attribute ODAttributeType, matchType ODMatchType, queryValueOrList TypeRef, returnAttributeOrList TypeRef, maxResults Index, error_ unsafe.Pointer) ODQueryRef {
	return _ODQueryCreateWithNodeType(allocator, nodeType, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, error_)
}/* debug [functions.gen.go/function]: ODQueryCreateWithNodeType */

// Returns the type ID for an Open Directory query.
//
// Added in macOS 10.6.
// Returns the type ID for an Open Directory query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryGetTypeID()
func ODQueryGetTypeID() TypeID {
	return _ODQueryGetTypeID()
}/* debug [functions.gen.go/function]: ODQueryGetTypeID */

// Retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// Added in macOS 10.6.
// Retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryScheduleWithRunLoop(_:_:_:)
func ODQueryScheduleWithRunLoop(query ODQueryRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_ODQueryScheduleWithRunLoop(query, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: ODQueryScheduleWithRunLoop */

// Sets the callback for an asynchronous query.
//
// Added in macOS 10.6.
// Sets the callback for an asynchronous query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuerySetCallback(_:_:_:)
func ODQuerySetCallback(query ODQueryRef, callback ODQueryCallback, userInfo unsafe.Pointer) {
	_ODQuerySetCallback(query, callback, userInfo)
}/* debug [functions.gen.go/function]: ODQuerySetCallback */

// Retrieves results from a query asynchronously by adding the query to a dispatch queue.
//
// Added in macOS 10.6.
// Retrieves results from a query asynchronously by adding the query to a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuerySetDispatchQueue(_:_:)
func ODQuerySetDispatchQueue(query ODQueryRef, queue unsafe.Pointer) {
	_ODQuerySetDispatchQueue(query, queue)
}/* debug [functions.gen.go/function]: ODQuerySetDispatchQueue */

// Restarts a query, disposing of any results it has obtained.
//
// Added in macOS 10.6.
// Restarts a query, disposing of any results it has obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuerySynchronize(_:)
func ODQuerySynchronize(query ODQueryRef) {
	_ODQuerySynchronize(query)
}/* debug [functions.gen.go/function]: ODQuerySynchronize */

// Removes a query from a specified run loop.
//
// Added in macOS 10.6.
// Removes a query from a specified run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryUnscheduleFromRunLoop(_:_:_:)
func ODQueryUnscheduleFromRunLoop(query ODQueryRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_ODQueryUnscheduleFromRunLoop(query, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: ODQueryUnscheduleFromRunLoop */

// ODRecordAddAccountPolicy is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddAccountPolicy(_:_:_:_:)
func ODRecordAddAccountPolicy(record ODRecordRef, policy DictionaryRef, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	return _ODRecordAddAccountPolicy(record, policy, category, error_)
}/* debug [functions.gen.go/function]: ODRecordAddAccountPolicy */

// Adds a record as a member of a group record.
//
// Added in macOS 10.6.
// Adds a record as a member of a group record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddMember(_:_:_:)
func ODRecordAddMember(group ODRecordRef, member ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordAddMember(group, member, error_)
}/* debug [functions.gen.go/function]: ODRecordAddMember */

// Adds a value to an attribute of a record.
//
// Added in macOS 10.6.
// Adds a value to an attribute of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddValue(_:_:_:_:)
func ODRecordAddValue(record ODRecordRef, attribute ODAttributeType, value TypeRef, error_ unsafe.Pointer) bool {
	return _ODRecordAddValue(record, attribute, value, error_)
}/* debug [functions.gen.go/function]: ODRecordAddValue */

// ODRecordAuthenticationAllowed is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordAuthenticationAllowed(_:_:)
func ODRecordAuthenticationAllowed(record ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordAuthenticationAllowed(record, error_)
}/* debug [functions.gen.go/function]: ODRecordAuthenticationAllowed */

// Changes the password of a record.
//
// Added in macOS 10.6.
// Changes the password of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordChangePassword(_:_:_:_:)
func ODRecordChangePassword(record ODRecordRef, oldPassword StringRef, newPassword StringRef, error_ unsafe.Pointer) bool {
	return _ODRecordChangePassword(record, oldPassword, newPassword, error_)
}/* debug [functions.gen.go/function]: ODRecordChangePassword */

// Returns whether a group record contains a given record.
//
// Added in macOS 10.6.
// Returns whether a group record contains a given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordContainsMember(_:_:_:)
func ODRecordContainsMember(group ODRecordRef, member ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordContainsMember(group, member, error_)
}/* debug [functions.gen.go/function]: ODRecordContainsMember */

// ODRecordCopyAccountPolicies is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyAccountPolicies(_:_:)
func ODRecordCopyAccountPolicies(record ODRecordRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopyAccountPolicies(record, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyAccountPolicies */

// Returns the values of a record’s attributes.
//
// Added in macOS 10.6.
// Returns the values of a record’s attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyDetails(_:_:_:)
func ODRecordCopyDetails(record ODRecordRef, attributes ArrayRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopyDetails(record, attributes, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyDetails */

// ODRecordCopyEffectivePolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyEffectivePolicies(_:_:)
func ODRecordCopyEffectivePolicies(record ODRecordRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopyEffectivePolicies(record, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyEffectivePolicies */

// Returns the password policies of a record.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.6.
// Returns the password policies of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyPasswordPolicy
func ODRecordCopyPasswordPolicy(allocator AllocatorRef, record ODRecordRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopyPasswordPolicy(allocator, record, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyPasswordPolicy */

// ODRecordCopyPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyPolicies(_:_:)
func ODRecordCopyPolicies(record ODRecordRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopyPolicies(record, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyPolicies */

// ODRecordCopySupportedPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopySupportedPolicies(_:_:)
func ODRecordCopySupportedPolicies(record ODRecordRef, error_ unsafe.Pointer) DictionaryRef {
	return _ODRecordCopySupportedPolicies(record, error_)
}/* debug [functions.gen.go/function]: ODRecordCopySupportedPolicies */

// Returns the value of a single attribute of a record.
//
// Added in macOS 10.6.
// Returns the value of a single attribute of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyValues(_:_:_:)
func ODRecordCopyValues(record ODRecordRef, attribute ODAttributeType, error_ unsafe.Pointer) ArrayRef {
	return _ODRecordCopyValues(record, attribute, error_)
}/* debug [functions.gen.go/function]: ODRecordCopyValues */

// Deletes a record from a node and invalidates the record.
//
// Added in macOS 10.6.
// Deletes a record from a node and invalidates the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordDelete(_:_:)
func ODRecordDelete(record ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordDelete(record, error_)
}/* debug [functions.gen.go/function]: ODRecordDelete */

// Returns the official name of a record.
//
// Added in macOS 10.6.
// Returns the official name of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetRecordName(_:)
func ODRecordGetRecordName(record ODRecordRef) StringRef {
	return _ODRecordGetRecordName(record)
}/* debug [functions.gen.go/function]: ODRecordGetRecordName */

// Returns the type of a record.
//
// Added in macOS 10.6.
// Returns the type of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetRecordType(_:)
func ODRecordGetRecordType(record ODRecordRef) StringRef {
	return _ODRecordGetRecordType(record)
}/* debug [functions.gen.go/function]: ODRecordGetRecordType */

// Returns the type ID for a record.
//
// Added in macOS 10.6.
// Returns the type ID for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetTypeID()
func ODRecordGetTypeID() TypeID {
	return _ODRecordGetTypeID()
}/* debug [functions.gen.go/function]: ODRecordGetTypeID */

// ODRecordPasswordChangeAllowed is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordPasswordChangeAllowed(_:_:_:)
func ODRecordPasswordChangeAllowed(record ODRecordRef, newPassword StringRef, error_ unsafe.Pointer) bool {
	return _ODRecordPasswordChangeAllowed(record, newPassword, error_)
}/* debug [functions.gen.go/function]: ODRecordPasswordChangeAllowed */

// ODRecordRemoveAccountPolicy is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveAccountPolicy(_:_:_:_:)
func ODRecordRemoveAccountPolicy(record ODRecordRef, policy DictionaryRef, category ODPolicyCategoryType, error_ unsafe.Pointer) bool {
	return _ODRecordRemoveAccountPolicy(record, policy, category, error_)
}/* debug [functions.gen.go/function]: ODRecordRemoveAccountPolicy */

// Removes a record as a member from a specified group record.
//
// Added in macOS 10.6.
// Removes a record as a member from a specified group record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveMember(_:_:_:)
func ODRecordRemoveMember(group ODRecordRef, member ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordRemoveMember(group, member, error_)
}/* debug [functions.gen.go/function]: ODRecordRemoveMember */

// ODRecordRemovePolicy is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemovePolicy(_:_:_:)
func ODRecordRemovePolicy(record ODRecordRef, policy ODPolicyType, error_ unsafe.Pointer) bool {
	return _ODRecordRemovePolicy(record, policy, error_)
}/* debug [functions.gen.go/function]: ODRecordRemovePolicy */

// Removes a value from a record’s attribute.
//
// Added in macOS 10.6.
// Removes a value from a record’s attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveValue(_:_:_:_:)
func ODRecordRemoveValue(record ODRecordRef, attribute ODAttributeType, value TypeRef, error_ unsafe.Pointer) bool {
	return _ODRecordRemoveValue(record, attribute, value, error_)
}/* debug [functions.gen.go/function]: ODRecordRemoveValue */

// ODRecordSecondsUntilAuthenticationsExpire is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSecondsUntilAuthenticationsExpire(_:)
func ODRecordSecondsUntilAuthenticationsExpire(record ODRecordRef) int64 {
	return _ODRecordSecondsUntilAuthenticationsExpire(record)
}/* debug [functions.gen.go/function]: ODRecordSecondsUntilAuthenticationsExpire */

// ODRecordSecondsUntilPasswordExpires is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSecondsUntilPasswordExpires(_:)
func ODRecordSecondsUntilPasswordExpires(record ODRecordRef) int64 {
	return _ODRecordSecondsUntilPasswordExpires(record)
}/* debug [functions.gen.go/function]: ODRecordSecondsUntilPasswordExpires */

// ODRecordSetAccountPolicies is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetAccountPolicies(_:_:_:)
func ODRecordSetAccountPolicies(record ODRecordRef, policies DictionaryRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetAccountPolicies(record, policies, error_)
}/* debug [functions.gen.go/function]: ODRecordSetAccountPolicies */

// Sets node authentication credentials for a given record.
//
// Added in macOS 10.6.
// Sets node authentication credentials for a given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentials(_:_:_:_:)
func ODRecordSetNodeCredentials(record ODRecordRef, username StringRef, password StringRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetNodeCredentials(record, username, password, error_)
}/* debug [functions.gen.go/function]: ODRecordSetNodeCredentials */

// Sets node authentication credentials for a record using a specified authentication method.
//
// Added in macOS 10.6.
// Sets node authentication credentials for a record using a specified authentication method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentialsExtended(_:_:_:_:_:_:_:)
func ODRecordSetNodeCredentialsExtended(record ODRecordRef, recordType ODRecordType, authType ODAuthenticationType, authItems ArrayRef, outAuthItems unsafe.Pointer, outContext unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ODRecordSetNodeCredentialsExtended(record, recordType, authType, authItems, outAuthItems, outContext, error_)
}/* debug [functions.gen.go/function]: ODRecordSetNodeCredentialsExtended */

// Sets credentials for interacting with a record’s node with the Kerberos cache.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.6.
// Sets credentials for interacting with a record’s node with the Kerberos cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentialsUsingKerberosCache
func ODRecordSetNodeCredentialsUsingKerberosCache(record ODRecordRef, cacheName StringRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetNodeCredentialsUsingKerberosCache(record, cacheName, error_)
}/* debug [functions.gen.go/function]: ODRecordSetNodeCredentialsUsingKerberosCache */

// ODRecordSetPolicies is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetPolicies(_:_:_:)
func ODRecordSetPolicies(record ODRecordRef, policies DictionaryRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetPolicies(record, policies, error_)
}/* debug [functions.gen.go/function]: ODRecordSetPolicies */

// ODRecordSetPolicy is a OpenDirectory function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetPolicy(_:_:_:_:)
func ODRecordSetPolicy(record ODRecordRef, policy ODPolicyType, value TypeRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetPolicy(record, policy, value, error_)
}/* debug [functions.gen.go/function]: ODRecordSetPolicy */

// Sets one or more attribute values of a record.
//
// Added in macOS 10.6.
// Sets one or more attribute values of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetValue(_:_:_:_:)
func ODRecordSetValue(record ODRecordRef, attribute ODAttributeType, valueOrValues TypeRef, error_ unsafe.Pointer) bool {
	return _ODRecordSetValue(record, attribute, valueOrValues, error_)
}/* debug [functions.gen.go/function]: ODRecordSetValue */

// Synchronizes a record with the directory to get current data and commit changes.
//
// Added in macOS 10.6.
// Synchronizes a record with the directory to get current data and commit changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordSynchronize(_:_:)
func ODRecordSynchronize(record ODRecordRef, error_ unsafe.Pointer) bool {
	return _ODRecordSynchronize(record, error_)
}/* debug [functions.gen.go/function]: ODRecordSynchronize */

// Verifies a given password for a record.
//
// Added in macOS 10.6.
// Verifies a given password for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordVerifyPassword(_:_:_:)
func ODRecordVerifyPassword(record ODRecordRef, password StringRef, error_ unsafe.Pointer) bool {
	return _ODRecordVerifyPassword(record, password, error_)
}/* debug [functions.gen.go/function]: ODRecordVerifyPassword */

// Verifies a given password for a record given a specified authentication method.
//
// Added in macOS 10.6.
// Verifies a given password for a record given a specified authentication method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordVerifyPasswordExtended(_:_:_:_:_:_:)
func ODRecordVerifyPasswordExtended(record ODRecordRef, authType ODAuthenticationType, authItems ArrayRef, outAuthItems unsafe.Pointer, outContext unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ODRecordVerifyPasswordExtended(record, authType, authItems, outAuthItems, outContext, error_)
}/* debug [functions.gen.go/function]: ODRecordVerifyPasswordExtended */

// ODRecordWillAuthenticationsExpire is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordWillAuthenticationsExpire(_:_:)
func ODRecordWillAuthenticationsExpire(record ODRecordRef, willExpireIn uint64) bool {
	return _ODRecordWillAuthenticationsExpire(record, willExpireIn)
}/* debug [functions.gen.go/function]: ODRecordWillAuthenticationsExpire */

// ODRecordWillPasswordExpire is a OpenDirectory function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordWillPasswordExpire(_:_:)
func ODRecordWillPasswordExpire(record ODRecordRef, willExpireIn uint64) bool {
	return _ODRecordWillPasswordExpire(record, willExpireIn)
}/* debug [functions.gen.go/function]: ODRecordWillPasswordExpire */

// Returns the names of nodes registered in a given session.
//
// Added in macOS 10.6.
// Returns the names of nodes registered in a given session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSessionCopyNodeNames(_:_:_:)
func ODSessionCopyNodeNames(allocator AllocatorRef, session ODSessionRef, error_ unsafe.Pointer) ArrayRef {
	return _ODSessionCopyNodeNames(allocator, session, error_)
}/* debug [functions.gen.go/function]: ODSessionCopyNodeNames */

// Creates a session to be passed to node functions.
//
// Added in macOS 10.6.
// Creates a session to be passed to node functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSessionCreate(_:_:_:)
func ODSessionCreate(allocator AllocatorRef, options DictionaryRef, error_ unsafe.Pointer) ODSessionRef {
	return _ODSessionCreate(allocator, options, error_)
}/* debug [functions.gen.go/function]: ODSessionCreate */

// Returns the type ID for a session.
//
// Added in macOS 10.6.
// Returns the type ID for a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSessionGetTypeID()
func ODSessionGetTypeID() TypeID {
	return _ODSessionGetTypeID()
}/* debug [functions.gen.go/function]: ODSessionGetTypeID */




