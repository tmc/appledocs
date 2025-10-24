// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNSocialProfile */


/* debug [class_header]: Header for CNSocialProfile */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNSocialProfile */
// An interface definition for the [CNSocialProfile] class.
type ICNSocialProfile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNSocialProfile */
	// properties:
	Service() objc.IObject /* cross-framework: NSString */
	UrlString() objc.IObject /* cross-framework: NSString */
	UserIdentifier() objc.IObject /* cross-framework: NSString */
	Username() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceFacebook() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceFlickr() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceGameCenter() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceKey() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceLinkedIn() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceMySpace() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceSinaWeibo() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceTencentWeibo() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceTwitter() objc.IObject /* cross-framework: NSString */
	CNSocialProfileServiceYelp() objc.IObject /* cross-framework: NSString */
	CNSocialProfileURLStringKey() objc.IObject /* cross-framework: NSString */
	CNSocialProfileUserIdentifierKey() objc.IObject /* cross-framework: NSString */
	CNSocialProfileUsernameKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNSocialProfile */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNSocialProfile */
// Alloc allocates a new instance without initialization.
func (cc _CNSocialProfileClass) Alloc() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNSocialProfile */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNSocialProfile */

// Initializes a new social profile object with the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/init(urlString:username:userIdentifier:service:)
func NewCNSocialProfileWithUrlStringUsernameUserIdentifierService(urlString objc.IObject /* cross-framework: NSString */, username objc.IObject /* cross-framework: NSString */, userIdentifier objc.IObject /* cross-framework: NSString */, service objc.IObject /* cross-framework: NSString */) CNSocialProfile {
	instance := getCNSocialProfileClass().Alloc()
	rv := objc.Send[CNSocialProfile](instance.ID, objc.Sel("initWithUrlString:username:userIdentifier:service:"), urlString, username, userIdentifier, service)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNSocialProfileWithUrlStringUsernameUserIdentifierService */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNSocialProfile */

// Returns the localized name of the property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/localizedString(forKey:)
func (cc _CNSocialProfileClass) LocalizedStringForKey(key objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForKey) */


// Returns the localized name of the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/localizedString(forService:)
func (cc _CNSocialProfileClass) LocalizedStringForService(service objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForService:"), service)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForService) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNSocialProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNSocialProfile */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNSocialProfile */

// The social profile’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/service
func (c_ CNSocialProfile) Service() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("service"))
	return rv
}/* debug [instance_properties/getter]: service */


// The URL associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/urlString
func (c_ CNSocialProfile) UrlString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("urlString"))
	return rv
}/* debug [instance_properties/getter]: urlString */


// The service’s user identifier associated with the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/userIdentifier
func (c_ CNSocialProfile) UserIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("userIdentifier"))
	return rv
}/* debug [instance_properties/getter]: userIdentifier */


// The user name for the social profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/username
func (c_ CNSocialProfile) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("username"))
	return rv
}/* debug [instance_properties/getter]: username */


// The Facebook social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicefacebook
func (c_ CNSocialProfile) CNSocialProfileServiceFacebook() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceFacebook"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceFacebook */


// The Flickr social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileserviceflickr
func (c_ CNSocialProfile) CNSocialProfileServiceFlickr() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceFlickr"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceFlickr */


// The Game Center social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicegamecenter
func (c_ CNSocialProfile) CNSocialProfileServiceGameCenter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceGameCenter"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceGameCenter */


// The social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicekey
func (c_ CNSocialProfile) CNSocialProfileServiceKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceKey"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceKey */


// The LinkedIn social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicelinkedin
func (c_ CNSocialProfile) CNSocialProfileServiceLinkedIn() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceLinkedIn"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceLinkedIn */


// The MySpace social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicemyspace
func (c_ CNSocialProfile) CNSocialProfileServiceMySpace() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceMySpace"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceMySpace */


// The Sina Weibo social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicesinaweibo
func (c_ CNSocialProfile) CNSocialProfileServiceSinaWeibo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceSinaWeibo"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceSinaWeibo */


// The Tencent Weibo social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicetencentweibo
func (c_ CNSocialProfile) CNSocialProfileServiceTencentWeibo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceTencentWeibo"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceTencentWeibo */


// The Twitter social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileservicetwitter
func (c_ CNSocialProfile) CNSocialProfileServiceTwitter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceTwitter"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceTwitter */


// The Yelp social profile service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileserviceyelp
func (c_ CNSocialProfile) CNSocialProfileServiceYelp() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileServiceYelp"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileServiceYelp */


// The social profile URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileurlstringkey
func (c_ CNSocialProfile) CNSocialProfileURLStringKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileURLStringKey"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileURLStringKey */


// The social profile user identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileuseridentifierkey
func (c_ CNSocialProfile) CNSocialProfileUserIdentifierKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileUserIdentifierKey"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileUserIdentifierKey */


// The social profile user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofileusernamekey
func (c_ CNSocialProfile) CNSocialProfileUsernameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNSocialProfileUsernameKey"))
	return rv
}/* debug [instance_properties/getter]: CNSocialProfileUsernameKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNSocialProfile */


