// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFChooseIdentityPanel */

/* debug [class_header]: Header for SFChooseIdentityPanel */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFChooseIdentityPanel */
// An interface definition for the [SFChooseIdentityPanel] class.
type ISFChooseIdentityPanel interface {
	appkit.IPanel

	/* debug [class_interface_properties]: Properties for SFChooseIdentityPanel */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFChooseIdentityPanel */
	// methods:
	BeginSheetForWindowModalDelegateDidEndSelectorContextInfoIdentitiesMessage(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, identities objc.IObject /* cross-framework: NSArray */, message objc.IObject /* cross-framework: NSString */)
	Domain() foundation.String
	HelpAnchor() foundation.String
	Identity() unsafe.Pointer
	InformativeText() foundation.String
	Policies() foundation.Array
	RunModalForIdentitiesMessage(identities objc.IObject /* cross-framework: NSArray */, message objc.IObject /* cross-framework: NSString */) int
	SetAlternateButtonTitle(title objc.IObject /* cross-framework: NSString */)
	SetDefaultButtonTitle(title objc.IObject /* cross-framework: NSString */)
	SetDomain(domainString objc.IObject /* cross-framework: NSString */)
	SetHelpAnchor(anchor objc.IObject /* cross-framework: NSString */)
	SetInformativeText(informativeText objc.IObject /* cross-framework: NSString */)
	SetPolicies(policies objc.IObject)
	SetShowsHelp(showsHelp bool)
	ShowsHelp() bool
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFChooseIdentityPanel */
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFChooseIdentityPanel */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFChooseIdentityPanel */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFChooseIdentityPanel */

// Returns a fully initialized, singleton choose identity panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/shared()
func (sc _SFChooseIdentityPanelClass) SharedChooseIdentityPanel() SFChooseIdentityPanel {
	rv := objc.Send[SFChooseIdentityPanel](objc.ID(sc.class), objc.Sel("sharedChooseIdentityPanel"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=SharedChooseIdentityPanel) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFChooseIdentityPanel */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFChooseIdentityPanel */

// Displays a list of identities in a modal sheet from which the user can select an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/beginSheet(for:modalDelegate:didEnd:contextInfo:identities:message:)
func (s_ SFChooseIdentityPanel) BeginSheetForWindowModalDelegateDidEndSelectorContextInfoIdentitiesMessage(docWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer, identities objc.IObject /* cross-framework: NSArray */, message objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetForWindow:modalDelegate:didEndSelector:contextInfo:identities:message:"), docWindow, delegate, didEndSelector, contextInfo, identities, message)
} /* debug [instance_methods/method]: BeginSheetForWindowModalDelegateDidEndSelectorContextInfoIdentitiesMessage */

// Returns the domain that will be associated with the chosen identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/domain()
func (s_ SFChooseIdentityPanel) Domain() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("domain"))
	return rv
} /* debug [instance_methods/method]: Domain */

// Returns the current help anchor string for the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/helpAnchor()
func (s_ SFChooseIdentityPanel) HelpAnchor() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("helpAnchor"))
	return rv
} /* debug [instance_methods/method]: HelpAnchor */

// Returns the identity that the user chose in the panel or sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/identity()
func (s_ SFChooseIdentityPanel) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("identity"))
	return rv
} /* debug [instance_methods/method]: Identity */

// Returns the informative text currently displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/informativeText()
func (s_ SFChooseIdentityPanel) InformativeText() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("informativeText"))
	return rv
} /* debug [instance_methods/method]: InformativeText */

// Returns an array of policies used to evaluate the status of the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/policies()
func (s_ SFChooseIdentityPanel) Policies() foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("policies"))
	return rv
} /* debug [instance_methods/method]: Policies */

// Displays a list of identities in a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/runModal(forIdentities:message:)
func (s_ SFChooseIdentityPanel) RunModalForIdentitiesMessage(identities objc.IObject /* cross-framework: NSArray */, message objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](s_.ID, objc.Sel("runModalForIdentities:message:"), identities, message)
	return rv
} /* debug [instance_methods/method]: RunModalForIdentitiesMessage */

// Customizes the title of the alternate button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setAlternateButtonTitle(_:)
func (s_ SFChooseIdentityPanel) SetAlternateButtonTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateButtonTitle:"), title)
} /* debug [instance_methods/method]: SetAlternateButtonTitle */

// Customizes the title of the default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setDefaultButtonTitle(_:)
func (s_ SFChooseIdentityPanel) SetDefaultButtonTitle(title objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultButtonTitle:"), title)
} /* debug [instance_methods/method]: SetDefaultButtonTitle */

// Sets an optional domain in which the identity is to be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setDomain(_:)
func (s_ SFChooseIdentityPanel) SetDomain(domainString objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDomain:"), domainString)
} /* debug [instance_methods/method]: SetDomain */

// Sets the help anchor string for the sheet or modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setHelpAnchor(_:)
func (s_ SFChooseIdentityPanel) SetHelpAnchor(anchor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHelpAnchor:"), anchor)
} /* debug [instance_methods/method]: SetHelpAnchor */

// Sets the optional informative text displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setInformativeText(_:)
func (s_ SFChooseIdentityPanel) SetInformativeText(informativeText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInformativeText:"), informativeText)
} /* debug [instance_methods/method]: SetInformativeText */

// Specifies one or more policies that apply to the displayed certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setPolicies(_:)
func (s_ SFChooseIdentityPanel) SetPolicies(policies objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPolicies:"), policies)
} /* debug [instance_methods/method]: SetPolicies */

// Displays a Help button in the sheet or panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/setShowsHelp(_:)
func (s_ SFChooseIdentityPanel) SetShowsHelp(showsHelp bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHelp:"), showsHelp)
} /* debug [instance_methods/method]: SetShowsHelp */

// Indicates whether the help button is currently set to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityPanel/showsHelp()
func (s_ SFChooseIdentityPanel) ShowsHelp() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHelp"))
	return rv
} /* debug [instance_methods/method]: ShowsHelp */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFChooseIdentityPanel */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFChooseIdentityPanel */
