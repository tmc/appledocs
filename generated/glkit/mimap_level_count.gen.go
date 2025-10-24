// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mimapLevelCount */


/* debug [class_header]: Header for mimapLevelCount */
// The class instance for the [mimapLevelCount] class.
var (
	MimapLevelCountClass     _mimapLevelCountClass
	MimapLevelCountClassOnce sync.Once
)

func getmimapLevelCountClass() _mimapLevelCountClass {
	MimapLevelCountClassOnce.Do(func() {
		MimapLevelCountClass = _mimapLevelCountClass{objc.GetClass("mimapLevelCount")}
	})
	return MimapLevelCountClass
}

type _mimapLevelCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mimapLevelCount */
// An interface definition for the [mimapLevelCount] class.
type ImimapLevelCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mimapLevelCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mimapLevelCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mimapLevelCount */
// Alloc allocates a new instance without initialization.
func (mc _mimapLevelCountClass) Alloc() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mimapLevelCountClass) New() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mimapLevelCount) Init() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mimapLevelCount) Autorelease() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmimapLevelCount creates a new mimapLevelCount instance.
func NewmimapLevelCount() mimapLevelCount {
	return getmimapLevelCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mimapLevelCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/mimapLevelCount-c.ivar
type mimapLevelCount struct {
	objectivec.Object
}

// mimapLevelCountFrom constructs a [mimapLevelCount] from an unsafe.Pointer.
func mimapLevelCountFrom(ptr unsafe.Pointer) mimapLevelCount {
	return mimapLevelCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mimapLevelCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mimapLevelCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mimapLevelCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mimapLevelCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mimapLevelCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mimapLevelCount */



