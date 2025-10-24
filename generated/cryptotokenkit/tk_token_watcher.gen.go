// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenWatcher */


/* debug [class_header]: Header for TKTokenWatcher */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenWatcher */
// An interface definition for the [TKTokenWatcher] class.
type ITKTokenWatcher interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenWatcher */
	// properties:
	TokenIDs() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenWatcher */
	// methods:
	AddRemovalHandlerForTokenID(removalHandler unsafe.Pointer, tokenID objc.IObject /* cross-framework: NSString */)
	SetInsertionHandler(insertionHandler unsafe.Pointer)
	TokenInfoForTokenID(tokenID objc.IObject /* cross-framework: NSString */) ITKTokenWatcherTokenInfo
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenWatcher */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenWatcherClass) Alloc() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenWatcher */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenWatcher */

// Initializes a token watcher with the specified insertion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/init(insertionHandler:)
func NewTKTokenWatcherWithInsertionHandler(insertionHandler unsafe.Pointer) TKTokenWatcher {
	instance := getTKTokenWatcherClass().Alloc()
	rv := objc.Send[TKTokenWatcher](instance.ID, objc.Sel("initWithInsertionHandler:"), insertionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenWatcherWithInsertionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenWatcher */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenWatcher */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenWatcher */

// Adds a removal handler for the specified token ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/addRemovalHandler(_:forTokenID:)
func (t_ TKTokenWatcher) AddRemovalHandlerForTokenID(removalHandler unsafe.Pointer, tokenID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addRemovalHandler:forTokenID:"), removalHandler, tokenID)
}/* debug [instance_methods/method]: AddRemovalHandlerForTokenID */


// Sets an insertion handler closure to be called when a new token is inserted into the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/setInsertionHandler(_:)
func (t_ TKTokenWatcher) SetInsertionHandler(insertionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionHandler:"), insertionHandler)
}/* debug [instance_methods/method]: SetInsertionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/tokenInfo(forTokenID:)
func (t_ TKTokenWatcher) TokenInfoForTokenID(tokenID objc.IObject /* cross-framework: NSString */) ITKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("tokenInfoForTokenID:"), tokenID)
	return rv
}/* debug [instance_methods/method]: TokenInfoForTokenID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenWatcher */

// The token IDs currently available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/tokenIDs
func (t_ TKTokenWatcher) TokenIDs() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("tokenIDs"))
	return rv
}/* debug [instance_properties/getter]: tokenIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenWatcher */


