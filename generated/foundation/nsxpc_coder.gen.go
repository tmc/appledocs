// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [XPCCoder] class.
type IXPCCoder interface {
	ICoder
	EncodeXPCObjectForKey(xpcObject unsafe.Pointer, key string)
	Connection() NSXPCConnection
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
}

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

// Alloc allocates a new instance without initialization.
func (xc _XPCCoderClass) Alloc() XPCCoder {
	rv := objc.Send[XPCCoder](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Encodes an object to send over an XPC connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/encodeXPCObject(_:forKey:)

func (x_ XPCCoder) EncodeXPCObjectForKey(xpcObject unsafe.Pointer, key string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("encodeXPCObject:forKey:"), xpcObject, objc.String(key))
}


// The connection currently performing encoding or decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/connection

func (x_ XPCCoder) Connection() NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x_.ID, objc.Sel("connection"))
	return rv
}


// An optional user information object associated with the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/userInfo

func (x_ XPCCoder) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("userInfo"))
	return rv
}


// An optional user information object associated with the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder/userInfo

func (x_ XPCCoder) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setUserInfo:"), value)
}



