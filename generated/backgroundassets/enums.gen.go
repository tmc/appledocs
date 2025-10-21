// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

// Enum types and constants
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
// BAAssetPackStatusOutOfDate - A status value that indicates that the downloaded asset pack is out of date.
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusOutOfDate
BAAssetPackStatusOutOfDate BAAssetPackStatus = 0
// BAAssetPackStatusUpToDate - A status value that indicates that the downloaded asset pack is up to date.
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusUpToDate
BAAssetPackStatusUpToDate BAAssetPackStatus = 0
// BAAssetPackStatusUpdateAvailable - A status value that indicates that an update to the asset pack is available to download.
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackStatus/BAAssetPackStatusUpdateAvailable
BAAssetPackStatusUpdateAvailable BAAssetPackStatus = 0
)

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

// BAErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode
type BAErrorCode uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callFromExtensionNotAllowed
BAErrorCodeCallFromExtensionNotAllowed BAErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callFromInactiveProcessNotAllowed
BAErrorCodeCallFromInactiveProcessNotAllowed BAErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callerConnectionInvalid
BAErrorCodeCallerConnectionInvalid BAErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAErrorCode/callerConnectionNotAccepted
BAErrorCodeCallerConnectionNotAccepted BAErrorCode = 0
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

// BAManagedErrorCode - An error code for a managed asset pack.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAManagedErrorCode
type BAManagedErrorCode uint

const (
// BAManagedErrorCodeAssetPackNotFound - An error code that indicates the system can’t find an asset pack with the given identifier.
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAManagedErrorCode/BAManagedErrorCodeAssetPackNotFound
BAManagedErrorCodeAssetPackNotFound BAManagedErrorCode = 0
// BAManagedErrorCodeFileNotFound - An error code that indicates the system can’t find a file at the specified path.
//
	// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAManagedErrorCode/BAManagedErrorCodeFileNotFound
BAManagedErrorCodeFileNotFound BAManagedErrorCode = 0
)


