// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEComposeContext */


/* debug [class_header]: Header for MEComposeContext */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEComposeContext */
// An interface definition for the [MEComposeContext] class.
type IMEComposeContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEComposeContext */
	// properties:
	Action() MEComposeUserAction
	ContextID() foundation.UUID
	IsEncrypted() bool
	IsSigned() bool
	OriginalMessage() IMEMessage
	ShouldEncrypt() bool
	ShouldSign() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEComposeContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEComposeContext */
// Alloc allocates a new instance without initialization.
func (mc _MEComposeContextClass) Alloc() MEComposeContext {
	rv := objc.Send[MEComposeContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEComposeContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext
type MEComposeContext struct {
	objectivec.Object
}

// MEComposeContextFrom constructs a [MEComposeContext] from an unsafe.Pointer.
func MEComposeContextFrom(ptr unsafe.Pointer) MEComposeContext {
	return MEComposeContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEComposeContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEComposeContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEComposeContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEComposeContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEComposeContext */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/action
func (m_ MEComposeContext) Action() MEComposeUserAction {
	rv := objc.Send[MEComposeUserAction](m_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/contextID
func (m_ MEComposeContext) ContextID() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("contextID"))
	return rv
}/* debug [instance_properties/getter]: contextID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/isEncrypted
func (m_ MEComposeContext) IsEncrypted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEncrypted"))
	return rv
}/* debug [instance_properties/getter]: isEncrypted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/isSigned
func (m_ MEComposeContext) IsSigned() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSigned"))
	return rv
}/* debug [instance_properties/getter]: isSigned */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/originalMessage
func (m_ MEComposeContext) OriginalMessage() IMEMessage {
	rv := objc.Send[MEMessage](m_.ID, objc.Sel("originalMessage"))
	return rv
}/* debug [instance_properties/getter]: originalMessage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/shouldEncrypt
func (m_ MEComposeContext) ShouldEncrypt() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldEncrypt"))
	return rv
}/* debug [instance_properties/getter]: shouldEncrypt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeContext/shouldSign
func (m_ MEComposeContext) ShouldSign() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldSign"))
	return rv
}/* debug [instance_properties/getter]: shouldSign */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEComposeContext */



