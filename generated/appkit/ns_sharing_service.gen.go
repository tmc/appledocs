// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSharingService */


/* debug [class_header]: Header for NSSharingService */
// The class instance for the [SharingService] class.
var (
	SharingServiceClass     _SharingServiceClass
	SharingServiceClassOnce sync.Once
)

func getSharingServiceClass() _SharingServiceClass {
	SharingServiceClassOnce.Do(func() {
		SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}
	})
	return SharingServiceClass
}

type _SharingServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SharingService */
// An interface definition for the [SharingService] class.
type ISharingService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SharingService */
	// properties:
	AccountName() objc.IObject /* cross-framework: NSString */
	AlternateImage() IImage
	AttachmentFileURLs() []foundation.URL
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Image() IImage
	MenuItemTitle() objc.IObject /* cross-framework: NSString */
	SetMenuItemTitle(value objc.IObject /* cross-framework: NSString */)
	MessageBody() objc.IObject /* cross-framework: NSString */
	PermanentLink() objc.IObject /* cross-framework: NSURL */
	Recipients() []string
	SetRecipients(value []string)
	Subject() objc.IObject /* cross-framework: NSString */
	SetSubject(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SharingService */
	// methods:
	CanPerformWithItems(items objc.IObject /* cross-framework: NSArray */) bool
	PerformWithItems(items objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SharingService */
// Alloc allocates a new instance without initialization.
func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharingServiceClass) New() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingService) Init() SharingService {
	rv := objc.Send[SharingService](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingService) Autorelease() SharingService {
	rv := objc.Send[SharingService](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingService creates a new SharingService instance.
func NewSharingService() SharingService {
	return getSharingServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SharingService */
// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
//
// An object provides a consistent user experience for sharing items— objects, objects, objects, video (through file URLs), of any object that implements the protocol—in macOS. For any item or group of items, the displays a sheet with the content to share. A sharing service can create a post on a social network like Twitter or Facebook, send a message by email or iMessage, upload videos to viewing services, or send a file using AirDrop. You can use objects directly in your app. The following example shows how to create a button that shares content directly to a social media service.


// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService
type SharingService struct {
	objectivec.Object
}

// SharingServiceFrom constructs a [SharingService] from an unsafe.Pointer.
//
// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
func SharingServiceFrom(ptr unsafe.Pointer) SharingService {
	return SharingService{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SharingService */

// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func NewSharingServiceNamed(serviceName SharingServiceName /* typedef */) SharingService {
	rv := objc.Send[SharingService](objc.ID(getSharingServiceClass().class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}/* debug [class_init_methods/constructor]: NewSharingServiceNamed */


// Creates a custom sharing service object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(title:image:alternateImage:handler:)
func NewSharingServiceWithTitleImageAlternateImageHandler(title objc.IObject /* cross-framework: NSString */, image IImage, alternateImage IImage, block unsafe.Pointer) SharingService {
	instance := getSharingServiceClass().Alloc()
	rv := objc.Send[SharingService](instance.ID, objc.Sel("initWithTitle:image:alternateImage:handler:"), title, image, alternateImage, block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSharingServiceWithTitleImageAlternateImageHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SharingService */

// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName /* typedef */) ISharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharingServiceNamed) */


// Returns a list of sharing services which could share all the provided items together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items objc.IObject /* cross-framework: NSArray */) []SharingService {
	rv := objc.Send[[]SharingService](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharingServicesForItems) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SharingService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SharingService */

// Returns whether the service can share all the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/canPerform(withItems:)
func (s_ SharingService) CanPerformWithItems(items objc.IObject /* cross-framework: NSArray */) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canPerformWithItems:"), items)
	return rv
}/* debug [instance_methods/method]: CanPerformWithItems */


// Manually performs the service on the provided items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/perform(withItems:)
func (s_ SharingService) PerformWithItems(items objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("performWithItems:"), items)
}/* debug [instance_methods/method]: PerformWithItems */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SharingService */

// The account name used for posting on Twitter or Sina Weibo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/accountName
func (s_ SharingService) AccountName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("accountName"))
	return rv
}/* debug [instance_properties/getter]: accountName */


// The alternate image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/alternateImage
func (s_ SharingService) AlternateImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}/* debug [instance_properties/getter]: alternateImage */


// An array of NSURL objects representing the files that were shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/attachmentFileURLs
func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](s_.ID, objc.Sel("attachmentFileURLs"))
	return rv
}/* debug [instance_properties/getter]: attachmentFileURLs */


// Specifies the delegate of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/delegate
func (s_ SharingService) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the delegate of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/delegate
func (s_ SharingService) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The primary image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/image
func (s_ SharingService) Image() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/menuItemTitle
func (s_ SharingService) MenuItemTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("menuItemTitle"))
	return rv
}/* debug [instance_properties/getter]: menuItemTitle */


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/menuItemTitle
func (s_ SharingService) SetMenuItemTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenuItemTitle:"), value)
}/* debug [instance_properties/setter]: menuItemTitle */


// The message body as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/messageBody
func (s_ SharingService) MessageBody() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("messageBody"))
	return rv
}/* debug [instance_properties/getter]: messageBody */


// A permanent URL (permalink) that your app can use to access the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/permanentLink
func (s_ SharingService) PermanentLink() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("permanentLink"))
	return rv
}/* debug [instance_properties/getter]: permanentLink */


// An array containing the user handles of the desired recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/recipients
func (s_ SharingService) Recipients() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("recipients"))
	return rv
}/* debug [instance_properties/getter]: recipients */


// An array containing the user handles of the desired recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/recipients
func (s_ SharingService) SetRecipients(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecipients:"), nsArray)
}/* debug [instance_properties/setter]: recipients */


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/subject
func (s_ SharingService) Subject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("subject"))
	return rv
}/* debug [instance_properties/getter]: subject */


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/subject
func (s_ SharingService) SetSubject(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubject:"), value)
}/* debug [instance_properties/setter]: subject */


// The title of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/title
func (s_ SharingService) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSharingService */


