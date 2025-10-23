// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNSocialProfile] class.
var (
	CNSocialProfileClass     _CNSocialProfileClass
	CNSocialProfileClassOnce sync.Once
)

func getCNSocialProfileClass() _CNSocialProfileClass {
	CNSocialProfileClassOnce.Do(func() {
		CNSocialProfileClass = _CNSocialProfileClass{objc.GetClass("CNSocialProfile")}
	})
	return CNSocialProfileClass
}

type _CNSocialProfileClass struct {
	class objc.Class
}

// An interface definition for the [CNSocialProfile] class.
type ICNSocialProfile interface {
	objectivec.IObject
	// properties:
	Service() string /* primitive/slice/pointer. */
	SetService(value string /* primitive/slice/pointer. */)
	UrlString() string /* primitive/slice/pointer. */
	SetUrlString(value string /* primitive/slice/pointer. */)
	UserIdentifier() string /* primitive/slice/pointer. */
	SetUserIdentifier(value string /* primitive/slice/pointer. */)
	Username() string /* primitive/slice/pointer. */
	SetUsername(value string /* primitive/slice/pointer. */)
	CNSocialProfileServiceFacebook() string /* primitive/slice/pointer. */
	CNSocialProfileServiceFlickr() string /* primitive/slice/pointer. */
	CNSocialProfileServiceGameCenter() string /* primitive/slice/pointer. */
	CNSocialProfileServiceKey() string /* primitive/slice/pointer. */
	CNSocialProfileServiceLinkedIn() string /* primitive/slice/pointer. */
	CNSocialProfileServiceMySpace() string /* primitive/slice/pointer. */
	CNSocialProfileServiceSinaWeibo() string /* primitive/slice/pointer. */
	CNSocialProfileServiceTencentWeibo() string /* primitive/slice/pointer. */
	CNSocialProfileServiceTwitter() string /* primitive/slice/pointer. */
	CNSocialProfileServiceYelp() string /* primitive/slice/pointer. */
	CNSocialProfileURLStringKey() string /* primitive/slice/pointer. */
	CNSocialProfileUserIdentifierKey() string /* primitive/slice/pointer. */
	CNSocialProfileUsernameKey() string /* primitive/slice/pointer. */
	// methods:
}

// An immutable object that represents one of the user’s social profiles.
//
// Some social profile services, such as Facebook and Twitter, are predefined in this class. You can also specify your own social profile service with the method. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents one of the user’s social profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile
type CNSocialProfile struct {
	objectivec.Object
}

// CNSocialProfileFrom constructs a [CNSocialProfile] from an unsafe.Pointer.
//
// An immutable object that represents one of the user’s social profiles.
func CNSocialProfileFrom(ptr unsafe.Pointer) CNSocialProfile {
	return CNSocialProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNSocialProfileClass) Alloc() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNSocialProfileClass) New() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNSocialProfile) Init() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNSocialProfile) Autorelease() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNSocialProfile creates a new CNSocialProfile instance.
func NewCNSocialProfile() CNSocialProfile {
	return getCNSocialProfileClass().New()
}



// Returns the localized name of the property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/localizedString(forKey:)
func (cc _CNSocialProfileClass) LocalizedStringForKey(key string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), objc.String(key))
	return rv
}


// The social profile’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/service
func (c_ CNSocialProfile) Service() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("service"))
	return rv
}


// The social profile’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/service
func (c_ CNSocialProfile) SetService(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setService:"), objc.String(value))
}


// The URL associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/urlstring
func (c_ CNSocialProfile) UrlString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("urlString"))
	return rv
}


// The URL associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/urlstring
func (c_ CNSocialProfile) SetUrlString(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrlString:"), objc.String(value))
}


// The service’s user identifier associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/useridentifier
func (c_ CNSocialProfile) UserIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("userIdentifier"))
	return rv
}


// The service’s user identifier associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/useridentifier
func (c_ CNSocialProfile) SetUserIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentifier:"), objc.String(value))
}


// The user name for the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/username
func (c_ CNSocialProfile) Username() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("username"))
	return rv
}


// The user name for the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/username
func (c_ CNSocialProfile) SetUsername(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsername:"), objc.String(value))
}


// The Facebook social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicefacebook
func (c_ CNSocialProfile) CNSocialProfileServiceFacebook() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceFacebook"))
	return rv
}


// The Flickr social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileserviceflickr
func (c_ CNSocialProfile) CNSocialProfileServiceFlickr() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceFlickr"))
	return rv
}


// The Game Center social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicegamecenter
func (c_ CNSocialProfile) CNSocialProfileServiceGameCenter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceGameCenter"))
	return rv
}


// The social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicekey
func (c_ CNSocialProfile) CNSocialProfileServiceKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceKey"))
	return rv
}


// The LinkedIn social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicelinkedin
func (c_ CNSocialProfile) CNSocialProfileServiceLinkedIn() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceLinkedIn"))
	return rv
}


// The MySpace social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicemyspace
func (c_ CNSocialProfile) CNSocialProfileServiceMySpace() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceMySpace"))
	return rv
}


// The Sina Weibo social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicesinaweibo
func (c_ CNSocialProfile) CNSocialProfileServiceSinaWeibo() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceSinaWeibo"))
	return rv
}


// The Tencent Weibo social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicetencentweibo
func (c_ CNSocialProfile) CNSocialProfileServiceTencentWeibo() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceTencentWeibo"))
	return rv
}


// The Twitter social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicetwitter
func (c_ CNSocialProfile) CNSocialProfileServiceTwitter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceTwitter"))
	return rv
}


// The Yelp social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileserviceyelp
func (c_ CNSocialProfile) CNSocialProfileServiceYelp() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileServiceYelp"))
	return rv
}


// The social profile URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileurlstringkey
func (c_ CNSocialProfile) CNSocialProfileURLStringKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileURLStringKey"))
	return rv
}


// The social profile user identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileuseridentifierkey
func (c_ CNSocialProfile) CNSocialProfileUserIdentifierKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileUserIdentifierKey"))
	return rv
}


// The social profile user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileusernamekey
func (c_ CNSocialProfile) CNSocialProfileUsernameKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNSocialProfileUsernameKey"))
	return rv
}




