// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TextStyleRule] class.
var (
	TextStyleRuleClass     _TextStyleRuleClass
	TextStyleRuleClassOnce sync.Once
)

func getTextStyleRuleClass() _TextStyleRuleClass {
	TextStyleRuleClassOnce.Do(func() {
		TextStyleRuleClass = _TextStyleRuleClass{objc.GetClass("AVTextStyleRule")}
	})
	return TextStyleRuleClass
}

type _TextStyleRuleClass struct {
	class objc.Class
}





// An interface definition for the [TextStyleRule] class.
type ITextStyleRule interface {
	objectivec.IObject
	

	// properties:
	TextMarkupAttributes() foundation.IDictionary
	TextSelector() objc.IObject /* cross-framework: NSString */
	TextStyleRules() IAVTextStyleRule
	SetTextStyleRules(value IAVTextStyleRule)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TextStyleRuleClass) Alloc() TextStyleRule {
	rv := objc.Send[TextStyleRule](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextStyleRuleClass) New() TextStyleRule {
	rv := objc.Send[TextStyleRule](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextStyleRule) Init() TextStyleRule {
	rv := objc.Send[TextStyleRule](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextStyleRule) Autorelease() TextStyleRule {
	rv := objc.Send[TextStyleRule](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextStyleRule creates a new TextStyleRule instance.
func NewTextStyleRule() TextStyleRule {
	return getTextStyleRuleClass().New()
}





// An object that represents the text styling rules to apply to a media item’s textual content.
//
// You use text style objects to format subtitles, closed captions, and other text-related content of the item. The system applies these rules to all or part of the text of the media item.


// An object that represents the text styling rules to apply to a media item’s textual content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule
type TextStyleRule struct {
	objectivec.Object
}

// TextStyleRuleFrom constructs a [TextStyleRule] from an unsafe.Pointer.
//
// An object that represents the text styling rules to apply to a media item’s textual content.
func TextStyleRuleFrom(ptr unsafe.Pointer) TextStyleRule {
	return TextStyleRule{objectivec.Object{objc.ID(ptr)}}
}






// Creates a text style rule object with the specified style attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:)
func NewTextStyleRuleWithTextMarkupAttributes(textMarkupAttributes foundation.IDictionary) TextStyleRule {
	instance := getTextStyleRuleClass().Alloc()
	rv := objc.Send[TextStyleRule](instance.ID, objc.Sel("initWithTextMarkupAttributes:"), textMarkupAttributes)
	rv.Autorelease()
	return rv
}


// Creates a text style rule object with the specified style attributes and text range information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:textSelector:)
func NewTextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.IDictionary, textSelector objc.IObject /* cross-framework: NSString */) TextStyleRule {
	instance := getTextStyleRuleClass().Alloc()
	rv := objc.Send[TextStyleRule](instance.ID, objc.Sel("initWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, textSelector)
	rv.Autorelease()
	return rv
}







// Converts one or more text style rules into a serializable property list object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/propertyList(for:)
func (tc _TextStyleRuleClass) PropertyListForTextStyleRules(textStyleRules []TextStyleRule) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("propertyListForTextStyleRules:"), textStyleRules)
	return rv
}


// Creates a new text style rule object using the style attributes in the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRuleWithTextMarkupAttributes:
func (tc _TextStyleRuleClass) TextStyleRuleWithTextMarkupAttributes(textMarkupAttributes foundation.IDictionary) ITextStyleRule {
	rv := objc.Send[TextStyleRule](objc.ID(tc.class), objc.Sel("textStyleRuleWithTextMarkupAttributes:"), textMarkupAttributes)
	return rv
}


// Creates a new text style rule object using the specified style attributes and text range information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRuleWithTextMarkupAttributes:textSelector:
func (tc _TextStyleRuleClass) TextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.IDictionary, textSelector objc.IObject /* cross-framework: NSString */) ITextStyleRule {
	rv := objc.Send[TextStyleRule](objc.ID(tc.class), objc.Sel("textStyleRuleWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, textSelector)
	return rv
}


// Creates an array of text style rule objects from the specified property-list object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRules(fromPropertyList:)
func (tc _TextStyleRuleClass) TextStyleRulesFromPropertyList(plist objc.IObject) []TextStyleRule {
	rv := objc.Send[[]TextStyleRule](objc.ID(tc.class), objc.Sel("textStyleRulesFromPropertyList:"), plist)
	return rv
}

















// A dictionary of text style attributes to apply to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textMarkupAttributes
func (t_ TextStyleRule) TextMarkupAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("textMarkupAttributes"))
	return rv
}


// A string that identifies the text to which the attributes should apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textSelector
func (t_ TextStyleRule) TextSelector() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("textSelector"))
	return rv
}


// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/textstylerules
func (t_ TextStyleRule) TextStyleRules() IAVTextStyleRule {
	rv := objc.Send[TextStyleRule](t_.ID, objc.Sel("textStyleRules"))
	return rv
}


// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/textstylerules
func (t_ TextStyleRule) SetTextStyleRules(value IAVTextStyleRule) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStyleRules:"), value)
}







