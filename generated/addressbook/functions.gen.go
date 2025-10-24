// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

/* debug [functions.gen.go]: Generating 153 functions for AddressBook */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AddressBook Functions (153 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_ABAddPropertiesAndTypes func(ABAddressBookRef, StringRef, DictionaryRef) Index
	_ABAddRecord func(ABAddressBookRef, ABRecordRef) bool
	_ABAddressBookAddRecord func(ABAddressBookRef, ABRecordRef, unsafe.Pointer) bool
	_ABAddressBookCopyArrayOfAllGroups func(ABAddressBookRef) ArrayRef
	_ABAddressBookCopyArrayOfAllGroupsInSource func(ABAddressBookRef, ABRecordRef) ArrayRef
	_ABAddressBookCopyArrayOfAllPeople func(ABAddressBookRef) ArrayRef
	_ABAddressBookCopyArrayOfAllPeopleInSource func(ABAddressBookRef, ABRecordRef) ArrayRef
	_ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering func(ABAddressBookRef, ABRecordRef, ABPersonSortOrdering) ArrayRef
	_ABAddressBookCopyArrayOfAllSources func(ABAddressBookRef) ArrayRef
	_ABAddressBookCopyDefaultSource func(ABAddressBookRef) ABRecordRef
	_ABAddressBookCopyLocalizedLabel func(StringRef) StringRef
	_ABAddressBookCopyPeopleWithName func(ABAddressBookRef, StringRef) ArrayRef
	_ABAddressBookCreate func() ABAddressBookRef
	_ABAddressBookCreateWithOptions func(DictionaryRef, unsafe.Pointer) ABAddressBookRef
	_ABAddressBookGetAuthorizationStatus func() ABAuthorizationStatus
	_ABAddressBookGetGroupCount func(ABAddressBookRef) Index
	_ABAddressBookGetGroupWithRecordID func(ABAddressBookRef, ABRecordID) ABRecordRef
	_ABAddressBookGetPersonCount func(ABAddressBookRef) Index
	_ABAddressBookGetPersonWithRecordID func(ABAddressBookRef, ABRecordID) ABRecordRef
	_ABAddressBookGetSourceWithRecordID func(ABAddressBookRef, ABRecordID) ABRecordRef
	_ABAddressBookHasUnsavedChanges func(ABAddressBookRef) bool
	_ABAddressBookRegisterExternalChangeCallback func(ABAddressBookRef, ABExternalChangeCallback, unsafe.Pointer)
	_ABAddressBookRemoveRecord func(ABAddressBookRef, ABRecordRef, unsafe.Pointer) bool
	_ABAddressBookRequestAccessWithCompletion func(ABAddressBookRef, unsafe.Pointer)
	_ABAddressBookRevert func(ABAddressBookRef)
	_ABAddressBookSave func(ABAddressBookRef, unsafe.Pointer) bool
	_ABAddressBookUnregisterExternalChangeCallback func(ABAddressBookRef, ABExternalChangeCallback, unsafe.Pointer)
	_ABBeginLoadingImageDataForClient func(ABPersonRef, ABImageClientCallback, unsafe.Pointer) Index
	_ABCancelLoadingImageDataForTag func(Index)
	_ABCopyArrayOfAllGroups func(ABAddressBookRef) ArrayRef
	_ABCopyArrayOfAllPeople func(ABAddressBookRef) ArrayRef
	_ABCopyArrayOfMatchingRecords func(ABAddressBookRef, ABSearchElementRef) ArrayRef
	_ABCopyArrayOfPropertiesForRecordType func(ABAddressBookRef, StringRef) ArrayRef
	_ABCopyDefaultCountryCode func(ABAddressBookRef) StringRef
	_ABCopyLocalizedPropertyOrLabel func(StringRef) StringRef
	_ABCopyRecordForUniqueId func(ABAddressBookRef, StringRef) ABRecordRef
	_ABCopyRecordTypeFromUniqueId func(ABAddressBookRef, StringRef) StringRef
	_ABCreateFormattedAddressFromDictionary func(ABAddressBookRef, DictionaryRef) StringRef
	_ABGetMe func(ABAddressBookRef) ABPersonRef
	_ABGetSharedAddressBook func() ABAddressBookRef
	_ABGroupAddGroup func(ABGroupRef, ABGroupRef) bool
	_ABGroupAddMember func(ABRecordRef, ABRecordRef, unsafe.Pointer) bool
	_ABGroupCopyArrayOfAllMembers func(ABGroupRef) ArrayRef
	_ABGroupCopyArrayOfAllMembersWithSortOrdering func(ABRecordRef, ABPersonSortOrdering) ArrayRef
	_ABGroupCopyArrayOfAllSubgroups func(ABGroupRef) ArrayRef
	_ABGroupCopyDistributionIdentifier func(ABGroupRef, ABPersonRef, StringRef) StringRef
	_ABGroupCopyParentGroups func(ABGroupRef) ArrayRef
	_ABGroupCopySource func(ABRecordRef) ABRecordRef
	_ABGroupCreate func() ABGroupRef
	_ABGroupCreateInSource func(ABRecordRef) ABRecordRef
	_ABGroupCreateSearchElement func(StringRef, StringRef, StringRef, TypeRef, ABSearchComparison) ABSearchElementRef
	_ABGroupRemoveGroup func(ABGroupRef, ABGroupRef) bool
	_ABGroupRemoveMember func(ABGroupRef, ABPersonRef) bool
	_ABGroupSetDistributionIdentifier func(ABGroupRef, ABPersonRef, StringRef, StringRef) bool
	_ABHasUnsavedChanges func(ABAddressBookRef) bool
	_ABLocalizedPropertyOrLabel func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueAdd func(ABMutableMultiValueRef, TypeRef, StringRef, unsafe.Pointer) bool
	_ABMultiValueAddValueAndLabel func(ABMutableMultiValueRef, TypeRef, StringRef, unsafe.Pointer) bool
	_ABMultiValueCopyArrayOfAllValues func(ABMultiValueRef) ArrayRef
	_ABMultiValueCopyIdentifierAtIndex func(ABMultiValueRef, Index) StringRef
	_ABMultiValueCopyLabelAtIndex func(ABMultiValueRef, Index) StringRef
	_ABMultiValueCopyPrimaryIdentifier func(ABMultiValueRef) StringRef
	_ABMultiValueCopyValueAtIndex func(ABMultiValueRef, Index) TypeRef
	_ABMultiValueCount func(ABMultiValueRef) Index
	_ABMultiValueCreate func() ABMultiValueRef
	_ABMultiValueCreateCopy func(ABMultiValueRef) ABMultiValueRef
	_ABMultiValueCreateMutable func() ABMutableMultiValueRef
	_ABMultiValueCreateMutableCopy func(ABMultiValueRef) ABMutableMultiValueRef
	_ABMultiValueGetCount func(ABMultiValueRef) Index
	_ABMultiValueGetFirstIndexOfValue func(ABMultiValueRef, TypeRef) Index
	_ABMultiValueGetIdentifierAtIndex func(ABMultiValueRef, Index) ABMultiValueIdentifier
	_ABMultiValueGetIndexForIdentifier func(ABMultiValueRef, ABMultiValueIdentifier) Index
	_ABMultiValueGetPropertyType func(ABMultiValueRef) ABPropertyType
	_ABMultiValueIndexForIdentifier func(ABMultiValueRef, StringRef) Index
	_ABMultiValueInsert func(ABMutableMultiValueRef, TypeRef, StringRef, Index, unsafe.Pointer) bool
	_ABMultiValueInsertValueAndLabelAtIndex func(ABMutableMultiValueRef, TypeRef, StringRef, Index, unsafe.Pointer) bool
	_ABMultiValuePropertyType func(ABMultiValueRef) ABPropertyType
	_ABMultiValueRemove func(ABMutableMultiValueRef, Index) bool
	_ABMultiValueRemoveValueAndLabelAtIndex func(ABMutableMultiValueRef, Index) bool
	_ABMultiValueReplaceLabel func(ABMutableMultiValueRef, StringRef, Index) bool
	_ABMultiValueReplaceLabelAtIndex func(ABMutableMultiValueRef, StringRef, Index) bool
	_ABMultiValueReplaceValue func(ABMutableMultiValueRef, TypeRef, Index) bool
	_ABMultiValueReplaceValueAtIndex func(ABMutableMultiValueRef, TypeRef, Index) bool
	_ABMultiValueSetPrimaryIdentifier func(ABMutableMultiValueRef, StringRef) bool
	_ABPersonComparePeopleByName func(ABRecordRef, ABRecordRef, ABPersonSortOrdering) ComparisonResult
	_ABPersonCopyArrayOfAllLinkedPeople func(ABRecordRef) ArrayRef
	_ABPersonCopyCompositeNameDelimiterForRecord func(ABRecordRef) StringRef
	_ABPersonCopyImageData func(ABPersonRef) DataRef
	_ABPersonCopyImageDataWithFormat func(ABRecordRef, unsafe.Pointer) DataRef
	_ABPersonCopyLocalizedPropertyName func(ABPropertyID) StringRef
	_ABPersonCopyParentGroups func(ABPersonRef) ArrayRef
	_ABPersonCopySource func(ABRecordRef) ABRecordRef
	_ABPersonCopyVCardRepresentation func(ABPersonRef) DataRef
	_ABPersonCreate func() ABRecordRef
	_ABPersonCreateInSource func(ABRecordRef) ABRecordRef
	_ABPersonCreatePeopleInSourceWithVCardRepresentation func(ABRecordRef, DataRef) ArrayRef
	_ABPersonCreateSearchElement func(StringRef, StringRef, StringRef, TypeRef, ABSearchComparison) ABSearchElementRef
	_ABPersonCreateVCardRepresentationWithPeople func(ArrayRef) DataRef
	_ABPersonCreateWithVCardRepresentation func(DataRef) ABPersonRef
	_ABPersonGetCompositeNameFormat func() ABPersonCompositeNameFormat
	_ABPersonGetCompositeNameFormatForRecord func(ABRecordRef) ABPersonCompositeNameFormat
	_ABPersonGetSortOrdering func() ABPersonSortOrdering
	_ABPersonGetTypeOfProperty func(ABPropertyID) ABPropertyType
	_ABPersonHasImageData func(ABRecordRef) bool
	_ABPersonRemoveImageData func(ABRecordRef, unsafe.Pointer) bool
	_ABPersonSetImageData func(ABPersonRef, DataRef) bool
	_ABPickerAddProperty func(unsafe.Pointer, StringRef)
	_ABPickerChangeAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_ABPickerClearSearchField func(unsafe.Pointer)
	_ABPickerCopyColumnTitle func(unsafe.Pointer, StringRef) StringRef
	_ABPickerCopyDisplayedProperty func(unsafe.Pointer) StringRef
	_ABPickerCopyProperties func(unsafe.Pointer) ArrayRef
	_ABPickerCopySelectedGroups func(unsafe.Pointer) ArrayRef
	_ABPickerCopySelectedIdentifiers func(unsafe.Pointer, ABPersonRef) ArrayRef
	_ABPickerCopySelectedRecords func(unsafe.Pointer) ArrayRef
	_ABPickerCopySelectedValues func(unsafe.Pointer) ArrayRef
	_ABPickerCreate func() unsafe.Pointer
	_ABPickerDeselectAll func(unsafe.Pointer)
	_ABPickerDeselectGroup func(unsafe.Pointer, ABGroupRef)
	_ABPickerDeselectIdentifier func(unsafe.Pointer, ABPersonRef, StringRef)
	_ABPickerDeselectRecord func(unsafe.Pointer, ABRecordRef)
	_ABPickerEditInAddressBook func(unsafe.Pointer)
	_ABPickerGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_ABPickerGetDelegate func(unsafe.Pointer) unsafe.Pointer
	_ABPickerGetFrame func(unsafe.Pointer, unsafe.Pointer)
	_ABPickerIsVisible func(unsafe.Pointer) bool
	_ABPickerRemoveProperty func(unsafe.Pointer, StringRef)
	_ABPickerSelectGroup func(unsafe.Pointer, ABGroupRef, bool)
	_ABPickerSelectIdentifier func(unsafe.Pointer, ABPersonRef, StringRef, bool)
	_ABPickerSelectInAddressBook func(unsafe.Pointer)
	_ABPickerSelectRecord func(unsafe.Pointer, ABRecordRef, bool)
	_ABPickerSetColumnTitle func(unsafe.Pointer, StringRef, StringRef)
	_ABPickerSetDelegate func(unsafe.Pointer, unsafe.Pointer)
	_ABPickerSetDisplayedProperty func(unsafe.Pointer, StringRef)
	_ABPickerSetFrame func(unsafe.Pointer, unsafe.Pointer)
	_ABPickerSetVisibility func(unsafe.Pointer, bool)
	_ABRecordCopyCompositeName func(ABRecordRef) StringRef
	_ABRecordCopyRecordType func(ABRecordRef) StringRef
	_ABRecordCopyUniqueId func(ABRecordRef) StringRef
	_ABRecordCopyValue func(ABRecordRef, StringRef) TypeRef
	_ABRecordCreateCopy func(ABRecordRef) ABRecordRef
	_ABRecordGetRecordID func(ABRecordRef) ABRecordID
	_ABRecordGetRecordType func(ABRecordRef) ABRecordType
	_ABRecordIsReadOnly func(ABRecordRef) bool
	_ABRecordRemoveValue func(ABRecordRef, StringRef) bool
	_ABRecordSetValue func(ABRecordRef, ABPropertyID, TypeRef, unsafe.Pointer) bool
	_ABRemoveProperties func(ABAddressBookRef, StringRef, ArrayRef) Index
	_ABRemoveRecord func(ABAddressBookRef, ABRecordRef) bool
	_ABSave func(ABAddressBookRef) bool
	_ABSearchElementCreateWithConjunction func(ABSearchConjunction, ArrayRef) ABSearchElementRef
	_ABSearchElementMatchesRecord func(ABSearchElementRef, ABRecordRef) bool
	_ABSetMe func(ABAddressBookRef, ABPersonRef)
	_ABTypeOfProperty func(ABAddressBookRef, StringRef, StringRef) ABPropertyType
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_ABAddPropertiesAndTypes, lib, "ABAddPropertiesAndTypes")
	tryRegister(&_ABAddRecord, lib, "ABAddRecord")
	tryRegister(&_ABAddressBookAddRecord, lib, "ABAddressBookAddRecord")
	tryRegister(&_ABAddressBookCopyArrayOfAllGroups, lib, "ABAddressBookCopyArrayOfAllGroups")
	tryRegister(&_ABAddressBookCopyArrayOfAllGroupsInSource, lib, "ABAddressBookCopyArrayOfAllGroupsInSource")
	tryRegister(&_ABAddressBookCopyArrayOfAllPeople, lib, "ABAddressBookCopyArrayOfAllPeople")
	tryRegister(&_ABAddressBookCopyArrayOfAllPeopleInSource, lib, "ABAddressBookCopyArrayOfAllPeopleInSource")
	tryRegister(&_ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering, lib, "ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering")
	tryRegister(&_ABAddressBookCopyArrayOfAllSources, lib, "ABAddressBookCopyArrayOfAllSources")
	tryRegister(&_ABAddressBookCopyDefaultSource, lib, "ABAddressBookCopyDefaultSource")
	tryRegister(&_ABAddressBookCopyLocalizedLabel, lib, "ABAddressBookCopyLocalizedLabel")
	tryRegister(&_ABAddressBookCopyPeopleWithName, lib, "ABAddressBookCopyPeopleWithName")
	tryRegister(&_ABAddressBookCreate, lib, "ABAddressBookCreate")
	tryRegister(&_ABAddressBookCreateWithOptions, lib, "ABAddressBookCreateWithOptions")
	tryRegister(&_ABAddressBookGetAuthorizationStatus, lib, "ABAddressBookGetAuthorizationStatus")
	tryRegister(&_ABAddressBookGetGroupCount, lib, "ABAddressBookGetGroupCount")
	tryRegister(&_ABAddressBookGetGroupWithRecordID, lib, "ABAddressBookGetGroupWithRecordID")
	tryRegister(&_ABAddressBookGetPersonCount, lib, "ABAddressBookGetPersonCount")
	tryRegister(&_ABAddressBookGetPersonWithRecordID, lib, "ABAddressBookGetPersonWithRecordID")
	tryRegister(&_ABAddressBookGetSourceWithRecordID, lib, "ABAddressBookGetSourceWithRecordID")
	tryRegister(&_ABAddressBookHasUnsavedChanges, lib, "ABAddressBookHasUnsavedChanges")
	tryRegister(&_ABAddressBookRegisterExternalChangeCallback, lib, "ABAddressBookRegisterExternalChangeCallback")
	tryRegister(&_ABAddressBookRemoveRecord, lib, "ABAddressBookRemoveRecord")
	tryRegister(&_ABAddressBookRequestAccessWithCompletion, lib, "ABAddressBookRequestAccessWithCompletion")
	tryRegister(&_ABAddressBookRevert, lib, "ABAddressBookRevert")
	tryRegister(&_ABAddressBookSave, lib, "ABAddressBookSave")
	tryRegister(&_ABAddressBookUnregisterExternalChangeCallback, lib, "ABAddressBookUnregisterExternalChangeCallback")
	tryRegister(&_ABBeginLoadingImageDataForClient, lib, "ABBeginLoadingImageDataForClient")
	tryRegister(&_ABCancelLoadingImageDataForTag, lib, "ABCancelLoadingImageDataForTag")
	tryRegister(&_ABCopyArrayOfAllGroups, lib, "ABCopyArrayOfAllGroups")
	tryRegister(&_ABCopyArrayOfAllPeople, lib, "ABCopyArrayOfAllPeople")
	tryRegister(&_ABCopyArrayOfMatchingRecords, lib, "ABCopyArrayOfMatchingRecords")
	tryRegister(&_ABCopyArrayOfPropertiesForRecordType, lib, "ABCopyArrayOfPropertiesForRecordType")
	tryRegister(&_ABCopyDefaultCountryCode, lib, "ABCopyDefaultCountryCode")
	tryRegister(&_ABCopyLocalizedPropertyOrLabel, lib, "ABCopyLocalizedPropertyOrLabel")
	tryRegister(&_ABCopyRecordForUniqueId, lib, "ABCopyRecordForUniqueId")
	tryRegister(&_ABCopyRecordTypeFromUniqueId, lib, "ABCopyRecordTypeFromUniqueId")
	tryRegister(&_ABCreateFormattedAddressFromDictionary, lib, "ABCreateFormattedAddressFromDictionary")
	tryRegister(&_ABGetMe, lib, "ABGetMe")
	tryRegister(&_ABGetSharedAddressBook, lib, "ABGetSharedAddressBook")
	tryRegister(&_ABGroupAddGroup, lib, "ABGroupAddGroup")
	tryRegister(&_ABGroupAddMember, lib, "ABGroupAddMember")
	tryRegister(&_ABGroupCopyArrayOfAllMembers, lib, "ABGroupCopyArrayOfAllMembers")
	tryRegister(&_ABGroupCopyArrayOfAllMembersWithSortOrdering, lib, "ABGroupCopyArrayOfAllMembersWithSortOrdering")
	tryRegister(&_ABGroupCopyArrayOfAllSubgroups, lib, "ABGroupCopyArrayOfAllSubgroups")
	tryRegister(&_ABGroupCopyDistributionIdentifier, lib, "ABGroupCopyDistributionIdentifier")
	tryRegister(&_ABGroupCopyParentGroups, lib, "ABGroupCopyParentGroups")
	tryRegister(&_ABGroupCopySource, lib, "ABGroupCopySource")
	tryRegister(&_ABGroupCreate, lib, "ABGroupCreate")
	tryRegister(&_ABGroupCreateInSource, lib, "ABGroupCreateInSource")
	tryRegister(&_ABGroupCreateSearchElement, lib, "ABGroupCreateSearchElement")
	tryRegister(&_ABGroupRemoveGroup, lib, "ABGroupRemoveGroup")
	tryRegister(&_ABGroupRemoveMember, lib, "ABGroupRemoveMember")
	tryRegister(&_ABGroupSetDistributionIdentifier, lib, "ABGroupSetDistributionIdentifier")
	tryRegister(&_ABHasUnsavedChanges, lib, "ABHasUnsavedChanges")
	tryRegister(&_ABLocalizedPropertyOrLabel, lib, "ABLocalizedPropertyOrLabel")
	tryRegister(&_ABMultiValueAdd, lib, "ABMultiValueAdd")
	tryRegister(&_ABMultiValueAddValueAndLabel, lib, "ABMultiValueAddValueAndLabel")
	tryRegister(&_ABMultiValueCopyArrayOfAllValues, lib, "ABMultiValueCopyArrayOfAllValues")
	tryRegister(&_ABMultiValueCopyIdentifierAtIndex, lib, "ABMultiValueCopyIdentifierAtIndex")
	tryRegister(&_ABMultiValueCopyLabelAtIndex, lib, "ABMultiValueCopyLabelAtIndex")
	tryRegister(&_ABMultiValueCopyPrimaryIdentifier, lib, "ABMultiValueCopyPrimaryIdentifier")
	tryRegister(&_ABMultiValueCopyValueAtIndex, lib, "ABMultiValueCopyValueAtIndex")
	tryRegister(&_ABMultiValueCount, lib, "ABMultiValueCount")
	tryRegister(&_ABMultiValueCreate, lib, "ABMultiValueCreate")
	tryRegister(&_ABMultiValueCreateCopy, lib, "ABMultiValueCreateCopy")
	tryRegister(&_ABMultiValueCreateMutable, lib, "ABMultiValueCreateMutable")
	tryRegister(&_ABMultiValueCreateMutableCopy, lib, "ABMultiValueCreateMutableCopy")
	tryRegister(&_ABMultiValueGetCount, lib, "ABMultiValueGetCount")
	tryRegister(&_ABMultiValueGetFirstIndexOfValue, lib, "ABMultiValueGetFirstIndexOfValue")
	tryRegister(&_ABMultiValueGetIdentifierAtIndex, lib, "ABMultiValueGetIdentifierAtIndex")
	tryRegister(&_ABMultiValueGetIndexForIdentifier, lib, "ABMultiValueGetIndexForIdentifier")
	tryRegister(&_ABMultiValueGetPropertyType, lib, "ABMultiValueGetPropertyType")
	tryRegister(&_ABMultiValueIndexForIdentifier, lib, "ABMultiValueIndexForIdentifier")
	tryRegister(&_ABMultiValueInsert, lib, "ABMultiValueInsert")
	tryRegister(&_ABMultiValueInsertValueAndLabelAtIndex, lib, "ABMultiValueInsertValueAndLabelAtIndex")
	tryRegister(&_ABMultiValuePropertyType, lib, "ABMultiValuePropertyType")
	tryRegister(&_ABMultiValueRemove, lib, "ABMultiValueRemove")
	tryRegister(&_ABMultiValueRemoveValueAndLabelAtIndex, lib, "ABMultiValueRemoveValueAndLabelAtIndex")
	tryRegister(&_ABMultiValueReplaceLabel, lib, "ABMultiValueReplaceLabel")
	tryRegister(&_ABMultiValueReplaceLabelAtIndex, lib, "ABMultiValueReplaceLabelAtIndex")
	tryRegister(&_ABMultiValueReplaceValue, lib, "ABMultiValueReplaceValue")
	tryRegister(&_ABMultiValueReplaceValueAtIndex, lib, "ABMultiValueReplaceValueAtIndex")
	tryRegister(&_ABMultiValueSetPrimaryIdentifier, lib, "ABMultiValueSetPrimaryIdentifier")
	tryRegister(&_ABPersonComparePeopleByName, lib, "ABPersonComparePeopleByName")
	tryRegister(&_ABPersonCopyArrayOfAllLinkedPeople, lib, "ABPersonCopyArrayOfAllLinkedPeople")
	tryRegister(&_ABPersonCopyCompositeNameDelimiterForRecord, lib, "ABPersonCopyCompositeNameDelimiterForRecord")
	tryRegister(&_ABPersonCopyImageData, lib, "ABPersonCopyImageData")
	tryRegister(&_ABPersonCopyImageDataWithFormat, lib, "ABPersonCopyImageDataWithFormat")
	tryRegister(&_ABPersonCopyLocalizedPropertyName, lib, "ABPersonCopyLocalizedPropertyName")
	tryRegister(&_ABPersonCopyParentGroups, lib, "ABPersonCopyParentGroups")
	tryRegister(&_ABPersonCopySource, lib, "ABPersonCopySource")
	tryRegister(&_ABPersonCopyVCardRepresentation, lib, "ABPersonCopyVCardRepresentation")
	tryRegister(&_ABPersonCreate, lib, "ABPersonCreate")
	tryRegister(&_ABPersonCreateInSource, lib, "ABPersonCreateInSource")
	tryRegister(&_ABPersonCreatePeopleInSourceWithVCardRepresentation, lib, "ABPersonCreatePeopleInSourceWithVCardRepresentation")
	tryRegister(&_ABPersonCreateSearchElement, lib, "ABPersonCreateSearchElement")
	tryRegister(&_ABPersonCreateVCardRepresentationWithPeople, lib, "ABPersonCreateVCardRepresentationWithPeople")
	tryRegister(&_ABPersonCreateWithVCardRepresentation, lib, "ABPersonCreateWithVCardRepresentation")
	tryRegister(&_ABPersonGetCompositeNameFormat, lib, "ABPersonGetCompositeNameFormat")
	tryRegister(&_ABPersonGetCompositeNameFormatForRecord, lib, "ABPersonGetCompositeNameFormatForRecord")
	tryRegister(&_ABPersonGetSortOrdering, lib, "ABPersonGetSortOrdering")
	tryRegister(&_ABPersonGetTypeOfProperty, lib, "ABPersonGetTypeOfProperty")
	tryRegister(&_ABPersonHasImageData, lib, "ABPersonHasImageData")
	tryRegister(&_ABPersonRemoveImageData, lib, "ABPersonRemoveImageData")
	tryRegister(&_ABPersonSetImageData, lib, "ABPersonSetImageData")
	tryRegister(&_ABPickerAddProperty, lib, "ABPickerAddProperty")
	tryRegister(&_ABPickerChangeAttributes, lib, "ABPickerChangeAttributes")
	tryRegister(&_ABPickerClearSearchField, lib, "ABPickerClearSearchField")
	tryRegister(&_ABPickerCopyColumnTitle, lib, "ABPickerCopyColumnTitle")
	tryRegister(&_ABPickerCopyDisplayedProperty, lib, "ABPickerCopyDisplayedProperty")
	tryRegister(&_ABPickerCopyProperties, lib, "ABPickerCopyProperties")
	tryRegister(&_ABPickerCopySelectedGroups, lib, "ABPickerCopySelectedGroups")
	tryRegister(&_ABPickerCopySelectedIdentifiers, lib, "ABPickerCopySelectedIdentifiers")
	tryRegister(&_ABPickerCopySelectedRecords, lib, "ABPickerCopySelectedRecords")
	tryRegister(&_ABPickerCopySelectedValues, lib, "ABPickerCopySelectedValues")
	tryRegister(&_ABPickerCreate, lib, "ABPickerCreate")
	tryRegister(&_ABPickerDeselectAll, lib, "ABPickerDeselectAll")
	tryRegister(&_ABPickerDeselectGroup, lib, "ABPickerDeselectGroup")
	tryRegister(&_ABPickerDeselectIdentifier, lib, "ABPickerDeselectIdentifier")
	tryRegister(&_ABPickerDeselectRecord, lib, "ABPickerDeselectRecord")
	tryRegister(&_ABPickerEditInAddressBook, lib, "ABPickerEditInAddressBook")
	tryRegister(&_ABPickerGetAttributes, lib, "ABPickerGetAttributes")
	tryRegister(&_ABPickerGetDelegate, lib, "ABPickerGetDelegate")
	tryRegister(&_ABPickerGetFrame, lib, "ABPickerGetFrame")
	tryRegister(&_ABPickerIsVisible, lib, "ABPickerIsVisible")
	tryRegister(&_ABPickerRemoveProperty, lib, "ABPickerRemoveProperty")
	tryRegister(&_ABPickerSelectGroup, lib, "ABPickerSelectGroup")
	tryRegister(&_ABPickerSelectIdentifier, lib, "ABPickerSelectIdentifier")
	tryRegister(&_ABPickerSelectInAddressBook, lib, "ABPickerSelectInAddressBook")
	tryRegister(&_ABPickerSelectRecord, lib, "ABPickerSelectRecord")
	tryRegister(&_ABPickerSetColumnTitle, lib, "ABPickerSetColumnTitle")
	tryRegister(&_ABPickerSetDelegate, lib, "ABPickerSetDelegate")
	tryRegister(&_ABPickerSetDisplayedProperty, lib, "ABPickerSetDisplayedProperty")
	tryRegister(&_ABPickerSetFrame, lib, "ABPickerSetFrame")
	tryRegister(&_ABPickerSetVisibility, lib, "ABPickerSetVisibility")
	tryRegister(&_ABRecordCopyCompositeName, lib, "ABRecordCopyCompositeName")
	tryRegister(&_ABRecordCopyRecordType, lib, "ABRecordCopyRecordType")
	tryRegister(&_ABRecordCopyUniqueId, lib, "ABRecordCopyUniqueId")
	tryRegister(&_ABRecordCopyValue, lib, "ABRecordCopyValue")
	tryRegister(&_ABRecordCreateCopy, lib, "ABRecordCreateCopy")
	tryRegister(&_ABRecordGetRecordID, lib, "ABRecordGetRecordID")
	tryRegister(&_ABRecordGetRecordType, lib, "ABRecordGetRecordType")
	tryRegister(&_ABRecordIsReadOnly, lib, "ABRecordIsReadOnly")
	tryRegister(&_ABRecordRemoveValue, lib, "ABRecordRemoveValue")
	tryRegister(&_ABRecordSetValue, lib, "ABRecordSetValue")
	tryRegister(&_ABRemoveProperties, lib, "ABRemoveProperties")
	tryRegister(&_ABRemoveRecord, lib, "ABRemoveRecord")
	tryRegister(&_ABSave, lib, "ABSave")
	tryRegister(&_ABSearchElementCreateWithConjunction, lib, "ABSearchElementCreateWithConjunction")
	tryRegister(&_ABSearchElementMatchesRecord, lib, "ABSearchElementMatchesRecord")
	tryRegister(&_ABSetMe, lib, "ABSetMe")
	tryRegister(&_ABTypeOfProperty, lib, "ABTypeOfProperty")
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



// Adds the given properties to all the records of the specified type in the Address Book database, and returns the number of properties successfully added.
//
// Added in macOS .
// Adds the given properties to all the records of the specified type in the Address Book database, and returns the number of properties successfully added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddPropertiesAndTypes(_:_:_:)
func ABAddPropertiesAndTypes(addressBook ABAddressBookRef, recordType StringRef, propertiesAndTypes DictionaryRef) Index {
	return _ABAddPropertiesAndTypes(addressBook, recordType, propertiesAndTypes)
}/* debug [functions.gen.go/function]: ABAddPropertiesAndTypes */

// Adds a record of the specified type to the Address Book database.
//
// Added in macOS .
// Adds a record of the specified type to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddRecord(_:_:)
func ABAddRecord(addressBook ABAddressBookRef, record ABRecordRef) bool {
	return _ABAddRecord(addressBook, record)
}/* debug [functions.gen.go/function]: ABAddRecord */

// Adds a record to an address book.

// Adds a record to an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookAddRecord(_:_:_:)
func ABAddressBookAddRecord(addressBook ABAddressBookRef, record ABRecordRef, error_ unsafe.Pointer) bool {
	return _ABAddressBookAddRecord(addressBook, record, error_)
}/* debug [functions.gen.go/function]: ABAddressBookAddRecord */

// Returns an array with all the groups in an address book.

// Returns an array with all the groups in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllGroups(_:)
func ABAddressBookCopyArrayOfAllGroups(addressBook ABAddressBookRef) ArrayRef {
	return _ABAddressBookCopyArrayOfAllGroups(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllGroups */

// Returns an array of all groups from a particular source.

// Returns an array of all groups from a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllGroupsInSource(_:_:)
func ABAddressBookCopyArrayOfAllGroupsInSource(addressBook ABAddressBookRef, source ABRecordRef) ArrayRef {
	return _ABAddressBookCopyArrayOfAllGroupsInSource(addressBook, source)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllGroupsInSource */

// Returns all the person records in an address book.

// Returns all the person records in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeople(_:)
func ABAddressBookCopyArrayOfAllPeople(addressBook ABAddressBookRef) ArrayRef {
	return _ABAddressBookCopyArrayOfAllPeople(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllPeople */

// Returns an array of all person records from a particular source.

// Returns an array of all person records from a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeopleInSource(_:_:)
func ABAddressBookCopyArrayOfAllPeopleInSource(addressBook ABAddressBookRef, source ABRecordRef) ArrayRef {
	return _ABAddressBookCopyArrayOfAllPeopleInSource(addressBook, source)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllPeopleInSource */

// Returns an array of all person records in the address book, sorted with the specified order.

// Returns an array of all person records in the address book, sorted with the specified order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(_:_:_:)
func ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(addressBook ABAddressBookRef, source ABRecordRef, sortOrdering ABPersonSortOrdering) ArrayRef {
	return _ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(addressBook, source, sortOrdering)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering */

// Returns an array of all sources in the address book.

// Returns an array of all sources in the address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllSources(_:)
func ABAddressBookCopyArrayOfAllSources(addressBook ABAddressBookRef) ArrayRef {
	return _ABAddressBookCopyArrayOfAllSources(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookCopyArrayOfAllSources */

// Returns the default source.

// Returns the default source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyDefaultSource(_:)
func ABAddressBookCopyDefaultSource(addressBook ABAddressBookRef) ABRecordRef {
	return _ABAddressBookCopyDefaultSource(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookCopyDefaultSource */

// Returns a localized version of a record-property label.

// Returns a localized version of a record-property label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyLocalizedLabel(_:)
func ABAddressBookCopyLocalizedLabel(label StringRef) StringRef {
	return _ABAddressBookCopyLocalizedLabel(label)
}/* debug [functions.gen.go/function]: ABAddressBookCopyLocalizedLabel */

// Performs a prefix search on the composite names of people in an address book and returns an array of persons that match the search criteria.

// Performs a prefix search on the composite names of people in an address book and returns an array of persons that match the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyPeopleWithName(_:_:)
func ABAddressBookCopyPeopleWithName(addressBook ABAddressBookRef, name StringRef) ArrayRef {
	return _ABAddressBookCopyPeopleWithName(addressBook, name)
}/* debug [functions.gen.go/function]: ABAddressBookCopyPeopleWithName */

// Creates a new address book object with data from the Address Book database.

// Creates a new address book object with data from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCreate()
func ABAddressBookCreate() ABAddressBookRef {
	return _ABAddressBookCreate()
}/* debug [functions.gen.go/function]: ABAddressBookCreate */

// Creates a new address book object with data from the Address Book database.

// Creates a new address book object with data from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCreateWithOptions(_:_:)
func ABAddressBookCreateWithOptions(options DictionaryRef, error_ unsafe.Pointer) ABAddressBookRef {
	return _ABAddressBookCreateWithOptions(options, error_)
}/* debug [functions.gen.go/function]: ABAddressBookCreateWithOptions */

// Returns the authorization status of your app for accessing address book data.

// Returns the authorization status of your app for accessing address book data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetAuthorizationStatus()
func ABAddressBookGetAuthorizationStatus() ABAuthorizationStatus {
	return _ABAddressBookGetAuthorizationStatus()
}/* debug [functions.gen.go/function]: ABAddressBookGetAuthorizationStatus */

// Returns the number of groups in an address book.

// Returns the number of groups in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetGroupCount(_:)
func ABAddressBookGetGroupCount(addressBook ABAddressBookRef) Index {
	return _ABAddressBookGetGroupCount(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookGetGroupCount */

// Returns the group with a given record ID.

// Returns the group with a given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetGroupWithRecordID(_:_:)
func ABAddressBookGetGroupWithRecordID(addressBook ABAddressBookRef, recordID ABRecordID) ABRecordRef {
	return _ABAddressBookGetGroupWithRecordID(addressBook, recordID)
}/* debug [functions.gen.go/function]: ABAddressBookGetGroupWithRecordID */

// Returns the number of person records in an address book.

// Returns the number of person records in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetPersonCount(_:)
func ABAddressBookGetPersonCount(addressBook ABAddressBookRef) Index {
	return _ABAddressBookGetPersonCount(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookGetPersonCount */

// Returns the person record with a given record ID.

// Returns the person record with a given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetPersonWithRecordID(_:_:)
func ABAddressBookGetPersonWithRecordID(addressBook ABAddressBookRef, recordID ABRecordID) ABRecordRef {
	return _ABAddressBookGetPersonWithRecordID(addressBook, recordID)
}/* debug [functions.gen.go/function]: ABAddressBookGetPersonWithRecordID */

// Returns the source record with the given record ID.

// Returns the source record with the given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetSourceWithRecordID(_:_:)
func ABAddressBookGetSourceWithRecordID(addressBook ABAddressBookRef, sourceID ABRecordID) ABRecordRef {
	return _ABAddressBookGetSourceWithRecordID(addressBook, sourceID)
}/* debug [functions.gen.go/function]: ABAddressBookGetSourceWithRecordID */

// Indicates whether an address book has changes that have not been saved to the Address Book database.

// Indicates whether an address book has changes that have not been saved to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookHasUnsavedChanges(_:)
func ABAddressBookHasUnsavedChanges(addressBook ABAddressBookRef) bool {
	return _ABAddressBookHasUnsavedChanges(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookHasUnsavedChanges */

// Registers a callback to receive notifications when the Address Book database is modified.

// Registers a callback to receive notifications when the Address Book database is modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRegisterExternalChangeCallback(_:_:_:)
func ABAddressBookRegisterExternalChangeCallback(addressBook ABAddressBookRef, callback ABExternalChangeCallback, context unsafe.Pointer) {
	_ABAddressBookRegisterExternalChangeCallback(addressBook, callback, context)
}/* debug [functions.gen.go/function]: ABAddressBookRegisterExternalChangeCallback */

// Removes a record from an address book.

// Removes a record from an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRemoveRecord(_:_:_:)
func ABAddressBookRemoveRecord(addressBook ABAddressBookRef, record ABRecordRef, error_ unsafe.Pointer) bool {
	return _ABAddressBookRemoveRecord(addressBook, record, error_)
}/* debug [functions.gen.go/function]: ABAddressBookRemoveRecord */

// Requests access to address book data from the user.

// Requests access to address book data from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRequestAccessWithCompletion(_:_:)
func ABAddressBookRequestAccessWithCompletion(addressBook ABAddressBookRef, completion unsafe.Pointer) {
	_ABAddressBookRequestAccessWithCompletion(addressBook, completion)
}/* debug [functions.gen.go/function]: ABAddressBookRequestAccessWithCompletion */

// Discards unsaved changes in an address book.

// Discards unsaved changes in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRevert(_:)
func ABAddressBookRevert(addressBook ABAddressBookRef) {
	_ABAddressBookRevert(addressBook)
}/* debug [functions.gen.go/function]: ABAddressBookRevert */

// Saves any unsaved changes to the Address Book database.

// Saves any unsaved changes to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookSave(_:_:)
func ABAddressBookSave(addressBook ABAddressBookRef, error_ unsafe.Pointer) bool {
	return _ABAddressBookSave(addressBook, error_)
}/* debug [functions.gen.go/function]: ABAddressBookSave */

// Unregisters a callback.

// Unregisters a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookUnregisterExternalChangeCallback(_:_:_:)
func ABAddressBookUnregisterExternalChangeCallback(addressBook ABAddressBookRef, callback ABExternalChangeCallback, context unsafe.Pointer) {
	_ABAddressBookUnregisterExternalChangeCallback(addressBook, callback, context)
}/* debug [functions.gen.go/function]: ABAddressBookUnregisterExternalChangeCallback */

// Starts an asynchronous fetch for image data in all locations, and returns a non-zero tag for tracking.
//
// Added in macOS .
// Starts an asynchronous fetch for image data in all locations, and returns a non-zero tag for tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABBeginLoadingImageDataForClient(_:_:_:)
func ABBeginLoadingImageDataForClient(person ABPersonRef, callback ABImageClientCallback, refcon unsafe.Pointer) Index {
	return _ABBeginLoadingImageDataForClient(person, callback, refcon)
}/* debug [functions.gen.go/function]: ABBeginLoadingImageDataForClient */

// Cancels an asynchronous fetch of an image for the given tag.
//
// Added in macOS .
// Cancels an asynchronous fetch of an image for the given tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCancelLoadingImageDataForTag(_:)
func ABCancelLoadingImageDataForTag(tag Index) {
	_ABCancelLoadingImageDataForTag(tag)
}/* debug [functions.gen.go/function]: ABCancelLoadingImageDataForTag */

// Returns an array of all the groups in the Address Book database.
//
// Added in macOS .
// Returns an array of all the groups in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfAllGroups(_:)
func ABCopyArrayOfAllGroups(addressBook ABAddressBookRef) ArrayRef {
	return _ABCopyArrayOfAllGroups(addressBook)
}/* debug [functions.gen.go/function]: ABCopyArrayOfAllGroups */

// Returns an array of all the people in the Address Book database.
//
// Added in macOS .
// Returns an array of all the people in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfAllPeople(_:)
func ABCopyArrayOfAllPeople(addressBook ABAddressBookRef) ArrayRef {
	return _ABCopyArrayOfAllPeople(addressBook)
}/* debug [functions.gen.go/function]: ABCopyArrayOfAllPeople */

// Returns an array of records that match the given search element, or an empty array if no records match the search element.
//
// Added in macOS .
// Returns an array of records that match the given search element, or an empty array if no records match the search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfMatchingRecords(_:_:)
func ABCopyArrayOfMatchingRecords(addressBook ABAddressBookRef, search ABSearchElementRef) ArrayRef {
	return _ABCopyArrayOfMatchingRecords(addressBook, search)
}/* debug [functions.gen.go/function]: ABCopyArrayOfMatchingRecords */

// Returns an array containing the names of all the properties for the specified record type.
//
// Added in macOS .
// Returns an array containing the names of all the properties for the specified record type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfPropertiesForRecordType(_:_:)
func ABCopyArrayOfPropertiesForRecordType(addressBook ABAddressBookRef, recordType StringRef) ArrayRef {
	return _ABCopyArrayOfPropertiesForRecordType(addressBook, recordType)
}/* debug [functions.gen.go/function]: ABCopyArrayOfPropertiesForRecordType */

// Returns the default country code for records with unspecified country codes.
//
// Added in macOS 10.3.
// Returns the default country code for records with unspecified country codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyDefaultCountryCode(_:)
func ABCopyDefaultCountryCode(addressBook ABAddressBookRef) StringRef {
	return _ABCopyDefaultCountryCode(addressBook)
}/* debug [functions.gen.go/function]: ABCopyDefaultCountryCode */

// Returns the localized version of a built in property,label, or key.
//
// Added in macOS .
// Returns the localized version of a built in property,label, or key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyLocalizedPropertyOrLabel(_:)
func ABCopyLocalizedPropertyOrLabel(labelOrProperty StringRef) StringRef {
	return _ABCopyLocalizedPropertyOrLabel(labelOrProperty)
}/* debug [functions.gen.go/function]: ABCopyLocalizedPropertyOrLabel */

// Returns the record that matches the given unique ID.
//
// Added in macOS .
// Returns the record that matches the given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyRecordForUniqueId(_:_:)
func ABCopyRecordForUniqueId(addressBook ABAddressBookRef, uniqueId StringRef) ABRecordRef {
	return _ABCopyRecordForUniqueId(addressBook, uniqueId)
}/* debug [functions.gen.go/function]: ABCopyRecordForUniqueId */

// Returns the type name of the record that matches a given unique ID.
//
// Added in macOS 10.3.
// Returns the type name of the record that matches a given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyRecordTypeFromUniqueId(_:_:)
func ABCopyRecordTypeFromUniqueId(addressBook ABAddressBookRef, uniqueId StringRef) StringRef {
	return _ABCopyRecordTypeFromUniqueId(addressBook, uniqueId)
}/* debug [functions.gen.go/function]: ABCopyRecordTypeFromUniqueId */

// Returns a string containing the formatted address.
//
// Added in macOS 10.3.
// Returns a string containing the formatted address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCreateFormattedAddressFromDictionary(_:_:)
func ABCreateFormattedAddressFromDictionary(addressBook ABAddressBookRef, address DictionaryRef) StringRef {
	return _ABCreateFormattedAddressFromDictionary(addressBook, address)
}/* debug [functions.gen.go/function]: ABCreateFormattedAddressFromDictionary */

// Returns the ABPerson object for the logged-in user.
//
// Added in macOS .
// Returns the ABPerson object for the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGetMe(_:)
func ABGetMe(addressBook ABAddressBookRef) ABPersonRef {
	return _ABGetMe(addressBook)
}/* debug [functions.gen.go/function]: ABGetMe */

// Returns the unique shared ABAddressBook object.
//
// Added in macOS .
// Returns the unique shared ABAddressBook object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGetSharedAddressBook()
func ABGetSharedAddressBook() ABAddressBookRef {
	return _ABGetSharedAddressBook()
}/* debug [functions.gen.go/function]: ABGetSharedAddressBook */

// Adds a subgroup to another group.
//
// Added in macOS .
// Adds a subgroup to another group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupAddGroup(_:_:)
func ABGroupAddGroup(group ABGroupRef, groupToAdd ABGroupRef) bool {
	return _ABGroupAddGroup(group, groupToAdd)
}/* debug [functions.gen.go/function]: ABGroupAddGroup */

// Adds a person to a group.
//
// Added in macOS .
// Adds a person to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupAddMember(_:_:_:)
func ABGroupAddMember(group ABRecordRef, person ABRecordRef, error_ unsafe.Pointer) bool {
	return _ABGroupAddMember(group, person, error_)
}/* debug [functions.gen.go/function]: ABGroupAddMember */

// Returns an array of persons in a group.
//
// Added in macOS .
// Returns an array of persons in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllMembers(_:)
func ABGroupCopyArrayOfAllMembers(group ABGroupRef) ArrayRef {
	return _ABGroupCopyArrayOfAllMembers(group)
}/* debug [functions.gen.go/function]: ABGroupCopyArrayOfAllMembers */

// Returns the records in a group, using a sort ordering.

// Returns the records in a group, using a sort ordering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllMembersWithSortOrdering(_:_:)
func ABGroupCopyArrayOfAllMembersWithSortOrdering(group ABRecordRef, sortOrdering ABPersonSortOrdering) ArrayRef {
	return _ABGroupCopyArrayOfAllMembersWithSortOrdering(group, sortOrdering)
}/* debug [functions.gen.go/function]: ABGroupCopyArrayOfAllMembersWithSortOrdering */

// Returns an array containing a group’s subgroups.
//
// Added in macOS .
// Returns an array containing a group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllSubgroups(_:)
func ABGroupCopyArrayOfAllSubgroups(group ABGroupRef) ArrayRef {
	return _ABGroupCopyArrayOfAllSubgroups(group)
}/* debug [functions.gen.go/function]: ABGroupCopyArrayOfAllSubgroups */

// Returns the distribution identifier for the given propertyand person.
//
// Added in macOS .
// Returns the distribution identifier for the given propertyand person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyDistributionIdentifier(_:_:_:)
func ABGroupCopyDistributionIdentifier(group ABGroupRef, person ABPersonRef, property StringRef) StringRef {
	return _ABGroupCopyDistributionIdentifier(group, person, property)
}/* debug [functions.gen.go/function]: ABGroupCopyDistributionIdentifier */

// Returns an array containing a group’s parents—thegroups that a group belongs to.
//
// Added in macOS .
// Returns an array containing a group’s parents—thegroups that a group belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyParentGroups(_:)
func ABGroupCopyParentGroups(group ABGroupRef) ArrayRef {
	return _ABGroupCopyParentGroups(group)
}/* debug [functions.gen.go/function]: ABGroupCopyParentGroups */

// Returns the source that the group is from.

// Returns the source that the group is from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopySource(_:)
func ABGroupCopySource(group ABRecordRef) ABRecordRef {
	return _ABGroupCopySource(group)
}/* debug [functions.gen.go/function]: ABGroupCopySource */

// Returns a new ABGroup object.
//
// Added in macOS .
// Returns a new ABGroup object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreate()
func ABGroupCreate() ABGroupRef {
	return _ABGroupCreate()
}/* debug [functions.gen.go/function]: ABGroupCreate */

// Creates a group in a particular source.

// Creates a group in a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreateInSource(_:)
func ABGroupCreateInSource(source ABRecordRef) ABRecordRef {
	return _ABGroupCreateInSource(source)
}/* debug [functions.gen.go/function]: ABGroupCreateInSource */

// Creates an ABSearchElement object that specifies a queryfor ABGroup records.
//
// Added in macOS .
// Creates an ABSearchElement object that specifies a queryfor ABGroup records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreateSearchElement(_:_:_:_:_:)
func ABGroupCreateSearchElement(property StringRef, label StringRef, key StringRef, value TypeRef, comparison ABSearchComparison) ABSearchElementRef {
	return _ABGroupCreateSearchElement(property, label, key, value, comparison)
}/* debug [functions.gen.go/function]: ABGroupCreateSearchElement */

// Removes a subgroup from a group.
//
// Added in macOS .
// Removes a subgroup from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupRemoveGroup(_:_:)
func ABGroupRemoveGroup(group ABGroupRef, groupToRemove ABGroupRef) bool {
	return _ABGroupRemoveGroup(group, groupToRemove)
}/* debug [functions.gen.go/function]: ABGroupRemoveGroup */

// Removes a person from a group.
//
// Added in macOS .
// Removes a person from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupRemoveMember(_:_:_:)
func ABGroupRemoveMember(group ABGroupRef, personToRemove ABPersonRef) bool {
	return _ABGroupRemoveMember(group, personToRemove)
}/* debug [functions.gen.go/function]: ABGroupRemoveMember */

// Assigning a specific distribution identifier for a person’smulti-value list property so that the group can be used as a distributionlist (mailing list, in the case of an email property).
//
// Added in macOS .
// Assigning a specific distribution identifier for a person’smulti-value list property so that the group can be used as a distributionlist (mailing list, in the case of an email property).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupSetDistributionIdentifier(_:_:_:_:)
func ABGroupSetDistributionIdentifier(group ABGroupRef, person ABPersonRef, property StringRef, identifier StringRef) bool {
	return _ABGroupSetDistributionIdentifier(group, person, property, identifier)
}/* debug [functions.gen.go/function]: ABGroupSetDistributionIdentifier */

// Returns whether if there are unsaved changes in the address book.
//
// Added in macOS .
// Returns whether if there are unsaved changes in the address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABHasUnsavedChanges(_:)
func ABHasUnsavedChanges(addressBook ABAddressBookRef) bool {
	return _ABHasUnsavedChanges(addressBook)
}/* debug [functions.gen.go/function]: ABHasUnsavedChanges */

// Returns the localized version of a built in property, label, or key.
//
// Added in macOS .
// Returns the localized version of a built in property, label, or key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABLocalizedPropertyOrLabel(_:)
func ABLocalizedPropertyOrLabel(propertyOrLabel unsafe.Pointer) unsafe.Pointer {
	return _ABLocalizedPropertyOrLabel(propertyOrLabel)
}/* debug [functions.gen.go/function]: ABLocalizedPropertyOrLabel */

// Adds a value and its label to a multi-value list.
//
// Added in macOS .
// Adds a value and its label to a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueAdd(_:_:_:_:)
func ABMultiValueAdd(multiValue ABMutableMultiValueRef, value TypeRef, label StringRef, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueAdd(multiValue, value, label, outIdentifier)
}/* debug [functions.gen.go/function]: ABMultiValueAdd */

// Adds a value and its corresponding label to a multivalue property.

// Adds a value and its corresponding label to a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueAddValueAndLabel(_:_:_:_:)
func ABMultiValueAddValueAndLabel(multiValue ABMutableMultiValueRef, value TypeRef, label StringRef, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueAddValueAndLabel(multiValue, value, label, outIdentifier)
}/* debug [functions.gen.go/function]: ABMultiValueAddValueAndLabel */

// Returns an array with the values in a multivalue property.

// Returns an array with the values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyArrayOfAllValues(_:)
func ABMultiValueCopyArrayOfAllValues(multiValue ABMultiValueRef) ArrayRef {
	return _ABMultiValueCopyArrayOfAllValues(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueCopyArrayOfAllValues */

// Returns the identifier at the given index.
//
// Added in macOS .
// Returns the identifier at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyIdentifierAtIndex(_:_:)
func ABMultiValueCopyIdentifierAtIndex(multiValue ABMultiValueRef, index Index) StringRef {
	return _ABMultiValueCopyIdentifierAtIndex(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueCopyIdentifierAtIndex */

// Returns the label for the given index.
//
// Added in macOS .
// Returns the label for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyLabelAtIndex(_:_:)
func ABMultiValueCopyLabelAtIndex(multiValue ABMultiValueRef, index Index) StringRef {
	return _ABMultiValueCopyLabelAtIndex(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueCopyLabelAtIndex */

// Returns the identifier for the primary value.
//
// Added in macOS .
// Returns the identifier for the primary value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyPrimaryIdentifier(_:)
func ABMultiValueCopyPrimaryIdentifier(multiValue ABMultiValueRef) StringRef {
	return _ABMultiValueCopyPrimaryIdentifier(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueCopyPrimaryIdentifier */

// Returns the value for the given index.
//
// Added in macOS .
// Returns the value for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyValueAtIndex(_:_:)
func ABMultiValueCopyValueAtIndex(multiValue ABMultiValueRef, index Index) TypeRef {
	return _ABMultiValueCopyValueAtIndex(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueCopyValueAtIndex */

// Returns the number of entries in a multi-value list.
//
// Added in macOS .
// Returns the number of entries in a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCount(_:)
func ABMultiValueCount(multiValue ABMultiValueRef) Index {
	return _ABMultiValueCount(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueCount */

// Returns a new ABMultiValue object.
//
// Added in macOS .
// Returns a new ABMultiValue object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreate()
func ABMultiValueCreate() ABMultiValueRef {
	return _ABMultiValueCreate()
}/* debug [functions.gen.go/function]: ABMultiValueCreate */

// Returns a copy of a multi-value object.
//
// Added in macOS .
// Returns a copy of a multi-value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateCopy(_:)
func ABMultiValueCreateCopy(multiValue ABMultiValueRef) ABMultiValueRef {
	return _ABMultiValueCreateCopy(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueCreateCopy */

// Returns a newly created mutable multi-value list object.
//
// Added in macOS .
// Returns a newly created mutable multi-value list object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateMutable(_:)
func ABMultiValueCreateMutable() ABMutableMultiValueRef {
	return _ABMultiValueCreateMutable()
}/* debug [functions.gen.go/function]: ABMultiValueCreateMutable */

// Returns a mutable copy of a multi-value object.
//
// Added in macOS .
// Returns a mutable copy of a multi-value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateMutableCopy(_:)
func ABMultiValueCreateMutableCopy(multiValue ABMultiValueRef) ABMutableMultiValueRef {
	return _ABMultiValueCreateMutableCopy(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueCreateMutableCopy */

// Returns the number of values in a multivalue property.

// Returns the number of values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetCount(_:)
func ABMultiValueGetCount(multiValue ABMultiValueRef) Index {
	return _ABMultiValueGetCount(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueGetCount */

// Returns the first location of a value in a multivalue property.

// Returns the first location of a value in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetFirstIndexOfValue(_:_:)
func ABMultiValueGetFirstIndexOfValue(multiValue ABMultiValueRef, value TypeRef) Index {
	return _ABMultiValueGetFirstIndexOfValue(multiValue, value)
}/* debug [functions.gen.go/function]: ABMultiValueGetFirstIndexOfValue */

// Returns the identifier of a value in a multivalue property.

// Returns the identifier of a value in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetIdentifierAtIndex(_:_:)
func ABMultiValueGetIdentifierAtIndex(multiValue ABMultiValueRef, index Index) ABMultiValueIdentifier {
	return _ABMultiValueGetIdentifierAtIndex(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueGetIdentifierAtIndex */

// Returns the location (within a multivalue property) of a value with a given identifier.

// Returns the location (within a multivalue property) of a value with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetIndexForIdentifier(_:_:)
func ABMultiValueGetIndexForIdentifier(multiValue ABMultiValueRef, identifier ABMultiValueIdentifier) Index {
	return _ABMultiValueGetIndexForIdentifier(multiValue, identifier)
}/* debug [functions.gen.go/function]: ABMultiValueGetIndexForIdentifier */

// Returns the type of the values contained in a multivalue property.

// Returns the type of the values contained in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetPropertyType(_:)
func ABMultiValueGetPropertyType(multiValue ABMultiValueRef) ABPropertyType {
	return _ABMultiValueGetPropertyType(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValueGetPropertyType */

// Returns the index for the given identifier.
//
// Added in macOS .
// Returns the index for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueIndexForIdentifier(_:_:)
func ABMultiValueIndexForIdentifier(multiValue ABMultiValueRef, identifier StringRef) Index {
	return _ABMultiValueIndexForIdentifier(multiValue, identifier)
}/* debug [functions.gen.go/function]: ABMultiValueIndexForIdentifier */

// Inserts a value and its label at the given index in amulti-value list.
//
// Added in macOS .
// Inserts a value and its label at the given index in amulti-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueInsert(_:_:_:_:_:)
func ABMultiValueInsert(multiValue ABMutableMultiValueRef, value TypeRef, label StringRef, index Index, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueInsert(multiValue, value, label, index, outIdentifier)
}/* debug [functions.gen.go/function]: ABMultiValueInsert */

// Inserts a value and a label into a multivalue property.

// Inserts a value and a label into a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueInsertValueAndLabelAtIndex(_:_:_:_:_:)
func ABMultiValueInsertValueAndLabelAtIndex(multiValue ABMutableMultiValueRef, value TypeRef, label StringRef, index Index, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueInsertValueAndLabelAtIndex(multiValue, value, label, index, outIdentifier)
}/* debug [functions.gen.go/function]: ABMultiValueInsertValueAndLabelAtIndex */

// Returns the type for the values in a multi-value list.
//
// Added in macOS .
// Returns the type for the values in a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValuePropertyType(_:)
func ABMultiValuePropertyType(multiValue ABMultiValueRef) ABPropertyType {
	return _ABMultiValuePropertyType(multiValue)
}/* debug [functions.gen.go/function]: ABMultiValuePropertyType */

// Removes the value and label at the given index.
//
// Added in macOS .
// Removes the value and label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueRemove(_:_:)
func ABMultiValueRemove(multiValue ABMutableMultiValueRef, index Index) bool {
	return _ABMultiValueRemove(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueRemove */

// Removes a value from a multivalue property.

// Removes a value from a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueRemoveValueAndLabelAtIndex(_:_:)
func ABMultiValueRemoveValueAndLabelAtIndex(multiValue ABMutableMultiValueRef, index Index) bool {
	return _ABMultiValueRemoveValueAndLabelAtIndex(multiValue, index)
}/* debug [functions.gen.go/function]: ABMultiValueRemoveValueAndLabelAtIndex */

// Replaces the label at the given index.
//
// Added in macOS .
// Replaces the label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceLabel(_:_:_:)
func ABMultiValueReplaceLabel(multiValue ABMutableMultiValueRef, label StringRef, index Index) bool {
	return _ABMultiValueReplaceLabel(multiValue, label, index)
}/* debug [functions.gen.go/function]: ABMultiValueReplaceLabel */

// Replaces a label in a multivalue property with another label.

// Replaces a label in a multivalue property with another label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceLabelAtIndex(_:_:_:)
func ABMultiValueReplaceLabelAtIndex(multiValue ABMutableMultiValueRef, label StringRef, index Index) bool {
	return _ABMultiValueReplaceLabelAtIndex(multiValue, label, index)
}/* debug [functions.gen.go/function]: ABMultiValueReplaceLabelAtIndex */

// Replaces the value at the given index.
//
// Added in macOS .
// Replaces the value at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceValue(_:_:_:)
func ABMultiValueReplaceValue(multiValue ABMutableMultiValueRef, value TypeRef, index Index) bool {
	return _ABMultiValueReplaceValue(multiValue, value, index)
}/* debug [functions.gen.go/function]: ABMultiValueReplaceValue */

// Replaces a value in a multivalue property with another value.

// Replaces a value in a multivalue property with another value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceValueAtIndex(_:_:_:)
func ABMultiValueReplaceValueAtIndex(multiValue ABMutableMultiValueRef, value TypeRef, index Index) bool {
	return _ABMultiValueReplaceValueAtIndex(multiValue, value, index)
}/* debug [functions.gen.go/function]: ABMultiValueReplaceValueAtIndex */

// Sets the primary value to be the value for the given identifier.
//
// Added in macOS .
// Sets the primary value to be the value for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueSetPrimaryIdentifier(_:_:)
func ABMultiValueSetPrimaryIdentifier(multiValue ABMutableMultiValueRef, identifier StringRef) bool {
	return _ABMultiValueSetPrimaryIdentifier(multiValue, identifier)
}/* debug [functions.gen.go/function]: ABMultiValueSetPrimaryIdentifier */

// Indicates how two person records get sorted.

// Indicates how two person records get sorted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonComparePeopleByName(_:_:_:)
func ABPersonComparePeopleByName(person1 ABRecordRef, person2 ABRecordRef, ordering ABPersonSortOrdering) ComparisonResult {
	return _ABPersonComparePeopleByName(person1, person2, ordering)
}/* debug [functions.gen.go/function]: ABPersonComparePeopleByName */

// Returns an array of all person records in the address book database that are linked to the given person record.

// Returns an array of all person records in the address book database that are linked to the given person record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyArrayOfAllLinkedPeople(_:)
func ABPersonCopyArrayOfAllLinkedPeople(person ABRecordRef) ArrayRef {
	return _ABPersonCopyArrayOfAllLinkedPeople(person)
}/* debug [functions.gen.go/function]: ABPersonCopyArrayOfAllLinkedPeople */

// Returns the delimiter to use between name components.

// Returns the delimiter to use between name components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyCompositeNameDelimiterForRecord(_:)
func ABPersonCopyCompositeNameDelimiterForRecord(record ABRecordRef) StringRef {
	return _ABPersonCopyCompositeNameDelimiterForRecord(record)
}/* debug [functions.gen.go/function]: ABPersonCopyCompositeNameDelimiterForRecord */

// Returns data that contains a picture of a person.
//
// Added in macOS .
// Returns data that contains a picture of a person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyImageData(_:)
func ABPersonCopyImageData(person ABPersonRef) DataRef {
	return _ABPersonCopyImageData(person)
}/* debug [functions.gen.go/function]: ABPersonCopyImageData */

// Returns the picture for a person record in the given format.

// Returns the picture for a person record in the given format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyImageDataWithFormat(_:_:)
func ABPersonCopyImageDataWithFormat(person ABRecordRef, format unsafe.Pointer) DataRef {
	return _ABPersonCopyImageDataWithFormat(person, format)
}/* debug [functions.gen.go/function]: ABPersonCopyImageDataWithFormat */

// Returns the localized name of a person property

// Returns the localized name of a person property
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyLocalizedPropertyName(_:)
func ABPersonCopyLocalizedPropertyName(property ABPropertyID) StringRef {
	return _ABPersonCopyLocalizedPropertyName(property)
}/* debug [functions.gen.go/function]: ABPersonCopyLocalizedPropertyName */

// Returns an array of groups that a person belongs to.
//
// Added in macOS .
// Returns an array of groups that a person belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyParentGroups(_:)
func ABPersonCopyParentGroups(person ABPersonRef) ArrayRef {
	return _ABPersonCopyParentGroups(person)
}/* debug [functions.gen.go/function]: ABPersonCopyParentGroups */

// Returns the source that the person record is from.

// Returns the source that the person record is from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopySource(_:)
func ABPersonCopySource(person ABRecordRef) ABRecordRef {
	return _ABPersonCopySource(person)
}/* debug [functions.gen.go/function]: ABPersonCopySource */

// Returns the vCard representation of the person as a data object in vCard format.
//
// Added in macOS .
// Returns the vCard representation of the person as a data object in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyVCardRepresentation(_:)
func ABPersonCopyVCardRepresentation(person ABPersonRef) DataRef {
	return _ABPersonCopyVCardRepresentation(person)
}/* debug [functions.gen.go/function]: ABPersonCopyVCardRepresentation */

// Returns a newly created person object.
//
// Added in macOS .
// Returns a newly created person object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreate()
func ABPersonCreate() ABRecordRef {
	return _ABPersonCreate()
}/* debug [functions.gen.go/function]: ABPersonCreate */

// Creates a new person record in a particular source.

// Creates a new person record in a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateInSource(_:)
func ABPersonCreateInSource(source ABRecordRef) ABRecordRef {
	return _ABPersonCreateInSource(source)
}/* debug [functions.gen.go/function]: ABPersonCreateInSource */

// Creates person records from the given vCard representation.

// Creates person records from the given vCard representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreatePeopleInSourceWithVCardRepresentation(_:_:)
func ABPersonCreatePeopleInSourceWithVCardRepresentation(source ABRecordRef, vCardData DataRef) ArrayRef {
	return _ABPersonCreatePeopleInSourceWithVCardRepresentation(source, vCardData)
}/* debug [functions.gen.go/function]: ABPersonCreatePeopleInSourceWithVCardRepresentation */

// Returns a search element object that specifies a query for records of this type.
//
// Added in macOS .
// Returns a search element object that specifies a query for records of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateSearchElement(_:_:_:_:_:)
func ABPersonCreateSearchElement(property StringRef, label StringRef, key StringRef, value TypeRef, comparison ABSearchComparison) ABSearchElementRef {
	return _ABPersonCreateSearchElement(property, label, key, value, comparison)
}/* debug [functions.gen.go/function]: ABPersonCreateSearchElement */

// Returns the vCard representation of the given person records.

// Returns the vCard representation of the given person records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateVCardRepresentationWithPeople(_:)
func ABPersonCreateVCardRepresentationWithPeople(people ArrayRef) DataRef {
	return _ABPersonCreateVCardRepresentationWithPeople(people)
}/* debug [functions.gen.go/function]: ABPersonCreateVCardRepresentationWithPeople */

// Returns a new ABPerson object initialized with the given data in vCard format.
//
// Added in macOS .
// Returns a new ABPerson object initialized with the given data in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateWithVCardRepresentation(_:)
func ABPersonCreateWithVCardRepresentation(vCard DataRef) ABPersonRef {
	return _ABPersonCreateWithVCardRepresentation(vCard)
}/* debug [functions.gen.go/function]: ABPersonCreateWithVCardRepresentation */

// Returns the person-name display format.

// Returns the person-name display format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetCompositeNameFormat()
func ABPersonGetCompositeNameFormat() ABPersonCompositeNameFormat {
	return _ABPersonGetCompositeNameFormat()
}/* debug [functions.gen.go/function]: ABPersonGetCompositeNameFormat */

// Returns the person-name display format to use for the given record.

// Returns the person-name display format to use for the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetCompositeNameFormatForRecord(_:)
func ABPersonGetCompositeNameFormatForRecord(record ABRecordRef) ABPersonCompositeNameFormat {
	return _ABPersonGetCompositeNameFormatForRecord(record)
}/* debug [functions.gen.go/function]: ABPersonGetCompositeNameFormatForRecord */

// Returns the user’s sort-ordering preference for lists of persons.

// Returns the user’s sort-ordering preference for lists of persons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetSortOrdering()
func ABPersonGetSortOrdering() ABPersonSortOrdering {
	return _ABPersonGetSortOrdering()
}/* debug [functions.gen.go/function]: ABPersonGetSortOrdering */

// Returns the type of a person property.

// Returns the type of a person property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetTypeOfProperty(_:)
func ABPersonGetTypeOfProperty(property ABPropertyID) ABPropertyType {
	return _ABPersonGetTypeOfProperty(property)
}/* debug [functions.gen.go/function]: ABPersonGetTypeOfProperty */

// Indicates whether a person has a picture.

// Indicates whether a person has a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonHasImageData(_:)
func ABPersonHasImageData(person ABRecordRef) bool {
	return _ABPersonHasImageData(person)
}/* debug [functions.gen.go/function]: ABPersonHasImageData */

// Removes a person’s picture.

// Removes a person’s picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonRemoveImageData(_:_:)
func ABPersonRemoveImageData(person ABRecordRef, error_ unsafe.Pointer) bool {
	return _ABPersonRemoveImageData(person, error_)
}/* debug [functions.gen.go/function]: ABPersonRemoveImageData */

// Sets the image for this person to the given data.
//
// Added in macOS .
// Sets the image for this person to the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonSetImageData(_:_:_:)
func ABPersonSetImageData(person ABPersonRef, imageData DataRef) bool {
	return _ABPersonSetImageData(person, imageData)
}/* debug [functions.gen.go/function]: ABPersonSetImageData */

// Adds a property to the group of properties available in the record list. Use to remove a property from the list and to obtain the list of properties available in the list.
//
// Added in macOS 10.3.
// Adds a property to the group of properties available in the record list. Use to remove a property from the list and to obtain the list of properties available in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerAddProperty
func ABPickerAddProperty(inPicker unsafe.Pointer, inProperty StringRef) {
	_ABPickerAddProperty(inPicker, inProperty)
}/* debug [functions.gen.go/function]: ABPickerAddProperty */

// Specifies the selection behaviors for a people-picker window. Use to obtain the selection behaviors specified for the window.
//
// Added in macOS 10.3.
// Specifies the selection behaviors for a people-picker window. Use to obtain the selection behaviors specified for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerChangeAttributes
func ABPickerChangeAttributes(inPicker unsafe.Pointer, inAttributesToSet unsafe.Pointer, inAttributesToClear unsafe.Pointer) {
	_ABPickerChangeAttributes(inPicker, inAttributesToSet, inAttributesToClear)
}/* debug [functions.gen.go/function]: ABPickerChangeAttributes */

// Clears the search field and resets the list of displayed records.
//
// Added in macOS 10.3.
// Clears the search field and resets the list of displayed records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerClearSearchField
func ABPickerClearSearchField(inPicker unsafe.Pointer) {
	_ABPickerClearSearchField(inPicker)
}/* debug [functions.gen.go/function]: ABPickerClearSearchField */

// Obtains the title of a custom property.
//
// Added in macOS 10.3.
// Obtains the title of a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyColumnTitle
func ABPickerCopyColumnTitle(inPicker unsafe.Pointer, inProperty StringRef) StringRef {
	return _ABPickerCopyColumnTitle(inPicker, inProperty)
}/* debug [functions.gen.go/function]: ABPickerCopyColumnTitle */

// Returns the name of the property currently displayed in the record list.
//
// Added in macOS 10.3.
// Returns the name of the property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyDisplayedProperty
func ABPickerCopyDisplayedProperty(inPicker unsafe.Pointer) StringRef {
	return _ABPickerCopyDisplayedProperty(inPicker)
}/* debug [functions.gen.go/function]: ABPickerCopyDisplayedProperty */

// Obtains the list of properties available in the record list. Use to add a property to the record list and to remove a property from the list.
//
// Added in macOS 10.3.
// Obtains the list of properties available in the record list. Use to add a property to the record list and to remove a property from the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyProperties
func ABPickerCopyProperties(inPicker unsafe.Pointer) ArrayRef {
	return _ABPickerCopyProperties(inPicker)
}/* debug [functions.gen.go/function]: ABPickerCopyProperties */

// Returns the groups selected in the group list as an array of objects.
//
// Added in macOS 10.3.
// Returns the groups selected in the group list as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedGroups
func ABPickerCopySelectedGroups(inPicker unsafe.Pointer) ArrayRef {
	return _ABPickerCopySelectedGroups(inPicker)
}/* debug [functions.gen.go/function]: ABPickerCopySelectedGroups */

// Returns the identifiers of the selected values in a multi-value property or an empty array if the property displayed is a single-value property.
//
// Added in macOS 10.3.
// Returns the identifiers of the selected values in a multi-value property or an empty array if the property displayed is a single-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedIdentifiers
func ABPickerCopySelectedIdentifiers(inPicker unsafe.Pointer, inPerson ABPersonRef) ArrayRef {
	return _ABPickerCopySelectedIdentifiers(inPicker, inPerson)
}/* debug [functions.gen.go/function]: ABPickerCopySelectedIdentifiers */

// Returns the selection in the record list as an array of ABGroup or objects.
//
// Added in macOS 10.3.
// Returns the selection in the record list as an array of ABGroup or objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedRecords
func ABPickerCopySelectedRecords(inPicker unsafe.Pointer) ArrayRef {
	return _ABPickerCopySelectedRecords(inPicker)
}/* debug [functions.gen.go/function]: ABPickerCopySelectedRecords */

// Returns the selected values in a multi-value property or an empty array if no values are selected or the property displayedis a single-value property.
//
// Added in macOS 10.3.
// Returns the selected values in a multi-value property or an empty array if no values are selected or the property displayedis a single-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedValues
func ABPickerCopySelectedValues(inPicker unsafe.Pointer) ArrayRef {
	return _ABPickerCopySelectedValues(inPicker)
}/* debug [functions.gen.go/function]: ABPickerCopySelectedValues */

// Creates an ABPickerRef. The corresponding window is hidden. Invoke to show it. Release with .
//
// Added in macOS 10.3.
// Creates an ABPickerRef. The corresponding window is hidden. Invoke to show it. Release with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCreate
func ABPickerCreate() unsafe.Pointer {
	return _ABPickerCreate()
}/* debug [functions.gen.go/function]: ABPickerCreate */

// Deselects all selected groups, records, and values in multi-value properties.
//
// Added in macOS 10.3.
// Deselects all selected groups, records, and values in multi-value properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectAll
func ABPickerDeselectAll(inPicker unsafe.Pointer) {
	_ABPickerDeselectAll(inPicker)
}/* debug [functions.gen.go/function]: ABPickerDeselectAll */

// Deselects a group in the group list.
//
// Added in macOS 10.3.
// Deselects a group in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectGroup
func ABPickerDeselectGroup(inPicker unsafe.Pointer, inGroup ABGroupRef) {
	_ABPickerDeselectGroup(inPicker, inGroup)
}/* debug [functions.gen.go/function]: ABPickerDeselectGroup */

// Deselects a value in multi-value property currently displayed in the record list.
//
// Added in macOS 10.3.
// Deselects a value in multi-value property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectIdentifier
func ABPickerDeselectIdentifier(inPicker unsafe.Pointer, inPerson ABPersonRef, inIdentifier StringRef) {
	_ABPickerDeselectIdentifier(inPicker, inPerson, inIdentifier)
}/* debug [functions.gen.go/function]: ABPickerDeselectIdentifier */

// Deselects a group in the record list.
//
// Added in macOS 10.3.
// Deselects a group in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectRecord
func ABPickerDeselectRecord(inPicker unsafe.Pointer, inRecord ABRecordRef) {
	_ABPickerDeselectRecord(inPicker, inRecord)
}/* debug [functions.gen.go/function]: ABPickerDeselectRecord */

// Launches Address Book to edit the item selected in the people-picker window.
//
// Added in macOS 10.3.
// Launches Address Book to edit the item selected in the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerEditInAddressBook
func ABPickerEditInAddressBook(inPicker unsafe.Pointer) {
	_ABPickerEditInAddressBook(inPicker)
}/* debug [functions.gen.go/function]: ABPickerEditInAddressBook */

// Indicates the selection behaviors selected a people-picker window. Use tospecify selection behaviors for the window.
//
// Added in macOS 10.3.
// Indicates the selection behaviors selected a people-picker window. Use tospecify selection behaviors for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetAttributes
func ABPickerGetAttributes(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerGetAttributes(inPicker)
}/* debug [functions.gen.go/function]: ABPickerGetAttributes */

// Obtains the delegate for a people-picker window.
//
// Added in macOS 10.3.
// Obtains the delegate for a people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetDelegate
func ABPickerGetDelegate(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerGetDelegate(inPicker)
}/* debug [functions.gen.go/function]: ABPickerGetDelegate */

// Returns the position and size of the people-picker window.
//
// Added in macOS 10.3.
// Returns the position and size of the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetFrame
func ABPickerGetFrame(inPicker unsafe.Pointer, outFrame unsafe.Pointer) {
	_ABPickerGetFrame(inPicker, outFrame)
}/* debug [functions.gen.go/function]: ABPickerGetFrame */

// Indicates whether the people-picker window is visible.
//
// Added in macOS 10.3.
// Indicates whether the people-picker window is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerIsVisible
func ABPickerIsVisible(inPicker unsafe.Pointer) bool {
	return _ABPickerIsVisible(inPicker)
}/* debug [functions.gen.go/function]: ABPickerIsVisible */

// Removes a property from the group of properties whose values are shown in the record list. Use to add a property to the record list and to obtain the list of properties shown in the record list.
//
// Added in macOS 10.3.
// Removes a property from the group of properties whose values are shown in the record list. Use to add a property to the record list and to obtain the list of properties shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerRemoveProperty
func ABPickerRemoveProperty(inPicker unsafe.Pointer, inProperty StringRef) {
	_ABPickerRemoveProperty(inPicker, inProperty)
}/* debug [functions.gen.go/function]: ABPickerRemoveProperty */

// Selects a group or a set of groups in the group list.
//
// Added in macOS 10.3.
// Selects a group or a set of groups in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectGroup
func ABPickerSelectGroup(inPicker unsafe.Pointer, inGroup ABGroupRef, inExtendSelection bool) {
	_ABPickerSelectGroup(inPicker, inGroup, inExtendSelection)
}/* debug [functions.gen.go/function]: ABPickerSelectGroup */

// Selects a value or a set of values in a multi-value property.
//
// Added in macOS 10.3.
// Selects a value or a set of values in a multi-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectIdentifier
func ABPickerSelectIdentifier(inPicker unsafe.Pointer, inPerson ABPersonRef, inIdentifier StringRef, inExtendSelection bool) {
	_ABPickerSelectIdentifier(inPicker, inPerson, inIdentifier, inExtendSelection)
}/* debug [functions.gen.go/function]: ABPickerSelectIdentifier */

// Launches Address Book and selects the item selected in the people-picker window.
//
// Added in macOS 10.3.
// Launches Address Book and selects the item selected in the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectInAddressBook
func ABPickerSelectInAddressBook(inPicker unsafe.Pointer) {
	_ABPickerSelectInAddressBook(inPicker)
}/* debug [functions.gen.go/function]: ABPickerSelectInAddressBook */

// Selects a record or a set of records in the record list.
//
// Added in macOS 10.3.
// Selects a record or a set of records in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectRecord
func ABPickerSelectRecord(inPicker unsafe.Pointer, inRecord ABRecordRef, inExtendSelection bool) {
	_ABPickerSelectRecord(inPicker, inRecord, inExtendSelection)
}/* debug [functions.gen.go/function]: ABPickerSelectRecord */

// Sets the title for a custom property.
//
// Added in macOS 10.3.
// Sets the title for a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetColumnTitle
func ABPickerSetColumnTitle(inPicker unsafe.Pointer, inTitle StringRef, inProperty StringRef) {
	_ABPickerSetColumnTitle(inPicker, inTitle, inProperty)
}/* debug [functions.gen.go/function]: ABPickerSetColumnTitle */

// Sets the event handler for people-picker events.
//
// Added in macOS 10.3.
// Sets the event handler for people-picker events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetDelegate
func ABPickerSetDelegate(inPicker unsafe.Pointer, inDelegate unsafe.Pointer) {
	_ABPickerSetDelegate(inPicker, inDelegate)
}/* debug [functions.gen.go/function]: ABPickerSetDelegate */

// Displays one of the properties whose values are shownin the record list.
//
// Added in macOS 10.3.
// Displays one of the properties whose values are shownin the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetDisplayedProperty
func ABPickerSetDisplayedProperty(inPicker unsafe.Pointer, inProperty StringRef) {
	_ABPickerSetDisplayedProperty(inPicker, inProperty)
}/* debug [functions.gen.go/function]: ABPickerSetDisplayedProperty */

// Specifies the position and size of the people-picker window.
//
// Added in macOS 10.3.
// Specifies the position and size of the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetFrame
func ABPickerSetFrame(inPicker unsafe.Pointer, inFrame unsafe.Pointer) {
	_ABPickerSetFrame(inPicker, inFrame)
}/* debug [functions.gen.go/function]: ABPickerSetFrame */

// Shows or hides a people-picker window.
//
// Added in macOS 10.3.
// Shows or hides a people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetVisibility
func ABPickerSetVisibility(inPicker unsafe.Pointer, visible bool) {
	_ABPickerSetVisibility(inPicker, visible)
}/* debug [functions.gen.go/function]: ABPickerSetVisibility */

// Returns an appropriate, human-friendly name for the record.

// Returns an appropriate, human-friendly name for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyCompositeName(_:)
func ABRecordCopyCompositeName(record ABRecordRef) StringRef {
	return _ABRecordCopyCompositeName(record)
}/* debug [functions.gen.go/function]: ABRecordCopyCompositeName */

// Returns the type of the given record.
//
// Added in macOS .
// Returns the type of the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyRecordType(_:)
func ABRecordCopyRecordType(record ABRecordRef) StringRef {
	return _ABRecordCopyRecordType(record)
}/* debug [functions.gen.go/function]: ABRecordCopyRecordType */

// Returns the unique ID of the receiver.
//
// Added in macOS .
// Returns the unique ID of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyUniqueId(_:)
func ABRecordCopyUniqueId(record ABRecordRef) StringRef {
	return _ABRecordCopyUniqueId(record)
}/* debug [functions.gen.go/function]: ABRecordCopyUniqueId */

// Returns the value of the given property.
//
// Added in macOS .
// Returns the value of the given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyValue(_:_:)
func ABRecordCopyValue(record ABRecordRef, property StringRef) TypeRef {
	return _ABRecordCopyValue(record, property)
}/* debug [functions.gen.go/function]: ABRecordCopyValue */

// Returns a copy of the given record.
//
// Added in macOS 10.4.
// Returns a copy of the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCreateCopy(_:)
func ABRecordCreateCopy(record ABRecordRef) ABRecordRef {
	return _ABRecordCreateCopy(record)
}/* debug [functions.gen.go/function]: ABRecordCreateCopy */

// Returns the unique ID of a record.

// Returns the unique ID of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordGetRecordID(_:)
func ABRecordGetRecordID(record ABRecordRef) ABRecordID {
	return _ABRecordGetRecordID(record)
}/* debug [functions.gen.go/function]: ABRecordGetRecordID */

// Returns the type of a record.

// Returns the type of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordGetRecordType(_:)
func ABRecordGetRecordType(record ABRecordRef) ABRecordType {
	return _ABRecordGetRecordType(record)
}/* debug [functions.gen.go/function]: ABRecordGetRecordType */

// Returns whether or not the record is read-only.
//
// Added in macOS 10.4.
// Returns whether or not the record is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordIsReadOnly(_:)
func ABRecordIsReadOnly(record ABRecordRef) bool {
	return _ABRecordIsReadOnly(record)
}/* debug [functions.gen.go/function]: ABRecordIsReadOnly */

// Removes the value of the given property.
//
// Added in macOS .
// Removes the value of the given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordRemoveValue(_:_:_:)
func ABRecordRemoveValue(record ABRecordRef, property StringRef) bool {
	return _ABRecordRemoveValue(record, property)
}/* debug [functions.gen.go/function]: ABRecordRemoveValue */

// Sets the value of a given property for a record.
//
// Added in macOS .
// Sets the value of a given property for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordSetValue(_:_:_:_:)
func ABRecordSetValue(record ABRecordRef, property ABPropertyID, value TypeRef, error_ unsafe.Pointer) bool {
	return _ABRecordSetValue(record, property, value, error_)
}/* debug [functions.gen.go/function]: ABRecordSetValue */

// Removes the given properties from all the records of this type in the Address Book database, and returns the number of properties successfully removed.
//
// Added in macOS .
// Removes the given properties from all the records of this type in the Address Book database, and returns the number of properties successfully removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRemoveProperties(_:_:_:)
func ABRemoveProperties(addressBook ABAddressBookRef, recordType StringRef, properties ArrayRef) Index {
	return _ABRemoveProperties(addressBook, recordType, properties)
}/* debug [functions.gen.go/function]: ABRemoveProperties */

// Removes the specified record from the Address Book database.
//
// Added in macOS .
// Removes the specified record from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRemoveRecord(_:_:)
func ABRemoveRecord(addressBook ABAddressBookRef, record ABRecordRef) bool {
	return _ABRemoveRecord(addressBook, record)
}/* debug [functions.gen.go/function]: ABRemoveRecord */

// Saves all the changes made since the last save.
//
// Added in macOS .
// Saves all the changes made since the last save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSave(_:)
func ABSave(addressBook ABAddressBookRef) bool {
	return _ABSave(addressBook)
}/* debug [functions.gen.go/function]: ABSave */

// Returns a compound search element created by combiningthe search elements in an array with the given conjunction.
//
// Added in macOS .
// Returns a compound search element created by combiningthe search elements in an array with the given conjunction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElementCreateWithConjunction(_:_:)
func ABSearchElementCreateWithConjunction(conjunction ABSearchConjunction, childrenSearchElement ArrayRef) ABSearchElementRef {
	return _ABSearchElementCreateWithConjunction(conjunction, childrenSearchElement)
}/* debug [functions.gen.go/function]: ABSearchElementCreateWithConjunction */

// Tests whether or not a record matches a search element.
//
// Added in macOS .
// Tests whether or not a record matches a search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElementMatchesRecord(_:_:)
func ABSearchElementMatchesRecord(searchElement ABSearchElementRef, record ABRecordRef) bool {
	return _ABSearchElementMatchesRecord(searchElement, record)
}/* debug [functions.gen.go/function]: ABSearchElementMatchesRecord */

// Sets the record that represents the logged-in user.
//
// Added in macOS .
// Sets the record that represents the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSetMe(_:_:)
func ABSetMe(addressBook ABAddressBookRef, moi ABPersonRef) {
	_ABSetMe(addressBook, moi)
}/* debug [functions.gen.go/function]: ABSetMe */

// Returns the type of a given property for a given record.
//
// Added in macOS .
// Returns the type of a given property for a given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABTypeOfProperty(_:_:_:)
func ABTypeOfProperty(addressBook ABAddressBookRef, recordType StringRef, property StringRef) ABPropertyType {
	return _ABTypeOfProperty(addressBook, recordType, property)
}/* debug [functions.gen.go/function]: ABTypeOfProperty */




