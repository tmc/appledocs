// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardATRInterfaceGroup */


/* debug [class_header]: Header for TKSmartCardATRInterfaceGroup */
// The class instance for the [TKSmartCardATRInterfaceGroup] class.
var (
	TKSmartCardATRInterfaceGroupClass     _TKSmartCardATRInterfaceGroupClass
	TKSmartCardATRInterfaceGroupClassOnce sync.Once
)

func getTKSmartCardATRInterfaceGroupClass() _TKSmartCardATRInterfaceGroupClass {
	TKSmartCardATRInterfaceGroupClassOnce.Do(func() {
		TKSmartCardATRInterfaceGroupClass = _TKSmartCardATRInterfaceGroupClass{objc.GetClass("TKSmartCardATRInterfaceGroup")}
	})
	return TKSmartCardATRInterfaceGroupClass
}

type _TKSmartCardATRInterfaceGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardATRInterfaceGroup */
// An interface definition for the [TKSmartCardATRInterfaceGroup] class.
type ITKSmartCardATRInterfaceGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardATRInterfaceGroup */
	// properties:
	Protocol() objc.IObject /* cross-framework: NSNumber */
	TA() objc.IObject /* cross-framework: NSNumber */
	TB() objc.IObject /* cross-framework: NSNumber */
	TC() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardATRInterfaceGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardATRInterfaceGroup */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardATRInterfaceGroupClass) Alloc() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardATRInterfaceGroupClass) New() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardATRInterfaceGroup) Init() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardATRInterfaceGroup) Autorelease() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardATRInterfaceGroup creates a new TKSmartCardATRInterfaceGroup instance.
func NewTKSmartCardATRInterfaceGroup() TKSmartCardATRInterfaceGroup {
	return getTKSmartCardATRInterfaceGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardATRInterfaceGroup */
// A single interface-bytes group for a Smart Card ATR (Answer to Reset).
//
// You access instances of this class by calling the and methods on an object.


// A single interface-bytes group for a Smart Card ATR (Answer to Reset).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup
type TKSmartCardATRInterfaceGroup struct {
	objectivec.Object
}

// TKSmartCardATRInterfaceGroupFrom constructs a [TKSmartCardATRInterfaceGroup] from an unsafe.Pointer.
//
// A single interface-bytes group for a Smart Card ATR (Answer to Reset).
func TKSmartCardATRInterfaceGroupFrom(ptr unsafe.Pointer) TKSmartCardATRInterfaceGroup {
	return TKSmartCardATRInterfaceGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardATRInterfaceGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardATRInterfaceGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardATRInterfaceGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardATRInterfaceGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardATRInterfaceGroup */

// The protocol for this group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/protocol
func (t_ TKSmartCardATRInterfaceGroup) Protocol() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](t_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// The TA interface byte of ATR group, or if TA is not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/ta
func (t_ TKSmartCardATRInterfaceGroup) TA() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](t_.ID, objc.Sel("TA"))
	return rv
}/* debug [instance_properties/getter]: TA */


// The TB interface byte of ATR group, or if TB is not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/tb
func (t_ TKSmartCardATRInterfaceGroup) TB() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](t_.ID, objc.Sel("TB"))
	return rv
}/* debug [instance_properties/getter]: TB */


// The TC interface byte of ATR group, or if TC is not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/tc
func (t_ TKSmartCardATRInterfaceGroup) TC() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](t_.ID, objc.Sel("TC"))
	return rv
}/* debug [instance_properties/getter]: TC */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardATRInterfaceGroup */



