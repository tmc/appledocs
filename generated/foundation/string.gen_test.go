// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewStringWithBytesLengthEncoding demonstrates how to create a String instance using NewStringWithBytesLengthEncoding.
// Returns an initialized   object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
func ExampleNewStringWithBytesLengthEncoding() {
	_ = foundation.NewStringWithBytesLengthEncoding(
		nil, // bytes unsafe.Pointer
		0, // len uint
		nil, // encoding unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithBytesNoCopyLengthEncodingFreeWhenDone demonstrates how to create a String instance using NewStringWithBytesNoCopyLengthEncodingFreeWhenDone.
// Returns an initialized   object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
func ExampleNewStringWithBytesNoCopyLengthEncodingFreeWhenDone() {
	_ = foundation.NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(
		nil, // bytes unsafe.Pointer
		0, // len uint
		nil, // encoding unsafe.Pointer
		false, // freeBuffer bool
	)
	// Output:
}

// ExampleNewStringWithContentsOfURL demonstrates how to create a String instance using NewStringWithContentsOfURL.
// Initializes the receiver, a newly allocated   object, by reading data from the location named by a given URL.
func ExampleNewStringWithContentsOfURL() {
	_ = foundation.NewStringWithContentsOfURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithContentsOfURLUsedEncodingError demonstrates how to create a String instance using NewStringWithContentsOfURLUsedEncodingError.
// Returns an   object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data.
func ExampleNewStringWithContentsOfURLUsedEncodingError() {
	_ = foundation.NewStringWithContentsOfURLUsedEncodingError(
		nil, // url unsafe.Pointer
		nil, // enc unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithCStringEncoding demonstrates how to create a String instance using NewStringWithCStringEncoding.
// Returns an   object initialized using the characters in a given C array, interpreted according to a given encoding.
func ExampleNewStringWithCStringEncoding() {
	_ = foundation.NewStringWithCStringEncoding(
		nil, // nullTerminatedCString unsafe.Pointer
		nil, // encoding unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithUTF8String demonstrates how to create a String instance using NewStringWithUTF8String.
// Returns an   object initialized by copying the characters from a given C array of UTF8-encoded bytes.
func ExampleNewStringWithUTF8String() {
	_ = foundation.NewStringWithUTF8String(
		nil, // nullTerminatedCString unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithCoder demonstrates how to create a String instance using NewStringWithCoder.
func ExampleNewStringWithCoder() {
	_ = foundation.NewStringWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithContentsOfFile demonstrates how to create a String instance using NewStringWithContentsOfFile.
// Initializes the receiver, a newly allocated   object, by reading data from the file named by  .
func ExampleNewStringWithContentsOfFile() {
	_ = foundation.NewStringWithContentsOfFile(
		"path", // path string
	)
	// Output:
}

// ExampleNewStringWithValidatedFormatValidFormatSpecifiersArgumentsError demonstrates how to create a String instance using NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError.
func ExampleNewStringWithValidatedFormatValidFormatSpecifiersArgumentsError() {
	_ = foundation.NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(
		"format", // format string
		"validFormatSpecifiers", // validFormatSpecifiers string
		nil, // argList unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError demonstrates how to create a String instance using NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError.
func ExampleNewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError() {
	_ = foundation.NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(
		"format", // format string
		"validFormatSpecifiers", // validFormatSpecifiers string
		0, // locale objc.ID
		nil, // argList unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithCString demonstrates how to create a String instance using NewStringWithCString.
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
func ExampleNewStringWithCString() {
	_ = foundation.NewStringWithCString(
		nil, // bytes unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithCStringNoCopyLengthFreeWhenDone demonstrates how to create a String instance using NewStringWithCStringNoCopyLengthFreeWhenDone.
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
func ExampleNewStringWithCStringNoCopyLengthFreeWhenDone() {
	_ = foundation.NewStringWithCStringNoCopyLengthFreeWhenDone(
		nil, // bytes unsafe.Pointer
		0, // length uint
		false, // freeBuffer bool
	)
	// Output:
}

// ExampleNewStringWithContentsOfFileEncodingError demonstrates how to create a String instance using NewStringWithContentsOfFileEncodingError.
// Returns an   object initialized by reading data from the file at a given path using a given encoding.
func ExampleNewStringWithContentsOfFileEncodingError() {
	_ = foundation.NewStringWithContentsOfFileEncodingError(
		"path", // path string
		nil, // enc unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithContentsOfFileUsedEncodingError demonstrates how to create a String instance using NewStringWithContentsOfFileUsedEncodingError.
// Returns an   object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
func ExampleNewStringWithContentsOfFileUsedEncodingError() {
	_ = foundation.NewStringWithContentsOfFileUsedEncodingError(
		"path", // path string
		nil, // enc unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithDataEncoding demonstrates how to create a String instance using NewStringWithDataEncoding.
// Returns an   object initialized by converting given data into UTF-16 code units using a given encoding.
func ExampleNewStringWithDataEncoding() {
	_ = foundation.NewStringWithDataEncoding(
		nil, // data unsafe.Pointer
		nil, // encoding unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithFormatArguments demonstrates how to create a String instance using NewStringWithFormatArguments.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
func ExampleNewStringWithFormatArguments() {
	_ = foundation.NewStringWithFormatArguments(
		"format", // format string
		nil, // argList unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithString demonstrates how to create a String instance using NewStringWithString.
// Returns an   object initialized by copying the characters from another given string.
func ExampleNewStringWithString() {
	_ = foundation.NewStringWithString(
		"aString", // aString string
	)
	// Output:
}

// ExampleNewStringWithValidatedFormatValidFormatSpecifiersLocaleError demonstrates how to create a String instance using NewStringWithValidatedFormatValidFormatSpecifiersLocaleError.
func ExampleNewStringWithValidatedFormatValidFormatSpecifiersLocaleError() {
	_ = foundation.NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(
		"format", // format string
		"validFormatSpecifiers", // validFormatSpecifiers string
		0, // locale objc.ID
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithFormatLocale demonstrates how to create a String instance using NewStringWithFormatLocale.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale.
func ExampleNewStringWithFormatLocale() {
	_ = foundation.NewStringWithFormatLocale(
		"format", // format string
		0, // locale objc.ID
	)
	// Output:
}

// ExampleNewStringWithCharactersLength demonstrates how to create a String instance using NewStringWithCharactersLength.
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units.
func ExampleNewStringWithCharactersLength() {
	_ = foundation.NewStringWithCharactersLength(
		nil, // characters unsafe.Pointer
		0, // length uint
	)
	// Output:
}

// ExampleNewStringWithCStringLength demonstrates how to create a String instance using NewStringWithCStringLength.
// Initializes the receiver, a newly allocated   object, by converting the data in a given C-string from the default C-string encoding into the Unicode character encoding.
func ExampleNewStringWithCStringLength() {
	_ = foundation.NewStringWithCStringLength(
		nil, // bytes unsafe.Pointer
		0, // length uint
	)
	// Output:
}

// ExampleNewStringWithFormatLocaleArguments demonstrates how to create a String instance using NewStringWithFormatLocaleArguments.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
func ExampleNewStringWithFormatLocaleArguments() {
	_ = foundation.NewStringWithFormatLocaleArguments(
		"format", // format string
		0, // locale objc.ID
		nil, // argList unsafe.Pointer
	)
	// Output:
}

// ExampleNewString demonstrates how to create a String instance.
// Returns an initialized   object that contains no characters.
func ExampleNewString() {
	_ = foundation.NewString()
	// Output:
}

// ExampleNewStringWithCharactersNoCopyLengthDeallocator demonstrates how to create a String instance using NewStringWithCharactersNoCopyLengthDeallocator.
func ExampleNewStringWithCharactersNoCopyLengthDeallocator() {
	_ = foundation.NewStringWithCharactersNoCopyLengthDeallocator(
		nil, // chars unsafe.Pointer
		0, // len uint
		nil, // deallocator unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithContentsOfURLEncodingError demonstrates how to create a String instance using NewStringWithContentsOfURLEncodingError.
// Returns an   object initialized by reading data from a given URL interpreted using a given encoding.
func ExampleNewStringWithContentsOfURLEncodingError() {
	_ = foundation.NewStringWithContentsOfURLEncodingError(
		nil, // url unsafe.Pointer
		nil, // enc unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithFormat demonstrates how to create a String instance using NewStringWithFormat.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted.
func ExampleNewStringWithFormat() {
	_ = foundation.NewStringWithFormat(
		"format", // format string
	)
	// Output:
}

// ExampleNewStringWithBytesNoCopyLengthEncodingDeallocator demonstrates how to create a String instance using NewStringWithBytesNoCopyLengthEncodingDeallocator.
func ExampleNewStringWithBytesNoCopyLengthEncodingDeallocator() {
	_ = foundation.NewStringWithBytesNoCopyLengthEncodingDeallocator(
		nil, // bytes unsafe.Pointer
		0, // len uint
		nil, // encoding unsafe.Pointer
		nil, // deallocator unsafe.Pointer
	)
	// Output:
}

// ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone demonstrates how to create a String instance using NewStringWithCharactersNoCopyLengthFreeWhenDone.
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units.
func ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone() {
	_ = foundation.NewStringWithCharactersNoCopyLengthFreeWhenDone(
		nil, // characters unsafe.Pointer
		0, // length uint
		false, // freeBuffer bool
	)
	// Output:
}

// ExampleNewStringWithValidatedFormatValidFormatSpecifiersError demonstrates how to create a String instance using NewStringWithValidatedFormatValidFormatSpecifiersError.
func ExampleNewStringWithValidatedFormatValidFormatSpecifiersError() {
	_ = foundation.NewStringWithValidatedFormatValidFormatSpecifiersError(
		"format", // format string
		"validFormatSpecifiers", // validFormatSpecifiers string
		nil, // error unsafe.Pointer
	)
	// Output:
}


