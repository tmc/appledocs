// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEComposeSession] class.
var (
	MEComposeSessionClass     _MEComposeSessionClass
	MEComposeSessionClassOnce sync.Once
)

func getMEComposeSessionClass() _MEComposeSessionClass {
	MEComposeSessionClassOnce.Do(func() {
		MEComposeSessionClass = _MEComposeSessionClass{objc.GetClass("MEComposeSession")}
	})
	return MEComposeSessionClass
}

type _MEComposeSessionClass struct {
	class objc.Class
}

// An interface definition for the [MEComposeSession] class.
type IMEComposeSession interface {
	objectivec.IObject
}

// An object that represents a single mail compose window.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSession
type MEComposeSession struct {
	objectivec.Object
}

// MEComposeSessionFrom constructs a [MEComposeSession] from an unsafe.Pointer.
//
// An object that represents a single mail compose window.
func MEComposeSessionFrom(ptr unsafe.Pointer) MEComposeSession {
	return MEComposeSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEComposeSessionClass) Alloc() MEComposeSession {
	rv := objc.Send[MEComposeSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEComposeSessionClass) New() MEComposeSession {
	rv := objc.Send[MEComposeSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEComposeSession) Init() MEComposeSession {
	rv := objc.Send[MEComposeSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEComposeSession) Autorelease() MEComposeSession {
	rv := objc.Send[MEComposeSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEComposeSession creates a new MEComposeSession instance.
func NewMEComposeSession() MEComposeSession {
	return getMEComposeSessionClass().New()
}




