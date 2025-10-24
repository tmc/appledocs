// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterAdditionalInfoStruct */


/* debug [class_header]: Header for MTRContentLauncherClusterAdditionalInfoStruct */
// The class instance for the [MTRContentLauncherClusterAdditionalInfoStruct] class.
var (
	MTRContentLauncherClusterAdditionalInfoStructClass     _MTRContentLauncherClusterAdditionalInfoStructClass
	MTRContentLauncherClusterAdditionalInfoStructClassOnce sync.Once
)

func getMTRContentLauncherClusterAdditionalInfoStructClass() _MTRContentLauncherClusterAdditionalInfoStructClass {
	MTRContentLauncherClusterAdditionalInfoStructClassOnce.Do(func() {
		MTRContentLauncherClusterAdditionalInfoStructClass = _MTRContentLauncherClusterAdditionalInfoStructClass{objc.GetClass("MTRContentLauncherClusterAdditionalInfoStruct")}
	})
	return MTRContentLauncherClusterAdditionalInfoStructClass
}

type _MTRContentLauncherClusterAdditionalInfoStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterAdditionalInfoStruct */
// An interface definition for the [MTRContentLauncherClusterAdditionalInfoStruct] class.
type IMTRContentLauncherClusterAdditionalInfoStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterAdditionalInfoStruct */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterAdditionalInfoStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterAdditionalInfoStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterAdditionalInfoStructClass) Alloc() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterAdditionalInfoStructClass) New() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Init() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Autorelease() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterAdditionalInfoStruct creates a new MTRContentLauncherClusterAdditionalInfoStruct instance.
func NewMTRContentLauncherClusterAdditionalInfoStruct() MTRContentLauncherClusterAdditionalInfoStruct {
	return getMTRContentLauncherClusterAdditionalInfoStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterAdditionalInfoStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct
type MTRContentLauncherClusterAdditionalInfoStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterAdditionalInfoStructFrom constructs a [MTRContentLauncherClusterAdditionalInfoStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterAdditionalInfoStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterAdditionalInfoStruct {
	return MTRContentLauncherClusterAdditionalInfoStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterAdditionalInfoStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterAdditionalInfoStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterAdditionalInfoStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterAdditionalInfoStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterAdditionalInfoStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct/name
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct/name
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct/value
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct/value
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterAdditionalInfoStruct */



