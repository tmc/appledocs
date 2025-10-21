// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEDecodedMessage] class.
var (
	MEDecodedMessageClass     _MEDecodedMessageClass
	MEDecodedMessageClassOnce sync.Once
)

func getMEDecodedMessageClass() _MEDecodedMessageClass {
	MEDecodedMessageClassOnce.Do(func() {
		MEDecodedMessageClass = _MEDecodedMessageClass{objc.GetClass("MEDecodedMessage")}
	})
	return MEDecodedMessageClass
}

type _MEDecodedMessageClass struct {
	class objc.Class
}

// An interface definition for the [MEDecodedMessage] class.
type IMEDecodedMessage interface {
	objectivec.IObject
}

// An object that contains the RFC 2822 data for a message, without encryption or digital signatures.
//
// When MailKit invokes your message security handler’s method, you decode the message data and return an instance of that contains unencrypted MIME data.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessage
type MEDecodedMessage struct {
	objectivec.Object
}

// MEDecodedMessageFrom constructs a [MEDecodedMessage] from an unsafe.Pointer.
//
// An object that contains the RFC 2822 data for a message, without encryption or digital signatures.
func MEDecodedMessageFrom(ptr unsafe.Pointer) MEDecodedMessage {
	return MEDecodedMessage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEDecodedMessageClass) Alloc() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEDecodedMessageClass) New() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEDecodedMessage) Init() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEDecodedMessage) Autorelease() MEDecodedMessage {
	rv := objc.Send[MEDecodedMessage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodedMessage creates a new MEDecodedMessage instance.
func NewMEDecodedMessage() MEDecodedMessage {
	return getMEDecodedMessageClass().New()
}




