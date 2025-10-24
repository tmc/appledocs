// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPersistentChangeFetchResult] class.
var (
	PHPersistentChangeFetchResultClass     _PHPersistentChangeFetchResultClass
	PHPersistentChangeFetchResultClassOnce sync.Once
)

func getPHPersistentChangeFetchResultClass() _PHPersistentChangeFetchResultClass {
	PHPersistentChangeFetchResultClassOnce.Do(func() {
		PHPersistentChangeFetchResultClass = _PHPersistentChangeFetchResultClass{objc.GetClass("PHPersistentChangeFetchResult")}
	})
	return PHPersistentChangeFetchResultClass
}

type _PHPersistentChangeFetchResultClass struct {
	class objc.Class
}

// An interface definition for the [PHPersistentChangeFetchResult] class.
type IPHPersistentChangeFetchResult interface {
	objectivec.IObject
	// properties:
	CurrentChangeToken() IPHPersistentChangeToken
	SetCurrentChangeToken(value IPHPersistentChangeToken)
	// methods:
}

// An object that represents a fetch result and allows you to enumerate a very large set of change records.

// An object that represents a fetch result and allows you to enumerate a very large set of change records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPersistentChangeFetchResult
type PHPersistentChangeFetchResult struct {
	objectivec.Object
}

// PHPersistentChangeFetchResultFrom constructs a [PHPersistentChangeFetchResult] from an unsafe.Pointer.
//
// An object that represents a fetch result and allows you to enumerate a very large set of change records.
func PHPersistentChangeFetchResultFrom(ptr unsafe.Pointer) PHPersistentChangeFetchResult {
	return PHPersistentChangeFetchResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPersistentChangeFetchResultClass) Alloc() PHPersistentChangeFetchResult {
	rv := objc.Send[PHPersistentChangeFetchResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPersistentChangeFetchResultClass) New() PHPersistentChangeFetchResult {
	rv := objc.Send[PHPersistentChangeFetchResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPersistentChangeFetchResult) Init() PHPersistentChangeFetchResult {
	rv := objc.Send[PHPersistentChangeFetchResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPersistentChangeFetchResult) Autorelease() PHPersistentChangeFetchResult {
	rv := objc.Send[PHPersistentChangeFetchResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPersistentChangeFetchResult creates a new PHPersistentChangeFetchResult instance.
func NewPHPersistentChangeFetchResult() PHPersistentChangeFetchResult {
	return getPHPersistentChangeFetchResultClass().New()
}

// The opaque token that represents the current state of the Photos library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phphotolibrary/currentchangetoken
func (p_ PHPersistentChangeFetchResult) CurrentChangeToken() IPHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](p_.ID, objc.Sel("currentChangeToken"))
	return rv
}

// The opaque token that represents the current state of the Photos library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phphotolibrary/currentchangetoken
func (p_ PHPersistentChangeFetchResult) SetCurrentChangeToken(value IPHPersistentChangeToken) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentChangeToken:"), value)
}
