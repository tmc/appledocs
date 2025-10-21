// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CXSetTranslatingCallAction] class.
var (
	CXSetTranslatingCallActionClass     _CXSetTranslatingCallActionClass
	CXSetTranslatingCallActionClassOnce sync.Once
)

func getCXSetTranslatingCallActionClass() _CXSetTranslatingCallActionClass {
	CXSetTranslatingCallActionClassOnce.Do(func() {
		CXSetTranslatingCallActionClass = _CXSetTranslatingCallActionClass{objc.GetClass("CXSetTranslatingCallAction")}
	})
	return CXSetTranslatingCallActionClass
}

type _CXSetTranslatingCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXSetTranslatingCallAction] class.
type ICXSetTranslatingCallAction interface {
	ICXCallAction
	FulfillUsingTranslationEngine(translationEngine unsafe.Pointer)
}

// An encapsulation of the act of translating a call.
//
// is a concrete subclass of . When a caller chooses to translate a conversation, the system provides translated captions, and a translated transcript of the call and the sends the to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction
type CXSetTranslatingCallAction struct {
	CXCallAction
}

// CXSetTranslatingCallActionFrom constructs a [CXSetTranslatingCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of translating a call.
func CXSetTranslatingCallActionFrom(ptr unsafe.Pointer) CXSetTranslatingCallAction {
	return CXSetTranslatingCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXSetTranslatingCallActionClass) Alloc() CXSetTranslatingCallAction {
	rv := objc.Send[CXSetTranslatingCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXSetTranslatingCallActionClass) New() CXSetTranslatingCallAction {
	rv := objc.Send[CXSetTranslatingCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetTranslatingCallAction) Init() CXSetTranslatingCallAction {
	rv := objc.Send[CXSetTranslatingCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetTranslatingCallAction) Autorelease() CXSetTranslatingCallAction {
	rv := objc.Send[CXSetTranslatingCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetTranslatingCallAction creates a new CXSetTranslatingCallAction instance.
func NewCXSetTranslatingCallAction() CXSetTranslatingCallAction {
	return getCXSetTranslatingCallActionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/init(call:isTranslating:localLanguage:remoteLanguage:)
func NewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage(uuid unsafe.Pointer, isTranslating bool, localLanguage string, remoteLanguage string) CXSetTranslatingCallAction {
	instance := getCXSetTranslatingCallActionClass().Alloc()
	rv := objc.Send[CXSetTranslatingCallAction](instance.ID, objc.Sel("initWithCallUUID:isTranslating:localLanguage:remoteLanguage:"), uuid, isTranslating, objc.String(localLanguage), objc.String(remoteLanguage))
	rv.Autorelease()
	return rv
}

// Creates a new action to start or stop translating a call with the provided data.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/init(coder:)
func NewCXSetTranslatingCallActionWithCoder(aDecoder unsafe.Pointer) CXSetTranslatingCallAction {
	instance := getCXSetTranslatingCallActionClass().Alloc()
	rv := objc.Send[CXSetTranslatingCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/fulfill(using:)
func (c_ CXSetTranslatingCallAction) FulfillUsingTranslationEngine(translationEngine unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillUsingTranslationEngine:"), translationEngine)
}

// A value that indicates whether translation is active for a call.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/isTranslating
func (c_ CXSetTranslatingCallAction) IsTranslating() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isTranslating"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/localLanguage
func (c_ CXSetTranslatingCallAction) LocalLanguage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("localLanguage"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/remoteLanguage
func (c_ CXSetTranslatingCallAction) RemoteLanguage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("remoteLanguage"))
	return rv
}


