// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariExtensionManager] class.
var (
	SFSafariExtensionManagerClass     _SFSafariExtensionManagerClass
	SFSafariExtensionManagerClassOnce sync.Once
)

func getSFSafariExtensionManagerClass() _SFSafariExtensionManagerClass {
	SFSafariExtensionManagerClassOnce.Do(func() {
		SFSafariExtensionManagerClass = _SFSafariExtensionManagerClass{objc.GetClass("SFSafariExtensionManager")}
	})
	return SFSafariExtensionManagerClass
}

type _SFSafariExtensionManagerClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariExtensionManager] class.
type ISFSafariExtensionManager interface {
	objectivec.IObject
	SFExtensionProfileKey() string
}

// A class that your app uses to find out the current state of a Safari app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionManager
type SFSafariExtensionManager struct {
	objectivec.Object
}

// SFSafariExtensionManagerFrom constructs a [SFSafariExtensionManager] from an unsafe.Pointer.
//
// A class that your app uses to find out the current state of a Safari app extension.
func SFSafariExtensionManagerFrom(ptr unsafe.Pointer) SFSafariExtensionManager {
	return SFSafariExtensionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionManagerClass) Alloc() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariExtensionManagerClass) New() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionManager) Init() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionManager) Autorelease() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionManager creates a new SFSafariExtensionManager instance.
func NewSFSafariExtensionManager() SFSafariExtensionManager {
	return getSFSafariExtensionManagerClass().New()
}


// Gets the current state of the Safari app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionManager/getStateOfSafariExtension(withIdentifier:completionHandler:)
func (sc _SFSafariExtensionManagerClass) GetStateOfSafariExtensionWithIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStateOfSafariExtensionWithIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}

// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariExtensionManager) SFExtensionProfileKey() string {
	rv := objc.Send[string](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}



