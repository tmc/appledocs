// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenWatcherTokenInfo */


/* debug [class_header]: Header for TKTokenWatcherTokenInfo */
// The class instance for the [TKTokenWatcherTokenInfo] class.
var (
	TKTokenWatcherTokenInfoClass     _TKTokenWatcherTokenInfoClass
	TKTokenWatcherTokenInfoClassOnce sync.Once
)

func getTKTokenWatcherTokenInfoClass() _TKTokenWatcherTokenInfoClass {
	TKTokenWatcherTokenInfoClassOnce.Do(func() {
		TKTokenWatcherTokenInfoClass = _TKTokenWatcherTokenInfoClass{objc.GetClass("TKTokenWatcherTokenInfo")}
	})
	return TKTokenWatcherTokenInfoClass
}

type _TKTokenWatcherTokenInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenWatcherTokenInfo */
// An interface definition for the [TKTokenWatcherTokenInfo] class.
type ITKTokenWatcherTokenInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenWatcherTokenInfo */
	// properties:
	DriverName() objc.IObject /* cross-framework: NSString */
	SlotName() objc.IObject /* cross-framework: NSString */
	TokenID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenWatcherTokenInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenWatcherTokenInfo */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenWatcherTokenInfoClass) Alloc() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenWatcherTokenInfoClass) New() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenWatcherTokenInfo) Init() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenWatcherTokenInfo) Autorelease() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenWatcherTokenInfo creates a new TKTokenWatcherTokenInfo instance.
func NewTKTokenWatcherTokenInfo() TKTokenWatcherTokenInfo {
	return getTKTokenWatcherTokenInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenWatcherTokenInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo
type TKTokenWatcherTokenInfo struct {
	objectivec.Object
}

// TKTokenWatcherTokenInfoFrom constructs a [TKTokenWatcherTokenInfo] from an unsafe.Pointer.
func TKTokenWatcherTokenInfoFrom(ptr unsafe.Pointer) TKTokenWatcherTokenInfo {
	return TKTokenWatcherTokenInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenWatcherTokenInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenWatcherTokenInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenWatcherTokenInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenWatcherTokenInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenWatcherTokenInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/driverName
func (t_ TKTokenWatcherTokenInfo) DriverName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("driverName"))
	return rv
}/* debug [instance_properties/getter]: driverName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/slotName
func (t_ TKTokenWatcherTokenInfo) SlotName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("slotName"))
	return rv
}/* debug [instance_properties/getter]: slotName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/tokenID
func (t_ TKTokenWatcherTokenInfo) TokenID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("tokenID"))
	return rv
}/* debug [instance_properties/getter]: tokenID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenWatcherTokenInfo */



