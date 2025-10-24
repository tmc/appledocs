// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMCSSStyleDeclaration */


/* debug [class_header]: Header for DOMCSSStyleDeclaration */
// The class instance for the [DOMCSSStyleDeclaration] class.
var (
	DOMCSSStyleDeclarationClass     _DOMCSSStyleDeclarationClass
	DOMCSSStyleDeclarationClassOnce sync.Once
)

func getDOMCSSStyleDeclarationClass() _DOMCSSStyleDeclarationClass {
	DOMCSSStyleDeclarationClassOnce.Do(func() {
		DOMCSSStyleDeclarationClass = _DOMCSSStyleDeclarationClass{objc.GetClass("DOMCSSStyleDeclaration")}
	})
	return DOMCSSStyleDeclarationClass
}

type _DOMCSSStyleDeclarationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMCSSStyleDeclaration */
// An interface definition for the [DOMCSSStyleDeclaration] class.
type IDOMCSSStyleDeclaration interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMCSSStyleDeclaration */
	// properties:
	CssText() objc.IObject /* cross-framework: NSString */
	SetCssText(value objc.IObject /* cross-framework: NSString */)
	Length() objectivec.IObject
	ParentRule() IDOMCSSRule
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMCSSStyleDeclaration */
	// methods:
	Azimuth() foundation.String
	Background() foundation.String
	BackgroundAttachment() foundation.String
	BackgroundColor() foundation.String
	BackgroundImage() foundation.String
	BackgroundPosition() foundation.String
	BackgroundRepeat() foundation.String
	Border() foundation.String
	BorderBottom() foundation.String
	BorderBottomColor() foundation.String
	BorderBottomStyle() foundation.String
	BorderBottomWidth() foundation.String
	BorderCollapse() foundation.String
	BorderColor() foundation.String
	BorderLeft() foundation.String
	BorderLeftColor() foundation.String
	BorderLeftStyle() foundation.String
	BorderLeftWidth() foundation.String
	BorderRight() foundation.String
	BorderRightColor() foundation.String
	BorderRightStyle() foundation.String
	BorderRightWidth() foundation.String
	BorderSpacing() foundation.String
	BorderStyle() foundation.String
	BorderTop() foundation.String
	BorderTopColor() foundation.String
	BorderTopStyle() foundation.String
	BorderTopWidth() foundation.String
	BorderWidth() foundation.String
	Bottom() foundation.String
	CaptionSide() foundation.String
	Clear() foundation.String
	Clip() foundation.String
	Color() foundation.String
	Content() foundation.String
	CounterIncrement() foundation.String
	CounterReset() foundation.String
	CssFloat() foundation.String
	Cue() foundation.String
	CueAfter() foundation.String
	CueBefore() foundation.String
	Cursor() foundation.String
	Direction() foundation.String
	Display() foundation.String
	Elevation() foundation.String
	EmptyCells() foundation.String
	Font() foundation.String
	FontFamily() foundation.String
	FontSize() foundation.String
	FontSizeAdjust() foundation.String
	FontStretch() foundation.String
	FontStyle() foundation.String
	FontVariant() foundation.String
	FontWeight() foundation.String
	Height() foundation.String
	Left() foundation.String
	LetterSpacing() foundation.String
	LineHeight() foundation.String
	ListStyle() foundation.String
	ListStyleImage() foundation.String
	ListStylePosition() foundation.String
	ListStyleType() foundation.String
	Margin() foundation.String
	MarginBottom() foundation.String
	MarginLeft() foundation.String
	MarginRight() foundation.String
	MarginTop() foundation.String
	MarkerOffset() foundation.String
	Marks() foundation.String
	MaxHeight() foundation.String
	MaxWidth() foundation.String
	MinHeight() foundation.String
	MinWidth() foundation.String
	Orphans() foundation.String
	Outline() foundation.String
	OutlineColor() foundation.String
	OutlineStyle() foundation.String
	OutlineWidth() foundation.String
	Overflow() foundation.String
	Padding() foundation.String
	PaddingBottom() foundation.String
	PaddingLeft() foundation.String
	PaddingRight() foundation.String
	PaddingTop() foundation.String
	Page() foundation.String
	PageBreakAfter() foundation.String
	PageBreakBefore() foundation.String
	PageBreakInside() foundation.String
	Pause() foundation.String
	PauseAfter() foundation.String
	PauseBefore() foundation.String
	Pitch() foundation.String
	PitchRange() foundation.String
	PlayDuring() foundation.String
	Position() foundation.String
	Quotes() foundation.String
	Richness() foundation.String
	Right() foundation.String
	SetAzimuth(azimuth objc.IObject /* cross-framework: NSString */)
	SetBackground(background objc.IObject /* cross-framework: NSString */)
	SetBackgroundAttachment(backgroundAttachment objc.IObject /* cross-framework: NSString */)
	SetBackgroundColor(backgroundColor objc.IObject /* cross-framework: NSString */)
	SetBackgroundImage(backgroundImage objc.IObject /* cross-framework: NSString */)
	SetBackgroundPosition(backgroundPosition objc.IObject /* cross-framework: NSString */)
	SetBackgroundRepeat(backgroundRepeat objc.IObject /* cross-framework: NSString */)
	SetBorder(border objc.IObject /* cross-framework: NSString */)
	SetBorderBottom(borderBottom objc.IObject /* cross-framework: NSString */)
	SetBorderBottomColor(borderBottomColor objc.IObject /* cross-framework: NSString */)
	SetBorderBottomStyle(borderBottomStyle objc.IObject /* cross-framework: NSString */)
	SetBorderBottomWidth(borderBottomWidth objc.IObject /* cross-framework: NSString */)
	SetBorderCollapse(borderCollapse objc.IObject /* cross-framework: NSString */)
	SetBorderColor(borderColor objc.IObject /* cross-framework: NSString */)
	SetBorderLeft(borderLeft objc.IObject /* cross-framework: NSString */)
	SetBorderLeftColor(borderLeftColor objc.IObject /* cross-framework: NSString */)
	SetBorderLeftStyle(borderLeftStyle objc.IObject /* cross-framework: NSString */)
	SetBorderLeftWidth(borderLeftWidth objc.IObject /* cross-framework: NSString */)
	SetBorderRight(borderRight objc.IObject /* cross-framework: NSString */)
	SetBorderRightColor(borderRightColor objc.IObject /* cross-framework: NSString */)
	SetBorderRightStyle(borderRightStyle objc.IObject /* cross-framework: NSString */)
	SetBorderRightWidth(borderRightWidth objc.IObject /* cross-framework: NSString */)
	SetBorderSpacing(borderSpacing objc.IObject /* cross-framework: NSString */)
	SetBorderStyle(borderStyle objc.IObject /* cross-framework: NSString */)
	SetBorderTop(borderTop objc.IObject /* cross-framework: NSString */)
	SetBorderTopColor(borderTopColor objc.IObject /* cross-framework: NSString */)
	SetBorderTopStyle(borderTopStyle objc.IObject /* cross-framework: NSString */)
	SetBorderTopWidth(borderTopWidth objc.IObject /* cross-framework: NSString */)
	SetBorderWidth(borderWidth objc.IObject /* cross-framework: NSString */)
	SetBottom(bottom objc.IObject /* cross-framework: NSString */)
	SetCaptionSide(captionSide objc.IObject /* cross-framework: NSString */)
	SetClear(clear objc.IObject /* cross-framework: NSString */)
	SetClip(clip objc.IObject /* cross-framework: NSString */)
	SetColor(color objc.IObject /* cross-framework: NSString */)
	SetContent(content objc.IObject /* cross-framework: NSString */)
	SetCounterIncrement(counterIncrement objc.IObject /* cross-framework: NSString */)
	SetCounterReset(counterReset objc.IObject /* cross-framework: NSString */)
	SetCssFloat(cssFloat objc.IObject /* cross-framework: NSString */)
	SetCue(cue objc.IObject /* cross-framework: NSString */)
	SetCueAfter(cueAfter objc.IObject /* cross-framework: NSString */)
	SetCueBefore(cueBefore objc.IObject /* cross-framework: NSString */)
	SetCursor(cursor objc.IObject /* cross-framework: NSString */)
	SetDirection(direction objc.IObject /* cross-framework: NSString */)
	SetDisplay(display objc.IObject /* cross-framework: NSString */)
	SetElevation(elevation objc.IObject /* cross-framework: NSString */)
	SetEmptyCells(emptyCells objc.IObject /* cross-framework: NSString */)
	SetFont(font objc.IObject /* cross-framework: NSString */)
	SetFontFamily(fontFamily objc.IObject /* cross-framework: NSString */)
	SetFontSize(fontSize objc.IObject /* cross-framework: NSString */)
	SetFontSizeAdjust(fontSizeAdjust objc.IObject /* cross-framework: NSString */)
	SetFontStretch(fontStretch objc.IObject /* cross-framework: NSString */)
	SetFontStyle(fontStyle objc.IObject /* cross-framework: NSString */)
	SetFontVariant(fontVariant objc.IObject /* cross-framework: NSString */)
	SetFontWeight(fontWeight objc.IObject /* cross-framework: NSString */)
	SetHeight(height objc.IObject /* cross-framework: NSString */)
	SetLeft(left objc.IObject /* cross-framework: NSString */)
	SetLetterSpacing(letterSpacing objc.IObject /* cross-framework: NSString */)
	SetLineHeight(lineHeight objc.IObject /* cross-framework: NSString */)
	SetListStyle(listStyle objc.IObject /* cross-framework: NSString */)
	SetListStyleImage(listStyleImage objc.IObject /* cross-framework: NSString */)
	SetListStylePosition(listStylePosition objc.IObject /* cross-framework: NSString */)
	SetListStyleType(listStyleType objc.IObject /* cross-framework: NSString */)
	SetMargin(margin objc.IObject /* cross-framework: NSString */)
	SetMarginBottom(marginBottom objc.IObject /* cross-framework: NSString */)
	SetMarginLeft(marginLeft objc.IObject /* cross-framework: NSString */)
	SetMarginRight(marginRight objc.IObject /* cross-framework: NSString */)
	SetMarginTop(marginTop objc.IObject /* cross-framework: NSString */)
	SetMarkerOffset(markerOffset objc.IObject /* cross-framework: NSString */)
	SetMarks(marks objc.IObject /* cross-framework: NSString */)
	SetMaxHeight(maxHeight objc.IObject /* cross-framework: NSString */)
	SetMaxWidth(maxWidth objc.IObject /* cross-framework: NSString */)
	SetMinHeight(minHeight objc.IObject /* cross-framework: NSString */)
	SetMinWidth(minWidth objc.IObject /* cross-framework: NSString */)
	SetOrphans(orphans objc.IObject /* cross-framework: NSString */)
	SetOutline(outline objc.IObject /* cross-framework: NSString */)
	SetOutlineColor(outlineColor objc.IObject /* cross-framework: NSString */)
	SetOutlineStyle(outlineStyle objc.IObject /* cross-framework: NSString */)
	SetOutlineWidth(outlineWidth objc.IObject /* cross-framework: NSString */)
	SetOverflow(overflow objc.IObject /* cross-framework: NSString */)
	SetPadding(padding objc.IObject /* cross-framework: NSString */)
	SetPaddingBottom(paddingBottom objc.IObject /* cross-framework: NSString */)
	SetPaddingLeft(paddingLeft objc.IObject /* cross-framework: NSString */)
	SetPaddingRight(paddingRight objc.IObject /* cross-framework: NSString */)
	SetPaddingTop(paddingTop objc.IObject /* cross-framework: NSString */)
	SetPage(page objc.IObject /* cross-framework: NSString */)
	SetPageBreakAfter(pageBreakAfter objc.IObject /* cross-framework: NSString */)
	SetPageBreakBefore(pageBreakBefore objc.IObject /* cross-framework: NSString */)
	SetPageBreakInside(pageBreakInside objc.IObject /* cross-framework: NSString */)
	SetPause(pause objc.IObject /* cross-framework: NSString */)
	SetPauseAfter(pauseAfter objc.IObject /* cross-framework: NSString */)
	SetPauseBefore(pauseBefore objc.IObject /* cross-framework: NSString */)
	SetPitch(pitch objc.IObject /* cross-framework: NSString */)
	SetPitchRange(pitchRange objc.IObject /* cross-framework: NSString */)
	SetPlayDuring(playDuring objc.IObject /* cross-framework: NSString */)
	SetPosition(position objc.IObject /* cross-framework: NSString */)
	SetQuotes(quotes objc.IObject /* cross-framework: NSString */)
	SetRichness(richness objc.IObject /* cross-framework: NSString */)
	SetRight(right objc.IObject /* cross-framework: NSString */)
	SetSize(size objc.IObject /* cross-framework: NSString */)
	SetSpeak(speak objc.IObject /* cross-framework: NSString */)
	SetSpeakHeader(speakHeader objc.IObject /* cross-framework: NSString */)
	SetSpeakNumeral(speakNumeral objc.IObject /* cross-framework: NSString */)
	SetSpeakPunctuation(speakPunctuation objc.IObject /* cross-framework: NSString */)
	SetSpeechRate(speechRate objc.IObject /* cross-framework: NSString */)
	SetStress(stress objc.IObject /* cross-framework: NSString */)
	SetTableLayout(tableLayout objc.IObject /* cross-framework: NSString */)
	SetTextAlign(textAlign objc.IObject /* cross-framework: NSString */)
	SetTextDecoration(textDecoration objc.IObject /* cross-framework: NSString */)
	SetTextIndent(textIndent objc.IObject /* cross-framework: NSString */)
	SetTextShadow(textShadow objc.IObject /* cross-framework: NSString */)
	SetTextTransform(textTransform objc.IObject /* cross-framework: NSString */)
	SetTop(top objc.IObject /* cross-framework: NSString */)
	SetUnicodeBidi(unicodeBidi objc.IObject /* cross-framework: NSString */)
	SetVerticalAlign(verticalAlign objc.IObject /* cross-framework: NSString */)
	SetVisibility(visibility objc.IObject /* cross-framework: NSString */)
	SetVoiceFamily(voiceFamily objc.IObject /* cross-framework: NSString */)
	SetVolume(volume objc.IObject /* cross-framework: NSString */)
	SetWhiteSpace(whiteSpace objc.IObject /* cross-framework: NSString */)
	SetWidows(widows objc.IObject /* cross-framework: NSString */)
	SetWidth(width objc.IObject /* cross-framework: NSString */)
	SetWordSpacing(wordSpacing objc.IObject /* cross-framework: NSString */)
	SetZIndex(zIndex objc.IObject /* cross-framework: NSString */)
	Size() foundation.String
	Speak() foundation.String
	SpeakHeader() foundation.String
	SpeakNumeral() foundation.String
	SpeakPunctuation() foundation.String
	SpeechRate() foundation.String
	Stress() foundation.String
	TableLayout() foundation.String
	TextAlign() foundation.String
	TextDecoration() foundation.String
	TextIndent() foundation.String
	TextShadow() foundation.String
	TextTransform() foundation.String
	Top() foundation.String
	UnicodeBidi() foundation.String
	VerticalAlign() foundation.String
	Visibility() foundation.String
	VoiceFamily() foundation.String
	Volume() foundation.String
	WhiteSpace() foundation.String
	Widows() foundation.String
	Width() foundation.String
	WordSpacing() foundation.String
	ZIndex() foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMCSSStyleDeclaration */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSStyleDeclarationClass) Alloc() DOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSStyleDeclarationClass) New() DOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSStyleDeclaration) Init() DOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSStyleDeclaration) Autorelease() DOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSStyleDeclaration creates a new DOMCSSStyleDeclaration instance.
func NewDOMCSSStyleDeclaration() DOMCSSStyleDeclaration {
	return getDOMCSSStyleDeclarationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMCSSStyleDeclaration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration
type DOMCSSStyleDeclaration struct {
	DOMObject
}

// DOMCSSStyleDeclarationFrom constructs a [DOMCSSStyleDeclaration] from an unsafe.Pointer.
func DOMCSSStyleDeclarationFrom(ptr unsafe.Pointer) DOMCSSStyleDeclaration {
	return DOMCSSStyleDeclaration{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMCSSStyleDeclaration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMCSSStyleDeclaration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMCSSStyleDeclaration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMCSSStyleDeclaration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/azimuth()
func (d_ DOMCSSStyleDeclaration) Azimuth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("azimuth"))
	return rv
}/* debug [instance_methods/method]: Azimuth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/background()
func (d_ DOMCSSStyleDeclaration) Background() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("background"))
	return rv
}/* debug [instance_methods/method]: Background */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/backgroundAttachment()
func (d_ DOMCSSStyleDeclaration) BackgroundAttachment() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("backgroundAttachment"))
	return rv
}/* debug [instance_methods/method]: BackgroundAttachment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/backgroundColor()
func (d_ DOMCSSStyleDeclaration) BackgroundColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_methods/method]: BackgroundColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/backgroundImage()
func (d_ DOMCSSStyleDeclaration) BackgroundImage() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("backgroundImage"))
	return rv
}/* debug [instance_methods/method]: BackgroundImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/backgroundPosition()
func (d_ DOMCSSStyleDeclaration) BackgroundPosition() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("backgroundPosition"))
	return rv
}/* debug [instance_methods/method]: BackgroundPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/backgroundRepeat()
func (d_ DOMCSSStyleDeclaration) BackgroundRepeat() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("backgroundRepeat"))
	return rv
}/* debug [instance_methods/method]: BackgroundRepeat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/border()
func (d_ DOMCSSStyleDeclaration) Border() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("border"))
	return rv
}/* debug [instance_methods/method]: Border */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderBottom()
func (d_ DOMCSSStyleDeclaration) BorderBottom() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderBottom"))
	return rv
}/* debug [instance_methods/method]: BorderBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderBottomColor()
func (d_ DOMCSSStyleDeclaration) BorderBottomColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderBottomColor"))
	return rv
}/* debug [instance_methods/method]: BorderBottomColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderBottomStyle()
func (d_ DOMCSSStyleDeclaration) BorderBottomStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderBottomStyle"))
	return rv
}/* debug [instance_methods/method]: BorderBottomStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderBottomWidth()
func (d_ DOMCSSStyleDeclaration) BorderBottomWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderBottomWidth"))
	return rv
}/* debug [instance_methods/method]: BorderBottomWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderCollapse()
func (d_ DOMCSSStyleDeclaration) BorderCollapse() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderCollapse"))
	return rv
}/* debug [instance_methods/method]: BorderCollapse */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderColor()
func (d_ DOMCSSStyleDeclaration) BorderColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderColor"))
	return rv
}/* debug [instance_methods/method]: BorderColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderLeft()
func (d_ DOMCSSStyleDeclaration) BorderLeft() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderLeft"))
	return rv
}/* debug [instance_methods/method]: BorderLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderLeftColor()
func (d_ DOMCSSStyleDeclaration) BorderLeftColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderLeftColor"))
	return rv
}/* debug [instance_methods/method]: BorderLeftColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderLeftStyle()
func (d_ DOMCSSStyleDeclaration) BorderLeftStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderLeftStyle"))
	return rv
}/* debug [instance_methods/method]: BorderLeftStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderLeftWidth()
func (d_ DOMCSSStyleDeclaration) BorderLeftWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderLeftWidth"))
	return rv
}/* debug [instance_methods/method]: BorderLeftWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderRight()
func (d_ DOMCSSStyleDeclaration) BorderRight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderRight"))
	return rv
}/* debug [instance_methods/method]: BorderRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderRightColor()
func (d_ DOMCSSStyleDeclaration) BorderRightColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderRightColor"))
	return rv
}/* debug [instance_methods/method]: BorderRightColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderRightStyle()
func (d_ DOMCSSStyleDeclaration) BorderRightStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderRightStyle"))
	return rv
}/* debug [instance_methods/method]: BorderRightStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderRightWidth()
func (d_ DOMCSSStyleDeclaration) BorderRightWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderRightWidth"))
	return rv
}/* debug [instance_methods/method]: BorderRightWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderSpacing()
func (d_ DOMCSSStyleDeclaration) BorderSpacing() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderSpacing"))
	return rv
}/* debug [instance_methods/method]: BorderSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderStyle()
func (d_ DOMCSSStyleDeclaration) BorderStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderStyle"))
	return rv
}/* debug [instance_methods/method]: BorderStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderTop()
func (d_ DOMCSSStyleDeclaration) BorderTop() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderTop"))
	return rv
}/* debug [instance_methods/method]: BorderTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderTopColor()
func (d_ DOMCSSStyleDeclaration) BorderTopColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderTopColor"))
	return rv
}/* debug [instance_methods/method]: BorderTopColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderTopStyle()
func (d_ DOMCSSStyleDeclaration) BorderTopStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderTopStyle"))
	return rv
}/* debug [instance_methods/method]: BorderTopStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderTopWidth()
func (d_ DOMCSSStyleDeclaration) BorderTopWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderTopWidth"))
	return rv
}/* debug [instance_methods/method]: BorderTopWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/borderWidth()
func (d_ DOMCSSStyleDeclaration) BorderWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("borderWidth"))
	return rv
}/* debug [instance_methods/method]: BorderWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/bottom()
func (d_ DOMCSSStyleDeclaration) Bottom() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("bottom"))
	return rv
}/* debug [instance_methods/method]: Bottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/captionSide()
func (d_ DOMCSSStyleDeclaration) CaptionSide() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("captionSide"))
	return rv
}/* debug [instance_methods/method]: CaptionSide */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/clear()
func (d_ DOMCSSStyleDeclaration) Clear() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("clear"))
	return rv
}/* debug [instance_methods/method]: Clear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/clip()
func (d_ DOMCSSStyleDeclaration) Clip() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("clip"))
	return rv
}/* debug [instance_methods/method]: Clip */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/color()
func (d_ DOMCSSStyleDeclaration) Color() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_methods/method]: Color */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/content()
func (d_ DOMCSSStyleDeclaration) Content() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_methods/method]: Content */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/counterIncrement()
func (d_ DOMCSSStyleDeclaration) CounterIncrement() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("counterIncrement"))
	return rv
}/* debug [instance_methods/method]: CounterIncrement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/counterReset()
func (d_ DOMCSSStyleDeclaration) CounterReset() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("counterReset"))
	return rv
}/* debug [instance_methods/method]: CounterReset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cssFloat()
func (d_ DOMCSSStyleDeclaration) CssFloat() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("cssFloat"))
	return rv
}/* debug [instance_methods/method]: CssFloat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cue()
func (d_ DOMCSSStyleDeclaration) Cue() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("cue"))
	return rv
}/* debug [instance_methods/method]: Cue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cueAfter()
func (d_ DOMCSSStyleDeclaration) CueAfter() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("cueAfter"))
	return rv
}/* debug [instance_methods/method]: CueAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cueBefore()
func (d_ DOMCSSStyleDeclaration) CueBefore() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("cueBefore"))
	return rv
}/* debug [instance_methods/method]: CueBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cursor()
func (d_ DOMCSSStyleDeclaration) Cursor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("cursor"))
	return rv
}/* debug [instance_methods/method]: Cursor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/direction()
func (d_ DOMCSSStyleDeclaration) Direction() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_methods/method]: Direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/display()
func (d_ DOMCSSStyleDeclaration) Display() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("display"))
	return rv
}/* debug [instance_methods/method]: Display */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/elevation()
func (d_ DOMCSSStyleDeclaration) Elevation() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("elevation"))
	return rv
}/* debug [instance_methods/method]: Elevation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/emptyCells()
func (d_ DOMCSSStyleDeclaration) EmptyCells() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("emptyCells"))
	return rv
}/* debug [instance_methods/method]: EmptyCells */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/font()
func (d_ DOMCSSStyleDeclaration) Font() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_methods/method]: Font */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontFamily()
func (d_ DOMCSSStyleDeclaration) FontFamily() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontFamily"))
	return rv
}/* debug [instance_methods/method]: FontFamily */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontSize()
func (d_ DOMCSSStyleDeclaration) FontSize() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontSize"))
	return rv
}/* debug [instance_methods/method]: FontSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontSizeAdjust()
func (d_ DOMCSSStyleDeclaration) FontSizeAdjust() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontSizeAdjust"))
	return rv
}/* debug [instance_methods/method]: FontSizeAdjust */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontStretch()
func (d_ DOMCSSStyleDeclaration) FontStretch() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontStretch"))
	return rv
}/* debug [instance_methods/method]: FontStretch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontStyle()
func (d_ DOMCSSStyleDeclaration) FontStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontStyle"))
	return rv
}/* debug [instance_methods/method]: FontStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontVariant()
func (d_ DOMCSSStyleDeclaration) FontVariant() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontVariant"))
	return rv
}/* debug [instance_methods/method]: FontVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/fontWeight()
func (d_ DOMCSSStyleDeclaration) FontWeight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fontWeight"))
	return rv
}/* debug [instance_methods/method]: FontWeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/height()
func (d_ DOMCSSStyleDeclaration) Height() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_methods/method]: Height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/left()
func (d_ DOMCSSStyleDeclaration) Left() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("left"))
	return rv
}/* debug [instance_methods/method]: Left */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/letterSpacing()
func (d_ DOMCSSStyleDeclaration) LetterSpacing() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("letterSpacing"))
	return rv
}/* debug [instance_methods/method]: LetterSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/lineHeight()
func (d_ DOMCSSStyleDeclaration) LineHeight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("lineHeight"))
	return rv
}/* debug [instance_methods/method]: LineHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/listStyle()
func (d_ DOMCSSStyleDeclaration) ListStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("listStyle"))
	return rv
}/* debug [instance_methods/method]: ListStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/listStyleImage()
func (d_ DOMCSSStyleDeclaration) ListStyleImage() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("listStyleImage"))
	return rv
}/* debug [instance_methods/method]: ListStyleImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/listStylePosition()
func (d_ DOMCSSStyleDeclaration) ListStylePosition() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("listStylePosition"))
	return rv
}/* debug [instance_methods/method]: ListStylePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/listStyleType()
func (d_ DOMCSSStyleDeclaration) ListStyleType() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("listStyleType"))
	return rv
}/* debug [instance_methods/method]: ListStyleType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/margin()
func (d_ DOMCSSStyleDeclaration) Margin() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("margin"))
	return rv
}/* debug [instance_methods/method]: Margin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/marginBottom()
func (d_ DOMCSSStyleDeclaration) MarginBottom() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("marginBottom"))
	return rv
}/* debug [instance_methods/method]: MarginBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/marginLeft()
func (d_ DOMCSSStyleDeclaration) MarginLeft() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("marginLeft"))
	return rv
}/* debug [instance_methods/method]: MarginLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/marginRight()
func (d_ DOMCSSStyleDeclaration) MarginRight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("marginRight"))
	return rv
}/* debug [instance_methods/method]: MarginRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/marginTop()
func (d_ DOMCSSStyleDeclaration) MarginTop() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("marginTop"))
	return rv
}/* debug [instance_methods/method]: MarginTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/markerOffset()
func (d_ DOMCSSStyleDeclaration) MarkerOffset() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("markerOffset"))
	return rv
}/* debug [instance_methods/method]: MarkerOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/marks()
func (d_ DOMCSSStyleDeclaration) Marks() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("marks"))
	return rv
}/* debug [instance_methods/method]: Marks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/maxHeight()
func (d_ DOMCSSStyleDeclaration) MaxHeight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("maxHeight"))
	return rv
}/* debug [instance_methods/method]: MaxHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/maxWidth()
func (d_ DOMCSSStyleDeclaration) MaxWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("maxWidth"))
	return rv
}/* debug [instance_methods/method]: MaxWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/minHeight()
func (d_ DOMCSSStyleDeclaration) MinHeight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("minHeight"))
	return rv
}/* debug [instance_methods/method]: MinHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/minWidth()
func (d_ DOMCSSStyleDeclaration) MinWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("minWidth"))
	return rv
}/* debug [instance_methods/method]: MinWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/orphans()
func (d_ DOMCSSStyleDeclaration) Orphans() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("orphans"))
	return rv
}/* debug [instance_methods/method]: Orphans */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/outline()
func (d_ DOMCSSStyleDeclaration) Outline() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("outline"))
	return rv
}/* debug [instance_methods/method]: Outline */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/outlineColor()
func (d_ DOMCSSStyleDeclaration) OutlineColor() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("outlineColor"))
	return rv
}/* debug [instance_methods/method]: OutlineColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/outlineStyle()
func (d_ DOMCSSStyleDeclaration) OutlineStyle() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("outlineStyle"))
	return rv
}/* debug [instance_methods/method]: OutlineStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/outlineWidth()
func (d_ DOMCSSStyleDeclaration) OutlineWidth() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("outlineWidth"))
	return rv
}/* debug [instance_methods/method]: OutlineWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/overflow()
func (d_ DOMCSSStyleDeclaration) Overflow() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("overflow"))
	return rv
}/* debug [instance_methods/method]: Overflow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/padding()
func (d_ DOMCSSStyleDeclaration) Padding() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("padding"))
	return rv
}/* debug [instance_methods/method]: Padding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/paddingBottom()
func (d_ DOMCSSStyleDeclaration) PaddingBottom() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("paddingBottom"))
	return rv
}/* debug [instance_methods/method]: PaddingBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/paddingLeft()
func (d_ DOMCSSStyleDeclaration) PaddingLeft() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("paddingLeft"))
	return rv
}/* debug [instance_methods/method]: PaddingLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/paddingRight()
func (d_ DOMCSSStyleDeclaration) PaddingRight() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("paddingRight"))
	return rv
}/* debug [instance_methods/method]: PaddingRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/paddingTop()
func (d_ DOMCSSStyleDeclaration) PaddingTop() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("paddingTop"))
	return rv
}/* debug [instance_methods/method]: PaddingTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/page()
func (d_ DOMCSSStyleDeclaration) Page() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("page"))
	return rv
}/* debug [instance_methods/method]: Page */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pageBreakAfter()
func (d_ DOMCSSStyleDeclaration) PageBreakAfter() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pageBreakAfter"))
	return rv
}/* debug [instance_methods/method]: PageBreakAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pageBreakBefore()
func (d_ DOMCSSStyleDeclaration) PageBreakBefore() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pageBreakBefore"))
	return rv
}/* debug [instance_methods/method]: PageBreakBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pageBreakInside()
func (d_ DOMCSSStyleDeclaration) PageBreakInside() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pageBreakInside"))
	return rv
}/* debug [instance_methods/method]: PageBreakInside */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pause()
func (d_ DOMCSSStyleDeclaration) Pause() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pause"))
	return rv
}/* debug [instance_methods/method]: Pause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pauseAfter()
func (d_ DOMCSSStyleDeclaration) PauseAfter() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pauseAfter"))
	return rv
}/* debug [instance_methods/method]: PauseAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pauseBefore()
func (d_ DOMCSSStyleDeclaration) PauseBefore() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pauseBefore"))
	return rv
}/* debug [instance_methods/method]: PauseBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pitch()
func (d_ DOMCSSStyleDeclaration) Pitch() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pitch"))
	return rv
}/* debug [instance_methods/method]: Pitch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/pitchRange()
func (d_ DOMCSSStyleDeclaration) PitchRange() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("pitchRange"))
	return rv
}/* debug [instance_methods/method]: PitchRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/playDuring()
func (d_ DOMCSSStyleDeclaration) PlayDuring() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("playDuring"))
	return rv
}/* debug [instance_methods/method]: PlayDuring */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/position()
func (d_ DOMCSSStyleDeclaration) Position() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_methods/method]: Position */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/quotes()
func (d_ DOMCSSStyleDeclaration) Quotes() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("quotes"))
	return rv
}/* debug [instance_methods/method]: Quotes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/richness()
func (d_ DOMCSSStyleDeclaration) Richness() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("richness"))
	return rv
}/* debug [instance_methods/method]: Richness */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/right()
func (d_ DOMCSSStyleDeclaration) Right() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("right"))
	return rv
}/* debug [instance_methods/method]: Right */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setAzimuth(_:)
func (d_ DOMCSSStyleDeclaration) SetAzimuth(azimuth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAzimuth:"), azimuth)
}/* debug [instance_methods/method]: SetAzimuth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackground(_:)
func (d_ DOMCSSStyleDeclaration) SetBackground(background objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackground:"), background)
}/* debug [instance_methods/method]: SetBackground */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackgroundAttachment(_:)
func (d_ DOMCSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundAttachment:"), backgroundAttachment)
}/* debug [instance_methods/method]: SetBackgroundAttachment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackgroundColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBackgroundColor(backgroundColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), backgroundColor)
}/* debug [instance_methods/method]: SetBackgroundColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackgroundImage(_:)
func (d_ DOMCSSStyleDeclaration) SetBackgroundImage(backgroundImage objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundImage:"), backgroundImage)
}/* debug [instance_methods/method]: SetBackgroundImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackgroundPosition(_:)
func (d_ DOMCSSStyleDeclaration) SetBackgroundPosition(backgroundPosition objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundPosition:"), backgroundPosition)
}/* debug [instance_methods/method]: SetBackgroundPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBackgroundRepeat(_:)
func (d_ DOMCSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundRepeat:"), backgroundRepeat)
}/* debug [instance_methods/method]: SetBackgroundRepeat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorder(_:)
func (d_ DOMCSSStyleDeclaration) SetBorder(border objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorder:"), border)
}/* debug [instance_methods/method]: SetBorder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderBottom(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderBottom(borderBottom objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderBottom:"), borderBottom)
}/* debug [instance_methods/method]: SetBorderBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderBottomColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderBottomColor(borderBottomColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderBottomColor:"), borderBottomColor)
}/* debug [instance_methods/method]: SetBorderBottomColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderBottomStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderBottomStyle:"), borderBottomStyle)
}/* debug [instance_methods/method]: SetBorderBottomStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderBottomWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderBottomWidth:"), borderBottomWidth)
}/* debug [instance_methods/method]: SetBorderBottomWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderCollapse(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderCollapse(borderCollapse objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderCollapse:"), borderCollapse)
}/* debug [instance_methods/method]: SetBorderCollapse */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderColor(borderColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderColor:"), borderColor)
}/* debug [instance_methods/method]: SetBorderColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderLeft(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderLeft(borderLeft objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderLeft:"), borderLeft)
}/* debug [instance_methods/method]: SetBorderLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderLeftColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderLeftColor(borderLeftColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderLeftColor:"), borderLeftColor)
}/* debug [instance_methods/method]: SetBorderLeftColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderLeftStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderLeftStyle:"), borderLeftStyle)
}/* debug [instance_methods/method]: SetBorderLeftStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderLeftWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderLeftWidth:"), borderLeftWidth)
}/* debug [instance_methods/method]: SetBorderLeftWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderRight(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderRight(borderRight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderRight:"), borderRight)
}/* debug [instance_methods/method]: SetBorderRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderRightColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderRightColor(borderRightColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderRightColor:"), borderRightColor)
}/* debug [instance_methods/method]: SetBorderRightColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderRightStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderRightStyle(borderRightStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderRightStyle:"), borderRightStyle)
}/* debug [instance_methods/method]: SetBorderRightStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderRightWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderRightWidth(borderRightWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderRightWidth:"), borderRightWidth)
}/* debug [instance_methods/method]: SetBorderRightWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderSpacing(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderSpacing(borderSpacing objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderSpacing:"), borderSpacing)
}/* debug [instance_methods/method]: SetBorderSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderStyle(borderStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderStyle:"), borderStyle)
}/* debug [instance_methods/method]: SetBorderStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderTop(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderTop(borderTop objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderTop:"), borderTop)
}/* debug [instance_methods/method]: SetBorderTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderTopColor(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderTopColor(borderTopColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderTopColor:"), borderTopColor)
}/* debug [instance_methods/method]: SetBorderTopColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderTopStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderTopStyle(borderTopStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderTopStyle:"), borderTopStyle)
}/* debug [instance_methods/method]: SetBorderTopStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderTopWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderTopWidth(borderTopWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderTopWidth:"), borderTopWidth)
}/* debug [instance_methods/method]: SetBorderTopWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBorderWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetBorderWidth(borderWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorderWidth:"), borderWidth)
}/* debug [instance_methods/method]: SetBorderWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setBottom(_:)
func (d_ DOMCSSStyleDeclaration) SetBottom(bottom objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBottom:"), bottom)
}/* debug [instance_methods/method]: SetBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCaptionSide(_:)
func (d_ DOMCSSStyleDeclaration) SetCaptionSide(captionSide objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCaptionSide:"), captionSide)
}/* debug [instance_methods/method]: SetCaptionSide */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setClear(_:)
func (d_ DOMCSSStyleDeclaration) SetClear(clear objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setClear:"), clear)
}/* debug [instance_methods/method]: SetClear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setClip(_:)
func (d_ DOMCSSStyleDeclaration) SetClip(clip objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setClip:"), clip)
}/* debug [instance_methods/method]: SetClip */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setColor(_:)
func (d_ DOMCSSStyleDeclaration) SetColor(color objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setColor:"), color)
}/* debug [instance_methods/method]: SetColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setContent(_:)
func (d_ DOMCSSStyleDeclaration) SetContent(content objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContent:"), content)
}/* debug [instance_methods/method]: SetContent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCounterIncrement(_:)
func (d_ DOMCSSStyleDeclaration) SetCounterIncrement(counterIncrement objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCounterIncrement:"), counterIncrement)
}/* debug [instance_methods/method]: SetCounterIncrement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCounterReset(_:)
func (d_ DOMCSSStyleDeclaration) SetCounterReset(counterReset objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCounterReset:"), counterReset)
}/* debug [instance_methods/method]: SetCounterReset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCssFloat(_:)
func (d_ DOMCSSStyleDeclaration) SetCssFloat(cssFloat objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCssFloat:"), cssFloat)
}/* debug [instance_methods/method]: SetCssFloat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCue(_:)
func (d_ DOMCSSStyleDeclaration) SetCue(cue objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCue:"), cue)
}/* debug [instance_methods/method]: SetCue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCueAfter(_:)
func (d_ DOMCSSStyleDeclaration) SetCueAfter(cueAfter objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCueAfter:"), cueAfter)
}/* debug [instance_methods/method]: SetCueAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCueBefore(_:)
func (d_ DOMCSSStyleDeclaration) SetCueBefore(cueBefore objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCueBefore:"), cueBefore)
}/* debug [instance_methods/method]: SetCueBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setCursor(_:)
func (d_ DOMCSSStyleDeclaration) SetCursor(cursor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCursor:"), cursor)
}/* debug [instance_methods/method]: SetCursor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setDirection(_:)
func (d_ DOMCSSStyleDeclaration) SetDirection(direction objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDirection:"), direction)
}/* debug [instance_methods/method]: SetDirection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setDisplay(_:)
func (d_ DOMCSSStyleDeclaration) SetDisplay(display objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplay:"), display)
}/* debug [instance_methods/method]: SetDisplay */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setElevation(_:)
func (d_ DOMCSSStyleDeclaration) SetElevation(elevation objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setElevation:"), elevation)
}/* debug [instance_methods/method]: SetElevation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setEmptyCells(_:)
func (d_ DOMCSSStyleDeclaration) SetEmptyCells(emptyCells objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEmptyCells:"), emptyCells)
}/* debug [instance_methods/method]: SetEmptyCells */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFont(_:)
func (d_ DOMCSSStyleDeclaration) SetFont(font objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFont:"), font)
}/* debug [instance_methods/method]: SetFont */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontFamily(_:)
func (d_ DOMCSSStyleDeclaration) SetFontFamily(fontFamily objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontFamily:"), fontFamily)
}/* debug [instance_methods/method]: SetFontFamily */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontSize(_:)
func (d_ DOMCSSStyleDeclaration) SetFontSize(fontSize objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontSize:"), fontSize)
}/* debug [instance_methods/method]: SetFontSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontSizeAdjust(_:)
func (d_ DOMCSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontSizeAdjust:"), fontSizeAdjust)
}/* debug [instance_methods/method]: SetFontSizeAdjust */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontStretch(_:)
func (d_ DOMCSSStyleDeclaration) SetFontStretch(fontStretch objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontStretch:"), fontStretch)
}/* debug [instance_methods/method]: SetFontStretch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetFontStyle(fontStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontStyle:"), fontStyle)
}/* debug [instance_methods/method]: SetFontStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontVariant(_:)
func (d_ DOMCSSStyleDeclaration) SetFontVariant(fontVariant objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontVariant:"), fontVariant)
}/* debug [instance_methods/method]: SetFontVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setFontWeight(_:)
func (d_ DOMCSSStyleDeclaration) SetFontWeight(fontWeight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFontWeight:"), fontWeight)
}/* debug [instance_methods/method]: SetFontWeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setHeight(_:)
func (d_ DOMCSSStyleDeclaration) SetHeight(height objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), height)
}/* debug [instance_methods/method]: SetHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setLeft(_:)
func (d_ DOMCSSStyleDeclaration) SetLeft(left objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLeft:"), left)
}/* debug [instance_methods/method]: SetLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setLetterSpacing(_:)
func (d_ DOMCSSStyleDeclaration) SetLetterSpacing(letterSpacing objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLetterSpacing:"), letterSpacing)
}/* debug [instance_methods/method]: SetLetterSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setLineHeight(_:)
func (d_ DOMCSSStyleDeclaration) SetLineHeight(lineHeight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLineHeight:"), lineHeight)
}/* debug [instance_methods/method]: SetLineHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setListStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetListStyle(listStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setListStyle:"), listStyle)
}/* debug [instance_methods/method]: SetListStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setListStyleImage(_:)
func (d_ DOMCSSStyleDeclaration) SetListStyleImage(listStyleImage objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setListStyleImage:"), listStyleImage)
}/* debug [instance_methods/method]: SetListStyleImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setListStylePosition(_:)
func (d_ DOMCSSStyleDeclaration) SetListStylePosition(listStylePosition objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setListStylePosition:"), listStylePosition)
}/* debug [instance_methods/method]: SetListStylePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setListStyleType(_:)
func (d_ DOMCSSStyleDeclaration) SetListStyleType(listStyleType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setListStyleType:"), listStyleType)
}/* debug [instance_methods/method]: SetListStyleType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMargin(_:)
func (d_ DOMCSSStyleDeclaration) SetMargin(margin objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMargin:"), margin)
}/* debug [instance_methods/method]: SetMargin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarginBottom(_:)
func (d_ DOMCSSStyleDeclaration) SetMarginBottom(marginBottom objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginBottom:"), marginBottom)
}/* debug [instance_methods/method]: SetMarginBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarginLeft(_:)
func (d_ DOMCSSStyleDeclaration) SetMarginLeft(marginLeft objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginLeft:"), marginLeft)
}/* debug [instance_methods/method]: SetMarginLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarginRight(_:)
func (d_ DOMCSSStyleDeclaration) SetMarginRight(marginRight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginRight:"), marginRight)
}/* debug [instance_methods/method]: SetMarginRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarginTop(_:)
func (d_ DOMCSSStyleDeclaration) SetMarginTop(marginTop objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginTop:"), marginTop)
}/* debug [instance_methods/method]: SetMarginTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarkerOffset(_:)
func (d_ DOMCSSStyleDeclaration) SetMarkerOffset(markerOffset objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarkerOffset:"), markerOffset)
}/* debug [instance_methods/method]: SetMarkerOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMarks(_:)
func (d_ DOMCSSStyleDeclaration) SetMarks(marks objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarks:"), marks)
}/* debug [instance_methods/method]: SetMarks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMaxHeight(_:)
func (d_ DOMCSSStyleDeclaration) SetMaxHeight(maxHeight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxHeight:"), maxHeight)
}/* debug [instance_methods/method]: SetMaxHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMaxWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetMaxWidth(maxWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxWidth:"), maxWidth)
}/* debug [instance_methods/method]: SetMaxWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMinHeight(_:)
func (d_ DOMCSSStyleDeclaration) SetMinHeight(minHeight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinHeight:"), minHeight)
}/* debug [instance_methods/method]: SetMinHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setMinWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetMinWidth(minWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinWidth:"), minWidth)
}/* debug [instance_methods/method]: SetMinWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOrphans(_:)
func (d_ DOMCSSStyleDeclaration) SetOrphans(orphans objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOrphans:"), orphans)
}/* debug [instance_methods/method]: SetOrphans */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOutline(_:)
func (d_ DOMCSSStyleDeclaration) SetOutline(outline objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOutline:"), outline)
}/* debug [instance_methods/method]: SetOutline */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOutlineColor(_:)
func (d_ DOMCSSStyleDeclaration) SetOutlineColor(outlineColor objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOutlineColor:"), outlineColor)
}/* debug [instance_methods/method]: SetOutlineColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOutlineStyle(_:)
func (d_ DOMCSSStyleDeclaration) SetOutlineStyle(outlineStyle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOutlineStyle:"), outlineStyle)
}/* debug [instance_methods/method]: SetOutlineStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOutlineWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetOutlineWidth(outlineWidth objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOutlineWidth:"), outlineWidth)
}/* debug [instance_methods/method]: SetOutlineWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setOverflow(_:)
func (d_ DOMCSSStyleDeclaration) SetOverflow(overflow objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOverflow:"), overflow)
}/* debug [instance_methods/method]: SetOverflow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPadding(_:)
func (d_ DOMCSSStyleDeclaration) SetPadding(padding objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPadding:"), padding)
}/* debug [instance_methods/method]: SetPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPaddingBottom(_:)
func (d_ DOMCSSStyleDeclaration) SetPaddingBottom(paddingBottom objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaddingBottom:"), paddingBottom)
}/* debug [instance_methods/method]: SetPaddingBottom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPaddingLeft(_:)
func (d_ DOMCSSStyleDeclaration) SetPaddingLeft(paddingLeft objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaddingLeft:"), paddingLeft)
}/* debug [instance_methods/method]: SetPaddingLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPaddingRight(_:)
func (d_ DOMCSSStyleDeclaration) SetPaddingRight(paddingRight objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaddingRight:"), paddingRight)
}/* debug [instance_methods/method]: SetPaddingRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPaddingTop(_:)
func (d_ DOMCSSStyleDeclaration) SetPaddingTop(paddingTop objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPaddingTop:"), paddingTop)
}/* debug [instance_methods/method]: SetPaddingTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPage(_:)
func (d_ DOMCSSStyleDeclaration) SetPage(page objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPage:"), page)
}/* debug [instance_methods/method]: SetPage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPageBreakAfter(_:)
func (d_ DOMCSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPageBreakAfter:"), pageBreakAfter)
}/* debug [instance_methods/method]: SetPageBreakAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPageBreakBefore(_:)
func (d_ DOMCSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPageBreakBefore:"), pageBreakBefore)
}/* debug [instance_methods/method]: SetPageBreakBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPageBreakInside(_:)
func (d_ DOMCSSStyleDeclaration) SetPageBreakInside(pageBreakInside objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPageBreakInside:"), pageBreakInside)
}/* debug [instance_methods/method]: SetPageBreakInside */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPause(_:)
func (d_ DOMCSSStyleDeclaration) SetPause(pause objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPause:"), pause)
}/* debug [instance_methods/method]: SetPause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPauseAfter(_:)
func (d_ DOMCSSStyleDeclaration) SetPauseAfter(pauseAfter objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPauseAfter:"), pauseAfter)
}/* debug [instance_methods/method]: SetPauseAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPauseBefore(_:)
func (d_ DOMCSSStyleDeclaration) SetPauseBefore(pauseBefore objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPauseBefore:"), pauseBefore)
}/* debug [instance_methods/method]: SetPauseBefore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPitch(_:)
func (d_ DOMCSSStyleDeclaration) SetPitch(pitch objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPitch:"), pitch)
}/* debug [instance_methods/method]: SetPitch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPitchRange(_:)
func (d_ DOMCSSStyleDeclaration) SetPitchRange(pitchRange objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPitchRange:"), pitchRange)
}/* debug [instance_methods/method]: SetPitchRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPlayDuring(_:)
func (d_ DOMCSSStyleDeclaration) SetPlayDuring(playDuring objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPlayDuring:"), playDuring)
}/* debug [instance_methods/method]: SetPlayDuring */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setPosition(_:)
func (d_ DOMCSSStyleDeclaration) SetPosition(position objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPosition:"), position)
}/* debug [instance_methods/method]: SetPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setQuotes(_:)
func (d_ DOMCSSStyleDeclaration) SetQuotes(quotes objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuotes:"), quotes)
}/* debug [instance_methods/method]: SetQuotes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setRichness(_:)
func (d_ DOMCSSStyleDeclaration) SetRichness(richness objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRichness:"), richness)
}/* debug [instance_methods/method]: SetRichness */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setRight(_:)
func (d_ DOMCSSStyleDeclaration) SetRight(right objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRight:"), right)
}/* debug [instance_methods/method]: SetRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSize(_:)
func (d_ DOMCSSStyleDeclaration) SetSize(size objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), size)
}/* debug [instance_methods/method]: SetSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSpeak(_:)
func (d_ DOMCSSStyleDeclaration) SetSpeak(speak objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpeak:"), speak)
}/* debug [instance_methods/method]: SetSpeak */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSpeakHeader(_:)
func (d_ DOMCSSStyleDeclaration) SetSpeakHeader(speakHeader objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpeakHeader:"), speakHeader)
}/* debug [instance_methods/method]: SetSpeakHeader */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSpeakNumeral(_:)
func (d_ DOMCSSStyleDeclaration) SetSpeakNumeral(speakNumeral objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpeakNumeral:"), speakNumeral)
}/* debug [instance_methods/method]: SetSpeakNumeral */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSpeakPunctuation(_:)
func (d_ DOMCSSStyleDeclaration) SetSpeakPunctuation(speakPunctuation objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpeakPunctuation:"), speakPunctuation)
}/* debug [instance_methods/method]: SetSpeakPunctuation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setSpeechRate(_:)
func (d_ DOMCSSStyleDeclaration) SetSpeechRate(speechRate objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpeechRate:"), speechRate)
}/* debug [instance_methods/method]: SetSpeechRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setStress(_:)
func (d_ DOMCSSStyleDeclaration) SetStress(stress objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStress:"), stress)
}/* debug [instance_methods/method]: SetStress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTableLayout(_:)
func (d_ DOMCSSStyleDeclaration) SetTableLayout(tableLayout objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTableLayout:"), tableLayout)
}/* debug [instance_methods/method]: SetTableLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTextAlign(_:)
func (d_ DOMCSSStyleDeclaration) SetTextAlign(textAlign objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextAlign:"), textAlign)
}/* debug [instance_methods/method]: SetTextAlign */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTextDecoration(_:)
func (d_ DOMCSSStyleDeclaration) SetTextDecoration(textDecoration objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextDecoration:"), textDecoration)
}/* debug [instance_methods/method]: SetTextDecoration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTextIndent(_:)
func (d_ DOMCSSStyleDeclaration) SetTextIndent(textIndent objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextIndent:"), textIndent)
}/* debug [instance_methods/method]: SetTextIndent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTextShadow(_:)
func (d_ DOMCSSStyleDeclaration) SetTextShadow(textShadow objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextShadow:"), textShadow)
}/* debug [instance_methods/method]: SetTextShadow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTextTransform(_:)
func (d_ DOMCSSStyleDeclaration) SetTextTransform(textTransform objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextTransform:"), textTransform)
}/* debug [instance_methods/method]: SetTextTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setTop(_:)
func (d_ DOMCSSStyleDeclaration) SetTop(top objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTop:"), top)
}/* debug [instance_methods/method]: SetTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setUnicodeBidi(_:)
func (d_ DOMCSSStyleDeclaration) SetUnicodeBidi(unicodeBidi objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUnicodeBidi:"), unicodeBidi)
}/* debug [instance_methods/method]: SetUnicodeBidi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setVerticalAlign(_:)
func (d_ DOMCSSStyleDeclaration) SetVerticalAlign(verticalAlign objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVerticalAlign:"), verticalAlign)
}/* debug [instance_methods/method]: SetVerticalAlign */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setVisibility(_:)
func (d_ DOMCSSStyleDeclaration) SetVisibility(visibility objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVisibility:"), visibility)
}/* debug [instance_methods/method]: SetVisibility */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setVoiceFamily(_:)
func (d_ DOMCSSStyleDeclaration) SetVoiceFamily(voiceFamily objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVoiceFamily:"), voiceFamily)
}/* debug [instance_methods/method]: SetVoiceFamily */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setVolume(_:)
func (d_ DOMCSSStyleDeclaration) SetVolume(volume objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVolume:"), volume)
}/* debug [instance_methods/method]: SetVolume */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setWhiteSpace(_:)
func (d_ DOMCSSStyleDeclaration) SetWhiteSpace(whiteSpace objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWhiteSpace:"), whiteSpace)
}/* debug [instance_methods/method]: SetWhiteSpace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setWidows(_:)
func (d_ DOMCSSStyleDeclaration) SetWidows(widows objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidows:"), widows)
}/* debug [instance_methods/method]: SetWidows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setWidth(_:)
func (d_ DOMCSSStyleDeclaration) SetWidth(width objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), width)
}/* debug [instance_methods/method]: SetWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setWordSpacing(_:)
func (d_ DOMCSSStyleDeclaration) SetWordSpacing(wordSpacing objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWordSpacing:"), wordSpacing)
}/* debug [instance_methods/method]: SetWordSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/setZIndex(_:)
func (d_ DOMCSSStyleDeclaration) SetZIndex(zIndex objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setZIndex:"), zIndex)
}/* debug [instance_methods/method]: SetZIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/size()
func (d_ DOMCSSStyleDeclaration) Size() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_methods/method]: Size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/speak()
func (d_ DOMCSSStyleDeclaration) Speak() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("speak"))
	return rv
}/* debug [instance_methods/method]: Speak */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/speakHeader()
func (d_ DOMCSSStyleDeclaration) SpeakHeader() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("speakHeader"))
	return rv
}/* debug [instance_methods/method]: SpeakHeader */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/speakNumeral()
func (d_ DOMCSSStyleDeclaration) SpeakNumeral() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("speakNumeral"))
	return rv
}/* debug [instance_methods/method]: SpeakNumeral */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/speakPunctuation()
func (d_ DOMCSSStyleDeclaration) SpeakPunctuation() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("speakPunctuation"))
	return rv
}/* debug [instance_methods/method]: SpeakPunctuation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/speechRate()
func (d_ DOMCSSStyleDeclaration) SpeechRate() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("speechRate"))
	return rv
}/* debug [instance_methods/method]: SpeechRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/stress()
func (d_ DOMCSSStyleDeclaration) Stress() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("stress"))
	return rv
}/* debug [instance_methods/method]: Stress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/tableLayout()
func (d_ DOMCSSStyleDeclaration) TableLayout() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("tableLayout"))
	return rv
}/* debug [instance_methods/method]: TableLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/textAlign()
func (d_ DOMCSSStyleDeclaration) TextAlign() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("textAlign"))
	return rv
}/* debug [instance_methods/method]: TextAlign */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/textDecoration()
func (d_ DOMCSSStyleDeclaration) TextDecoration() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("textDecoration"))
	return rv
}/* debug [instance_methods/method]: TextDecoration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/textIndent()
func (d_ DOMCSSStyleDeclaration) TextIndent() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("textIndent"))
	return rv
}/* debug [instance_methods/method]: TextIndent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/textShadow()
func (d_ DOMCSSStyleDeclaration) TextShadow() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("textShadow"))
	return rv
}/* debug [instance_methods/method]: TextShadow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/textTransform()
func (d_ DOMCSSStyleDeclaration) TextTransform() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("textTransform"))
	return rv
}/* debug [instance_methods/method]: TextTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/top()
func (d_ DOMCSSStyleDeclaration) Top() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("top"))
	return rv
}/* debug [instance_methods/method]: Top */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/unicodeBidi()
func (d_ DOMCSSStyleDeclaration) UnicodeBidi() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("unicodeBidi"))
	return rv
}/* debug [instance_methods/method]: UnicodeBidi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/verticalAlign()
func (d_ DOMCSSStyleDeclaration) VerticalAlign() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("verticalAlign"))
	return rv
}/* debug [instance_methods/method]: VerticalAlign */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/visibility()
func (d_ DOMCSSStyleDeclaration) Visibility() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("visibility"))
	return rv
}/* debug [instance_methods/method]: Visibility */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/voiceFamily()
func (d_ DOMCSSStyleDeclaration) VoiceFamily() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("voiceFamily"))
	return rv
}/* debug [instance_methods/method]: VoiceFamily */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/volume()
func (d_ DOMCSSStyleDeclaration) Volume() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_methods/method]: Volume */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/whiteSpace()
func (d_ DOMCSSStyleDeclaration) WhiteSpace() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("whiteSpace"))
	return rv
}/* debug [instance_methods/method]: WhiteSpace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/widows()
func (d_ DOMCSSStyleDeclaration) Widows() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("widows"))
	return rv
}/* debug [instance_methods/method]: Widows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/width()
func (d_ DOMCSSStyleDeclaration) Width() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_methods/method]: Width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/wordSpacing()
func (d_ DOMCSSStyleDeclaration) WordSpacing() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("wordSpacing"))
	return rv
}/* debug [instance_methods/method]: WordSpacing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/zIndex()
func (d_ DOMCSSStyleDeclaration) ZIndex() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("zIndex"))
	return rv
}/* debug [instance_methods/method]: ZIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMCSSStyleDeclaration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cssText
func (d_ DOMCSSStyleDeclaration) CssText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cssText"))
	return rv
}/* debug [instance_properties/getter]: cssText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/cssText
func (d_ DOMCSSStyleDeclaration) SetCssText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCssText:"), value)
}/* debug [instance_properties/setter]: cssText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/length
func (d_ DOMCSSStyleDeclaration) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleDeclaration/parentRule
func (d_ DOMCSSStyleDeclaration) ParentRule() IDOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("parentRule"))
	return rv
}/* debug [instance_properties/getter]: parentRule */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMCSSStyleDeclaration */



