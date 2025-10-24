// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PWebExtensionTab is the WKWebExtensionTab protocol interface.
//
// A protocol with methods that represent a tab to web extensions.
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - visionOS 2.4+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKWebExtensionTab
type PWebExtensionTab interface {
	// Optional methods
	ActivateForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasActivateForWebExtensionContextCompletionHandler() bool
	CloseForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasCloseForWebExtensionContextCompletionHandler() bool
	DetectWebpageLocaleForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasDetectWebpageLocaleForWebExtensionContextCompletionHandler() bool
	DuplicateUsingConfigurationForWebExtensionContextCompletionHandler(configuration IWKWebExtensionTabConfiguration, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasDuplicateUsingConfigurationForWebExtensionContextCompletionHandler() bool
	GoBackForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasGoBackForWebExtensionContextCompletionHandler() bool
	GoForwardForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasGoForwardForWebExtensionContextCompletionHandler() bool
	IndexInWindowForWebExtensionContext(context IWKWebExtensionContext) uint
	HasIndexInWindowForWebExtensionContext() bool
	IsLoadingCompleteForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsLoadingCompleteForWebExtensionContext() bool
	IsMutedForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsMutedForWebExtensionContext() bool
	IsPinnedForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsPinnedForWebExtensionContext() bool
	IsPlayingAudioForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsPlayingAudioForWebExtensionContext() bool
	IsReaderModeActiveForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsReaderModeActiveForWebExtensionContext() bool
	IsReaderModeAvailableForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsReaderModeAvailableForWebExtensionContext() bool
	IsSelectedForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsSelectedForWebExtensionContext() bool
	LoadURLForWebExtensionContextCompletionHandler(url objc.IObject /* cross-framework: NSURL */, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasLoadURLForWebExtensionContextCompletionHandler() bool
	ParentTabForWebExtensionContext(context IWKWebExtensionContext) unsafe.Pointer
	HasParentTabForWebExtensionContext() bool
	PendingURLForWebExtensionContext(context IWKWebExtensionContext) foundation.URL
	HasPendingURLForWebExtensionContext() bool
	ReloadFromOriginForWebExtensionContextCompletionHandler(fromOrigin bool, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasReloadFromOriginForWebExtensionContextCompletionHandler() bool
	SetMutedForWebExtensionContextCompletionHandler(muted bool, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetMutedForWebExtensionContextCompletionHandler() bool
	SetParentTabForWebExtensionContextCompletionHandler(parentTab unsafe.Pointer, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetParentTabForWebExtensionContextCompletionHandler() bool
	SetPinnedForWebExtensionContextCompletionHandler(pinned bool, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetPinnedForWebExtensionContextCompletionHandler() bool
	SetReaderModeActiveForWebExtensionContextCompletionHandler(active bool, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetReaderModeActiveForWebExtensionContextCompletionHandler() bool
	SetSelectedForWebExtensionContextCompletionHandler(selected bool, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetSelectedForWebExtensionContextCompletionHandler() bool
	SetZoomFactorForWebExtensionContextCompletionHandler(zoomFactor float64, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetZoomFactorForWebExtensionContextCompletionHandler() bool
	ShouldBypassPermissionsForWebExtensionContext(context IWKWebExtensionContext) bool
	HasShouldBypassPermissionsForWebExtensionContext() bool
	ShouldGrantPermissionsOnUserGestureForWebExtensionContext(context IWKWebExtensionContext) bool
	HasShouldGrantPermissionsOnUserGestureForWebExtensionContext() bool
	SizeForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGSize
	HasSizeForWebExtensionContext() bool
	TakeSnapshotUsingConfigurationForWebExtensionContextCompletionHandler(configuration IWKSnapshotConfiguration, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasTakeSnapshotUsingConfigurationForWebExtensionContextCompletionHandler() bool
	TitleForWebExtensionContext(context IWKWebExtensionContext) foundation.String
	HasTitleForWebExtensionContext() bool
	UrlForWebExtensionContext(context IWKWebExtensionContext) foundation.URL
	HasUrlForWebExtensionContext() bool
	WebViewForWebExtensionContext(context IWKWebExtensionContext) WebView
	HasWebViewForWebExtensionContext() bool
	WindowForWebExtensionContext(context IWKWebExtensionContext) unsafe.Pointer
	HasWindowForWebExtensionContext() bool
	ZoomFactorForWebExtensionContext(context IWKWebExtensionContext) float64
	HasZoomFactorForWebExtensionContext() bool
}
