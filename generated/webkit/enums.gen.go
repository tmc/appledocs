// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// Enum types and constants
// WKDownloadPlaceholderPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy
type WKDownloadPlaceholderPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy/disable
	WKDownloadPlaceholderPolicyDisable WKDownloadPlaceholderPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy/enable
	WKDownloadPlaceholderPolicyEnable WKDownloadPlaceholderPolicy = 0
)

// WKErrorCode - Possible error values that WebKit APIs can return.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code
type WKErrorCode uint

const (
	// WKErrorContentRuleListStoreRemoveFailed - An error that indicates a failure to remove a content rule list from the rule list data store object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreRemoveFailed
	WKErrorContentRuleListStoreRemoveFailed WKErrorCode = 0
)

// WKMediaCaptureState - An enumeration that describes whether a media device, like a camera or microphone, is currently capturing audio or video.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState
type WKMediaCaptureState uint

const (
	// WKMediaCaptureStateActive - The media device is actively capturing audio or video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/active
	WKMediaCaptureStateActive WKMediaCaptureState = 0
	// WKMediaCaptureStateMuted - The media device is muted, and not actively capturing audio or video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/muted
	WKMediaCaptureStateMuted WKMediaCaptureState = 0
	// WKMediaCaptureStateNone - The media device is off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/none
	WKMediaCaptureStateNone WKMediaCaptureState = 0
)

// WKMediaPlaybackState - An enumeration that describes whether an audio or video presentation is playing, paused, or suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState
type WKMediaPlaybackState uint

const (
	// WKMediaPlaybackStateNone - There is no media to play back.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/none
	WKMediaPlaybackStateNone WKMediaPlaybackState = 0
	// WKMediaPlaybackStatePaused - The media playback is paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/paused
	WKMediaPlaybackStatePaused WKMediaPlaybackState = 0
	// WKMediaPlaybackStatePlaying - The media is playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/playing
	WKMediaPlaybackStatePlaying WKMediaPlaybackState = 0
	// WKMediaPlaybackStateSuspended - The media is not playing, and cannot be resumed until the user revokes the suspension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/suspended
	WKMediaPlaybackStateSuspended WKMediaPlaybackState = 0
)

// WKNavigationActionPolicy - Constants that indicate whether to allow or cancel navigation to a webpage from an action.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
type WKNavigationActionPolicy uint

// WKNavigationResponsePolicy - Constants that indicate whether to allow or cancel navigation to a webpage from a response.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy
type WKNavigationResponsePolicy uint

// WKWebViewDataType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType
type WKWebViewDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType/sessionStorage
	WKWebViewDataTypeSessionStorage WKWebViewDataType = 0
)

// WKContentMode - Constants that indicate how to render web view content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode
type WKContentMode uint

// WKWebpagePreferencesUpgradeToHTTPSPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy
type WKWebpagePreferencesUpgradeToHTTPSPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/automaticFallbackToHTTP
	WKWebpagePreferencesUpgradeToHTTPSPolicyAutomaticFallbackToHTTP WKWebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/errorOnFailure
	WKWebpagePreferencesUpgradeToHTTPSPolicyErrorOnFailure WKWebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/keepAsRequested
	WKWebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested WKWebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/userMediatedFallbackToHTTP
	WKWebpagePreferencesUpgradeToHTTPSPolicyUserMediatedFallbackToHTTP WKWebpagePreferencesUpgradeToHTTPSPolicy = 0
)

// WebNavigationType - Possible values for the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType
type WebNavigationType uint

// WebViewInsertAction - The type of user action that initiated a delegate message.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebViewInsertAction
type WebViewInsertAction uint


