// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEOutgoingMessageEncodingStatus */


/* debug [class_header]: Header for MEOutgoingMessageEncodingStatus */
// The class instance for the [MEOutgoingMessageEncodingStatus] class.
var (
	MEOutgoingMessageEncodingStatusClass     _MEOutgoingMessageEncodingStatusClass
	MEOutgoingMessageEncodingStatusClassOnce sync.Once
)

func getMEOutgoingMessageEncodingStatusClass() _MEOutgoingMessageEncodingStatusClass {
	MEOutgoingMessageEncodingStatusClassOnce.Do(func() {
		MEOutgoingMessageEncodingStatusClass = _MEOutgoingMessageEncodingStatusClass{objc.GetClass("MEOutgoingMessageEncodingStatus")}
	})
	return MEOutgoingMessageEncodingStatusClass
}

type _MEOutgoingMessageEncodingStatusClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEOutgoingMessageEncodingStatus */
// An interface definition for the [MEOutgoingMessageEncodingStatus] class.
type IMEOutgoingMessageEncodingStatus interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEOutgoingMessageEncodingStatus */
	// properties:
	AddressesFailingEncryption() []MEEmailAddress
	CanEncrypt() bool
	CanSign() bool
	SecurityError() objc.IObject /* cross-framework: Error */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEOutgoingMessageEncodingStatus */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEOutgoingMessageEncodingStatus */
// Alloc allocates a new instance without initialization.
func (mc _MEOutgoingMessageEncodingStatusClass) Alloc() MEOutgoingMessageEncodingStatus {
	rv := objc.Send[MEOutgoingMessageEncodingStatus](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEOutgoingMessageEncodingStatusClass) New() MEOutgoingMessageEncodingStatus {
	rv := objc.Send[MEOutgoingMessageEncodingStatus](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEOutgoingMessageEncodingStatus) Init() MEOutgoingMessageEncodingStatus {
	rv := objc.Send[MEOutgoingMessageEncodingStatus](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEOutgoingMessageEncodingStatus) Autorelease() MEOutgoingMessageEncodingStatus {
	rv := objc.Send[MEOutgoingMessageEncodingStatus](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEOutgoingMessageEncodingStatus creates a new MEOutgoingMessageEncodingStatus instance.
func NewMEOutgoingMessageEncodingStatus() MEOutgoingMessageEncodingStatus {
	return getMEOutgoingMessageEncodingStatusClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEOutgoingMessageEncodingStatus */
// An object that contains information about security measures the user can apply when composing a message.
//
// As a user composes a new message, MailKit requests the encoding status from your message security handler. The handler provides an that contains: Boolean values that indicate if the handler can sign or encrypt the message An error if verifying the security status fails An array of recipient addresses for which the handler can’t encrypt the message


// An object that contains information about security measures the user can apply when composing a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus
type MEOutgoingMessageEncodingStatus struct {
	objectivec.Object
}

// MEOutgoingMessageEncodingStatusFrom constructs a [MEOutgoingMessageEncodingStatus] from an unsafe.Pointer.
//
// An object that contains information about security measures the user can apply when composing a message.
func MEOutgoingMessageEncodingStatusFrom(ptr unsafe.Pointer) MEOutgoingMessageEncodingStatus {
	return MEOutgoingMessageEncodingStatus{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEOutgoingMessageEncodingStatus */

// Creates an object that describes whether the message security handler can encrypt or sign an outgoing message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus/init(canSign:canEncrypt:securityError:addressesFailingEncryption:)
func NewMEOutgoingMessageEncodingStatusWithCanSignCanEncryptSecurityErrorAddressesFailingEncryption(canSign bool, canEncrypt bool, securityError objc.IObject /* cross-framework: Error */, addressesFailingEncryption []MEEmailAddress) MEOutgoingMessageEncodingStatus {
	instance := getMEOutgoingMessageEncodingStatusClass().Alloc()
	rv := objc.Send[MEOutgoingMessageEncodingStatus](instance.ID, objc.Sel("initWithCanSign:canEncrypt:securityError:addressesFailingEncryption:"), canSign, canEncrypt, securityError, addressesFailingEncryption)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEOutgoingMessageEncodingStatusWithCanSignCanEncryptSecurityErrorAddressesFailingEncryption */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEOutgoingMessageEncodingStatus */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEOutgoingMessageEncodingStatus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEOutgoingMessageEncodingStatus */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEOutgoingMessageEncodingStatus */

// An array of email addresses that prevent the message security handler from signing the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus/addressesFailingEncryption
func (m_ MEOutgoingMessageEncodingStatus) AddressesFailingEncryption() []MEEmailAddress {
	rv := objc.Send[[]MEEmailAddress](m_.ID, objc.Sel("addressesFailingEncryption"))
	return rv
}/* debug [instance_properties/getter]: addressesFailingEncryption */


// A Boolean value that indicates the message security handler can encrypt the outgoing message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus/canEncrypt
func (m_ MEOutgoingMessageEncodingStatus) CanEncrypt() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canEncrypt"))
	return rv
}/* debug [instance_properties/getter]: canEncrypt */


// A Boolean value that indicates the message security handler can digitally sign the outgoing message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus/canSign
func (m_ MEOutgoingMessageEncodingStatus) CanSign() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canSign"))
	return rv
}/* debug [instance_properties/getter]: canSign */


// An error that the message encoder encountered while determining the encoding status for the outgoing message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEOutgoingMessageEncodingStatus/securityError
func (m_ MEOutgoingMessageEncodingStatus) SecurityError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("securityError"))
	return rv
}/* debug [instance_properties/getter]: securityError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEOutgoingMessageEncodingStatus */


