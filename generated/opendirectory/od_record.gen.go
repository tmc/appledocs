// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	RecordName() objc.IObject /* cross-framework: NSString */
	RecordType() objc.IObject /* cross-framework: NSString */
	SecondsUntilAuthenticationsExpire() int64
	SecondsUntilPasswordExpires() int64
	// methods:
	AccountPoliciesAndReturnError(error_ unsafe.Pointer) objc.IObject /* cross-framework: Dictionary */
	AddAccountPolicyToCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool
	AddMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool
	AddValueToAttributeError(inValue objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool
	AuthenticationAllowedAndReturnError(error_ unsafe.Pointer) bool
	ChangePasswordToPasswordError(oldPassword objc.IObject /* cross-framework: NSString */, newPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	DeleteRecordAndReturnError(outError unsafe.Pointer) bool
	IsMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool
	PasswordChangeAllowedError(newPassword objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	RecordDetailsForAttributesError(inAttributes objc.IObject /* cross-framework: NSArray */, outError unsafe.Pointer) objc.IObject /* cross-framework: Dictionary */
	RemoveAccountPolicyFromCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool
	RemoveMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool
	RemoveValueFromAttributeError(inValue objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool
	RemoveValuesForAttributeError(inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool
	SetAccountPoliciesError(policies objc.IObject /* cross-framework: NSDictionary */, error_ unsafe.Pointer) bool
	SetNodeCredentialsPasswordError(inUsername objc.IObject /* cross-framework: NSString */, inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	SetNodeCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool
	SetValueForAttributeError(inValueOrValues objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool
	SynchronizeAndReturnError(outError unsafe.Pointer) bool
	ValuesForAttributeError(inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) objc.IObject /* cross-framework: Array */
	VerifyExtendedWithAuthenticationTypeAuthenticationItemsContinueItemsContextError(inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool
	VerifyPasswordError(inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	WillAuthenticationsExpire(willExpireIn uint64) bool
	WillPasswordExpire(willExpireIn uint64) bool
}

// An object serves as a Cocoa wrapper for an Open Directory record.


// An object serves as a Cocoa wrapper for an Open Directory record.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/accountPolicies()
func (o_ ODRecord) AccountPoliciesAndReturnError(error_ unsafe.Pointer) objc.IObject /* cross-framework: Dictionary */ {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("accountPoliciesAndReturnError:"), error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addAccountPolicy(_:toCategory:)
func (o_ ODRecord) AddAccountPolicyToCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addAccountPolicy:toCategory:error:"), policy, category, error_)
	return rv
}


// Adds a member record to this group record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addMemberRecord(_:)
func (o_ ODRecord) AddMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addMemberRecord:error:"), inRecord, outError)
	return rv
}


// Adds a value to an attribute of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addValue(_:toAttribute:)
func (o_ ODRecord) AddValueToAttributeError(inValue objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addValue:toAttribute:error:"), inValue, inAttribute, outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/authenticationAllowed()
func (o_ ODRecord) AuthenticationAllowedAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("authenticationAllowedAndReturnError:"), error_)
	return rv
}


// Changes the record’s password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/changePassword(_:toPassword:)
func (o_ ODRecord) ChangePasswordToPasswordError(oldPassword objc.IObject /* cross-framework: NSString */, newPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("changePassword:toPassword:error:"), oldPassword, newPassword, outError)
	return rv
}


// Deletes the record from its node and invalidates it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/delete()
func (o_ ODRecord) DeleteRecordAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteRecordAndReturnError:"), outError)
	return rv
}


// Determines whether a given record is a member of this group record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/isMemberRecord(_:)
func (o_ ODRecord) IsMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isMemberRecord:error:"), inRecord, outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/passwordChangeAllowed(_:)
func (o_ ODRecord) PasswordChangeAllowedError(newPassword objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("passwordChangeAllowed:error:"), newPassword, error_)
	return rv
}


// Returns a dictionary of attributes with their respective values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordDetails(forAttributes:)
func (o_ ODRecord) RecordDetailsForAttributesError(inAttributes objc.IObject /* cross-framework: NSArray */, outError unsafe.Pointer) objc.IObject /* cross-framework: Dictionary */ {
	rv := objc.Send[foundation.Dictionary](o_.ID, objc.Sel("recordDetailsForAttributes:error:"), inAttributes, outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeAccountPolicy(_:fromCategory:)
func (o_ ODRecord) RemoveAccountPolicyFromCategoryError(policy objc.IObject /* cross-framework: NSDictionary */, category ODPolicyCategoryType /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeAccountPolicy:fromCategory:error:"), policy, category, error_)
	return rv
}


// Removes a record as a member of this group record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeMemberRecord(_:)
func (o_ ODRecord) RemoveMemberRecordError(inRecord IODRecord, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeMemberRecord:error:"), inRecord, outError)
	return rv
}


// Removes a value from an attribute of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValue(_:fromAttribute:)
func (o_ ODRecord) RemoveValueFromAttributeError(inValue objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeValue:fromAttribute:error:"), inValue, inAttribute, outError)
	return rv
}


// Removes all values from an attribute of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValues(forAttribute:)
func (o_ ODRecord) RemoveValuesForAttributeError(inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeValuesForAttribute:error:"), inAttribute, outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setAccountPolicies(_:)
func (o_ ODRecord) SetAccountPoliciesError(policies objc.IObject /* cross-framework: NSDictionary */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setAccountPolicies:error:"), policies, error_)
	return rv
}


// Sets credentials for the record’s node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentials(_:password:)
func (o_ ODRecord) SetNodeCredentialsPasswordError(inUsername objc.IObject /* cross-framework: NSString */, inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setNodeCredentials:password:error:"), inUsername, inPassword, outError)
	return rv
}


// Sets the credentials for interaction with the record’s node using other types of authentication available to Open Directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentialsWithRecordType(_:authenticationType:authenticationItems:continueItems:context:)
func (o_ ODRecord) SetNodeCredentialsWithRecordTypeAuthenticationTypeAuthenticationItemsContinueItemsContextError(inRecordType unsafe.Pointer, inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setNodeCredentialsWithRecordType:authenticationType:authenticationItems:continueItems:context:error:"), inRecordType, inType, inItems, outItems, outContext, outError)
	return rv
}


// Sets the values of an attribute of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setValue(_:forAttribute:)
func (o_ ODRecord) SetValueForAttributeError(inValueOrValues objectivec.IObject, inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setValue:forAttribute:error:"), inValueOrValues, inAttribute, outError)
	return rv
}


// Synchronizes the record from the directory to get current data and commit changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/synchronize()
func (o_ ODRecord) SynchronizeAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("synchronizeAndReturnError:"), outError)
	return rv
}


// Returns the values of an attribute of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/values(forAttribute:)
func (o_ ODRecord) ValuesForAttributeError(inAttribute ODAttributeType /* typedef */, outError unsafe.Pointer) objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("valuesForAttribute:error:"), inAttribute, outError)
	return rv
}


// Verifies the credentials for interaction with the record’s node using other types of authentication available to Open Directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/verifyExtended(withAuthenticationType:authenticationItems:continueItems:context:)
func (o_ ODRecord) VerifyExtendedWithAuthenticationTypeAuthenticationItemsContinueItemsContextError(inType ODAuthenticationType /* typedef */, inItems objc.IObject /* cross-framework: NSArray */, outItems objc.IObject /* cross-framework: NSArray */, outContext unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("verifyExtendedWithAuthenticationType:authenticationItems:continueItems:context:error:"), inType, inItems, outItems, outContext, outError)
	return rv
}


// Verifies the password for interaction with the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/verifyPassword(_:)
func (o_ ODRecord) VerifyPasswordError(inPassword objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("verifyPassword:error:"), inPassword, outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willAuthenticationsExpire(_:)
func (o_ ODRecord) WillAuthenticationsExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("willAuthenticationsExpire:"), willExpireIn)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willPasswordExpire(_:)
func (o_ ODRecord) WillPasswordExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("willPasswordExpire:"), willExpireIn)
	return rv
}


// The official name of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordName
func (o_ ODRecord) RecordName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("recordName"))
	return rv
}


// The record’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordType
func (o_ ODRecord) RecordType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("recordType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilAuthenticationsExpire
func (o_ ODRecord) SecondsUntilAuthenticationsExpire() int64 {
	rv := objc.Send[int64](o_.ID, objc.Sel("secondsUntilAuthenticationsExpire"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilPasswordExpires
func (o_ ODRecord) SecondsUntilPasswordExpires() int64 {
	rv := objc.Send[int64](o_.ID, objc.Sel("secondsUntilPasswordExpires"))
	return rv
}



