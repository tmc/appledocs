// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterMessageStruct] class.
var (
	MTRMessagesClusterMessageStructClass     _MTRMessagesClusterMessageStructClass
	MTRMessagesClusterMessageStructClassOnce sync.Once
)

func getMTRMessagesClusterMessageStructClass() _MTRMessagesClusterMessageStructClass {
	MTRMessagesClusterMessageStructClassOnce.Do(func() {
		MTRMessagesClusterMessageStructClass = _MTRMessagesClusterMessageStructClass{objc.GetClass("MTRMessagesClusterMessageStruct")}
	})
	return MTRMessagesClusterMessageStructClass
}

type _MTRMessagesClusterMessageStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterMessageStruct] class.
type IMTRMessagesClusterMessageStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct
type MTRMessagesClusterMessageStruct struct {
	objectivec.Object
}

// MTRMessagesClusterMessageStructFrom constructs a [MTRMessagesClusterMessageStruct] from an unsafe.Pointer.
func MTRMessagesClusterMessageStructFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageStruct {
	return MTRMessagesClusterMessageStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageStructClass) Alloc() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterMessageStructClass) New() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageStruct) Init() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageStruct) Autorelease() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageStruct creates a new MTRMessagesClusterMessageStruct instance.
func NewMTRMessagesClusterMessageStruct() MTRMessagesClusterMessageStruct {
	return getMTRMessagesClusterMessageStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageControl
func (m_ MTRMessagesClusterMessageStruct) MessageControl() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("messageControl"))
	return rv
}


// SetMessageControl sets the value of the messageControl property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageControl
func (m_ MTRMessagesClusterMessageStruct) SetMessageControl(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageID
func (m_ MTRMessagesClusterMessageStruct) MessageID() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// SetMessageID sets the value of the messageID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageID
func (m_ MTRMessagesClusterMessageStruct) SetMessageID(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageText
func (m_ MTRMessagesClusterMessageStruct) MessageText() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("messageText"))
	return rv
}


// SetMessageText sets the value of the messageText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageText
func (m_ MTRMessagesClusterMessageStruct) SetMessageText(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/priority
func (m_ MTRMessagesClusterMessageStruct) Priority() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("priority"))
	return rv
}


// SetPriority sets the value of the priority property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/priority
func (m_ MTRMessagesClusterMessageStruct) SetPriority(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/responses
func (m_ MTRMessagesClusterMessageStruct) Responses() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("responses"))
	return rv
}


// SetResponses sets the value of the responses property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/responses
func (m_ MTRMessagesClusterMessageStruct) SetResponses(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/startTime
func (m_ MTRMessagesClusterMessageStruct) StartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/startTime
func (m_ MTRMessagesClusterMessageStruct) SetStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}



