// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [SharingService] class.
type ISharingService interface {
	objectivec.IObject
	

	// properties:
	AccountName() foundation.foundation.INSString
	AlternateImage() IImage
	AttachmentFileURLs() []foundation.URL
	Image() IImage
	MenuItemTitle() foundation.foundation.INSString
	SetMenuItemTitle(value foundation.foundation.INSString)
	MessageBody() foundation.foundation.INSString
	PermanentLink() foundation.foundation.INSURL
	Recipients() []string
	SetRecipients(value []string)
	Subject() foundation.foundation.INSString
	SetSubject(value foundation.foundation.INSString)
	Title() foundation.foundation.INSString


	

	// methods:
	CanPerformWithItems(items foundation.foundation.INSArray) bool
	PerformWithItems(items foundation.foundation.INSArray)


}





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






// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func NewSharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := objc.Send[SharingService](objc.ID(getSharingServiceClass().class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}


// Creates a custom sharing service object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(title:image:alternateImage:handler:)
func NewSharingServiceWithTitleImageAlternateImageHandler(title foundation.foundation.INSString, image IImage, alternateImage IImage, block unsafe.Pointer) SharingService {
	instance := getSharingServiceClass().Alloc()
	rv := objc.Send[SharingService](instance.ID, objc.Sel("initWithTitle:image:alternateImage:handler:"), title, image, alternateImage, block)
	rv.Autorelease()
	return rv
}







// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) ISharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}


// Returns a list of sharing services which could share all the provided items together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items foundation.foundation.INSArray) []SharingService {
	rv := objc.Send[[]SharingService](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}












// Returns whether the service can share all the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/canPerform(withItems:)
func (s_ SharingService) CanPerformWithItems(items foundation.foundation.INSArray) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canPerformWithItems:"), items)
	return rv
}


// Manually performs the service on the provided items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/perform(withItems:)
func (s_ SharingService) PerformWithItems(items foundation.foundation.INSArray) {
	objc.Send[objc.ID](s_.ID, objc.Sel("performWithItems:"), items)
}







// The account name used for posting on Twitter or Sina Weibo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/accountName
func (s_ SharingService) AccountName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("accountName"))
	return rv
}


// The alternate image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/alternateImage
func (s_ SharingService) AlternateImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// An array of NSURL objects representing the files that were shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/attachmentFileURLs
func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](s_.ID, objc.Sel("attachmentFileURLs"))
	return rv
}


// The primary image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/image
func (s_ SharingService) Image() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/menuItemTitle
func (s_ SharingService) MenuItemTitle() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("menuItemTitle"))
	return rv
}


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/menuItemTitle
func (s_ SharingService) SetMenuItemTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenuItemTitle:"), value)
}


// The message body as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/messageBody
func (s_ SharingService) MessageBody() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("messageBody"))
	return rv
}


// A permanent URL (permalink) that your app can use to access the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/permanentLink
func (s_ SharingService) PermanentLink() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("permanentLink"))
	return rv
}


// An array containing the user handles of the desired recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/recipients
func (s_ SharingService) Recipients() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("recipients"))
	return rv
}


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
}


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/subject
func (s_ SharingService) Subject() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/subject
func (s_ SharingService) SetSubject(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubject:"), value)
}


// The title of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/title
func (s_ SharingService) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}







