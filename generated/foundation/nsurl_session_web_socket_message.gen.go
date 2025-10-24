// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionWebSocketMessage */


/* debug [class_header]: Header for NSURLSessionWebSocketMessage */
// The class instance for the [URLSessionWebSocketMessage] class.
var (
	URLSessionWebSocketMessageClass     _URLSessionWebSocketMessageClass
	URLSessionWebSocketMessageClassOnce sync.Once
)

func getURLSessionWebSocketMessageClass() _URLSessionWebSocketMessageClass {
	URLSessionWebSocketMessageClassOnce.Do(func() {
		URLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}
	})
	return URLSessionWebSocketMessageClass
}

type _URLSessionWebSocketMessageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionWebSocketMessage */
// An interface definition for the [URLSessionWebSocketMessage] class.
type IURLSessionWebSocketMessage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionWebSocketMessage */
	// properties:
	Data() IData
	String() IString
	Type() URLSessionWebSocketMessageType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionWebSocketMessage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionWebSocketMessage */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionWebSocketMessageClass) Alloc() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionWebSocketMessageClass) New() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionWebSocketMessage) Init() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionWebSocketMessage) Autorelease() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionWebSocketMessage creates a new URLSessionWebSocketMessage instance.
func NewURLSessionWebSocketMessage() URLSessionWebSocketMessage {
	return getURLSessionWebSocketMessageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionWebSocketMessage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage
type URLSessionWebSocketMessage struct {
	objectivec.Object
}

// URLSessionWebSocketMessageFrom constructs a [URLSessionWebSocketMessage] from an unsafe.Pointer.
func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionWebSocketMessage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithData:
func NewURLSessionWebSocketMessageWithData(data IData) URLSessionWebSocketMessage {
	instance := getURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[URLSessionWebSocketMessage](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLSessionWebSocketMessageWithData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithString:
func NewURLSessionWebSocketMessageWithString(string_ IString) URLSessionWebSocketMessage {
	instance := getURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[URLSessionWebSocketMessage](instance.ID, objc.Sel("initWithString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLSessionWebSocketMessageWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionWebSocketMessage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionWebSocketMessage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionWebSocketMessage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionWebSocketMessage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/data
func (u_ URLSessionWebSocketMessage) Data() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/string
func (u_ URLSessionWebSocketMessage) String() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/type
func (u_ URLSessionWebSocketMessage) Type() URLSessionWebSocketMessageType {
	rv := objc.Send[URLSessionWebSocketMessageType](u_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionWebSocketMessage */


