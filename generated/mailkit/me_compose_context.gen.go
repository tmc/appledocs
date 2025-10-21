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
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldencrypt
func (m_ MEComposeContext) ShouldEncrypt() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldEncrypt"))
	return rv
}


// SetShouldEncrypt sets the value of the shouldEncrypt property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldencrypt
func (m_ MEComposeContext) SetShouldEncrypt(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldEncrypt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/issigned
func (m_ MEComposeContext) IsSigned() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSigned"))
	return rv
}


// SetIsSigned sets the value of the isSigned property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/issigned
func (m_ MEComposeContext) SetIsSigned(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSigned:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/contextid
func (m_ MEComposeContext) ContextID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("contextID"))
	return rv
}


// SetContextID sets the value of the contextID property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/contextid
func (m_ MEComposeContext) SetContextID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContextID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldsign
func (m_ MEComposeContext) ShouldSign() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldSign"))
	return rv
}


// SetShouldSign sets the value of the shouldSign property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldsign
func (m_ MEComposeContext) SetShouldSign(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldSign:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/isencrypted
func (m_ MEComposeContext) IsEncrypted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEncrypted"))
	return rv
}


// SetIsEncrypted sets the value of the isEncrypted property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/isencrypted
func (m_ MEComposeContext) SetIsEncrypted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEncrypted:"), value)
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



