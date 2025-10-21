// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

// Enum types and constants
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

// PDFActionNamedName enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName
type PDFActionNamedName uint

// PDFBorderStyle - PDF Kit annotation borders may have the following styles.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle
type PDFBorderStyle uint

// PDFDisplayBox - The following box types may be used with 
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox
type PDFDisplayBox uint

// PDFDisplayDirection enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDisplayDirection
type PDFDisplayDirection uint

// PDFDocumentPermissions - An enumeration that specifies document permissions status.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions
type PDFDocumentPermissions uint

// PDFMarkupType - The styles available for markup annotations in PDFKit.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFMarkupType
type PDFMarkupType uint

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

// PDFTextAnnotationIconType - The types of icons that a text annotation can use.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType
type PDFTextAnnotationIconType uint

// PDFThumbnailLayoutMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailLayoutMode
type PDFThumbnailLayoutMode uint

// PDFWidgetCellState - The state of a button annotation, either on, off, or mixed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState
type PDFWidgetCellState uint

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
)


