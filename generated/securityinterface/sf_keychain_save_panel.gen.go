// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFKeychainSavePanel] class.
var (
	SFKeychainSavePanelClass     _SFKeychainSavePanelClass
	SFKeychainSavePanelClassOnce sync.Once
)

func getSFKeychainSavePanelClass() _SFKeychainSavePanelClass {
	SFKeychainSavePanelClassOnce.Do(func() {
		SFKeychainSavePanelClass = _SFKeychainSavePanelClass{objc.GetClass("SFKeychainSavePanel")}
	})
	return SFKeychainSavePanelClass
}

type _SFKeychainSavePanelClass struct {
	class objc.Class
}

// An interface definition for the [SFKeychainSavePanel] class.
type ISFKeychainSavePanel interface {
	appkit.ISavePanel
	BeginSheetForDirectoryFileModalForWindowModalDelegateDidEndSelectorContextInfo(path appkit.string, name appkit.string, docWindow appkit.IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	Error() foundation.Error
	Keychain() unsafe.Pointer
	RunModalForDirectoryFile(path appkit.string, name appkit.string) int
	SetPassword(password appkit.string)
}

// A panel or sheet that allows the user to create a keychain.
//
// The following figure shows an example of a keychain save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel
type SFKeychainSavePanel struct {
	appkit.SavePanel
}

// SFKeychainSavePanelFrom constructs a [SFKeychainSavePanel] from an unsafe.Pointer.
//
// A panel or sheet that allows the user to create a keychain.
func SFKeychainSavePanelFrom(ptr unsafe.Pointer) SFKeychainSavePanel {
	return SFKeychainSavePanel{
		SavePanel: appkit.SavePanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFKeychainSavePanelClass) Alloc() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFKeychainSavePanelClass) New() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFKeychainSavePanel) Init() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFKeychainSavePanel) Autorelease() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFKeychainSavePanel creates a new SFKeychainSavePanel instance.
func NewSFKeychainSavePanel() SFKeychainSavePanel {
	return getSFKeychainSavePanelClass().New()
}


// Returns a shared keychain save panel object.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/shared()
func (sc _SFKeychainSavePanelClass) SharedKeychainSavePanel() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](objc.ID(sc.class), objc.Sel("sharedKeychainSavePanel"))
	return rv
}

// Displays a sheet that allows a user to create a new keychain.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/beginSheet(forDirectory:file:modalFor:modalDelegate:didEnd:contextInfo:)
func (s_ SFKeychainSavePanel) BeginSheetForDirectoryFileModalForWindowModalDelegateDidEndSelectorContextInfo(path appkit.string, name appkit.string, docWindow appkit.IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForDirectory:file:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), path, name, docWindow, delegate, didEndSelector, contextInfo)
}

// Returns the last error encountered by the keychain save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/error()
func (s_ SFKeychainSavePanel) Error() foundation.Error {
	rv := objc.Send[foundation.Error](s_.ID, objc.Sel("error"))
	return rv
}

// Returns the keychain created by the keychain save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/keychain()
func (s_ SFKeychainSavePanel) Keychain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("keychain"))
	return rv
}

// Displays a panel that allows a user to create a new keychain.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/runModal(forDirectory:file:)
func (s_ SFKeychainSavePanel) RunModalForDirectoryFile(path appkit.string, name appkit.string) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForDirectory:file:"), path, name)
	return rv
}

// Specifies the password for the keychain that will be created.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/setPassword(_:)
func (s_ SFKeychainSavePanel) SetPassword(password appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPassword:"), password)
}



