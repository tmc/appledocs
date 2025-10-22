// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AccountName() string
	SetAccountName(value string)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	AttachmentFileURLs() foundation.URL
	SetAttachmentFileURLs(value foundation.IURL)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Image() Image
	SetImage(value IImage)
	MenuItemTitle() string
	SetMenuItemTitle(value string)
	MessageBody() string
	SetMessageBody(value string)
	PermanentLink() foundation.URL
	SetPermanentLink(value foundation.IURL)
	Recipients() string
	SetRecipients(value string)
	Subject() string
	SetSubject(value string)
	Title() string
	SetTitle(value string)
}

// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
//
// An object provides a consistent user experience for sharing items— objects, objects, objects, video (through file URLs), of any object that implements the protocol—in macOS. For any item or group of items, the displays a sheet with the content to share. A sharing service can create a post on a social network like Twitter or Facebook, send a message by email or iMessage, upload videos to viewing services, or send a file using AirDrop. You can use objects directly in your app. The following example shows how to create a button that shares content directly to a social media service.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func NewSharingServiceNamed(serviceName ISharingServiceName) SharingService {
	rv := objc.Send[SharingService](objc.ID(getSharingServiceClass().class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}



// Creates a custom sharing service object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(title:image:alternateImage:handler:)
func NewSharingServiceWithTitleImageAlternateImageHandler(title string, image IImage, alternateImage IImage, block unsafe.Pointer) SharingService {
	instance := getSharingServiceClass().Alloc()
	rv := objc.Send[SharingService](instance.ID, objc.Sel("initWithTitle:image:alternateImage:handler:"), objc.String(title), image, alternateImage, block)
	rv.Autorelease()
	return rv
}


// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func (sc _SharingServiceClass) SharingServiceNamed(serviceName ISharingServiceName) SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}

// Returns a list of sharing services which could share all the provided items together.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items objectivec.IObject) []SharingService {
	rv := objc.Send[[]SharingService](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}

// The account name used for posting on Twitter or Sina Weibo.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/accountname
func (s_ SharingService) AccountName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("accountName"))
	return rv
}


// SetAccountName sets the value of the accountName property.
// The account name used for posting on Twitter or Sina Weibo.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/accountname
func (s_ SharingService) SetAccountName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccountName:"), objc.String(value))
}

// The alternate image representing the sharing service.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/alternateimage
func (s_ SharingService) AlternateImage() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// SetAlternateImage sets the value of the alternateImage property.
// The alternate image representing the sharing service.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/alternateimage
func (s_ SharingService) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateImage:"), value)
}

// An array of NSURL objects representing the files that were shared.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/attachmentfileurls
func (s_ SharingService) AttachmentFileURLs() foundation.URL {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("attachmentFileURLs"))
	return rv
}


// SetAttachmentFileURLs sets the value of the attachmentFileURLs property.
// An array of NSURL objects representing the files that were shared.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/attachmentfileurls
func (s_ SharingService) SetAttachmentFileURLs(value foundation.IURL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttachmentFileURLs:"), value)
}

// Specifies the delegate of the sharing service.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/delegate
func (s_ SharingService) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Specifies the delegate of the sharing service.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/delegate
func (s_ SharingService) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// The primary image representing the sharing service.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/image
func (s_ SharingService) Image() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The primary image representing the sharing service.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/image
func (s_ SharingService) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}

// The title of the service in the Share menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/menuitemtitle
func (s_ SharingService) MenuItemTitle() string {
	rv := objc.Send[string](s_.ID, objc.Sel("menuItemTitle"))
	return rv
}


// SetMenuItemTitle sets the value of the menuItemTitle property.
// The title of the service in the Share menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/menuitemtitle
func (s_ SharingService) SetMenuItemTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenuItemTitle:"), objc.String(value))
}

// The message body as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/messagebody
func (s_ SharingService) MessageBody() string {
	rv := objc.Send[string](s_.ID, objc.Sel("messageBody"))
	return rv
}


// SetMessageBody sets the value of the messageBody property.
// The message body as a string.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/messagebody
func (s_ SharingService) SetMessageBody(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMessageBody:"), objc.String(value))
}

// A permanent URL (permalink) that your app can use to access the post.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/permanentlink
func (s_ SharingService) PermanentLink() foundation.URL {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("permanentLink"))
	return rv
}


// SetPermanentLink sets the value of the permanentLink property.
// A permanent URL (permalink) that your app can use to access the post.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/permanentlink
func (s_ SharingService) SetPermanentLink(value foundation.IURL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPermanentLink:"), value)
}

// An array containing the user handles of the desired recipients.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/recipients
func (s_ SharingService) Recipients() string {
	rv := objc.Send[string](s_.ID, objc.Sel("recipients"))
	return rv
}


// SetRecipients sets the value of the recipients property.
// An array containing the user handles of the desired recipients.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/recipients
func (s_ SharingService) SetRecipients(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecipients:"), objc.String(value))
}

// The subject of the post.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/subject
func (s_ SharingService) Subject() string {
	rv := objc.Send[string](s_.ID, objc.Sel("subject"))
	return rv
}


// SetSubject sets the value of the subject property.
// The subject of the post.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/subject
func (s_ SharingService) SetSubject(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubject:"), objc.String(value))
}

// The title of the sharing service.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/title
func (s_ SharingService) Title() string {
	rv := objc.Send[string](s_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the sharing service.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/title
func (s_ SharingService) SetTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), objc.String(value))
}


