// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

/* debug [enums.gen.go]: Generating 5 enums for BackgroundAssets */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum BAAssetPackStatus (4 cases) */
// BAAssetPackStatus - The status of an asset pack.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus
type BAAssetPackStatus uint

const (
	// BAAssetPackStatusDownloadAvailable - A status value that indicates that the asset pack is available to download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusDownloadAvailable
	BAAssetPackStatusDownloadAvailable BAAssetPackStatus = 0
	// BAAssetPackStatusDownloaded - A status value that indicates that the system finished downloading the asset pack.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusDownloaded
	BAAssetPackStatusDownloaded BAAssetPackStatus = 0
	// BAAssetPackStatusDownloading - A status value that indicates that the system is currently downloading the asset pack.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusDownloading
	BAAssetPackStatusDownloading BAAssetPackStatus = 0
	// BAAssetPackStatusObsolete - A status value that indicates that the asset pack is no longer available to download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusObsolete
	BAAssetPackStatusObsolete BAAssetPackStatus = 0
)

/* debug [enums.gen.go]: Processing enum BADownloadState (5 cases) */
// BADownloadState - Constants that indicate the state of a download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum
type BADownloadState uint

const (
	// BADownloadStateCreated - A state that indicates a created download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum/created
	BADownloadStateCreated BADownloadState = 0
	// BADownloadStateDownloading - A state that indicates a download is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum/downloading
	BADownloadStateDownloading BADownloadState = 0
	// BADownloadStateFailed - A state that indicates a failed download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum/failed
	BADownloadStateFailed BADownloadState = 0
	// BADownloadStateFinished - A state that indicates a finished download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum/finished
	BADownloadStateFinished BADownloadState = 0
	// BADownloadStateWaiting - A state that indicates a download is waiting to execute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/State-swift.enum/waiting
	BADownloadStateWaiting BADownloadState = 0
)

/* debug [enums.gen.go]: Processing enum BAErrorCode (17 cases) */
// BAErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode
type BAErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callerConnectionInvalid
	BAErrorCodeCallerConnectionInvalid BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callerConnectionNotAccepted
	BAErrorCodeCallerConnectionNotAccepted BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callFromExtensionNotAllowed
	BAErrorCodeCallFromExtensionNotAllowed BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callFromInactiveProcessNotAllowed
	BAErrorCodeCallFromInactiveProcessNotAllowed BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadAlreadyFailed
	BAErrorCodeDownloadAlreadyFailed BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadAlreadyScheduled
	BAErrorCodeDownloadAlreadyScheduled BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadBackgroundActivityProhibited
	BAErrorCodeDownloadBackgroundActivityProhibited BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadDoesNotExist
	BAErrorCodeDownloadDoesNotExist BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadEssentialDownloadNotPermitted
	BAErrorCodeDownloadEssentialDownloadNotPermitted BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadFailedToStart
	BAErrorCodeDownloadFailedToStart BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadInvalid
	BAErrorCodeDownloadInvalid BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadNotScheduled
	BAErrorCodeDownloadNotScheduled BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/downloadWouldExceedAllowance
	BAErrorCodeDownloadWouldExceedAllowance BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/sessionDownloadAllowanceExceeded
	BAErrorCodeSessionDownloadAllowanceExceeded BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/sessionDownloadDisallowedByAllowance
	BAErrorCodeSessionDownloadDisallowedByAllowance BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/sessionDownloadDisallowedByDomain
	BAErrorCodeSessionDownloadDisallowedByDomain BAErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/sessionDownloadNotPermittedBeforeAppLaunch
	BAErrorCodeSessionDownloadNotPermittedBeforeAppLaunch BAErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum BAManagedErrorCode (0 cases) */
// BAManagedErrorCode - An error code for a managed asset pack.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAManagedErrorCode
type BAManagedErrorCode int

/* debug [enums.gen.go]: Processing enum BAContentRequest (3 cases) */
// BAContentRequest enum type
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAContentRequest
type BAContentRequest uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAContentRequest/install
	BAContentRequestInstall BAContentRequest = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAContentRequest/periodic
	BAContentRequestPeriodic BAContentRequest = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAContentRequest/update
	BAContentRequestUpdate BAContentRequest = 0
)


