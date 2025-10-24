// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class BAAppExtensionInfo */


/* debug [class_header]: Header for BAAppExtensionInfo */
// The class instance for the [BAAppExtensionInfo] class.
var (
	BAAppExtensionInfoClass     _BAAppExtensionInfoClass
	BAAppExtensionInfoClassOnce sync.Once
)

func getBAAppExtensionInfoClass() _BAAppExtensionInfoClass {
	BAAppExtensionInfoClassOnce.Do(func() {
		BAAppExtensionInfoClass = _BAAppExtensionInfoClass{objc.GetClass("BAAppExtensionInfo")}
	})
	return BAAppExtensionInfoClass
}

type _BAAppExtensionInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BAAppExtensionInfo */
// An interface definition for the [BAAppExtensionInfo] class.
type IBAAppExtensionInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BAAppExtensionInfo */
	// properties:
	RestrictedDownloadSizeRemaining() int
	SetRestrictedDownloadSizeRemaining(value int)
	RestrictedEssentialDownloadSizeRemaining() int
	SetRestrictedEssentialDownloadSizeRemaining(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BAAppExtensionInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BAAppExtensionInfo */
// Alloc allocates a new instance without initialization.
func (bc _BAAppExtensionInfoClass) Alloc() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BAAppExtensionInfoClass) New() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAppExtensionInfo) Init() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAppExtensionInfo) Autorelease() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAppExtensionInfo creates a new BAAppExtensionInfo instance.
func NewBAAppExtensionInfo() BAAppExtensionInfo {
	return getBAAppExtensionInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BAAppExtensionInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAppExtensionInfo
type BAAppExtensionInfo struct {
	objectivec.Object
}

// BAAppExtensionInfoFrom constructs a [BAAppExtensionInfo] from an unsafe.Pointer.
func BAAppExtensionInfoFrom(ptr unsafe.Pointer) BAAppExtensionInfo {
	return BAAppExtensionInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BAAppExtensionInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BAAppExtensionInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BAAppExtensionInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BAAppExtensionInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BAAppExtensionInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/backgroundassets/baappextensioninfo/restricteddownloadsizeremaining-4hea4
func (b_ BAAppExtensionInfo) RestrictedDownloadSizeRemaining() int {
	rv := objc.Send[int](b_.ID, objc.Sel("restrictedDownloadSizeRemaining"))
	return rv
}/* debug [instance_properties/getter]: restrictedDownloadSizeRemaining */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/backgroundassets/baappextensioninfo/restricteddownloadsizeremaining-4hea4
func (b_ BAAppExtensionInfo) SetRestrictedDownloadSizeRemaining(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setRestrictedDownloadSizeRemaining:"), value)
}/* debug [instance_properties/setter]: restrictedDownloadSizeRemaining */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/backgroundassets/baappextensioninfo/restrictedessentialdownloadsizeremaining-5r8v0
func (b_ BAAppExtensionInfo) RestrictedEssentialDownloadSizeRemaining() int {
	rv := objc.Send[int](b_.ID, objc.Sel("restrictedEssentialDownloadSizeRemaining"))
	return rv
}/* debug [instance_properties/getter]: restrictedEssentialDownloadSizeRemaining */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/backgroundassets/baappextensioninfo/restrictedessentialdownloadsizeremaining-5r8v0
func (b_ BAAppExtensionInfo) SetRestrictedEssentialDownloadSizeRemaining(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setRestrictedEssentialDownloadSizeRemaining:"), value)
}/* debug [instance_properties/setter]: restrictedEssentialDownloadSizeRemaining */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BAAppExtensionInfo */





