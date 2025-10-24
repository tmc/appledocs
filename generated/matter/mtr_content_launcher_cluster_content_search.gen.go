// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterContentSearch */


/* debug [class_header]: Header for MTRContentLauncherClusterContentSearch */
// The class instance for the [MTRContentLauncherClusterContentSearch] class.
var (
	MTRContentLauncherClusterContentSearchClass     _MTRContentLauncherClusterContentSearchClass
	MTRContentLauncherClusterContentSearchClassOnce sync.Once
)

func getMTRContentLauncherClusterContentSearchClass() _MTRContentLauncherClusterContentSearchClass {
	MTRContentLauncherClusterContentSearchClassOnce.Do(func() {
		MTRContentLauncherClusterContentSearchClass = _MTRContentLauncherClusterContentSearchClass{objc.GetClass("MTRContentLauncherClusterContentSearch")}
	})
	return MTRContentLauncherClusterContentSearchClass
}

type _MTRContentLauncherClusterContentSearchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterContentSearch */
// An interface definition for the [MTRContentLauncherClusterContentSearch] class.
type IMTRContentLauncherClusterContentSearch interface {
	IMTRContentLauncherClusterContentSearchStruct
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterContentSearch */
	// properties:
	ParameterList() objc.IObject /* cross-framework: NSArray */
	SetParameterList(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterContentSearch */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterContentSearch */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterContentSearchClass) Alloc() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterContentSearchClass) New() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterContentSearch) Init() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterContentSearch) Autorelease() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterContentSearch creates a new MTRContentLauncherClusterContentSearch instance.
func NewMTRContentLauncherClusterContentSearch() MTRContentLauncherClusterContentSearch {
	return getMTRContentLauncherClusterContentSearchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterContentSearch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearch
type MTRContentLauncherClusterContentSearch struct {
	MTRContentLauncherClusterContentSearchStruct
}

// MTRContentLauncherClusterContentSearchFrom constructs a [MTRContentLauncherClusterContentSearch] from an unsafe.Pointer.
func MTRContentLauncherClusterContentSearchFrom(ptr unsafe.Pointer) MTRContentLauncherClusterContentSearch {
	return MTRContentLauncherClusterContentSearch{
		MTRContentLauncherClusterContentSearchStruct: MTRContentLauncherClusterContentSearchStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterContentSearch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterContentSearch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterContentSearch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterContentSearch */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterContentSearch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearch/parameterList
func (m_ MTRContentLauncherClusterContentSearch) ParameterList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("parameterList"))
	return rv
}/* debug [instance_properties/getter]: parameterList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearch/parameterList
func (m_ MTRContentLauncherClusterContentSearch) SetParameterList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameterList:"), value)
}/* debug [instance_properties/setter]: parameterList */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterContentSearch */



