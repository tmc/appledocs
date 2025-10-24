// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class TKSmartCardUserInteractionForSecurePINVerification */


/* debug [class_header]: Header for TKSmartCardUserInteractionForSecurePINVerification */
// The class instance for the [TKSmartCardUserInteractionForSecurePINVerification] class.
var (
	TKSmartCardUserInteractionForSecurePINVerificationClass     _TKSmartCardUserInteractionForSecurePINVerificationClass
	TKSmartCardUserInteractionForSecurePINVerificationClassOnce sync.Once
)

func getTKSmartCardUserInteractionForSecurePINVerificationClass() _TKSmartCardUserInteractionForSecurePINVerificationClass {
	TKSmartCardUserInteractionForSecurePINVerificationClassOnce.Do(func() {
		TKSmartCardUserInteractionForSecurePINVerificationClass = _TKSmartCardUserInteractionForSecurePINVerificationClass{objc.GetClass("TKSmartCardUserInteractionForSecurePINVerification")}
	})
	return TKSmartCardUserInteractionForSecurePINVerificationClass
}

type _TKSmartCardUserInteractionForSecurePINVerificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardUserInteractionForSecurePINVerification */
// An interface definition for the [TKSmartCardUserInteractionForSecurePINVerification] class.
type ITKSmartCardUserInteractionForSecurePINVerification interface {
	ITKSmartCardUserInteractionForPINOperation
	
/* debug [class_interface_properties]: Properties for TKSmartCardUserInteractionForSecurePINVerification */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardUserInteractionForSecurePINVerification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardUserInteractionForSecurePINVerification */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardUserInteractionForSecurePINVerificationClass) Alloc() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardUserInteractionForSecurePINVerificationClass) New() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardUserInteractionForSecurePINVerification) Init() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardUserInteractionForSecurePINVerification) Autorelease() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForSecurePINVerification creates a new TKSmartCardUserInteractionForSecurePINVerification instance.
func NewTKSmartCardUserInteractionForSecurePINVerification() TKSmartCardUserInteractionForSecurePINVerification {
	return getTKSmartCardUserInteractionForSecurePINVerificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardUserInteractionForSecurePINVerification */
// A representation of the user interaction for secure PIN change verification on a Smart Card reader.
//
// The result of a user interaction is available once the interaction has completed.


// A representation of the user interaction for secure PIN change verification on a Smart Card reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINVerification
type TKSmartCardUserInteractionForSecurePINVerification struct {
	TKSmartCardUserInteractionForPINOperation
}

// TKSmartCardUserInteractionForSecurePINVerificationFrom constructs a [TKSmartCardUserInteractionForSecurePINVerification] from an unsafe.Pointer.
//
// A representation of the user interaction for secure PIN change verification on a Smart Card reader.
func TKSmartCardUserInteractionForSecurePINVerificationFrom(ptr unsafe.Pointer) TKSmartCardUserInteractionForSecurePINVerification {
	return TKSmartCardUserInteractionForSecurePINVerification{
		TKSmartCardUserInteractionForPINOperation: TKSmartCardUserInteractionForPINOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardUserInteractionForSecurePINVerification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardUserInteractionForSecurePINVerification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardUserInteractionForSecurePINVerification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardUserInteractionForSecurePINVerification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardUserInteractionForSecurePINVerification */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardUserInteractionForSecurePINVerification */



