// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASCredentialProviderViewController */


/* debug [class_header]: Header for ASCredentialProviderViewController */
// The class instance for the [CredentialProviderViewController] class.
var (
	CredentialProviderViewControllerClass     _CredentialProviderViewControllerClass
	CredentialProviderViewControllerClassOnce sync.Once
)

func getCredentialProviderViewControllerClass() _CredentialProviderViewControllerClass {
	CredentialProviderViewControllerClassOnce.Do(func() {
		CredentialProviderViewControllerClass = _CredentialProviderViewControllerClass{objc.GetClass("ASCredentialProviderViewController")}
	})
	return CredentialProviderViewControllerClass
}

type _CredentialProviderViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CredentialProviderViewController */
// An interface definition for the [CredentialProviderViewController] class.
type ICredentialProviderViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for CredentialProviderViewController */
	// properties:
	ExtensionContext() IASCredentialProviderExtensionContext
	ASExtensionErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CredentialProviderViewController */
	// methods:
	PerformPasskeyRegistrationWithoutUserInteractionIfPossible(registrationRequest IASPasskeyCredentialRequest)
	PrepareCredentialListForServiceIdentifiers(serviceIdentifiers []CredentialServiceIdentifier)
	PrepareCredentialListForServiceIdentifiersRequestParameters(serviceIdentifiers []CredentialServiceIdentifier, requestParameters IASPasskeyCredentialRequestParameters)
	PrepareInterfaceForPasskeyRegistration(registrationRequest unsafe.Pointer)
	PrepareInterfaceForExtensionConfiguration()
	PrepareInterfaceToProvideCredentialForRequest(credentialRequest unsafe.Pointer)
	PrepareOneTimeCodeCredentialListForServiceIdentifiers(serviceIdentifiers []CredentialServiceIdentifier)
	ProvideCredentialWithoutUserInteractionForRequest(credentialRequest unsafe.Pointer)
	ReportAllAcceptedPublicKeyCredentialsForRelyingPartyUserHandleAcceptedCredentialIDs(relyingParty objc.IObject /* cross-framework: NSString */, userHandle objc.IObject /* cross-framework: NSData */, acceptedCredentialIDs []foundation.Data)
	ReportPublicKeyCredentialUpdateForRelyingPartyUserHandleNewName(relyingParty objc.IObject /* cross-framework: NSString */, userHandle objc.IObject /* cross-framework: NSData */, newName objc.IObject /* cross-framework: NSString */)
	ReportUnknownPublicKeyCredentialForRelyingPartyCredentialID(relyingParty objc.IObject /* cross-framework: NSString */, credentialID objc.IObject /* cross-framework: NSData */)
	ReportUnusedPasswordCredentialForDomainUserName(domain objc.IObject /* cross-framework: NSString */, userName objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CredentialProviderViewController */
// Alloc allocates a new instance without initialization.
func (cc _CredentialProviderViewControllerClass) Alloc() CredentialProviderViewController {
	rv := objc.Send[CredentialProviderViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CredentialProviderViewControllerClass) New() CredentialProviderViewController {
	rv := objc.Send[CredentialProviderViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CredentialProviderViewController) Init() CredentialProviderViewController {
	rv := objc.Send[CredentialProviderViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CredentialProviderViewController) Autorelease() CredentialProviderViewController {
	rv := objc.Send[CredentialProviderViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCredentialProviderViewController creates a new CredentialProviderViewController instance.
func NewCredentialProviderViewController() CredentialProviderViewController {
	return getCredentialProviderViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CredentialProviderViewController */
// A view controller that a credential manager app uses to extend AutoFill.
//
// To integrate a password, passkey, or one-time passcode manager app with AutoFill: Add a Credential Provider Extension target to your project that subclasses . Add the to both the extension and its containing app. Override the view controller’s method to prepare a view with a list of credentials that the person can choose from after opening your extension from the AutoFill suggestions list. Optionally add and instances to the shared to make identities available directly in the AutoFill suggestions list. Then override the method to provide the associated credentials when the person taps a suggestion. Optionally, override the method to specify a configuration interface that you can show when people first enable your credentials manager in Settings.


// A view controller that a credential manager app uses to extend AutoFill.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController
type CredentialProviderViewController struct {
	ViewController
}

// CredentialProviderViewControllerFrom constructs a [CredentialProviderViewController] from an unsafe.Pointer.
//
// A view controller that a credential manager app uses to extend AutoFill.
func CredentialProviderViewControllerFrom(ptr unsafe.Pointer) CredentialProviderViewController {
	return CredentialProviderViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CredentialProviderViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CredentialProviderViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CredentialProviderViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CredentialProviderViewController */

// Perform a conditional passkey registration, if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/performWithoutUserInteractionIfPossible(passkeyRegistration:)
func (c_ CredentialProviderViewController) PerformPasskeyRegistrationWithoutUserInteractionIfPossible(registrationRequest IASPasskeyCredentialRequest) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performPasskeyRegistrationWithoutUserInteractionIfPossible:"), registrationRequest)
}/* debug [instance_methods/method]: PerformPasskeyRegistrationWithoutUserInteractionIfPossible */


// Prepares the interface to display a list of credentials from which the user can select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareCredentialList(for:)
func (c_ CredentialProviderViewController) PrepareCredentialListForServiceIdentifiers(serviceIdentifiers []CredentialServiceIdentifier) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareCredentialListForServiceIdentifiers:"), serviceIdentifiers)
}/* debug [instance_methods/method]: PrepareCredentialListForServiceIdentifiers */


// Prepares the interface to display a list of passkey and password credentials from which the user can select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareCredentialList(for:requestParameters:)
func (c_ CredentialProviderViewController) PrepareCredentialListForServiceIdentifiersRequestParameters(serviceIdentifiers []CredentialServiceIdentifier, requestParameters IASPasskeyCredentialRequestParameters) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareCredentialListForServiceIdentifiers:requestParameters:"), serviceIdentifiers, requestParameters)
}/* debug [instance_methods/method]: PrepareCredentialListForServiceIdentifiersRequestParameters */


// Prepare the view controller to show user interface for registering a new passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterface(forPasskeyRegistration:)
func (c_ CredentialProviderViewController) PrepareInterfaceForPasskeyRegistration(registrationRequest unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceForPasskeyRegistration:"), registrationRequest)
}/* debug [instance_methods/method]: PrepareInterfaceForPasskeyRegistration */


// Prepares the interface to enable the user to configure the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterfaceForExtensionConfiguration()
func (c_ CredentialProviderViewController) PrepareInterfaceForExtensionConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceForExtensionConfiguration"))
}/* debug [instance_methods/method]: PrepareInterfaceForExtensionConfiguration */


// Prepare the view controller to show user interface for providing the requested credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterfaceToProvideCredential(for:)-68qpo
func (c_ CredentialProviderViewController) PrepareInterfaceToProvideCredentialForRequest(credentialRequest unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceToProvideCredentialForRequest:"), credentialRequest)
}/* debug [instance_methods/method]: PrepareInterfaceToProvideCredentialForRequest */


// Prepares the interface to display a list of one-time passcodes (OTPs) that people can select from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareOneTimeCodeCredentialList(for:)
func (c_ CredentialProviderViewController) PrepareOneTimeCodeCredentialListForServiceIdentifiers(serviceIdentifiers []CredentialServiceIdentifier) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareOneTimeCodeCredentialListForServiceIdentifiers:"), serviceIdentifiers)
}/* debug [instance_methods/method]: PrepareOneTimeCodeCredentialListForServiceIdentifiers */


// Attempts to provide the user-requested credential with no further user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/provideCredentialWithoutUserInteraction(for:)-3mo23
func (c_ CredentialProviderViewController) ProvideCredentialWithoutUserInteractionForRequest(credentialRequest unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("provideCredentialWithoutUserInteractionForRequest:"), credentialRequest)
}/* debug [instance_methods/method]: ProvideCredentialWithoutUserInteractionForRequest */


// Receives a report from the system that a relying party sent a snapshot of all accepted credentials for an account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/reportAllAcceptedPublicKeyCredentials(forRelyingParty:userHandle:acceptedCredentialIDs:)
func (c_ CredentialProviderViewController) ReportAllAcceptedPublicKeyCredentialsForRelyingPartyUserHandleAcceptedCredentialIDs(relyingParty objc.IObject /* cross-framework: NSString */, userHandle objc.IObject /* cross-framework: NSData */, acceptedCredentialIDs []foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportAllAcceptedPublicKeyCredentialsForRelyingParty:userHandle:acceptedCredentialIDs:"), relyingParty, userHandle, acceptedCredentialIDs)
}/* debug [instance_methods/method]: ReportAllAcceptedPublicKeyCredentialsForRelyingPartyUserHandleAcceptedCredentialIDs */


// Receives a report from the system that a relying party indicated that a passkey’s user name updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/reportPublicKeyCredentialUpdate(forRelyingParty:userHandle:newName:)
func (c_ CredentialProviderViewController) ReportPublicKeyCredentialUpdateForRelyingPartyUserHandleNewName(relyingParty objc.IObject /* cross-framework: NSString */, userHandle objc.IObject /* cross-framework: NSData */, newName objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportPublicKeyCredentialUpdateForRelyingParty:userHandle:newName:"), relyingParty, userHandle, newName)
}/* debug [instance_methods/method]: ReportPublicKeyCredentialUpdateForRelyingPartyUserHandleNewName */


// Receives a report from the system that a relying party indicated a passkey credential is invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/reportUnknownPublicKeyCredential(forRelyingParty:credentialID:)
func (c_ CredentialProviderViewController) ReportUnknownPublicKeyCredentialForRelyingPartyCredentialID(relyingParty objc.IObject /* cross-framework: NSString */, credentialID objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportUnknownPublicKeyCredentialForRelyingParty:credentialID:"), relyingParty, credentialID)
}/* debug [instance_methods/method]: ReportUnknownPublicKeyCredentialForRelyingPartyCredentialID */


// Receives a report from the system that a relying party indicatd that a password credential isn’t needed anymore for a given user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/reportUnusedPasswordCredential(forDomain:userName:)
func (c_ CredentialProviderViewController) ReportUnusedPasswordCredentialForDomainUserName(domain objc.IObject /* cross-framework: NSString */, userName objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportUnusedPasswordCredentialForDomain:userName:"), domain, userName)
}/* debug [instance_methods/method]: ReportUnusedPasswordCredentialForDomainUserName */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CredentialProviderViewController */

// The context your credential provider extension uses to provide information to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/extensionContext
func (c_ CredentialProviderViewController) ExtensionContext() IASCredentialProviderExtensionContext {
	rv := objc.Send[CredentialProviderExtensionContext](c_.ID, objc.Sel("extensionContext"))
	return rv
}/* debug [instance_properties/getter]: extensionContext */


// The domain for a credential provider extension error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asextensionerrordomain
func (c_ CredentialProviderViewController) ASExtensionErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ASExtensionErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: ASExtensionErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASCredentialProviderViewController */


