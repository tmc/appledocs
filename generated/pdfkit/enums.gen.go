// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

/* debug [enums.gen.go]: Generating 17 enums for PDFKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum PDFAccessPermissions (8 cases) */
// PDFAccessPermissions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions
type PDFAccessPermissions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsCommenting
	PDFAllowsCommenting PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsContentAccessibility
	PDFAllowsContentAccessibility PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsContentCopying
	PDFAllowsContentCopying PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsDocumentAssembly
	PDFAllowsDocumentAssembly PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsDocumentChanges
	PDFAllowsDocumentChanges PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsFormFieldEntry
	PDFAllowsFormFieldEntry PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsHighQualityPrinting
	PDFAllowsHighQualityPrinting PDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions/allowsLowQualityPrinting
	PDFAllowsLowQualityPrinting PDFAccessPermissions = 0
)

/* debug [enums.gen.go]: Processing enum PDFActionNamedName (12 cases) */
// PDFActionNamedName enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName
type PDFActionNamedName uint

const (
	// kPDFActionNamedFind - The Find action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/find
	kPDFActionNamedFind PDFActionNamedName = 0
	// kPDFActionNamedFirstPage - The First Page action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/firstPage
	kPDFActionNamedFirstPage PDFActionNamedName = 0
	// kPDFActionNamedGoBack - The Go Back action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/goBack
	kPDFActionNamedGoBack PDFActionNamedName = 0
	// kPDFActionNamedGoForward - The Go Forward action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/goForward
	kPDFActionNamedGoForward PDFActionNamedName = 0
	// kPDFActionNamedGoToPage - The Go to Page action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/goToPage
	kPDFActionNamedGoToPage PDFActionNamedName = 0
	// kPDFActionNamedLastPage - The Last Page action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/lastPage
	kPDFActionNamedLastPage PDFActionNamedName = 0
	// kPDFActionNamedNextPage - The Next Page action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/nextPage
	kPDFActionNamedNextPage PDFActionNamedName = 0
	// kPDFActionNamedNone - The action has no name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/none
	kPDFActionNamedNone PDFActionNamedName = 0
	// kPDFActionNamedPreviousPage - The Previous Page action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/previousPage
	kPDFActionNamedPreviousPage PDFActionNamedName = 0
	// kPDFActionNamedPrint - The Print action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/print
	kPDFActionNamedPrint PDFActionNamedName = 0
	// kPDFActionNamedZoomIn - The Zoom In action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/zoomIn
	kPDFActionNamedZoomIn PDFActionNamedName = 0
	// kPDFActionNamedZoomOut - The Zoom Out action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName/zoomOut
	kPDFActionNamedZoomOut PDFActionNamedName = 0
)

/* debug [enums.gen.go]: Processing enum PDFAreaOfInterest (11 cases) */
// PDFAreaOfInterest - The mouse position over PDF view areas.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest
type PDFAreaOfInterest uint

const (
	// kPDFAnnotationArea - The mouse is over an annotation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/annotationArea
	kPDFAnnotationArea PDFAreaOfInterest = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/anyArea
	kPDFAnyArea PDFAreaOfInterest = 0
	// kPDFControlArea - The mouse is over a control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/controlArea
	kPDFControlArea PDFAreaOfInterest = 0
	// kPDFIconArea - The mouse is over an icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/iconArea
	kPDFIconArea PDFAreaOfInterest = 0
	// kPDFImageArea - The mouse is over an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/imageArea
	kPDFImageArea PDFAreaOfInterest = 0
	// kPDFNoArea - The mouse is over an undefined area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/kPDFNoArea
	kPDFNoArea PDFAreaOfInterest = 0
	// kPDFLinkArea - The mouse is over a link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/linkArea
	kPDFLinkArea PDFAreaOfInterest = 0
	// kPDFPageArea - The mouse is over a page.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/pageArea
	kPDFPageArea PDFAreaOfInterest = 0
	// kPDFPopupArea - The mouse is over a popup menu.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/popupArea
	kPDFPopupArea PDFAreaOfInterest = 0
	// kPDFTextArea - The mouse is over text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/textArea
	kPDFTextArea PDFAreaOfInterest = 0
	// kPDFTextFieldArea - The mouse is over a text field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest/textFieldArea
	kPDFTextFieldArea PDFAreaOfInterest = 0
)

/* debug [enums.gen.go]: Processing enum PDFBorderStyle (5 cases) */
// PDFBorderStyle - PDF Kit annotation borders may have the following styles.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle
type PDFBorderStyle uint

const (
	// kPDFBorderStyleBeveled - Beveled border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle/beveled
	kPDFBorderStyleBeveled PDFBorderStyle = 0
	// kPDFBorderStyleDashed - Dashed border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle/dashed
	kPDFBorderStyleDashed PDFBorderStyle = 0
	// kPDFBorderStyleInset - Inset border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle/inset
	kPDFBorderStyleInset PDFBorderStyle = 0
	// kPDFBorderStyleSolid - Solid border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle/solid
	kPDFBorderStyleSolid PDFBorderStyle = 0
	// kPDFBorderStyleUnderline - Underline border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle/underline
	kPDFBorderStyleUnderline PDFBorderStyle = 0
)

/* debug [enums.gen.go]: Processing enum PDFDisplayBox (5 cases) */
// PDFDisplayBox - The following box types may be used with 
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox
type PDFDisplayBox uint

const (
	// kPDFDisplayBoxArtBox - A rectangle defining the boundaries of the page’s meaningful content including surrounding white space intended for display. Default value equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox/artBox
	kPDFDisplayBoxArtBox PDFDisplayBox = 0
	// kPDFDisplayBoxBleedBox - A rectangle defining the boundaries of the clip region for the page contents in a production environment. Default value equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox/bleedBox
	kPDFDisplayBoxBleedBox PDFDisplayBox = 0
	// kPDFDisplayBoxCropBox - A rectangle defining the boundaries of the visible region , expressed in default user-space units. Default value equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox/cropBox
	kPDFDisplayBoxCropBox PDFDisplayBox = 0
	// kPDFDisplayBoxMediaBox - A rectangle defining the boundaries of the physical medium for display or printing, expressed in default user-space units.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox/mediaBox
	kPDFDisplayBoxMediaBox PDFDisplayBox = 0
	// kPDFDisplayBoxTrimBox - A rectangle defining the intended boundaries of the finished page. Default value equal to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox/trimBox
	kPDFDisplayBoxTrimBox PDFDisplayBox = 0
)

/* debug [enums.gen.go]: Processing enum PDFDisplayDirection (2 cases) */
// PDFDisplayDirection enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayDirection
type PDFDisplayDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayDirection/horizontal
	kPDFDisplayDirectionHorizontal PDFDisplayDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayDirection/vertical
	kPDFDisplayDirectionVertical PDFDisplayDirection = 0
)

/* debug [enums.gen.go]: Processing enum PDFDisplayMode (4 cases) */
// PDFDisplayMode - A wrapper for the chosen display mode constant.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode
type PDFDisplayMode uint

const (
	// kPDFDisplaySinglePage - A display mode where the document displays one page at a time horizontally and vertically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode/singlePage
	kPDFDisplaySinglePage PDFDisplayMode = 0
	// kPDFDisplaySinglePageContinuous - A display mode where the document displays in continuous mode vertically, with single-page width horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode/singlePageContinuous
	kPDFDisplaySinglePageContinuous PDFDisplayMode = 0
	// kPDFDisplayTwoUp - A display mode where the document displays two pages side-by-side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode/twoUp
	kPDFDisplayTwoUp PDFDisplayMode = 0
	// kPDFDisplayTwoUpContinuous - A display mode where the document displays in continuous mode vertically and displays two pages side-by-side horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode/twoUpContinuous
	kPDFDisplayTwoUpContinuous PDFDisplayMode = 0
)

/* debug [enums.gen.go]: Processing enum PDFDocumentPermissions (3 cases) */
// PDFDocumentPermissions - An enumeration that specifies document permissions status.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions
type PDFDocumentPermissions uint

const (
	// kPDFDocumentPermissionsNone - The status that indicates no document permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions/none
	kPDFDocumentPermissionsNone PDFDocumentPermissions = 0
	// kPDFDocumentPermissionsOwner - The status that indicates owner document permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions/owner
	kPDFDocumentPermissionsOwner PDFDocumentPermissions = 0
	// kPDFDocumentPermissionsUser - The status that indicates user document permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions/user
	kPDFDocumentPermissionsUser PDFDocumentPermissions = 0
)

/* debug [enums.gen.go]: Processing enum PDFInterpolationQuality (3 cases) */
// PDFInterpolationQuality - A wrapper for the specified interpolation quality.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFInterpolationQuality
type PDFInterpolationQuality uint

const (
	// kPDFInterpolationQualityHigh - The case specifying high interpolation quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFInterpolationQuality/high
	kPDFInterpolationQualityHigh PDFInterpolationQuality = 0
	// kPDFInterpolationQualityLow - The case specifying low interpolation quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFInterpolationQuality/low
	kPDFInterpolationQualityLow PDFInterpolationQuality = 0
	// kPDFInterpolationQualityNone - The case where no interpolation quality is specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFInterpolationQuality/none
	kPDFInterpolationQualityNone PDFInterpolationQuality = 0
)

/* debug [enums.gen.go]: Processing enum PDFLineStyle (6 cases) */
// PDFLineStyle - The following constants specify the available line ending styles.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle
type PDFLineStyle uint

const (
	// kPDFLineStyleCircle - A circular line ending filled with the annotation’s interior color, if any.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/circle
	kPDFLineStyleCircle PDFLineStyle = 0
	// kPDFLineStyleClosedArrow - A closed arrowhead line ending, consisting of a triangle with the acute vertex at the line end and filled with the annotation’s interior color, if any.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/closedArrow
	kPDFLineStyleClosedArrow PDFLineStyle = 0
	// kPDFLineStyleDiamond - A diamond-shaped line ending filled with the annotation’s interior color, if any.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/diamond
	kPDFLineStyleDiamond PDFLineStyle = 0
	// kPDFLineStyleNone - No line ending.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/none
	kPDFLineStyleNone PDFLineStyle = 0
	// kPDFLineStyleOpenArrow - An open arrowhead line ending, composed from two short lines meeting in an acute angle at the line end.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/openArrow
	kPDFLineStyleOpenArrow PDFLineStyle = 0
	// kPDFLineStyleSquare - A square line ending filled with the annotation’s interior color, if any.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFLineStyle/square
	kPDFLineStyleSquare PDFLineStyle = 0
)

/* debug [enums.gen.go]: Processing enum PDFMarkupType (4 cases) */
// PDFMarkupType - The styles available for markup annotations in PDFKit.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType
type PDFMarkupType uint

const (
	// kPDFMarkupTypeHighlight - Highlight style for the markup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType/highlight
	kPDFMarkupTypeHighlight PDFMarkupType = 0
	// kPDFMarkupTypeRedact - The redaction style for markup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType/redact
	kPDFMarkupTypeRedact PDFMarkupType = 0
	// kPDFMarkupTypeStrikeOut - Strikethrough style for the markup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType/strikeOut
	kPDFMarkupTypeStrikeOut PDFMarkupType = 0
	// kPDFMarkupTypeUnderline - Underline style for the markup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType/underline
	kPDFMarkupTypeUnderline PDFMarkupType = 0
)

/* debug [enums.gen.go]: Processing enum PDFPrintScalingMode (3 cases) */
// PDFPrintScalingMode - The type of scaling to be used when printing a page (see 
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPrintScalingMode
type PDFPrintScalingMode uint

const (
	// kPDFPrintPageScaleDownToFit - Scale large pages down to fit the paper size (smaller pages do not get scaled up).
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPrintScalingMode/pageScaleDownToFit
	kPDFPrintPageScaleDownToFit PDFPrintScalingMode = 0
	// kPDFPrintPageScaleNone - Do not apply scaling to the page when printing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPrintScalingMode/pageScaleNone
	kPDFPrintPageScaleNone PDFPrintScalingMode = 0
	// kPDFPrintPageScaleToFit - Scale each page up or down to best fit the paper size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPrintScalingMode/pageScaleToFit
	kPDFPrintPageScaleToFit PDFPrintScalingMode = 0
)

/* debug [enums.gen.go]: Processing enum PDFSelectionGranularity (3 cases) */
// PDFSelectionGranularity enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelectionGranularity
type PDFSelectionGranularity uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelectionGranularity/character
	PDFSelectionGranularityCharacter PDFSelectionGranularity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelectionGranularity/line
	PDFSelectionGranularityLine PDFSelectionGranularity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelectionGranularity/word
	PDFSelectionGranularityWord PDFSelectionGranularity = 0
)

/* debug [enums.gen.go]: Processing enum PDFTextAnnotationIconType (7 cases) */
// PDFTextAnnotationIconType - The types of icons that a text annotation can use.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType
type PDFTextAnnotationIconType uint

const (
	// kPDFTextAnnotationIconComment - Comment annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/comment
	kPDFTextAnnotationIconComment PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconHelp - Help annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/help
	kPDFTextAnnotationIconHelp PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconInsert - Insert annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/insert
	kPDFTextAnnotationIconInsert PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconKey - Key annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/key
	kPDFTextAnnotationIconKey PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconNewParagraph - New Paragraph annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/newParagraph
	kPDFTextAnnotationIconNewParagraph PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconNote - Note annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/note
	kPDFTextAnnotationIconNote PDFTextAnnotationIconType = 0
	// kPDFTextAnnotationIconParagraph - Paragraph annotation icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType/paragraph
	kPDFTextAnnotationIconParagraph PDFTextAnnotationIconType = 0
)

/* debug [enums.gen.go]: Processing enum PDFThumbnailLayoutMode (2 cases) */
// PDFThumbnailLayoutMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailLayoutMode
type PDFThumbnailLayoutMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailLayoutMode/horizontal
	PDFThumbnailLayoutModeHorizontal PDFThumbnailLayoutMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailLayoutMode/vertical
	PDFThumbnailLayoutModeVertical PDFThumbnailLayoutMode = 0
)

/* debug [enums.gen.go]: Processing enum PDFWidgetCellState (3 cases) */
// PDFWidgetCellState - The state of a button annotation, either on, off, or mixed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState
type PDFWidgetCellState uint

const (
	// kPDFWidgetMixedState - The button widget is in a mixed state, neither on nor off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState/mixedState
	kPDFWidgetMixedState PDFWidgetCellState = 0
	// kPDFWidgetOffState - The button widget is in an unselected state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState/offState
	kPDFWidgetOffState PDFWidgetCellState = 0
	// kPDFWidgetOnState - The button widget is in a selected state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState/onState
	kPDFWidgetOnState PDFWidgetCellState = 0
)

/* debug [enums.gen.go]: Processing enum PDFWidgetControlType (4 cases) */
// PDFWidgetControlType - The types of annotation buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType
type PDFWidgetControlType uint

const (
	// kPDFWidgetCheckBoxControl - Check box control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType/checkBoxControl
	kPDFWidgetCheckBoxControl PDFWidgetControlType = 0
	// kPDFWidgetPushButtonControl - Push button control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType/pushButtonControl
	kPDFWidgetPushButtonControl PDFWidgetControlType = 0
	// kPDFWidgetRadioButtonControl - Radio button control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType/radioButtonControl
	kPDFWidgetRadioButtonControl PDFWidgetControlType = 0
	// kPDFWidgetUnknownControl - Unknown control type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType/unknownControl
	kPDFWidgetUnknownControl PDFWidgetControlType = 0
)


