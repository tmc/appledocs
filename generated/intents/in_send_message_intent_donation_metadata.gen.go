// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendMessageIntentDonationMetadata] class.
var (
	INSendMessageIntentDonationMetadataClass     _INSendMessageIntentDonationMetadataClass
	INSendMessageIntentDonationMetadataClassOnce sync.Once
)

func getINSendMessageIntentDonationMetadataClass() _INSendMessageIntentDonationMetadataClass {
	INSendMessageIntentDonationMetadataClassOnce.Do(func() {
		INSendMessageIntentDonationMetadataClass = _INSendMessageIntentDonationMetadataClass{objc.GetClass("INSendMessageIntentDonationMetadata")}
	})
	return INSendMessageIntentDonationMetadataClass
}

type _INSendMessageIntentDonationMetadataClass struct {
	class objc.Class
}

// An interface definition for the [INSendMessageIntentDonationMetadata] class.
type IINSendMessageIntentDonationMetadata interface {
	IINIntentDonationMetadata
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata
type INSendMessageIntentDonationMetadata struct {
	INIntentDonationMetadata
}

// INSendMessageIntentDonationMetadataFrom constructs a [INSendMessageIntentDonationMetadata] from an unsafe.Pointer.
func INSendMessageIntentDonationMetadataFrom(ptr unsafe.Pointer) INSendMessageIntentDonationMetadata {
	return INSendMessageIntentDonationMetadata{
		INIntentDonationMetadata: INIntentDonationMetadataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendMessageIntentDonationMetadataClass) Alloc() INSendMessageIntentDonationMetadata {
	rv := objc.Send[INSendMessageIntentDonationMetadata](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendMessageIntentDonationMetadataClass) New() INSendMessageIntentDonationMetadata {
	rv := objc.Send[INSendMessageIntentDonationMetadata](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendMessageIntentDonationMetadata) Init() INSendMessageIntentDonationMetadata {
	rv := objc.Send[INSendMessageIntentDonationMetadata](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendMessageIntentDonationMetadata) Autorelease() INSendMessageIntentDonationMetadata {
	rv := objc.Send[INSendMessageIntentDonationMetadata](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendMessageIntentDonationMetadata creates a new INSendMessageIntentDonationMetadata instance.
func NewINSendMessageIntentDonationMetadata() INSendMessageIntentDonationMetadata {
	return getINSendMessageIntentDonationMetadataClass().New()
}



//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/isReplyToCurrentUser
func (i_ INSendMessageIntentDonationMetadata) ReplyToCurrentUser() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("replyToCurrentUser"))
	return rv
}


// SetReplyToCurrentUser sets the value of the replyToCurrentUser property.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/isReplyToCurrentUser
func (i_ INSendMessageIntentDonationMetadata) SetReplyToCurrentUser(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReplyToCurrentUser:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/mentionsCurrentUser
func (i_ INSendMessageIntentDonationMetadata) MentionsCurrentUser() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("mentionsCurrentUser"))
	return rv
}


// SetMentionsCurrentUser sets the value of the mentionsCurrentUser property.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/mentionsCurrentUser
func (i_ INSendMessageIntentDonationMetadata) SetMentionsCurrentUser(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMentionsCurrentUser:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/notifyRecipientAnyway
func (i_ INSendMessageIntentDonationMetadata) NotifyRecipientAnyway() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("notifyRecipientAnyway"))
	return rv
}


// SetNotifyRecipientAnyway sets the value of the notifyRecipientAnyway property.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/notifyRecipientAnyway
func (i_ INSendMessageIntentDonationMetadata) SetNotifyRecipientAnyway(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNotifyRecipientAnyway:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/recipientCount
func (i_ INSendMessageIntentDonationMetadata) RecipientCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("recipientCount"))
	return rv
}


// SetRecipientCount sets the value of the recipientCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentDonationMetadata/recipientCount
func (i_ INSendMessageIntentDonationMetadata) SetRecipientCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecipientCount:"), value)
}

