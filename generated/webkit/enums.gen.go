// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

/* debug [enums.gen.go]: Generating 41 enums for WebKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum WKDownloadPlaceholderPolicy (2 cases) */
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

/* debug [enums.gen.go]: Processing enum WKDownloadRedirectPolicy (2 cases) */
// WKDownloadRedirectPolicy - An enumeration with cases that indicate whether to proceed with a redirect.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/RedirectPolicy
type WKDownloadRedirectPolicy uint

const (
	// WKDownloadRedirectPolicyAllow - Allow a redirect to proceed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/RedirectPolicy/allow
	WKDownloadRedirectPolicyAllow WKDownloadRedirectPolicy = 0
	// WKDownloadRedirectPolicyCancel - Cancel the redirect action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/RedirectPolicy/cancel
	WKDownloadRedirectPolicyCancel WKDownloadRedirectPolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKWebpagePreferencesUpgradeToHTTPSPolicy (4 cases) */
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

/* debug [enums.gen.go]: Processing enum WKContentMode (3 cases) */
// WKContentMode - Constants that indicate how to render web view content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode
type WKContentMode uint

const (
	// WKContentModeDesktop - The content mode that represents a desktop experience.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode/desktop
	WKContentModeDesktop WKContentMode = 0
	// WKContentModeMobile - The content mode that represents a mobile experience.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode/mobile
	WKContentModeMobile WKContentMode = 0
	// WKContentModeRecommended - The content mode that is appropriate for the current device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode/recommended
	WKContentModeRecommended WKContentMode = 0
)

/* debug [enums.gen.go]: Processing enum DOMEventExceptionCode (1 cases) */
// DOMEventExceptionCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEventExceptionCode
type DOMEventExceptionCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_UNSPECIFIED_EVENT_TYPE_ERR
	DOM_UNSPECIFIED_EVENT_TYPE_ERR DOMEventExceptionCode = 0
)

/* debug [enums.gen.go]: Processing enum DOMExceptionCode (15 cases) */
// DOMExceptionCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMExceptionCode
type DOMExceptionCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_DOMSTRING_SIZE_ERR
	DOM_DOMSTRING_SIZE_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_HIERARCHY_REQUEST_ERR
	DOM_HIERARCHY_REQUEST_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INDEX_SIZE_ERR
	DOM_INDEX_SIZE_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INUSE_ATTRIBUTE_ERR
	DOM_INUSE_ATTRIBUTE_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_ACCESS_ERR
	DOM_INVALID_ACCESS_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_CHARACTER_ERR
	DOM_INVALID_CHARACTER_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_MODIFICATION_ERR
	DOM_INVALID_MODIFICATION_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_STATE_ERR
	DOM_INVALID_STATE_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_NAMESPACE_ERR
	DOM_NAMESPACE_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_NO_DATA_ALLOWED_ERR
	DOM_NO_DATA_ALLOWED_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_NO_MODIFICATION_ALLOWED_ERR
	DOM_NO_MODIFICATION_ALLOWED_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_NOT_FOUND_ERR
	DOM_NOT_FOUND_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_NOT_SUPPORTED_ERR
	DOM_NOT_SUPPORTED_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_SYNTAX_ERR
	DOM_SYNTAX_ERR DOMExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_WRONG_DOCUMENT_ERR
	DOM_WRONG_DOCUMENT_ERR DOMExceptionCode = 0
)

/* debug [enums.gen.go]: Processing enum DOMRangeExceptionCode (2 cases) */
// DOMRangeExceptionCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRangeExceptionCode
type DOMRangeExceptionCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_BAD_BOUNDARYPOINTS_ERR
	DOM_BAD_BOUNDARYPOINTS_ERR DOMRangeExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_NODE_TYPE_ERR
	DOM_INVALID_NODE_TYPE_ERR DOMRangeExceptionCode = 0
)

/* debug [enums.gen.go]: Processing enum DOMXPathExceptionCode (2 cases) */
// DOMXPathExceptionCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathExceptionCode
type DOMXPathExceptionCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_INVALID_EXPRESSION_ERR
	DOM_INVALID_EXPRESSION_ERR DOMXPathExceptionCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOM_TYPE_ERR
	DOM_TYPE_ERR DOMXPathExceptionCode = 0
)

/* debug [enums.gen.go]: Processing enum WebCacheModel (3 cases) */
// WebCacheModel - Specifies the caching model for a web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebCacheModel
type WebCacheModel uint

const (
	// WebCacheModelDocumentBrowser - Caches a reasonable number of resources and previously viewed documents in memory and on disk. This model is appropriate for displaying and navigating between multiple documents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebCacheModel/documentBrowser
	WebCacheModelDocumentBrowser WebCacheModel = 0
	// WebCacheModelDocumentViewer - Releases resources when they are no longer referenced and caches remote resources on disk. This model is appropriate for displaying a static document with no navigation user interface. This is the most memory-efficient model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebCacheModel/documentViewer
	WebCacheModelDocumentViewer WebCacheModel = 0
	// WebCacheModelPrimaryWebBrowser - Caches a large number of resources and previously viewed documents in memory and on disk. This model is appropriate for a web view that behaves like a web browser.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebCacheModel/primaryWebBrowser
	WebCacheModelPrimaryWebBrowser WebCacheModel = 0
)

/* debug [enums.gen.go]: Processing enum WebDragDestinationAction (5 cases) */
// WebDragDestinationAction - Actions that the destination object of a drag operation can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction
type WebDragDestinationAction uint

const (
	// WebDragDestinationActionAny - Allows any defined action to occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction/any
	WebDragDestinationActionAny WebDragDestinationAction = 0
	// WebDragDestinationActionDHTML - Allows DHTML (such as JavaScript) to handle the drag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction/DHTML
	WebDragDestinationActionDHTML WebDragDestinationAction = 0
	// WebDragDestinationActionEdit - Allows editable documents to be changed by the drag operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction/edit
	WebDragDestinationActionEdit WebDragDestinationAction = 0
	// WebDragDestinationActionLoad - Allows the drag operation to change the location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction/load
	WebDragDestinationActionLoad WebDragDestinationAction = 0
	// WebDragDestinationActionNone - No action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction/WebDragDestinationActionNone
	WebDragDestinationActionNone WebDragDestinationAction = 0
)

/* debug [enums.gen.go]: Processing enum WebDragSourceAction (6 cases) */
// WebDragSourceAction - Actions that the source object of a drag operation can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction
type WebDragSourceAction uint

const (
	// WebDragSourceActionAny - Allows any defined action to occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/any
	WebDragSourceActionAny WebDragSourceAction = 0
	// WebDragSourceActionDHTML - Allows DHTML (such as JavaScript) in the source object to initiate a drag operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/DHTML
	WebDragSourceActionDHTML WebDragSourceAction = 0
	// WebDragSourceActionImage - Allows the user to drag an image in the source object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/image
	WebDragSourceActionImage WebDragSourceAction = 0
	// WebDragSourceActionLink - Allows the user to drag a link in the source object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/link
	WebDragSourceActionLink WebDragSourceAction = 0
	// WebDragSourceActionSelection - Allows the user to drag a selection in the source object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/selection
	WebDragSourceActionSelection WebDragSourceAction = 0
	// WebDragSourceActionNone - No action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDragSourceAction/WebDragSourceActionNone
	WebDragSourceActionNone WebDragSourceAction = 0
)

/* debug [enums.gen.go]: Processing enum WebNavigationType (6 cases) */
// WebNavigationType - Possible values for the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType
type WebNavigationType uint

const (
	// WebNavigationTypeBackForward - The user clicked back or forward button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/backForward
	WebNavigationTypeBackForward WebNavigationType = 0
	// WebNavigationTypeFormResubmitted - A form was resubmitted (through a back, forward or reload action).
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/formResubmitted
	WebNavigationTypeFormResubmitted WebNavigationType = 0
	// WebNavigationTypeFormSubmitted - A form was submitted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/formSubmitted
	WebNavigationTypeFormSubmitted WebNavigationType = 0
	// WebNavigationTypeLinkClicked - A link (an  ) was clicked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/linkClicked
	WebNavigationTypeLinkClicked WebNavigationType = 0
	// WebNavigationTypeOther - Navigation is taking place for some other reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/other
	WebNavigationTypeOther WebNavigationType = 0
	// WebNavigationTypeReload - The user hit the reload button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebNavigationType/reload
	WebNavigationTypeReload WebNavigationType = 0
)

/* debug [enums.gen.go]: Processing enum WebViewInsertAction (3 cases) */
// WebViewInsertAction - The type of user action that initiated a delegate message.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebViewInsertAction
type WebViewInsertAction uint

const (
	// WebViewInsertActionDropped - Indicates the user inserted content by dropping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebViewInsertAction/dropped
	WebViewInsertActionDropped WebViewInsertAction = 0
	// WebViewInsertActionPasted - Indicates the user inserted content by pasting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebViewInsertAction/pasted
	WebViewInsertActionPasted WebViewInsertAction = 0
	// WebViewInsertActionTyped - Indicates the user inserted content by typing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebViewInsertAction/typed
	WebViewInsertActionTyped WebViewInsertAction = 0
)

/* debug [enums.gen.go]: Processing enum WKAudiovisualMediaTypes (4 cases) */
// WKAudiovisualMediaTypes - The media types that require a user gesture to begin playing.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes
type WKAudiovisualMediaTypes uint

const (
	// WKAudiovisualMediaTypeAll - All media types require a user gesture to begin playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes/all
	WKAudiovisualMediaTypeAll WKAudiovisualMediaTypes = 0
	// WKAudiovisualMediaTypeAudio - Media types that contain audio require a user gesture to begin playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes/audio
	WKAudiovisualMediaTypeAudio WKAudiovisualMediaTypes = 0
	// WKAudiovisualMediaTypeVideo - Media types that contain video require a user gesture to begin playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes/video
	WKAudiovisualMediaTypeVideo WKAudiovisualMediaTypes = 0
	// WKAudiovisualMediaTypeNone - No media types require a user gesture to begin playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes/WKAudiovisualMediaTypeNone
	WKAudiovisualMediaTypeNone WKAudiovisualMediaTypes = 0
)

/* debug [enums.gen.go]: Processing enum WKDataDetectorTypes (10 cases) */
// WKDataDetectorTypes - The data detector types.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes
type WKDataDetectorTypes uint

const (
	// WKDataDetectorTypeAddress - Detect addresses in text and turn them into links to display the location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/address
	WKDataDetectorTypeAddress WKDataDetectorTypes = 0
	// WKDataDetectorTypeAll - Detect all data types and turn them into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/all
	WKDataDetectorTypeAll WKDataDetectorTypes = 0
	// WKDataDetectorTypeCalendarEvent - Turn future dates and times into links to create calendar events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/calendarEvent
	WKDataDetectorTypeCalendarEvent WKDataDetectorTypes = 0
	// WKDataDetectorTypeFlightNumber - Detect flight numbers in text and turn them into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/flightNumber
	WKDataDetectorTypeFlightNumber WKDataDetectorTypes = 0
	// WKDataDetectorTypeLink - Detect URLs in text and turn them into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/link
	WKDataDetectorTypeLink WKDataDetectorTypes = 0
	// WKDataDetectorTypeLookupSuggestion - Detect Spotlight suggestions and turn them into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/lookupSuggestion
	WKDataDetectorTypeLookupSuggestion WKDataDetectorTypes = 0
	// WKDataDetectorTypePhoneNumber - Detect phone numbers in text and create a link to call the specified number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/phoneNumber
	WKDataDetectorTypePhoneNumber WKDataDetectorTypes = 0
	// WKDataDetectorTypeSpotlightSuggestion - Spotlight suggestions are detected and turned into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/spotlightSuggestion
	WKDataDetectorTypeSpotlightSuggestion WKDataDetectorTypes = 0
	// WKDataDetectorTypeTrackingNumber - Detect tracking numbers in text and turn them into links.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/trackingNumber
	WKDataDetectorTypeTrackingNumber WKDataDetectorTypes = 0
	// WKDataDetectorTypeNone - No data detection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes/WKDataDetectorTypeNone
	WKDataDetectorTypeNone WKDataDetectorTypes = 0
)

/* debug [enums.gen.go]: Processing enum WKDialogResult (3 cases) */
// WKDialogResult - An enumeration that lists the possible ways a delegate handled displaying a custom Lockdown Mode first use dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult
type WKDialogResult uint

const (
	// WKDialogResultAskAgain - A result that indicates the delegate didn’t display a message, so other web views should check again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult/askAgain
	WKDialogResultAskAgain WKDialogResult = 0
	// WKDialogResultHandled - A result that indicates the delegate displayed the first use message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult/handled
	WKDialogResultHandled WKDialogResult = 0
	// WKDialogResultShowDefault - A result that indicates the delegate didn’t display a message, so the web view should show the default Lockdown Mode message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDialogResult/showDefault
	WKDialogResultShowDefault WKDialogResult = 0
)

/* debug [enums.gen.go]: Processing enum WKErrorCode (17 cases) */
// WKErrorCode - Possible error values that WebKit APIs can return.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code
type WKErrorCode uint

const (
	// WKErrorAttributedStringContentFailedToLoad - An error that indicates the failure to navigate to web content from an attributed string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/attributedStringContentFailedToLoad
	WKErrorAttributedStringContentFailedToLoad WKErrorCode = 0
	// WKErrorAttributedStringContentLoadTimedOut - An error that indicates a timeout occurred while trying to load web content from an attributed string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/attributedStringContentLoadTimedOut
	WKErrorAttributedStringContentLoadTimedOut WKErrorCode = 0
	// WKErrorContentRuleListStoreCompileFailed - An error that indicates the compilation of a rule list failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreCompileFailed
	WKErrorContentRuleListStoreCompileFailed WKErrorCode = 0
	// WKErrorContentRuleListStoreLookUpFailed - An error that indicates a content rule list data store didn’t find a rule list with the specified identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreLookUpFailed
	WKErrorContentRuleListStoreLookUpFailed WKErrorCode = 0
	// WKErrorContentRuleListStoreRemoveFailed - An error that indicates a failure to remove a content rule list from the rule list data store object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreRemoveFailed
	WKErrorContentRuleListStoreRemoveFailed WKErrorCode = 0
	// WKErrorContentRuleListStoreVersionMismatch - An error that indicates the rule list version is outdated and cannot be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/contentRuleListStoreVersionMismatch
	WKErrorContentRuleListStoreVersionMismatch WKErrorCode = 0
	// WKErrorCredentialNotFound - An error that indicates the system could not find a passkey during an export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/credentialNotFound
	WKErrorCredentialNotFound WKErrorCode = 0
	// WKErrorDuplicateCredential - An error that indicates the system found a duplicate passkey during an import.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/duplicateCredential
	WKErrorDuplicateCredential WKErrorCode = 0
	// WKErrorJavaScriptAppBoundDomain - An error that indicates JavaScript execution failed due to an app-bound domain restriction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/javaScriptAppBoundDomain
	WKErrorJavaScriptAppBoundDomain WKErrorCode = 0
	// WKErrorJavaScriptExceptionOccurred - An error that indicates a JavaScript exception occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/javaScriptExceptionOccurred
	WKErrorJavaScriptExceptionOccurred WKErrorCode = 0
	// WKErrorJavaScriptInvalidFrameTarget - An error that indicates your content referenced an invalid web frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/javaScriptInvalidFrameTarget
	WKErrorJavaScriptInvalidFrameTarget WKErrorCode = 0
	// WKErrorJavaScriptResultTypeIsUnsupported - An error that indicates the result of JavaScript execution could not be returned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/javaScriptResultTypeIsUnsupported
	WKErrorJavaScriptResultTypeIsUnsupported WKErrorCode = 0
	// WKErrorMalformedCredential - An error that indicates the system could not parse passkey data during an import.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/malformedCredential
	WKErrorMalformedCredential WKErrorCode = 0
	// WKErrorNavigationAppBoundDomain - An error that indicates navigation failed due to an app-bound domain restriction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/navigationAppBoundDomain
	WKErrorNavigationAppBoundDomain WKErrorCode = 0
	// WKErrorUnknown - An error that indicates an unknown issue occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/unknown
	WKErrorUnknown WKErrorCode = 0
	// WKErrorWebContentProcessTerminated - An error that indicates the web process that contains the content is no longer running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/webContentProcessTerminated
	WKErrorWebContentProcessTerminated WKErrorCode = 0
	// WKErrorWebViewInvalidated - An error that indicates the web view was invalidated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKError/Code/webViewInvalidated
	WKErrorWebViewInvalidated WKErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum WKCookiePolicy (2 cases) */
// WKCookiePolicy - An enumeration with cases that indicate whether a cookie store allows cookie storage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/CookiePolicy
type WKCookiePolicy uint

const (
	// WKCookiePolicyAllow - A case that indicates the cookie store allows cookie storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/CookiePolicy/allow
	WKCookiePolicyAllow WKCookiePolicy = 0
	// WKCookiePolicyDisallow - A case that indicates the cookie store does not allow cookie storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/CookiePolicy/disallow
	WKCookiePolicyDisallow WKCookiePolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKMediaCaptureState (3 cases) */
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

/* debug [enums.gen.go]: Processing enum WKMediaCaptureType (3 cases) */
// WKMediaCaptureType - An enumeration listing the types of media devices that can capture audio, video, or both.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureType
type WKMediaCaptureType uint

const (
	// WKMediaCaptureTypeCamera - A media device that can capture video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureType/camera
	WKMediaCaptureTypeCamera WKMediaCaptureType = 0
	// WKMediaCaptureTypeCameraAndMicrophone - A media device or devices that can capture audio and video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureType/cameraAndMicrophone
	WKMediaCaptureTypeCameraAndMicrophone WKMediaCaptureType = 0
	// WKMediaCaptureTypeMicrophone - A media device that can capture audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKMediaCaptureType/microphone
	WKMediaCaptureTypeMicrophone WKMediaCaptureType = 0
)

/* debug [enums.gen.go]: Processing enum WKMediaPlaybackState (4 cases) */
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

/* debug [enums.gen.go]: Processing enum WKNavigationActionPolicy (3 cases) */
// WKNavigationActionPolicy - Constants that indicate whether to allow or cancel navigation to a webpage from an action.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
type WKNavigationActionPolicy uint

const (
	// WKNavigationActionPolicyAllow - Allow the navigation to continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy/allow
	WKNavigationActionPolicyAllow WKNavigationActionPolicy = 0
	// WKNavigationActionPolicyCancel - Cancel the navigation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy/cancel
	WKNavigationActionPolicyCancel WKNavigationActionPolicy = 0
	// WKNavigationActionPolicyDownload - Allow the download to proceed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy/download
	WKNavigationActionPolicyDownload WKNavigationActionPolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKNavigationResponsePolicy (3 cases) */
// WKNavigationResponsePolicy - Constants that indicate whether to allow or cancel navigation to a webpage from a response.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy
type WKNavigationResponsePolicy uint

const (
	// WKNavigationResponsePolicyAllow - Allow the navigation to continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy/allow
	WKNavigationResponsePolicyAllow WKNavigationResponsePolicy = 0
	// WKNavigationResponsePolicyCancel - Cancel the navigation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy/cancel
	WKNavigationResponsePolicyCancel WKNavigationResponsePolicy = 0
	// WKNavigationResponsePolicyDownload - Allow the download to proceed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy/download
	WKNavigationResponsePolicyDownload WKNavigationResponsePolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKNavigationType (6 cases) */
// WKNavigationType - The type of action that triggered the navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType
type WKNavigationType uint

const (
	// WKNavigationTypeBackForward - A request for the frame’s next or previous item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/backForward
	WKNavigationTypeBackForward WKNavigationType = 0
	// WKNavigationTypeFormResubmitted - A request to resubmit a form.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/formResubmitted
	WKNavigationTypeFormResubmitted WKNavigationType = 0
	// WKNavigationTypeFormSubmitted - A request to submit a form.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/formSubmitted
	WKNavigationTypeFormSubmitted WKNavigationType = 0
	// WKNavigationTypeLinkActivated - A link activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/linkActivated
	WKNavigationTypeLinkActivated WKNavigationType = 0
	// WKNavigationTypeOther - A navigation request that originates for some other reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/other
	WKNavigationTypeOther WKNavigationType = 0
	// WKNavigationTypeReload - A request to reload the webpage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationType/reload
	WKNavigationTypeReload WKNavigationType = 0
)

/* debug [enums.gen.go]: Processing enum WKPermissionDecision (3 cases) */
// WKPermissionDecision - An enumeration of possible permission decisions for device resource access.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPermissionDecision
type WKPermissionDecision uint

const (
	// WKPermissionDecisionDeny - Deny permission for the requested resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPermissionDecision/deny
	WKPermissionDecisionDeny WKPermissionDecision = 0
	// WKPermissionDecisionGrant - Grant permission for the requested resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPermissionDecision/grant
	WKPermissionDecisionGrant WKPermissionDecision = 0
	// WKPermissionDecisionPrompt - Prompt the user for permission for the requested resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPermissionDecision/prompt
	WKPermissionDecisionPrompt WKPermissionDecision = 0
)

/* debug [enums.gen.go]: Processing enum WKInactiveSchedulingPolicy (3 cases) */
// WKInactiveSchedulingPolicy - An enumeration that lists policies for how a web view that’s not in a window handles tasks.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum
type WKInactiveSchedulingPolicy uint

const (
	// WKInactiveSchedulingPolicyNone - A policy where a web view that’s not in a window runs tasks normally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum/none
	WKInactiveSchedulingPolicyNone WKInactiveSchedulingPolicy = 0
	// WKInactiveSchedulingPolicySuspend - A policy where a web view that’s not in a window fully suspends tasks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum/suspend
	WKInactiveSchedulingPolicySuspend WKInactiveSchedulingPolicy = 0
	// WKInactiveSchedulingPolicyThrottle - A policy where a web view that’s not in a window limits processing, but does not fully suspend tasks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum/throttle
	WKInactiveSchedulingPolicyThrottle WKInactiveSchedulingPolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKSelectionGranularity (2 cases) */
// WKSelectionGranularity - The granularity with which the user can select and modify web view content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSelectionGranularity
type WKSelectionGranularity uint

const (
	// WKSelectionGranularityCharacter - Granularity that allows the user to place selection endpoints at any character boundary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSelectionGranularity/character
	WKSelectionGranularityCharacter WKSelectionGranularity = 0
	// WKSelectionGranularityDynamic - Granularity that varies automatically depending on the selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSelectionGranularity/dynamic
	WKSelectionGranularityDynamic WKSelectionGranularity = 0
)

/* debug [enums.gen.go]: Processing enum WKUserInterfaceDirectionPolicy (2 cases) */
// WKUserInterfaceDirectionPolicy - The policy that determines the directionality of user interface elements in a web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy
type WKUserInterfaceDirectionPolicy uint

const (
	// WKUserInterfaceDirectionPolicyContent - The directionality follows the CSS/HTML/XHTML specifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy/content
	WKUserInterfaceDirectionPolicyContent WKUserInterfaceDirectionPolicy = 0
	// WKUserInterfaceDirectionPolicySystem - The directionality follows the view’s user interface layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy/system
	WKUserInterfaceDirectionPolicySystem WKUserInterfaceDirectionPolicy = 0
)

/* debug [enums.gen.go]: Processing enum WKUserScriptInjectionTime (2 cases) */
// WKUserScriptInjectionTime - Constants for the times at which to inject script content into a webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
type WKUserScriptInjectionTime uint

const (
	// WKUserScriptInjectionTimeAtDocumentEnd - A constant to inject the script after the document finishes loading, but before loading any other subresources.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime/atDocumentEnd
	WKUserScriptInjectionTimeAtDocumentEnd WKUserScriptInjectionTime = 0
	// WKUserScriptInjectionTimeAtDocumentStart - A constant to inject the script after the creation of the webpage’s document element, but before loading any other content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime/atDocumentStart
	WKUserScriptInjectionTimeAtDocumentStart WKUserScriptInjectionTime = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionDataRecordError (4 cases) */
// WKWebExtensionDataRecordError - Constants that indicate errors in the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code
type WKWebExtensionDataRecordError uint

const (
	// WKWebExtensionDataRecordErrorLocalStorageFailed - Indicates a failure occurred when either deleting or calculating local storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code/localStorageFailed
	WKWebExtensionDataRecordErrorLocalStorageFailed WKWebExtensionDataRecordError = 0
	// WKWebExtensionDataRecordErrorSessionStorageFailed - Indicates a failure occurred when either deleting or calculating session storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code/sessionStorageFailed
	WKWebExtensionDataRecordErrorSessionStorageFailed WKWebExtensionDataRecordError = 0
	// WKWebExtensionDataRecordErrorSynchronizedStorageFailed - Indicates a failure occurred when either deleting or calculating synchronized storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code/synchronizedStorageFailed
	WKWebExtensionDataRecordErrorSynchronizedStorageFailed WKWebExtensionDataRecordError = 0
	// WKWebExtensionDataRecordErrorUnknown - Indicates that an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code/unknown
	WKWebExtensionDataRecordErrorUnknown WKWebExtensionDataRecordError = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionError (9 cases) */
// WKWebExtensionError - Constants that indicate errors in the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code
type WKWebExtensionError uint

const (
	// WKWebExtensionErrorInvalidArchive - Indicates that the archive file is invalid or corrupt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidArchive
	WKWebExtensionErrorInvalidArchive WKWebExtensionError = 0
	// WKWebExtensionErrorInvalidBackgroundPersistence - Indicates that the extension specified background persistence that was not compatible with the platform or features requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidBackgroundPersistence
	WKWebExtensionErrorInvalidBackgroundPersistence WKWebExtensionError = 0
	// WKWebExtensionErrorInvalidDeclarativeNetRequestEntry - Indicates that an invalid declarative net request entry was encountered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidDeclarativeNetRequestEntry
	WKWebExtensionErrorInvalidDeclarativeNetRequestEntry WKWebExtensionError = 0
	// WKWebExtensionErrorInvalidManifest - Indicates that an invalid   was encountered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidManifest
	WKWebExtensionErrorInvalidManifest WKWebExtensionError = 0
	// WKWebExtensionErrorInvalidManifestEntry - Indicates that an invalid manifest entry was encountered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidManifestEntry
	WKWebExtensionErrorInvalidManifestEntry WKWebExtensionError = 0
	// WKWebExtensionErrorInvalidResourceCodeSignature - Indicates that a resource failed the bundle’s code signature checks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/invalidResourceCodeSignature
	WKWebExtensionErrorInvalidResourceCodeSignature WKWebExtensionError = 0
	// WKWebExtensionErrorResourceNotFound - Indicates that a specified resource was not found on disk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/resourceNotFound
	WKWebExtensionErrorResourceNotFound WKWebExtensionError = 0
	// WKWebExtensionErrorUnknown - Indicates that an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/unknown
	WKWebExtensionErrorUnknown WKWebExtensionError = 0
	// WKWebExtensionErrorUnsupportedManifestVersion - Indicates that the manifest version is not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code/unsupportedManifestVersion
	WKWebExtensionErrorUnsupportedManifestVersion WKWebExtensionError = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionMatchPatternError (4 cases) */
// WKWebExtensionMatchPatternError - Constants that indicate errors in the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code
type WKWebExtensionMatchPatternError uint

const (
	// WKWebExtensionMatchPatternErrorInvalidHost - Indicates that the host component was invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code/invalidHost
	WKWebExtensionMatchPatternErrorInvalidHost WKWebExtensionMatchPatternError = 0
	// WKWebExtensionMatchPatternErrorInvalidPath - Indicates that the path component was invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code/invalidPath
	WKWebExtensionMatchPatternErrorInvalidPath WKWebExtensionMatchPatternError = 0
	// WKWebExtensionMatchPatternErrorInvalidScheme - Indicates that the scheme component was invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code/invalidScheme
	WKWebExtensionMatchPatternErrorInvalidScheme WKWebExtensionMatchPatternError = 0
	// WKWebExtensionMatchPatternErrorUnknown - Indicates that an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code/unknown
	WKWebExtensionMatchPatternErrorUnknown WKWebExtensionMatchPatternError = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionMatchPatternOptions (4 cases) */
// WKWebExtensionMatchPatternOptions - Constants used by 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Options
type WKWebExtensionMatchPatternOptions uint

const (
	// WKWebExtensionMatchPatternOptionsIgnorePaths - Indicates that the host components should be ignored while matching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Options/ignorePaths
	WKWebExtensionMatchPatternOptionsIgnorePaths WKWebExtensionMatchPatternOptions = 0
	// WKWebExtensionMatchPatternOptionsIgnoreSchemes - Indicates that the scheme components should be ignored while matching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Options/ignoreSchemes
	WKWebExtensionMatchPatternOptionsIgnoreSchemes WKWebExtensionMatchPatternOptions = 0
	// WKWebExtensionMatchPatternOptionsMatchBidirectionally - Indicates that two patterns should be checked in either direction while matching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Options/matchBidirectionally
	WKWebExtensionMatchPatternOptionsMatchBidirectionally WKWebExtensionMatchPatternOptions = 0
	// WKWebExtensionMatchPatternOptionsNone - Indicates no special matching options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionMatchPatternOptions/WKWebExtensionMatchPatternOptionsNone
	WKWebExtensionMatchPatternOptionsNone WKWebExtensionMatchPatternOptions = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionMessagePortError (3 cases) */
// WKWebExtensionMessagePortError - Constants that indicate errors in the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/Error/Code
type WKWebExtensionMessagePortError uint

const (
	// WKWebExtensionMessagePortErrorMessageInvalid - Indicates that the message is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/Error/Code/messageInvalid
	WKWebExtensionMessagePortErrorMessageInvalid WKWebExtensionMessagePortError = 0
	// WKWebExtensionMessagePortErrorNotConnected - Indicates that the message port is disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/Error/Code/notConnected
	WKWebExtensionMessagePortErrorNotConnected WKWebExtensionMessagePortError = 0
	// WKWebExtensionMessagePortErrorUnknown - Indicates that an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/Error/Code/unknown
	WKWebExtensionMessagePortErrorUnknown WKWebExtensionMessagePortError = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionTabChangedProperties (10 cases) */
// WKWebExtensionTabChangedProperties - Constants the web extension controller and web extension context use to indicate tab changes.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties
type WKWebExtensionTabChangedProperties uint

const (
	// WKWebExtensionTabChangedPropertiesLoading - Indicates the loading state changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/loading
	WKWebExtensionTabChangedPropertiesLoading WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesMuted - Indicates the muted state changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/muted
	WKWebExtensionTabChangedPropertiesMuted WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesPinned - Indicates the pinned state changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/pinned
	WKWebExtensionTabChangedPropertiesPinned WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesPlayingAudio - Indicates the audio playback state changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/playingAudio
	WKWebExtensionTabChangedPropertiesPlayingAudio WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesReaderMode - Indicates the reader mode state changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/readerMode
	WKWebExtensionTabChangedPropertiesReaderMode WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesSize - Indicates the size changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/size
	WKWebExtensionTabChangedPropertiesSize WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesTitle - Indicates the title changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/title
	WKWebExtensionTabChangedPropertiesTitle WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesURL - Indicates the URL changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/URL
	WKWebExtensionTabChangedPropertiesURL WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesZoomFactor - Indicates the zoom factor changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties/zoomFactor
	WKWebExtensionTabChangedPropertiesZoomFactor WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesNone - Indicates nothing changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionTabChangedProperties/WKWebExtensionTabChangedPropertiesNone
	WKWebExtensionTabChangedPropertiesNone WKWebExtensionTabChangedProperties = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionWindowState (4 cases) */
// WKWebExtensionWindowState - Constants used by 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState
type WKWebExtensionWindowState uint

const (
	// WKWebExtensionWindowStateFullscreen - Indicates a window is in full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState/fullscreen
	WKWebExtensionWindowStateFullscreen WKWebExtensionWindowState = 0
	// WKWebExtensionWindowStateMaximized - Indicates a window is maximized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState/maximized
	WKWebExtensionWindowStateMaximized WKWebExtensionWindowState = 0
	// WKWebExtensionWindowStateMinimized - Indicates a window is minimized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState/minimized
	WKWebExtensionWindowStateMinimized WKWebExtensionWindowState = 0
	// WKWebExtensionWindowStateNormal - Indicates a window is in its normal state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState/normal
	WKWebExtensionWindowStateNormal WKWebExtensionWindowState = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionWindowType (2 cases) */
// WKWebExtensionWindowType - Constants used by 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowType
type WKWebExtensionWindowType uint

const (
	// WKWebExtensionWindowTypeNormal - Indicates a normal window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowType/normal
	WKWebExtensionWindowTypeNormal WKWebExtensionWindowType = 0
	// WKWebExtensionWindowTypePopup - Indicates a pop-up window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowType/popup
	WKWebExtensionWindowTypePopup WKWebExtensionWindowType = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionContextError (6 cases) */
// WKWebExtensionContextError - Constants that indicate errors in the 
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code
type WKWebExtensionContextError uint

const (
	// WKWebExtensionContextErrorAlreadyLoaded - Indicates that the context is already loaded by a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/alreadyLoaded
	WKWebExtensionContextErrorAlreadyLoaded WKWebExtensionContextError = 0
	// WKWebExtensionContextErrorBackgroundContentFailedToLoad - Indicates that an error occurred loading the background content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/backgroundContentFailedToLoad
	WKWebExtensionContextErrorBackgroundContentFailedToLoad WKWebExtensionContextError = 0
	// WKWebExtensionContextErrorBaseURLAlreadyInUse - Indicates that another context is already using the specified base URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/baseURLAlreadyInUse
	WKWebExtensionContextErrorBaseURLAlreadyInUse WKWebExtensionContextError = 0
	// WKWebExtensionContextErrorNoBackgroundContent - Indicates that the extension does not have background content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/noBackgroundContent
	WKWebExtensionContextErrorNoBackgroundContent WKWebExtensionContextError = 0
	// WKWebExtensionContextErrorNotLoaded - Indicates that the context is not loaded by a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/notLoaded
	WKWebExtensionContextErrorNotLoaded WKWebExtensionContextError = 0
	// WKWebExtensionContextErrorUnknown - Indicates that an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code/unknown
	WKWebExtensionContextErrorUnknown WKWebExtensionContextError = 0
)

/* debug [enums.gen.go]: Processing enum WKWebExtensionContextPermissionStatus (7 cases) */
// WKWebExtensionContextPermissionStatus - Constants used to indicate permission status in web extension context.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus
type WKWebExtensionContextPermissionStatus uint

const (
	// WKWebExtensionContextPermissionStatusDeniedExplicitly - Indicates that the permission was explicitly denied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/deniedExplicitly
	WKWebExtensionContextPermissionStatusDeniedExplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusDeniedImplicitly - Indicates that the permission was implicitly denied because of another explicitly denied permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/deniedImplicitly
	WKWebExtensionContextPermissionStatusDeniedImplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusGrantedExplicitly - Indicates that the permission was explicitly granted permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/grantedExplicitly
	WKWebExtensionContextPermissionStatusGrantedExplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusGrantedImplicitly - Indicates that the permission was implicitly granted because of another explicitly granted permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/grantedImplicitly
	WKWebExtensionContextPermissionStatusGrantedImplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusRequestedExplicitly - Indicates that the permission was explicitly requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/requestedExplicitly
	WKWebExtensionContextPermissionStatusRequestedExplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusRequestedImplicitly - Indicates that the permission was implicitly requested because of another explicitly requested permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/requestedImplicitly
	WKWebExtensionContextPermissionStatusRequestedImplicitly WKWebExtensionContextPermissionStatus = 0
	// WKWebExtensionContextPermissionStatusUnknown - Indicates an unknown permission status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus/unknown
	WKWebExtensionContextPermissionStatusUnknown WKWebExtensionContextPermissionStatus = 0
)

/* debug [enums.gen.go]: Processing enum WKFullscreenState (4 cases) */
// WKFullscreenState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum
type WKFullscreenState uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/enteringFullscreen
	WKFullscreenStateEnteringFullscreen WKFullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/exitingFullscreen
	WKFullscreenStateExitingFullscreen WKFullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/inFullscreen
	WKFullscreenStateInFullscreen WKFullscreenState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum/notInFullscreen
	WKFullscreenStateNotInFullscreen WKFullscreenState = 0
)

/* debug [enums.gen.go]: Processing enum WKWebViewDataType (1 cases) */
// WKWebViewDataType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType
type WKWebViewDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewDataType/sessionStorage
	WKWebViewDataTypeSessionStorage WKWebViewDataType = 0
)


