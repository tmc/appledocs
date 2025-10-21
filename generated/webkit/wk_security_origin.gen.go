// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SecurityOrigin] class.
type ISecurityOrigin interface {
	objectivec.IObject
}

// An object that identifies the origin of a particular resource.
//
// A object is a transient, data-only object that identifies the host name, protocol, and port number associated with a particular resource. You don’t create objects directly. Instead, WebKit creates them for the resources it loads. A load is any load URL has the same security origin as the requesting web site. First-party webpages can access each other’s resources, such as scripts and databases. Because a object is transient, it doesn’t uniquely identify a security origin across multiple delegate method calls.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SecurityOriginClass) Alloc() SecurityOrigin {
	rv := objc.Send[SecurityOrigin](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The security origin’s port.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/port
func (s_ SecurityOrigin) Port() int {
	rv := objc.Send[int](s_.ID, objc.Sel("port"))
	return rv
}


// SetPort sets the value of the port property.
// The security origin’s port.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/port
func (s_ SecurityOrigin) SetPort(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPort:"), value)
}

// The security origin’s protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/protocol
func (s_ SecurityOrigin) `protocol`() string {
	rv := objc.Send[string](s_.ID, objc.Sel("`protocol`"))
	return rv
}


// Set`protocol` sets the value of the `protocol` property.
// The security origin’s protocol.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/protocol
func (s_ SecurityOrigin) Set`protocol`(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("set`protocol`:"), objc.String(value))
}

// The security origin’s host.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/host
func (s_ SecurityOrigin) Host() string {
	rv := objc.Send[string](s_.ID, objc.Sel("host"))
	return rv
}



