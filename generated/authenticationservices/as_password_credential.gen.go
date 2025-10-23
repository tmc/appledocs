// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasswordCredential] class.
var (
	PasswordCredentialClass     _PasswordCredentialClass
	PasswordCredentialClassOnce sync.Once
)

func getPasswordCredentialClass() _PasswordCredentialClass {
	PasswordCredentialClassOnce.Do(func() {
		PasswordCredentialClass = _PasswordCredentialClass{objc.GetClass("ASPasswordCredential")}
	})
	return PasswordCredentialClass
}

type _PasswordCredentialClass struct {
	class objc.Class
}

// An interface definition for the [PasswordCredential] class.
type IPasswordCredential interface {
	objectivec.IObject
	Password() string
	User() string
	SetUser(value string)
}

// A password credential.


// A password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential
type PasswordCredential struct {
	objectivec.Object
}

// PasswordCredentialFrom constructs a [PasswordCredential] from an unsafe.Pointer.
//
// A password credential.
func PasswordCredentialFrom(ptr unsafe.Pointer) PasswordCredential {
	return PasswordCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasswordCredentialClass) Alloc() PasswordCredential {
	rv := objc.Send[PasswordCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasswordCredentialClass) New() PasswordCredential {
	rv := objc.Send[PasswordCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasswordCredential) Init() PasswordCredential {
	rv := objc.Send[PasswordCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasswordCredential) Autorelease() PasswordCredential {
	rv := objc.Send[PasswordCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasswordCredential creates a new PasswordCredential instance.
func NewPasswordCredential() PasswordCredential {
	return getPasswordCredentialClass().New()
}



// Creates a password credential instance with a given user name and password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/credentialWithUser:password:
func (pc _PasswordCredentialClass) CredentialWithUserPassword(user string, password string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("credentialWithUser:password:"), objc.String(user), objc.String(password))
	return rv
}


// The password for a password credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/password
func (p_ PasswordCredential) Password() string {
	rv := objc.Send[string](p_.ID, objc.Sel("password"))
	return rv
}


// The user for a password credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasswordcredential/user
func (p_ PasswordCredential) User() string {
	rv := objc.Send[string](p_.ID, objc.Sel("user"))
	return rv
}


// The user for a password credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasswordcredential/user
func (p_ PasswordCredential) SetUser(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUser:"), objc.String(value))
}



