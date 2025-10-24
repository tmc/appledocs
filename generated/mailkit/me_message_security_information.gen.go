// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessageSecurityInformation */


/* debug [class_header]: Header for MEMessageSecurityInformation */
// The class instance for the [MEMessageSecurityInformation] class.
var (
	MEMessageSecurityInformationClass     _MEMessageSecurityInformationClass
	MEMessageSecurityInformationClassOnce sync.Once
)

func getMEMessageSecurityInformationClass() _MEMessageSecurityInformationClass {
	MEMessageSecurityInformationClassOnce.Do(func() {
		MEMessageSecurityInformationClass = _MEMessageSecurityInformationClass{objc.GetClass("MEMessageSecurityInformation")}
	})
	return MEMessageSecurityInformationClass
}

type _MEMessageSecurityInformationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessageSecurityInformation */
// An interface definition for the [MEMessageSecurityInformation] class.
type IMEMessageSecurityInformation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessageSecurityInformation */
	// properties:
	EncryptionError() objc.IObject /* cross-framework: Error */
	IsEncrypted() bool
	LocalizedRemoteContentBlockingReason() objc.IObject /* cross-framework: NSString */
	ShouldBlockRemoteContent() bool
	Signers() []MEMessageSigner
	SigningError() objc.IObject /* cross-framework: Error */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessageSecurityInformation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessageSecurityInformation */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageSecurityInformationClass) Alloc() MEMessageSecurityInformation {
	rv := objc.Send[MEMessageSecurityInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEMessageSecurityInformationClass) New() MEMessageSecurityInformation {
	rv := objc.Send[MEMessageSecurityInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessageSecurityInformation) Init() MEMessageSecurityInformation {
	rv := objc.Send[MEMessageSecurityInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessageSecurityInformation) Autorelease() MEMessageSecurityInformation {
	rv := objc.Send[MEMessageSecurityInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessageSecurityInformation creates a new MEMessageSecurityInformation instance.
func NewMEMessageSecurityInformation() MEMessageSecurityInformation {
	return getMEMessageSecurityInformationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessageSecurityInformation */
// An object that contains details about a message’s content, such as if it’s encrypted and who digitally signed it.


// An object that contains details about a message’s content, such as if it’s encrypted and who digitally signed it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation
type MEMessageSecurityInformation struct {
	objectivec.Object
}

// MEMessageSecurityInformationFrom constructs a [MEMessageSecurityInformation] from an unsafe.Pointer.
//
// An object that contains details about a message’s content, such as if it’s encrypted and who digitally signed it.
func MEMessageSecurityInformationFrom(ptr unsafe.Pointer) MEMessageSecurityInformation {
	return MEMessageSecurityInformation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessageSecurityInformation */

// Creates a message security information object that indicates if a message is encrypted, who signed it, or if an error occurred when decoding the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/init(signers:isEncrypted:signingError:encryptionError:)
func NewMEMessageSecurityInformationWithSignersIsEncryptedSigningErrorEncryptionError(signers []MEMessageSigner, isEncrypted bool, signingError objc.IObject /* cross-framework: Error */, encryptionError objc.IObject /* cross-framework: Error */) MEMessageSecurityInformation {
	instance := getMEMessageSecurityInformationClass().Alloc()
	rv := objc.Send[MEMessageSecurityInformation](instance.ID, objc.Sel("initWithSigners:isEncrypted:signingError:encryptionError:"), signers, isEncrypted, signingError, encryptionError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEMessageSecurityInformationWithSignersIsEncryptedSigningErrorEncryptionError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/init(signers:isEncrypted:signingError:encryptionError:shouldBlockRemoteContent:localizedRemoteContentBlockingReason:)
func NewMEMessageSecurityInformationWithSignersIsEncryptedSigningErrorEncryptionErrorShouldBlockRemoteContentLocalizedRemoteContentBlockingReason(signers []MEMessageSigner, isEncrypted bool, signingError objc.IObject /* cross-framework: Error */, encryptionError objc.IObject /* cross-framework: Error */, shouldBlockRemoteContent bool, localizedRemoteContentBlockingReason objc.IObject /* cross-framework: NSString */) MEMessageSecurityInformation {
	instance := getMEMessageSecurityInformationClass().Alloc()
	rv := objc.Send[MEMessageSecurityInformation](instance.ID, objc.Sel("initWithSigners:isEncrypted:signingError:encryptionError:shouldBlockRemoteContent:localizedRemoteContentBlockingReason:"), signers, isEncrypted, signingError, encryptionError, shouldBlockRemoteContent, localizedRemoteContentBlockingReason)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEMessageSecurityInformationWithSignersIsEncryptedSigningErrorEncryptionErrorShouldBlockRemoteContentLocalizedRemoteContentBlockingReason */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessageSecurityInformation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessageSecurityInformation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessageSecurityInformation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessageSecurityInformation */

// An error that indicates the security handler couldn’t decrypt the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/encryptionError
func (m_ MEMessageSecurityInformation) EncryptionError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("encryptionError"))
	return rv
}/* debug [instance_properties/getter]: encryptionError */


// A Boolean value that indicates if the sender encrypted the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/isEncrypted
func (m_ MEMessageSecurityInformation) IsEncrypted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEncrypted"))
	return rv
}/* debug [instance_properties/getter]: isEncrypted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/localizedRemoteContentBlockingReason
func (m_ MEMessageSecurityInformation) LocalizedRemoteContentBlockingReason() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("localizedRemoteContentBlockingReason"))
	return rv
}/* debug [instance_properties/getter]: localizedRemoteContentBlockingReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/shouldBlockRemoteContent
func (m_ MEMessageSecurityInformation) ShouldBlockRemoteContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldBlockRemoteContent"))
	return rv
}/* debug [instance_properties/getter]: shouldBlockRemoteContent */


// An array of objects that contain information about who signed the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/signers
func (m_ MEMessageSecurityInformation) Signers() []MEMessageSigner {
	rv := objc.Send[[]MEMessageSigner](m_.ID, objc.Sel("signers"))
	return rv
}/* debug [instance_properties/getter]: signers */


// An error that indicates the security handler couldn’t decode the message’s digital signatures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageSecurityInformation/signingError
func (m_ MEMessageSecurityInformation) SigningError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("signingError"))
	return rv
}/* debug [instance_properties/getter]: signingError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessageSecurityInformation */


