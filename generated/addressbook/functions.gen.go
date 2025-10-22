// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AddressBook Functions (153 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_ABAddPropertiesAndTypes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddRecord func(unsafe.Pointer, unsafe.Pointer) bool
	_ABAddressBookAddRecord func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABAddressBookCopyArrayOfAllGroups func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyArrayOfAllGroupsInSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyArrayOfAllPeople func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyArrayOfAllPeopleInSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyArrayOfAllSources func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyDefaultSource func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyLocalizedLabel func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCopyPeopleWithName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookCreate func() unsafe.Pointer
	_ABAddressBookCreateWithOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookGetAuthorizationStatus func() unsafe.Pointer
	_ABAddressBookGetGroupCount func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookGetGroupWithRecordID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookGetPersonCount func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookGetPersonWithRecordID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookGetSourceWithRecordID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookHasUnsavedChanges func(unsafe.Pointer) bool
	_ABAddressBookRegisterExternalChangeCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookRemoveRecord func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABAddressBookRequestAccessWithCompletion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABAddressBookRevert func(unsafe.Pointer) unsafe.Pointer
	_ABAddressBookSave func(unsafe.Pointer, unsafe.Pointer) bool
	_ABAddressBookUnregisterExternalChangeCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABBeginLoadingImageDataForClient func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABCancelLoadingImageDataForTag func(unsafe.Pointer) unsafe.Pointer
	_ABCopyArrayOfAllGroups func(unsafe.Pointer) unsafe.Pointer
	_ABCopyArrayOfAllPeople func(unsafe.Pointer) unsafe.Pointer
	_ABCopyArrayOfMatchingRecords func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABCopyArrayOfPropertiesForRecordType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABCopyDefaultCountryCode func(unsafe.Pointer) unsafe.Pointer
	_ABCopyLocalizedPropertyOrLabel func(unsafe.Pointer) unsafe.Pointer
	_ABCopyRecordForUniqueId func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABCopyRecordTypeFromUniqueId func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABCreateFormattedAddressFromDictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABGetMe func(unsafe.Pointer) unsafe.Pointer
	_ABGetSharedAddressBook func() unsafe.Pointer
	_ABGroupAddGroup func(unsafe.Pointer, unsafe.Pointer) bool
	_ABGroupAddMember func(unsafe.Pointer, unsafe.Pointer) bool
	_ABGroupCopyArrayOfAllMembers func(unsafe.Pointer) unsafe.Pointer
	_ABGroupCopyArrayOfAllMembersWithSortOrdering func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABGroupCopyArrayOfAllSubgroups func(unsafe.Pointer) unsafe.Pointer
	_ABGroupCopyDistributionIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABGroupCopyParentGroups func(unsafe.Pointer) unsafe.Pointer
	_ABGroupCopySource func(unsafe.Pointer) unsafe.Pointer
	_ABGroupCreate func() unsafe.Pointer
	_ABGroupCreateInSource func(unsafe.Pointer) unsafe.Pointer
	_ABGroupCreateSearchElement func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABGroupRemoveGroup func(unsafe.Pointer, unsafe.Pointer) bool
	_ABGroupRemoveMember func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABGroupSetDistributionIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABHasUnsavedChanges func(unsafe.Pointer) bool
	_ABLocalizedPropertyOrLabel func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueAddValueAndLabel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueCopyArrayOfAllValues func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCopyIdentifierAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCopyLabelAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCopyPrimaryIdentifier func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCopyValueAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCount func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCreate func() unsafe.Pointer
	_ABMultiValueCreateCopy func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueCreateMutable func() unsafe.Pointer
	_ABMultiValueCreateMutableCopy func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueGetCount func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueGetFirstIndexOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueGetIdentifierAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueGetIndexForIdentifier func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueGetPropertyType func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueIndexForIdentifier func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABMultiValueInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueInsertValueAndLabelAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValuePropertyType func(unsafe.Pointer) unsafe.Pointer
	_ABMultiValueRemove func(unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueRemoveValueAndLabelAtIndex func(unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueReplaceLabel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueReplaceLabelAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueReplaceValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueReplaceValueAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABMultiValueSetPrimaryIdentifier func(unsafe.Pointer, unsafe.Pointer) bool
	_ABPersonComparePeopleByName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyArrayOfAllLinkedPeople func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyCompositeNameDelimiterForRecord func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyImageData func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyImageDataWithFormat func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyLocalizedPropertyName func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyParentGroups func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopySource func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCopyVCardRepresentation func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCreate func() unsafe.Pointer
	_ABPersonCreateInSource func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCreatePeopleInSourceWithVCardRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPersonCreateSearchElement func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPersonCreateVCardRepresentationWithPeople func(unsafe.Pointer) unsafe.Pointer
	_ABPersonCreateWithVCardRepresentation func(unsafe.Pointer) unsafe.Pointer
	_ABPersonGetCompositeNameFormat func() unsafe.Pointer
	_ABPersonGetCompositeNameFormatForRecord func(unsafe.Pointer) unsafe.Pointer
	_ABPersonGetSortOrdering func() unsafe.Pointer
	_ABPersonGetTypeOfProperty func(unsafe.Pointer) unsafe.Pointer
	_ABPersonHasImageData func(unsafe.Pointer) bool
	_ABPersonRemoveImageData func(unsafe.Pointer, unsafe.Pointer) bool
	_ABPersonSetImageData func(unsafe.Pointer, unsafe.Pointer) bool
	_ABPickerAddProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerChangeAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerClearSearchField func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCopyColumnTitle func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerCopyDisplayedProperty func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCopyProperties func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCopySelectedGroups func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCopySelectedIdentifiers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerCopySelectedRecords func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCopySelectedValues func(unsafe.Pointer) unsafe.Pointer
	_ABPickerCreate func() unsafe.Pointer
	_ABPickerDeselectAll func(unsafe.Pointer) unsafe.Pointer
	_ABPickerDeselectGroup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerDeselectIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerDeselectRecord func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerEditInAddressBook func(unsafe.Pointer) unsafe.Pointer
	_ABPickerGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_ABPickerGetDelegate func(unsafe.Pointer) unsafe.Pointer
	_ABPickerGetFrame func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerIsVisible func(unsafe.Pointer) bool
	_ABPickerRemoveProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerSelectGroup func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_ABPickerSelectIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_ABPickerSelectInAddressBook func(unsafe.Pointer) unsafe.Pointer
	_ABPickerSelectRecord func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_ABPickerSetColumnTitle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerSetDelegate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerSetDisplayedProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerSetFrame func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABPickerSetVisibility func(unsafe.Pointer, bool) unsafe.Pointer
	_ABRecordCopyCompositeName func(unsafe.Pointer) unsafe.Pointer
	_ABRecordCopyRecordType func(unsafe.Pointer) unsafe.Pointer
	_ABRecordCopyUniqueId func(unsafe.Pointer) unsafe.Pointer
	_ABRecordCopyValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABRecordCreateCopy func(unsafe.Pointer) unsafe.Pointer
	_ABRecordGetRecordID func(unsafe.Pointer) unsafe.Pointer
	_ABRecordGetRecordType func(unsafe.Pointer) unsafe.Pointer
	_ABRecordIsReadOnly func(unsafe.Pointer) bool
	_ABRecordRemoveValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABRecordSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ABRemoveProperties func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABRemoveRecord func(unsafe.Pointer, unsafe.Pointer) bool
	_ABSave func(unsafe.Pointer) bool
	_ABSearchElementCreateWithConjunction func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABSearchElementMatchesRecord func(unsafe.Pointer, unsafe.Pointer) bool
	_ABSetMe func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ABTypeOfProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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

// Adds the given properties to all the records of the specified type in the Address Book database, and returns the number of properties successfully added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddPropertiesAndTypes(_:_:_:)

func ABAddPropertiesAndTypes(addressBook unsafe.Pointer, recordType unsafe.Pointer, propertiesAndTypes unsafe.Pointer) unsafe.Pointer {
	return _ABAddPropertiesAndTypes(addressBook, recordType, propertiesAndTypes)
	}


// Adds a record of the specified type to the Address Book database.

// Adds a record of the specified type to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddRecord(_:_:)

func ABAddRecord(addressBook unsafe.Pointer, record unsafe.Pointer) bool {
	return _ABAddRecord(addressBook, record)
	}


// Adds a record to an address book.

// Adds a record to an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookAddRecord(_:_:_:)

func ABAddressBookAddRecord(addressBook unsafe.Pointer, record unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABAddressBookAddRecord(addressBook, record, error_)
	}


// Returns an array with all the groups in an address book.

// Returns an array with all the groups in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllGroups(_:)

func ABAddressBookCopyArrayOfAllGroups(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllGroups(addressBook)
	}


// Returns an array of all groups from a particular source.

// Returns an array of all groups from a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllGroupsInSource(_:_:)

func ABAddressBookCopyArrayOfAllGroupsInSource(addressBook unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllGroupsInSource(addressBook, source)
	}


// Returns all the person records in an address book.

// Returns all the person records in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeople(_:)

func ABAddressBookCopyArrayOfAllPeople(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllPeople(addressBook)
	}


// Returns an array of all person records from a particular source.

// Returns an array of all person records from a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeopleInSource(_:_:)

func ABAddressBookCopyArrayOfAllPeopleInSource(addressBook unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllPeopleInSource(addressBook, source)
	}


// Returns an array of all person records in the address book, sorted with the specified order.

// Returns an array of all person records in the address book, sorted with the specified order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(_:_:_:)

func ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(addressBook unsafe.Pointer, source unsafe.Pointer, sortOrdering unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllPeopleInSourceWithSortOrdering(addressBook, source, sortOrdering)
	}


// Returns an array of all sources in the address book.

// Returns an array of all sources in the address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyArrayOfAllSources(_:)

func ABAddressBookCopyArrayOfAllSources(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyArrayOfAllSources(addressBook)
	}


// Returns the default source.

// Returns the default source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyDefaultSource(_:)

func ABAddressBookCopyDefaultSource(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyDefaultSource(addressBook)
	}


// Returns a localized version of a record-property label.

// Returns a localized version of a record-property label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyLocalizedLabel(_:)

func ABAddressBookCopyLocalizedLabel(label unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyLocalizedLabel(label)
	}


// Performs a prefix search on the composite names of people in an address book and returns an array of persons that match the search criteria.

// Performs a prefix search on the composite names of people in an address book and returns an array of persons that match the search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCopyPeopleWithName(_:_:)

func ABAddressBookCopyPeopleWithName(addressBook unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCopyPeopleWithName(addressBook, name)
	}


// Creates a new address book object with data from the Address Book database.

// Creates a new address book object with data from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCreate()

func ABAddressBookCreate() unsafe.Pointer {
	return _ABAddressBookCreate()
	}


// Creates a new address book object with data from the Address Book database.

// Creates a new address book object with data from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookCreateWithOptions(_:_:)

func ABAddressBookCreateWithOptions(options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookCreateWithOptions(options, error_)
	}


// Returns the authorization status of your app for accessing address book data.

// Returns the authorization status of your app for accessing address book data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetAuthorizationStatus()

func ABAddressBookGetAuthorizationStatus() unsafe.Pointer {
	return _ABAddressBookGetAuthorizationStatus()
	}


// Returns the number of groups in an address book.

// Returns the number of groups in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetGroupCount(_:)

func ABAddressBookGetGroupCount(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookGetGroupCount(addressBook)
	}


// Returns the group with a given record ID.

// Returns the group with a given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetGroupWithRecordID(_:_:)

func ABAddressBookGetGroupWithRecordID(addressBook unsafe.Pointer, recordID unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookGetGroupWithRecordID(addressBook, recordID)
	}


// Returns the number of person records in an address book.

// Returns the number of person records in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetPersonCount(_:)

func ABAddressBookGetPersonCount(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookGetPersonCount(addressBook)
	}


// Returns the person record with a given record ID.

// Returns the person record with a given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetPersonWithRecordID(_:_:)

func ABAddressBookGetPersonWithRecordID(addressBook unsafe.Pointer, recordID unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookGetPersonWithRecordID(addressBook, recordID)
	}


// Returns the source record with the given record ID.

// Returns the source record with the given record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookGetSourceWithRecordID(_:_:)

func ABAddressBookGetSourceWithRecordID(addressBook unsafe.Pointer, sourceID unsafe.Pointer) unsafe.Pointer {
	return _ABAddressBookGetSourceWithRecordID(addressBook, sourceID)
	}


// Indicates whether an address book has changes that have not been saved to the Address Book database.

// Indicates whether an address book has changes that have not been saved to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookHasUnsavedChanges(_:)

func ABAddressBookHasUnsavedChanges(addressBook unsafe.Pointer) bool {
	return _ABAddressBookHasUnsavedChanges(addressBook)
	}


// Registers a callback to receive notifications when the Address Book database is modified.

// Registers a callback to receive notifications when the Address Book database is modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRegisterExternalChangeCallback(_:_:_:)

func ABAddressBookRegisterExternalChangeCallback(addressBook unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_ABAddressBookRegisterExternalChangeCallback(addressBook, callback, context)
	}


// Removes a record from an address book.

// Removes a record from an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRemoveRecord(_:_:_:)

func ABAddressBookRemoveRecord(addressBook unsafe.Pointer, record unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABAddressBookRemoveRecord(addressBook, record, error_)
	}


// Requests access to address book data from the user.

// Requests access to address book data from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRequestAccessWithCompletion(_:_:)

func ABAddressBookRequestAccessWithCompletion(addressBook unsafe.Pointer, completion unsafe.Pointer) {
	_ABAddressBookRequestAccessWithCompletion(addressBook, completion)
	}


// Discards unsaved changes in an address book.

// Discards unsaved changes in an address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRevert(_:)

func ABAddressBookRevert(addressBook unsafe.Pointer) {
	_ABAddressBookRevert(addressBook)
	}


// Saves any unsaved changes to the Address Book database.

// Saves any unsaved changes to the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookSave(_:_:)

func ABAddressBookSave(addressBook unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABAddressBookSave(addressBook, error_)
	}


// Unregisters a callback.

// Unregisters a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookUnregisterExternalChangeCallback(_:_:_:)

func ABAddressBookUnregisterExternalChangeCallback(addressBook unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_ABAddressBookUnregisterExternalChangeCallback(addressBook, callback, context)
	}


// Starts an asynchronous fetch for image data in all locations, and returns a non-zero tag for tracking.

// Starts an asynchronous fetch for image data in all locations, and returns a non-zero tag for tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABBeginLoadingImageDataForClient(_:_:_:)

func ABBeginLoadingImageDataForClient(person unsafe.Pointer, callback unsafe.Pointer, refcon unsafe.Pointer) unsafe.Pointer {
	return _ABBeginLoadingImageDataForClient(person, callback, refcon)
	}


// Cancels an asynchronous fetch of an image for the given tag.

// Cancels an asynchronous fetch of an image for the given tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCancelLoadingImageDataForTag(_:)

func ABCancelLoadingImageDataForTag(tag unsafe.Pointer) {
	_ABCancelLoadingImageDataForTag(tag)
	}


// Returns an array of all the groups in the Address Book database.

// Returns an array of all the groups in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfAllGroups(_:)

func ABCopyArrayOfAllGroups(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABCopyArrayOfAllGroups(addressBook)
	}


// Returns an array of all the people in the Address Book database.

// Returns an array of all the people in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfAllPeople(_:)

func ABCopyArrayOfAllPeople(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABCopyArrayOfAllPeople(addressBook)
	}


// Returns an array of records that match the given search element, or an empty array if no records match the search element.

// Returns an array of records that match the given search element, or an empty array if no records match the search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfMatchingRecords(_:_:)

func ABCopyArrayOfMatchingRecords(addressBook unsafe.Pointer, search unsafe.Pointer) unsafe.Pointer {
	return _ABCopyArrayOfMatchingRecords(addressBook, search)
	}


// Returns an array containing the names of all the properties for the specified record type.

// Returns an array containing the names of all the properties for the specified record type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyArrayOfPropertiesForRecordType(_:_:)

func ABCopyArrayOfPropertiesForRecordType(addressBook unsafe.Pointer, recordType unsafe.Pointer) unsafe.Pointer {
	return _ABCopyArrayOfPropertiesForRecordType(addressBook, recordType)
	}


// Returns the default country code for records with unspecified country codes.
//
// Added in macOS 10.3.

// Returns the default country code for records with unspecified country codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyDefaultCountryCode(_:)

func ABCopyDefaultCountryCode(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABCopyDefaultCountryCode(addressBook)
	}


// Returns the localized version of a built in property,label, or key.

// Returns the localized version of a built in property,label, or key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyLocalizedPropertyOrLabel(_:)

func ABCopyLocalizedPropertyOrLabel(labelOrProperty unsafe.Pointer) unsafe.Pointer {
	return _ABCopyLocalizedPropertyOrLabel(labelOrProperty)
	}


// Returns the record that matches the given unique ID.

// Returns the record that matches the given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyRecordForUniqueId(_:_:)

func ABCopyRecordForUniqueId(addressBook unsafe.Pointer, uniqueId unsafe.Pointer) unsafe.Pointer {
	return _ABCopyRecordForUniqueId(addressBook, uniqueId)
	}


// Returns the type name of the record that matches a given unique ID.
//
// Added in macOS 10.3.

// Returns the type name of the record that matches a given unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCopyRecordTypeFromUniqueId(_:_:)

func ABCopyRecordTypeFromUniqueId(addressBook unsafe.Pointer, uniqueId unsafe.Pointer) unsafe.Pointer {
	return _ABCopyRecordTypeFromUniqueId(addressBook, uniqueId)
	}


// Returns a string containing the formatted address.
//
// Added in macOS 10.3.

// Returns a string containing the formatted address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABCreateFormattedAddressFromDictionary(_:_:)

func ABCreateFormattedAddressFromDictionary(addressBook unsafe.Pointer, address unsafe.Pointer) unsafe.Pointer {
	return _ABCreateFormattedAddressFromDictionary(addressBook, address)
	}


// Returns the ABPerson object for the logged-in user.

// Returns the ABPerson object for the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGetMe(_:)

func ABGetMe(addressBook unsafe.Pointer) unsafe.Pointer {
	return _ABGetMe(addressBook)
	}


// Returns the unique shared ABAddressBook object.

// Returns the unique shared ABAddressBook object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGetSharedAddressBook()

func ABGetSharedAddressBook() unsafe.Pointer {
	return _ABGetSharedAddressBook()
	}


// Adds a subgroup to another group.

// Adds a subgroup to another group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupAddGroup(_:_:)

func ABGroupAddGroup(group unsafe.Pointer, groupToAdd unsafe.Pointer) bool {
	return _ABGroupAddGroup(group, groupToAdd)
	}


// Adds a person to a group.

// Adds a person to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupAddMember(_:_:)

func ABGroupAddMember(group unsafe.Pointer, personToAdd unsafe.Pointer) bool {
	return _ABGroupAddMember(group, personToAdd)
	}


// Returns an array of persons in a group.

// Returns an array of persons in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllMembers(_:)

func ABGroupCopyArrayOfAllMembers(group unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopyArrayOfAllMembers(group)
	}


// Returns the records in a group, using a sort ordering.

// Returns the records in a group, using a sort ordering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllMembersWithSortOrdering(_:_:)

func ABGroupCopyArrayOfAllMembersWithSortOrdering(group unsafe.Pointer, sortOrdering unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopyArrayOfAllMembersWithSortOrdering(group, sortOrdering)
	}


// Returns an array containing a group’s subgroups.

// Returns an array containing a group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyArrayOfAllSubgroups(_:)

func ABGroupCopyArrayOfAllSubgroups(group unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopyArrayOfAllSubgroups(group)
	}


// Returns the distribution identifier for the given propertyand person.

// Returns the distribution identifier for the given propertyand person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyDistributionIdentifier(_:_:_:)

func ABGroupCopyDistributionIdentifier(group unsafe.Pointer, person unsafe.Pointer, property unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopyDistributionIdentifier(group, person, property)
	}


// Returns an array containing a group’s parents—thegroups that a group belongs to.

// Returns an array containing a group’s parents—thegroups that a group belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopyParentGroups(_:)

func ABGroupCopyParentGroups(group unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopyParentGroups(group)
	}


// Returns the source that the group is from.

// Returns the source that the group is from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCopySource(_:)

func ABGroupCopySource(group unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCopySource(group)
	}


// Returns a new ABGroup object.

// Returns a new ABGroup object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreate()

func ABGroupCreate() unsafe.Pointer {
	return _ABGroupCreate()
	}


// Creates a group in a particular source.

// Creates a group in a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreateInSource(_:)

func ABGroupCreateInSource(source unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCreateInSource(source)
	}


// Creates an ABSearchElement object that specifies a queryfor ABGroup records.

// Creates an ABSearchElement object that specifies a queryfor ABGroup records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupCreateSearchElement(_:_:_:_:_:)

func ABGroupCreateSearchElement(property unsafe.Pointer, label unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, comparison unsafe.Pointer) unsafe.Pointer {
	return _ABGroupCreateSearchElement(property, label, key, value, comparison)
	}


// Removes a subgroup from a group.

// Removes a subgroup from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupRemoveGroup(_:_:)

func ABGroupRemoveGroup(group unsafe.Pointer, groupToRemove unsafe.Pointer) bool {
	return _ABGroupRemoveGroup(group, groupToRemove)
	}


// Removes a person from a group.

// Removes a person from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupRemoveMember(_:_:)

func ABGroupRemoveMember(group unsafe.Pointer, member unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABGroupRemoveMember(group, member, error_)
	}


// Assigning a specific distribution identifier for a person’smulti-value list property so that the group can be used as a distributionlist (mailing list, in the case of an email property).

// Assigning a specific distribution identifier for a person’smulti-value list property so that the group can be used as a distributionlist (mailing list, in the case of an email property).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupSetDistributionIdentifier(_:_:_:_:)

func ABGroupSetDistributionIdentifier(group unsafe.Pointer, person unsafe.Pointer, property unsafe.Pointer, identifier unsafe.Pointer) bool {
	return _ABGroupSetDistributionIdentifier(group, person, property, identifier)
	}


// Returns whether if there are unsaved changes in the address book.

// Returns whether if there are unsaved changes in the address book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABHasUnsavedChanges(_:)

func ABHasUnsavedChanges(addressBook unsafe.Pointer) bool {
	return _ABHasUnsavedChanges(addressBook)
	}


// Returns the localized version of a built in property, label, or key.

// Returns the localized version of a built in property, label, or key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABLocalizedPropertyOrLabel(_:)

func ABLocalizedPropertyOrLabel(propertyOrLabel unsafe.Pointer) unsafe.Pointer {
	return _ABLocalizedPropertyOrLabel(propertyOrLabel)
	}


// Adds a value and its label to a multi-value list.

// Adds a value and its label to a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueAdd(_:_:_:_:)

func ABMultiValueAdd(multiValue unsafe.Pointer, value unsafe.Pointer, label unsafe.Pointer, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueAdd(multiValue, value, label, outIdentifier)
	}


// Adds a value and its corresponding label to a multivalue property.

// Adds a value and its corresponding label to a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueAddValueAndLabel(_:_:_:_:)

func ABMultiValueAddValueAndLabel(multiValue unsafe.Pointer, value unsafe.Pointer, label unsafe.Pointer, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueAddValueAndLabel(multiValue, value, label, outIdentifier)
	}


// Returns an array with the values in a multivalue property.

// Returns an array with the values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyArrayOfAllValues(_:)

func ABMultiValueCopyArrayOfAllValues(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCopyArrayOfAllValues(multiValue)
	}


// Returns the identifier at the given index.

// Returns the identifier at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyIdentifierAtIndex(_:_:)

func ABMultiValueCopyIdentifierAtIndex(multiValue unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCopyIdentifierAtIndex(multiValue, index)
	}


// Returns the label for the given index.

// Returns the label for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyLabelAtIndex(_:_:)

func ABMultiValueCopyLabelAtIndex(multiValue unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCopyLabelAtIndex(multiValue, index)
	}


// Returns the identifier for the primary value.

// Returns the identifier for the primary value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyPrimaryIdentifier(_:)

func ABMultiValueCopyPrimaryIdentifier(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCopyPrimaryIdentifier(multiValue)
	}


// Returns the value for the given index.

// Returns the value for the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCopyValueAtIndex(_:_:)

func ABMultiValueCopyValueAtIndex(multiValue unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCopyValueAtIndex(multiValue, index)
	}


// Returns the number of entries in a multi-value list.

// Returns the number of entries in a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCount(_:)

func ABMultiValueCount(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCount(multiValue)
	}


// Returns a new ABMultiValue object.

// Returns a new ABMultiValue object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreate()

func ABMultiValueCreate() unsafe.Pointer {
	return _ABMultiValueCreate()
	}


// Returns a copy of a multi-value object.

// Returns a copy of a multi-value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateCopy(_:)

func ABMultiValueCreateCopy(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCreateCopy(multiValue)
	}


// Returns a newly created mutable multi-value list object.

// Returns a newly created mutable multi-value list object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateMutable()

func ABMultiValueCreateMutable() unsafe.Pointer {
	return _ABMultiValueCreateMutable()
	}


// Returns a mutable copy of a multi-value object.

// Returns a mutable copy of a multi-value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueCreateMutableCopy(_:)

func ABMultiValueCreateMutableCopy(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueCreateMutableCopy(multiValue)
	}


// Returns the number of values in a multivalue property.

// Returns the number of values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetCount(_:)

func ABMultiValueGetCount(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueGetCount(multiValue)
	}


// Returns the first location of a value in a multivalue property.

// Returns the first location of a value in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetFirstIndexOfValue(_:_:)

func ABMultiValueGetFirstIndexOfValue(multiValue unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueGetFirstIndexOfValue(multiValue, value)
	}


// Returns the identifier of a value in a multivalue property.

// Returns the identifier of a value in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetIdentifierAtIndex(_:_:)

func ABMultiValueGetIdentifierAtIndex(multiValue unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueGetIdentifierAtIndex(multiValue, index)
	}


// Returns the location (within a multivalue property) of a value with a given identifier.

// Returns the location (within a multivalue property) of a value with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetIndexForIdentifier(_:_:)

func ABMultiValueGetIndexForIdentifier(multiValue unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueGetIndexForIdentifier(multiValue, identifier)
	}


// Returns the type of the values contained in a multivalue property.

// Returns the type of the values contained in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueGetPropertyType(_:)

func ABMultiValueGetPropertyType(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueGetPropertyType(multiValue)
	}


// Returns the index for the given identifier.

// Returns the index for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueIndexForIdentifier(_:_:)

func ABMultiValueIndexForIdentifier(multiValue unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValueIndexForIdentifier(multiValue, identifier)
	}


// Inserts a value and its label at the given index in amulti-value list.

// Inserts a value and its label at the given index in amulti-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueInsert(_:_:_:_:_:)

func ABMultiValueInsert(multiValue unsafe.Pointer, value unsafe.Pointer, label unsafe.Pointer, index unsafe.Pointer, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueInsert(multiValue, value, label, index, outIdentifier)
	}


// Inserts a value and a label into a multivalue property.

// Inserts a value and a label into a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueInsertValueAndLabelAtIndex(_:_:_:_:_:)

func ABMultiValueInsertValueAndLabelAtIndex(multiValue unsafe.Pointer, value unsafe.Pointer, label unsafe.Pointer, index unsafe.Pointer, outIdentifier unsafe.Pointer) bool {
	return _ABMultiValueInsertValueAndLabelAtIndex(multiValue, value, label, index, outIdentifier)
	}


// Returns the type for the values in a multi-value list.

// Returns the type for the values in a multi-value list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValuePropertyType(_:)

func ABMultiValuePropertyType(multiValue unsafe.Pointer) unsafe.Pointer {
	return _ABMultiValuePropertyType(multiValue)
	}


// Removes the value and label at the given index.

// Removes the value and label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueRemove(_:_:)

func ABMultiValueRemove(multiValue unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueRemove(multiValue, index)
	}


// Removes a value from a multivalue property.

// Removes a value from a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueRemoveValueAndLabelAtIndex(_:_:)

func ABMultiValueRemoveValueAndLabelAtIndex(multiValue unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueRemoveValueAndLabelAtIndex(multiValue, index)
	}


// Replaces the label at the given index.

// Replaces the label at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceLabel(_:_:_:)

func ABMultiValueReplaceLabel(multiValue unsafe.Pointer, label unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueReplaceLabel(multiValue, label, index)
	}


// Replaces a label in a multivalue property with another label.

// Replaces a label in a multivalue property with another label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceLabelAtIndex(_:_:_:)

func ABMultiValueReplaceLabelAtIndex(multiValue unsafe.Pointer, label unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueReplaceLabelAtIndex(multiValue, label, index)
	}


// Replaces the value at the given index.

// Replaces the value at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceValue(_:_:_:)

func ABMultiValueReplaceValue(multiValue unsafe.Pointer, value unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueReplaceValue(multiValue, value, index)
	}


// Replaces a value in a multivalue property with another value.

// Replaces a value in a multivalue property with another value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueReplaceValueAtIndex(_:_:_:)

func ABMultiValueReplaceValueAtIndex(multiValue unsafe.Pointer, value unsafe.Pointer, index unsafe.Pointer) bool {
	return _ABMultiValueReplaceValueAtIndex(multiValue, value, index)
	}


// Sets the primary value to be the value for the given identifier.

// Sets the primary value to be the value for the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueSetPrimaryIdentifier(_:_:)

func ABMultiValueSetPrimaryIdentifier(multiValue unsafe.Pointer, identifier unsafe.Pointer) bool {
	return _ABMultiValueSetPrimaryIdentifier(multiValue, identifier)
	}


// Indicates how two person records get sorted.

// Indicates how two person records get sorted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonComparePeopleByName(_:_:_:)

func ABPersonComparePeopleByName(person1 unsafe.Pointer, person2 unsafe.Pointer, ordering unsafe.Pointer) unsafe.Pointer {
	return _ABPersonComparePeopleByName(person1, person2, ordering)
	}


// Returns an array of all person records in the address book database that are linked to the given person record.

// Returns an array of all person records in the address book database that are linked to the given person record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyArrayOfAllLinkedPeople(_:)

func ABPersonCopyArrayOfAllLinkedPeople(person unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyArrayOfAllLinkedPeople(person)
	}


// Returns the delimiter to use between name components.

// Returns the delimiter to use between name components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyCompositeNameDelimiterForRecord(_:)

func ABPersonCopyCompositeNameDelimiterForRecord(record unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyCompositeNameDelimiterForRecord(record)
	}


// Returns data that contains a picture of a person.

// Returns data that contains a picture of a person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyImageData(_:)

func ABPersonCopyImageData(person unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyImageData(person)
	}


// Returns the picture for a person record in the given format.

// Returns the picture for a person record in the given format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyImageDataWithFormat(_:_:)

func ABPersonCopyImageDataWithFormat(person unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyImageDataWithFormat(person, format)
	}


// Returns the localized name of a person property

// Returns the localized name of a person property
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyLocalizedPropertyName(_:)

func ABPersonCopyLocalizedPropertyName(property unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyLocalizedPropertyName(property)
	}


// Returns an array of groups that a person belongs to.

// Returns an array of groups that a person belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyParentGroups(_:)

func ABPersonCopyParentGroups(person unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyParentGroups(person)
	}


// Returns the source that the person record is from.

// Returns the source that the person record is from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopySource(_:)

func ABPersonCopySource(person unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopySource(person)
	}


// Returns the vCard representation of the person as a data object in vCard format.

// Returns the vCard representation of the person as a data object in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCopyVCardRepresentation(_:)

func ABPersonCopyVCardRepresentation(person unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCopyVCardRepresentation(person)
	}


// Returns a newly created person object.

// Returns a newly created person object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreate()

func ABPersonCreate() unsafe.Pointer {
	return _ABPersonCreate()
	}


// Creates a new person record in a particular source.

// Creates a new person record in a particular source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateInSource(_:)

func ABPersonCreateInSource(source unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCreateInSource(source)
	}


// Creates person records from the given vCard representation.

// Creates person records from the given vCard representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreatePeopleInSourceWithVCardRepresentation(_:_:)

func ABPersonCreatePeopleInSourceWithVCardRepresentation(source unsafe.Pointer, vCardData unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCreatePeopleInSourceWithVCardRepresentation(source, vCardData)
	}


// Returns a search element object that specifies a query for records of this type.

// Returns a search element object that specifies a query for records of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateSearchElement(_:_:_:_:_:)

func ABPersonCreateSearchElement(property unsafe.Pointer, label unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, comparison unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCreateSearchElement(property, label, key, value, comparison)
	}


// Returns the vCard representation of the given person records.

// Returns the vCard representation of the given person records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateVCardRepresentationWithPeople(_:)

func ABPersonCreateVCardRepresentationWithPeople(people unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCreateVCardRepresentationWithPeople(people)
	}


// Returns a new ABPerson object initialized with the given data in vCard format.

// Returns a new ABPerson object initialized with the given data in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCreateWithVCardRepresentation(_:)

func ABPersonCreateWithVCardRepresentation(vCard unsafe.Pointer) unsafe.Pointer {
	return _ABPersonCreateWithVCardRepresentation(vCard)
	}


// Returns the person-name display format.

// Returns the person-name display format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetCompositeNameFormat()

func ABPersonGetCompositeNameFormat() unsafe.Pointer {
	return _ABPersonGetCompositeNameFormat()
	}


// Returns the person-name display format to use for the given record.

// Returns the person-name display format to use for the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetCompositeNameFormatForRecord(_:)

func ABPersonGetCompositeNameFormatForRecord(record unsafe.Pointer) unsafe.Pointer {
	return _ABPersonGetCompositeNameFormatForRecord(record)
	}


// Returns the user’s sort-ordering preference for lists of persons.

// Returns the user’s sort-ordering preference for lists of persons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetSortOrdering()

func ABPersonGetSortOrdering() unsafe.Pointer {
	return _ABPersonGetSortOrdering()
	}


// Returns the type of a person property.

// Returns the type of a person property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonGetTypeOfProperty(_:)

func ABPersonGetTypeOfProperty(property unsafe.Pointer) unsafe.Pointer {
	return _ABPersonGetTypeOfProperty(property)
	}


// Indicates whether a person has a picture.

// Indicates whether a person has a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonHasImageData(_:)

func ABPersonHasImageData(person unsafe.Pointer) bool {
	return _ABPersonHasImageData(person)
	}


// Removes a person’s picture.

// Removes a person’s picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonRemoveImageData(_:_:)

func ABPersonRemoveImageData(person unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABPersonRemoveImageData(person, error_)
	}


// Sets the image for this person to the given data.

// Sets the image for this person to the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonSetImageData(_:_:)

func ABPersonSetImageData(person unsafe.Pointer, imageData unsafe.Pointer) bool {
	return _ABPersonSetImageData(person, imageData)
	}


// Adds a property to the group of properties available in the record list. Use to remove a property from the list and to obtain the list of properties available in the list.
//
// Added in macOS 10.3.

// Adds a property to the group of properties available in the record list. Use to remove a property from the list and to obtain the list of properties available in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerAddProperty

func ABPickerAddProperty(inPicker unsafe.Pointer, inProperty unsafe.Pointer) {
	_ABPickerAddProperty(inPicker, inProperty)
	}


// Specifies the selection behaviors for a people-picker window. Use to obtain the selection behaviors specified for the window.
//
// Added in macOS 10.3.

// Specifies the selection behaviors for a people-picker window. Use to obtain the selection behaviors specified for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerChangeAttributes

func ABPickerChangeAttributes(inPicker unsafe.Pointer, inAttributesToSet unsafe.Pointer, inAttributesToClear unsafe.Pointer) {
	_ABPickerChangeAttributes(inPicker, inAttributesToSet, inAttributesToClear)
	}


// Clears the search field and resets the list of displayed records.
//
// Added in macOS 10.3.

// Clears the search field and resets the list of displayed records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerClearSearchField

func ABPickerClearSearchField(inPicker unsafe.Pointer) {
	_ABPickerClearSearchField(inPicker)
	}


// Obtains the title of a custom property.
//
// Added in macOS 10.3.

// Obtains the title of a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyColumnTitle

func ABPickerCopyColumnTitle(inPicker unsafe.Pointer, inProperty unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopyColumnTitle(inPicker, inProperty)
	}


// Returns the name of the property currently displayed in the record list.
//
// Added in macOS 10.3.

// Returns the name of the property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyDisplayedProperty

func ABPickerCopyDisplayedProperty(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopyDisplayedProperty(inPicker)
	}


// Obtains the list of properties available in the record list. Use to add a property to the record list and to remove a property from the list.
//
// Added in macOS 10.3.

// Obtains the list of properties available in the record list. Use to add a property to the record list and to remove a property from the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopyProperties

func ABPickerCopyProperties(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopyProperties(inPicker)
	}


// Returns the groups selected in the group list as an array of objects.
//
// Added in macOS 10.3.

// Returns the groups selected in the group list as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedGroups

func ABPickerCopySelectedGroups(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopySelectedGroups(inPicker)
	}


// Returns the identifiers of the selected values in a multi-value property or an empty array if the property displayed is a single-value property.
//
// Added in macOS 10.3.

// Returns the identifiers of the selected values in a multi-value property or an empty array if the property displayed is a single-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedIdentifiers

func ABPickerCopySelectedIdentifiers(inPicker unsafe.Pointer, inPerson unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopySelectedIdentifiers(inPicker, inPerson)
	}


// Returns the selection in the record list as an array of ABGroup or objects.
//
// Added in macOS 10.3.

// Returns the selection in the record list as an array of ABGroup or objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedRecords

func ABPickerCopySelectedRecords(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopySelectedRecords(inPicker)
	}


// Returns the selected values in a multi-value property or an empty array if no values are selected or the property displayedis a single-value property.
//
// Added in macOS 10.3.

// Returns the selected values in a multi-value property or an empty array if no values are selected or the property displayedis a single-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCopySelectedValues

func ABPickerCopySelectedValues(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerCopySelectedValues(inPicker)
	}


// Creates an ABPickerRef. The corresponding window is hidden. Invoke to show it. Release with .
//
// Added in macOS 10.3.

// Creates an ABPickerRef. The corresponding window is hidden. Invoke to show it. Release with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerCreate

func ABPickerCreate() unsafe.Pointer {
	return _ABPickerCreate()
	}


// Deselects all selected groups, records, and values in multi-value properties.
//
// Added in macOS 10.3.

// Deselects all selected groups, records, and values in multi-value properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectAll

func ABPickerDeselectAll(inPicker unsafe.Pointer) {
	_ABPickerDeselectAll(inPicker)
	}


// Deselects a group in the group list.
//
// Added in macOS 10.3.

// Deselects a group in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectGroup

func ABPickerDeselectGroup(inPicker unsafe.Pointer, inGroup unsafe.Pointer) {
	_ABPickerDeselectGroup(inPicker, inGroup)
	}


// Deselects a value in multi-value property currently displayed in the record list.
//
// Added in macOS 10.3.

// Deselects a value in multi-value property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectIdentifier

func ABPickerDeselectIdentifier(inPicker unsafe.Pointer, inPerson unsafe.Pointer, inIdentifier unsafe.Pointer) {
	_ABPickerDeselectIdentifier(inPicker, inPerson, inIdentifier)
	}


// Deselects a group in the record list.
//
// Added in macOS 10.3.

// Deselects a group in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerDeselectRecord

func ABPickerDeselectRecord(inPicker unsafe.Pointer, inRecord unsafe.Pointer) {
	_ABPickerDeselectRecord(inPicker, inRecord)
	}


// Launches Address Book to edit the item selected in the people-picker window.
//
// Added in macOS 10.3.

// Launches Address Book to edit the item selected in the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerEditInAddressBook

func ABPickerEditInAddressBook(inPicker unsafe.Pointer) {
	_ABPickerEditInAddressBook(inPicker)
	}


// Indicates the selection behaviors selected a people-picker window. Use tospecify selection behaviors for the window.
//
// Added in macOS 10.3.

// Indicates the selection behaviors selected a people-picker window. Use tospecify selection behaviors for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetAttributes

func ABPickerGetAttributes(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerGetAttributes(inPicker)
	}


// Obtains the delegate for a people-picker window.
//
// Added in macOS 10.3.

// Obtains the delegate for a people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetDelegate

func ABPickerGetDelegate(inPicker unsafe.Pointer) unsafe.Pointer {
	return _ABPickerGetDelegate(inPicker)
	}


// Returns the position and size of the people-picker window.
//
// Added in macOS 10.3.

// Returns the position and size of the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerGetFrame

func ABPickerGetFrame(inPicker unsafe.Pointer, outFrame unsafe.Pointer) {
	_ABPickerGetFrame(inPicker, outFrame)
	}


// Indicates whether the people-picker window is visible.
//
// Added in macOS 10.3.

// Indicates whether the people-picker window is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerIsVisible

func ABPickerIsVisible(inPicker unsafe.Pointer) bool {
	return _ABPickerIsVisible(inPicker)
	}


// Removes a property from the group of properties whose values are shown in the record list. Use to add a property to the record list and to obtain the list of properties shown in the record list.
//
// Added in macOS 10.3.

// Removes a property from the group of properties whose values are shown in the record list. Use to add a property to the record list and to obtain the list of properties shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerRemoveProperty

func ABPickerRemoveProperty(inPicker unsafe.Pointer, inProperty unsafe.Pointer) {
	_ABPickerRemoveProperty(inPicker, inProperty)
	}


// Selects a group or a set of groups in the group list.
//
// Added in macOS 10.3.

// Selects a group or a set of groups in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectGroup

func ABPickerSelectGroup(inPicker unsafe.Pointer, inGroup unsafe.Pointer, inExtendSelection bool) {
	_ABPickerSelectGroup(inPicker, inGroup, inExtendSelection)
	}


// Selects a value or a set of values in a multi-value property.
//
// Added in macOS 10.3.

// Selects a value or a set of values in a multi-value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectIdentifier

func ABPickerSelectIdentifier(inPicker unsafe.Pointer, inPerson unsafe.Pointer, inIdentifier unsafe.Pointer, inExtendSelection bool) {
	_ABPickerSelectIdentifier(inPicker, inPerson, inIdentifier, inExtendSelection)
	}


// Launches Address Book and selects the item selected in the people-picker window.
//
// Added in macOS 10.3.

// Launches Address Book and selects the item selected in the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectInAddressBook

func ABPickerSelectInAddressBook(inPicker unsafe.Pointer) {
	_ABPickerSelectInAddressBook(inPicker)
	}


// Selects a record or a set of records in the record list.
//
// Added in macOS 10.3.

// Selects a record or a set of records in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSelectRecord

func ABPickerSelectRecord(inPicker unsafe.Pointer, inRecord unsafe.Pointer, inExtendSelection bool) {
	_ABPickerSelectRecord(inPicker, inRecord, inExtendSelection)
	}


// Sets the title for a custom property.
//
// Added in macOS 10.3.

// Sets the title for a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetColumnTitle

func ABPickerSetColumnTitle(inPicker unsafe.Pointer, inTitle unsafe.Pointer, inProperty unsafe.Pointer) {
	_ABPickerSetColumnTitle(inPicker, inTitle, inProperty)
	}


// Sets the event handler for people-picker events.
//
// Added in macOS 10.3.

// Sets the event handler for people-picker events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetDelegate

func ABPickerSetDelegate(inPicker unsafe.Pointer, inDelegate unsafe.Pointer) {
	_ABPickerSetDelegate(inPicker, inDelegate)
	}


// Displays one of the properties whose values are shownin the record list.
//
// Added in macOS 10.3.

// Displays one of the properties whose values are shownin the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetDisplayedProperty

func ABPickerSetDisplayedProperty(inPicker unsafe.Pointer, inProperty unsafe.Pointer) {
	_ABPickerSetDisplayedProperty(inPicker, inProperty)
	}


// Specifies the position and size of the people-picker window.
//
// Added in macOS 10.3.

// Specifies the position and size of the people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetFrame

func ABPickerSetFrame(inPicker unsafe.Pointer, inFrame unsafe.Pointer) {
	_ABPickerSetFrame(inPicker, inFrame)
	}


// Shows or hides a people-picker window.
//
// Added in macOS 10.3.

// Shows or hides a people-picker window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPickerSetVisibility

func ABPickerSetVisibility(inPicker unsafe.Pointer, visible bool) {
	_ABPickerSetVisibility(inPicker, visible)
	}


// Returns an appropriate, human-friendly name for the record.

// Returns an appropriate, human-friendly name for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyCompositeName(_:)

func ABRecordCopyCompositeName(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordCopyCompositeName(record)
	}


// Returns the type of the given record.

// Returns the type of the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyRecordType(_:)

func ABRecordCopyRecordType(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordCopyRecordType(record)
	}


// Returns the unique ID of the receiver.

// Returns the unique ID of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyUniqueId(_:)

func ABRecordCopyUniqueId(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordCopyUniqueId(record)
	}


// Returns the value of the given property.

// Returns the value of the given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCopyValue(_:_:)

func ABRecordCopyValue(record unsafe.Pointer, property unsafe.Pointer) unsafe.Pointer {
	return _ABRecordCopyValue(record, property)
	}


// Returns a copy of the given record.
//
// Added in macOS 10.4.

// Returns a copy of the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordCreateCopy(_:)

func ABRecordCreateCopy(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordCreateCopy(record)
	}


// Returns the unique ID of a record.

// Returns the unique ID of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordGetRecordID(_:)

func ABRecordGetRecordID(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordGetRecordID(record)
	}


// Returns the type of a record.

// Returns the type of a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordGetRecordType(_:)

func ABRecordGetRecordType(record unsafe.Pointer) unsafe.Pointer {
	return _ABRecordGetRecordType(record)
	}


// Returns whether or not the record is read-only.
//
// Added in macOS 10.4.

// Returns whether or not the record is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordIsReadOnly(_:)

func ABRecordIsReadOnly(record unsafe.Pointer) bool {
	return _ABRecordIsReadOnly(record)
	}


// Removes the value of the given property.

// Removes the value of the given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordRemoveValue(_:_:)

func ABRecordRemoveValue(record unsafe.Pointer, property unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ABRecordRemoveValue(record, property, error_)
	}


// Sets the value of a given property for a record.

// Sets the value of a given property for a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordSetValue(_:_:_:)

func ABRecordSetValue(record unsafe.Pointer, property unsafe.Pointer, value unsafe.Pointer) bool {
	return _ABRecordSetValue(record, property, value)
	}


// Removes the given properties from all the records of this type in the Address Book database, and returns the number of properties successfully removed.

// Removes the given properties from all the records of this type in the Address Book database, and returns the number of properties successfully removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRemoveProperties(_:_:_:)

func ABRemoveProperties(addressBook unsafe.Pointer, recordType unsafe.Pointer, properties unsafe.Pointer) unsafe.Pointer {
	return _ABRemoveProperties(addressBook, recordType, properties)
	}


// Removes the specified record from the Address Book database.

// Removes the specified record from the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRemoveRecord(_:_:)

func ABRemoveRecord(addressBook unsafe.Pointer, record unsafe.Pointer) bool {
	return _ABRemoveRecord(addressBook, record)
	}


// Saves all the changes made since the last save.

// Saves all the changes made since the last save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSave(_:)

func ABSave(addressBook unsafe.Pointer) bool {
	return _ABSave(addressBook)
	}


// Returns a compound search element created by combiningthe search elements in an array with the given conjunction.

// Returns a compound search element created by combiningthe search elements in an array with the given conjunction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElementCreateWithConjunction(_:_:)

func ABSearchElementCreateWithConjunction(conjunction unsafe.Pointer, childrenSearchElement unsafe.Pointer) unsafe.Pointer {
	return _ABSearchElementCreateWithConjunction(conjunction, childrenSearchElement)
	}


// Tests whether or not a record matches a search element.

// Tests whether or not a record matches a search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElementMatchesRecord(_:_:)

func ABSearchElementMatchesRecord(searchElement unsafe.Pointer, record unsafe.Pointer) bool {
	return _ABSearchElementMatchesRecord(searchElement, record)
	}


// Sets the record that represents the logged-in user.

// Sets the record that represents the logged-in user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSetMe(_:_:)

func ABSetMe(addressBook unsafe.Pointer, moi unsafe.Pointer) {
	_ABSetMe(addressBook, moi)
	}


// Returns the type of a given property for a given record.

// Returns the type of a given property for a given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABTypeOfProperty(_:_:_:)

func ABTypeOfProperty(addressBook unsafe.Pointer, recordType unsafe.Pointer, property unsafe.Pointer) unsafe.Pointer {
	return _ABTypeOfProperty(addressBook, recordType, property)
	}




