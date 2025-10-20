// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [SFCertificateTrustPanel] class.
type ISFCertificateTrustPanel interface {
	ISFCertificatePanel
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustMessage(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, message string)
	InformativeText() unsafe.Pointer
	RunModalForTrustMessage(trust unsafe.Pointer, message string) int
	SetInformativeText(informativeText string)
}

// A panel or sheet that lets the user edit the trust settings in any of the certificates in a certificate chain.
//
// The following figure shows an example of a certificate trust panel. You can use this class to enable a user to make trust decisions when one or more certificates required for an operation are invalid or cannot be verified. To display a certificate in a panel or sheet without editable trust settings, use the class. To display certificates in a custom view, use the class.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFCertificateTrustPanelClass) Alloc() SFCertificateTrustPanel {
	rv := objc.Send[SFCertificateTrustPanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a fully initialized, singleton certificate trust panel object.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/shared()
func (sc _SFCertificateTrustPanelClass) SharedCertificateTrustPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharedCertificateTrustPanel"))
	return rv
}

// Displays a modal sheet that shows the results of a certificate trust evaluation and that allows the user to edit trust settings.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:trust:message:)
func (s_ SFCertificateTrustPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoTrustMessage(docWindow unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer, trust unsafe.Pointer, message string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:trust:message:"), docWindow, delegate, didEndSelector, contextInfo, trust, objc.String(message))
}

// Returns the (optional) informative text currently displayed in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/informativeText()
func (s_ SFCertificateTrustPanel) InformativeText() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("informativeText"))
	return rv
}

// Displays a modal panel that shows the results of a certificate trust evaluation and that allows the user to edit trust settings.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/runModal(for:message:)
func (s_ SFCertificateTrustPanel) RunModalForTrustMessage(trust unsafe.Pointer, message string) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForTrust:message:"), trust, objc.String(message))
	return rv
}

// Sets the (optional) informative text displayed in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateTrustPanel/setInformativeText(_:)
func (s_ SFCertificateTrustPanel) SetInformativeText(informativeText string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInformativeText:"), objc.String(informativeText))
}



