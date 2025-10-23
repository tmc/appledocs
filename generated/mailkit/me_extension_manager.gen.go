// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEExtensionManager] class.
var (
	MEExtensionManagerClass     _MEExtensionManagerClass
	MEExtensionManagerClassOnce sync.Once
)

func getMEExtensionManagerClass() _MEExtensionManagerClass {
	MEExtensionManagerClassOnce.Do(func() {
		MEExtensionManagerClass = _MEExtensionManagerClass{objc.GetClass("MEExtensionManager")}
	})
	return MEExtensionManagerClass
}

type _MEExtensionManagerClass struct {
	class objc.Class
}

// An interface definition for the [MEExtensionManager] class.
type IMEExtensionManager interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionManager
type MEExtensionManager struct {
	objectivec.Object
}

// MEExtensionManagerFrom constructs a [MEExtensionManager] from an unsafe.Pointer.
func MEExtensionManagerFrom(ptr unsafe.Pointer) MEExtensionManager {
	return MEExtensionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEExtensionManagerClass) Alloc() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEExtensionManagerClass) New() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEExtensionManager) Init() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEExtensionManager) Autorelease() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEExtensionManager creates a new MEExtensionManager instance.
func NewMEExtensionManager() MEExtensionManager {
	return getMEExtensionManagerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionManager/reloadContentBlocker(withIdentifier:completionHandler:)
func (mc _MEExtensionManagerClass) ReloadContentBlockerWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("reloadContentBlockerWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}



