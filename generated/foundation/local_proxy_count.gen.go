// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class localProxyCount */


/* debug [class_header]: Header for localProxyCount */
// The class instance for the [localProxyCount] class.
var (
	LocalProxyCountClass     _localProxyCountClass
	LocalProxyCountClassOnce sync.Once
)

func getlocalProxyCountClass() _localProxyCountClass {
	LocalProxyCountClassOnce.Do(func() {
		LocalProxyCountClass = _localProxyCountClass{objc.GetClass("localProxyCount")}
	})
	return LocalProxyCountClass
}

type _localProxyCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for localProxyCount */
// An interface definition for the [localProxyCount] class.
type IlocalProxyCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for localProxyCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for localProxyCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for localProxyCount */
// Alloc allocates a new instance without initialization.
func (lc _localProxyCountClass) Alloc() localProxyCount {
	rv := objc.Send[localProxyCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _localProxyCountClass) New() localProxyCount {
	rv := objc.Send[localProxyCount](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ localProxyCount) Init() localProxyCount {
	rv := objc.Send[localProxyCount](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ localProxyCount) Autorelease() localProxyCount {
	rv := objc.Send[localProxyCount](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlocalProxyCount creates a new localProxyCount instance.
func NewlocalProxyCount() localProxyCount {
	return getlocalProxyCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for localProxyCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/localProxyCount
type localProxyCount struct {
	objectivec.Object
}

// localProxyCountFrom constructs a [localProxyCount] from an unsafe.Pointer.
func localProxyCountFrom(ptr unsafe.Pointer) localProxyCount {
	return localProxyCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for localProxyCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for localProxyCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for localProxyCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for localProxyCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for localProxyCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class localProxyCount */



