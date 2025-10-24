// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class TKSmartCardUserInteractionForSecurePINChange */


/* debug [class_header]: Header for TKSmartCardUserInteractionForSecurePINChange */
// The class instance for the [TKSmartCardUserInteractionForSecurePINChange] class.
var (
	TKSmartCardUserInteractionForSecurePINChangeClass     _TKSmartCardUserInteractionForSecurePINChangeClass
	TKSmartCardUserInteractionForSecurePINChangeClassOnce sync.Once
)

func getTKSmartCardUserInteractionForSecurePINChangeClass() _TKSmartCardUserInteractionForSecurePINChangeClass {
	TKSmartCardUserInteractionForSecurePINChangeClassOnce.Do(func() {
		TKSmartCardUserInteractionForSecurePINChangeClass = _TKSmartCardUserInteractionForSecurePINChangeClass{objc.GetClass("TKSmartCardUserInteractionForSecurePINChange")}
	})
	return TKSmartCardUserInteractionForSecurePINChangeClass
}

type _TKSmartCardUserInteractionForSecurePINChangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardUserInteractionForSecurePINChange */
// An interface definition for the [TKSmartCardUserInteractionForSecurePINChange] class.
type ITKSmartCardUserInteractionForSecurePINChange interface {
	ITKSmartCardUserInteractionForPINOperation
	
/* debug [class_interface_properties]: Properties for TKSmartCardUserInteractionForSecurePINChange */
	// properties:
	PINConfirmation() TKSmartCardPINConfirmation
	SetPINConfirmation(value TKSmartCardPINConfirmation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardUserInteractionForSecurePINChange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardUserInteractionForSecurePINChange */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardUserInteractionForSecurePINChangeClass) Alloc() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardUserInteractionForSecurePINChangeClass) New() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardUserInteractionForSecurePINChange) Init() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardUserInteractionForSecurePINChange) Autorelease() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForSecurePINChange creates a new TKSmartCardUserInteractionForSecurePINChange instance.
func NewTKSmartCardUserInteractionForSecurePINChange() TKSmartCardUserInteractionForSecurePINChange {
	return getTKSmartCardUserInteractionForSecurePINChangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardUserInteractionForSecurePINChange */
// A representation of the user interaction for secure PIN change operations on a Smart Card reader.
//
// The result of a user interaction is available once the interaction has completed.


// A representation of the user interaction for secure PIN change operations on a Smart Card reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange
type TKSmartCardUserInteractionForSecurePINChange struct {
	TKSmartCardUserInteractionForPINOperation
}

// TKSmartCardUserInteractionForSecurePINChangeFrom constructs a [TKSmartCardUserInteractionForSecurePINChange] from an unsafe.Pointer.
//
// A representation of the user interaction for secure PIN change operations on a Smart Card reader.
func TKSmartCardUserInteractionForSecurePINChangeFrom(ptr unsafe.Pointer) TKSmartCardUserInteractionForSecurePINChange {
	return TKSmartCardUserInteractionForSecurePINChange{
		TKSmartCardUserInteractionForPINOperation: TKSmartCardUserInteractionForPINOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardUserInteractionForSecurePINChange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardUserInteractionForSecurePINChange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardUserInteractionForSecurePINChange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardUserInteractionForSecurePINChange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardUserInteractionForSecurePINChange */

// The way PIN confirmation is requested. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/pinConfirmation
func (t_ TKSmartCardUserInteractionForSecurePINChange) PINConfirmation() TKSmartCardPINConfirmation {
	rv := objc.Send[TKSmartCardPINConfirmation](t_.ID, objc.Sel("PINConfirmation"))
	return rv
}/* debug [instance_properties/getter]: PINConfirmation */


// The way PIN confirmation is requested. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/pinConfirmation
func (t_ TKSmartCardUserInteractionForSecurePINChange) SetPINConfirmation(value TKSmartCardPINConfirmation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINConfirmation:"), value)
}/* debug [instance_properties/setter]: PINConfirmation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardUserInteractionForSecurePINChange */



