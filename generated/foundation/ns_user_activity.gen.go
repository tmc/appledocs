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
	// properties:
	Interaction() objectivec.IObject
	TVUserActivityTypeBrowsingChannelGuide() IString
	ActivityItemsConfiguration() objectivec.IObject
	SetActivityItemsConfiguration(value objectivec.IObject)
	ActivityType() IString
	SetActivityType(value IString)
	AppClipActivationPayload() unsafe.Pointer
	SetAppClipActivationPayload(value unsafe.Pointer)
	AppEntityIdentifier() unsafe.Pointer
	SetAppEntityIdentifier(value unsafe.Pointer)
	ContentAttributeSet() objectivec.IObject
	SetContentAttributeSet(value objectivec.IObject)
	ContextIdentifierPath() IString
	SetContextIdentifierPath(value IString)
	Delegate() UserActivityDelegate /* not a class type */
	SetDelegate(value UserActivityDelegate /* not a class type */)
	DetectedBarcodeDescriptor() objectivec.IObject
	SetDetectedBarcodeDescriptor(value objectivec.IObject)
	ExpirationDate() IDate
	SetExpirationDate(value IDate)
	ExternalMediaContentIdentifier() IString
	SetExternalMediaContentIdentifier(value IString)
	IsClassKitDeepLink() bool
	SetIsClassKitDeepLink(value bool)
	IsEligibleForHandoff() bool
	SetIsEligibleForHandoff(value bool)
	IsEligibleForPrediction() bool
	SetIsEligibleForPrediction(value bool)
	IsEligibleForPublicIndexing() bool
	SetIsEligibleForPublicIndexing(value bool)
	IsEligibleForSearch() bool
	SetIsEligibleForSearch(value bool)
	Keywords() IString
	SetKeywords(value IString)
	MapItem() objectivec.IObject
	SetMapItem(value objectivec.IObject)
	NdefMessagePayload() unsafe.Pointer
	SetNdefMessagePayload(value unsafe.Pointer)
	NeedsSave() bool
	SetNeedsSave(value bool)
	PersistentIdentifier() UserActivityPersistentIdentifier /* not a class type */
	SetPersistentIdentifier(value UserActivityPersistentIdentifier /* not a class type */)
	ReferrerURL() IURL
	SetReferrerURL(value IURL)
	RequiredUserInfoKeys() IString
	SetRequiredUserInfoKeys(value IString)
	ShortcutAvailability() unsafe.Pointer
	SetShortcutAvailability(value unsafe.Pointer)
	SuggestedInvocationPhrase() IString
	SetSuggestedInvocationPhrase(value IString)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	TargetContentIdentifier() IString
	SetTargetContentIdentifier(value IString)
	Title() IString
	SetTitle(value IString)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	WebpageURL() IURL
	SetWebpageURL(value IURL)
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
	NSUserActivityTypeBrowsingWeb() IString
	// methods:
}

// A representation of the state of your app at a moment in time.
//
// An object provides a lightweight way to capture the state of your app and put it to use later. Create this object to capture information about what a person was doing, such as viewing app content, editing a document, viewing a web page, or watching a video. When the system launches your app and an activity object is available, your app can use the information in that object to restore itself to an appropriate state. Spotlight also uses these objects to improve search results for people. To allow people to continue an activity on another device, see .


// A representation of the state of your app at a moment in time.
//
// [Full Topic]
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



// The SiriKit interaction object to use when configuring your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/interaction
func (u_ UserActivity) Interaction() objectivec.IObject {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("interaction"))
	return rv
}


// An activity for viewing your app’s channel guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/TVServices/TVUserActivityTypeBrowsingChannelGuide
func (u_ UserActivity) TVUserActivityTypeBrowsingChannelGuide() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("TVUserActivityTypeBrowsingChannelGuide"))
	return rv
}


// An object or value that specifies items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) ActivityItemsConfiguration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}


// An object or value that specifies items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) SetActivityItemsConfiguration(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}


// The user activity object’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/activitytype
func (u_ UserActivity) ActivityType() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("activityType"))
	return rv
}


// The user activity object’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/activitytype
func (u_ UserActivity) SetActivityType(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivityType:"), value)
}


// An object containing the payload information that launches an App Clip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appclipactivationpayload
func (u_ UserActivity) AppClipActivationPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("appClipActivationPayload"))
	return rv
}


// An object containing the payload information that launches an App Clip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appclipactivationpayload
func (u_ UserActivity) SetAppClipActivationPayload(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppClipActivationPayload:"), value)
}


// The identifier of an app entity that you associate with the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) AppEntityIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("appEntityIdentifier"))
	return rv
}


// The identifier of an app entity that you associate with the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) SetAppEntityIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppEntityIdentifier:"), value)
}


// A set of properties that describe the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contentattributeset
func (u_ UserActivity) ContentAttributeSet() objectivec.IObject {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("contentAttributeSet"))
	return rv
}


// A set of properties that describe the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contentattributeset
func (u_ UserActivity) SetContentAttributeSet(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContentAttributeSet:"), value)
}


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contextidentifierpath
func (u_ UserActivity) ContextIdentifierPath() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("contextIdentifierPath"))
	return rv
}


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contextidentifierpath
func (u_ UserActivity) SetContextIdentifierPath(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContextIdentifierPath:"), value)
}


// The user activity object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/delegate
func (u_ UserActivity) Delegate() UserActivityDelegate /* not a class type */ {
	rv := objc.Send[UserActivityDelegate](u_.ID, objc.Sel("delegate"))
	return rv
}


// The user activity object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/delegate
func (u_ UserActivity) SetDelegate(value UserActivityDelegate /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}


// The barcode that the system scanner passes in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/detectedbarcodedescriptor
func (u_ UserActivity) DetectedBarcodeDescriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("detectedBarcodeDescriptor"))
	return rv
}


// The barcode that the system scanner passes in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/detectedbarcodedescriptor
func (u_ UserActivity) SetDetectedBarcodeDescriptor(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDetectedBarcodeDescriptor:"), value)
}


// The date after which the activity is no longer eligible for Handoff or indexing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/expirationdate
func (u_ UserActivity) ExpirationDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("expirationDate"))
	return rv
}


// The date after which the activity is no longer eligible for Handoff or indexing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/expirationdate
func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExpirationDate:"), value)
}


// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/externalmediacontentidentifier
func (u_ UserActivity) ExternalMediaContentIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("externalMediaContentIdentifier"))
	return rv
}


// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/externalmediacontentidentifier
func (u_ UserActivity) SetExternalMediaContentIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExternalMediaContentIdentifier:"), value)
}


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/isclasskitdeeplink
func (u_ UserActivity) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/isclasskitdeeplink
func (u_ UserActivity) SetIsClassKitDeepLink(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsClassKitDeepLink:"), value)
}


// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForHandoff"))
	return rv
}


// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) SetIsEligibleForHandoff(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForHandoff:"), value)
}


// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) IsEligibleForPrediction() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPrediction"))
	return rv
}


// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) SetIsEligibleForPrediction(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPrediction:"), value)
}


// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPublicIndexing"))
	return rv
}


// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) SetIsEligibleForPublicIndexing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPublicIndexing:"), value)
}


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForSearch"))
	return rv
}


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) SetIsEligibleForSearch(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForSearch:"), value)
}


// A set of localized keywords that can help users find the activity in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/keywords
func (u_ UserActivity) Keywords() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("keywords"))
	return rv
}


// A set of localized keywords that can help users find the activity in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/keywords
func (u_ UserActivity) SetKeywords(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setKeywords:"), value)
}


// Attaches the specified map item to a user activity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/mapitem
func (u_ UserActivity) MapItem() objectivec.IObject {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("mapItem"))
	return rv
}


// Attaches the specified map item to a user activity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/mapitem
func (u_ UserActivity) SetMapItem(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMapItem:"), value)
}


// The NDEF message read by the system in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/ndefmessagepayload
func (u_ UserActivity) NdefMessagePayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ndefMessagePayload"))
	return rv
}


// The NDEF message read by the system in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/ndefmessagepayload
func (u_ UserActivity) SetNdefMessagePayload(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNdefMessagePayload:"), value)
}


// A Boolean value that indicates whether the state of the activity needs to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/needssave
func (u_ UserActivity) NeedsSave() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("needsSave"))
	return rv
}


// A Boolean value that indicates whether the state of the activity needs to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/needssave
func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNeedsSave:"), value)
}


// A value used to identify the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/persistentidentifier
func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier /* not a class type */ {
	rv := objc.Send[UserActivityPersistentIdentifier](u_.ID, objc.Sel("persistentIdentifier"))
	return rv
}


// A value used to identify the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/persistentidentifier
func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPersistentIdentifier:"), value)
}


// The URL of the webpage that linked to the webpage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/referrerurl
func (u_ UserActivity) ReferrerURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("referrerURL"))
	return rv
}


// The URL of the webpage that linked to the webpage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/referrerurl
func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setReferrerURL:"), value)
}


// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/requireduserinfokeys
func (u_ UserActivity) RequiredUserInfoKeys() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("requiredUserInfoKeys"))
	return rv
}


// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/requireduserinfokeys
func (u_ UserActivity) SetRequiredUserInfoKeys(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiredUserInfoKeys:"), value)
}


// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/shortcutavailability
func (u_ UserActivity) ShortcutAvailability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("shortcutAvailability"))
	return rv
}


// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/shortcutavailability
func (u_ UserActivity) SetShortcutAvailability(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShortcutAvailability:"), value)
}


// A phrase suggested to the user when they create a shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/suggestedinvocationphrase
func (u_ UserActivity) SuggestedInvocationPhrase() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("suggestedInvocationPhrase"))
	return rv
}


// A phrase suggested to the user when they create a shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/suggestedinvocationphrase
func (u_ UserActivity) SetSuggestedInvocationPhrase(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSuggestedInvocationPhrase:"), value)
}


// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/supportscontinuationstreams
func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("supportsContinuationStreams"))
	return rv
}


// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/supportscontinuationstreams
func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSupportsContinuationStreams:"), value)
}


// A string that identifies the user activity’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/targetcontentidentifier
func (u_ UserActivity) TargetContentIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}


// A string that identifies the user activity’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/targetcontentidentifier
func (u_ UserActivity) SetTargetContentIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), value)
}


// An optional, user-visible title for this activity, such as a document name or web page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/title
func (u_ UserActivity) Title() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("title"))
	return rv
}


// An optional, user-visible title for this activity, such as a document name or web page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/title
func (u_ UserActivity) SetTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}


// A dictionary containing app-specific state information needed to continue an activity on another device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/userinfo
func (u_ UserActivity) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("userInfo"))
	return rv
}


// A dictionary containing app-specific state information needed to continue an activity on another device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/userinfo
func (u_ UserActivity) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}


// The URL of the webpage to load in a browser to continue the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/webpageurl
func (u_ UserActivity) WebpageURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("webpageURL"))
	return rv
}


// The URL of the webpage to load in a browser to continue the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/webpageurl
func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWebpageURL:"), value)
}


// The user activity couldn’t be continued because a required connection wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) NSUserActivityConnectionUnavailableError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityConnectionUnavailableError"))
	return rv
}


// The user activity couldn’t be continued because a required connection wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) SetNSUserActivityConnectionUnavailableError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityConnectionUnavailableError:"), value)
}


// The end of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) NSUserActivityErrorMaximum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMaximum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMaximum:"), value)
}


// The start of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) NSUserActivityErrorMinimum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMinimum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMinimum:"), value)
}


// The data for the user activity wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) NSUserActivityHandoffFailedError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffFailedError"))
	return rv
}


// The data for the user activity wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffFailedError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffFailedError:"), value)
}


// The user info dictionary was too large to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) NSUserActivityHandoffUserInfoTooLargeError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffUserInfoTooLargeError"))
	return rv
}


// The user info dictionary was too large to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffUserInfoTooLargeError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffUserInfoTooLargeError:"), value)
}


// The remote application failed to send data within the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) NSUserActivityRemoteApplicationTimedOutError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityRemoteApplicationTimedOutError"))
	return rv
}


// The remote application failed to send data within the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) SetNSUserActivityRemoteApplicationTimedOutError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityRemoteApplicationTimedOutError:"), value)
}


// An activity that continues from Handoff or a universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitytypebrowsingweb
func (u_ UserActivity) NSUserActivityTypeBrowsingWeb() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUserActivityTypeBrowsingWeb"))
	return rv
}



