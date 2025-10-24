// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterAdditionalInfo */


/* debug [class_header]: Header for MTRContentLauncherClusterAdditionalInfo */
// The class instance for the [MTRContentLauncherClusterAdditionalInfo] class.
var (
	MTRContentLauncherClusterAdditionalInfoClass     _MTRContentLauncherClusterAdditionalInfoClass
	MTRContentLauncherClusterAdditionalInfoClassOnce sync.Once
)

func getMTRContentLauncherClusterAdditionalInfoClass() _MTRContentLauncherClusterAdditionalInfoClass {
	MTRContentLauncherClusterAdditionalInfoClassOnce.Do(func() {
		MTRContentLauncherClusterAdditionalInfoClass = _MTRContentLauncherClusterAdditionalInfoClass{objc.GetClass("MTRContentLauncherClusterAdditionalInfo")}
	})
	return MTRContentLauncherClusterAdditionalInfoClass
}

type _MTRContentLauncherClusterAdditionalInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterAdditionalInfo */
// An interface definition for the [MTRContentLauncherClusterAdditionalInfo] class.
type IMTRContentLauncherClusterAdditionalInfo interface {
	IMTRContentLauncherClusterAdditionalInfoStruct
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterAdditionalInfo */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterAdditionalInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterAdditionalInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterAdditionalInfoClass) Alloc() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterAdditionalInfoClass) New() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterAdditionalInfo) Init() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterAdditionalInfo) Autorelease() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterAdditionalInfo creates a new MTRContentLauncherClusterAdditionalInfo instance.
func NewMTRContentLauncherClusterAdditionalInfo() MTRContentLauncherClusterAdditionalInfo {
	return getMTRContentLauncherClusterAdditionalInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterAdditionalInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo
type MTRContentLauncherClusterAdditionalInfo struct {
	MTRContentLauncherClusterAdditionalInfoStruct
}

// MTRContentLauncherClusterAdditionalInfoFrom constructs a [MTRContentLauncherClusterAdditionalInfo] from an unsafe.Pointer.
func MTRContentLauncherClusterAdditionalInfoFrom(ptr unsafe.Pointer) MTRContentLauncherClusterAdditionalInfo {
	return MTRContentLauncherClusterAdditionalInfo{
		MTRContentLauncherClusterAdditionalInfoStruct: MTRContentLauncherClusterAdditionalInfoStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterAdditionalInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterAdditionalInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterAdditionalInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterAdditionalInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterAdditionalInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterAdditionalInfo */



