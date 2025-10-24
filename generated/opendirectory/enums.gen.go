// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

/* debug [enums.gen.go]: Generating 1 enums for OpenDirectory */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ODFrameworkErrors (64 cases) */
// ODFrameworkErrors enum type
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODFrameworkErrors
type ODFrameworkErrors uint

const (
	// kODErrorCredentialsContactMaster - The authentication server contacted is not the primary server, and the requested operation requires the primary server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODFrameworkErrors/kODErrorCredentialsContactMaster
	kODErrorCredentialsContactMaster ODFrameworkErrors = 0
	// kODErrorCredentialsAccountDisabled - The account is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountDisabled
	kODErrorCredentialsAccountDisabled ODFrameworkErrors = 0
	// kODErrorCredentialsAccountExpired - The account is expired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountExpired
	kODErrorCredentialsAccountExpired ODFrameworkErrors = 0
	// kODErrorCredentialsAccountInactive - The account is inactive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountInactive
	kODErrorCredentialsAccountInactive ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountLocked
	kODErrorCredentialsAccountLocked ODFrameworkErrors = 0
	// kODErrorCredentialsAccountNotFound - The authentication server could not find the provided account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountNotFound
	kODErrorCredentialsAccountNotFound ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsAccountTemporarilyLocked
	kODErrorCredentialsAccountTemporarilyLocked ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsContactPrimary
	kODErrorCredentialsContactPrimary ODFrameworkErrors = 0
	// kODErrorCredentialsInvalid - The provided credentials are invalid with the current node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsInvalid
	kODErrorCredentialsInvalid ODFrameworkErrors = 0
	// kODErrorCredentialsInvalidComputer - The account is not permitted to log into this computer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsInvalidComputer
	kODErrorCredentialsInvalidComputer ODFrameworkErrors = 0
	// kODErrorCredentialsInvalidLogonHours - The logon attempt was not within set logon hours.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsInvalidLogonHours
	kODErrorCredentialsInvalidLogonHours ODFrameworkErrors = 0
	// kODErrorCredentialsMethodNotSupported - The extended authentication method is not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsMethodNotSupported
	kODErrorCredentialsMethodNotSupported ODFrameworkErrors = 0
	// kODErrorCredentialsNotAuthorized - The operation, such as changing a password, is not permitted with current privileges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsNotAuthorized
	kODErrorCredentialsNotAuthorized ODFrameworkErrors = 0
	// kODErrorCredentialsOperationFailed - The requested operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsOperationFailed
	kODErrorCredentialsOperationFailed ODFrameworkErrors = 0
	// kODErrorCredentialsParameterError - An invalid parameter was provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsParameterError
	kODErrorCredentialsParameterError ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordChangeRequired - The password must be changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordChangeRequired
	kODErrorCredentialsPasswordChangeRequired ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordChangeTooSoon - The password was changed too recently to be changed again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordChangeTooSoon
	kODErrorCredentialsPasswordChangeTooSoon ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordExpired - The password has expired and must be changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordExpired
	kODErrorCredentialsPasswordExpired ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordNeedsDigit - The provided password needs at least one digit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordNeedsDigit
	kODErrorCredentialsPasswordNeedsDigit ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordNeedsLetter - The provided password needs at least one letter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordNeedsLetter
	kODErrorCredentialsPasswordNeedsLetter ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordQualityFailed - The provided password did not meet minimum quality requirements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordQualityFailed
	kODErrorCredentialsPasswordQualityFailed ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordTooLong - The provided password is too long.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordTooLong
	kODErrorCredentialsPasswordTooLong ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordTooShort - The provided password is too short.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordTooShort
	kODErrorCredentialsPasswordTooShort ODFrameworkErrors = 0
	// kODErrorCredentialsPasswordUnrecoverable - The password could not be recovered from the authentication database.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsPasswordUnrecoverable
	kODErrorCredentialsPasswordUnrecoverable ODFrameworkErrors = 0
	// kODErrorCredentialsServerCommunicationError - The authentication server encountered a communication error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsServerCommunicationError
	kODErrorCredentialsServerCommunicationError ODFrameworkErrors = 0
	// kODErrorCredentialsServerError - The authentication server encountered an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsServerError
	kODErrorCredentialsServerError ODFrameworkErrors = 0
	// kODErrorCredentialsServerNotFound - The authentication server could not be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsServerNotFound
	kODErrorCredentialsServerNotFound ODFrameworkErrors = 0
	// kODErrorCredentialsServerTimeout - The authentication server timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsServerTimeout
	kODErrorCredentialsServerTimeout ODFrameworkErrors = 0
	// kODErrorCredentialsServerUnreachable - The authentication server could not be reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorCredentialsServerUnreachable
	kODErrorCredentialsServerUnreachable ODFrameworkErrors = 0
	// kODErrorDaemonError - The daemon has encountered an undefined error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorDaemonError
	kODErrorDaemonError ODFrameworkErrors = 0
	// kODErrorNodeConnectionFailed - The node connection failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorNodeConnectionFailed
	kODErrorNodeConnectionFailed ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorNodeDisabled
	kODErrorNodeDisabled ODFrameworkErrors = 0
	// kODErrorNodeUnknownHost - The host provided is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorNodeUnknownHost
	kODErrorNodeUnknownHost ODFrameworkErrors = 0
	// kODErrorNodeUnknownName - The node name provided does not exist and cannot be opened.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorNodeUnknownName
	kODErrorNodeUnknownName ODFrameworkErrors = 0
	// kODErrorNodeUnknownType - The node type provided is not a known value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorNodeUnknownType
	kODErrorNodeUnknownType ODFrameworkErrors = 0
	// kODErrorPluginError - A plug-in has encountered an undefined error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorPluginError
	kODErrorPluginError ODFrameworkErrors = 0
	// kODErrorPluginOperationNotSupported - The plug-in does not support the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorPluginOperationNotSupported
	kODErrorPluginOperationNotSupported ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorPluginOperationTimeout
	kODErrorPluginOperationTimeout ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorPolicyOutOfRange
	kODErrorPolicyOutOfRange ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorPolicyUnsupported
	kODErrorPolicyUnsupported ODFrameworkErrors = 0
	// kODErrorQueryInvalidMatchType - An invalid match type was provided in the query.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorQueryInvalidMatchType
	kODErrorQueryInvalidMatchType ODFrameworkErrors = 0
	// kODErrorQuerySynchronize - A query synchronization has been initiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorQuerySynchronize
	kODErrorQuerySynchronize ODFrameworkErrors = 0
	// kODErrorQueryTimeout - The query timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorQueryTimeout
	kODErrorQueryTimeout ODFrameworkErrors = 0
	// kODErrorQueryUnsupportedMatchType - An unsupported match type was provided in the query.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorQueryUnsupportedMatchType
	kODErrorQueryUnsupportedMatchType ODFrameworkErrors = 0
	// kODErrorRecordAlreadyExists - The record create failed because the record already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordAlreadyExists
	kODErrorRecordAlreadyExists ODFrameworkErrors = 0
	// kODErrorRecordAttributeNotFound - The requested attribute could not be found in the record.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordAttributeNotFound
	kODErrorRecordAttributeNotFound ODFrameworkErrors = 0
	// kODErrorRecordAttributeUnknownType - The attribute type is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordAttributeUnknownType
	kODErrorRecordAttributeUnknownType ODFrameworkErrors = 0
	// kODErrorRecordAttributeValueNotFound - The requested attribute value could not be found in the record.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordAttributeValueNotFound
	kODErrorRecordAttributeValueNotFound ODFrameworkErrors = 0
	// kODErrorRecordAttributeValueSchemaError - The attribute value does not meet schema requirements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordAttributeValueSchemaError
	kODErrorRecordAttributeValueSchemaError ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordInvalidType
	kODErrorRecordInvalidType ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordNoLongerExists
	kODErrorRecordNoLongerExists ODFrameworkErrors = 0
	// kODErrorRecordParameterError - An invalid parameter was provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordParameterError
	kODErrorRecordParameterError ODFrameworkErrors = 0
	// kODErrorRecordPermissionError - The changes were denied due to insufficient permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordPermissionError
	kODErrorRecordPermissionError ODFrameworkErrors = 0
	// kODErrorRecordReadOnlyNode - The record cannot be modified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordReadOnlyNode
	kODErrorRecordReadOnlyNode ODFrameworkErrors = 0
	// kODErrorRecordTypeDisabled - The record type is disabled by policy for a plug-in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorRecordTypeDisabled
	kODErrorRecordTypeDisabled ODFrameworkErrors = 0
	// kODErrorSessionDaemonNotRunning - The daemon is not running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionDaemonNotRunning
	kODErrorSessionDaemonNotRunning ODFrameworkErrors = 0
	// kODErrorSessionDaemonRefused - The daemon refused the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionDaemonRefused
	kODErrorSessionDaemonRefused ODFrameworkErrors = 0
	// kODErrorSessionLocalOnlyDaemonInUse - A normal request was issued when the local-only daemon was in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionLocalOnlyDaemonInUse
	kODErrorSessionLocalOnlyDaemonInUse ODFrameworkErrors = 0
	// kODErrorSessionNormalDaemonInUse - A local-only request was issued when the normal daemon was in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionNormalDaemonInUse
	kODErrorSessionNormalDaemonInUse ODFrameworkErrors = 0
	// kODErrorSessionProxyCommunicationError - There was a communication error with the remote daemon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionProxyCommunicationError
	kODErrorSessionProxyCommunicationError ODFrameworkErrors = 0
	// kODErrorSessionProxyIPUnreachable - The proxy did not respond.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionProxyIPUnreachable
	kODErrorSessionProxyIPUnreachable ODFrameworkErrors = 0
	// kODErrorSessionProxyUnknownHost - The proxy could not be resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionProxyUnknownHost
	kODErrorSessionProxyUnknownHost ODFrameworkErrors = 0
	// kODErrorSessionProxyVersionMismatch - Versions mismatch between the remote daemon and the local framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSessionProxyVersionMismatch
	kODErrorSessionProxyVersionMismatch ODFrameworkErrors = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/kODErrorSuccess
	kODErrorSuccess ODFrameworkErrors = 0
)


