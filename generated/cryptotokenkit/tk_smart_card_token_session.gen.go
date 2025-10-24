// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class TKSmartCardTokenSession */


/* debug [class_header]: Header for TKSmartCardTokenSession */
// The class instance for the [TKSmartCardTokenSession] class.
var (
	TKSmartCardTokenSessionClass     _TKSmartCardTokenSessionClass
	TKSmartCardTokenSessionClassOnce sync.Once
)

func getTKSmartCardTokenSessionClass() _TKSmartCardTokenSessionClass {
	TKSmartCardTokenSessionClassOnce.Do(func() {
		TKSmartCardTokenSessionClass = _TKSmartCardTokenSessionClass{objc.GetClass("TKSmartCardTokenSession")}
	})
	return TKSmartCardTokenSessionClass
}

type _TKSmartCardTokenSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardTokenSession */
// An interface definition for the [TKSmartCardTokenSession] class.
type ITKSmartCardTokenSession interface {
	ITKTokenSession
	
/* debug [class_interface_properties]: Properties for TKSmartCardTokenSession */
	// properties:
	SmartCard() ITKSmartCard
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardTokenSession */
	// methods:
	GetSmartCardWithError(error_ unsafe.Pointer) ITKSmartCard
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardTokenSession */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenSessionClass) Alloc() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardTokenSessionClass) New() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenSession) Init() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenSession) Autorelease() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenSession creates a new TKSmartCardTokenSession instance.
func NewTKSmartCardTokenSession() TKSmartCardTokenSession {
	return getTKSmartCardTokenSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardTokenSession */
// A token session that is based on a smart card token.
//
// You can use the property to access and send APDUs to the underlying smart card.


// A token session that is based on a smart card token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession
type TKSmartCardTokenSession struct {
	TKTokenSession
}

// TKSmartCardTokenSessionFrom constructs a [TKSmartCardTokenSession] from an unsafe.Pointer.
//
// A token session that is based on a smart card token.
func TKSmartCardTokenSessionFrom(ptr unsafe.Pointer) TKSmartCardTokenSession {
	return TKSmartCardTokenSession{
		TKTokenSession: TKTokenSessionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardTokenSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardTokenSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardTokenSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardTokenSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/getSmartCard()
func (t_ TKSmartCardTokenSession) GetSmartCardWithError(error_ unsafe.Pointer) ITKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("getSmartCardWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: GetSmartCardWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardTokenSession */

// The smart card for the active exclusive session and selected application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/smartCard
func (t_ TKSmartCardTokenSession) SmartCard() ITKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("smartCard"))
	return rv
}/* debug [instance_properties/getter]: smartCard */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardTokenSession */



