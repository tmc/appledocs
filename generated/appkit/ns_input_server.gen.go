// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInputServer */


/* debug [class_header]: Header for NSInputServer */
// The class instance for the [InputServer] class.
var (
	InputServerClass     _InputServerClass
	InputServerClassOnce sync.Once
)

func getInputServerClass() _InputServerClass {
	InputServerClassOnce.Do(func() {
		InputServerClass = _InputServerClass{objc.GetClass("NSInputServer")}
	})
	return InputServerClass
}

type _InputServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InputServer */
// An interface definition for the [InputServer] class.
type IInputServer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for InputServer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InputServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InputServer */
// Alloc allocates a new instance without initialization.
func (ic _InputServerClass) Alloc() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InputServerClass) New() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputServer) Init() InputServer {
	rv := objc.Send[InputServer](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputServer) Autorelease() InputServer {
	rv := objc.Send[InputServer](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputServer creates a new InputServer instance.
func NewInputServer() InputServer {
	return getInputServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InputServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputServer
type InputServer struct {
	objectivec.Object
}

// InputServerFrom constructs a [InputServer] from an unsafe.Pointer.
func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InputServer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputServer/initWithDelegate:name:
func NewInputServerWithDelegateName(delegate objc.IObject, name objc.IObject /* cross-framework: NSString */) InputServer {
	instance := getInputServerClass().Alloc()
	rv := objc.Send[InputServer](instance.ID, objc.Sel("initWithDelegate:name:"), delegate, name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInputServerWithDelegateName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InputServer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InputServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InputServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InputServer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInputServer */


