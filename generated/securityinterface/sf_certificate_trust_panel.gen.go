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

/* debug [class.gen.go]: Generating class SFCertificateTrustPanel */


/* debug [class_header]: Header for SFCertificateTrustPanel */
// The class instance for the [SFCertificateTrustPanel] class.
var (
	SFCertificateTrustPanelClass     _SFCertificateTrustPanelClass
	SFCertificateTrustPanelClassOnce sync.Once
)

func getSFCertificateTrustPanelClass() _SFCertificateTrustPanelClass {
	SFCertificateTrustPanelClassOnce.Do(func() {
		SFCertificateTrustPanelClass = _SFCertificateTrustPanelClass{objc.GetClass("SFCertificateTrustPanel")}
	})
	return SFCertificateTrustPanelClass
}

type _SFCertificateTrustPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFCertificateTrustPanel */
// An interface definition for the [SFCertificateTrustPanel] class.
type ISFCertificateTrustPanel interface {
	ISFCertificatePanel
	
/* debug [class_interface_properties]: Properties for SFCertificateTrustPanel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFCertificateTrustPanel */
	// methods:
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustMessage(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, message objc.IObject /* cross-framework: NSString */)
	InformativeText() foundation.String
	RunModalForTrustMessage(trust unsafe.Pointer, message objc.IObject /* cross-framework: NSString */) int
	SetInformativeText(informativeText objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFCertificateTrustPanel */
// Alloc allocates a new instance without initialization.
func (sc _SFCertificateTrustPanelClass) Alloc() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFCertificateTrustPanelClass) New() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFCertificateTrustPanel) Init() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFCertificateTrustPanel) Autorelease() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFCertificateTrustPanel creates a new SFCertificateTrustPanel instance.
func NewSFCertificateTrustPanel() SFCertificateTrustPanel {
	return getSFCertificateTrustPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFCertificateTrustPanel */
// A panel or sheet that lets the user edit the trust settings in any of the certificates in a certificate chain.
//
// The following figure shows an example of a certificate trust panel. You can use this class to enable a user to make trust decisions when one or more certificates required for an operation are invalid or cannot be verified. To display a certificate in a panel or sheet without editable trust settings, use the class. To display certificates in a custom view, use the class.


// A panel or sheet that lets the user edit the trust settings in any of the certificates in a certificate chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel
type SFCertificateTrustPanel struct {
	SFCertificatePanel
}

// SFCertificateTrustPanelFrom constructs a [SFCertificateTrustPanel] from an unsafe.Pointer.
//
// A panel or sheet that lets the user edit the trust settings in any of the certificates in a certificate chain.
func SFCertificateTrustPanelFrom(ptr unsafe.Pointer) SFCertificateTrustPanel {
	return SFCertificateTrustPanel{
		SFCertificatePanel: SFCertificatePanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFCertificateTrustPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFCertificateTrustPanel */

// Returns a fully initialized, singleton certificate trust panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/shared()
func (sc _SFCertificateTrustPanelClass) SharedCertificateTrustPanel() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](objc.ID(sc.class), objc.Sel("sharedCertificateTrustPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCertificateTrustPanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFCertificateTrustPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFCertificateTrustPanel */

// Displays a modal sheet that shows the results of a certificate trust evaluation and that allows the user to edit trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:trust:message:)
func (s_ SFCertificateTrustPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustMessage(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, message objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:trust:message:"), docWindow, delegate, didEndSelector, contextInfo, trust, message)
}/* debug [instance_methods/method]: BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustMessage */


// Returns the (optional) informative text currently displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/informativeText()
func (s_ SFCertificateTrustPanel) InformativeText() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("informativeText"))
	return rv
}/* debug [instance_methods/method]: InformativeText */


// Displays a modal panel that shows the results of a certificate trust evaluation and that allows the user to edit trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/runModal(for:message:)
func (s_ SFCertificateTrustPanel) RunModalForTrustMessage(trust unsafe.Pointer, message objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForTrust:message:"), trust, message)
	return rv
}/* debug [instance_methods/method]: RunModalForTrustMessage */


// Sets the (optional) informative text displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/setInformativeText(_:)
func (s_ SFCertificateTrustPanel) SetInformativeText(informativeText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInformativeText:"), informativeText)
}/* debug [instance_methods/method]: SetInformativeText */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFCertificateTrustPanel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFCertificateTrustPanel */



