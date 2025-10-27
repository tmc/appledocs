// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SFChooseIdentityPanel] class.
var (
	SFChooseIdentityPanelClass     _SFChooseIdentityPanelClass
	SFChooseIdentityPanelClassOnce sync.Once
)

func getSFChooseIdentityPanelClass() _SFChooseIdentityPanelClass {
	SFChooseIdentityPanelClassOnce.Do(func() {
		SFChooseIdentityPanelClass = _SFChooseIdentityPanelClass{objc.GetClass("SFChooseIdentityPanel")}
	})
	return SFChooseIdentityPanelClass
}

type _SFChooseIdentityPanelClass struct {
	class objc.Class
}





// An interface definition for the [SFChooseIdentityPanel] class.
type ISFChooseIdentityPanel interface {
	appkit.IPanel
	

	// properties:


	

	// methods:
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoIdentitiesMessage(docWindow appkit.Window, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, identities foundation.foundation.INSArray, message foundation.foundation.INSString)
	Domain() foundation.String
	HelpAnchor() foundation.String
	Identity() unsafe.Pointer
	InformativeText() foundation.String
	Policies() foundation.Array
	RunModalForIdentitiesMessage(identities foundation.foundation.INSArray, message foundation.foundation.INSString) int
	SetAlternateButtonTitle(title foundation.foundation.INSString)
	SetDefaultButtonTitle(title foundation.foundation.INSString)
	SetDomain(domainString foundation.foundation.INSString)
	SetHelpAnchor(anchor foundation.foundation.INSString)
	SetInformativeText(informativeText foundation.foundation.INSString)
	SetPolicies(policies objectivec.IObject)
	SetShowsHelp(showsHelp bool)
	ShowsHelp() bool


}





// Alloc allocates a new instance without initialization.
func (sc _SFChooseIdentityPanelClass) Alloc() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFChooseIdentityPanelClass) New() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFChooseIdentityPanel) Init() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFChooseIdentityPanel) Autorelease() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFChooseIdentityPanel creates a new SFChooseIdentityPanel instance.
func NewSFChooseIdentityPanel() SFChooseIdentityPanel {
	return getSFChooseIdentityPanelClass().New()
}





// A panel or sheet containing a list of identities that a user can choose from.
//
// An identity is a digital certificate together with its associated private key. This class also allows the user to display the contents of any certificate in the list. The following figure shows an example of a choose identity panel.


// A panel or sheet containing a list of identities that a user can choose from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel
type SFChooseIdentityPanel struct {
	appkit.Panel
}

// SFChooseIdentityPanelFrom constructs a [SFChooseIdentityPanel] from an unsafe.Pointer.
//
// A panel or sheet containing a list of identities that a user can choose from.
func SFChooseIdentityPanelFrom(ptr unsafe.Pointer) SFChooseIdentityPanel {
	return SFChooseIdentityPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}










// Returns a fully initialized, singleton choose identity panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/shared()
func (sc _SFChooseIdentityPanelClass) SharedChooseIdentityPanel() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](objc.ID(sc.class), objc.Sel("sharedChooseIdentityPanel"))
	return rv
}












// Displays a list of identities in a modal sheet from which the user can select an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:identities:message:)
func (s_ SFChooseIdentityPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoIdentitiesMessage(docWindow appkit.Window, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, identities foundation.foundation.INSArray, message foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:identities:message:"), docWindow, delegate, didEndSelector, contextInfo, identities, message)
}


// Returns the domain that will be associated with the chosen identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/domain()
func (s_ SFChooseIdentityPanel) Domain() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("domain"))
	return rv
}


// Returns the current help anchor string for the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/helpAnchor()
func (s_ SFChooseIdentityPanel) HelpAnchor() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("helpAnchor"))
	return rv
}


// Returns the identity that the user chose in the panel or sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/identity()
func (s_ SFChooseIdentityPanel) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("identity"))
	return rv
}


// Returns the informative text currently displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/informativeText()
func (s_ SFChooseIdentityPanel) InformativeText() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("informativeText"))
	return rv
}


// Returns an array of policies used to evaluate the status of the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/policies()
func (s_ SFChooseIdentityPanel) Policies() foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("policies"))
	return rv
}


// Displays a list of identities in a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/runModal(forIdentities:message:)
func (s_ SFChooseIdentityPanel) RunModalForIdentitiesMessage(identities foundation.foundation.INSArray, message foundation.foundation.INSString) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForIdentities:message:"), identities, message)
	return rv
}


// Customizes the title of the alternate button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setAlternateButtonTitle(_:)
func (s_ SFChooseIdentityPanel) SetAlternateButtonTitle(title foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateButtonTitle:"), title)
}


// Customizes the title of the default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setDefaultButtonTitle(_:)
func (s_ SFChooseIdentityPanel) SetDefaultButtonTitle(title foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultButtonTitle:"), title)
}


// Sets an optional domain in which the identity is to be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setDomain(_:)
func (s_ SFChooseIdentityPanel) SetDomain(domainString foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDomain:"), domainString)
}


// Sets the help anchor string for the sheet or modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setHelpAnchor(_:)
func (s_ SFChooseIdentityPanel) SetHelpAnchor(anchor foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHelpAnchor:"), anchor)
}


// Sets the optional informative text displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setInformativeText(_:)
func (s_ SFChooseIdentityPanel) SetInformativeText(informativeText foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInformativeText:"), informativeText)
}


// Specifies one or more policies that apply to the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setPolicies(_:)
func (s_ SFChooseIdentityPanel) SetPolicies(policies objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPolicies:"), policies)
}


// Displays a Help button in the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setShowsHelp(_:)
func (s_ SFChooseIdentityPanel) SetShowsHelp(showsHelp bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHelp:"), showsHelp)
}


// Indicates whether the help button is currently set to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/showsHelp()
func (s_ SFChooseIdentityPanel) ShowsHelp() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHelp"))
	return rv
}













