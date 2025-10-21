// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ODRecord] class.
var (
	ODRecordClass     _ODRecordClass
	ODRecordClassOnce sync.Once
)

func getODRecordClass() _ODRecordClass {
	ODRecordClassOnce.Do(func() {
		ODRecordClass = _ODRecordClass{objc.GetClass("ODRecord")}
	})
	return ODRecordClass
}

type _ODRecordClass struct {
	class objc.Class
}

// An interface definition for the [ODRecord] class.
type IODRecord interface {
	objectivec.IObject
	AccountPoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
	AddAccountPolicyToCategoryError(policy objc.ID, category unsafe.Pointer, error_ unsafe.Pointer) bool
	AddMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool
	AddValueToAttributeError(inValue objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool
	AuthenticationAllowedAndReturnError(error_ unsafe.Pointer) bool
	ChangePasswordToPasswordError(oldPassword string, newPassword string, outError unsafe.Pointer) bool
	DeleteRecordAndReturnError(outError unsafe.Pointer) bool
	EffectivePoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
	IsMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool
	PasswordChangeAllowedError(newPassword string, error_ unsafe.Pointer) bool
	PasswordPolicyAndReturnError(outError unsafe.Pointer) unsafe.Pointer
	PoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
	RecordDetailsForAttributesError(inAttributes objc.ID, outError unsafe.Pointer) unsafe.Pointer
	RemoveAccountPolicyFromCategoryError(policy objc.ID, category unsafe.Pointer, error_ unsafe.Pointer) bool
	RemoveMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool
	RemovePolicyError(policy unsafe.Pointer, error_ unsafe.Pointer) bool
	RemoveValueFromAttributeError(inValue objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool
	RemoveValuesForAttributeError(inAttribute unsafe.Pointer, outError unsafe.Pointer) bool
	SetAccountPoliciesError(policies objc.ID, error_ unsafe.Pointer) bool
	SetNodeCredentialsPasswordError(inUsername string, inPassword string, outError unsafe.Pointer) bool
	SetNodeCredentialsUsingKerberosCacheError(inCacheName string, outError unsafe.Pointer) bool
	SetNodeCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType unsafe.Pointer, inItems objc.ID, outItems objc.ID, outContext objc.ID, outError unsafe.Pointer) bool
	SetPoliciesError(policies objc.ID, error_ unsafe.Pointer) bool
	SetPolicyValueError(policy unsafe.Pointer, value objc.ID, error_ unsafe.Pointer) bool
	SetValueForAttributeError(inValueOrValues objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool
	SupportedPoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
	SynchronizeAndReturnError(outError unsafe.Pointer) bool
	ValuesForAttributeError(inAttribute unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer
	VerifyExtendedWithAuthenticationTypeAuthenticationItemsContinueItemsContextError(inType unsafe.Pointer, inItems objc.ID, outItems objc.ID, outContext objc.ID, outError unsafe.Pointer) bool
	VerifyPasswordError(inPassword string, outError unsafe.Pointer) bool
	WillAuthenticationsExpire(willExpireIn uint64) bool
	WillPasswordExpire(willExpireIn uint64) bool
}

// An object serves as a Cocoa wrapper for an Open Directory record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord
type ODRecord struct {
	objectivec.Object
}

// ODRecordFrom constructs a [ODRecord] from an unsafe.Pointer.
//
// An object serves as a Cocoa wrapper for an Open Directory record.
func ODRecordFrom(ptr unsafe.Pointer) ODRecord {
	return ODRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODRecordClass) Alloc() ODRecord {
	rv := objc.Send[ODRecord](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODRecordClass) New() ODRecord {
	rv := objc.Send[ODRecord](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODRecord) Init() ODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODRecord) Autorelease() ODRecord {
	rv := objc.Send[ODRecord](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODRecord creates a new ODRecord instance.
func NewODRecord() ODRecord {
	return getODRecordClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/accountPolicies()
func (o_ ODRecord) AccountPoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accountPoliciesAndReturnError:"), error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addAccountPolicy(_:toCategory:)
func (o_ ODRecord) AddAccountPolicyToCategoryError(policy objc.ID, category unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addAccountPolicy:toCategory:error:"), policy, category, error_)
	return rv
}

// Adds a member record to this group record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addMemberRecord(_:)
func (o_ ODRecord) AddMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addMemberRecord:error:"), inRecord, outError)
	return rv
}

// Adds a value to an attribute of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addValue(_:toAttribute:)
func (o_ ODRecord) AddValueToAttributeError(inValue objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addValue:toAttribute:error:"), inValue, inAttribute, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/authenticationAllowed()
func (o_ ODRecord) AuthenticationAllowedAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("authenticationAllowedAndReturnError:"), error_)
	return rv
}

// Changes the record’s password.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/changePassword(_:toPassword:)
func (o_ ODRecord) ChangePasswordToPasswordError(oldPassword string, newPassword string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("changePassword:toPassword:error:"), objc.String(oldPassword), objc.String(newPassword), outError)
	return rv
}

// Deletes the record from its node and invalidates it.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/delete()
func (o_ ODRecord) DeleteRecordAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteRecordAndReturnError:"), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/effectivePolicies()
func (o_ ODRecord) EffectivePoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("effectivePoliciesAndReturnError:"), error_)
	return rv
}

// Determines whether a given record is a member of this group record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/isMemberRecord(_:)
func (o_ ODRecord) IsMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isMemberRecord:error:"), inRecord, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/passwordChangeAllowed(_:)
func (o_ ODRecord) PasswordChangeAllowedError(newPassword string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("passwordChangeAllowed:error:"), objc.String(newPassword), error_)
	return rv
}

// Returns a dictionary containing the password policy for the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/passwordPolicyAndReturnError:
func (o_ ODRecord) PasswordPolicyAndReturnError(outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("passwordPolicyAndReturnError:"), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/policies()
func (o_ ODRecord) PoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("policiesAndReturnError:"), error_)
	return rv
}

// Returns a dictionary of attributes with their respective values.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordDetails(forAttributes:)
func (o_ ODRecord) RecordDetailsForAttributesError(inAttributes objc.ID, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("recordDetailsForAttributes:error:"), inAttributes, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeAccountPolicy(_:fromCategory:)
func (o_ ODRecord) RemoveAccountPolicyFromCategoryError(policy objc.ID, category unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeAccountPolicy:fromCategory:error:"), policy, category, error_)
	return rv
}

// Removes a record as a member of this group record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeMemberRecord(_:)
func (o_ ODRecord) RemoveMemberRecordError(inRecord unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeMemberRecord:error:"), inRecord, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removePolicy(_:)
func (o_ ODRecord) RemovePolicyError(policy unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removePolicy:error:"), policy, error_)
	return rv
}

// Removes a value from an attribute of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValue(_:fromAttribute:)
func (o_ ODRecord) RemoveValueFromAttributeError(inValue objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeValue:fromAttribute:error:"), inValue, inAttribute, outError)
	return rv
}

// Removes all values from an attribute of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValues(forAttribute:)
func (o_ ODRecord) RemoveValuesForAttributeError(inAttribute unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeValuesForAttribute:error:"), inAttribute, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setAccountPolicies(_:)
func (o_ ODRecord) SetAccountPoliciesError(policies objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setAccountPolicies:error:"), policies, error_)
	return rv
}

// Sets credentials for the record’s node.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentials(_:password:)
func (o_ ODRecord) SetNodeCredentialsPasswordError(inUsername string, inPassword string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setNodeCredentials:password:error:"), objc.String(inUsername), objc.String(inPassword), outError)
	return rv
}

// Sets the credentials for interaction with the record’s node using a Kerberos cache.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentialsUsingKerberosCache:error:
func (o_ ODRecord) SetNodeCredentialsUsingKerberosCacheError(inCacheName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setNodeCredentialsUsingKerberosCache:error:"), objc.String(inCacheName), outError)
	return rv
}

// Sets the credentials for interaction with the record’s node using other types of authentication available to Open Directory.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentialsWithRecordType(_:authenticationType:authenticationItems:continueItems:context:)
func (o_ ODRecord) SetNodeCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType unsafe.Pointer, inItems objc.ID, outItems objc.ID, outContext objc.ID, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setNodeCredentialsWithRecordType:authenticationType:authenticationItems:continueItems:context:error:"), inRecordType, inType, inItems, outItems, outContext, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setPolicies(_:)
func (o_ ODRecord) SetPoliciesError(policies objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setPolicies:error:"), policies, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setPolicy(_:value:)
func (o_ ODRecord) SetPolicyValueError(policy unsafe.Pointer, value objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setPolicy:value:error:"), policy, value, error_)
	return rv
}

// Sets the values of an attribute of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setValue(_:forAttribute:)
func (o_ ODRecord) SetValueForAttributeError(inValueOrValues objc.ID, inAttribute unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setValue:forAttribute:error:"), inValueOrValues, inAttribute, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/supportedPolicies()
func (o_ ODRecord) SupportedPoliciesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("supportedPoliciesAndReturnError:"), error_)
	return rv
}

// Synchronizes the record from the directory to get current data and commit changes.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/synchronize()
func (o_ ODRecord) SynchronizeAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("synchronizeAndReturnError:"), outError)
	return rv
}

// Returns the values of an attribute of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/values(forAttribute:)
func (o_ ODRecord) ValuesForAttributeError(inAttribute unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("valuesForAttribute:error:"), inAttribute, outError)
	return rv
}

// Verifies the credentials for interaction with the record’s node using other types of authentication available to Open Directory.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/verifyExtended(withAuthenticationType:authenticationItems:continueItems:context:)
func (o_ ODRecord) VerifyExtendedWithAuthenticationTypeAuthenticationItemsContinueItemsContextError(inType unsafe.Pointer, inItems objc.ID, outItems objc.ID, outContext objc.ID, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("verifyExtendedWithAuthenticationType:authenticationItems:continueItems:context:error:"), inType, inItems, outItems, outContext, outError)
	return rv
}

// Verifies the password for interaction with the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/verifyPassword(_:)
func (o_ ODRecord) VerifyPasswordError(inPassword string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("verifyPassword:error:"), objc.String(inPassword), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willAuthenticationsExpire(_:)
func (o_ ODRecord) WillAuthenticationsExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("willAuthenticationsExpire:"), willExpireIn)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willPasswordExpire(_:)
func (o_ ODRecord) WillPasswordExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("willPasswordExpire:"), willExpireIn)
	return rv
}

// The official name of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordName
func (o_ ODRecord) RecordName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("recordName"))
	return rv
}

// The record’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordType
func (o_ ODRecord) RecordType() string {
	rv := objc.Send[string](o_.ID, objc.Sel("recordType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilAuthenticationsExpire
func (o_ ODRecord) SecondsUntilAuthenticationsExpire() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("secondsUntilAuthenticationsExpire"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilPasswordExpires
func (o_ ODRecord) SecondsUntilPasswordExpires() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("secondsUntilPasswordExpires"))
	return rv
}



