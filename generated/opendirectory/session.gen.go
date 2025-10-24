// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class session */


/* debug [class_header]: Header for session */
// The class instance for the [session] class.
var (
	SessionClass     _sessionClass
	SessionClassOnce sync.Once
)

func getsessionClass() _sessionClass {
	SessionClassOnce.Do(func() {
		SessionClass = _sessionClass{objc.GetClass("session")}
	})
	return SessionClass
}

type _sessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for session */
// An interface definition for the [session] class.
type Isession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for session */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for session */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for session */
// Alloc allocates a new instance without initialization.
func (sc _sessionClass) Alloc() session {
	rv := objc.Send[session](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _sessionClass) New() session {
	rv := objc.Send[session](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ session) Init() session {
	rv := objc.Send[session](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ session) Autorelease() session {
	rv := objc.Send[session](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newsession creates a new session instance.
func Newsession() session {
	return getsessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for session */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/session
type session struct {
	objectivec.Object
}

// sessionFrom constructs a [session] from an unsafe.Pointer.
func sessionFrom(ptr unsafe.Pointer) session {
	return session{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for session *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for session */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for session */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for session */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for session */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class session */



