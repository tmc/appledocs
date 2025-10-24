// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	MessageControl() objc.IObject /* cross-framework: NSNumber */
	SetMessageControl(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() objc.IObject /* cross-framework: NSData */
	SetMessageID(value objc.IObject /* cross-framework: NSData */)
	MessageText() objc.IObject /* cross-framework: NSString */
	SetMessageText(value objc.IObject /* cross-framework: NSString */)
	Priority() objc.IObject /* cross-framework: NSNumber */
	SetPriority(value objc.IObject /* cross-framework: NSNumber */)
	Responses() objc.IObject /* cross-framework: NSArray */
	SetResponses(value objc.IObject /* cross-framework: NSArray */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageControl
func (m_ MTRMessagesClusterMessageStruct) MessageControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageControl"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageControl
func (m_ MTRMessagesClusterMessageStruct) SetMessageControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageID
func (m_ MTRMessagesClusterMessageStruct) MessageID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageID
func (m_ MTRMessagesClusterMessageStruct) SetMessageID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageText
func (m_ MTRMessagesClusterMessageStruct) MessageText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("messageText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/messageText
func (m_ MTRMessagesClusterMessageStruct) SetMessageText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/priority
func (m_ MTRMessagesClusterMessageStruct) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/priority
func (m_ MTRMessagesClusterMessageStruct) SetPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/responses
func (m_ MTRMessagesClusterMessageStruct) Responses() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("responses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/responses
func (m_ MTRMessagesClusterMessageStruct) SetResponses(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/startTime
func (m_ MTRMessagesClusterMessageStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/startTime
func (m_ MTRMessagesClusterMessageStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}



