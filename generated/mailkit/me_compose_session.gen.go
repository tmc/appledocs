// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEComposeSession */


/* debug [class_header]: Header for MEComposeSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEComposeSession */
// An interface definition for the [MEComposeSession] class.
type IMEComposeSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEComposeSession */
	// properties:
	ComposeContext() IMEComposeContext
	MailMessage() IMEMessage
	SessionID() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEComposeSession */
	// methods:
	ReloadSession()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEComposeSession */
// Alloc allocates a new instance without initialization.
func (mc _MEComposeSessionClass) Alloc() MEComposeSession {
	rv := objc.Send[MEComposeSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEComposeSession */
// An object that represents a single mail compose window.


// An object that represents a single mail compose window.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEComposeSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEComposeSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEComposeSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEComposeSession */

// Refreshes the compose session with the extension’s new information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSession/reload()
func (m_ MEComposeSession) ReloadSession() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reloadSession"))
}/* debug [instance_methods/method]: ReloadSession */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEComposeSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSession/composeContext
func (m_ MEComposeSession) ComposeContext() IMEComposeContext {
	rv := objc.Send[MEComposeContext](m_.ID, objc.Sel("composeContext"))
	return rv
}/* debug [instance_properties/getter]: composeContext */


// The properties of the mail message, such as the subject and recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSession/mailMessage
func (m_ MEComposeSession) MailMessage() IMEMessage {
	rv := objc.Send[MEMessage](m_.ID, objc.Sel("mailMessage"))
	return rv
}/* debug [instance_properties/getter]: mailMessage */


// A unique identifier for the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEComposeSession/sessionID
func (m_ MEComposeSession) SessionID() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEComposeSession */



