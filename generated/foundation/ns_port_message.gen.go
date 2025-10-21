// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PortMessage] class.
var (
	PortMessageClass     _PortMessageClass
	PortMessageClassOnce sync.Once
)

func getPortMessageClass() _PortMessageClass {
	PortMessageClassOnce.Do(func() {
		PortMessageClass = _PortMessageClass{objc.GetClass("NSPortMessage")}
	})
	return PortMessageClass
}

type _PortMessageClass struct {
	class objc.Class
}

// An interface definition for the [PortMessage] class.
type IPortMessage interface {
	objectivec.IObject
	SendBeforeDate(date unsafe.Pointer) bool
}

// A low-level, operating system-independent type for inter-application (and inter-thread) messages.
//
// Port messages are used primarily by the distributed objects system. You should implement inter-application communication using distributed objects whenever possible and use only when necessary. An object has three major parts: the send and receive ports, which are objects that link the sender of the message to the receiver, and the components, which form the body of the message. The components are held as an object containing and objects. The message sends the components out through the send port; any replies to the message arrive on the receive port. See the class specification for information on handling incoming messages. An instance can be initialized with a pair of objects and an array of components. A port message’s body can contain only objects or objects. In the distributed objects system the byte/character arrays are usually encoded objects that are being forwarded from a proxy to the corresponding real object. An object also maintains a message identifier, which can be used to indicate the class of a message, such as an Objective-C method invocation, a connection request, an error, and so on. Use the and methods to access the identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage
type PortMessage struct {
	objectivec.Object
}

// PortMessageFrom constructs a [PortMessage] from an unsafe.Pointer.
//
// A low-level, operating system-independent type for inter-application (and inter-thread) messages.
func PortMessageFrom(ptr unsafe.Pointer) PortMessage {
	return PortMessage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PortMessageClass) Alloc() PortMessage {
	rv := objc.Send[PortMessage](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PortMessageClass) New() PortMessage {
	rv := objc.Send[PortMessage](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PortMessage) Init() PortMessage {
	rv := objc.Send[PortMessage](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PortMessage) Autorelease() PortMessage {
	rv := objc.Send[PortMessage](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortMessage creates a new PortMessage instance.
func NewPortMessage() PortMessage {
	return getPortMessageClass().New()
}




// Initializes a newly allocated object to send given data on a given port and to receiver replies on another given port.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/init(send:receive:components:)
func NewPortMessageWithSendPortReceivePortComponents(sendPort unsafe.Pointer, replyPort unsafe.Pointer, components objc.ID) PortMessage {
	instance := getPortMessageClass().Alloc()
	rv := objc.Send[PortMessage](instance.ID, objc.Sel("initWithSendPort:receivePort:components:"), sendPort, replyPort, components)
	rv.Autorelease()
	return rv
}


// Attempts to send the message before the specified date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/send(before:)
func (p_ PortMessage) SendBeforeDate(date unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("sendBeforeDate:"), date)
	return rv
}

// Returns the data components of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/components
func (p_ PortMessage) Components() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("components"))
	return rv
}

// Returns the identifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/msgid
func (p_ PortMessage) Msgid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("msgid"))
	return rv
}


// SetMsgid sets the value of the msgid property.
// Returns the identifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/msgid
func (p_ PortMessage) SetMsgid(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMsgid:"), value)
}
// For an outgoing message, returns the port on which replies to the receiver will arrive. For an incoming message, returns the port the receiver did arrive on.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/receivePort
func (p_ PortMessage) ReceivePort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("receivePort"))
	return rv
}

// For an outgoing message, returns the port the receiver will send itself through. For an incoming message, returns the port replies to the receiver should be sent through.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PortMessage/sendPort
func (p_ PortMessage) SendPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sendPort"))
	return rv
}


