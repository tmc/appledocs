// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenSession */


/* debug [class_header]: Header for TKTokenSession */
// The class instance for the [TKTokenSession] class.
var (
	TKTokenSessionClass     _TKTokenSessionClass
	TKTokenSessionClassOnce sync.Once
)

func getTKTokenSessionClass() _TKTokenSessionClass {
	TKTokenSessionClassOnce.Do(func() {
		TKTokenSessionClass = _TKTokenSessionClass{objc.GetClass("TKTokenSession")}
	})
	return TKTokenSessionClass
}

type _TKTokenSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenSession */
// An interface definition for the [TKTokenSession] class.
type ITKTokenSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenSession */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Token() ITKToken
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenSession */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenSessionClass) Alloc() TKTokenSession {
	rv := objc.Send[TKTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenSessionClass) New() TKTokenSession {
	rv := objc.Send[TKTokenSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenSession) Init() TKTokenSession {
	rv := objc.Send[TKTokenSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenSession) Autorelease() TKTokenSession {
	rv := objc.Send[TKTokenSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenSession creates a new TKTokenSession instance.
func NewTKTokenSession() TKTokenSession {
	return getTKTokenSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenSession */
// A token session that manages the authentication state of a token.
//
// A token session communicates with its delegate to perform operations with its token that are bound to the authentication state. A session is always instantiated by a instance through the token’s delegate when the framework detects access to the token from a new authentication session.


// A token session that manages the authentication state of a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession
type TKTokenSession struct {
	objectivec.Object
}

// TKTokenSessionFrom constructs a [TKTokenSession] from an unsafe.Pointer.
//
// A token session that manages the authentication state of a token.
func TKTokenSessionFrom(ptr unsafe.Pointer) TKTokenSession {
	return TKTokenSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenSession */

// Initializes a token session with the specified token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/init(token:)
func NewTKTokenSessionWithToken(token ITKToken) TKTokenSession {
	instance := getTKTokenSessionClass().Alloc()
	rv := objc.Send[TKTokenSession](instance.ID, objc.Sel("initWithToken:"), token)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenSessionWithToken */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenSession */

// The token session delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/delegate
func (t_ TKTokenSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The token session delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/delegate
func (t_ TKTokenSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The token to which the session is bound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/token
func (t_ TKTokenSession) Token() ITKToken {
	rv := objc.Send[TKToken](t_.ID, objc.Sel("token"))
	return rv
}/* debug [instance_properties/getter]: token */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenSession */


