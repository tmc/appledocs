// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessageEncodingResult */


/* debug [class_header]: Header for MEMessageEncodingResult */
// The class instance for the [MEMessageEncodingResult] class.
var (
	MEMessageEncodingResultClass     _MEMessageEncodingResultClass
	MEMessageEncodingResultClassOnce sync.Once
)

func getMEMessageEncodingResultClass() _MEMessageEncodingResultClass {
	MEMessageEncodingResultClassOnce.Do(func() {
		MEMessageEncodingResultClass = _MEMessageEncodingResultClass{objc.GetClass("MEMessageEncodingResult")}
	})
	return MEMessageEncodingResultClass
}

type _MEMessageEncodingResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessageEncodingResult */
// An interface definition for the [MEMessageEncodingResult] class.
type IMEMessageEncodingResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessageEncodingResult */
	// properties:
	EncodedMessage() IMEEncodedOutgoingMessage
	EncryptionError() objc.IObject /* cross-framework: Error */
	SigningError() objc.IObject /* cross-framework: Error */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessageEncodingResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessageEncodingResult */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageEncodingResultClass) Alloc() MEMessageEncodingResult {
	rv := objc.Send[MEMessageEncodingResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEMessageEncodingResultClass) New() MEMessageEncodingResult {
	rv := objc.Send[MEMessageEncodingResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessageEncodingResult) Init() MEMessageEncodingResult {
	rv := objc.Send[MEMessageEncodingResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessageEncodingResult) Autorelease() MEMessageEncodingResult {
	rv := objc.Send[MEMessageEncodingResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessageEncodingResult creates a new MEMessageEncodingResult instance.
func NewMEMessageEncodingResult() MEMessageEncodingResult {
	return getMEMessageEncodingResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessageEncodingResult */
// An object that contains a signed or encrypted message, or errors that indicate failure to encode the message.


// An object that contains a signed or encrypted message, or errors that indicate failure to encode the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncodingResult
type MEMessageEncodingResult struct {
	objectivec.Object
}

// MEMessageEncodingResultFrom constructs a [MEMessageEncodingResult] from an unsafe.Pointer.
//
// An object that contains a signed or encrypted message, or errors that indicate failure to encode the message.
func MEMessageEncodingResultFrom(ptr unsafe.Pointer) MEMessageEncodingResult {
	return MEMessageEncodingResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessageEncodingResult */

// Creates an encoding result object with a signed or encrypted message, or errors if the message encoder fails to encode the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncodingResult/init(encodedMessage:signingError:encryptionError:)
func NewMEMessageEncodingResultWithEncodedMessageSigningErrorEncryptionError(encodedMessage IMEEncodedOutgoingMessage, signingError objc.IObject /* cross-framework: Error */, encryptionError objc.IObject /* cross-framework: Error */) MEMessageEncodingResult {
	instance := getMEMessageEncodingResultClass().Alloc()
	rv := objc.Send[MEMessageEncodingResult](instance.ID, objc.Sel("initWithEncodedMessage:signingError:encryptionError:"), encodedMessage, signingError, encryptionError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEMessageEncodingResultWithEncodedMessageSigningErrorEncryptionError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessageEncodingResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessageEncodingResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessageEncodingResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessageEncodingResult */

// A signed or encrypted message, if the message security handler needs to encode the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncodingResult/encodedMessage
func (m_ MEMessageEncodingResult) EncodedMessage() IMEEncodedOutgoingMessage {
	rv := objc.Send[MEEncodedOutgoingMessage](m_.ID, objc.Sel("encodedMessage"))
	return rv
}/* debug [instance_properties/getter]: encodedMessage */


// An error that occurred while the message encoder encrypted the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncodingResult/encryptionError
func (m_ MEMessageEncodingResult) EncryptionError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("encryptionError"))
	return rv
}/* debug [instance_properties/getter]: encryptionError */


// An error that occurred while the message encoder signed the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageEncodingResult/signingError
func (m_ MEMessageEncodingResult) SigningError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("signingError"))
	return rv
}/* debug [instance_properties/getter]: signingError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessageEncodingResult */


