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

/* debug [class.gen.go]: Generating class SFCertificatePanel */


/* debug [class_header]: Header for SFCertificatePanel */
// The class instance for the [SFCertificatePanel] class.
var (
	SFCertificatePanelClass     _SFCertificatePanelClass
	SFCertificatePanelClassOnce sync.Once
)

func getSFCertificatePanelClass() _SFCertificatePanelClass {
	SFCertificatePanelClassOnce.Do(func() {
		SFCertificatePanelClass = _SFCertificatePanelClass{objc.GetClass("SFCertificatePanel")}
	})
	return SFCertificatePanelClass
}

type _SFCertificatePanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFCertificatePanel */
// An interface definition for the [SFCertificatePanel] class.
type ISFCertificatePanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for SFCertificatePanel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFCertificatePanel */
	// methods:
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoCertificatesShowGroup(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, certificates objc.IObject /* cross-framework: NSArray */, showGroup bool)
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustShowGroup(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, showGroup bool)
	CertificateView() ISFCertificateView
	HelpAnchor() foundation.String
	Policies() foundation.Array
	RunModalForTrustShowGroup(trust unsafe.Pointer, showGroup bool) int
	RunModalForCertificatesShowGroup(certificates objc.IObject /* cross-framework: NSArray */, showGroup bool) int
	SetAlternateButtonTitle(title objc.IObject /* cross-framework: NSString */)
	SetDefaultButtonTitle(title objc.IObject /* cross-framework: NSString */)
	SetHelpAnchor(anchor objc.IObject /* cross-framework: NSString */)
	SetPolicies(policies objc.IObject)
	SetShowsHelp(showsHelp bool)
	ShowsHelp() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFCertificatePanel */
// Alloc allocates a new instance without initialization.
func (sc _SFCertificatePanelClass) Alloc() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFCertificatePanelClass) New() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFCertificatePanel) Init() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFCertificatePanel) Autorelease() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFCertificatePanel creates a new SFCertificatePanel instance.
func NewSFCertificatePanel() SFCertificatePanel {
	return getSFCertificatePanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFCertificatePanel */
// A panel or sheet that displays one or more certificates.
//
// The following figure shows an example of a certificate panel. An can optionally display all of the certificates in a certificate chain. This class displays certificate details, but not trust settings. To display a certificate with editable trust settings in a panel or sheet, use the class. To display certificates in a custom view, use the class. Note that for macOS 10.4 and later, this class displays the evaluation status for each certificate. You can modify how the certificates are evaluated by calling the method.


// A panel or sheet that displays one or more certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel
type SFCertificatePanel struct {
	appkit.Panel
}

// SFCertificatePanelFrom constructs a [SFCertificatePanel] from an unsafe.Pointer.
//
// A panel or sheet that displays one or more certificates.
func SFCertificatePanelFrom(ptr unsafe.Pointer) SFCertificatePanel {
	return SFCertificatePanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFCertificatePanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFCertificatePanel */

// Returns a fully initialized, singleton certificate panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/shared()
func (sc _SFCertificatePanelClass) SharedCertificatePanel() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](objc.ID(sc.class), objc.Sel("sharedCertificatePanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCertificatePanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFCertificatePanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFCertificatePanel */

// Displays one or more certificates in a modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/beginSheet(for:modalDelegate:didEnd:contextInfo:certificates:showGroup:)
func (s_ SFCertificatePanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoCertificatesShowGroup(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, certificates objc.IObject /* cross-framework: NSArray */, showGroup bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:certificates:showGroup:"), docWindow, delegate, didEndSelector, contextInfo, certificates, showGroup)
}/* debug [instance_methods/method]: BeginSheetForWindowModalDelegateDidEndSelectorContextInfoCertificatesShowGroup */


// Displays a certificate chain in a modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/beginSheet(for:modalDelegate:didEnd:contextInfo:trust:showGroup:)
func (s_ SFCertificatePanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustShowGroup(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, showGroup bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:trust:showGroup:"), docWindow, delegate, didEndSelector, contextInfo, trust, showGroup)
}/* debug [instance_methods/method]: BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustShowGroup */


// Returns the certificate view for the modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/certificateView()
func (s_ SFCertificatePanel) CertificateView() ISFCertificateView {
	rv := objc.Send[SFCertificateView](s_.ID, objc.Sel("certificateView"))
	return rv
}/* debug [instance_methods/method]: CertificateView */


// Returns the current help anchor string for the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/helpAnchor()
func (s_ SFCertificatePanel) HelpAnchor() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("helpAnchor"))
	return rv
}/* debug [instance_methods/method]: HelpAnchor */


// Returns an array of policies used to evaluate the status of the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/policies()
func (s_ SFCertificatePanel) Policies() foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("policies"))
	return rv
}/* debug [instance_methods/method]: Policies */


// Displays a certificate chain in a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/runModal(for:showGroup:)
func (s_ SFCertificatePanel) RunModalForTrustShowGroup(trust unsafe.Pointer, showGroup bool) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForTrust:showGroup:"), trust, showGroup)
	return rv
}/* debug [instance_methods/method]: RunModalForTrustShowGroup */


// Displays one or more specified certificates in a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/runModal(forCertificates:showGroup:)
func (s_ SFCertificatePanel) RunModalForCertificatesShowGroup(certificates objc.IObject /* cross-framework: NSArray */, showGroup bool) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForCertificates:showGroup:"), certificates, showGroup)
	return rv
}/* debug [instance_methods/method]: RunModalForCertificatesShowGroup */


// Customizes the title of the alternate button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setAlternateButtonTitle(_:)
func (s_ SFCertificatePanel) SetAlternateButtonTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateButtonTitle:"), title)
}/* debug [instance_methods/method]: SetAlternateButtonTitle */


// Customizes the title of the default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setDefaultButtonTitle(_:)
func (s_ SFCertificatePanel) SetDefaultButtonTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultButtonTitle:"), title)
}/* debug [instance_methods/method]: SetDefaultButtonTitle */


// Sets the help anchor string for the sheet or modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setHelpAnchor(_:)
func (s_ SFCertificatePanel) SetHelpAnchor(anchor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHelpAnchor:"), anchor)
}/* debug [instance_methods/method]: SetHelpAnchor */


// Specifies one or more policies that apply to the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setPolicies(_:)
func (s_ SFCertificatePanel) SetPolicies(policies objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPolicies:"), policies)
}/* debug [instance_methods/method]: SetPolicies */


// Displays a Help button in the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setShowsHelp(_:)
func (s_ SFCertificatePanel) SetShowsHelp(showsHelp bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHelp:"), showsHelp)
}/* debug [instance_methods/method]: SetShowsHelp */


// Indicates whether the help button is currently set to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/showsHelp()
func (s_ SFCertificatePanel) ShowsHelp() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHelp"))
	return rv
}/* debug [instance_methods/method]: ShowsHelp */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFCertificatePanel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFCertificatePanel */



