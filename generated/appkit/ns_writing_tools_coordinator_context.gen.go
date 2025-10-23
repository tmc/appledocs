// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinatorContext] class.
var (
	WritingToolsCoordinatorContextClass     _WritingToolsCoordinatorContextClass
	WritingToolsCoordinatorContextClassOnce sync.Once
)

func getWritingToolsCoordinatorContextClass() _WritingToolsCoordinatorContextClass {
	WritingToolsCoordinatorContextClassOnce.Do(func() {
		WritingToolsCoordinatorContextClass = _WritingToolsCoordinatorContextClass{objc.GetClass("NSWritingToolsCoordinatorContext")}
	})
	return WritingToolsCoordinatorContextClass
}

type _WritingToolsCoordinatorContextClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinatorContext] class.
type IWritingToolsCoordinatorContext interface {
	objectivec.IObject
	Range() foundation.Range
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
	Identifier() foundation.UUID
	SetIdentifier(value foundation.IUUID)
	ResolvedRange() foundation.Range
	SetResolvedRange(value foundation.Range)
}

// A data object that you use to share your custom view’s text with Writing Tools.
//
// At the start of every Writing Tools operation, you create one or more objects with a copy of the text you want Writing Tools to evaluate. Each Writing Tools operation starts with a call to the method of your object. Use the parameters of that method to determine how much of your view’s text to provide. For some operations, Writing Tools asks for all of your view’s text, but in others it asks for only a portion of the text. When Writing Tools finishes its evaluation, it reports changes back to your delegate relative to the context objects you provided. When Writing Tools asks for your view’s text, create one or more objects with the requested content. If your view contains only one text storage object, create only one context object for the request. However, if you use multiple text storage objects to manage different parts of your view’s content, you might need to create multiple context objects. The actual number depends on how much of your text Writing Tools asks for. For example, when Writing Tools asks for all of your view’s content, you return one context object for each text storage object in your view. However, if Writing Tools asks for the current selection, and one text storage object contains all of the selected text, you create only one context object for the content. Writing Tools uses your context objects as the starting point for its evaluations, and as a reference point for any changes. Because Writing Tools doesn’t know anything about your view or its content, it makes suggestions only relative to your context objects. It’s your responsibility to take those suggestions and incorporate them back into your view’s text storage. In some cases, you might need to store additional information to update your storage correctly. For example, you might need to store, and update as needed, the offset from the start of your document to the start of the text in your context object. When Writing Tools asks for the currently selected text in your view, include some of the surrounding text in your context object as well. Supply a string that includes the selection and any text up to the nearest paragraph boundary. When creating your context object, specify a range value that represents the portion of that string that corresponds to the text selection. Providing some additional text in your context object can help Writing Tools improve its evaluation of your content. Writing Tools uses the property of your context object to indicate what text it considered. If your context object includes text that you don’t want Writing Tools to evaluate, add the attribute to the corresponding characters of your object. You might add this attribute if the text string includes a code listing or readonly content that you don’t want Writing Tools to change.


// A data object that you use to share your custom view’s text with Writing Tools.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context
type WritingToolsCoordinatorContext struct {
	objectivec.Object
}

// WritingToolsCoordinatorContextFrom constructs a [WritingToolsCoordinatorContext] from an unsafe.Pointer.
//
// A data object that you use to share your custom view’s text with Writing Tools.
func WritingToolsCoordinatorContextFrom(ptr unsafe.Pointer) WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorContextClass) Alloc() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WritingToolsCoordinatorContextClass) New() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorContext) Init() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WritingToolsCoordinatorContext) Autorelease() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWritingToolsCoordinatorContext creates a new WritingToolsCoordinatorContext instance.
func NewWritingToolsCoordinatorContext() WritingToolsCoordinatorContext {
	return getWritingToolsCoordinatorContextClass().New()
}



// The unique identifier of the context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/range
func (w_ WritingToolsCoordinatorContext) Range() foundation.Range {
	rv := objc.Send[foundation.Range](w_.ID, objc.Sel("range"))
	return rv
}


// The portion of your view’s text to evaluate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/attributedstring
func (w_ WritingToolsCoordinatorContext) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](w_.ID, objc.Sel("attributedString"))
	return rv
}


// The portion of your view’s text to evaluate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/attributedstring
func (w_ WritingToolsCoordinatorContext) SetAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAttributedString:"), value)
}


// The unique identifier of the context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/identifier
func (w_ WritingToolsCoordinatorContext) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](w_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier of the context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/identifier
func (w_ WritingToolsCoordinatorContext) SetIdentifier(value foundation.IUUID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIdentifier:"), value)
}


// The actual range of text that Writing Tools might change, which can
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/resolvedrange
func (w_ WritingToolsCoordinatorContext) ResolvedRange() foundation.Range {
	rv := objc.Send[foundation.Range](w_.ID, objc.Sel("resolvedRange"))
	return rv
}


// The actual range of text that Writing Tools might change, which can
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinator/context/resolvedrange
func (w_ WritingToolsCoordinatorContext) SetResolvedRange(value foundation.Range) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResolvedRange:"), value)
}




