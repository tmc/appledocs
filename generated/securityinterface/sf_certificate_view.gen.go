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

/* debug [class.gen.go]: Generating class SFCertificateView */


/* debug [class_header]: Header for SFCertificateView */
// The class instance for the [SFCertificateView] class.
var (
	SFCertificateViewClass     _SFCertificateViewClass
	SFCertificateViewClassOnce sync.Once
)

func getSFCertificateViewClass() _SFCertificateViewClass {
	SFCertificateViewClassOnce.Do(func() {
		SFCertificateViewClass = _SFCertificateViewClass{objc.GetClass("SFCertificateView")}
	})
	return SFCertificateViewClass
}

type _SFCertificateViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFCertificateView */
// An interface definition for the [SFCertificateView] class.
type ISFCertificateView interface {
	appkit.IVisualEffectView
	
/* debug [class_interface_properties]: Properties for SFCertificateView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFCertificateView */
	// methods:
	Certificate() unsafe.Pointer
	DetailsDisclosed() bool
	DetailsDisplayed() bool
	IsEditable() bool
	IsTrustDisplayed() bool
	Policies() foundation.Array
	PoliciesDisclosed() bool
	SaveTrustSettings()
	SetCertificate(certificate unsafe.Pointer)
	SetDetailsDisclosed(disclosed bool)
	SetDisplayDetails(display bool)
	SetDisplayTrust(display bool)
	SetEditableTrust(editable bool)
	SetPolicies(policies objc.IObject)
	SetPoliciesDisclosed(disclosed bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFCertificateView */
// Alloc allocates a new instance without initialization.
func (sc _SFCertificateViewClass) Alloc() SFCertificateView {
	rv := objc.Send[SFCertificateView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFCertificateViewClass) New() SFCertificateView {
	rv := objc.Send[SFCertificateView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFCertificateView) Init() SFCertificateView {
	rv := objc.Send[SFCertificateView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFCertificateView) Autorelease() SFCertificateView {
	rv := objc.Send[SFCertificateView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFCertificateView creates a new SFCertificateView instance.
func NewSFCertificateView() SFCertificateView {
	return getSFCertificateViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFCertificateView */
// A view that displays the contents of a certificate, with options to display certificate details, display trust settings, and allow users to edit a certificate’s trust settings.
//
// The following figure shows a certificate view that includes editable trust settings and certificate details.


// A view that displays the contents of a certificate, with options to display certificate details, display trust settings, and allow users to edit a certificate’s trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView
type SFCertificateView struct {
	appkit.VisualEffectView
}

// SFCertificateViewFrom constructs a [SFCertificateView] from an unsafe.Pointer.
//
// A view that displays the contents of a certificate, with options to display certificate details, display trust settings, and allow users to edit a certificate’s trust settings.
func SFCertificateViewFrom(ptr unsafe.Pointer) SFCertificateView {
	return SFCertificateView{
		VisualEffectView: appkit.VisualEffectViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFCertificateView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFCertificateView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFCertificateView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFCertificateView */

// Returns the certificate currently displayed in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/certificate()
func (s_ SFCertificateView) Certificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("certificate"))
	return rv
}/* debug [instance_methods/method]: Certificate */


// Returns whether the view currently shows the certificate’s details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/detailsDisclosed()
func (s_ SFCertificateView) DetailsDisclosed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("detailsDisclosed"))
	return rv
}/* debug [instance_methods/method]: DetailsDisclosed */


// Indicates if the view currently shows the certificate’s details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/detailsDisplayed()
func (s_ SFCertificateView) DetailsDisplayed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("detailsDisplayed"))
	return rv
}/* debug [instance_methods/method]: DetailsDisplayed */


// Indicates if the view allows the user to edit the certificate’s trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/isEditable()
func (s_ SFCertificateView) IsEditable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_methods/method]: IsEditable */


// Indicates if the view currently shows the certificate’s trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/isTrustDisplayed()
func (s_ SFCertificateView) IsTrustDisplayed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isTrustDisplayed"))
	return rv
}/* debug [instance_methods/method]: IsTrustDisplayed */


// Returns an array of policies used to evaluate the status of the displayed certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/policies()
func (s_ SFCertificateView) Policies() foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("policies"))
	return rv
}/* debug [instance_methods/method]: Policies */


// Returns whether the trust policy subview is disclosed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/policiesDisclosed()
func (s_ SFCertificateView) PoliciesDisclosed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("policiesDisclosed"))
	return rv
}/* debug [instance_methods/method]: PoliciesDisclosed */


// Saves the user’s current trust settings for the displayed certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/saveTrustSettings()
func (s_ SFCertificateView) SaveTrustSettings() {
	objc.Send[objc.ID](s_.ID, objc.Sel("saveTrustSettings"))
}/* debug [instance_methods/method]: SaveTrustSettings */


// Specifies the certificate that’s displayed in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setCertificate(_:)
func (s_ SFCertificateView) SetCertificate(certificate unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCertificate:"), certificate)
}/* debug [instance_methods/method]: SetCertificate */


// Sets whether the certificate details subview is disclosed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setDetailsDisclosed(_:)
func (s_ SFCertificateView) SetDetailsDisclosed(disclosed bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDetailsDisclosed:"), disclosed)
}/* debug [instance_methods/method]: SetDetailsDisclosed */


// Specifies whether the user can see the certificate details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setDisplayDetails(_:)
func (s_ SFCertificateView) SetDisplayDetails(display bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayDetails:"), display)
}/* debug [instance_methods/method]: SetDisplayDetails */


// Specifies whether the user can see the certificate’s trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setDisplayTrust(_:)
func (s_ SFCertificateView) SetDisplayTrust(display bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayTrust:"), display)
}/* debug [instance_methods/method]: SetDisplayTrust */


// Specifies whether the user can edit the certificate’s trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setEditableTrust(_:)
func (s_ SFCertificateView) SetEditableTrust(editable bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEditableTrust:"), editable)
}/* debug [instance_methods/method]: SetEditableTrust */


// Specifies the policies to use when evaluating this certificate’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setPolicies(_:)
func (s_ SFCertificateView) SetPolicies(policies objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPolicies:"), policies)
}/* debug [instance_methods/method]: SetPolicies */


// Specifies whether the trust policy settings subview is disclosed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFCertificateView/setPoliciesDisclosed(_:)
func (s_ SFCertificateView) SetPoliciesDisclosed(disclosed bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPoliciesDisclosed:"), disclosed)
}/* debug [instance_methods/method]: SetPoliciesDisclosed */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFCertificateView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFCertificateView */



