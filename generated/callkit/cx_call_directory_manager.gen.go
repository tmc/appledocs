// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallDirectoryManager] class.
var (
	CXCallDirectoryManagerClass     _CXCallDirectoryManagerClass
	CXCallDirectoryManagerClassOnce sync.Once
)

func getCXCallDirectoryManagerClass() _CXCallDirectoryManagerClass {
	CXCallDirectoryManagerClassOnce.Do(func() {
		CXCallDirectoryManagerClass = _CXCallDirectoryManagerClass{objc.GetClass("CXCallDirectoryManager")}
	})
	return CXCallDirectoryManagerClass
}

type _CXCallDirectoryManagerClass struct {
	class objc.Class
}

// An interface definition for the [CXCallDirectoryManager] class.
type ICXCallDirectoryManager interface {
	objectivec.IObject
	GetEnabledStatusForExtensionWithIdentifierCompletionHandler(identifier string, completion unsafe.Pointer)
	OpenSettingsWithCompletionHandler(completion unsafe.Pointer)
	ReloadExtensionWithIdentifierCompletionHandler(identifier string, completion unsafe.Pointer)
}

// The programmatic interface to an object that manages a Call Directory app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager
type CXCallDirectoryManager struct {
	objectivec.Object
}

// CXCallDirectoryManagerFrom constructs a [CXCallDirectoryManager] from an unsafe.Pointer.
//
// The programmatic interface to an object that manages a Call Directory app extension.
func CXCallDirectoryManagerFrom(ptr unsafe.Pointer) CXCallDirectoryManager {
	return CXCallDirectoryManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryManagerClass) Alloc() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallDirectoryManagerClass) New() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryManager) Init() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryManager) Autorelease() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryManager creates a new CXCallDirectoryManager instance.
func NewCXCallDirectoryManager() CXCallDirectoryManager {
	return getCXCallDirectoryManagerClass().New()
}


// Returns the shared call directory manager instance for the app.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/sharedInstance
func (cc _CXCallDirectoryManagerClass) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sharedInstance"))
	return rv
}
// Asynchronously returns the enabled status of the extension with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/getEnabledStatusForExtension(withIdentifier:completionHandler:)
func (c_ CXCallDirectoryManager) GetEnabledStatusForExtensionWithIdentifierCompletionHandler(identifier string, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEnabledStatusForExtensionWithIdentifier:completionHandler:"), objc.String(identifier), completion)
}

// Opens the iOS Settings app and shows the Call Blocking & Identification settings.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/openSettings(completionHandler:)
func (c_ CXCallDirectoryManager) OpenSettingsWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("openSettingsWithCompletionHandler:"), completion)
}

// Asynchronously reloads the extension with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/reloadExtension(withIdentifier:completionHandler:)
func (c_ CXCallDirectoryManager) ReloadExtensionWithIdentifierCompletionHandler(identifier string, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadExtensionWithIdentifier:completionHandler:"), objc.String(identifier), completion)
}

// Returns the shared call directory manager instance for the app.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/sharedInstance
func (c_ CXCallDirectoryManager) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sharedInstance"))
	return rv
}



