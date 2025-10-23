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
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	AccountName() objc.IObject /* cross-framework: NSString */
	SetAccountName(value objc.IObject /* cross-framework: NSString */)
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	AttachmentFileURLs() objc.IObject /* cross-framework: URL */
	SetAttachmentFileURLs(value objc.IObject /* cross-framework: URL */)
	Image() IImage
	SetImage(value IImage)
	MenuItemTitle() objc.IObject /* cross-framework: NSString */
	SetMenuItemTitle(value objc.IObject /* cross-framework: NSString */)
	MessageBody() objc.IObject /* cross-framework: NSString */
	SetMessageBody(value objc.IObject /* cross-framework: NSString */)
	PermanentLink() objc.IObject /* cross-framework: URL */
	SetPermanentLink(value objc.IObject /* cross-framework: URL */)
	Recipients() objc.IObject /* cross-framework: NSString */
	SetRecipients(value objc.IObject /* cross-framework: NSString */)
	Subject() objc.IObject /* cross-framework: NSString */
	SetSubject(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
	CanPerformWithItems(items objc.IObject /* cross-framework NSArray */) bool /* primitive/slice/pointer. */
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



// Returns whether the service can share all the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/canPerform(withItems:)
func (s_ SharingService) CanPerformWithItems(items objc.IObject /* cross-framework NSArray */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("canPerformWithItems:"), items)
	return rv
}


// Specifies the delegate of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/delegate
func (s_ SharingService) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// Specifies the delegate of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/delegate
func (s_ SharingService) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// The account name used for posting on Twitter or Sina Weibo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/accountname
func (s_ SharingService) AccountName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("accountName"))
	return rv
}


// The account name used for posting on Twitter or Sina Weibo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/accountname
func (s_ SharingService) SetAccountName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccountName:"), value)
}


// The alternate image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/alternateimage
func (s_ SharingService) AlternateImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// The alternate image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/alternateimage
func (s_ SharingService) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateImage:"), value)
}


// An array of NSURL objects representing the files that were shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/attachmentfileurls
func (s_ SharingService) AttachmentFileURLs() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("attachmentFileURLs"))
	return rv
}


// An array of NSURL objects representing the files that were shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/attachmentfileurls
func (s_ SharingService) SetAttachmentFileURLs(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttachmentFileURLs:"), value)
}


// The primary image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/image
func (s_ SharingService) Image() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// The primary image representing the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/image
func (s_ SharingService) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/menuitemtitle
func (s_ SharingService) MenuItemTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("menuItemTitle"))
	return rv
}


// The title of the service in the Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/menuitemtitle
func (s_ SharingService) SetMenuItemTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenuItemTitle:"), value)
}


// The message body as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/messagebody
func (s_ SharingService) MessageBody() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("messageBody"))
	return rv
}


// The message body as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/messagebody
func (s_ SharingService) SetMessageBody(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMessageBody:"), value)
}


// A permanent URL (permalink) that your app can use to access the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/permanentlink
func (s_ SharingService) PermanentLink() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("permanentLink"))
	return rv
}


// A permanent URL (permalink) that your app can use to access the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/permanentlink
func (s_ SharingService) SetPermanentLink(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPermanentLink:"), value)
}


// An array containing the user handles of the desired recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/recipients
func (s_ SharingService) Recipients() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("recipients"))
	return rv
}


// An array containing the user handles of the desired recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/recipients
func (s_ SharingService) SetRecipients(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecipients:"), value)
}


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/subject
func (s_ SharingService) Subject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the post.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/subject
func (s_ SharingService) SetSubject(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubject:"), value)
}


// The title of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/title
func (s_ SharingService) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}


// The title of the sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/title
func (s_ SharingService) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}



