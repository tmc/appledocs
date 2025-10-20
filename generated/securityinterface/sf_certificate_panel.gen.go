// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [SFCertificatePanel] class.
type ISFCertificatePanel interface {
	appkit.IPanel
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoCertificatesShowGroup(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, certificates objc.ID, showGroup bool)
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustShowGroup(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, showGroup bool)
	CertificateView() unsafe.Pointer
	HelpAnchor() unsafe.Pointer
	Policies() unsafe.Pointer
	RunModalForTrustShowGroup(trust unsafe.Pointer, showGroup bool) int
	RunModalForCertificatesShowGroup(certificates objc.ID, showGroup bool) int
	SetAlternateButtonTitle(title string)
	SetDefaultButtonTitle(title string)
	SetHelpAnchor(anchor string)
	SetPolicies(policies objc.ID)
	SetShowsHelp(showsHelp bool)
	ShowsHelp() bool
}

// A panel or sheet that displays one or more certificates.
//
// The following figure shows an example of a certificate panel. An can optionally display all of the certificates in a certificate chain. This class displays certificate details, but not trust settings. To display a certificate with editable trust settings in a panel or sheet, use the class. To display certificates in a custom view, use the class. Note that for macOS 10.4 and later, this class displays the evaluation status for each certificate. You can modify how the certificates are evaluated by calling the method.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFCertificatePanelClass) Alloc() SFCertificatePanel {
	rv := objc.Send[SFCertificatePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a fully initialized, singleton certificate panel object.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/shared()
func (sc _SFCertificatePanelClass) SharedCertificatePanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharedCertificatePanel"))
	return rv
}

// Displays one or more certificates in a modal sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/beginSheet(for:modalDelegate:didEnd:contextInfo:certificates:showGroup:)
func (s_ SFCertificatePanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoCertificatesShowGroup(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, certificates objc.ID, showGroup bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:certificates:showGroup:"), docWindow, delegate, didEndSelector, contextInfo, certificates, showGroup)
}

// Displays a certificate chain in a modal sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/beginSheet(for:modalDelegate:didEnd:contextInfo:trust:showGroup:)
func (s_ SFCertificatePanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustShowGroup(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, showGroup bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:trust:showGroup:"), docWindow, delegate, didEndSelector, contextInfo, trust, showGroup)
}

// Returns the certificate view for the modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/certificateView()
func (s_ SFCertificatePanel) CertificateView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("certificateView"))
	return rv
}

// Returns the current help anchor string for the sheet or panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/helpAnchor()
func (s_ SFCertificatePanel) HelpAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("helpAnchor"))
	return rv
}

// Returns an array of policies used to evaluate the status of the displayed certificates.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/policies()
func (s_ SFCertificatePanel) Policies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("policies"))
	return rv
}

// Displays a certificate chain in a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/runModal(for:showGroup:)
func (s_ SFCertificatePanel) RunModalForTrustShowGroup(trust unsafe.Pointer, showGroup bool) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForTrust:showGroup:"), trust, showGroup)
	return rv
}

// Displays one or more specified certificates in a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/runModal(forCertificates:showGroup:)
func (s_ SFCertificatePanel) RunModalForCertificatesShowGroup(certificates objc.ID, showGroup bool) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForCertificates:showGroup:"), certificates, showGroup)
	return rv
}

// Customizes the title of the alternate button.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setAlternateButtonTitle(_:)
func (s_ SFCertificatePanel) SetAlternateButtonTitle(title string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateButtonTitle:"), objc.String(title))
}

// Customizes the title of the default button.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setDefaultButtonTitle(_:)
func (s_ SFCertificatePanel) SetDefaultButtonTitle(title string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultButtonTitle:"), objc.String(title))
}

// Sets the help anchor string for the sheet or modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setHelpAnchor(_:)
func (s_ SFCertificatePanel) SetHelpAnchor(anchor string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHelpAnchor:"), objc.String(anchor))
}

// Specifies one or more policies that apply to the displayed certificates.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setPolicies(_:)
func (s_ SFCertificatePanel) SetPolicies(policies objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPolicies:"), policies)
}

// Displays a Help button in the sheet or panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/setShowsHelp(_:)
func (s_ SFCertificatePanel) SetShowsHelp(showsHelp bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHelp:"), showsHelp)
}

// Indicates whether the help button is currently set to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificatePanel/showsHelp()
func (s_ SFCertificatePanel) ShowsHelp() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHelp"))
	return rv
}



