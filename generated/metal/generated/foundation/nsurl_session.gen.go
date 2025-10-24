// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSession] class.
var (
	URLSessionClass     _URLSessionClass
	URLSessionClassOnce sync.Once
)

func getURLSessionClass() _URLSessionClass {
	URLSessionClassOnce.Do(func() {
		URLSessionClass = _URLSessionClass{objc.GetClass("NSURLSession")}
	})
	return URLSessionClass
}

type _URLSessionClass struct {
	class objc.Class
}

// An interface definition for the [URLSession] class.
type IURLSession interface {
	objectivec.IObject
	// properties:
	Configuration() IURLSessionConfiguration
	SetConfiguration(value IURLSessionConfiguration)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DelegateQueue() IOperationQueue
	SetDelegateQueue(value IOperationQueue)
	SessionDescription() IString
	SetSessionDescription(value IString)
	// methods:
}

// An object that coordinates a group of related, network data transfer tasks.
//
// The class and related classes provide an API for downloading data from and uploading data to endpoints indicated by URLs. Your app can also use this API to perform background downloads when your app isn’t running or, in iOS, while your app is suspended. You can use the related and to support authentication and receive events like redirection and task completion. Your app creates one or more instances, each of which coordinates a group of related data-transfer tasks. For example, if you’re creating a web browser, your app might create one session per tab or window, or one session for interactive use and another for background downloads. Within each session, your app adds a series of tasks, each of which represents a request for a specific URL (following HTTP redirects, if necessary).


// An object that coordinates a group of related, network data transfer tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession
type URLSession struct {
	objectivec.Object
}

// URLSessionFrom constructs a [URLSession] from an unsafe.Pointer.
//
// An object that coordinates a group of related, network data transfer tasks.
func URLSessionFrom(ptr unsafe.Pointer) URLSession {
	return URLSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionClass) Alloc() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionClass) New() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSession) Init() URLSession {
	rv := objc.Send[URLSession](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSession) Autorelease() URLSession {
	rv := objc.Send[URLSession](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSession creates a new URLSession instance.
func NewURLSession() URLSession {
	return getURLSessionClass().New()
}



// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSession) Configuration() IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("configuration"))
	return rv
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSession) SetConfiguration(value IURLSessionConfiguration) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfiguration:"), value)
}


// The delegate assigned when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/delegate
func (u_ URLSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate assigned when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/delegate
func (u_ URLSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}


// The operation queue provided when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/delegatequeue
func (u_ URLSession) DelegateQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](u_.ID, objc.Sel("delegateQueue"))
	return rv
}


// The operation queue provided when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/delegatequeue
func (u_ URLSession) SetDelegateQueue(value IOperationQueue) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegateQueue:"), value)
}


// An app-defined descriptive label for the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/sessiondescription
func (u_ URLSession) SessionDescription() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("sessionDescription"))
	return rv
}


// An app-defined descriptive label for the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/sessiondescription
func (u_ URLSession) SetSessionDescription(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionDescription:"), value)
}



