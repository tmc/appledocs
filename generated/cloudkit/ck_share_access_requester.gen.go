// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKShareAccessRequester] class.
var (
	CKShareAccessRequesterClass     _CKShareAccessRequesterClass
	CKShareAccessRequesterClassOnce sync.Once
)

func getCKShareAccessRequesterClass() _CKShareAccessRequesterClass {
	CKShareAccessRequesterClassOnce.Do(func() {
		CKShareAccessRequesterClass = _CKShareAccessRequesterClass{objc.GetClass("CKShareAccessRequester")}
	})
	return CKShareAccessRequesterClass
}

type _CKShareAccessRequesterClass struct {
	class objc.Class
}

// An interface definition for the [CKShareAccessRequester] class.
type ICKShareAccessRequester interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester
type CKShareAccessRequester struct {
	objectivec.Object
}

// CKShareAccessRequesterFrom constructs a [CKShareAccessRequester] from an unsafe.Pointer.
func CKShareAccessRequesterFrom(ptr unsafe.Pointer) CKShareAccessRequester {
	return CKShareAccessRequester{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKShareAccessRequesterClass) Alloc() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKShareAccessRequesterClass) New() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareAccessRequester) Init() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareAccessRequester) Autorelease() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareAccessRequester creates a new CKShareAccessRequester instance.
func NewCKShareAccessRequester() CKShareAccessRequester {
	return getCKShareAccessRequesterClass().New()
}


// A displayable representing the requester.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/contact
func (c_ CKShareAccessRequester) Contact() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contact"))
	return rv
}

// Lookup information for the requester.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/participantLookupInfo
func (c_ CKShareAccessRequester) ParticipantLookupInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("participantLookupInfo"))
	return rv
}

// The identity of the user requesting access to the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/userIdentity
func (c_ CKShareAccessRequester) UserIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userIdentity"))
	return rv
}



