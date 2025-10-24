// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/avfoundation"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/corevideo"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayerViewControllerDelegate is the AVPlayerViewControllerDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to player view controller events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPlayerViewControllerDelegate
type PPlayerViewControllerDelegate interface {
	// Required methods
	PlayerViewControllerDidSelectExternalSubtitleOptionLanguage(playerViewController IAVPlayerViewController, language objc.IObject /* cross-framework: NSString */)
	// Optional methods
	NextChannelInterstitialViewControllerForPlayerViewController(playerViewController IAVPlayerViewController) ViewController
	HasNextChannelInterstitialViewControllerForPlayerViewController() bool
	PlayerViewControllerDidAcceptContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal)
	HasPlayerViewControllerDidAcceptContentProposal() bool
	PlayerViewControllerDidPresentInterstitialTimeRange(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)
	HasPlayerViewControllerDidPresentInterstitialTimeRange() bool
	PlayerViewControllerDidRejectContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal)
	HasPlayerViewControllerDidRejectContentProposal() bool
	PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup(playerViewController IAVPlayerViewController, mediaSelectionOption avfoundation.MediaSelectionOption, mediaSelectionGroup avfoundation.MediaSelectionGroup)
	HasPlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup() bool
	PlayerViewControllerFailedToStartPictureInPictureWithError(playerViewController IAVPlayerViewController, error_ objc.IObject /* cross-framework: Error */)
	HasPlayerViewControllerFailedToStartPictureInPictureWithError() bool
	PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)
	HasPlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool
	PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)
	HasPlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool
	PlayerViewControllerShouldPresentContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal) bool
	HasPlayerViewControllerShouldPresentContentProposal() bool
	PlayerViewControllerSkipToNextChannel(playerViewController IAVPlayerViewController, completion unsafe.Pointer)
	HasPlayerViewControllerSkipToNextChannel() bool
	PlayerViewControllerSkipToPreviousChannel(playerViewController IAVPlayerViewController, completion unsafe.Pointer)
	HasPlayerViewControllerSkipToPreviousChannel() bool
	PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */) corevideo.Time
	HasPlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime() bool
	PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)
	HasPlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator() bool
	PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)
	HasPlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator() bool
	PlayerViewControllerWillPresentInterstitialTimeRange(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)
	HasPlayerViewControllerWillPresentInterstitialTimeRange() bool
	PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */)
	HasPlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime() bool
	PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator(playerViewController IAVPlayerViewController, visible bool, coordinator unsafe.Pointer)
	HasPlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator() bool
	PlayerViewControllerDidEndDismissalTransition(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerDidEndDismissalTransition() bool
	PlayerViewControllerDidStartPictureInPicture(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerDidStartPictureInPicture() bool
	PlayerViewControllerDidStopPictureInPicture(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerDidStopPictureInPicture() bool
	PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart(playerViewController IAVPlayerViewController) bool
	HasPlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart() bool
	PlayerViewControllerShouldDismiss(playerViewController IAVPlayerViewController) bool
	HasPlayerViewControllerShouldDismiss() bool
	PlayerViewControllerWillBeginDismissalTransition(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerWillBeginDismissalTransition() bool
	PlayerViewControllerWillStartPictureInPicture(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerWillStartPictureInPicture() bool
	PlayerViewControllerWillStopPictureInPicture(playerViewController IAVPlayerViewController)
	HasPlayerViewControllerWillStopPictureInPicture() bool
	PreviousChannelInterstitialViewControllerForPlayerViewController(playerViewController IAVPlayerViewController) ViewController
	HasPreviousChannelInterstitialViewControllerForPlayerViewController() bool
	SkipToNextItemForPlayerViewController(playerViewController IAVPlayerViewController)
	HasSkipToNextItemForPlayerViewController() bool
	SkipToPreviousItemForPlayerViewController(playerViewController IAVPlayerViewController)
	HasSkipToPreviousItemForPlayerViewController() bool
}

// PlayerViewControllerDelegate is a delegate implementation builder for the PPlayerViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerViewControllerDelegate struct {
	_NextChannelInterstitialViewControllerForPlayerViewController func(playerViewController IAVPlayerViewController) ViewController
	_PlayerViewControllerDidAcceptContentProposal func(playerViewController IAVPlayerViewController, proposal IAVContentProposal)
	_PlayerViewControllerDidPresentInterstitialTimeRange func(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)
	_PlayerViewControllerDidRejectContentProposal func(playerViewController IAVPlayerViewController, proposal IAVContentProposal)
	_PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup func(playerViewController IAVPlayerViewController, mediaSelectionOption avfoundation.MediaSelectionOption, mediaSelectionGroup avfoundation.MediaSelectionGroup)
	_PlayerViewControllerFailedToStartPictureInPictureWithError func(playerViewController IAVPlayerViewController, error_ objc.IObject /* cross-framework: Error */)
	_PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler func(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)
	_PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler func(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)
	_PlayerViewControllerShouldPresentContentProposal func(playerViewController IAVPlayerViewController, proposal IAVContentProposal) bool
	_PlayerViewControllerSkipToNextChannel func(playerViewController IAVPlayerViewController, completion unsafe.Pointer)
	_PlayerViewControllerSkipToPreviousChannel func(playerViewController IAVPlayerViewController, completion unsafe.Pointer)
	_PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime func(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */) corevideo.Time
	_PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator func(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)
	_PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator func(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)
	_PlayerViewControllerWillPresentInterstitialTimeRange func(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)
	_PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime func(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */)
	_PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator func(playerViewController IAVPlayerViewController, visible bool, coordinator unsafe.Pointer)
	_PlayerViewControllerDidEndDismissalTransition func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerDidStartPictureInPicture func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerDidStopPictureInPicture func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart func(playerViewController IAVPlayerViewController) bool
	_PlayerViewControllerShouldDismiss func(playerViewController IAVPlayerViewController) bool
	_PlayerViewControllerWillBeginDismissalTransition func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerWillStartPictureInPicture func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerWillStopPictureInPicture func(playerViewController IAVPlayerViewController)
	_PreviousChannelInterstitialViewControllerForPlayerViewController func(playerViewController IAVPlayerViewController) ViewController
	_SkipToNextItemForPlayerViewController func(playerViewController IAVPlayerViewController)
	_SkipToPreviousItemForPlayerViewController func(playerViewController IAVPlayerViewController)
	_PlayerViewControllerDidSelectExternalSubtitleOptionLanguage func(playerViewController IAVPlayerViewController, language objc.IObject /* cross-framework: NSString */)
}

// SetNextChannelInterstitialViewControllerForPlayerViewController sets the handler for the NextChannelInterstitialViewControllerForPlayerViewController delegate method.
//
// Asks the delegate for a view controller that describes the layout of the next channel’s interstitial view.
func (d *PlayerViewControllerDelegate) SetNextChannelInterstitialViewControllerForPlayerViewController(f func(playerViewController IAVPlayerViewController) ViewController) {
	d._NextChannelInterstitialViewControllerForPlayerViewController = f
}

// SetPlayerViewControllerDidAcceptContentProposal sets the handler for the PlayerViewControllerDidAcceptContentProposal delegate method.
//
// Tells the delegate when the user accepts the proposed content.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidAcceptContentProposal(f func(playerViewController IAVPlayerViewController, proposal IAVContentProposal)) {
	d._PlayerViewControllerDidAcceptContentProposal = f
}

// SetPlayerViewControllerDidPresentInterstitialTimeRange sets the handler for the PlayerViewControllerDidPresentInterstitialTimeRange delegate method.
//
// Tells the delegate when the player view controller finishes playing a range of interstitial content.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidPresentInterstitialTimeRange(f func(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)) {
	d._PlayerViewControllerDidPresentInterstitialTimeRange = f
}

// SetPlayerViewControllerDidRejectContentProposal sets the handler for the PlayerViewControllerDidRejectContentProposal delegate method.
//
// Tells the delegate when the user rejects the proposed content.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidRejectContentProposal(f func(playerViewController IAVPlayerViewController, proposal IAVContentProposal)) {
	d._PlayerViewControllerDidRejectContentProposal = f
}

// SetPlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup sets the handler for the PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup delegate method.
//
// Tells the delegate when the user selects a media option from a media selection group.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup(f func(playerViewController IAVPlayerViewController, mediaSelectionOption avfoundation.MediaSelectionOption, mediaSelectionGroup avfoundation.MediaSelectionGroup)) {
	d._PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup = f
}

// SetPlayerViewControllerFailedToStartPictureInPictureWithError sets the handler for the PlayerViewControllerFailedToStartPictureInPictureWithError delegate method.
//
// Tells the delegate when Picture in Picture fails to start.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerFailedToStartPictureInPictureWithError(f func(playerViewController IAVPlayerViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._PlayerViewControllerFailedToStartPictureInPictureWithError = f
}

// SetPlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler sets the handler for the PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler delegate method.
//
// Tells the delegate to restore the app’s user interface after returning from a full-screen presentation.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler(f func(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)) {
	d._PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler = f
}

// SetPlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler sets the handler for the PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler delegate method.
//
// Tells the delegate when Picture in Picture is about to stop so you can restore your app’s user interface.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(f func(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer)) {
	d._PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler = f
}

// SetPlayerViewControllerShouldPresentContentProposal sets the handler for the PlayerViewControllerShouldPresentContentProposal delegate method.
//
// Asks the delegate whether the player view controller presents a content proposal.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerShouldPresentContentProposal(f func(playerViewController IAVPlayerViewController, proposal IAVContentProposal) bool) {
	d._PlayerViewControllerShouldPresentContentProposal = f
}

// SetPlayerViewControllerSkipToNextChannel sets the handler for the PlayerViewControllerSkipToNextChannel delegate method.
//
// Tells the delegate when the user wants to skip to the next channel.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerSkipToNextChannel(f func(playerViewController IAVPlayerViewController, completion unsafe.Pointer)) {
	d._PlayerViewControllerSkipToNextChannel = f
}

// SetPlayerViewControllerSkipToPreviousChannel sets the handler for the PlayerViewControllerSkipToPreviousChannel delegate method.
//
// Tells the delegate when the user wants to skip to the previous channel.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerSkipToPreviousChannel(f func(playerViewController IAVPlayerViewController, completion unsafe.Pointer)) {
	d._PlayerViewControllerSkipToPreviousChannel = f
}

// SetPlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime sets the handler for the PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime delegate method.
//
// Tells the delegate when the user skips, scrubs, or otherwise navigates to a new time and wants to resume playback at the target time.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime(f func(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */) corevideo.Time) {
	d._PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime = f
}

// SetPlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator sets the handler for the PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator delegate method.
//
// Tells the delegate when the player view controller is about to start full-screen display.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator(f func(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)) {
	d._PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator = f
}

// SetPlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator sets the handler for the PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator delegate method.
//
// Tells the delegate when the player view controller is about to end full-screen display.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator(f func(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer)) {
	d._PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator = f
}

// SetPlayerViewControllerWillPresentInterstitialTimeRange sets the handler for the PlayerViewControllerWillPresentInterstitialTimeRange delegate method.
//
// Tells the delegate when the player view controller is about to start playing a range of interstitial content.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillPresentInterstitialTimeRange(f func(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange)) {
	d._PlayerViewControllerWillPresentInterstitialTimeRange = f
}

// SetPlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime sets the handler for the PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime delegate method.
//
// Tells the delegate when the user navigates to a new time and playback is about to begin.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime(f func(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */)) {
	d._PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime = f
}

// SetPlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator sets the handler for the PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator delegate method.
//
// Tells the delegate when the transport bar’s visibility is about to change.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator(f func(playerViewController IAVPlayerViewController, visible bool, coordinator unsafe.Pointer)) {
	d._PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator = f
}

// SetPlayerViewControllerDidEndDismissalTransition sets the handler for the PlayerViewControllerDidEndDismissalTransition delegate method.
//
// Tells the delegate when the player view controller ends its dismissal transition.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidEndDismissalTransition(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerDidEndDismissalTransition = f
}

// SetPlayerViewControllerDidStartPictureInPicture sets the handler for the PlayerViewControllerDidStartPictureInPicture delegate method.
//
// Tells the delegate when Picture in Picture starts.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidStartPictureInPicture(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerDidStartPictureInPicture = f
}

// SetPlayerViewControllerDidStopPictureInPicture sets the handler for the PlayerViewControllerDidStopPictureInPicture delegate method.
//
// Tells the delegate when Picture in Picture stops.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidStopPictureInPicture(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerDidStopPictureInPicture = f
}

// SetPlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart sets the handler for the PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart delegate method.
//
// Asks the delegate whether the player view controller automatically dismisses itself when Picture in Picture starts.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart(f func(playerViewController IAVPlayerViewController) bool) {
	d._PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart = f
}

// SetPlayerViewControllerShouldDismiss sets the handler for the PlayerViewControllerShouldDismiss delegate method.
//
// Asks the delegate object whether the player view controller dismisses itself upon request.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerShouldDismiss(f func(playerViewController IAVPlayerViewController) bool) {
	d._PlayerViewControllerShouldDismiss = f
}

// SetPlayerViewControllerWillBeginDismissalTransition sets the handler for the PlayerViewControllerWillBeginDismissalTransition delegate method.
//
// Tells the delegate when the player view controller is about to start its dismissal transition.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillBeginDismissalTransition(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerWillBeginDismissalTransition = f
}

// SetPlayerViewControllerWillStartPictureInPicture sets the handler for the PlayerViewControllerWillStartPictureInPicture delegate method.
//
// Tells the delegate when Picture in Picture is about to start.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillStartPictureInPicture(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerWillStartPictureInPicture = f
}

// SetPlayerViewControllerWillStopPictureInPicture sets the handler for the PlayerViewControllerWillStopPictureInPicture delegate method.
//
// Tells the delegate when Picture in Picture is about to stop.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerWillStopPictureInPicture(f func(playerViewController IAVPlayerViewController)) {
	d._PlayerViewControllerWillStopPictureInPicture = f
}

// SetPreviousChannelInterstitialViewControllerForPlayerViewController sets the handler for the PreviousChannelInterstitialViewControllerForPlayerViewController delegate method.
//
// Asks the delegate for a view controller that describes the layout of the previous channel’s interstitial view.
func (d *PlayerViewControllerDelegate) SetPreviousChannelInterstitialViewControllerForPlayerViewController(f func(playerViewController IAVPlayerViewController) ViewController) {
	d._PreviousChannelInterstitialViewControllerForPlayerViewController = f
}

// SetSkipToNextItemForPlayerViewController sets the handler for the SkipToNextItemForPlayerViewController delegate method.
//
// Tells the delegate when the user requests skipping to the next item in the timeline.
func (d *PlayerViewControllerDelegate) SetSkipToNextItemForPlayerViewController(f func(playerViewController IAVPlayerViewController)) {
	d._SkipToNextItemForPlayerViewController = f
}

// SetSkipToPreviousItemForPlayerViewController sets the handler for the SkipToPreviousItemForPlayerViewController delegate method.
//
// Tells the delegate when the user requests skipping to the previous item in the timeline.
func (d *PlayerViewControllerDelegate) SetSkipToPreviousItemForPlayerViewController(f func(playerViewController IAVPlayerViewController)) {
	d._SkipToPreviousItemForPlayerViewController = f
}

// SetPlayerViewControllerDidSelectExternalSubtitleOptionLanguage sets the handler for the PlayerViewControllerDidSelectExternalSubtitleOptionLanguage delegate method.
//
// Tells the delegate when the user selects a specific subtitle option.
func (d *PlayerViewControllerDelegate) SetPlayerViewControllerDidSelectExternalSubtitleOptionLanguage(f func(playerViewController IAVPlayerViewController, language objc.IObject /* cross-framework: NSString */)) {
	d._PlayerViewControllerDidSelectExternalSubtitleOptionLanguage = f
}

// NextChannelInterstitialViewControllerForPlayerViewController implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) NextChannelInterstitialViewControllerForPlayerViewController(playerViewController IAVPlayerViewController) ViewController {
	if d._NextChannelInterstitialViewControllerForPlayerViewController != nil {
		return d._NextChannelInterstitialViewControllerForPlayerViewController(playerViewController)
	}
	var zero ViewController
	return zero
}

// HasNextChannelInterstitialViewControllerForPlayerViewController returns true if a handler for NextChannelInterstitialViewControllerForPlayerViewController has been set.
func (d *PlayerViewControllerDelegate) HasNextChannelInterstitialViewControllerForPlayerViewController() bool {
	return d._NextChannelInterstitialViewControllerForPlayerViewController != nil
}

// PlayerViewControllerDidAcceptContentProposal implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidAcceptContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal) {
	if d._PlayerViewControllerDidAcceptContentProposal != nil {
		d._PlayerViewControllerDidAcceptContentProposal(playerViewController, proposal)
	}
}

// HasPlayerViewControllerDidAcceptContentProposal returns true if a handler for PlayerViewControllerDidAcceptContentProposal has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidAcceptContentProposal() bool {
	return d._PlayerViewControllerDidAcceptContentProposal != nil
}

// PlayerViewControllerDidPresentInterstitialTimeRange implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidPresentInterstitialTimeRange(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange) {
	if d._PlayerViewControllerDidPresentInterstitialTimeRange != nil {
		d._PlayerViewControllerDidPresentInterstitialTimeRange(playerViewController, interstitial)
	}
}

// HasPlayerViewControllerDidPresentInterstitialTimeRange returns true if a handler for PlayerViewControllerDidPresentInterstitialTimeRange has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidPresentInterstitialTimeRange() bool {
	return d._PlayerViewControllerDidPresentInterstitialTimeRange != nil
}

// PlayerViewControllerDidRejectContentProposal implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidRejectContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal) {
	if d._PlayerViewControllerDidRejectContentProposal != nil {
		d._PlayerViewControllerDidRejectContentProposal(playerViewController, proposal)
	}
}

// HasPlayerViewControllerDidRejectContentProposal returns true if a handler for PlayerViewControllerDidRejectContentProposal has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidRejectContentProposal() bool {
	return d._PlayerViewControllerDidRejectContentProposal != nil
}

// PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup(playerViewController IAVPlayerViewController, mediaSelectionOption avfoundation.MediaSelectionOption, mediaSelectionGroup avfoundation.MediaSelectionGroup) {
	if d._PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup != nil {
		d._PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup(playerViewController, mediaSelectionOption, mediaSelectionGroup)
	}
}

// HasPlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup returns true if a handler for PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup() bool {
	return d._PlayerViewControllerDidSelectMediaSelectionOptionInMediaSelectionGroup != nil
}

// PlayerViewControllerFailedToStartPictureInPictureWithError implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerFailedToStartPictureInPictureWithError(playerViewController IAVPlayerViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._PlayerViewControllerFailedToStartPictureInPictureWithError != nil {
		d._PlayerViewControllerFailedToStartPictureInPictureWithError(playerViewController, error_)
	}
}

// HasPlayerViewControllerFailedToStartPictureInPictureWithError returns true if a handler for PlayerViewControllerFailedToStartPictureInPictureWithError has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerFailedToStartPictureInPictureWithError() bool {
	return d._PlayerViewControllerFailedToStartPictureInPictureWithError != nil
}

// PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer) {
	if d._PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler != nil {
		d._PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerViewController, completionHandler)
	}
}

// HasPlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler returns true if a handler for PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool {
	return d._PlayerViewControllerRestoreUserInterfaceForFullScreenExitWithCompletionHandler != nil
}

// PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerViewController IAVPlayerViewController, completionHandler unsafe.Pointer) {
	if d._PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil {
		d._PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerViewController, completionHandler)
	}
}

// HasPlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler returns true if a handler for PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool {
	return d._PlayerViewControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil
}

// PlayerViewControllerShouldPresentContentProposal implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerShouldPresentContentProposal(playerViewController IAVPlayerViewController, proposal IAVContentProposal) bool {
	if d._PlayerViewControllerShouldPresentContentProposal != nil {
		return d._PlayerViewControllerShouldPresentContentProposal(playerViewController, proposal)
	}
	var zero bool
	return zero
}

// HasPlayerViewControllerShouldPresentContentProposal returns true if a handler for PlayerViewControllerShouldPresentContentProposal has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerShouldPresentContentProposal() bool {
	return d._PlayerViewControllerShouldPresentContentProposal != nil
}

// PlayerViewControllerSkipToNextChannel implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerSkipToNextChannel(playerViewController IAVPlayerViewController, completion unsafe.Pointer) {
	if d._PlayerViewControllerSkipToNextChannel != nil {
		d._PlayerViewControllerSkipToNextChannel(playerViewController, completion)
	}
}

// HasPlayerViewControllerSkipToNextChannel returns true if a handler for PlayerViewControllerSkipToNextChannel has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerSkipToNextChannel() bool {
	return d._PlayerViewControllerSkipToNextChannel != nil
}

// PlayerViewControllerSkipToPreviousChannel implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerSkipToPreviousChannel(playerViewController IAVPlayerViewController, completion unsafe.Pointer) {
	if d._PlayerViewControllerSkipToPreviousChannel != nil {
		d._PlayerViewControllerSkipToPreviousChannel(playerViewController, completion)
	}
}

// HasPlayerViewControllerSkipToPreviousChannel returns true if a handler for PlayerViewControllerSkipToPreviousChannel has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerSkipToPreviousChannel() bool {
	return d._PlayerViewControllerSkipToPreviousChannel != nil
}

// PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */) corevideo.Time {
	if d._PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime != nil {
		return d._PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime(playerViewController, oldTime, targetTime)
	}
	var zero corevideo.Time
	return zero
}

// HasPlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime returns true if a handler for PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime() bool {
	return d._PlayerViewControllerTimeToSeekAfterUserNavigatedFromTimeToTime != nil
}

// PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer) {
	if d._PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator != nil {
		d._PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator(playerViewController, coordinator)
	}
}

// HasPlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator returns true if a handler for PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator() bool {
	return d._PlayerViewControllerWillBeginFullScreenPresentationWithAnimationCoordinator != nil
}

// PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator(playerViewController IAVPlayerViewController, coordinator unsafe.Pointer) {
	if d._PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator != nil {
		d._PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator(playerViewController, coordinator)
	}
}

// HasPlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator returns true if a handler for PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator() bool {
	return d._PlayerViewControllerWillEndFullScreenPresentationWithAnimationCoordinator != nil
}

// PlayerViewControllerWillPresentInterstitialTimeRange implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillPresentInterstitialTimeRange(playerViewController IAVPlayerViewController, interstitial IAVInterstitialTimeRange) {
	if d._PlayerViewControllerWillPresentInterstitialTimeRange != nil {
		d._PlayerViewControllerWillPresentInterstitialTimeRange(playerViewController, interstitial)
	}
}

// HasPlayerViewControllerWillPresentInterstitialTimeRange returns true if a handler for PlayerViewControllerWillPresentInterstitialTimeRange has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillPresentInterstitialTimeRange() bool {
	return d._PlayerViewControllerWillPresentInterstitialTimeRange != nil
}

// PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime(playerViewController IAVPlayerViewController, oldTime objc.IObject /* cross-framework: Time */, targetTime objc.IObject /* cross-framework: Time */) {
	if d._PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime != nil {
		d._PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime(playerViewController, oldTime, targetTime)
	}
}

// HasPlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime returns true if a handler for PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime() bool {
	return d._PlayerViewControllerWillResumePlaybackAfterUserNavigatedFromTimeToTime != nil
}

// PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator(playerViewController IAVPlayerViewController, visible bool, coordinator unsafe.Pointer) {
	if d._PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator != nil {
		d._PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator(playerViewController, visible, coordinator)
	}
}

// HasPlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator returns true if a handler for PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator() bool {
	return d._PlayerViewControllerWillTransitionToVisibilityOfTransportBarWithAnimationCoordinator != nil
}

// PlayerViewControllerDidEndDismissalTransition implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidEndDismissalTransition(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerDidEndDismissalTransition != nil {
		d._PlayerViewControllerDidEndDismissalTransition(playerViewController)
	}
}

// HasPlayerViewControllerDidEndDismissalTransition returns true if a handler for PlayerViewControllerDidEndDismissalTransition has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidEndDismissalTransition() bool {
	return d._PlayerViewControllerDidEndDismissalTransition != nil
}

// PlayerViewControllerDidStartPictureInPicture implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidStartPictureInPicture(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerDidStartPictureInPicture != nil {
		d._PlayerViewControllerDidStartPictureInPicture(playerViewController)
	}
}

// HasPlayerViewControllerDidStartPictureInPicture returns true if a handler for PlayerViewControllerDidStartPictureInPicture has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidStartPictureInPicture() bool {
	return d._PlayerViewControllerDidStartPictureInPicture != nil
}

// PlayerViewControllerDidStopPictureInPicture implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidStopPictureInPicture(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerDidStopPictureInPicture != nil {
		d._PlayerViewControllerDidStopPictureInPicture(playerViewController)
	}
}

// HasPlayerViewControllerDidStopPictureInPicture returns true if a handler for PlayerViewControllerDidStopPictureInPicture has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidStopPictureInPicture() bool {
	return d._PlayerViewControllerDidStopPictureInPicture != nil
}

// PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart(playerViewController IAVPlayerViewController) bool {
	if d._PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart != nil {
		return d._PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart(playerViewController)
	}
	var zero bool
	return zero
}

// HasPlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart returns true if a handler for PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart() bool {
	return d._PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart != nil
}

// PlayerViewControllerShouldDismiss implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerShouldDismiss(playerViewController IAVPlayerViewController) bool {
	if d._PlayerViewControllerShouldDismiss != nil {
		return d._PlayerViewControllerShouldDismiss(playerViewController)
	}
	var zero bool
	return zero
}

// HasPlayerViewControllerShouldDismiss returns true if a handler for PlayerViewControllerShouldDismiss has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerShouldDismiss() bool {
	return d._PlayerViewControllerShouldDismiss != nil
}

// PlayerViewControllerWillBeginDismissalTransition implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillBeginDismissalTransition(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerWillBeginDismissalTransition != nil {
		d._PlayerViewControllerWillBeginDismissalTransition(playerViewController)
	}
}

// HasPlayerViewControllerWillBeginDismissalTransition returns true if a handler for PlayerViewControllerWillBeginDismissalTransition has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillBeginDismissalTransition() bool {
	return d._PlayerViewControllerWillBeginDismissalTransition != nil
}

// PlayerViewControllerWillStartPictureInPicture implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillStartPictureInPicture(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerWillStartPictureInPicture != nil {
		d._PlayerViewControllerWillStartPictureInPicture(playerViewController)
	}
}

// HasPlayerViewControllerWillStartPictureInPicture returns true if a handler for PlayerViewControllerWillStartPictureInPicture has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillStartPictureInPicture() bool {
	return d._PlayerViewControllerWillStartPictureInPicture != nil
}

// PlayerViewControllerWillStopPictureInPicture implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerWillStopPictureInPicture(playerViewController IAVPlayerViewController) {
	if d._PlayerViewControllerWillStopPictureInPicture != nil {
		d._PlayerViewControllerWillStopPictureInPicture(playerViewController)
	}
}

// HasPlayerViewControllerWillStopPictureInPicture returns true if a handler for PlayerViewControllerWillStopPictureInPicture has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerWillStopPictureInPicture() bool {
	return d._PlayerViewControllerWillStopPictureInPicture != nil
}

// PreviousChannelInterstitialViewControllerForPlayerViewController implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PreviousChannelInterstitialViewControllerForPlayerViewController(playerViewController IAVPlayerViewController) ViewController {
	if d._PreviousChannelInterstitialViewControllerForPlayerViewController != nil {
		return d._PreviousChannelInterstitialViewControllerForPlayerViewController(playerViewController)
	}
	var zero ViewController
	return zero
}

// HasPreviousChannelInterstitialViewControllerForPlayerViewController returns true if a handler for PreviousChannelInterstitialViewControllerForPlayerViewController has been set.
func (d *PlayerViewControllerDelegate) HasPreviousChannelInterstitialViewControllerForPlayerViewController() bool {
	return d._PreviousChannelInterstitialViewControllerForPlayerViewController != nil
}

// SkipToNextItemForPlayerViewController implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) SkipToNextItemForPlayerViewController(playerViewController IAVPlayerViewController) {
	if d._SkipToNextItemForPlayerViewController != nil {
		d._SkipToNextItemForPlayerViewController(playerViewController)
	}
}

// HasSkipToNextItemForPlayerViewController returns true if a handler for SkipToNextItemForPlayerViewController has been set.
func (d *PlayerViewControllerDelegate) HasSkipToNextItemForPlayerViewController() bool {
	return d._SkipToNextItemForPlayerViewController != nil
}

// SkipToPreviousItemForPlayerViewController implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) SkipToPreviousItemForPlayerViewController(playerViewController IAVPlayerViewController) {
	if d._SkipToPreviousItemForPlayerViewController != nil {
		d._SkipToPreviousItemForPlayerViewController(playerViewController)
	}
}

// HasSkipToPreviousItemForPlayerViewController returns true if a handler for SkipToPreviousItemForPlayerViewController has been set.
func (d *PlayerViewControllerDelegate) HasSkipToPreviousItemForPlayerViewController() bool {
	return d._SkipToPreviousItemForPlayerViewController != nil
}

// PlayerViewControllerDidSelectExternalSubtitleOptionLanguage implements the PPlayerViewControllerDelegate interface.
func (d *PlayerViewControllerDelegate) PlayerViewControllerDidSelectExternalSubtitleOptionLanguage(playerViewController IAVPlayerViewController, language objc.IObject /* cross-framework: NSString */) {
	if d._PlayerViewControllerDidSelectExternalSubtitleOptionLanguage != nil {
		d._PlayerViewControllerDidSelectExternalSubtitleOptionLanguage(playerViewController, language)
	}
}

// HasPlayerViewControllerDidSelectExternalSubtitleOptionLanguage returns true if a handler for PlayerViewControllerDidSelectExternalSubtitleOptionLanguage has been set.
func (d *PlayerViewControllerDelegate) HasPlayerViewControllerDidSelectExternalSubtitleOptionLanguage() bool {
	return d._PlayerViewControllerDidSelectExternalSubtitleOptionLanguage != nil
}
