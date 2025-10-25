// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class cursor */


/* debug [class_header]: Header for cursor */
// The class instance for the [cursor] class.
var (
	CursorClass     _cursorClass
	CursorClassOnce sync.Once
)

func getcursorClass() _cursorClass {
	CursorClassOnce.Do(func() {
		CursorClass = _cursorClass{objc.GetClass("cursor")}
	})
	return CursorClass
}

type _cursorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for cursor */
// An interface definition for the [cursor] class.
type Icursor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for cursor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for cursor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for cursor */
// Alloc allocates a new instance without initialization.
func (cc _cursorClass) Alloc() cursor {
	rv := objc.Send[cursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _cursorClass) New() cursor {
	rv := objc.Send[cursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ cursor) Init() cursor {
	rv := objc.Send[cursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ cursor) Autorelease() cursor {
	rv := objc.Send[cursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcursor creates a new cursor instance.
func Newcursor() cursor {
	return getcursorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for cursor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/cursor
type cursor struct {
	objectivec.Object
}

// cursorFrom constructs a [cursor] from an unsafe.Pointer.
func cursorFrom(ptr unsafe.Pointer) cursor {
	return cursor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for cursor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for cursor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for cursor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for cursor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for cursor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class cursor */



