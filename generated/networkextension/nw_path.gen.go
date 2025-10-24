// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NWPath */


/* debug [class_header]: Header for NWPath */
// The class instance for the [NWPath] class.
var (
	NWPathClass     _NWPathClass
	NWPathClassOnce sync.Once
)

func getNWPathClass() _NWPathClass {
	NWPathClassOnce.Do(func() {
		NWPathClass = _NWPathClass{objc.GetClass("NWPath")}
	})
	return NWPathClass
}

type _NWPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NWPath */
// An interface definition for the [NWPath] class.
type INWPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NWPath */
	// properties:
	Constrained() bool
	Expensive() bool
	Status() NWPathStatus
	IsConstrained() bool
	SetIsConstrained(value bool)
	IsExpensive() bool
	SetIsExpensive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NWPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NWPath */
// Alloc allocates a new instance without initialization.
func (nc _NWPathClass) Alloc() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NWPathClass) New() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWPath) Init() NWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWPath) Autorelease() NWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWPath creates a new NWPath instance.
func NewNWPath() NWPath {
	return getNWPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NWPath */
// The path made by a network connection, including information about its viability.
//
// For example, if the path status is , then a connection attempt will be made. When attached to a specific connection, a path takes all of the connection parameters into account. For example, if the route for a connection changes or is removed, the path will reflect that change. Note that every path is evaluated within the context of the process it is running in, and may be different across processes. is a static object, and properties of the path will never change. To monitor changing network status, use Key-Value Observing (KVO) to watch a path property on another object. For information about KVO, see .


// The path made by a network connection, including information about its viability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPath
type NWPath struct {
	objectivec.Object
}

// NWPathFrom constructs a [NWPath] from an unsafe.Pointer.
//
// The path made by a network connection, including information about its viability.
func NWPathFrom(ptr unsafe.Pointer) NWPath {
	return NWPath{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NWPath *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NWPath */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NWPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NWPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NWPath */

// A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPath/isConstrained
func (n_ NWPath) Constrained() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("constrained"))
	return rv
}/* debug [instance_properties/getter]: constrained */


// A Boolean that indicates whether or not the path uses an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPath/isExpensive
func (n_ NWPath) Expensive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("expensive"))
	return rv
}/* debug [instance_properties/getter]: expensive */


// The evaluated status of the network path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPath/status
func (n_ NWPath) Status() NWPathStatus {
	rv := objc.Send[NWPathStatus](n_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwpath/isconstrained
func (n_ NWPath) IsConstrained() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isConstrained"))
	return rv
}/* debug [instance_properties/getter]: isConstrained */


// A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwpath/isconstrained
func (n_ NWPath) SetIsConstrained(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsConstrained:"), value)
}/* debug [instance_properties/setter]: isConstrained */


// A Boolean that indicates whether or not the path uses an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwpath/isexpensive
func (n_ NWPath) IsExpensive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isExpensive"))
	return rv
}/* debug [instance_properties/getter]: isExpensive */


// A Boolean that indicates whether or not the path uses an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwpath/isexpensive
func (n_ NWPath) SetIsExpensive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsExpensive:"), value)
}/* debug [instance_properties/setter]: isExpensive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NWPath */



