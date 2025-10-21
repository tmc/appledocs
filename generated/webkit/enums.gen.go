// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// Enum types and constants
// WKDialogResult - An enumeration that lists the possible ways a delegate handled displaying a custom Lockdown Mode first use dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult
type DialogResult uint

const (
	// DialogResultShowDefault - A result that indicates the delegate didn’t display a message, so the web view should show the default Lockdown Mode message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult/showDefault
	DialogResultShowDefault DialogResult = 0
)

// WKDownloadPlaceholderPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy
type DownloadPlaceholderPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy/disable
	DownloadPlaceholderPolicyDisable DownloadPlaceholderPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy/enable
	DownloadPlaceholderPolicyEnable DownloadPlaceholderPolicy = 0
)

// WKErrorCode - Possible error values that WebKit APIs can return.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code
type ErrorCode uint

const (
	// ErrorContentRuleListStoreRemoveFailed - An error that indicates a failure to remove a content rule list from the rule list data store object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreRemoveFailed
	ErrorContentRuleListStoreRemoveFailed ErrorCode = 0
)

// WKMediaCaptureState - An enumeration that describes whether a media device, like a camera or microphone, is currently capturing audio or video.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState
type MediaCaptureState uint

const (
	// MediaCaptureStateActive - The media device is actively capturing audio or video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/active
	MediaCaptureStateActive MediaCaptureState = 0
	// MediaCaptureStateMuted - The media device is muted, and not actively capturing audio or video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/muted
	MediaCaptureStateMuted MediaCaptureState = 0
	// MediaCaptureStateNone - The media device is off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState/none
	MediaCaptureStateNone MediaCaptureState = 0
)

// WKMediaPlaybackState - An enumeration that describes whether an audio or video presentation is playing, paused, or suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState
type MediaPlaybackState uint

const (
	// MediaPlaybackStateNone - There is no media to play back.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/none
	MediaPlaybackStateNone MediaPlaybackState = 0
	// MediaPlaybackStatePaused - The media playback is paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/paused
	MediaPlaybackStatePaused MediaPlaybackState = 0
	// MediaPlaybackStatePlaying - The media is playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/playing
	MediaPlaybackStatePlaying MediaPlaybackState = 0
	// MediaPlaybackStateSuspended - The media is not playing, and cannot be resumed until the user revokes the suspension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState/suspended
	MediaPlaybackStateSuspended MediaPlaybackState = 0
)

// WKNavigationActionPolicy - Constants that indicate whether to allow or cancel navigation to a webpage from an action.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
type NavigationActionPolicy uint

const (
	// NavigationActionPolicyAllow - Allow the navigation to continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy/allow
	NavigationActionPolicyAllow NavigationActionPolicy = 0
	// NavigationActionPolicyCancel - Cancel the navigation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy/cancel
	NavigationActionPolicyCancel NavigationActionPolicy = 0
)

// WKPermissionDecision - An enumeration of possible permission decisions for device resource access.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPermissionDecision
type PermissionDecision uint

// WKInactiveSchedulingPolicy - An enumeration that lists policies for how a web view that’s not in a window handles tasks.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum
type InactiveSchedulingPolicy uint

// WKUserInterfaceDirectionPolicy - The policy that determines the directionality of user interface elements in a web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy
type UserInterfaceDirectionPolicy uint

// WKUserScriptInjectionTime - Constants for the times at which to inject script content into a webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
type UserScriptInjectionTime uint

// WKFullscreenState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum
type FullscreenState uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/enteringFullscreen
	FullscreenStateEnteringFullscreen FullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/exitingFullscreen
	FullscreenStateExitingFullscreen FullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/inFullscreen
	FullscreenStateInFullscreen FullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/notInFullscreen
	FullscreenStateNotInFullscreen FullscreenState = 0
)

// WKWebViewDataType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType
type WebViewDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType/sessionStorage
	WebViewDataTypeSessionStorage WebViewDataType = 0
)

// WKContentMode - Constants that indicate how to render web view content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode
type ContentMode uint

// WKWebpagePreferencesUpgradeToHTTPSPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy
type WebpagePreferencesUpgradeToHTTPSPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/automaticFallbackToHTTP
	WebpagePreferencesUpgradeToHTTPSPolicyAutomaticFallbackToHTTP WebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/errorOnFailure
	WebpagePreferencesUpgradeToHTTPSPolicyErrorOnFailure WebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/keepAsRequested
	WebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested WebpagePreferencesUpgradeToHTTPSPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy/userMediatedFallbackToHTTP
	WebpagePreferencesUpgradeToHTTPSPolicyUserMediatedFallbackToHTTP WebpagePreferencesUpgradeToHTTPSPolicy = 0
)


