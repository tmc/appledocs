// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserActivity] class.
var (
	UserActivityClass     _UserActivityClass
	UserActivityClassOnce sync.Once
)

func getUserActivityClass() _UserActivityClass {
	UserActivityClassOnce.Do(func() {
		UserActivityClass = _UserActivityClass{objc.GetClass("NSUserActivity")}
	})
	return UserActivityClass
}

type _UserActivityClass struct {
	class objc.Class
}

// An interface definition for the [UserActivity] class.
type IUserActivity interface {
	objectivec.IObject
	AddUserInfoEntriesFromDictionary(otherDictionary objectivec.IObject)
	BecomeCurrent()
	GetContinuationStreamsWithCompletionHandler(completionHandler unsafe.Pointer)
	Invalidate()
	ResignCurrent()
	ActivityType() string
	AppClipActivationPayload() unsafe.Pointer
	ContextIdentifierPath() []string
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ExpirationDate() NSDate
	SetExpirationDate(value IDate)
	ExternalMediaContentIdentifier() string
	SetExternalMediaContentIdentifier(value string)
	IsClassKitDeepLink() bool
	EligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	EligibleForPrediction() bool
	SetEligibleForPrediction(value bool)
	EligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	EligibleForSearch() bool
	SetEligibleForSearch(value bool)
	Keywords() unsafe.Pointer
	SetKeywords(value unsafe.Pointer)
	NdefMessagePayload() unsafe.Pointer
	NeedsSave() bool
	SetNeedsSave(value bool)
	PersistentIdentifier() UserActivityPersistentIdentifier
	SetPersistentIdentifier(value IUserActivityPersistentIdentifier)
	ReferrerURL() URL
	SetReferrerURL(value IURL)
	RequiredUserInfoKeys() unsafe.Pointer
	SetRequiredUserInfoKeys(value unsafe.Pointer)
	ShortcutAvailability() unsafe.Pointer
	SetShortcutAvailability(value unsafe.Pointer)
	SuggestedInvocationPhrase() string
	SetSuggestedInvocationPhrase(value string)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	Title() string
	SetTitle(value string)
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
	WebpageURL() URL
	SetWebpageURL(value IURL)
	TVUserActivityTypeBrowsingChannelGuide() string
	ActivityItemsConfiguration() unsafe.Pointer
	SetActivityItemsConfiguration(value unsafe.Pointer)
	AppEntityIdentifier() unsafe.Pointer
	SetAppEntityIdentifier(value unsafe.Pointer)
	IsEligibleForHandoff() bool
	SetIsEligibleForHandoff(value bool)
	IsEligibleForPrediction() bool
	SetIsEligibleForPrediction(value bool)
	IsEligibleForPublicIndexing() bool
	SetIsEligibleForPublicIndexing(value bool)
	IsEligibleForSearch() bool
	SetIsEligibleForSearch(value bool)
	NSUserActivityConnectionUnavailableError() int
	SetNSUserActivityConnectionUnavailableError(value int)
	NSUserActivityErrorMaximum() int
	SetNSUserActivityErrorMaximum(value int)
	NSUserActivityErrorMinimum() int
	SetNSUserActivityErrorMinimum(value int)
	NSUserActivityHandoffFailedError() int
	SetNSUserActivityHandoffFailedError(value int)
	NSUserActivityHandoffUserInfoTooLargeError() int
	SetNSUserActivityHandoffUserInfoTooLargeError(value int)
	NSUserActivityRemoteApplicationTimedOutError() int
	SetNSUserActivityRemoteApplicationTimedOutError(value int)
	NSUserActivityTypeBrowsingWeb() string
}

// A representation of the state of your app at a moment in time.
//
// An object provides a lightweight way to capture the state of your app and put it to use later. Create this object to capture information about what a person was doing, such as viewing app content, editing a document, viewing a web page, or watching a video. When the system launches your app and an activity object is available, your app can use the information in that object to restore itself to an appropriate state. Spotlight also uses these objects to improve search results for people. To allow people to continue an activity on another device, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity
type UserActivity struct {
	objectivec.Object
}

// UserActivityFrom constructs a [UserActivity] from an unsafe.Pointer.
//
// A representation of the state of your app at a moment in time.
func UserActivityFrom(ptr unsafe.Pointer) UserActivity {
	return UserActivity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.Send[UserActivity](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserActivityClass) New() UserActivity {
	rv := objc.Send[UserActivity](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserActivity) Init() UserActivity {
	rv := objc.Send[UserActivity](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserActivity) Autorelease() UserActivity {
	rv := objc.Send[UserActivity](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserActivity creates a new UserActivity instance.
func NewUserActivity() UserActivity {
	return getUserActivityClass().New()
}




// Creates a user activity object with the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/init(activityType:)
func NewUserActivityWithActivityType(activityType string) UserActivity {
	instance := getUserActivityClass().Alloc()
	rv := objc.Send[UserActivity](instance.ID, objc.Sel("initWithActivityType:"), objc.String(activityType))
	rv.Autorelease()
	return rv
}


// Deletes all user activities created by your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteAllSavedUserActivities(completionHandler:)
func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("deleteAllSavedUserActivitiesWithCompletionHandler:"), handler)
}

// Deletes user activities created by your app that have the specified persistent identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteSavedUserActivities(withPersistentIdentifiers:completionHandler:)
func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []string, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, handler)
}

// Adds the contents of the specified dictionary to the user info dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/addUserInfoEntries(from:)
func (u_ UserActivity) AddUserInfoEntriesFromDictionary(otherDictionary objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addUserInfoEntriesFromDictionary:"), otherDictionary)
}

// Marks the activity as currently in use by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/becomeCurrent()
func (u_ UserActivity) BecomeCurrent() {
	objc.Send[objc.ID](u_.ID, objc.Sel("becomeCurrent"))
}

// Requests streams back to the originating app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/getContinuationStreams(completionHandler:)
func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getContinuationStreamsWithCompletionHandler:"), completionHandler)
}

// Invalidates an activity and marks it as no longer eligible for continuation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/invalidate()
func (u_ UserActivity) Invalidate() {
	objc.Send[objc.ID](u_.ID, objc.Sel("invalidate"))
}

// Marks this activity object as inactive without invalidating it.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/resignCurrent()
func (u_ UserActivity) ResignCurrent() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resignCurrent"))
}

// The user activity object’s activity type.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/activityType
func (u_ UserActivity) ActivityType() string {
	rv := objc.Send[string](u_.ID, objc.Sel("activityType"))
	return rv
}

// An object containing the payload information that launches an App Clip.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/appClipActivationPayload
func (u_ UserActivity) AppClipActivationPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("appClipActivationPayload"))
	return rv
}

// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (u_ UserActivity) ContextIdentifierPath() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("contextIdentifierPath"))
	return rv
}

// The user activity object’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/delegate
func (u_ UserActivity) Delegate() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The user activity object’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/delegate
func (u_ UserActivity) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}

// The date after which the activity is no longer eligible for Handoff or indexing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/expirationDate
func (u_ UserActivity) ExpirationDate() NSDate {
	rv := objc.Send[NSDate](u_.ID, objc.Sel("expirationDate"))
	return rv
}


// SetExpirationDate sets the value of the expirationDate property.
// The date after which the activity is no longer eligible for Handoff or indexing.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/expirationDate
func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExpirationDate:"), value)
}

// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/externalMediaContentIdentifier
func (u_ UserActivity) ExternalMediaContentIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("externalMediaContentIdentifier"))
	return rv
}


// SetExternalMediaContentIdentifier sets the value of the externalMediaContentIdentifier property.
// A unique identifier from the app’s media content catalog for the currently displayed media item.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/externalMediaContentIdentifier
func (u_ UserActivity) SetExternalMediaContentIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExternalMediaContentIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (u_ UserActivity) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}

// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForHandoff
func (u_ UserActivity) EligibleForHandoff() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("eligibleForHandoff"))
	return rv
}


// SetEligibleForHandoff sets the value of the eligibleForHandoff property.
// A Boolean value that indicates whether the activity can be continued on another device using Handoff.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForHandoff
func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEligibleForHandoff:"), value)
}

// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPrediction
func (u_ UserActivity) EligibleForPrediction() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("eligibleForPrediction"))
	return rv
}


// SetEligibleForPrediction sets the value of the eligibleForPrediction property.
// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPrediction
func (u_ UserActivity) SetEligibleForPrediction(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEligibleForPrediction:"), value)
}

// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPublicIndexing
func (u_ UserActivity) EligibleForPublicIndexing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("eligibleForPublicIndexing"))
	return rv
}


// SetEligibleForPublicIndexing sets the value of the eligibleForPublicIndexing property.
// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPublicIndexing
func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEligibleForPublicIndexing:"), value)
}

// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForSearch
func (u_ UserActivity) EligibleForSearch() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("eligibleForSearch"))
	return rv
}


// SetEligibleForSearch sets the value of the eligibleForSearch property.
// A Boolean value that indicates whether the activity should be added to the on-device index.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForSearch
func (u_ UserActivity) SetEligibleForSearch(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEligibleForSearch:"), value)
}

// A set of localized keywords that can help users find the activity in search results.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/keywords
func (u_ UserActivity) Keywords() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("keywords"))
	return rv
}


// SetKeywords sets the value of the keywords property.
// A set of localized keywords that can help users find the activity in search results.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/keywords
func (u_ UserActivity) SetKeywords(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setKeywords:"), value)
}

// The NDEF message read by the system in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/ndefMessagePayload
func (u_ UserActivity) NdefMessagePayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ndefMessagePayload"))
	return rv
}

// A Boolean value that indicates whether the state of the activity needs to be updated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/needsSave
func (u_ UserActivity) NeedsSave() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("needsSave"))
	return rv
}


// SetNeedsSave sets the value of the needsSave property.
// A Boolean value that indicates whether the state of the activity needs to be updated.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/needsSave
func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNeedsSave:"), value)
}

// A value used to identify the user activity.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/persistentIdentifier
func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := objc.Send[UserActivityPersistentIdentifier](u_.ID, objc.Sel("persistentIdentifier"))
	return rv
}


// SetPersistentIdentifier sets the value of the persistentIdentifier property.
// A value used to identify the user activity.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/persistentIdentifier
func (u_ UserActivity) SetPersistentIdentifier(value IUserActivityPersistentIdentifier) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPersistentIdentifier:"), value)
}

// The URL of the webpage that linked to the webpage URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/referrerURL
func (u_ UserActivity) ReferrerURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("referrerURL"))
	return rv
}


// SetReferrerURL sets the value of the referrerURL property.
// The URL of the webpage that linked to the webpage URL.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/referrerURL
func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setReferrerURL:"), value)
}

// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/requiredUserInfoKeys
func (u_ UserActivity) RequiredUserInfoKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("requiredUserInfoKeys"))
	return rv
}


// SetRequiredUserInfoKeys sets the value of the requiredUserInfoKeys property.
// A set of keys that represent the minimal information about the activity that should be stored for later restoration.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/requiredUserInfoKeys
func (u_ UserActivity) SetRequiredUserInfoKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiredUserInfoKeys:"), value)
}

// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/shortcutAvailability
func (u_ UserActivity) ShortcutAvailability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("shortcutAvailability"))
	return rv
}


// SetShortcutAvailability sets the value of the shortcutAvailability property.
// A set of defined contexts in which an intent or activity might be relevant to a user.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/shortcutAvailability
func (u_ UserActivity) SetShortcutAvailability(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShortcutAvailability:"), value)
}

// A phrase suggested to the user when they create a shortcut.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/suggestedInvocationPhrase
func (u_ UserActivity) SuggestedInvocationPhrase() string {
	rv := objc.Send[string](u_.ID, objc.Sel("suggestedInvocationPhrase"))
	return rv
}


// SetSuggestedInvocationPhrase sets the value of the suggestedInvocationPhrase property.
// A phrase suggested to the user when they create a shortcut.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/suggestedInvocationPhrase
func (u_ UserActivity) SetSuggestedInvocationPhrase(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSuggestedInvocationPhrase:"), objc.String(value))
}

// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/supportsContinuationStreams
func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("supportsContinuationStreams"))
	return rv
}


// SetSupportsContinuationStreams sets the value of the supportsContinuationStreams property.
// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/supportsContinuationStreams
func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSupportsContinuationStreams:"), value)
}

// A string that identifies the user activity’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/targetContentIdentifier
func (u_ UserActivity) TargetContentIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}


// SetTargetContentIdentifier sets the value of the targetContentIdentifier property.
// A string that identifies the user activity’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/targetContentIdentifier
func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), objc.String(value))
}

// An optional, user-visible title for this activity, such as a document name or web page title.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/title
func (u_ UserActivity) Title() string {
	rv := objc.Send[string](u_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// An optional, user-visible title for this activity, such as a document name or web page title.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/title
func (u_ UserActivity) SetTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A dictionary containing app-specific state information needed to continue an activity on another device.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
func (u_ UserActivity) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// A dictionary containing app-specific state information needed to continue an activity on another device.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
func (u_ UserActivity) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}

// The URL of the webpage to load in a browser to continue the activity.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/webpageURL
func (u_ UserActivity) WebpageURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("webpageURL"))
	return rv
}


// SetWebpageURL sets the value of the webpageURL property.
// The URL of the webpage to load in a browser to continue the activity.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/webpageURL
func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWebpageURL:"), value)
}

// An activity for viewing your app’s channel guide.
//
// [Full Topic]: https://developer.apple.com/documentation/TVServices/TVUserActivityTypeBrowsingChannelGuide
func (u_ UserActivity) TVUserActivityTypeBrowsingChannelGuide() string {
	rv := objc.Send[string](u_.ID, objc.Sel("TVUserActivityTypeBrowsingChannelGuide"))
	return rv
}

// An object or value that specifies items to share.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) ActivityItemsConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}


// SetActivityItemsConfiguration sets the value of the activityItemsConfiguration property.
// An object or value that specifies items to share.

//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) SetActivityItemsConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}

// The identifier of an app entity that you associate with the user activity.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) AppEntityIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("appEntityIdentifier"))
	return rv
}


// SetAppEntityIdentifier sets the value of the appEntityIdentifier property.
// The identifier of an app entity that you associate with the user activity.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) SetAppEntityIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppEntityIdentifier:"), value)
}

// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForHandoff"))
	return rv
}


// SetIsEligibleForHandoff sets the value of the isEligibleForHandoff property.
// A Boolean value that indicates whether the activity can be continued on another device using Handoff.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) SetIsEligibleForHandoff(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForHandoff:"), value)
}

// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) IsEligibleForPrediction() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPrediction"))
	return rv
}


// SetIsEligibleForPrediction sets the value of the isEligibleForPrediction property.
// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) SetIsEligibleForPrediction(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPrediction:"), value)
}

// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPublicIndexing"))
	return rv
}


// SetIsEligibleForPublicIndexing sets the value of the isEligibleForPublicIndexing property.
// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) SetIsEligibleForPublicIndexing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPublicIndexing:"), value)
}

// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForSearch"))
	return rv
}


// SetIsEligibleForSearch sets the value of the isEligibleForSearch property.
// A Boolean value that indicates whether the activity should be added to the on-device index.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) SetIsEligibleForSearch(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForSearch:"), value)
}

// The user activity couldn’t be continued because a required connection wasn’t available.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) NSUserActivityConnectionUnavailableError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityConnectionUnavailableError"))
	return rv
}


// SetNSUserActivityConnectionUnavailableError sets the value of the NSUserActivityConnectionUnavailableError property.
// The user activity couldn’t be continued because a required connection wasn’t available.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) SetNSUserActivityConnectionUnavailableError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityConnectionUnavailableError:"), value)
}

// The end of the range of error codes reserved for user activity errors.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) NSUserActivityErrorMaximum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMaximum"))
	return rv
}


// SetNSUserActivityErrorMaximum sets the value of the NSUserActivityErrorMaximum property.
// The end of the range of error codes reserved for user activity errors.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMaximum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMaximum:"), value)
}

// The start of the range of error codes reserved for user activity errors.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) NSUserActivityErrorMinimum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMinimum"))
	return rv
}


// SetNSUserActivityErrorMinimum sets the value of the NSUserActivityErrorMinimum property.
// The start of the range of error codes reserved for user activity errors.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMinimum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMinimum:"), value)
}

// The data for the user activity wasn’t available.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) NSUserActivityHandoffFailedError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffFailedError"))
	return rv
}


// SetNSUserActivityHandoffFailedError sets the value of the NSUserActivityHandoffFailedError property.
// The data for the user activity wasn’t available.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffFailedError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffFailedError:"), value)
}

// The user info dictionary was too large to receive.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) NSUserActivityHandoffUserInfoTooLargeError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffUserInfoTooLargeError"))
	return rv
}


// SetNSUserActivityHandoffUserInfoTooLargeError sets the value of the NSUserActivityHandoffUserInfoTooLargeError property.
// The user info dictionary was too large to receive.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffUserInfoTooLargeError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffUserInfoTooLargeError:"), value)
}

// The remote application failed to send data within the specified time.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) NSUserActivityRemoteApplicationTimedOutError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityRemoteApplicationTimedOutError"))
	return rv
}


// SetNSUserActivityRemoteApplicationTimedOutError sets the value of the NSUserActivityRemoteApplicationTimedOutError property.
// The remote application failed to send data within the specified time.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) SetNSUserActivityRemoteApplicationTimedOutError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityRemoteApplicationTimedOutError:"), value)
}

// An activity that continues from Handoff or a universal link.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitytypebrowsingweb
func (u_ UserActivity) NSUserActivityTypeBrowsingWeb() string {
	rv := objc.Send[string](u_.ID, objc.Sel("NSUserActivityTypeBrowsingWeb"))
	return rv
}



