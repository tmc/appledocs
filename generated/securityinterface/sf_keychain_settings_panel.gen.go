// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [SFKeychainSettingsPanel] class.
type ISFKeychainSettingsPanel interface {
	appkit.IPanel
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoSettingsKeychain(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, settings unsafe.Pointer, keychain unsafe.Pointer)
	RunModalForSettingsKeychain(settings unsafe.Pointer, keychain unsafe.Pointer) int
}

// A panel or sheet that allows users to change their keychain settings.
//
// Keychain settings include: Lock after a set period of inactivity Lock on sleep Synchronize using .Mac The following figure shows an example of a keychain settings panel. For more information, see .
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFKeychainSettingsPanelClass) Alloc() SFKeychainSettingsPanel {
	rv := objc.Send[SFKeychainSettingsPanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a shared keychain settings panel object.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/shared()
func (sc _SFKeychainSettingsPanelClass) SharedKeychainSettingsPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharedKeychainSettingsPanel"))
	return rv
}

// Displays a sheet that allows users to change keychain settings.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:settings:keychain:)
func (s_ SFKeychainSettingsPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoSettingsKeychain(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, settings unsafe.Pointer, keychain unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:settings:keychain:"), docWindow, delegate, didEndSelector, contextInfo, settings, keychain)
}

// Displays a panel that allows users to change keychain settings.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSettingsPanel/runModal(for:keychain:)
func (s_ SFKeychainSettingsPanel) RunModalForSettingsKeychain(settings unsafe.Pointer, keychain unsafe.Pointer) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForSettings:keychain:"), settings, keychain)
	return rv
}




