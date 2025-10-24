// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterBrandingInformation */


/* debug [class_header]: Header for MTRContentLauncherClusterBrandingInformation */
// The class instance for the [MTRContentLauncherClusterBrandingInformation] class.
var (
	MTRContentLauncherClusterBrandingInformationClass     _MTRContentLauncherClusterBrandingInformationClass
	MTRContentLauncherClusterBrandingInformationClassOnce sync.Once
)

func getMTRContentLauncherClusterBrandingInformationClass() _MTRContentLauncherClusterBrandingInformationClass {
	MTRContentLauncherClusterBrandingInformationClassOnce.Do(func() {
		MTRContentLauncherClusterBrandingInformationClass = _MTRContentLauncherClusterBrandingInformationClass{objc.GetClass("MTRContentLauncherClusterBrandingInformation")}
	})
	return MTRContentLauncherClusterBrandingInformationClass
}

type _MTRContentLauncherClusterBrandingInformationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterBrandingInformation */
// An interface definition for the [MTRContentLauncherClusterBrandingInformation] class.
type IMTRContentLauncherClusterBrandingInformation interface {
	IMTRContentLauncherClusterBrandingInformationStruct
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterBrandingInformation */
	// properties:
	Background() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */
	SetBackground(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */)
	Logo() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */
	SetLogo(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */)
	ProgressBar() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */
	SetProgressBar(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */)
	ProviderName() objc.IObject /* cross-framework: NSString */
	SetProviderName(value objc.IObject /* cross-framework: NSString */)
	Splash() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */
	SetSplash(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */)
	WaterMark() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */
	SetWaterMark(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterBrandingInformation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterBrandingInformation */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterBrandingInformationClass) Alloc() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterBrandingInformationClass) New() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterBrandingInformation) Init() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterBrandingInformation) Autorelease() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterBrandingInformation creates a new MTRContentLauncherClusterBrandingInformation instance.
func NewMTRContentLauncherClusterBrandingInformation() MTRContentLauncherClusterBrandingInformation {
	return getMTRContentLauncherClusterBrandingInformationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterBrandingInformation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation
type MTRContentLauncherClusterBrandingInformation struct {
	MTRContentLauncherClusterBrandingInformationStruct
}

// MTRContentLauncherClusterBrandingInformationFrom constructs a [MTRContentLauncherClusterBrandingInformation] from an unsafe.Pointer.
func MTRContentLauncherClusterBrandingInformationFrom(ptr unsafe.Pointer) MTRContentLauncherClusterBrandingInformation {
	return MTRContentLauncherClusterBrandingInformation{
		MTRContentLauncherClusterBrandingInformationStruct: MTRContentLauncherClusterBrandingInformationStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterBrandingInformation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterBrandingInformation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterBrandingInformation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterBrandingInformation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterBrandingInformation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/background
func (m_ MTRContentLauncherClusterBrandingInformation) Background() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("background"))
	return rv
}/* debug [instance_properties/getter]: background */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/background
func (m_ MTRContentLauncherClusterBrandingInformation) SetBackground(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackground:"), value)
}/* debug [instance_properties/setter]: background */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/logo
func (m_ MTRContentLauncherClusterBrandingInformation) Logo() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("logo"))
	return rv
}/* debug [instance_properties/getter]: logo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/logo
func (m_ MTRContentLauncherClusterBrandingInformation) SetLogo(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogo:"), value)
}/* debug [instance_properties/setter]: logo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/progressBar
func (m_ MTRContentLauncherClusterBrandingInformation) ProgressBar() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("progressBar"))
	return rv
}/* debug [instance_properties/getter]: progressBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/progressBar
func (m_ MTRContentLauncherClusterBrandingInformation) SetProgressBar(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressBar:"), value)
}/* debug [instance_properties/setter]: progressBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/providerName
func (m_ MTRContentLauncherClusterBrandingInformation) ProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("providerName"))
	return rv
}/* debug [instance_properties/getter]: providerName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/providerName
func (m_ MTRContentLauncherClusterBrandingInformation) SetProviderName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderName:"), value)
}/* debug [instance_properties/setter]: providerName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/splash
func (m_ MTRContentLauncherClusterBrandingInformation) Splash() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("splash"))
	return rv
}/* debug [instance_properties/getter]: splash */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/splash
func (m_ MTRContentLauncherClusterBrandingInformation) SetSplash(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSplash:"), value)
}/* debug [instance_properties/setter]: splash */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/waterMark
func (m_ MTRContentLauncherClusterBrandingInformation) WaterMark() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("waterMark"))
	return rv
}/* debug [instance_properties/getter]: waterMark */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation/waterMark
func (m_ MTRContentLauncherClusterBrandingInformation) SetWaterMark(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWaterMark:"), value)
}/* debug [instance_properties/setter]: waterMark */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterBrandingInformation */



