// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

/* debug [enums.gen.go]: Generating 2 enums for Accounts */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ACErrorCode (23 cases) */
// ACErrorCode - Codes for errors that may occur.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorCode
type ACErrorCode uint

const (
	// ACErrorAccessDeniedByProtectionPolicy - Error code that indicates due to the current protection policy, the credentials couldn’t be fetched.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccessDeniedByProtectionPolicy
	ACErrorAccessDeniedByProtectionPolicy ACErrorCode = 0
	// ACErrorAccessInfoInvalid - Error code that indicates the client’s access info dictionary has incorrect or missing values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccessInfoInvalid
	ACErrorAccessInfoInvalid ACErrorCode = 0
	// ACErrorAccountAlreadyExists - Error code that indicates an account wasn’t added because it already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccountAlreadyExists
	ACErrorAccountAlreadyExists ACErrorCode = 0
	// ACErrorAccountAuthenticationFailed - Error code that indicates an account wasn’t saved because authentication of its credential failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccountAuthenticationFailed
	ACErrorAccountAuthenticationFailed ACErrorCode = 0
	// ACErrorAccountMissingRequiredProperty - Error code that indicates an account wasn’t saved because a required property is missing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccountMissingRequiredProperty
	ACErrorAccountMissingRequiredProperty ACErrorCode = 0
	// ACErrorAccountNotFound - Error code that indicates an account wasn’t deleted because it couldn’t be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccountNotFound
	ACErrorAccountNotFound ACErrorCode = 0
	// ACErrorAccountTypeInvalid - Error code that indicates an account wasn’t saved because its account type is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorAccountTypeInvalid
	ACErrorAccountTypeInvalid ACErrorCode = 0
	// ACErrorClientPermissionDenied - Error code that indicates the client doesn’t have access to the requested data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorClientPermissionDenied
	ACErrorClientPermissionDenied ACErrorCode = 0
	// ACErrorCoreDataSaveFailed - Error code that indicates an error occurred while trying to save to a Core Data store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorCoreDataSaveFailed
	ACErrorCoreDataSaveFailed ACErrorCode = 0
	// ACErrorCredentialItemNotExpired - Error code that indicates a credential item wasn’t removed because it hasn’t yet expired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorCredentialItemNotExpired
	ACErrorCredentialItemNotExpired ACErrorCode = 0
	// ACErrorCredentialItemNotFound - Error code that indicates a credential item wasn’t saved because it couldn’t be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorCredentialItemNotFound
	ACErrorCredentialItemNotFound ACErrorCode = 0
	// ACErrorCredentialNotFound - Error code that indicates no credentials were found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorCredentialNotFound
	ACErrorCredentialNotFound ACErrorCode = 0
	// ACErrorDeniedByPlugin - Error code that indicates a plugin prevented the expected action from occurring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorDeniedByPlugin
	ACErrorDeniedByPlugin ACErrorCode = 0
	// ACErrorFailedSerializingAccountInfo - Error code that indicates an account’s information couldn’t be serialized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorFailedSerializingAccountInfo
	ACErrorFailedSerializingAccountInfo ACErrorCode = 0
	// ACErrorFetchCredentialFailed - Error code that indicates the credentials couldn’t be fetched from Keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorFetchCredentialFailed
	ACErrorFetchCredentialFailed ACErrorCode = 0
	// ACErrorInvalidClientBundleID - Error code that indicates the client making the request doesn’t have a valid bundle ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorInvalidClientBundleID
	ACErrorInvalidClientBundleID ACErrorCode = 0
	// ACErrorInvalidCommand - Error code that indicates an invalid command was attempted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorInvalidCommand
	ACErrorInvalidCommand ACErrorCode = 0
	// ACErrorMissingTransportMessageID - Error code that indicates an expected message identifier wasn’t found while performing a command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorMissingTransportMessageID
	ACErrorMissingTransportMessageID ACErrorCode = 0
	// ACErrorPermissionDenied - Error code that indicates the operation failed because the application doesn’t have permission to perform the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorPermissionDenied
	ACErrorPermissionDenied ACErrorCode = 0
	// ACErrorRemoveCredentialFailed - Error code that indicates the credentials couldn’t be removed from Keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorRemoveCredentialFailed
	ACErrorRemoveCredentialFailed ACErrorCode = 0
	// ACErrorStoreCredentialFailed - Error code that indicates the credentials couldn’t be stored in Keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorStoreCredentialFailed
	ACErrorStoreCredentialFailed ACErrorCode = 0
	// ACErrorUnknown - Error code that indicates an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorUnknown
	ACErrorUnknown ACErrorCode = 0
	// ACErrorUpdatingNonexistentAccount - Error code that indicates an account save failed because the account being updated has been removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACErrorUpdatingNonexistentAccount
	ACErrorUpdatingNonexistentAccount ACErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ACAccountCredentialRenewResult (3 cases) */
// ACAccountCredentialRenewResult - Status codes of credential renewal requests.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredentialRenewResult
type ACAccountCredentialRenewResult uint

const (
	// ACAccountCredentialRenewResultFailed - A non-user-initiated cancel of the prompt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredentialRenewResult/failed
	ACAccountCredentialRenewResultFailed ACAccountCredentialRenewResult = 0
	// ACAccountCredentialRenewResultRejected - Renewal failed because the user revoked your access to their account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredentialRenewResult/rejected
	ACAccountCredentialRenewResultRejected ACAccountCredentialRenewResult = 0
	// ACAccountCredentialRenewResultRenewed - The account’s credentials have been renewed and are now associated with the account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredentialRenewResult/renewed
	ACAccountCredentialRenewResultRenewed ACAccountCredentialRenewResult = 0
)


