
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [message] class.
var messageClass _messageClass

func init() {
	messageClass = _messageClass{objc.GetClass("message")}
}

type _messageClass struct {
	objc.Class
}

// An interface definition for the [message] class.
type Imessage interface {
	ID() objc.ID
}

type message struct {
	id objc.ID
}

func messageFrom(ptr unsafe.Pointer) message {
	return message{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ message) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _messageClass) Alloc() message {
	rv := objc.Send[message](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _messageClass) New() message {
	rv := objc.Send[message](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmessage creates and returns a new initialized instance.
func Newmessage() message {
	return messageClass.New()
}

// Init initializes the instance.
func (m_ message) Init() message {
	rv := objc.Send[message](m_.ID(), selInit)
	return rv
}
