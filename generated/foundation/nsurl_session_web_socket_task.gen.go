// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [URLSessionWebSocketTask] class.
var (
	URLSessionWebSocketTaskClass     _URLSessionWebSocketTaskClass
	URLSessionWebSocketTaskClassOnce sync.Once
)

func getURLSessionWebSocketTaskClass() _URLSessionWebSocketTaskClass {
	URLSessionWebSocketTaskClassOnce.Do(func() {
		URLSessionWebSocketTaskClass = _URLSessionWebSocketTaskClass{objc.GetClass("NSURLSessionWebSocketTask")}
	})
	return URLSessionWebSocketTaskClass
}

type _URLSessionWebSocketTaskClass struct {
	class objc.Class
}





// An interface definition for the [URLSessionWebSocketTask] class.
type IURLSessionWebSocketTask interface {
	IURLSessionTask
	

	// properties:
	CloseCode() URLSessionWebSocketCloseCode
	CloseReason() IData
	MaximumMessageSize() int
	SetMaximumMessageSize(value int)
	HttpCookieStorage() IHTTPCookieStorage
	SetHttpCookieStorage(value IHTTPCookieStorage)


	

	// methods:
	ReceiveMessageWithCompletionHandler(completionHandler unsafe.Pointer)
	SendMessageCompletionHandler(message IURLSessionWebSocketMessage, completionHandler unsafe.Pointer)
	CancelWithCloseCodeReason(closeCode URLSessionWebSocketCloseCode, reason IData)
	SendPingWithPongReceiveHandler(pongReceiveHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (uc _URLSessionWebSocketTaskClass) Alloc() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionWebSocketTaskClass) New() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionWebSocketTask) Init() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionWebSocketTask) Autorelease() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionWebSocketTask creates a new URLSessionWebSocketTask instance.
func NewURLSessionWebSocketTask() URLSessionWebSocketTask {
	return getURLSessionWebSocketTaskClass().New()
}





// A URL session task that communicates over the WebSockets protocol standard.
//
// is a concrete subclass of that provides a message-oriented transport protocol over TCP and TLS in the form of WebSocket framing. It follows the WebSocket Protocol defined in . You create a with either a or URL. When creating the task, you can also provide a list of protocols to advertise during the handshake phase. Once the handshake completes, your app receives notifications through the session’s . You send data with and receive data with . The task performs reads and writes asynchronously, and allows you to send and receive messages that contain both binary frames and UTF-8 encoded text frames. The task enqueues any reads or writes you perform prior to the handshake’s completion, and executes them after the handshake completes. supports redirection and authentication like other types of tasks do, using the methods in . The WebSocket task calls the redirection and authentication delegate methods prior to completing the handshake. The WebSocket task also supports cookies, by storing cookies to the session configuration’s , and attaches cookies to outgoing HTTP handshake requests.


// A URL session task that communicates over the WebSockets protocol standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask
type URLSessionWebSocketTask struct {
	URLSessionTask
}

// URLSessionWebSocketTaskFrom constructs a [URLSessionWebSocketTask] from an unsafe.Pointer.
//
// A URL session task that communicates over the WebSockets protocol standard.
func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}




















// Reads a WebSocket message once all the frames of the message are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketTask/receiveMessageWithCompletionHandler:
func (u_ URLSessionWebSocketTask) ReceiveMessageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("receiveMessageWithCompletionHandler:"), completionHandler)
}


// Sends a WebSocket message, receiving the result in a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketTask/sendMessage:completionHandler:
func (u_ URLSessionWebSocketTask) SendMessageCompletionHandler(message IURLSessionWebSocketMessage, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("sendMessage:completionHandler:"), message, completionHandler)
}


// Sends a close frame with the given close code and optional close reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/cancel(with:reason:)
func (u_ URLSessionWebSocketTask) CancelWithCloseCodeReason(closeCode URLSessionWebSocketCloseCode, reason IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancelWithCloseCode:reason:"), closeCode, reason)
}


// Sends a ping frame from the client side, with a closure to receive the pong from the server endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/sendPing(pongReceiveHandler:)
func (u_ URLSessionWebSocketTask) SendPingWithPongReceiveHandler(pongReceiveHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("sendPingWithPongReceiveHandler:"), pongReceiveHandler)
}







// A code that indicates the reason a connection closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/closeCode-swift.property
func (u_ URLSessionWebSocketTask) CloseCode() URLSessionWebSocketCloseCode {
	rv := objc.Send[URLSessionWebSocketCloseCode](u_.ID, objc.Sel("closeCode"))
	return rv
}


// A block of data that provides further information about why a connection closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/closeReason
func (u_ URLSessionWebSocketTask) CloseReason() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("closeReason"))
	return rv
}


// The maximum number of bytes to buffer before the receive call fails with an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/maximumMessageSize
func (u_ URLSessionWebSocketTask) MaximumMessageSize() int {
	rv := objc.Send[int](u_.ID, objc.Sel("maximumMessageSize"))
	return rv
}


// The maximum number of bytes to buffer before the receive call fails with an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/maximumMessageSize
func (u_ URLSessionWebSocketTask) SetMaximumMessageSize(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMaximumMessageSize:"), value)
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionWebSocketTask) HttpCookieStorage() IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](u_.ID, objc.Sel("httpCookieStorage"))
	return rv
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionWebSocketTask) SetHttpCookieStorage(value IHTTPCookieStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpCookieStorage:"), value)
}








