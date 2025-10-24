// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKSecurityOrigin */


/* debug [class_header]: Header for WKSecurityOrigin */
// The class instance for the [SecurityOrigin] class.
var (
	SecurityOriginClass     _SecurityOriginClass
	SecurityOriginClassOnce sync.Once
)

func getSecurityOriginClass() _SecurityOriginClass {
	SecurityOriginClassOnce.Do(func() {
		SecurityOriginClass = _SecurityOriginClass{objc.GetClass("WKSecurityOrigin")}
	})
	return SecurityOriginClass
}

type _SecurityOriginClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SecurityOrigin */
// An interface definition for the [SecurityOrigin] class.
type ISecurityOrigin interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SecurityOrigin */
	// properties:
	Host() objc.IObject /* cross-framework: NSString */
	Port() int
	Protocol() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SecurityOrigin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SecurityOrigin */
// Alloc allocates a new instance without initialization.
func (sc _SecurityOriginClass) Alloc() SecurityOrigin {
	rv := objc.Send[SecurityOrigin](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SecurityOriginClass) New() SecurityOrigin {
	rv := objc.Send[SecurityOrigin](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecurityOrigin) Init() SecurityOrigin {
	rv := objc.Send[SecurityOrigin](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecurityOrigin) Autorelease() SecurityOrigin {
	rv := objc.Send[SecurityOrigin](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecurityOrigin creates a new SecurityOrigin instance.
func NewSecurityOrigin() SecurityOrigin {
	return getSecurityOriginClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SecurityOrigin */
// An object that identifies the origin of a particular resource.
//
// A object is a transient, data-only object that identifies the host name, protocol, and port number associated with a particular resource. You don’t create objects directly. Instead, WebKit creates them for the resources it loads. A load is any load URL has the same security origin as the requesting web site. First-party webpages can access each other’s resources, such as scripts and databases. Because a object is transient, it doesn’t uniquely identify a security origin across multiple delegate method calls.


// An object that identifies the origin of a particular resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin
type SecurityOrigin struct {
	objectivec.Object
}

// SecurityOriginFrom constructs a [SecurityOrigin] from an unsafe.Pointer.
//
// An object that identifies the origin of a particular resource.
func SecurityOriginFrom(ptr unsafe.Pointer) SecurityOrigin {
	return SecurityOrigin{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SecurityOrigin *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SecurityOrigin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SecurityOrigin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SecurityOrigin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SecurityOrigin */

// The security origin’s host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/host
func (s_ SecurityOrigin) Host() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("host"))
	return rv
}/* debug [instance_properties/getter]: host */


// The security origin’s port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/port
func (s_ SecurityOrigin) Port() int {
	rv := objc.Send[int](s_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// The security origin’s protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/protocol
func (s_ SecurityOrigin) Protocol() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKSecurityOrigin */



