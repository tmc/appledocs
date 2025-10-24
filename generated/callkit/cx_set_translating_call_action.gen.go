// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXSetTranslatingCallAction */


/* debug [class_header]: Header for CXSetTranslatingCallAction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXSetTranslatingCallAction */
// An interface definition for the [CXSetTranslatingCallAction] class.
type ICXSetTranslatingCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXSetTranslatingCallAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXSetTranslatingCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXSetTranslatingCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXSetTranslatingCallActionClass) Alloc() CXSetTranslatingCallAction {
	rv := objc.Send[CXSetTranslatingCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXSetTranslatingCallAction */
// An encapsulation of the act of translating a call.
//
// is a concrete subclass of . When a caller chooses to translate a conversation, the system provides translated captions, and a translated transcript of the call and the sends the to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed.


// An encapsulation of the act of translating a call.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXSetTranslatingCallAction */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/init(call:isTranslating:localLanguage:remoteLanguage:)
func NewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage(uuid foundation.UUID, isTranslating bool, localLanguage objc.IObject /* cross-framework: NSString */, remoteLanguage objc.IObject /* cross-framework: NSString */) CXSetTranslatingCallAction {
	instance := getCXSetTranslatingCallActionClass().Alloc()
	rv := objc.Send[CXSetTranslatingCallAction](instance.ID, objc.Sel("initWithCallUUID:isTranslating:localLanguage:remoteLanguage:"), uuid, isTranslating, localLanguage, remoteLanguage)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage */


// Creates a new action to start or stop translating a call with the provided data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/init(coder:)
func NewCXSetTranslatingCallActionWithCoder(aDecoder foundation.Coder) CXSetTranslatingCallAction {
	instance := getCXSetTranslatingCallActionClass().Alloc()
	rv := objc.Send[CXSetTranslatingCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetTranslatingCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXSetTranslatingCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXSetTranslatingCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXSetTranslatingCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXSetTranslatingCallAction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXSetTranslatingCallAction */


