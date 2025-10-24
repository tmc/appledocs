// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/security"
)

/* debug [class.gen.go]: Generating class SFKeychainSettingsPanel */


/* debug [class_header]: Header for SFKeychainSettingsPanel */
// The class instance for the [SFKeychainSettingsPanel] class.
var (
	SFKeychainSettingsPanelClass     _SFKeychainSettingsPanelClass
	SFKeychainSettingsPanelClassOnce sync.Once
)

func getSFKeychainSettingsPanelClass() _SFKeychainSettingsPanelClass {
	SFKeychainSettingsPanelClassOnce.Do(func() {
		SFKeychainSettingsPanelClass = _SFKeychainSettingsPanelClass{objc.GetClass("SFKeychainSettingsPanel")}
	})
	return SFKeychainSettingsPanelClass
}

type _SFKeychainSettingsPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFKeychainSettingsPanel */
// An interface definition for the [SFKeychainSettingsPanel] class.
type ISFKeychainSettingsPanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for SFKeychainSettingsPanel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFKeychainSettingsPanel */
	// methods:
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoSettingsKeychain(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, settings security.SecKeychainSettings, keychain unsafe.Pointer)
	RunModalForSettingsKeychain(settings security.SecKeychainSettings, keychain unsafe.Pointer) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFKeychainSettingsPanel */
// Alloc allocates a new instance without initialization.
func (sc _SFKeychainSettingsPanelClass) Alloc() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFKeychainSettingsPanelClass) New() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFKeychainSettingsPanel) Init() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFKeychainSettingsPanel) Autorelease() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFKeychainSettingsPanel creates a new SFKeychainSettingsPanel instance.
func NewSFKeychainSettingsPanel() SFKeychainSettingsPanel {
	return getSFKeychainSettingsPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFKeychainSettingsPanel */
// A panel or sheet that allows users to change their keychain settings.
//
// Keychain settings include: Lock after a set period of inactivity Lock on sleep Synchronize using .Mac The following figure shows an example of a keychain settings panel. For more information, see .


// A panel or sheet that allows users to change their keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel
type SFKeychainSettingsPanel struct {
	appkit.Panel
}

// SFKeychainSettingsPanelFrom constructs a [SFKeychainSettingsPanel] from an unsafe.Pointer.
//
// A panel or sheet that allows users to change their keychain settings.
func SFKeychainSettingsPanelFrom(ptr unsafe.Pointer) SFKeychainSettingsPanel {
	return SFKeychainSettingsPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFKeychainSettingsPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFKeychainSettingsPanel */

// Returns a shared keychain settings panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/shared()
func (sc _SFKeychainSettingsPanelClass) SharedKeychainSettingsPanel() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](objc.ID(sc.class), objc.Sel("sharedKeychainSettingsPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedKeychainSettingsPanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFKeychainSettingsPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFKeychainSettingsPanel */

// Displays a sheet that allows users to change keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:settings:keychain:)
func (s_ SFKeychainSettingsPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoSettingsKeychain(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, settings security.SecKeychainSettings, keychain unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:settings:keychain:"), docWindow, delegate, didEndSelector, contextInfo, settings, keychain)
}/* debug [instance_methods/method]: BeginSheetForWindowModalDelegateDidEndSelectorContextInfoSettingsKeychain */


// Displays a panel that allows users to change keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/runModal(for:keychain:)
func (s_ SFKeychainSettingsPanel) RunModalForSettingsKeychain(settings security.SecKeychainSettings, keychain unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForSettings:keychain:"), settings, keychain)
	return rv
}/* debug [instance_methods/method]: RunModalForSettingsKeychain */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFKeychainSettingsPanel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFKeychainSettingsPanel */






