// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

// Enum types and constants
// AMErrorCode - Automator error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code
type AMErrorCode uint

const (
	// AMActionApplicationResourceError - An error that indicates an app required by the action is not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionApplicationResourceError
	AMActionApplicationResourceError AMErrorCode = 0
	// AMActionApplicationVersionResourceError - An error that indicates an app required by the action is the wrong version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionApplicationVersionResourceError
	AMActionApplicationVersionResourceError AMErrorCode = 0
	// AMActionArchitectureMismatchError - An error that indicates the action’s binary is not compatible with the current processor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionArchitectureMismatchError
	AMActionArchitectureMismatchError AMErrorCode = 0
	// AMActionExceptionError - An error that indicates an action encounters an exception while running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionExceptionError
	AMActionExceptionError AMErrorCode = 0
	// AMActionExecutionError - An error that indicates an action encounters an error while running (reason unknown).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionExecutionError
	AMActionExecutionError AMErrorCode = 0
	// AMActionFailedGatekeeperError - An error that indicates the action doesn’t meet the Gatekeeper security policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionFailedGatekeeperError
	AMActionFailedGatekeeperError AMErrorCode = 0
	// AMActionFileResourceError - An error that indicates a file required by the action is not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionFileResourceError
	AMActionFileResourceError AMErrorCode = 0
	// AMActionInitializationError - An error that indicates Automator is unable to initialize an action (reason unknown).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionInitializationError
	AMActionInitializationError AMErrorCode = 0
	// AMActionInsufficientDataError - An error that indicates the action requires input data to run, but none was supplied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionInsufficientDataError
	AMActionInsufficientDataError AMErrorCode = 0
	// AMActionIsDeprecatedError - An error that indicates the action has been deprecated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionIsDeprecatedError
	AMActionIsDeprecatedError AMErrorCode = 0
	// AMActionLicenseResourceError - An error that indicates a license required by the action was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionLicenseResourceError
	AMActionLicenseResourceError AMErrorCode = 0
	// AMActionLinkError - An error that indicates the action’s executable failed to load due to linking issues.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionLinkError
	AMActionLinkError AMErrorCode = 0
	// AMActionLoadError - An error that indicates the action’s executable failed to load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionLoadError
	AMActionLoadError AMErrorCode = 0
	// AMActionMalwareError - An error that indicates the action has been identified as malware by XProtect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionMalwareError
	AMActionMalwareError AMErrorCode = 0
	// AMActionNotLoadableError - An error that indicates the action’s executable is of a type that is not loadable in the current process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionNotLoadableError
	AMActionNotLoadableError AMErrorCode = 0
	// AMActionPropertyListInvalidError - An error that indicates the property list for an action is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionPropertyListInvalidError
	AMActionPropertyListInvalidError AMErrorCode = 0
	// AMActionQuarantineError - An error that indicates action has been quarantined by XProtect, the antimalware system on the Mac.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionQuarantineError
	AMActionQuarantineError AMErrorCode = 0
	// AMActionRequiredActionResourceError - An error that indicates an action required by the action is not loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionRequiredActionResourceError
	AMActionRequiredActionResourceError AMErrorCode = 0
	// AMActionRuntimeMismatchError - An error that indicates an attempt was made to load an action that is not compiled in a way that is compatible with the current app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionRuntimeMismatchError
	AMActionRuntimeMismatchError AMErrorCode = 0
	// AMActionSignatureCorruptError - An error that indicates developer signature for this action is corrupted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionSignatureCorruptError
	AMActionSignatureCorruptError AMErrorCode = 0
	// AMActionThirdPartyActionsNotAllowedError - An error that indicates the action is a third party action, and loading it has not been allowed by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionThirdPartyActionsNotAllowedError
	AMActionThirdPartyActionsNotAllowedError AMErrorCode = 0
	// AMActionXPCError - An error that indicates the remote process running the action has crashed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionXPCError
	AMActionXPCError AMErrorCode = 0
	// AMActionXProtectError - An error that indicates XProtect is unable to successfully analyze the action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/actionXProtectError
	AMActionXProtectError AMErrorCode = 0
	// AMConversionFailedError - An error that occurs when, for example, the converter encounters an error converting data from one type to another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/conversionFailedError
	AMConversionFailedError AMErrorCode = 0
	// AMConversionNoDataError - An error that occurs when the converter determines that the conversion, though possible, would produce a nil result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/conversionNoDataError
	AMConversionNoDataError AMErrorCode = 0
	// AMConversionNotPossibleError - An error that occurs when the converter determines that it is unable to convert from one data type to another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/conversionNotPossibleError
	AMConversionNotPossibleError AMErrorCode = 0
	// AMNoSuchActionError - An error that indicates the action could not be located on the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/noSuchActionError
	AMNoSuchActionError AMErrorCode = 0
	// AMUserCanceledError - An error that indicates the user cancelled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/userCanceledError
	AMUserCanceledError AMErrorCode = 0
	// AMWorkflowActionsNotLoadedError - An error that indicates one of the actions of the workflow couldn’t be loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowActionsNotLoadedError
	AMWorkflowActionsNotLoadedError AMErrorCode = 0
	// AMWorkflowNewerActionVersionError - An error that indicates an action in a workflow is newer than the installed action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowNewerActionVersionError
	AMWorkflowNewerActionVersionError AMErrorCode = 0
	// AMWorkflowNewerVersionError - An error that indicates an attempt to open a workflow document that was saved with a newer version of Automator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowNewerVersionError
	AMWorkflowNewerVersionError AMErrorCode = 0
	// AMWorkflowNoEnabledActionsError - An error that indicates there are no enabled actions in the workflow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowNoEnabledActionsError
	AMWorkflowNoEnabledActionsError AMErrorCode = 0
	// AMWorkflowOlderActionVersionError - An error that indicates an action in a workflow is older than the installed action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowOlderActionVersionError
	AMWorkflowOlderActionVersionError AMErrorCode = 0
	// AMWorkflowPropertyListInvalidError - An error that indicates an attempt to open a workflow document whose property list couldn’t be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMError/Code/workflowPropertyListInvalidError
	AMWorkflowPropertyListInvalidError AMErrorCode = 0
)

// AMLogLevel - Logging levels that Automator supports.
//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMLogLevel
type AMLogLevel uint

const (
	// AMLogLevelDebug - The debug log level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMLogLevel/debug
	AMLogLevelDebug AMLogLevel = 0
	// AMLogLevelError - The error log level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMLogLevel/error
	AMLogLevelError AMLogLevel = 0
	// AMLogLevelInfo - The informational log level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMLogLevel/info
	AMLogLevelInfo AMLogLevel = 0
	// AMLogLevelWarn - The warning log level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Automator/AMLogLevel/warn
	AMLogLevelWarn AMLogLevel = 0
)


