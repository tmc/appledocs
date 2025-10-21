// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

// Type aliases and typedefs
// ABAddressBookRef - A reference to an ABAddressBook object.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAddressBookRef
// ABAddressBookRef has base type: struct __ABAddressBookRef *
type ABAddressBookRef uintptr
// ABExternalChangeCallback - Prototype for a function callback invoked on an address book when the Address Book database is modified by another address book instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABExternalChangeCallback
// ABExternalChangeCallback has base type: void (*)(const void *, const struct __CFDictionary *, void *)
type ABExternalChangeCallback uintptr
// ABGroupRef - A reference to an ABGroup object.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroupRef
// ABGroupRef has base type: struct __ABGroup *
type ABGroupRef uintptr
// ABImageClientCallback - Prototype of a callback function used to notify an application when an asynchronous image fetch is complete.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABImageClientCallback
// ABImageClientCallback has base type: void (*)(const struct __CFData *, long, void *)
type ABImageClientCallback uintptr
// ABMultiValueIdentifier - Identifies multivalue properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueIdentifier
// ABMultiValueIdentifier has base type: int32_t
type ABMultiValueIdentifier uintptr
// ABMultiValueRef - A reference to an   or  .
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMultiValueRef
// ABMultiValueRef has base type: const struct __ABMultiValue *
type ABMultiValueRef uintptr
// ABMutableMultiValueRef - A reference to an ABMutableMultiValue object.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABMutableMultiValueRef
// ABMutableMultiValueRef has base type: CFTypeRef
type ABMutableMultiValueRef uintptr
// ABPersonCompositeNameFormat - Indicates a person-name display format.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonCompositeNameFormat
// ABPersonCompositeNameFormat has base type: uint32_t
type ABPersonCompositeNameFormat uintptr
// ABPersonRef - A reference to an ABPerson object.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonRef
// ABPersonRef has base type: struct __ABPerson *
type ABPersonRef uintptr
// ABPersonSortOrdering - Indicates a person sort ordering.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonSortOrdering
// ABPersonSortOrdering has base type: uint32_t
type ABPersonSortOrdering uintptr
// ABPropertyID - Integer that identifies a record property.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPropertyID
// ABPropertyID has base type: int32_t
type ABPropertyID uintptr
// ABPropertyType - These are the possible types of ABRecord properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPropertyType
// ABPropertyType has base type: CFIndex
type ABPropertyType uintptr
// ABRecordID - Integer that identifies a record.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordID
// ABRecordID has base type: int32_t
type ABRecordID uintptr
// ABRecordRef - A reference to an ABRecord object or any of its derivedopaque types.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordRef
// ABRecordRef has base type: void *
type ABRecordRef uintptr
// ABRecordType - Integer that identifies a record type.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABRecordType
// ABRecordType has base type: uint32_t
type ABRecordType uintptr
// ABSearchComparison - Constants used to specify the type of comparison beingmade.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchComparison
// ABSearchComparison has base type: CFIndex
type ABSearchComparison uintptr
// ABSearchConjunction - Constants used to create compound search elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchConjunction
// ABSearchConjunction has base type: CFIndex
type ABSearchConjunction uintptr
// ABSearchElementRef - A reference to an ABSearchElement object.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElementRef
// ABSearchElementRef has base type: struct __ABSearchElementRef *
type ABSearchElementRef uintptr
// ABSourceType - Indicates a source type. See  .
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSourceType
type ABSourceType int32

