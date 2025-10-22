// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCardSlotNFCSession] class.
var (
	TKSmartCardSlotNFCSessionClass     _TKSmartCardSlotNFCSessionClass
	TKSmartCardSlotNFCSessionClassOnce sync.Once
)

func getTKSmartCardSlotNFCSessionClass() _TKSmartCardSlotNFCSessionClass {
	TKSmartCardSlotNFCSessionClassOnce.Do(func() {
		TKSmartCardSlotNFCSessionClass = _TKSmartCardSlotNFCSessionClass{objc.GetClass("TKSmartCardSlotNFCSession")}
	})
	return TKSmartCardSlotNFCSessionClass
}

type _TKSmartCardSlotNFCSessionClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardSlotNFCSession] class.
type ITKSmartCardSlotNFCSession interface {
	objectivec.IObject
	EndSession()
	UpdateWithMessageError(message string, error_ unsafe.Pointer) bool
	SlotName() string
}

// NFC session that’s related to NFC smart card slot which was created.
//
// Lifetime of this session object is tied to the NFC smart card slot lifetime and once the NFC slot disappears (eg. after a user cancellation, calling end session, or an NFC timeout) the functions will start to fail and return error.


// NFC session that’s related to NFC smart card slot which was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession

type TKSmartCardSlotNFCSession struct {
	objectivec.Object
}

// TKSmartCardSlotNFCSessionFrom constructs a [TKSmartCardSlotNFCSession] from an unsafe.Pointer.
//
// NFC session that’s related to NFC smart card slot which was created.
func TKSmartCardSlotNFCSessionFrom(ptr unsafe.Pointer) TKSmartCardSlotNFCSession {
	return TKSmartCardSlotNFCSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotNFCSessionClass) Alloc() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardSlotNFCSessionClass) New() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlotNFCSession) Init() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlotNFCSession) Autorelease() TKSmartCardSlotNFCSession {
	rv := objc.Send[TKSmartCardSlotNFCSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlotNFCSession creates a new TKSmartCardSlotNFCSession instance.
func NewTKSmartCardSlotNFCSession() TKSmartCardSlotNFCSession {
	return getTKSmartCardSlotNFCSessionClass().New()
}




// Ends the NFC slot session and dismisses the system-presented NFC UI (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/end()

func (t_ TKSmartCardSlotNFCSession) EndSession() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endSession"))
}



// Updates the message of the system-presented NFC UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/update(message:)

func (t_ TKSmartCardSlotNFCSession) UpdateWithMessageError(message string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("updateWithMessage:error:"), objc.String(message), error_)
	return rv
}


// Smart card slot name of the NFC slot that was created together with this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/slotName

func (t_ TKSmartCardSlotNFCSession) SlotName() string {
	rv := objc.Send[string](t_.ID, objc.Sel("slotName"))
	return rv
}



