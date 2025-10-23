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
	// properties:
	ContextID() objc.IObject /* cross-framework: UUID */
	OriginalMessage() IMEMessage
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	IsEncrypted() bool /* primitive/slice/pointer. */
	SetIsEncrypted(value bool /* primitive/slice/pointer. */)
	IsSigned() bool /* primitive/slice/pointer. */
	SetIsSigned(value bool /* primitive/slice/pointer. */)
	ShouldEncrypt() bool /* primitive/slice/pointer. */
	SetShouldEncrypt(value bool /* primitive/slice/pointer. */)
	ShouldSign() bool /* primitive/slice/pointer. */
	SetShouldSign(value bool /* primitive/slice/pointer. */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/contextID
func (m_ MEComposeContext) ContextID() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[UUID](m_.ID, objc.Sel("contextID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/originalMessage
func (m_ MEComposeContext) OriginalMessage() IMEMessage {
	rv := objc.Send[MEMessage](m_.ID, objc.Sel("originalMessage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/action
func (m_ MEComposeContext) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/action
func (m_ MEComposeContext) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/isencrypted
func (m_ MEComposeContext) IsEncrypted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEncrypted"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/isencrypted
func (m_ MEComposeContext) SetIsEncrypted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEncrypted:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/issigned
func (m_ MEComposeContext) IsSigned() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSigned"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/issigned
func (m_ MEComposeContext) SetIsSigned(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSigned:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldencrypt
func (m_ MEComposeContext) ShouldEncrypt() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldEncrypt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldencrypt
func (m_ MEComposeContext) SetShouldEncrypt(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldEncrypt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldsign
func (m_ MEComposeContext) ShouldSign() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldSign"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/mecomposecontext/shouldsign
func (m_ MEComposeContext) SetShouldSign(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldSign:"), value)
}



