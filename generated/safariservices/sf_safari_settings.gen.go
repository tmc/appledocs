// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariSettings */


/* debug [class_header]: Header for SFSafariSettings */
// The class instance for the [SFSafariSettings] class.
var (
	SFSafariSettingsClass     _SFSafariSettingsClass
	SFSafariSettingsClassOnce sync.Once
)

func getSFSafariSettingsClass() _SFSafariSettingsClass {
	SFSafariSettingsClassOnce.Do(func() {
		SFSafariSettingsClass = _SFSafariSettingsClass{objc.GetClass("SFSafariSettings")}
	})
	return SFSafariSettingsClass
}

type _SFSafariSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariSettings */
// An interface definition for the [SFSafariSettings] class.
type ISFSafariSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariSettings */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariSettings */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariSettingsClass) Alloc() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariSettingsClass) New() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariSettings) Init() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariSettings) Autorelease() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariSettings creates a new SFSafariSettings instance.
func NewSFSafariSettings() SFSafariSettings {
	return getSFSafariSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariSettings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariSettings
type SFSafariSettings struct {
	objectivec.Object
}

// SFSafariSettingsFrom constructs a [SFSafariSettings] from an unsafe.Pointer.
func SFSafariSettingsFrom(ptr unsafe.Pointer) SFSafariSettings {
	return SFSafariSettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariSettings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariSettings/openExportBrowsingDataSettings(completionHandler:)
func (sc _SFSafariSettingsClass) OpenExportBrowsingDataSettingsWithCompletionHandler(completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openExportBrowsingDataSettingsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenExportBrowsingDataSettingsWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariSettings */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariSettings */


