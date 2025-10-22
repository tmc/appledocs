// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasswordCredentialIdentity] class.
var (
	PasswordCredentialIdentityClass     _PasswordCredentialIdentityClass
	PasswordCredentialIdentityClassOnce sync.Once
)

func getPasswordCredentialIdentityClass() _PasswordCredentialIdentityClass {
	PasswordCredentialIdentityClassOnce.Do(func() {
		PasswordCredentialIdentityClass = _PasswordCredentialIdentityClass{objc.GetClass("ASPasswordCredentialIdentity")}
	})
	return PasswordCredentialIdentityClass
}

type _PasswordCredentialIdentityClass struct {
	class objc.Class
}

// An interface definition for the [PasswordCredentialIdentity] class.
type IPasswordCredentialIdentity interface {
	objectivec.IObject
	Rank() int
	SetRank(value int)
	RecordIdentifier() string
	ServiceIdentifier() unsafe.Pointer
	User() string
}

// A description that uniquely identifies a particular password credential.


// A description that uniquely identifies a particular password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity

type PasswordCredentialIdentity struct {
	objectivec.Object
}

// PasswordCredentialIdentityFrom constructs a [PasswordCredentialIdentity] from an unsafe.Pointer.
//
// A description that uniquely identifies a particular password credential.
func PasswordCredentialIdentityFrom(ptr unsafe.Pointer) PasswordCredentialIdentity {
	return PasswordCredentialIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasswordCredentialIdentityClass) Alloc() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasswordCredentialIdentityClass) New() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasswordCredentialIdentity) Init() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasswordCredentialIdentity) Autorelease() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasswordCredentialIdentity creates a new PasswordCredentialIdentity instance.
func NewPasswordCredentialIdentity() PasswordCredentialIdentity {
	return getPasswordCredentialIdentityClass().New()
}




// Initializes a password credential identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/init(serviceIdentifier:user:recordIdentifier:)

func NewPasswordCredentialIdentityWithServiceIdentifierUserRecordIdentifier(serviceIdentifier unsafe.Pointer, user string, recordIdentifier string) PasswordCredentialIdentity {
	instance := getPasswordCredentialIdentityClass().Alloc()
	rv := objc.Send[PasswordCredentialIdentity](instance.ID, objc.Sel("initWithServiceIdentifier:user:recordIdentifier:"), serviceIdentifier, objc.String(user), objc.String(recordIdentifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a password credential identity object with a service identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/identityWithServiceIdentifier:user:recordIdentifier:

func (pc _PasswordCredentialIdentityClass) IdentityWithServiceIdentifierUserRecordIdentifier(serviceIdentifier unsafe.Pointer, user string, recordIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("identityWithServiceIdentifier:user:recordIdentifier:"), serviceIdentifier, objc.String(user), objc.String(recordIdentifier))
	return rv
}


// An indicator that enables you to prioritze credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/rank

func (p_ PasswordCredentialIdentity) Rank() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rank"))
	return rv
}


// An indicator that enables you to prioritze credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/rank

func (p_ PasswordCredentialIdentity) SetRank(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRank:"), value)
}


// A string used to correlate this identity to a record in your app’s own database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/recordIdentifier

func (p_ PasswordCredentialIdentity) RecordIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("recordIdentifier"))
	return rv
}


// An identifier that helps the system know with which apps or websites to associate this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/serviceIdentifier

func (p_ PasswordCredentialIdentity) ServiceIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("serviceIdentifier"))
	return rv
}


// The username associated with the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/user

func (p_ PasswordCredentialIdentity) User() string {
	rv := objc.Send[string](p_.ID, objc.Sel("user"))
	return rv
}


