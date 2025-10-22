// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [msgid] class.
var (
	MsgidClass     _msgidClass
	MsgidClassOnce sync.Once
)

func getmsgidClass() _msgidClass {
	MsgidClassOnce.Do(func() {
		MsgidClass = _msgidClass{objc.GetClass("msgid")}
	})
	return MsgidClass
}

type _msgidClass struct {
	class objc.Class
}

// An interface definition for the [msgid] class.
type Imsgid interface {
	objectivec.IObject
}

//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/msgid-c.ivar
type msgid struct {
	objectivec.Object
}

// msgidFrom constructs a [msgid] from an unsafe.Pointer.
func msgidFrom(ptr unsafe.Pointer) msgid {
	return msgid{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _msgidClass) Alloc() msgid {
	rv := objc.Send[msgid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _msgidClass) New() msgid {
	rv := objc.Send[msgid](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ msgid) Init() msgid {
	rv := objc.Send[msgid](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ msgid) Autorelease() msgid {
	rv := objc.Send[msgid](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmsgid creates a new msgid instance.
func Newmsgid() msgid {
	return getmsgidClass().New()
}




