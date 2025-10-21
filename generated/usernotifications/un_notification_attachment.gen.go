// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationAttachment] class.
var (
	UNNotificationAttachmentClass     _UNNotificationAttachmentClass
	UNNotificationAttachmentClassOnce sync.Once
)

func getUNNotificationAttachmentClass() _UNNotificationAttachmentClass {
	UNNotificationAttachmentClassOnce.Do(func() {
		UNNotificationAttachmentClass = _UNNotificationAttachmentClass{objc.GetClass("UNNotificationAttachment")}
	})
	return UNNotificationAttachmentClass
}

type _UNNotificationAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationAttachment] class.
type IUNNotificationAttachment interface {
	objectivec.IObject
}

// A media file associated with a notification.
//
// Create a object when you want to include audio, image, or video content together in an alert-based notification. When creating the object, the file you specify must be on disk, and the file format must be one of the supported types. You’re responsible for supplying attachments before the system displays your notification’s alert. For local notifications, add attachments when creating the notification’s content. For remote notifications, use a notification service app extension to download the attached files and then add them to the notification’s content before delivery. The system validates attachments before displaying the associated notification. If you attach a file to a local notification request that’s corrupted, invalid, or of an unsupported file type, the system doesn’t schedule your request. For remote notifications, the system validates attachments after your notification service app extension finishes. Once validated, the system moves the attached files into the attachment data store so that the appropriate processes can access the files. The system copies attachments located inside an app’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment
type UNNotificationAttachment struct {
	objectivec.Object
}

// UNNotificationAttachmentFrom constructs a [UNNotificationAttachment] from an unsafe.Pointer.
//
// A media file associated with a notification.
func UNNotificationAttachmentFrom(ptr unsafe.Pointer) UNNotificationAttachment {
	return UNNotificationAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationAttachmentClass) Alloc() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationAttachmentClass) New() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationAttachment) Init() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationAttachment) Autorelease() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAttachment creates a new UNNotificationAttachment instance.
func NewUNNotificationAttachment() UNNotificationAttachment {
	return getUNNotificationAttachmentClass().New()
}




// Creates an attachment object from the specified file and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/init(identifier:url:options:)
func NewUNNotificationAttachmentWithIdentifierURLOptionsError(identifier string, URL unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](objc.ID(getUNNotificationAttachmentClass().class), objc.Sel("attachmentWithIdentifier:URL:options:error:"), objc.String(identifier), URL, options, error_)
	return rv
}


// Creates an attachment object from the specified file and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/init(identifier:url:options:)
func (uc _UNNotificationAttachmentClass) AttachmentWithIdentifierURLOptionsError(identifier string, URL unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("attachmentWithIdentifier:URL:options:error:"), objc.String(identifier), URL, options, error_)
	return rv
}

// The unique identifier for the attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/identifier
func (u_ UNNotificationAttachment) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}

// The UTI type of the attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/type
func (u_ UNNotificationAttachment) Type() string {
	rv := objc.Send[string](u_.ID, objc.Sel("type"))
	return rv
}

// The URL of the file for this attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/url
func (u_ UNNotificationAttachment) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URL"))
	return rv
}


