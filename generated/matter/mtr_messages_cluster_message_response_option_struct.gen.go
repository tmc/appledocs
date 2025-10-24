// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterMessageResponseOptionStruct] class.
var (
	MTRMessagesClusterMessageResponseOptionStructClass     _MTRMessagesClusterMessageResponseOptionStructClass
	MTRMessagesClusterMessageResponseOptionStructClassOnce sync.Once
)

func getMTRMessagesClusterMessageResponseOptionStructClass() _MTRMessagesClusterMessageResponseOptionStructClass {
	MTRMessagesClusterMessageResponseOptionStructClassOnce.Do(func() {
		MTRMessagesClusterMessageResponseOptionStructClass = _MTRMessagesClusterMessageResponseOptionStructClass{objc.GetClass("MTRMessagesClusterMessageResponseOptionStruct")}
	})
	return MTRMessagesClusterMessageResponseOptionStructClass
}

type _MTRMessagesClusterMessageResponseOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterMessageResponseOptionStruct] class.
type IMTRMessagesClusterMessageResponseOptionStruct interface {
	objectivec.IObject
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MessageResponseID() objc.IObject /* cross-framework: NSNumber */
	SetMessageResponseID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct
type MTRMessagesClusterMessageResponseOptionStruct struct {
	objectivec.Object
}

// MTRMessagesClusterMessageResponseOptionStructFrom constructs a [MTRMessagesClusterMessageResponseOptionStruct] from an unsafe.Pointer.
func MTRMessagesClusterMessageResponseOptionStructFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageResponseOptionStruct {
	return MTRMessagesClusterMessageResponseOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageResponseOptionStructClass) Alloc() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterMessageResponseOptionStructClass) New() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Init() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Autorelease() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageResponseOptionStruct creates a new MTRMessagesClusterMessageResponseOptionStruct instance.
func NewMTRMessagesClusterMessageResponseOptionStruct() MTRMessagesClusterMessageResponseOptionStruct {
	return getMTRMessagesClusterMessageResponseOptionStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/label
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/label
func (m_ MTRMessagesClusterMessageResponseOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/messageResponseID
func (m_ MTRMessagesClusterMessageResponseOptionStruct) MessageResponseID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageResponseID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/messageResponseID
func (m_ MTRMessagesClusterMessageResponseOptionStruct) SetMessageResponseID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageResponseID:"), value)
}



