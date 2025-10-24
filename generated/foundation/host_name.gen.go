// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hostName */


/* debug [class_header]: Header for hostName */
// The class instance for the [hostName] class.
var (
	HostNameClass     _hostNameClass
	HostNameClassOnce sync.Once
)

func gethostNameClass() _hostNameClass {
	HostNameClassOnce.Do(func() {
		HostNameClass = _hostNameClass{objc.GetClass("hostName")}
	})
	return HostNameClass
}

type _hostNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hostName */
// An interface definition for the [hostName] class.
type IhostName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hostName */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hostName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hostName */
// Alloc allocates a new instance without initialization.
func (hc _hostNameClass) Alloc() hostName {
	rv := objc.Send[hostName](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hostNameClass) New() hostName {
	rv := objc.Send[hostName](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hostName) Init() hostName {
	rv := objc.Send[hostName](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hostName) Autorelease() hostName {
	rv := objc.Send[hostName](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhostName creates a new hostName instance.
func NewhostName() hostName {
	return gethostNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hostName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/hostName-c.ivar
type hostName struct {
	objectivec.Object
}

// hostNameFrom constructs a [hostName] from an unsafe.Pointer.
func hostNameFrom(ptr unsafe.Pointer) hostName {
	return hostName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hostName *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hostName */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hostName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hostName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hostName */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hostName */



