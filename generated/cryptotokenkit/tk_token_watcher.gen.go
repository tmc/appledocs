// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKTokenWatcher] class.
var (
	TKTokenWatcherClass     _TKTokenWatcherClass
	TKTokenWatcherClassOnce sync.Once
)

func getTKTokenWatcherClass() _TKTokenWatcherClass {
	TKTokenWatcherClassOnce.Do(func() {
		TKTokenWatcherClass = _TKTokenWatcherClass{objc.GetClass("TKTokenWatcher")}
	})
	return TKTokenWatcherClass
}

type _TKTokenWatcherClass struct {
	class objc.Class
}

// An interface definition for the [TKTokenWatcher] class.
type ITKTokenWatcher interface {
	objectivec.IObject
	// properties:
	TokenIDs() string /* primitive/slice/pointer. */
	SetTokenIDs(value string /* primitive/slice/pointer. */)
	// methods:
	SetInsertionHandler(insertionHandler unsafe.Pointer)
	TokenInfoForTokenID(tokenID string /* primitive/slice/pointer. */) ITKTokenWatcherTokenInfo
}

// An object that tracks the tokens available in the system.
//
// Create a token watcher and register an insertion handler to be notified when tokens are added to the system. You can also add removal handlers for specific tokens to be notified when those tokens are removed from the system.


// An object that tracks the tokens available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher
type TKTokenWatcher struct {
	objectivec.Object
}

// TKTokenWatcherFrom constructs a [TKTokenWatcher] from an unsafe.Pointer.
//
// An object that tracks the tokens available in the system.
func TKTokenWatcherFrom(ptr unsafe.Pointer) TKTokenWatcher {
	return TKTokenWatcher{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenWatcherClass) Alloc() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenWatcherClass) New() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenWatcher) Init() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenWatcher) Autorelease() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenWatcher creates a new TKTokenWatcher instance.
func NewTKTokenWatcher() TKTokenWatcher {
	return getTKTokenWatcherClass().New()
}




// Sets an insertion handler closure to be called when a new token is inserted into the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/setInsertionHandler(_:)
func (t_ TKTokenWatcher) SetInsertionHandler(insertionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionHandler:"), insertionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/tokenInfo(forTokenID:)
func (t_ TKTokenWatcher) TokenInfoForTokenID(tokenID string /* primitive/slice/pointer. */) ITKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("tokenInfoForTokenID:"), objc.String(tokenID))
	return rv
}


// The token IDs currently available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenwatcher/tokenids
func (t_ TKTokenWatcher) TokenIDs() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](t_.ID, objc.Sel("tokenIDs"))
	return rv
}


// The token IDs currently available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenwatcher/tokenids
func (t_ TKTokenWatcher) SetTokenIDs(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenIDs:"), objc.String(value))
}


