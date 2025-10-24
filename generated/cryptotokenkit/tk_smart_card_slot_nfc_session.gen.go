// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardSlotNFCSession */


/* debug [class_header]: Header for TKSmartCardSlotNFCSession */
// The class instance for the [TKSmartCardSlotNFCSession] class.
var (
	TKSmartCardSlotNFCSessionClass     _TKSmartCardSlotNFCSessionClass
	TKSmartCardSlotNFCSessionClassOnce sync.Once
)

func getTKSmartCardSlotNFCSessionClass() _TKSmartCardSlotNFCSessionClass {
	TKSmartCardSlotNFCSessionClassOnce.Do(func() {
		TKSmartCardSlotNFCSessionClass = _TKSmartCardSlotNFCSessionClass{objc.GetClass("TKSmartCardSlotNFCSession")}
	})
	return TKSmartCardSlotNFCSessionClass
}

type _TKSmartCardSlotNFCSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardSlotNFCSession */
// An interface definition for the [TKSmartCardSlotNFCSession] class.
type ITKSmartCardSlotNFCSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardSlotNFCSession */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardSlotNFCSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardSlotNFCSession */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotNFCSessionClass) Alloc() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardSlotNFCSessionClass) New() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlotNFCSession) Init() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlotNFCSession) Autorelease() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlotNFCSession creates a new TKSmartCardSlotNFCSession instance.
func NewTKSmartCardSlotNFCSession() TKSmartCardSlotNFCSession {
	return getTKSmartCardSlotNFCSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardSlotNFCSession */
// NFC session that’s related to NFC smart card slot which was created.
//
// Lifetime of this session object is tied to the NFC smart card slot lifetime and once the NFC slot disappears (eg. after a user cancellation, calling end session, or an NFC timeout) the functions will start to fail and return error.


// NFC session that’s related to NFC smart card slot which was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession
type TKSmartCardSlotNFCSession struct {
	objectivec.Object
}

// TKSmartCardSlotNFCSessionFrom constructs a [TKSmartCardSlotNFCSession] from an unsafe.Pointer.
//
// NFC session that’s related to NFC smart card slot which was created.
func TKSmartCardSlotNFCSessionFrom(ptr unsafe.Pointer) TKSmartCardSlotNFCSession {
	return TKSmartCardSlotNFCSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardSlotNFCSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardSlotNFCSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardSlotNFCSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardSlotNFCSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardSlotNFCSession */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardSlotNFCSession */


