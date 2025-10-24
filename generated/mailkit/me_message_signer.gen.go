// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessageSigner */


/* debug [class_header]: Header for MEMessageSigner */
// The class instance for the [MEMessageSigner] class.
var (
	MEMessageSignerClass     _MEMessageSignerClass
	MEMessageSignerClassOnce sync.Once
)

func getMEMessageSignerClass() _MEMessageSignerClass {
	MEMessageSignerClassOnce.Do(func() {
		MEMessageSignerClass = _MEMessageSignerClass{objc.GetClass("MEMessageSigner")}
	})
	return MEMessageSignerClass
}

type _MEMessageSignerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessageSigner */
// An interface definition for the [MEMessageSigner] class.
type IMEMessageSigner interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessageSigner */
	// properties:
	Context() objc.IObject /* cross-framework: NSData */
	EmailAddresses() []MEEmailAddress
	Label() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessageSigner */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessageSigner */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageSignerClass) Alloc() MEMessageSigner {
	rv := objc.Send[MEMessageSigner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEMessageSignerClass) New() MEMessageSigner {
	rv := objc.Send[MEMessageSigner](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessageSigner) Init() MEMessageSigner {
	rv := objc.Send[MEMessageSigner](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessageSigner) Autorelease() MEMessageSigner {
	rv := objc.Send[MEMessageSigner](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessageSigner creates a new MEMessageSigner instance.
func NewMEMessageSigner() MEMessageSigner {
	return getMEMessageSignerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessageSigner */
// An object that contains details about the person who signed a message.


// An object that contains details about the person who signed a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSigner
type MEMessageSigner struct {
	objectivec.Object
}

// MEMessageSignerFrom constructs a [MEMessageSigner] from an unsafe.Pointer.
//
// An object that contains details about the person who signed a message.
func MEMessageSignerFrom(ptr unsafe.Pointer) MEMessageSigner {
	return MEMessageSigner{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessageSigner */

// Creates a new message signer object that contains the email addresses of the signers, a label, and context data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSigner/init(emailAddresses:signatureLabel:context:)
func NewMEMessageSignerWithEmailAddressesSignatureLabelContext(emailAddresses []MEEmailAddress, label objc.IObject /* cross-framework: NSString */, context objc.IObject /* cross-framework: NSData */) MEMessageSigner {
	instance := getMEMessageSignerClass().Alloc()
	rv := objc.Send[MEMessageSigner](instance.ID, objc.Sel("initWithEmailAddresses:signatureLabel:context:"), emailAddresses, label, context)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEMessageSignerWithEmailAddressesSignatureLabelContext */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessageSigner */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessageSigner */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessageSigner */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessageSigner */

// Data related to the message signature, such as the signing certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSigner/context
func (m_ MEMessageSigner) Context() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// An array of email addresses associated with the signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSigner/emailAddresses
func (m_ MEMessageSigner) EmailAddresses() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("emailAddresses"))
	return rv
}/* debug [instance_properties/getter]: emailAddresses */


// A string that the message’s headers use to display the message signer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSigner/label
func (m_ MEMessageSigner) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessageSigner */


