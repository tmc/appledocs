// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterBrandingInformationStruct */


/* debug [class_header]: Header for MTRContentLauncherClusterBrandingInformationStruct */
// The class instance for the [MTRContentLauncherClusterBrandingInformationStruct] class.
var (
	MTRContentLauncherClusterBrandingInformationStructClass     _MTRContentLauncherClusterBrandingInformationStructClass
	MTRContentLauncherClusterBrandingInformationStructClassOnce sync.Once
)

func getMTRContentLauncherClusterBrandingInformationStructClass() _MTRContentLauncherClusterBrandingInformationStructClass {
	MTRContentLauncherClusterBrandingInformationStructClassOnce.Do(func() {
		MTRContentLauncherClusterBrandingInformationStructClass = _MTRContentLauncherClusterBrandingInformationStructClass{objc.GetClass("MTRContentLauncherClusterBrandingInformationStruct")}
	})
	return MTRContentLauncherClusterBrandingInformationStructClass
}

type _MTRContentLauncherClusterBrandingInformationStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterBrandingInformationStruct */
// An interface definition for the [MTRContentLauncherClusterBrandingInformationStruct] class.
type IMTRContentLauncherClusterBrandingInformationStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterBrandingInformationStruct */
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

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterBrandingInformationStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterBrandingInformationStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterBrandingInformationStructClass) Alloc() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterBrandingInformationStructClass) New() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Init() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Autorelease() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterBrandingInformationStruct creates a new MTRContentLauncherClusterBrandingInformationStruct instance.
func NewMTRContentLauncherClusterBrandingInformationStruct() MTRContentLauncherClusterBrandingInformationStruct {
	return getMTRContentLauncherClusterBrandingInformationStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterBrandingInformationStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct
type MTRContentLauncherClusterBrandingInformationStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterBrandingInformationStructFrom constructs a [MTRContentLauncherClusterBrandingInformationStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterBrandingInformationStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterBrandingInformationStruct {
	return MTRContentLauncherClusterBrandingInformationStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterBrandingInformationStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterBrandingInformationStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterBrandingInformationStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterBrandingInformationStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterBrandingInformationStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Background() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("background"))
	return rv
}/* debug [instance_properties/getter]: background */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetBackground(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackground:"), value)
}/* debug [instance_properties/setter]: background */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Logo() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("logo"))
	return rv
}/* debug [instance_properties/getter]: logo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetLogo(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogo:"), value)
}/* debug [instance_properties/setter]: logo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/progressBar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProgressBar() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("progressBar"))
	return rv
}/* debug [instance_properties/getter]: progressBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/progressBar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProgressBar(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressBar:"), value)
}/* debug [instance_properties/setter]: progressBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/providerName
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("providerName"))
	return rv
}/* debug [instance_properties/getter]: providerName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/providerName
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProviderName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderName:"), value)
}/* debug [instance_properties/setter]: providerName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Splash() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("splash"))
	return rv
}/* debug [instance_properties/getter]: splash */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetSplash(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSplash:"), value)
}/* debug [instance_properties/setter]: splash */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/waterMark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) WaterMark() objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("waterMark"))
	return rv
}/* debug [instance_properties/getter]: waterMark */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct/waterMark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetWaterMark(value objc.IObject /* cross-framework: MTRContentLauncherClusterStyleInformationStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWaterMark:"), value)
}/* debug [instance_properties/setter]: waterMark */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterBrandingInformationStruct */



