// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtocol] class.
var (
	URLProtocolClass     _URLProtocolClass
	URLProtocolClassOnce sync.Once
)

func getURLProtocolClass() _URLProtocolClass {
	URLProtocolClassOnce.Do(func() {
		URLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}
	})
	return URLProtocolClass
}

type _URLProtocolClass struct {
	class objc.Class
}

// An interface definition for the [URLProtocol] class.
type IURLProtocol interface {
	objectivec.IObject
}

// An abstract class that handles the loading of protocol-specific URL data.
//
// Don’t instantiate a subclass directly. Instead, create subclasses for any custom protocols or URL schemes that your app supports. When a download starts, the system creates the appropriate protocol object to handle the corresponding URL request. You define your protocol class and call the class method during your app’s launch time so that the system is aware of your protocol. To support the customization of protocol-specific requests, create extensions to the class to provide any custom API that you need. You can store and retrieve protocol-specific request data by using ’s class methods and . Create a for each request your subclass processes successfully. You may want to create a custom class to provide protocol specific information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol
type URLProtocol struct {
	objectivec.Object
}

// URLProtocolFrom constructs a [URLProtocol] from an unsafe.Pointer.
//
// An abstract class that handles the loading of protocol-specific URL data.
func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLProtocolClass) Alloc() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLProtocolClass) New() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLProtocol) Init() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLProtocol) Autorelease() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtocol creates a new URLProtocol instance.
func NewURLProtocol() URLProtocol {
	return getURLProtocolClass().New()
}

// Determines whether the protocol subclass can handle the specified request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/canInit(with:)-76brg
func (uc _URLProtocolClass) CanInitWithRequest(request unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canInitWithRequest:"), request)
	return rv
}

// Returns a canonical version of the specified request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/canonicalRequest(for:)
func (uc _URLProtocolClass) CanonicalRequestForRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("canonicalRequestForRequest:"), request)
	return rv
}

// Fetches the property associated with the specified key in the specified request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/property(forKey:in:)
func (uc _URLProtocolClass) PropertyForKeyInRequest(key string, request unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("propertyForKey:inRequest:"), objc.String(key), request)
	return rv
}

// Attempts to register a subclass of , making it visible to the URL loading system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/registerClass(_:)
func (uc _URLProtocolClass) RegisterClass(protocolClass objc.Class) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("registerClass:"), protocolClass)
	return rv
}

// Removes the property associated with the specified key in the specified request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/removeProperty(forKey:in:)
func (uc _URLProtocolClass) RemovePropertyForKeyInRequest(key string, request unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("removePropertyForKey:inRequest:"), objc.String(key), request)
}

// Sets the property associated with the specified key in the specified request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/setProperty(_:forKey:in:)
func (uc _URLProtocolClass) SetPropertyForKeyInRequest(value objc.ID, key string, request unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("setProperty:forKey:inRequest:"), value, objc.String(key), request)
}

// Unregisters the specified subclass of .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/unregisterClass(_:)
func (uc _URLProtocolClass) UnregisterClass(protocolClass objc.Class) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unregisterClass:"), protocolClass)
}
