// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEComposeContext] class.
var (
	MEComposeContextClass     _MEComposeContextClass
	MEComposeContextClassOnce sync.Once
)

func getMEComposeContextClass() _MEComposeContextClass {
	MEComposeContextClassOnce.Do(func() {
		MEComposeContextClass = _MEComposeContextClass{objc.GetClass("MEComposeContext")}
	})
	return MEComposeContextClass
}

type _MEComposeContextClass struct {
	class objc.Class
}

// An interface definition for the [MEComposeContext] class.
type IMEComposeContext interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext
type MEComposeContext struct {
	objectivec.Object
}

// MEComposeContextFrom constructs a [MEComposeContext] from an unsafe.Pointer.
func MEComposeContextFrom(ptr unsafe.Pointer) MEComposeContext {
	return MEComposeContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEComposeContextClass) Alloc() MEComposeContext {
	rv := objc.Send[MEComposeContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEComposeContextClass) New() MEComposeContext {
	rv := objc.Send[MEComposeContext](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEComposeContext) Init() MEComposeContext {
	rv := objc.Send[MEComposeContext](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEComposeContext) Autorelease() MEComposeContext {
	rv := objc.Send[MEComposeContext](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEComposeContext creates a new MEComposeContext instance.
func NewMEComposeContext() MEComposeContext {
	return getMEComposeContextClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/action
func (m_ MEComposeContext) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("action"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/originalMessage
func (m_ MEComposeContext) OriginalMessage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("originalMessage"))
	return rv
}



