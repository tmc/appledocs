// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

// Enum types and constants
// CNAuthorizationStatus - An authorization status the user can grant for an app to access the specified entity type.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus
type CNAuthorizationStatus uint

const (
	// CNAuthorizationStatusAuthorized - The application is authorized to access contact data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus/authorized
	CNAuthorizationStatusAuthorized CNAuthorizationStatus = 0
	// CNAuthorizationStatusDenied - The user explicitly denied access to contact data for the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus/denied
	CNAuthorizationStatusDenied CNAuthorizationStatus = 0
	// CNAuthorizationStatusLimited - The app has access to a limited subset of contacts, chosen by the person using   the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus/limited
	CNAuthorizationStatusLimited CNAuthorizationStatus = 0
	// CNAuthorizationStatusNotDetermined - The user has not yet made a choice regarding whether the application may access contact data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus/notDetermined
	CNAuthorizationStatusNotDetermined CNAuthorizationStatus = 0
	// CNAuthorizationStatusRestricted - The application is not authorized to access contact data. The user cannot change this application’s status, possibly due to active restrictions such as parental controls being in place.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNAuthorizationStatus/restricted
	CNAuthorizationStatusRestricted CNAuthorizationStatus = 0
)

// CNContactDisplayNameOrder - The formatting orders for contact names component.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactDisplayNameOrder
type CNContactDisplayNameOrder uint

const (
	// CNContactDisplayNameOrderFamilyNameFirst - Display name order by family name first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactDisplayNameOrder/familyNameFirst
	CNContactDisplayNameOrderFamilyNameFirst CNContactDisplayNameOrder = 0
	// CNContactDisplayNameOrderGivenNameFirst - Display name order by given name first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactDisplayNameOrder/givenNameFirst
	CNContactDisplayNameOrderGivenNameFirst CNContactDisplayNameOrder = 0
	// CNContactDisplayNameOrderUserDefault - Display name order by user default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactDisplayNameOrder/userDefault
	CNContactDisplayNameOrderUserDefault CNContactDisplayNameOrder = 0
)

// CNContactFormatterStyle - The formatting styles for contact names.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatterStyle
type CNContactFormatterStyle uint

const (
	// CNContactFormatterStyleFullName - Combines the contact name components into a full name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatterStyle/fullName
	CNContactFormatterStyleFullName CNContactFormatterStyle = 0
	// CNContactFormatterStylePhoneticFullName - Combines the contact phonetic name components into a phonetic full name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFormatterStyle/phoneticFullName
	CNContactFormatterStylePhoneticFullName CNContactFormatterStyle = 0
)

// CNContactSortOrder - Indicates the sorting order for contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactSortOrder
type CNContactSortOrder uint

const (
	// CNContactSortOrderFamilyName - Sorting contacts by family name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactSortOrder/familyName
	CNContactSortOrderFamilyName CNContactSortOrder = 0
	// CNContactSortOrderGivenName - Sorting contacts by given name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactSortOrder/givenName
	CNContactSortOrderGivenName CNContactSortOrder = 0
	// CNContactSortOrderNone - No sorting order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactSortOrder/none
	CNContactSortOrderNone CNContactSortOrder = 0
	// CNContactSortOrderUserDefault - The user’s default sorting order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactSortOrder/userDefault
	CNContactSortOrderUserDefault CNContactSortOrder = 0
)

// CNContactType - The types a contact can be.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactType
type CNContactType uint

const (
	// CNContactTypeOrganization - The contact is an Organization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactType/organization
	CNContactTypeOrganization CNContactType = 0
	// CNContactTypePerson - The contact is a person.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactType/person
	CNContactTypePerson CNContactType = 0
)

// CNContainerType - The container may be local on the device or associated with a server account that has contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainerType
type CNContainerType uint

const (
	// CNContainerTypeCardDAV - A container for contacts stored in an CardDAV server, such as iCloud.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainerType/cardDAV
	CNContainerTypeCardDAV CNContainerType = 0
	// CNContainerTypeExchange - A container for contacts stored in an Exchange folder from an Exchange server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainerType/exchange
	CNContainerTypeExchange CNContainerType = 0
	// CNContainerTypeLocal - A container for contacts only stored locally on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainerType/local
	CNContainerTypeLocal CNContainerType = 0
	// CNContainerTypeUnassigned - A container where the system hasn’t assigned the container type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainerType/unassigned
	CNContainerTypeUnassigned CNContainerType = 0
)

// CNEntityType - The entities the user can grant access to.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNEntityType
type CNEntityType uint

const (
	// CNEntityTypeContacts - The user’s contacts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNEntityType/contacts
	CNEntityTypeContacts CNEntityType = 0
)

// CNErrorCode - Error codes that the system may return when you use Contacts framework methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code
type CNErrorCode uint

const (
	// CNErrorCodeAuthorizationDenied - An error that indicates the system denied authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/authorizationDenied
	CNErrorCodeAuthorizationDenied CNErrorCode = 0
	// CNErrorCodeChangeHistoryExpired - An error that indicates the change history expired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/changeHistoryExpired
	CNErrorCodeChangeHistoryExpired CNErrorCode = 0
	// CNErrorCodeChangeHistoryInvalidAnchor - An error that indicates a change history anchor is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/changeHistoryInvalidAnchor
	CNErrorCodeChangeHistoryInvalidAnchor CNErrorCode = 0
	// CNErrorCodeChangeHistoryInvalidFetchRequest - An error that indicates a change history fetch request is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/changeHistoryInvalidFetchRequest
	CNErrorCodeChangeHistoryInvalidFetchRequest CNErrorCode = 0
	// CNErrorCodeClientIdentifierCollision - An error that indicates a client identifier collision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/clientIdentifierCollision
	CNErrorCodeClientIdentifierCollision CNErrorCode = 0
	// CNErrorCodeClientIdentifierDoesNotExist - An error that indicates the client identifier doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/clientIdentifierDoesNotExist
	CNErrorCodeClientIdentifierDoesNotExist CNErrorCode = 0
	// CNErrorCodeClientIdentifierInvalid - An error that indicates the client identifier is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/clientIdentifierInvalid
	CNErrorCodeClientIdentifierInvalid CNErrorCode = 0
	// CNErrorCodeCommunicationError - An error that indicates a communication error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/communicationError
	CNErrorCodeCommunicationError CNErrorCode = 0
	// CNErrorCodeContainmentCycle - An error with the containment cycle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/containmentCycle
	CNErrorCodeContainmentCycle CNErrorCode = 0
	// CNErrorCodeContainmentScope - An error with containment scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/containmentScope
	CNErrorCodeContainmentScope CNErrorCode = 0
	// CNErrorCodeDataAccessError - An error with data access.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/dataAccessError
	CNErrorCodeDataAccessError CNErrorCode = 0
	// CNErrorCodeFeatureDisabledByUser - An error that indicates the user disabled the feature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/featureDisabledByUser
	CNErrorCodeFeatureDisabledByUser CNErrorCode = 0
	// CNErrorCodeFeatureNotAvailable - An error that indicates the feature isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/featureNotAvailable
	CNErrorCodeFeatureNotAvailable CNErrorCode = 0
	// CNErrorCodeInsertedRecordAlreadyExists - An error that indicates the inserted record already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/insertedRecordAlreadyExists
	CNErrorCodeInsertedRecordAlreadyExists CNErrorCode = 0
	// CNErrorCodeNoAccessableWritableContainers - An error that indicates there are no accessible writable containers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/noAccessableWritableContainers
	CNErrorCodeNoAccessableWritableContainers CNErrorCode = 0
	// CNErrorCodeParentContainerNotWritable - An error that indicates the parent container isn’t writable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/parentContainerNotWritable
	CNErrorCodeParentContainerNotWritable CNErrorCode = 0
	// CNErrorCodeParentRecordDoesNotExist - An error that indicates the parent record doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/parentRecordDoesNotExist
	CNErrorCodeParentRecordDoesNotExist CNErrorCode = 0
	// CNErrorCodePolicyViolation - An error that indicates a policy violation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/policyViolation
	CNErrorCodePolicyViolation CNErrorCode = 0
	// CNErrorCodePredicateInvalid - An error that indicates an invalid predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/predicateInvalid
	CNErrorCodePredicateInvalid CNErrorCode = 0
	// CNErrorCodeRecordDoesNotExist - An error that indicates a record doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/recordDoesNotExist
	CNErrorCodeRecordDoesNotExist CNErrorCode = 0
	// CNErrorCodeRecordIdentifierInvalid - An error that indicates a record identifier is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/recordIdentifierInvalid
	CNErrorCodeRecordIdentifierInvalid CNErrorCode = 0
	// CNErrorCodeRecordNotWritable - An error that indicates a record isn’t writable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/recordNotWritable
	CNErrorCodeRecordNotWritable CNErrorCode = 0
	// CNErrorCodeUnauthorizedKeys - An error that indicates unauthorized keys usage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/unauthorizedKeys
	CNErrorCodeUnauthorizedKeys CNErrorCode = 0
	// CNErrorCodeVCardMalformed - An error that indicates a malformed vCard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/vCardMalformed
	CNErrorCodeVCardMalformed CNErrorCode = 0
	// CNErrorCodeVCardSummarizationError - An error that indicates a vCard summarization problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/vCardSummarizationError
	CNErrorCodeVCardSummarizationError CNErrorCode = 0
	// CNErrorCodeValidationConfigurationError - An error with validation configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/validationConfigurationError
	CNErrorCodeValidationConfigurationError CNErrorCode = 0
	// CNErrorCodeValidationMultipleErrors - An error that indicates the system encountered multiple validation errors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/validationMultipleErrors
	CNErrorCodeValidationMultipleErrors CNErrorCode = 0
	// CNErrorCodeValidationTypeMismatch - A validation error that indicates a type mismatch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNError/Code/validationTypeMismatch
	CNErrorCodeValidationTypeMismatch CNErrorCode = 0
)

// CNPostalAddressFormatterStyle - Constants for postal formatting styles.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatterStyle
type CNPostalAddressFormatterStyle uint

const (
	// CNPostalAddressFormatterStyleMailingAddress - A style that combines the postal address components into a multi-line mailing address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddressFormatterStyle/mailingAddress
	CNPostalAddressFormatterStyleMailingAddress CNPostalAddressFormatterStyle = 0
)


