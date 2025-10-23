// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCardSlotManager] class.
var (
	TKSmartCardSlotManagerClass     _TKSmartCardSlotManagerClass
	TKSmartCardSlotManagerClassOnce sync.Once
)

func getTKSmartCardSlotManagerClass() _TKSmartCardSlotManagerClass {
	TKSmartCardSlotManagerClassOnce.Do(func() {
		TKSmartCardSlotManagerClass = _TKSmartCardSlotManagerClass{objc.GetClass("TKSmartCardSlotManager")}
	})
	return TKSmartCardSlotManagerClass
}

type _TKSmartCardSlotManagerClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardSlotManager] class.
type ITKSmartCardSlotManager interface {
	objectivec.IObject
	// properties:
	SlotNames() []string /* primitive/slice/pointer. */
	// methods:
	CreateNFCSlotWithMessageCompletion(message string /* primitive/slice/pointer. */, completion unsafe.Pointer)
	GetSlotWithNameReply(name string /* primitive/slice/pointer. */, reply unsafe.Pointer)
	SlotNamed(name string /* primitive/slice/pointer. */) ITKSmartCardSlot
}

// An interface to all available smart card reader slots.
//
// Get a list of all known smart card reader slots in the system using the property, and access individual slots by name using the method.


// An interface to all available smart card reader slots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager
type TKSmartCardSlotManager struct {
	objectivec.Object
}

// TKSmartCardSlotManagerFrom constructs a [TKSmartCardSlotManager] from an unsafe.Pointer.
//
// An interface to all available smart card reader slots.
func TKSmartCardSlotManagerFrom(ptr unsafe.Pointer) TKSmartCardSlotManager {
	return TKSmartCardSlotManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotManagerClass) Alloc() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardSlotManagerClass) New() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlotManager) Init() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlotManager) Autorelease() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlotManager creates a new TKSmartCardSlotManager instance.
func NewTKSmartCardSlotManager() TKSmartCardSlotManager {
	return getTKSmartCardSlotManagerClass().New()
}



// The shared singleton Smart Card reader slot manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/default
func (tc _TKSmartCardSlotManagerClass) DefaultManager() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("defaultManager"))
	return rv
}

// Creates an NFC smart card slot using the device’s hardware and presents a system UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/createNFCSlot(message:completion:)
func (t_ TKSmartCardSlotManager) CreateNFCSlotWithMessageCompletion(message string /* primitive/slice/pointer. */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("createNFCSlotWithMessage:completion:"), objc.String(message), completion)
}


// Asynchronously calls a block with a Smart Card reader slot for a specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/getSlot(withName:reply:)
func (t_ TKSmartCardSlotManager) GetSlotWithNameReply(name string /* primitive/slice/pointer. */, reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getSlotWithName:reply:"), objc.String(name), reply)
}


// Returns the Smart Card slot with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNamed(_:)
func (t_ TKSmartCardSlotManager) SlotNamed(name string /* primitive/slice/pointer. */) ITKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("slotNamed:"), objc.String(name))
	return rv
}


// The shared singleton Smart Card reader slot manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/default
func (t_ TKSmartCardSlotManager) DefaultManager() ITKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("defaultManager"))
	return rv
}


// A list of identifiers for all the Smart Card reader slots available to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNames
func (t_ TKSmartCardSlotManager) SlotNames() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](t_.ID, objc.Sel("slotNames"))
	return rv
}



