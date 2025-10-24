// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUserActivity */


/* debug [class_header]: Header for NSUserActivity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserActivity */
// An interface definition for the [UserActivity] class.
type IUserActivity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserActivity */
	// properties:
	TVUserActivityTypeBrowsingChannelGuide() IString
	ActivityItemsConfiguration() objectivec.IObject
	SetActivityItemsConfiguration(value objectivec.IObject)
	ActivityType() IString
	SetActivityType(value IString)
	AppClipActivationPayload() objectivec.IObject
	SetAppClipActivationPayload(value objectivec.IObject)
	AppEntityIdentifier() objectivec.IObject
	SetAppEntityIdentifier(value objectivec.IObject)
	ContextIdentifierPath() IString
	SetContextIdentifierPath(value IString)
	Delegate() objc.IObject /* cross-framework: UserActivityDelegate */
	SetDelegate(value objc.IObject /* cross-framework: UserActivityDelegate */)
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
	NdefMessagePayload() objectivec.IObject
	SetNdefMessagePayload(value objectivec.IObject)
	NeedsSave() bool
	SetNeedsSave(value bool)
	PersistentIdentifier() UserActivityPersistentIdentifier /* not a class type */
	SetPersistentIdentifier(value UserActivityPersistentIdentifier /* not a class type */)
	ReferrerURL() IURL
	SetReferrerURL(value IURL)
	RequiredUserInfoKeys() IString
	SetRequiredUserInfoKeys(value IString)
	ShortcutAvailability() objectivec.IObject
	SetShortcutAvailability(value objectivec.IObject)
	SuggestedInvocationPhrase() IString
	SetSuggestedInvocationPhrase(value IString)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	TargetContentIdentifier() IString
	SetTargetContentIdentifier(value IString)
	Title() IString
	SetTitle(value IString)
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserActivity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserActivity */
// Alloc allocates a new instance without initialization.
func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.Send[UserActivity](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserActivity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserActivity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserActivity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserActivity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserActivity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserActivity */

// An activity for viewing your app’s channel guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/TVServices/TVUserActivityTypeBrowsingChannelGuide
func (u_ UserActivity) TVUserActivityTypeBrowsingChannelGuide() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("TVUserActivityTypeBrowsingChannelGuide"))
	return rv
}/* debug [instance_properties/getter]: TVUserActivityTypeBrowsingChannelGuide */


// An object or value that specifies items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) ActivityItemsConfiguration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}/* debug [instance_properties/getter]: activityItemsConfiguration */


// An object or value that specifies items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u_ UserActivity) SetActivityItemsConfiguration(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}/* debug [instance_properties/setter]: activityItemsConfiguration */


// The user activity object’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/activitytype
func (u_ UserActivity) ActivityType() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("activityType"))
	return rv
}/* debug [instance_properties/getter]: activityType */


// The user activity object’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/activitytype
func (u_ UserActivity) SetActivityType(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivityType:"), value)
}/* debug [instance_properties/setter]: activityType */


// An object containing the payload information that launches an App Clip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appclipactivationpayload
func (u_ UserActivity) AppClipActivationPayload() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("appClipActivationPayload"))
	return rv
}/* debug [instance_properties/getter]: appClipActivationPayload */


// An object containing the payload information that launches an App Clip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appclipactivationpayload
func (u_ UserActivity) SetAppClipActivationPayload(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppClipActivationPayload:"), value)
}/* debug [instance_properties/setter]: appClipActivationPayload */


// The identifier of an app entity that you associate with the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) AppEntityIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("appEntityIdentifier"))
	return rv
}/* debug [instance_properties/getter]: appEntityIdentifier */


// The identifier of an app entity that you associate with the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u_ UserActivity) SetAppEntityIdentifier(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppEntityIdentifier:"), value)
}/* debug [instance_properties/setter]: appEntityIdentifier */


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contextidentifierpath
func (u_ UserActivity) ContextIdentifierPath() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("contextIdentifierPath"))
	return rv
}/* debug [instance_properties/getter]: contextIdentifierPath */


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/contextidentifierpath
func (u_ UserActivity) SetContextIdentifierPath(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContextIdentifierPath:"), value)
}/* debug [instance_properties/setter]: contextIdentifierPath */


// The user activity object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/delegate
func (u_ UserActivity) Delegate() objc.IObject /* cross-framework: UserActivityDelegate */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The user activity object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/delegate
func (u_ UserActivity) SetDelegate(value objc.IObject /* cross-framework: UserActivityDelegate */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The barcode that the system scanner passes in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/detectedbarcodedescriptor
func (u_ UserActivity) DetectedBarcodeDescriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("detectedBarcodeDescriptor"))
	return rv
}/* debug [instance_properties/getter]: detectedBarcodeDescriptor */


// The barcode that the system scanner passes in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/detectedbarcodedescriptor
func (u_ UserActivity) SetDetectedBarcodeDescriptor(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDetectedBarcodeDescriptor:"), value)
}/* debug [instance_properties/setter]: detectedBarcodeDescriptor */


// The date after which the activity is no longer eligible for Handoff or indexing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/expirationdate
func (u_ UserActivity) ExpirationDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("expirationDate"))
	return rv
}/* debug [instance_properties/getter]: expirationDate */


// The date after which the activity is no longer eligible for Handoff or indexing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/expirationdate
func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExpirationDate:"), value)
}/* debug [instance_properties/setter]: expirationDate */


// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/externalmediacontentidentifier
func (u_ UserActivity) ExternalMediaContentIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("externalMediaContentIdentifier"))
	return rv
}/* debug [instance_properties/getter]: externalMediaContentIdentifier */


// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/externalmediacontentidentifier
func (u_ UserActivity) SetExternalMediaContentIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExternalMediaContentIdentifier:"), value)
}/* debug [instance_properties/setter]: externalMediaContentIdentifier */


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/isclasskitdeeplink
func (u_ UserActivity) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}/* debug [instance_properties/getter]: isClassKitDeepLink */


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/isclasskitdeeplink
func (u_ UserActivity) SetIsClassKitDeepLink(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsClassKitDeepLink:"), value)
}/* debug [instance_properties/setter]: isClassKitDeepLink */


// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForHandoff"))
	return rv
}/* debug [instance_properties/getter]: isEligibleForHandoff */


// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforhandoff
func (u_ UserActivity) SetIsEligibleForHandoff(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForHandoff:"), value)
}/* debug [instance_properties/setter]: isEligibleForHandoff */


// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) IsEligibleForPrediction() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPrediction"))
	return rv
}/* debug [instance_properties/getter]: isEligibleForPrediction */


// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforprediction
func (u_ UserActivity) SetIsEligibleForPrediction(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPrediction:"), value)
}/* debug [instance_properties/setter]: isEligibleForPrediction */


// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForPublicIndexing"))
	return rv
}/* debug [instance_properties/getter]: isEligibleForPublicIndexing */


// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforpublicindexing
func (u_ UserActivity) SetIsEligibleForPublicIndexing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForPublicIndexing:"), value)
}/* debug [instance_properties/setter]: isEligibleForPublicIndexing */


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isEligibleForSearch"))
	return rv
}/* debug [instance_properties/getter]: isEligibleForSearch */


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/iseligibleforsearch
func (u_ UserActivity) SetIsEligibleForSearch(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsEligibleForSearch:"), value)
}/* debug [instance_properties/setter]: isEligibleForSearch */


// A set of localized keywords that can help users find the activity in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/keywords
func (u_ UserActivity) Keywords() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("keywords"))
	return rv
}/* debug [instance_properties/getter]: keywords */


// A set of localized keywords that can help users find the activity in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/keywords
func (u_ UserActivity) SetKeywords(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setKeywords:"), value)
}/* debug [instance_properties/setter]: keywords */


// The NDEF message read by the system in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/ndefmessagepayload
func (u_ UserActivity) NdefMessagePayload() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("ndefMessagePayload"))
	return rv
}/* debug [instance_properties/getter]: ndefMessagePayload */


// The NDEF message read by the system in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/ndefmessagepayload
func (u_ UserActivity) SetNdefMessagePayload(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNdefMessagePayload:"), value)
}/* debug [instance_properties/setter]: ndefMessagePayload */


// A Boolean value that indicates whether the state of the activity needs to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/needssave
func (u_ UserActivity) NeedsSave() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("needsSave"))
	return rv
}/* debug [instance_properties/getter]: needsSave */


// A Boolean value that indicates whether the state of the activity needs to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/needssave
func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNeedsSave:"), value)
}/* debug [instance_properties/setter]: needsSave */


// A value used to identify the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/persistentidentifier
func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier /* not a class type */ {
	rv := objc.Send[UserActivityPersistentIdentifier](u_.ID, objc.Sel("persistentIdentifier"))
	return rv
}/* debug [instance_properties/getter]: persistentIdentifier */


// A value used to identify the user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/persistentidentifier
func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPersistentIdentifier:"), value)
}/* debug [instance_properties/setter]: persistentIdentifier */


// The URL of the webpage that linked to the webpage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/referrerurl
func (u_ UserActivity) ReferrerURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("referrerURL"))
	return rv
}/* debug [instance_properties/getter]: referrerURL */


// The URL of the webpage that linked to the webpage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/referrerurl
func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setReferrerURL:"), value)
}/* debug [instance_properties/setter]: referrerURL */


// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/requireduserinfokeys
func (u_ UserActivity) RequiredUserInfoKeys() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("requiredUserInfoKeys"))
	return rv
}/* debug [instance_properties/getter]: requiredUserInfoKeys */


// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/requireduserinfokeys
func (u_ UserActivity) SetRequiredUserInfoKeys(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiredUserInfoKeys:"), value)
}/* debug [instance_properties/setter]: requiredUserInfoKeys */


// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/shortcutavailability
func (u_ UserActivity) ShortcutAvailability() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("shortcutAvailability"))
	return rv
}/* debug [instance_properties/getter]: shortcutAvailability */


// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/shortcutavailability
func (u_ UserActivity) SetShortcutAvailability(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShortcutAvailability:"), value)
}/* debug [instance_properties/setter]: shortcutAvailability */


// A phrase suggested to the user when they create a shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/suggestedinvocationphrase
func (u_ UserActivity) SuggestedInvocationPhrase() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("suggestedInvocationPhrase"))
	return rv
}/* debug [instance_properties/getter]: suggestedInvocationPhrase */


// A phrase suggested to the user when they create a shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/suggestedinvocationphrase
func (u_ UserActivity) SetSuggestedInvocationPhrase(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSuggestedInvocationPhrase:"), value)
}/* debug [instance_properties/setter]: suggestedInvocationPhrase */


// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/supportscontinuationstreams
func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("supportsContinuationStreams"))
	return rv
}/* debug [instance_properties/getter]: supportsContinuationStreams */


// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/supportscontinuationstreams
func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSupportsContinuationStreams:"), value)
}/* debug [instance_properties/setter]: supportsContinuationStreams */


// A string that identifies the user activity’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/targetcontentidentifier
func (u_ UserActivity) TargetContentIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}/* debug [instance_properties/getter]: targetContentIdentifier */


// A string that identifies the user activity’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/targetcontentidentifier
func (u_ UserActivity) SetTargetContentIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), value)
}/* debug [instance_properties/setter]: targetContentIdentifier */


// An optional, user-visible title for this activity, such as a document name or web page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/title
func (u_ UserActivity) Title() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// An optional, user-visible title for this activity, such as a document name or web page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/title
func (u_ UserActivity) SetTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A dictionary containing app-specific state information needed to continue an activity on another device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/userinfo
func (u_ UserActivity) UserInfo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A dictionary containing app-specific state information needed to continue an activity on another device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/userinfo
func (u_ UserActivity) SetUserInfo(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */


// The URL of the webpage to load in a browser to continue the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/webpageurl
func (u_ UserActivity) WebpageURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("webpageURL"))
	return rv
}/* debug [instance_properties/getter]: webpageURL */


// The URL of the webpage to load in a browser to continue the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/webpageurl
func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWebpageURL:"), value)
}/* debug [instance_properties/setter]: webpageURL */


// The user activity couldn’t be continued because a required connection wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) NSUserActivityConnectionUnavailableError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityConnectionUnavailableError"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityConnectionUnavailableError */


// The user activity couldn’t be continued because a required connection wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u_ UserActivity) SetNSUserActivityConnectionUnavailableError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityConnectionUnavailableError:"), value)
}/* debug [instance_properties/setter]: NSUserActivityConnectionUnavailableError */


// The end of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) NSUserActivityErrorMaximum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityErrorMaximum */


// The end of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMaximum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSUserActivityErrorMaximum */


// The start of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) NSUserActivityErrorMinimum() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityErrorMinimum */


// The start of the range of error codes reserved for user activity errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u_ UserActivity) SetNSUserActivityErrorMinimum(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSUserActivityErrorMinimum */


// The data for the user activity wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) NSUserActivityHandoffFailedError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffFailedError"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityHandoffFailedError */


// The data for the user activity wasn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffFailedError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffFailedError:"), value)
}/* debug [instance_properties/setter]: NSUserActivityHandoffFailedError */


// The user info dictionary was too large to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) NSUserActivityHandoffUserInfoTooLargeError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityHandoffUserInfoTooLargeError"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityHandoffUserInfoTooLargeError */


// The user info dictionary was too large to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u_ UserActivity) SetNSUserActivityHandoffUserInfoTooLargeError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityHandoffUserInfoTooLargeError:"), value)
}/* debug [instance_properties/setter]: NSUserActivityHandoffUserInfoTooLargeError */


// The remote application failed to send data within the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) NSUserActivityRemoteApplicationTimedOutError() int {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUserActivityRemoteApplicationTimedOutError"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityRemoteApplicationTimedOutError */


// The remote application failed to send data within the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u_ UserActivity) SetNSUserActivityRemoteApplicationTimedOutError(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUserActivityRemoteApplicationTimedOutError:"), value)
}/* debug [instance_properties/setter]: NSUserActivityRemoteApplicationTimedOutError */


// An activity that continues from Handoff or a universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitytypebrowsingweb
func (u_ UserActivity) NSUserActivityTypeBrowsingWeb() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUserActivityTypeBrowsingWeb"))
	return rv
}/* debug [instance_properties/getter]: NSUserActivityTypeBrowsingWeb */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserActivity */



