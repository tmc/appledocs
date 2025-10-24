// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFKeychainSavePanel */

/* debug [class_header]: Header for SFKeychainSavePanel */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFKeychainSavePanel */
// An interface definition for the [SFKeychainSavePanel] class.
type ISFKeychainSavePanel interface {
	appkit.ISavePanel

	/* debug [class_interface_properties]: Properties for SFKeychainSavePanel */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFKeychainSavePanel */
	// methods:
	BeginSheetForDirectoryFileModalForWindowModalDelegateDidEndSelectorContextInfo(path objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	Error() objc.IObject /* cross-framework: Error */
	Keychain() unsafe.Pointer
	RunModalForDirectoryFile(path objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */) int
	SetPassword(password objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFKeychainSavePanel */
// Alloc allocates a new instance without initialization.
func (sc _SFKeychainSavePanelClass) Alloc() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFKeychainSavePanel */
// A panel or sheet that allows the user to create a keychain.
//
// The following figure shows an example of a keychain save panel.

// A panel or sheet that allows the user to create a keychain.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFKeychainSavePanel */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFKeychainSavePanel */

// Returns a shared keychain save panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/shared()
func (sc _SFKeychainSavePanelClass) SharedKeychainSavePanel() SFKeychainSavePanel {
	rv := objc.Send[SFKeychainSavePanel](objc.ID(sc.class), objc.Sel("sharedKeychainSavePanel"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=SharedKeychainSavePanel) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFKeychainSavePanel */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFKeychainSavePanel */

// Displays a sheet that allows a user to create a new keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/beginSheet(forDirectory:file:modalFor:modalDelegate:didEnd:contextInfo:)
func (s_ SFKeychainSavePanel) BeginSheetForDirectoryFileModalForWindowModalDelegateDidEndSelectorContextInfo(path objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForDirectory:file:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), path, name, docWindow, delegate, didEndSelector, contextInfo)
} /* debug [instance_methods/method]: BeginSheetForDirectoryFileModalForWindowModalDelegateDidEndSelectorContextInfo */

// Returns the last error encountered by the keychain save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/error()
func (s_ SFKeychainSavePanel) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](s_.ID, objc.Sel("error"))
	return rv
} /* debug [instance_methods/method]: Error */

// Returns the keychain created by the keychain save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/keychain()
func (s_ SFKeychainSavePanel) Keychain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("keychain"))
	return rv
} /* debug [instance_methods/method]: Keychain */

// Displays a panel that allows a user to create a new keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/runModal(forDirectory:file:)
func (s_ SFKeychainSavePanel) RunModalForDirectoryFile(path objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForDirectory:file:"), path, name)
	return rv
} /* debug [instance_methods/method]: RunModalForDirectoryFile */

// Specifies the password for the keychain that will be created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFKeychainSavePanel/setPassword(_:)
func (s_ SFKeychainSavePanel) SetPassword(password objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPassword:"), password)
} /* debug [instance_methods/method]: SetPassword */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFKeychainSavePanel */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFKeychainSavePanel */
