// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLConnection] class.
var (
	URLConnectionClass     _URLConnectionClass
	URLConnectionClassOnce sync.Once
)

func getURLConnectionClass() _URLConnectionClass {
	URLConnectionClassOnce.Do(func() {
		URLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}
	})
	return URLConnectionClass
}

type _URLConnectionClass struct {
	class objc.Class
}

// An interface definition for the [URLConnection] class.
type IURLConnection interface {
	objectivec.IObject
	// properties:
	CurrentRequest() objc.IObject /* cross-framework: URLRequest */
	SetCurrentRequest(value objc.IObject /* cross-framework: URLRequest */)
	OriginalRequest() objc.IObject /* cross-framework: URLRequest */
	SetOriginalRequest(value objc.IObject /* cross-framework: URLRequest */)
	// methods:
}

// An object that enables you to start and stop URL requests.
//
// An object lets you load the contents of a URL by providing a URL request object. The interface for is sparse, providing only the controls to start and cancel asynchronous loads of a URL request. You perform most of your configuration on the URL request object itself. The class provides convenience class methods to load URL requests both asynchronously using a callback block and synchronously. For greater control, you can create a URL connection object with a delegate object that conforms to the and protocols. The connection calls methods on that delegate to provide you with progress and status as the URL request is loaded asynchronously. The connection also calls delegate methods to let you override the connection’s default behavior (for example, specifying how a particular redirect should be handled). These delegate methods are called on the thread that initiated the asynchronous load operation. For more information about errors, see the header, , and URL Loading System Error Codes in .


// An object that enables you to start and stop URL requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection
type URLConnection struct {
	objectivec.Object
}

// URLConnectionFrom constructs a [URLConnection] from an unsafe.Pointer.
//
// An object that enables you to start and stop URL requests.
func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLConnectionClass) Alloc() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLConnectionClass) New() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLConnection) Init() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLConnection) Autorelease() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLConnection creates a new URLConnection instance.
func NewURLConnection() URLConnection {
	return getURLConnectionClass().New()
}



// The current connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/currentrequest
func (u_ URLConnection) CurrentRequest() objc.IObject /* cross-framework: URLRequest */ {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("currentRequest"))
	return rv
}


// The current connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/currentrequest
func (u_ URLConnection) SetCurrentRequest(value objc.IObject /* cross-framework: URLRequest */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentRequest:"), value)
}


// A deep copy of the original connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/originalrequest
func (u_ URLConnection) OriginalRequest() objc.IObject /* cross-framework: URLRequest */ {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("originalRequest"))
	return rv
}


// A deep copy of the original connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/originalrequest
func (u_ URLConnection) SetOriginalRequest(value objc.IObject /* cross-framework: URLRequest */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOriginalRequest:"), value)
}



