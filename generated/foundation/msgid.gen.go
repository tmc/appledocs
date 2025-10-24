// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class msgid */


/* debug [class_header]: Header for msgid */
// The class instance for the [msgid] class.
var (
	MsgidClass     _msgidClass
	MsgidClassOnce sync.Once
)

func getmsgidClass() _msgidClass {
	MsgidClassOnce.Do(func() {
		MsgidClass = _msgidClass{objc.GetClass("msgid")}
	})
	return MsgidClass
}

type _msgidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for msgid */
// An interface definition for the [msgid] class.
type Imsgid interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for msgid */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for msgid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for msgid */
// Alloc allocates a new instance without initialization.
func (mc _msgidClass) Alloc() msgid {
	rv := objc.Send[msgid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _msgidClass) New() msgid {
	rv := objc.Send[msgid](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ msgid) Init() msgid {
	rv := objc.Send[msgid](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ msgid) Autorelease() msgid {
	rv := objc.Send[msgid](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmsgid creates a new msgid instance.
func Newmsgid() msgid {
	return getmsgidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for msgid */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/msgid-c.ivar
type msgid struct {
	objectivec.Object
}

// msgidFrom constructs a [msgid] from an unsafe.Pointer.
func msgidFrom(ptr unsafe.Pointer) msgid {
	return msgid{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for msgid *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for msgid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for msgid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for msgid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for msgid */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class msgid */



