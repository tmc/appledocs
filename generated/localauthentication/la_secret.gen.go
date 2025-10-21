// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Secret] class.
var (
	SecretClass     _SecretClass
	SecretClassOnce sync.Once
)

func getSecretClass() _SecretClass {
	SecretClassOnce.Do(func() {
		SecretClass = _SecretClass{objc.GetClass("LASecret")}
	})
	return SecretClass
}

type _SecretClass struct {
	class objc.Class
}

// An interface definition for the [Secret] class.
type ISecret interface {
	objectivec.IObject
	LoadDataWithCompletion(handler unsafe.Pointer)
}

// Data that’s protected by a persisted right.
//
// You create instances when you store an ; you can’t create them directly.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LASecret
type Secret struct {
	objectivec.Object
}

// SecretFrom constructs a [Secret] from an unsafe.Pointer.
//
// Data that’s protected by a persisted right.
func SecretFrom(ptr unsafe.Pointer) Secret {
	return Secret{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SecretClass) Alloc() Secret {
	rv := objc.Send[Secret](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SecretClass) New() Secret {
	rv := objc.Send[Secret](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Secret) Init() Secret {
	rv := objc.Send[Secret](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Secret) Autorelease() Secret {
	rv := objc.Send[Secret](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecret creates a new Secret instance.
func NewSecret() Secret {
	return getSecretClass().New()
}


// Retrieves data stored in a secret.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LASecret/loadData(completion:)
func (s_ Secret) LoadDataWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadDataWithCompletion:"), handler)
}




