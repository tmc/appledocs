// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKShareRequestAccessOperation] class.
var (
	CKShareRequestAccessOperationClass     _CKShareRequestAccessOperationClass
	CKShareRequestAccessOperationClassOnce sync.Once
)

func getCKShareRequestAccessOperationClass() _CKShareRequestAccessOperationClass {
	CKShareRequestAccessOperationClassOnce.Do(func() {
		CKShareRequestAccessOperationClass = _CKShareRequestAccessOperationClass{objc.GetClass("CKShareRequestAccessOperation")}
	})
	return CKShareRequestAccessOperationClass
}

type _CKShareRequestAccessOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKShareRequestAccessOperation] class.
type ICKShareRequestAccessOperation interface {
	ICKOperation
	PerShareAccessRequestCompletionBlock() unsafe.Pointer
	SetPerShareAccessRequestCompletionBlock(value unsafe.Pointer)
	ShareRequestAccessCompletionBlock() func(error objc.ID)
	SetShareRequestAccessCompletionBlock(value func(error objc.ID))
	ShareURLs() []foundation.URL
	SetShareURLs(value []foundation.URL)
	PerShareAccessRequestResultBlock() unsafe.Pointer
	SetPerShareAccessRequestResultBlock(value unsafe.Pointer)
	ShareAccessRequestResultBlock() unsafe.Pointer
	SetShareAccessRequestResultBlock(value unsafe.Pointer)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation
type CKShareRequestAccessOperation struct {
	CKOperation
}

// CKShareRequestAccessOperationFrom constructs a [CKShareRequestAccessOperation] from an unsafe.Pointer.
func CKShareRequestAccessOperationFrom(ptr unsafe.Pointer) CKShareRequestAccessOperation {
	return CKShareRequestAccessOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKShareRequestAccessOperationClass) Alloc() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKShareRequestAccessOperationClass) New() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareRequestAccessOperation) Init() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareRequestAccessOperation) Autorelease() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareRequestAccessOperation creates a new CKShareRequestAccessOperation instance.
func NewCKShareRequestAccessOperation() CKShareRequestAccessOperation {
	return getCKShareRequestAccessOperationClass().New()
}



// Creates a share request access operation configured with specified share URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/init(shareURLs:)
func NewCKShareRequestAccessOperationWithShareURLs(shareURLs []foundation.URL) CKShareRequestAccessOperation {
	instance := getCKShareRequestAccessOperationClass().Alloc()
	rv := objc.Send[CKShareRequestAccessOperation](instance.ID, objc.Sel("initWithShareURLs:"), shareURLs)
	rv.Autorelease()
	return rv
}



// A completion block called once for each processed share URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/perShareAccessRequestCompletionBlock
func (c_ CKShareRequestAccessOperation) PerShareAccessRequestCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareAccessRequestCompletionBlock"))
	return rv
}


// A completion block called once for each processed share URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/perShareAccessRequestCompletionBlock
func (c_ CKShareRequestAccessOperation) SetPerShareAccessRequestCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareAccessRequestCompletionBlock:"), value)
}


// A completion block called when the entire operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareRequestAccessCompletionBlock
func (c_ CKShareRequestAccessOperation) ShareRequestAccessCompletionBlock() func(error objc.ID) {
	rv := objc.Send[func(error objc.ID)](c_.ID, objc.Sel("shareRequestAccessCompletionBlock"))
	return rv
}


// A completion block called when the entire operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareRequestAccessCompletionBlock
func (c_ CKShareRequestAccessOperation) SetShareRequestAccessCompletionBlock(value func(error objc.ID)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareRequestAccessCompletionBlock:"), value)
}


// The URLs of the shares to request access to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareURLs
func (c_ CKShareRequestAccessOperation) ShareURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](c_.ID, objc.Sel("shareURLs"))
	return rv
}


// The URLs of the shares to request access to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareURLs
func (c_ CKShareRequestAccessOperation) SetShareURLs(value []foundation.URL) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareURLs:"), nsArray)
}


// A block called once for each share URL processed by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/pershareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) PerShareAccessRequestResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareAccessRequestResultBlock"))
	return rv
}


// A block called once for each share URL processed by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/pershareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) SetPerShareAccessRequestResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareAccessRequestResultBlock:"), value)
}


// A block called when the entire share access request operation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/shareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) ShareAccessRequestResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("shareAccessRequestResultBlock"))
	return rv
}


// A block called when the entire share access request operation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/shareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) SetShareAccessRequestResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareAccessRequestResultBlock:"), value)
}


