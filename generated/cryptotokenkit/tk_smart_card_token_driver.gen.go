// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class TKSmartCardTokenDriver */


/* debug [class_header]: Header for TKSmartCardTokenDriver */
// The class instance for the [TKSmartCardTokenDriver] class.
var (
	TKSmartCardTokenDriverClass     _TKSmartCardTokenDriverClass
	TKSmartCardTokenDriverClassOnce sync.Once
)

func getTKSmartCardTokenDriverClass() _TKSmartCardTokenDriverClass {
	TKSmartCardTokenDriverClassOnce.Do(func() {
		TKSmartCardTokenDriverClass = _TKSmartCardTokenDriverClass{objc.GetClass("TKSmartCardTokenDriver")}
	})
	return TKSmartCardTokenDriverClass
}

type _TKSmartCardTokenDriverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardTokenDriver */
// An interface definition for the [TKSmartCardTokenDriver] class.
type ITKSmartCardTokenDriver interface {
	ITKTokenDriver
	
/* debug [class_interface_properties]: Properties for TKSmartCardTokenDriver */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardTokenDriver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardTokenDriver */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenDriverClass) Alloc() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardTokenDriverClass) New() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenDriver) Init() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenDriver) Autorelease() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenDriver creates a new TKSmartCardTokenDriver instance.
func NewTKSmartCardTokenDriver() TKSmartCardTokenDriver {
	return getTKSmartCardTokenDriverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardTokenDriver */
// The driver that acts as an entry point for smart card app extensions.


// The driver that acts as an entry point for smart card app extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriver
type TKSmartCardTokenDriver struct {
	TKTokenDriver
}

// TKSmartCardTokenDriverFrom constructs a [TKSmartCardTokenDriver] from an unsafe.Pointer.
//
// The driver that acts as an entry point for smart card app extensions.
func TKSmartCardTokenDriverFrom(ptr unsafe.Pointer) TKSmartCardTokenDriver {
	return TKSmartCardTokenDriver{
		TKTokenDriver: TKTokenDriverFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardTokenDriver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardTokenDriver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardTokenDriver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardTokenDriver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardTokenDriver */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardTokenDriver */



