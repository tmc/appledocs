// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterContentSearchStruct */


/* debug [class_header]: Header for MTRContentLauncherClusterContentSearchStruct */
// The class instance for the [MTRContentLauncherClusterContentSearchStruct] class.
var (
	MTRContentLauncherClusterContentSearchStructClass     _MTRContentLauncherClusterContentSearchStructClass
	MTRContentLauncherClusterContentSearchStructClassOnce sync.Once
)

func getMTRContentLauncherClusterContentSearchStructClass() _MTRContentLauncherClusterContentSearchStructClass {
	MTRContentLauncherClusterContentSearchStructClassOnce.Do(func() {
		MTRContentLauncherClusterContentSearchStructClass = _MTRContentLauncherClusterContentSearchStructClass{objc.GetClass("MTRContentLauncherClusterContentSearchStruct")}
	})
	return MTRContentLauncherClusterContentSearchStructClass
}

type _MTRContentLauncherClusterContentSearchStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterContentSearchStruct */
// An interface definition for the [MTRContentLauncherClusterContentSearchStruct] class.
type IMTRContentLauncherClusterContentSearchStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterContentSearchStruct */
	// properties:
	ParameterList() objc.IObject /* cross-framework: NSArray */
	SetParameterList(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterContentSearchStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterContentSearchStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterContentSearchStructClass) Alloc() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterContentSearchStructClass) New() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterContentSearchStruct) Init() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterContentSearchStruct) Autorelease() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterContentSearchStruct creates a new MTRContentLauncherClusterContentSearchStruct instance.
func NewMTRContentLauncherClusterContentSearchStruct() MTRContentLauncherClusterContentSearchStruct {
	return getMTRContentLauncherClusterContentSearchStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterContentSearchStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearchStruct
type MTRContentLauncherClusterContentSearchStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterContentSearchStructFrom constructs a [MTRContentLauncherClusterContentSearchStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterContentSearchStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterContentSearchStruct {
	return MTRContentLauncherClusterContentSearchStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterContentSearchStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterContentSearchStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterContentSearchStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterContentSearchStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterContentSearchStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearchStruct/parameterList
func (m_ MTRContentLauncherClusterContentSearchStruct) ParameterList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("parameterList"))
	return rv
}/* debug [instance_properties/getter]: parameterList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearchStruct/parameterList
func (m_ MTRContentLauncherClusterContentSearchStruct) SetParameterList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameterList:"), value)
}/* debug [instance_properties/setter]: parameterList */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterContentSearchStruct */



