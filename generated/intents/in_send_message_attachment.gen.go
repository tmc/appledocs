// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INSendMessageAttachment] class.
var (
	INSendMessageAttachmentClass     _INSendMessageAttachmentClass
	INSendMessageAttachmentClassOnce sync.Once
)

func getINSendMessageAttachmentClass() _INSendMessageAttachmentClass {
	INSendMessageAttachmentClassOnce.Do(func() {
		INSendMessageAttachmentClass = _INSendMessageAttachmentClass{objc.GetClass("INSendMessageAttachment")}
	})
	return INSendMessageAttachmentClass
}

type _INSendMessageAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [INSendMessageAttachment] class.
type IINSendMessageAttachment interface {
	objectivec.IObject
}

// A file to include in a message.
//
// When the type of a message is , the intent includes the audio recording as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageAttachment
type INSendMessageAttachment struct {
	objectivec.Object
}

// INSendMessageAttachmentFrom constructs a [INSendMessageAttachment] from an unsafe.Pointer.
//
// A file to include in a message.
func INSendMessageAttachmentFrom(ptr unsafe.Pointer) INSendMessageAttachment {
	return INSendMessageAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendMessageAttachmentClass) Alloc() INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendMessageAttachmentClass) New() INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendMessageAttachment) Init() INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendMessageAttachment) Autorelease() INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendMessageAttachment creates a new INSendMessageAttachment instance.
func NewINSendMessageAttachment() INSendMessageAttachment {
	return getINSendMessageAttachmentClass().New()
}




// Creates a message attachment with an audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageAttachment/init(audioMessageFile:)
func NewINSendMessageAttachmentWithAudioMessageFile(audioMessageFile unsafe.Pointer) INSendMessageAttachment {
	rv := objc.Send[INSendMessageAttachment](objc.ID(getINSendMessageAttachmentClass().class), objc.Sel("attachmentWithAudioMessageFile:"), audioMessageFile)
	return rv
}


// Creates a message attachment with an audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageAttachment/init(audioMessageFile:)
func (ic _INSendMessageAttachmentClass) AttachmentWithAudioMessageFile(audioMessageFile unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("attachmentWithAudioMessageFile:"), audioMessageFile)
	return rv
}

// The attachment’s recorded message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageAttachment/audioMessageFile
func (i_ INSendMessageAttachment) AudioMessageFile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("audioMessageFile"))
	return rv
}


