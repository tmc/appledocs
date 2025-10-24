// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/cloudkit"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PApplicationDelegate is the NSApplicationDelegate protocol interface.
//
// A set of methods that manage your app’s life cycle and its interaction with common system services.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSApplicationDelegate
type PApplicationDelegate interface {
	// Optional methods
	ApplicationContinueUserActivityRestorationHandler(application IApplication, userActivity foundation.UserActivity, restorationHandler unsafe.Pointer) bool
	HasApplicationContinueUserActivityRestorationHandler() bool
	ApplicationDelegateHandlesKey(sender IApplication, key objc.IObject /* cross-framework: NSString */) bool
	HasApplicationDelegateHandlesKey() bool
	ApplicationDidDecodeRestorableState(app IApplication, coder foundation.Coder)
	HasApplicationDidDecodeRestorableState() bool
	ApplicationDidFailToContinueUserActivityWithTypeError(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	HasApplicationDidFailToContinueUserActivityWithTypeError() bool
	ApplicationDidFailToRegisterForRemoteNotificationsWithError(application IApplication, error_ objc.IObject /* cross-framework: Error */)
	HasApplicationDidFailToRegisterForRemoteNotificationsWithError() bool
	ApplicationDidReceiveRemoteNotification(application IApplication, userInfo foundation.IDictionary)
	HasApplicationDidReceiveRemoteNotification() bool
	ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken objc.IObject /* cross-framework: NSData */)
	HasApplicationDidRegisterForRemoteNotificationsWithDeviceToken() bool
	ApplicationDidUpdateUserActivity(application IApplication, userActivity foundation.UserActivity)
	HasApplicationDidUpdateUserActivity() bool
	ApplicationHandlerForIntent(application IApplication, intent objectivec.IObject) objc.ID
	HasApplicationHandlerForIntent() bool
	ApplicationOpenURLs(application IApplication, urls []foundation.URL)
	HasApplicationOpenURLs() bool
	ApplicationOpenFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	HasApplicationOpenFile() bool
	ApplicationOpenFiles(sender IApplication, filenames []string)
	HasApplicationOpenFiles() bool
	ApplicationOpenFileWithoutUI(sender objc.IObject, filename objc.IObject /* cross-framework: NSString */) bool
	HasApplicationOpenFileWithoutUI() bool
	ApplicationOpenTempFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	HasApplicationOpenTempFile() bool
	ApplicationPrintFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	HasApplicationPrintFile() bool
	ApplicationPrintFilesWithSettingsShowPrintPanels(application IApplication, fileNames []string, printSettings foundation.IDictionary, showPrintPanels bool) ApplicationPrintReply
	HasApplicationPrintFilesWithSettingsShowPrintPanels() bool
	ApplicationUserDidAcceptCloudKitShareWithMetadata(application IApplication, metadata objc.IObject)
	HasApplicationUserDidAcceptCloudKitShareWithMetadata() bool
	ApplicationWillContinueUserActivityWithType(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */) bool
	HasApplicationWillContinueUserActivityWithType() bool
	ApplicationWillEncodeRestorableState(app IApplication, coder foundation.Coder)
	HasApplicationWillEncodeRestorableState() bool
	ApplicationWillPresentError(application IApplication, error_ objc.IObject /* cross-framework: Error */) coretelephony.Error
	HasApplicationWillPresentError() bool
	ApplicationDidFinishLaunching(notification foundation.Notification)
	HasApplicationDidFinishLaunching() bool
	ApplicationWillBecomeActive(notification foundation.Notification)
	HasApplicationWillBecomeActive() bool
	ApplicationWillFinishLaunching(notification foundation.Notification)
	HasApplicationWillFinishLaunching() bool
	ApplicationDidBecomeActive(notification foundation.Notification)
	HasApplicationDidBecomeActive() bool
	ApplicationDidChangeOcclusionState(notification foundation.Notification)
	HasApplicationDidChangeOcclusionState() bool
	ApplicationDidChangeScreenParameters(notification foundation.Notification)
	HasApplicationDidChangeScreenParameters() bool
	ApplicationDidHide(notification foundation.Notification)
	HasApplicationDidHide() bool
	ApplicationDidResignActive(notification foundation.Notification)
	HasApplicationDidResignActive() bool
	ApplicationDidUnhide(notification foundation.Notification)
	HasApplicationDidUnhide() bool
	ApplicationDidUpdate(notification foundation.Notification)
	HasApplicationDidUpdate() bool
	ApplicationDockMenu(sender IApplication) Menu
	HasApplicationDockMenu() bool
	ApplicationOpenUntitledFile(sender IApplication) bool
	HasApplicationOpenUntitledFile() bool
	ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification)
	HasApplicationProtectedDataDidBecomeAvailable() bool
	ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification)
	HasApplicationProtectedDataWillBecomeUnavailable() bool
	ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool
	HasApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool
	ApplicationShouldHandleReopenHasVisibleWindows(sender IApplication, hasVisibleWindows bool) bool
	HasApplicationShouldHandleReopenHasVisibleWindows() bool
	ApplicationShouldOpenUntitledFile(sender IApplication) bool
	HasApplicationShouldOpenUntitledFile() bool
	ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply
	HasApplicationShouldTerminate() bool
	ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool
	HasApplicationShouldTerminateAfterLastWindowClosed() bool
	ApplicationSupportsSecureRestorableState(app IApplication) bool
	HasApplicationSupportsSecureRestorableState() bool
	ApplicationWillHide(notification foundation.Notification)
	HasApplicationWillHide() bool
	ApplicationWillResignActive(notification foundation.Notification)
	HasApplicationWillResignActive() bool
	ApplicationWillTerminate(notification foundation.Notification)
	HasApplicationWillTerminate() bool
	ApplicationWillUnhide(notification foundation.Notification)
	HasApplicationWillUnhide() bool
	ApplicationWillUpdate(notification foundation.Notification)
	HasApplicationWillUpdate() bool
}

// ApplicationDelegate is a delegate implementation builder for the PApplicationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ApplicationDelegate struct {
	_ApplicationContinueUserActivityRestorationHandler func(application IApplication, userActivity foundation.UserActivity, restorationHandler unsafe.Pointer) bool
	_ApplicationDelegateHandlesKey func(sender IApplication, key objc.IObject /* cross-framework: NSString */) bool
	_ApplicationDidDecodeRestorableState func(app IApplication, coder foundation.Coder)
	_ApplicationDidFailToContinueUserActivityWithTypeError func(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	_ApplicationDidFailToRegisterForRemoteNotificationsWithError func(application IApplication, error_ objc.IObject /* cross-framework: Error */)
	_ApplicationDidReceiveRemoteNotification func(application IApplication, userInfo foundation.IDictionary)
	_ApplicationDidRegisterForRemoteNotificationsWithDeviceToken func(application IApplication, deviceToken objc.IObject /* cross-framework: NSData */)
	_ApplicationDidUpdateUserActivity func(application IApplication, userActivity foundation.UserActivity)
	_ApplicationHandlerForIntent func(application IApplication, intent objectivec.IObject) objc.ID
	_ApplicationOpenURLs func(application IApplication, urls []foundation.URL)
	_ApplicationOpenFile func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	_ApplicationOpenFiles func(sender IApplication, filenames []string)
	_ApplicationOpenFileWithoutUI func(sender objc.IObject, filename objc.IObject /* cross-framework: NSString */) bool
	_ApplicationOpenTempFile func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	_ApplicationPrintFile func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool
	_ApplicationPrintFilesWithSettingsShowPrintPanels func(application IApplication, fileNames []string, printSettings foundation.IDictionary, showPrintPanels bool) ApplicationPrintReply
	_ApplicationUserDidAcceptCloudKitShareWithMetadata func(application IApplication, metadata objc.IObject)
	_ApplicationWillContinueUserActivityWithType func(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */) bool
	_ApplicationWillEncodeRestorableState func(app IApplication, coder foundation.Coder)
	_ApplicationWillPresentError func(application IApplication, error_ objc.IObject /* cross-framework: Error */) coretelephony.Error
	_ApplicationDidFinishLaunching func(notification foundation.Notification)
	_ApplicationWillBecomeActive func(notification foundation.Notification)
	_ApplicationWillFinishLaunching func(notification foundation.Notification)
	_ApplicationDidBecomeActive func(notification foundation.Notification)
	_ApplicationDidChangeOcclusionState func(notification foundation.Notification)
	_ApplicationDidChangeScreenParameters func(notification foundation.Notification)
	_ApplicationDidHide func(notification foundation.Notification)
	_ApplicationDidResignActive func(notification foundation.Notification)
	_ApplicationDidUnhide func(notification foundation.Notification)
	_ApplicationDidUpdate func(notification foundation.Notification)
	_ApplicationDockMenu func(sender IApplication) Menu
	_ApplicationOpenUntitledFile func(sender IApplication) bool
	_ApplicationProtectedDataDidBecomeAvailable func(notification foundation.Notification)
	_ApplicationProtectedDataWillBecomeUnavailable func(notification foundation.Notification)
	_ApplicationShouldAutomaticallyLocalizeKeyEquivalents func(application IApplication) bool
	_ApplicationShouldHandleReopenHasVisibleWindows func(sender IApplication, hasVisibleWindows bool) bool
	_ApplicationShouldOpenUntitledFile func(sender IApplication) bool
	_ApplicationShouldTerminate func(sender IApplication) ApplicationTerminateReply
	_ApplicationShouldTerminateAfterLastWindowClosed func(sender IApplication) bool
	_ApplicationSupportsSecureRestorableState func(app IApplication) bool
	_ApplicationWillHide func(notification foundation.Notification)
	_ApplicationWillResignActive func(notification foundation.Notification)
	_ApplicationWillTerminate func(notification foundation.Notification)
	_ApplicationWillUnhide func(notification foundation.Notification)
	_ApplicationWillUpdate func(notification foundation.Notification)
}

// SetApplicationContinueUserActivityRestorationHandler sets the handler for the ApplicationContinueUserActivityRestorationHandler delegate method.
//
// Returns a Boolean value that indicates if the app successfully recreates the specified activity.
func (d *ApplicationDelegate) SetApplicationContinueUserActivityRestorationHandler(f func(application IApplication, userActivity foundation.UserActivity, restorationHandler unsafe.Pointer) bool) {
	d._ApplicationContinueUserActivityRestorationHandler = f
}

// SetApplicationDelegateHandlesKey sets the handler for the ApplicationDelegateHandlesKey delegate method.
//
// Returns a Boolean value that indicates if the app supports the specified scripting key.
func (d *ApplicationDelegate) SetApplicationDelegateHandlesKey(f func(sender IApplication, key objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationDelegateHandlesKey = f
}

// SetApplicationDidDecodeRestorableState sets the handler for the ApplicationDidDecodeRestorableState delegate method.
//
// Tells the delegate when the app finished decoding its restorable state.
func (d *ApplicationDelegate) SetApplicationDidDecodeRestorableState(f func(app IApplication, coder foundation.Coder)) {
	d._ApplicationDidDecodeRestorableState = f
}

// SetApplicationDidFailToContinueUserActivityWithTypeError sets the handler for the ApplicationDidFailToContinueUserActivityWithTypeError delegate method.
//
// Tells the delegate that the app couldn’t continue the specified activity.
func (d *ApplicationDelegate) SetApplicationDidFailToContinueUserActivityWithTypeError(f func(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)) {
	d._ApplicationDidFailToContinueUserActivityWithTypeError = f
}

// SetApplicationDidFailToRegisterForRemoteNotificationsWithError sets the handler for the ApplicationDidFailToRegisterForRemoteNotificationsWithError delegate method.
//
// Tells the delegate that the app was unable to register for Apple Push Services.
func (d *ApplicationDelegate) SetApplicationDidFailToRegisterForRemoteNotificationsWithError(f func(application IApplication, error_ objc.IObject /* cross-framework: Error */)) {
	d._ApplicationDidFailToRegisterForRemoteNotificationsWithError = f
}

// SetApplicationDidReceiveRemoteNotification sets the handler for the ApplicationDidReceiveRemoteNotification delegate method.
//
// Tells the delegate when the app receives a remote notification.
func (d *ApplicationDelegate) SetApplicationDidReceiveRemoteNotification(f func(application IApplication, userInfo foundation.IDictionary)) {
	d._ApplicationDidReceiveRemoteNotification = f
}

// SetApplicationDidRegisterForRemoteNotificationsWithDeviceToken sets the handler for the ApplicationDidRegisterForRemoteNotificationsWithDeviceToken delegate method.
//
// Tells the delegate that the app registered for Apple Push Services.
func (d *ApplicationDelegate) SetApplicationDidRegisterForRemoteNotificationsWithDeviceToken(f func(application IApplication, deviceToken objc.IObject /* cross-framework: NSData */)) {
	d._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken = f
}

// SetApplicationDidUpdateUserActivity sets the handler for the ApplicationDidUpdateUserActivity delegate method.
//
// Tells the delegate that there are changes to the specified activity.
func (d *ApplicationDelegate) SetApplicationDidUpdateUserActivity(f func(application IApplication, userActivity foundation.UserActivity)) {
	d._ApplicationDidUpdateUserActivity = f
}

// SetApplicationHandlerForIntent sets the handler for the ApplicationHandlerForIntent delegate method.
//
// Returns an intent handler that’s capable of handling the specified intent.
func (d *ApplicationDelegate) SetApplicationHandlerForIntent(f func(application IApplication, intent objectivec.IObject) objc.ID) {
	d._ApplicationHandlerForIntent = f
}

// SetApplicationOpenURLs sets the handler for the ApplicationOpenURLs delegate method.
//
// Tells the delegate to open the resource at the specified URL.
func (d *ApplicationDelegate) SetApplicationOpenURLs(f func(application IApplication, urls []foundation.URL)) {
	d._ApplicationOpenURLs = f
}

// SetApplicationOpenFile sets the handler for the ApplicationOpenFile delegate method.
//
// Returns a Boolean value that indicates if the app opens the specified file.
func (d *ApplicationDelegate) SetApplicationOpenFile(f func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationOpenFile = f
}

// SetApplicationOpenFiles sets the handler for the ApplicationOpenFiles delegate method.
//
// Tells the delegate to open the specified files.
func (d *ApplicationDelegate) SetApplicationOpenFiles(f func(sender IApplication, filenames []string)) {
	d._ApplicationOpenFiles = f
}

// SetApplicationOpenFileWithoutUI sets the handler for the ApplicationOpenFileWithoutUI delegate method.
//
// Returns a Boolean value that indicates if the app opens the specified file without showing its user interface.
func (d *ApplicationDelegate) SetApplicationOpenFileWithoutUI(f func(sender objc.IObject, filename objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationOpenFileWithoutUI = f
}

// SetApplicationOpenTempFile sets the handler for the ApplicationOpenTempFile delegate method.
//
// Returns a Boolean value that indicates if the app opens the specified temporary file.
func (d *ApplicationDelegate) SetApplicationOpenTempFile(f func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationOpenTempFile = f
}

// SetApplicationPrintFile sets the handler for the ApplicationPrintFile delegate method.
//
// Returns a Boolean value that indicates if the app prints the specified file in its entirety.
func (d *ApplicationDelegate) SetApplicationPrintFile(f func(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationPrintFile = f
}

// SetApplicationPrintFilesWithSettingsShowPrintPanels sets the handler for the ApplicationPrintFilesWithSettingsShowPrintPanels delegate method.
//
// Returns a value that indicates if the app prints the specified files.
func (d *ApplicationDelegate) SetApplicationPrintFilesWithSettingsShowPrintPanels(f func(application IApplication, fileNames []string, printSettings foundation.IDictionary, showPrintPanels bool) ApplicationPrintReply) {
	d._ApplicationPrintFilesWithSettingsShowPrintPanels = f
}

// SetApplicationUserDidAcceptCloudKitShareWithMetadata sets the handler for the ApplicationUserDidAcceptCloudKitShareWithMetadata delegate method.
//
// Tells the delegate when the user accepts a CloudKit sharing invitation.
func (d *ApplicationDelegate) SetApplicationUserDidAcceptCloudKitShareWithMetadata(f func(application IApplication, metadata objc.IObject)) {
	d._ApplicationUserDidAcceptCloudKitShareWithMetadata = f
}

// SetApplicationWillContinueUserActivityWithType sets the handler for the ApplicationWillContinueUserActivityWithType delegate method.
//
// Returns a Boolean value that indicates if the app can continue the specified activity.
func (d *ApplicationDelegate) SetApplicationWillContinueUserActivityWithType(f func(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */) bool) {
	d._ApplicationWillContinueUserActivityWithType = f
}

// SetApplicationWillEncodeRestorableState sets the handler for the ApplicationWillEncodeRestorableState delegate method.
//
// Tells the delegate that the app is about to encode its restorable state.
func (d *ApplicationDelegate) SetApplicationWillEncodeRestorableState(f func(app IApplication, coder foundation.Coder)) {
	d._ApplicationWillEncodeRestorableState = f
}

// SetApplicationWillPresentError sets the handler for the ApplicationWillPresentError delegate method.
//
// Returns an error for the app to display to the user.
func (d *ApplicationDelegate) SetApplicationWillPresentError(f func(application IApplication, error_ objc.IObject /* cross-framework: Error */) coretelephony.Error) {
	d._ApplicationWillPresentError = f
}

// SetApplicationDidFinishLaunching sets the handler for the ApplicationDidFinishLaunching delegate method.
//
// Tells the delegate that the app’s initialization is complete but it hasn’t received its first event.
func (d *ApplicationDelegate) SetApplicationDidFinishLaunching(f func(notification foundation.Notification)) {
	d._ApplicationDidFinishLaunching = f
}

// SetApplicationWillBecomeActive sets the handler for the ApplicationWillBecomeActive delegate method.
//
// Tells the delegate that the app is about to become active.
func (d *ApplicationDelegate) SetApplicationWillBecomeActive(f func(notification foundation.Notification)) {
	d._ApplicationWillBecomeActive = f
}

// SetApplicationWillFinishLaunching sets the handler for the ApplicationWillFinishLaunching delegate method.
//
// Tells the delegate that the app’s initialization is about to complete.
func (d *ApplicationDelegate) SetApplicationWillFinishLaunching(f func(notification foundation.Notification)) {
	d._ApplicationWillFinishLaunching = f
}

// SetApplicationDidBecomeActive sets the handler for the ApplicationDidBecomeActive delegate method.
//
// Tells the delegate that the app is now active.
func (d *ApplicationDelegate) SetApplicationDidBecomeActive(f func(notification foundation.Notification)) {
	d._ApplicationDidBecomeActive = f
}

// SetApplicationDidChangeOcclusionState sets the handler for the ApplicationDidChangeOcclusionState delegate method.
//
// Tells the delegate about changes to the app’s occlusion state.
func (d *ApplicationDelegate) SetApplicationDidChangeOcclusionState(f func(notification foundation.Notification)) {
	d._ApplicationDidChangeOcclusionState = f
}

// SetApplicationDidChangeScreenParameters sets the handler for the ApplicationDidChangeScreenParameters delegate method.
//
// Tells the delegate about changes to the configuration of any attached displays.
func (d *ApplicationDelegate) SetApplicationDidChangeScreenParameters(f func(notification foundation.Notification)) {
	d._ApplicationDidChangeScreenParameters = f
}

// SetApplicationDidHide sets the handler for the ApplicationDidHide delegate method.
//
// Tells the delegate that the app is now hidden.
func (d *ApplicationDelegate) SetApplicationDidHide(f func(notification foundation.Notification)) {
	d._ApplicationDidHide = f
}

// SetApplicationDidResignActive sets the handler for the ApplicationDidResignActive delegate method.
//
// Tells the delegate that the app is no longer active and doesn’t have focus.
func (d *ApplicationDelegate) SetApplicationDidResignActive(f func(notification foundation.Notification)) {
	d._ApplicationDidResignActive = f
}

// SetApplicationDidUnhide sets the handler for the ApplicationDidUnhide delegate method.
//
// Tells the delegate that the app is now visible.
func (d *ApplicationDelegate) SetApplicationDidUnhide(f func(notification foundation.Notification)) {
	d._ApplicationDidUnhide = f
}

// SetApplicationDidUpdate sets the handler for the ApplicationDidUpdate delegate method.
//
// Tells the delegate that the app’s windows did update.
func (d *ApplicationDelegate) SetApplicationDidUpdate(f func(notification foundation.Notification)) {
	d._ApplicationDidUpdate = f
}

// SetApplicationDockMenu sets the handler for the ApplicationDockMenu delegate method.
//
// Returns the app’s dock menu.
func (d *ApplicationDelegate) SetApplicationDockMenu(f func(sender IApplication) Menu) {
	d._ApplicationDockMenu = f
}

// SetApplicationOpenUntitledFile sets the handler for the ApplicationOpenUntitledFile delegate method.
//
// Returns a Boolean value that indicates if the app opens an untitled file.
func (d *ApplicationDelegate) SetApplicationOpenUntitledFile(f func(sender IApplication) bool) {
	d._ApplicationOpenUntitledFile = f
}

// SetApplicationProtectedDataDidBecomeAvailable sets the handler for the ApplicationProtectedDataDidBecomeAvailable delegate method.
//
// Tells the delegate that protected data is now available.
func (d *ApplicationDelegate) SetApplicationProtectedDataDidBecomeAvailable(f func(notification foundation.Notification)) {
	d._ApplicationProtectedDataDidBecomeAvailable = f
}

// SetApplicationProtectedDataWillBecomeUnavailable sets the handler for the ApplicationProtectedDataWillBecomeUnavailable delegate method.
//
// Tells the delegate that protected data is about to become unavailable.
func (d *ApplicationDelegate) SetApplicationProtectedDataWillBecomeUnavailable(f func(notification foundation.Notification)) {
	d._ApplicationProtectedDataWillBecomeUnavailable = f
}

// SetApplicationShouldAutomaticallyLocalizeKeyEquivalents sets the handler for the ApplicationShouldAutomaticallyLocalizeKeyEquivalents delegate method.
//
// Returns a Boolean value that tells the system whether to remap menu shortcuts to support localized keyboards.
func (d *ApplicationDelegate) SetApplicationShouldAutomaticallyLocalizeKeyEquivalents(f func(application IApplication) bool) {
	d._ApplicationShouldAutomaticallyLocalizeKeyEquivalents = f
}

// SetApplicationShouldHandleReopenHasVisibleWindows sets the handler for the ApplicationShouldHandleReopenHasVisibleWindows delegate method.
//
// Returns a Boolean value that indicates if the app responds to reopen AppleEvents.
func (d *ApplicationDelegate) SetApplicationShouldHandleReopenHasVisibleWindows(f func(sender IApplication, hasVisibleWindows bool) bool) {
	d._ApplicationShouldHandleReopenHasVisibleWindows = f
}

// SetApplicationShouldOpenUntitledFile sets the handler for the ApplicationShouldOpenUntitledFile delegate method.
//
// Returns a Boolean value that indicates if the app can open an untitled file.
func (d *ApplicationDelegate) SetApplicationShouldOpenUntitledFile(f func(sender IApplication) bool) {
	d._ApplicationShouldOpenUntitledFile = f
}

// SetApplicationShouldTerminate sets the handler for the ApplicationShouldTerminate delegate method.
//
// Returns a value that indicates if the app should terminate.
func (d *ApplicationDelegate) SetApplicationShouldTerminate(f func(sender IApplication) ApplicationTerminateReply) {
	d._ApplicationShouldTerminate = f
}

// SetApplicationShouldTerminateAfterLastWindowClosed sets the handler for the ApplicationShouldTerminateAfterLastWindowClosed delegate method.
//
// Returns a Boolean value that indicates if the app terminates once the last window closes.
func (d *ApplicationDelegate) SetApplicationShouldTerminateAfterLastWindowClosed(f func(sender IApplication) bool) {
	d._ApplicationShouldTerminateAfterLastWindowClosed = f
}

// SetApplicationSupportsSecureRestorableState sets the handler for the ApplicationSupportsSecureRestorableState delegate method.
//
// Returns a Boolean value that indicates if the app supports secure state restoration.
func (d *ApplicationDelegate) SetApplicationSupportsSecureRestorableState(f func(app IApplication) bool) {
	d._ApplicationSupportsSecureRestorableState = f
}

// SetApplicationWillHide sets the handler for the ApplicationWillHide delegate method.
//
// Tells the delegate that the app is about to be hidden.
func (d *ApplicationDelegate) SetApplicationWillHide(f func(notification foundation.Notification)) {
	d._ApplicationWillHide = f
}

// SetApplicationWillResignActive sets the handler for the ApplicationWillResignActive delegate method.
//
// Tells the delegate that the app is about to become inactive and will lose focus.
func (d *ApplicationDelegate) SetApplicationWillResignActive(f func(notification foundation.Notification)) {
	d._ApplicationWillResignActive = f
}

// SetApplicationWillTerminate sets the handler for the ApplicationWillTerminate delegate method.
//
// Tells the delegate that the app is about to terminate.
func (d *ApplicationDelegate) SetApplicationWillTerminate(f func(notification foundation.Notification)) {
	d._ApplicationWillTerminate = f
}

// SetApplicationWillUnhide sets the handler for the ApplicationWillUnhide delegate method.
//
// Tells the delegate that the app is about to become visible.
func (d *ApplicationDelegate) SetApplicationWillUnhide(f func(notification foundation.Notification)) {
	d._ApplicationWillUnhide = f
}

// SetApplicationWillUpdate sets the handler for the ApplicationWillUpdate delegate method.
//
// Tells the delegate that the app is about to update its windows.
func (d *ApplicationDelegate) SetApplicationWillUpdate(f func(notification foundation.Notification)) {
	d._ApplicationWillUpdate = f
}

// ApplicationContinueUserActivityRestorationHandler implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationContinueUserActivityRestorationHandler(application IApplication, userActivity foundation.UserActivity, restorationHandler unsafe.Pointer) bool {
	if d._ApplicationContinueUserActivityRestorationHandler != nil {
		return d._ApplicationContinueUserActivityRestorationHandler(application, userActivity, restorationHandler)
	}
	var zero bool
	return zero
}

// HasApplicationContinueUserActivityRestorationHandler returns true if a handler for ApplicationContinueUserActivityRestorationHandler has been set.
func (d *ApplicationDelegate) HasApplicationContinueUserActivityRestorationHandler() bool {
	return d._ApplicationContinueUserActivityRestorationHandler != nil
}

// ApplicationDelegateHandlesKey implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDelegateHandlesKey(sender IApplication, key objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationDelegateHandlesKey != nil {
		return d._ApplicationDelegateHandlesKey(sender, key)
	}
	var zero bool
	return zero
}

// HasApplicationDelegateHandlesKey returns true if a handler for ApplicationDelegateHandlesKey has been set.
func (d *ApplicationDelegate) HasApplicationDelegateHandlesKey() bool {
	return d._ApplicationDelegateHandlesKey != nil
}

// ApplicationDidDecodeRestorableState implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidDecodeRestorableState(app IApplication, coder foundation.Coder) {
	if d._ApplicationDidDecodeRestorableState != nil {
		d._ApplicationDidDecodeRestorableState(app, coder)
	}
}

// HasApplicationDidDecodeRestorableState returns true if a handler for ApplicationDidDecodeRestorableState has been set.
func (d *ApplicationDelegate) HasApplicationDidDecodeRestorableState() bool {
	return d._ApplicationDidDecodeRestorableState != nil
}

// ApplicationDidFailToContinueUserActivityWithTypeError implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidFailToContinueUserActivityWithTypeError(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */) {
	if d._ApplicationDidFailToContinueUserActivityWithTypeError != nil {
		d._ApplicationDidFailToContinueUserActivityWithTypeError(application, userActivityType, error_)
	}
}

// HasApplicationDidFailToContinueUserActivityWithTypeError returns true if a handler for ApplicationDidFailToContinueUserActivityWithTypeError has been set.
func (d *ApplicationDelegate) HasApplicationDidFailToContinueUserActivityWithTypeError() bool {
	return d._ApplicationDidFailToContinueUserActivityWithTypeError != nil
}

// ApplicationDidFailToRegisterForRemoteNotificationsWithError implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidFailToRegisterForRemoteNotificationsWithError(application IApplication, error_ objc.IObject /* cross-framework: Error */) {
	if d._ApplicationDidFailToRegisterForRemoteNotificationsWithError != nil {
		d._ApplicationDidFailToRegisterForRemoteNotificationsWithError(application, error_)
	}
}

// HasApplicationDidFailToRegisterForRemoteNotificationsWithError returns true if a handler for ApplicationDidFailToRegisterForRemoteNotificationsWithError has been set.
func (d *ApplicationDelegate) HasApplicationDidFailToRegisterForRemoteNotificationsWithError() bool {
	return d._ApplicationDidFailToRegisterForRemoteNotificationsWithError != nil
}

// ApplicationDidReceiveRemoteNotification implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidReceiveRemoteNotification(application IApplication, userInfo foundation.IDictionary) {
	if d._ApplicationDidReceiveRemoteNotification != nil {
		d._ApplicationDidReceiveRemoteNotification(application, userInfo)
	}
}

// HasApplicationDidReceiveRemoteNotification returns true if a handler for ApplicationDidReceiveRemoteNotification has been set.
func (d *ApplicationDelegate) HasApplicationDidReceiveRemoteNotification() bool {
	return d._ApplicationDidReceiveRemoteNotification != nil
}

// ApplicationDidRegisterForRemoteNotificationsWithDeviceToken implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken objc.IObject /* cross-framework: NSData */) {
	if d._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken != nil {
		d._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application, deviceToken)
	}
}

// HasApplicationDidRegisterForRemoteNotificationsWithDeviceToken returns true if a handler for ApplicationDidRegisterForRemoteNotificationsWithDeviceToken has been set.
func (d *ApplicationDelegate) HasApplicationDidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return d._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken != nil
}

// ApplicationDidUpdateUserActivity implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidUpdateUserActivity(application IApplication, userActivity foundation.UserActivity) {
	if d._ApplicationDidUpdateUserActivity != nil {
		d._ApplicationDidUpdateUserActivity(application, userActivity)
	}
}

// HasApplicationDidUpdateUserActivity returns true if a handler for ApplicationDidUpdateUserActivity has been set.
func (d *ApplicationDelegate) HasApplicationDidUpdateUserActivity() bool {
	return d._ApplicationDidUpdateUserActivity != nil
}

// ApplicationHandlerForIntent implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationHandlerForIntent(application IApplication, intent objectivec.IObject) objc.ID {
	if d._ApplicationHandlerForIntent != nil {
		return d._ApplicationHandlerForIntent(application, intent)
	}
	var zero objc.ID
	return zero
}

// HasApplicationHandlerForIntent returns true if a handler for ApplicationHandlerForIntent has been set.
func (d *ApplicationDelegate) HasApplicationHandlerForIntent() bool {
	return d._ApplicationHandlerForIntent != nil
}

// ApplicationOpenURLs implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenURLs(application IApplication, urls []foundation.URL) {
	if d._ApplicationOpenURLs != nil {
		d._ApplicationOpenURLs(application, urls)
	}
}

// HasApplicationOpenURLs returns true if a handler for ApplicationOpenURLs has been set.
func (d *ApplicationDelegate) HasApplicationOpenURLs() bool {
	return d._ApplicationOpenURLs != nil
}

// ApplicationOpenFile implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationOpenFile != nil {
		return d._ApplicationOpenFile(sender, filename)
	}
	var zero bool
	return zero
}

// HasApplicationOpenFile returns true if a handler for ApplicationOpenFile has been set.
func (d *ApplicationDelegate) HasApplicationOpenFile() bool {
	return d._ApplicationOpenFile != nil
}

// ApplicationOpenFiles implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenFiles(sender IApplication, filenames []string) {
	if d._ApplicationOpenFiles != nil {
		d._ApplicationOpenFiles(sender, filenames)
	}
}

// HasApplicationOpenFiles returns true if a handler for ApplicationOpenFiles has been set.
func (d *ApplicationDelegate) HasApplicationOpenFiles() bool {
	return d._ApplicationOpenFiles != nil
}

// ApplicationOpenFileWithoutUI implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenFileWithoutUI(sender objc.IObject, filename objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationOpenFileWithoutUI != nil {
		return d._ApplicationOpenFileWithoutUI(sender, filename)
	}
	var zero bool
	return zero
}

// HasApplicationOpenFileWithoutUI returns true if a handler for ApplicationOpenFileWithoutUI has been set.
func (d *ApplicationDelegate) HasApplicationOpenFileWithoutUI() bool {
	return d._ApplicationOpenFileWithoutUI != nil
}

// ApplicationOpenTempFile implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenTempFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationOpenTempFile != nil {
		return d._ApplicationOpenTempFile(sender, filename)
	}
	var zero bool
	return zero
}

// HasApplicationOpenTempFile returns true if a handler for ApplicationOpenTempFile has been set.
func (d *ApplicationDelegate) HasApplicationOpenTempFile() bool {
	return d._ApplicationOpenTempFile != nil
}

// ApplicationPrintFile implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationPrintFile(sender IApplication, filename objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationPrintFile != nil {
		return d._ApplicationPrintFile(sender, filename)
	}
	var zero bool
	return zero
}

// HasApplicationPrintFile returns true if a handler for ApplicationPrintFile has been set.
func (d *ApplicationDelegate) HasApplicationPrintFile() bool {
	return d._ApplicationPrintFile != nil
}

// ApplicationPrintFilesWithSettingsShowPrintPanels implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationPrintFilesWithSettingsShowPrintPanels(application IApplication, fileNames []string, printSettings foundation.IDictionary, showPrintPanels bool) ApplicationPrintReply {
	if d._ApplicationPrintFilesWithSettingsShowPrintPanels != nil {
		return d._ApplicationPrintFilesWithSettingsShowPrintPanels(application, fileNames, printSettings, showPrintPanels)
	}
	var zero ApplicationPrintReply
	return zero
}

// HasApplicationPrintFilesWithSettingsShowPrintPanels returns true if a handler for ApplicationPrintFilesWithSettingsShowPrintPanels has been set.
func (d *ApplicationDelegate) HasApplicationPrintFilesWithSettingsShowPrintPanels() bool {
	return d._ApplicationPrintFilesWithSettingsShowPrintPanels != nil
}

// ApplicationUserDidAcceptCloudKitShareWithMetadata implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationUserDidAcceptCloudKitShareWithMetadata(application IApplication, metadata objc.IObject) {
	if d._ApplicationUserDidAcceptCloudKitShareWithMetadata != nil {
		d._ApplicationUserDidAcceptCloudKitShareWithMetadata(application, metadata)
	}
}

// HasApplicationUserDidAcceptCloudKitShareWithMetadata returns true if a handler for ApplicationUserDidAcceptCloudKitShareWithMetadata has been set.
func (d *ApplicationDelegate) HasApplicationUserDidAcceptCloudKitShareWithMetadata() bool {
	return d._ApplicationUserDidAcceptCloudKitShareWithMetadata != nil
}

// ApplicationWillContinueUserActivityWithType implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillContinueUserActivityWithType(application IApplication, userActivityType objc.IObject /* cross-framework: NSString */) bool {
	if d._ApplicationWillContinueUserActivityWithType != nil {
		return d._ApplicationWillContinueUserActivityWithType(application, userActivityType)
	}
	var zero bool
	return zero
}

// HasApplicationWillContinueUserActivityWithType returns true if a handler for ApplicationWillContinueUserActivityWithType has been set.
func (d *ApplicationDelegate) HasApplicationWillContinueUserActivityWithType() bool {
	return d._ApplicationWillContinueUserActivityWithType != nil
}

// ApplicationWillEncodeRestorableState implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillEncodeRestorableState(app IApplication, coder foundation.Coder) {
	if d._ApplicationWillEncodeRestorableState != nil {
		d._ApplicationWillEncodeRestorableState(app, coder)
	}
}

// HasApplicationWillEncodeRestorableState returns true if a handler for ApplicationWillEncodeRestorableState has been set.
func (d *ApplicationDelegate) HasApplicationWillEncodeRestorableState() bool {
	return d._ApplicationWillEncodeRestorableState != nil
}

// ApplicationWillPresentError implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillPresentError(application IApplication, error_ objc.IObject /* cross-framework: Error */) coretelephony.Error {
	if d._ApplicationWillPresentError != nil {
		return d._ApplicationWillPresentError(application, error_)
	}
	var zero coretelephony.Error
	return zero
}

// HasApplicationWillPresentError returns true if a handler for ApplicationWillPresentError has been set.
func (d *ApplicationDelegate) HasApplicationWillPresentError() bool {
	return d._ApplicationWillPresentError != nil
}

// ApplicationDidFinishLaunching implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	if d._ApplicationDidFinishLaunching != nil {
		d._ApplicationDidFinishLaunching(notification)
	}
}

// HasApplicationDidFinishLaunching returns true if a handler for ApplicationDidFinishLaunching has been set.
func (d *ApplicationDelegate) HasApplicationDidFinishLaunching() bool {
	return d._ApplicationDidFinishLaunching != nil
}

// ApplicationWillBecomeActive implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillBecomeActive(notification foundation.Notification) {
	if d._ApplicationWillBecomeActive != nil {
		d._ApplicationWillBecomeActive(notification)
	}
}

// HasApplicationWillBecomeActive returns true if a handler for ApplicationWillBecomeActive has been set.
func (d *ApplicationDelegate) HasApplicationWillBecomeActive() bool {
	return d._ApplicationWillBecomeActive != nil
}

// ApplicationWillFinishLaunching implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillFinishLaunching(notification foundation.Notification) {
	if d._ApplicationWillFinishLaunching != nil {
		d._ApplicationWillFinishLaunching(notification)
	}
}

// HasApplicationWillFinishLaunching returns true if a handler for ApplicationWillFinishLaunching has been set.
func (d *ApplicationDelegate) HasApplicationWillFinishLaunching() bool {
	return d._ApplicationWillFinishLaunching != nil
}

// ApplicationDidBecomeActive implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidBecomeActive(notification foundation.Notification) {
	if d._ApplicationDidBecomeActive != nil {
		d._ApplicationDidBecomeActive(notification)
	}
}

// HasApplicationDidBecomeActive returns true if a handler for ApplicationDidBecomeActive has been set.
func (d *ApplicationDelegate) HasApplicationDidBecomeActive() bool {
	return d._ApplicationDidBecomeActive != nil
}

// ApplicationDidChangeOcclusionState implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidChangeOcclusionState(notification foundation.Notification) {
	if d._ApplicationDidChangeOcclusionState != nil {
		d._ApplicationDidChangeOcclusionState(notification)
	}
}

// HasApplicationDidChangeOcclusionState returns true if a handler for ApplicationDidChangeOcclusionState has been set.
func (d *ApplicationDelegate) HasApplicationDidChangeOcclusionState() bool {
	return d._ApplicationDidChangeOcclusionState != nil
}

// ApplicationDidChangeScreenParameters implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidChangeScreenParameters(notification foundation.Notification) {
	if d._ApplicationDidChangeScreenParameters != nil {
		d._ApplicationDidChangeScreenParameters(notification)
	}
}

// HasApplicationDidChangeScreenParameters returns true if a handler for ApplicationDidChangeScreenParameters has been set.
func (d *ApplicationDelegate) HasApplicationDidChangeScreenParameters() bool {
	return d._ApplicationDidChangeScreenParameters != nil
}

// ApplicationDidHide implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidHide(notification foundation.Notification) {
	if d._ApplicationDidHide != nil {
		d._ApplicationDidHide(notification)
	}
}

// HasApplicationDidHide returns true if a handler for ApplicationDidHide has been set.
func (d *ApplicationDelegate) HasApplicationDidHide() bool {
	return d._ApplicationDidHide != nil
}

// ApplicationDidResignActive implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidResignActive(notification foundation.Notification) {
	if d._ApplicationDidResignActive != nil {
		d._ApplicationDidResignActive(notification)
	}
}

// HasApplicationDidResignActive returns true if a handler for ApplicationDidResignActive has been set.
func (d *ApplicationDelegate) HasApplicationDidResignActive() bool {
	return d._ApplicationDidResignActive != nil
}

// ApplicationDidUnhide implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidUnhide(notification foundation.Notification) {
	if d._ApplicationDidUnhide != nil {
		d._ApplicationDidUnhide(notification)
	}
}

// HasApplicationDidUnhide returns true if a handler for ApplicationDidUnhide has been set.
func (d *ApplicationDelegate) HasApplicationDidUnhide() bool {
	return d._ApplicationDidUnhide != nil
}

// ApplicationDidUpdate implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDidUpdate(notification foundation.Notification) {
	if d._ApplicationDidUpdate != nil {
		d._ApplicationDidUpdate(notification)
	}
}

// HasApplicationDidUpdate returns true if a handler for ApplicationDidUpdate has been set.
func (d *ApplicationDelegate) HasApplicationDidUpdate() bool {
	return d._ApplicationDidUpdate != nil
}

// ApplicationDockMenu implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationDockMenu(sender IApplication) Menu {
	if d._ApplicationDockMenu != nil {
		return d._ApplicationDockMenu(sender)
	}
	var zero Menu
	return zero
}

// HasApplicationDockMenu returns true if a handler for ApplicationDockMenu has been set.
func (d *ApplicationDelegate) HasApplicationDockMenu() bool {
	return d._ApplicationDockMenu != nil
}

// ApplicationOpenUntitledFile implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationOpenUntitledFile(sender IApplication) bool {
	if d._ApplicationOpenUntitledFile != nil {
		return d._ApplicationOpenUntitledFile(sender)
	}
	var zero bool
	return zero
}

// HasApplicationOpenUntitledFile returns true if a handler for ApplicationOpenUntitledFile has been set.
func (d *ApplicationDelegate) HasApplicationOpenUntitledFile() bool {
	return d._ApplicationOpenUntitledFile != nil
}

// ApplicationProtectedDataDidBecomeAvailable implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification) {
	if d._ApplicationProtectedDataDidBecomeAvailable != nil {
		d._ApplicationProtectedDataDidBecomeAvailable(notification)
	}
}

// HasApplicationProtectedDataDidBecomeAvailable returns true if a handler for ApplicationProtectedDataDidBecomeAvailable has been set.
func (d *ApplicationDelegate) HasApplicationProtectedDataDidBecomeAvailable() bool {
	return d._ApplicationProtectedDataDidBecomeAvailable != nil
}

// ApplicationProtectedDataWillBecomeUnavailable implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification) {
	if d._ApplicationProtectedDataWillBecomeUnavailable != nil {
		d._ApplicationProtectedDataWillBecomeUnavailable(notification)
	}
}

// HasApplicationProtectedDataWillBecomeUnavailable returns true if a handler for ApplicationProtectedDataWillBecomeUnavailable has been set.
func (d *ApplicationDelegate) HasApplicationProtectedDataWillBecomeUnavailable() bool {
	return d._ApplicationProtectedDataWillBecomeUnavailable != nil
}

// ApplicationShouldAutomaticallyLocalizeKeyEquivalents implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	if d._ApplicationShouldAutomaticallyLocalizeKeyEquivalents != nil {
		return d._ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application)
	}
	var zero bool
	return zero
}

// HasApplicationShouldAutomaticallyLocalizeKeyEquivalents returns true if a handler for ApplicationShouldAutomaticallyLocalizeKeyEquivalents has been set.
func (d *ApplicationDelegate) HasApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return d._ApplicationShouldAutomaticallyLocalizeKeyEquivalents != nil
}

// ApplicationShouldHandleReopenHasVisibleWindows implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationShouldHandleReopenHasVisibleWindows(sender IApplication, hasVisibleWindows bool) bool {
	if d._ApplicationShouldHandleReopenHasVisibleWindows != nil {
		return d._ApplicationShouldHandleReopenHasVisibleWindows(sender, hasVisibleWindows)
	}
	var zero bool
	return zero
}

// HasApplicationShouldHandleReopenHasVisibleWindows returns true if a handler for ApplicationShouldHandleReopenHasVisibleWindows has been set.
func (d *ApplicationDelegate) HasApplicationShouldHandleReopenHasVisibleWindows() bool {
	return d._ApplicationShouldHandleReopenHasVisibleWindows != nil
}

// ApplicationShouldOpenUntitledFile implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	if d._ApplicationShouldOpenUntitledFile != nil {
		return d._ApplicationShouldOpenUntitledFile(sender)
	}
	var zero bool
	return zero
}

// HasApplicationShouldOpenUntitledFile returns true if a handler for ApplicationShouldOpenUntitledFile has been set.
func (d *ApplicationDelegate) HasApplicationShouldOpenUntitledFile() bool {
	return d._ApplicationShouldOpenUntitledFile != nil
}

// ApplicationShouldTerminate implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	if d._ApplicationShouldTerminate != nil {
		return d._ApplicationShouldTerminate(sender)
	}
	var zero ApplicationTerminateReply
	return zero
}

// HasApplicationShouldTerminate returns true if a handler for ApplicationShouldTerminate has been set.
func (d *ApplicationDelegate) HasApplicationShouldTerminate() bool {
	return d._ApplicationShouldTerminate != nil
}

// ApplicationShouldTerminateAfterLastWindowClosed implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	if d._ApplicationShouldTerminateAfterLastWindowClosed != nil {
		return d._ApplicationShouldTerminateAfterLastWindowClosed(sender)
	}
	var zero bool
	return zero
}

// HasApplicationShouldTerminateAfterLastWindowClosed returns true if a handler for ApplicationShouldTerminateAfterLastWindowClosed has been set.
func (d *ApplicationDelegate) HasApplicationShouldTerminateAfterLastWindowClosed() bool {
	return d._ApplicationShouldTerminateAfterLastWindowClosed != nil
}

// ApplicationSupportsSecureRestorableState implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	if d._ApplicationSupportsSecureRestorableState != nil {
		return d._ApplicationSupportsSecureRestorableState(app)
	}
	var zero bool
	return zero
}

// HasApplicationSupportsSecureRestorableState returns true if a handler for ApplicationSupportsSecureRestorableState has been set.
func (d *ApplicationDelegate) HasApplicationSupportsSecureRestorableState() bool {
	return d._ApplicationSupportsSecureRestorableState != nil
}

// ApplicationWillHide implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillHide(notification foundation.Notification) {
	if d._ApplicationWillHide != nil {
		d._ApplicationWillHide(notification)
	}
}

// HasApplicationWillHide returns true if a handler for ApplicationWillHide has been set.
func (d *ApplicationDelegate) HasApplicationWillHide() bool {
	return d._ApplicationWillHide != nil
}

// ApplicationWillResignActive implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillResignActive(notification foundation.Notification) {
	if d._ApplicationWillResignActive != nil {
		d._ApplicationWillResignActive(notification)
	}
}

// HasApplicationWillResignActive returns true if a handler for ApplicationWillResignActive has been set.
func (d *ApplicationDelegate) HasApplicationWillResignActive() bool {
	return d._ApplicationWillResignActive != nil
}

// ApplicationWillTerminate implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillTerminate(notification foundation.Notification) {
	if d._ApplicationWillTerminate != nil {
		d._ApplicationWillTerminate(notification)
	}
}

// HasApplicationWillTerminate returns true if a handler for ApplicationWillTerminate has been set.
func (d *ApplicationDelegate) HasApplicationWillTerminate() bool {
	return d._ApplicationWillTerminate != nil
}

// ApplicationWillUnhide implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillUnhide(notification foundation.Notification) {
	if d._ApplicationWillUnhide != nil {
		d._ApplicationWillUnhide(notification)
	}
}

// HasApplicationWillUnhide returns true if a handler for ApplicationWillUnhide has been set.
func (d *ApplicationDelegate) HasApplicationWillUnhide() bool {
	return d._ApplicationWillUnhide != nil
}

// ApplicationWillUpdate implements the PApplicationDelegate interface.
func (d *ApplicationDelegate) ApplicationWillUpdate(notification foundation.Notification) {
	if d._ApplicationWillUpdate != nil {
		d._ApplicationWillUpdate(notification)
	}
}

// HasApplicationWillUpdate returns true if a handler for ApplicationWillUpdate has been set.
func (d *ApplicationDelegate) HasApplicationWillUpdate() bool {
	return d._ApplicationWillUpdate != nil
}
