// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewAttributedString

// ExampleNewAttributedStringWithAdaptiveImageGlyphAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithAdaptiveImageGlyphAttributes.
// Creates an attributed string with an adaptive image glyph and applies the specified attributes to it.
func ExampleNewAttributedStringWithAdaptiveImageGlyphAttributes() {
	_ = foundation.NewAttributedStringWithAdaptiveImageGlyphAttributes(
		foundation.AdaptiveImageGlyph{}, // adaptiveImageGlyph AdaptiveImageGlyph
		foundation.IDictionary{}, // attributes IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithAttachment demonstrates how to create a AttributedString instance using NewAttributedStringWithAttachment.
// Creates an attributed string with an attachment.
func ExampleNewAttributedStringWithAttachment() {
	_ = foundation.NewAttributedStringWithAttachment(
		foundation.TextAttachment{}, // attachment TextAttachment
	)
	// Output:
}
// ExampleNewAttributedStringWithAttachmentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithAttachmentAttributes.
// Creates an attributed string with an attachment and applies the specified attributes to it.
func ExampleNewAttributedStringWithAttachmentAttributes() {
	_ = foundation.NewAttributedStringWithAttachmentAttributes(
		foundation.TextAttachment{}, // attachment TextAttachment
		foundation.IDictionary{}, // attributes IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithAttributedString demonstrates how to create a AttributedString instance using NewAttributedStringWithAttributedString.
// Creates a new attributed string from the contents of another attributed string.
func ExampleNewAttributedStringWithAttributedString() {
	_ = foundation.NewAttributedStringWithAttributedString(
		foundation.NSAttributedString{}, // attrStr NSAttributedString
	)
	// Output:
}
// ExampleNewAttributedStringWithDocFormatDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithDocFormatDocumentAttributes.
// Creates an attributed string from Microsoft Word format data in the specified data object.
func ExampleNewAttributedStringWithDocFormatDocumentAttributes() {
	_ = foundation.NewAttributedStringWithDocFormatDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithHTMLBaseURLDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithHTMLBaseURLDocumentAttributes.
// Creates an attributed string from the HTML in the specified data object and base URL.
func ExampleNewAttributedStringWithHTMLBaseURLDocumentAttributes() {
	_ = foundation.NewAttributedStringWithHTMLBaseURLDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.URL{}, // base URL
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithHTMLDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithHTMLDocumentAttributes.
// Creates an attributed string from the HTML in the specified data object.
func ExampleNewAttributedStringWithHTMLDocumentAttributes() {
	_ = foundation.NewAttributedStringWithHTMLDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithHTMLOptionsDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithHTMLOptionsDocumentAttributes.
// Creates an attributed string from the HTML in the specified data object.
func ExampleNewAttributedStringWithHTMLOptionsDocumentAttributes() {
	_ = foundation.NewAttributedStringWithHTMLOptionsDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.IDictionary{}, // options IDictionary
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithRTFDDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithRTFDDocumentAttributes.
// Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object.
func ExampleNewAttributedStringWithRTFDDocumentAttributes() {
	_ = foundation.NewAttributedStringWithRTFDDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithRTFDFileWrapperDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithRTFDFileWrapperDocumentAttributes.
// Creates an attributed string from the specified file wrapper that contains an RTFD document.
func ExampleNewAttributedStringWithRTFDFileWrapperDocumentAttributes() {
	_ = foundation.NewAttributedStringWithRTFDFileWrapperDocumentAttributes(
		foundation.FileWrapper{}, // wrapper FileWrapper
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithRTFDocumentAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithRTFDocumentAttributes.
// Creates an attributed string by decoding the stream of RTF commands and data in the specified data object.
func ExampleNewAttributedStringWithRTFDocumentAttributes() {
	_ = foundation.NewAttributedStringWithRTFDocumentAttributes(
		foundation.NSData{}, // data NSData
		foundation.IDictionary{}, // dict IDictionary
	)
	// Output:
}
// ExampleNewAttributedStringWithString demonstrates how to create a AttributedString instance using NewAttributedStringWithString.
// Creates an attributed string with the specified text and no attribute information.
func ExampleNewAttributedStringWithString() {
	_ = foundation.NewAttributedStringWithString(
		"str", // str string
	)
	// Output:
}
// ExampleNewAttributedStringWithStringAttributes demonstrates how to create a AttributedString instance using NewAttributedStringWithStringAttributes.
// Creates an attributed string with the specified text and attributes.
func ExampleNewAttributedStringWithStringAttributes() {
	_ = foundation.NewAttributedStringWithStringAttributes(
		"str", // str string
		foundation.IDictionary{}, // attrs IDictionary
	)
	// Output:
}
