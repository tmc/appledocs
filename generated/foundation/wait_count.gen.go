// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class waitCount */


/* debug [class_header]: Header for waitCount */
// The class instance for the [waitCount] class.
var (
	WaitCountClass     _waitCountClass
	WaitCountClassOnce sync.Once
)

func getwaitCountClass() _waitCountClass {
	WaitCountClassOnce.Do(func() {
		WaitCountClass = _waitCountClass{objc.GetClass("waitCount")}
	})
	return WaitCountClass
}

type _waitCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for waitCount */
// An interface definition for the [waitCount] class.
type IwaitCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for waitCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for waitCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for waitCount */
// Alloc allocates a new instance without initialization.
func (wc _waitCountClass) Alloc() waitCount {
	rv := objc.Send[waitCount](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _waitCountClass) New() waitCount {
	rv := objc.Send[waitCount](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ waitCount) Init() waitCount {
	rv := objc.Send[waitCount](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ waitCount) Autorelease() waitCount {
	rv := objc.Send[waitCount](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwaitCount creates a new waitCount instance.
func NewwaitCount() waitCount {
	return getwaitCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for waitCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/waitCount
type waitCount struct {
	objectivec.Object
}

// waitCountFrom constructs a [waitCount] from an unsafe.Pointer.
func waitCountFrom(ptr unsafe.Pointer) waitCount {
	return waitCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for waitCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for waitCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for waitCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for waitCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for waitCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class waitCount */



