// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXPCCoder */


/* debug [class_header]: Header for NSXPCCoder */
// The class instance for the [XPCCoder] class.
var (
	XPCCoderClass     _XPCCoderClass
	XPCCoderClassOnce sync.Once
)

func getXPCCoderClass() _XPCCoderClass {
	XPCCoderClassOnce.Do(func() {
		XPCCoderClass = _XPCCoderClass{objc.GetClass("NSXPCCoder")}
	})
	return XPCCoderClass
}

type _XPCCoderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XPCCoder */
// An interface definition for the [XPCCoder] class.
type IXPCCoder interface {
	ICoder
	
/* debug [class_interface_properties]: Properties for XPCCoder */
	// properties:
	Connection() IXPCConnection
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XPCCoder */
	// methods:
	DecodeXPCObjectOfTypeForKey(type_ objectivec.IObject, key IString) objectivec.IObject
	EncodeXPCObjectForKey(xpcObject objectivec.IObject, key IString)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XPCCoder */
// Alloc allocates a new instance without initialization.
func (xc _XPCCoderClass) Alloc() XPCCoder {
	rv := objc.Send[XPCCoder](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (xc _XPCCoderClass) New() XPCCoder {
	rv := objc.Send[XPCCoder](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCCoder) Init() XPCCoder {
	rv := objc.Send[XPCCoder](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCCoder) Autorelease() XPCCoder {
	rv := objc.Send[XPCCoder](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCCoder creates a new XPCCoder instance.
func NewXPCCoder() XPCCoder {
	return getXPCCoderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XPCCoder */
// A coder that encodes and decodes objects that your app sends over an XPC connection.
//
// If you want to perform custom encoding or decoding of objects that your app sends over an , use to determine if the coder provided to your object is a kind of .


// A coder that encodes and decodes objects that your app sends over an XPC connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder
type XPCCoder struct {
	Coder
}

// XPCCoderFrom constructs a [XPCCoder] from an unsafe.Pointer.
//
// A coder that encodes and decodes objects that your app sends over an XPC connection.
func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		Coder: CoderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XPCCoder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XPCCoder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XPCCoder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XPCCoder */

// Decodes an object and validates that its type matches the type a service provides over XPC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/decodeXPCObject(ofType:forKey:)
func (x_ XPCCoder) DecodeXPCObjectOfTypeForKey(type_ objectivec.IObject, key IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("decodeXPCObjectOfType:forKey:"), type_, key)
	return rv
}/* debug [instance_methods/method]: DecodeXPCObjectOfTypeForKey */


// Encodes an object to send over an XPC connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/encodeXPCObject(_:forKey:)
func (x_ XPCCoder) EncodeXPCObjectForKey(xpcObject objectivec.IObject, key IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("encodeXPCObject:forKey:"), xpcObject, key)
}/* debug [instance_methods/method]: EncodeXPCObjectForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XPCCoder */

// The connection currently performing encoding or decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/connection
func (x_ XPCCoder) Connection() IXPCConnection {
	rv := objc.Send[XPCConnection](x_.ID, objc.Sel("connection"))
	return rv
}/* debug [instance_properties/getter]: connection */


// An optional user information object associated with the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/userInfo
func (x_ XPCCoder) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// An optional user information object associated with the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/userInfo
func (x_ XPCCoder) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXPCCoder */



