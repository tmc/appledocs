// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [CredentialProviderViewController] class.
type ICredentialProviderViewController interface {
	appkit.IViewController
	PerformPasskeyRegistrationWithoutUserInteractionIfPossible(registrationRequest unsafe.Pointer)
	PrepareCredentialListForServiceIdentifiers(serviceIdentifiers unsafe.Pointer)
	PrepareInterfaceForUserChoosingTextToInsert()
	PrepareInterfaceToProvideCredentialForRequest(credentialRequest objc.ID)
	ProvideCredentialWithoutUserInteractionForRequest(credentialRequest objc.ID)
}

// A view controller that a credential manager app uses to extend AutoFill.
//
// To integrate a password, passkey, or one-time passcode manager app with AutoFill: Add a Credential Provider Extension target to your project that subclasses . Add the to both the extension and its containing app. Override the view controller’s method to prepare a view with a list of credentials that the person can choose from after opening your extension from the AutoFill suggestions list. Optionally add and instances to the shared to make identities available directly in the AutoFill suggestions list. Then override the method to provide the associated credentials when the person taps a suggestion. Optionally, override the method to specify a configuration interface that you can show when people first enable your credentials manager in Settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController
type CredentialProviderViewController struct {
	appkit.ViewController
}

// CredentialProviderViewControllerFrom constructs a [CredentialProviderViewController] from an unsafe.Pointer.
//
// A view controller that a credential manager app uses to extend AutoFill.
func CredentialProviderViewControllerFrom(ptr unsafe.Pointer) CredentialProviderViewController {
	return CredentialProviderViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CredentialProviderViewControllerClass) Alloc() CredentialProviderViewController {
	rv := objc.Send[CredentialProviderViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Perform a conditional passkey registration, if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/performWithoutUserInteractionIfPossible(passkeyRegistration:)
func (c_ CredentialProviderViewController) PerformPasskeyRegistrationWithoutUserInteractionIfPossible(registrationRequest unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performPasskeyRegistrationWithoutUserInteractionIfPossible:"), registrationRequest)
}

// Prepares the interface to display a list of credentials from which the user can select.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareCredentialList(for:)
func (c_ CredentialProviderViewController) PrepareCredentialListForServiceIdentifiers(serviceIdentifiers unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareCredentialListForServiceIdentifiers:"), serviceIdentifiers)
}

// Prepare the view controller to show a list of all insertable text with user selectable fields.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterfaceForUserChoosingTextToInsert()
func (c_ CredentialProviderViewController) PrepareInterfaceForUserChoosingTextToInsert() {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceForUserChoosingTextToInsert"))
}

// Prepare the view controller to show user interface for providing the requested credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterfaceToProvideCredential(for:)-68qpo
func (c_ CredentialProviderViewController) PrepareInterfaceToProvideCredentialForRequest(credentialRequest objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceToProvideCredentialForRequest:"), credentialRequest)
}

// Attempts to provide the user-requested credential with no further user interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/provideCredentialWithoutUserInteraction(for:)-3mo23
func (c_ CredentialProviderViewController) ProvideCredentialWithoutUserInteractionForRequest(credentialRequest objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("provideCredentialWithoutUserInteractionForRequest:"), credentialRequest)
}

// The context your credential provider extension uses to provide information to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/ascredentialproviderviewcontroller/extensioncontext
func (c_ CredentialProviderViewController) ExtensionContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("extensionContext"))
	return rv
}


// SetExtensionContext sets the value of the extensionContext property.
// The context your credential provider extension uses to provide information to the system.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/ascredentialproviderviewcontroller/extensioncontext
func (c_ CredentialProviderViewController) SetExtensionContext(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtensionContext:"), value)
}

// The domain for a credential provider extension error.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asextensionerrordomain
func (c_ CredentialProviderViewController) ASExtensionErrorDomain() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ASExtensionErrorDomain"))
	return rv
}



