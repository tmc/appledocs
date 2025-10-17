// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewDictionary demonstrates how to create a Dictionary instance.
// Initializes a newly allocated dictionary.
func ExampleNewDictionary() {
	_ = foundation.NewDictionary()
	// Output:
}

// ExampleNewDictionaryWithObjectForKey demonstrates how to create a Dictionary instance using NewDictionaryWithObjectForKey.
// Creates a dictionary containing a given key and value.
func ExampleNewDictionaryWithObjectForKey() {
	_ = foundation.NewDictionaryWithObjectForKey(
		nil, // object unsafe.Pointer
		nil, // key unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithObjectsForKeys demonstrates how to create a Dictionary instance using NewDictionaryWithObjectsForKeys.
// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
func ExampleNewDictionaryWithObjectsForKeys() {
	_ = foundation.NewDictionaryWithObjectsForKeys(
		nil, // objects unsafe.Pointer
		nil, // keys unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithCoder demonstrates how to create a Dictionary instance using NewDictionaryWithCoder.
// Creates a dictionary initialized from data in the provided unarchiver.
func ExampleNewDictionaryWithCoder() {
	_ = foundation.NewDictionaryWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithContentsOfFile demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfFile.
// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
func ExampleNewDictionaryWithContentsOfFile() {
	_ = foundation.NewDictionaryWithContentsOfFile(
		"path", // path string
	)
	// Output:
}

// ExampleNewDictionaryWithContentsOfURL demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfURL.
// Initializes a newly allocated dictionary using the keys and values found at a given URL.
func ExampleNewDictionaryWithContentsOfURL() {
	_ = foundation.NewDictionaryWithContentsOfURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithContentsOfURLError demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfURLError.
// Initializes a newly allocated dictionary using the keys and values found at a given URL.
func ExampleNewDictionaryWithContentsOfURLError() {
	_ = foundation.NewDictionaryWithContentsOfURLError(
		nil, // url unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithDictionary demonstrates how to create a Dictionary instance using NewDictionaryWithDictionary.
// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
func ExampleNewDictionaryWithDictionary() {
	_ = foundation.NewDictionaryWithDictionary(
		nil, // otherDictionary unsafe.Pointer
	)
	// Output:
}

// ExampleNewDictionaryWithDictionaryCopyItems demonstrates how to create a Dictionary instance using NewDictionaryWithDictionaryCopyItems.
// Initializes a newly allocated dictionary using the objects contained in another given dictionary.
func ExampleNewDictionaryWithDictionaryCopyItems() {
	_ = foundation.NewDictionaryWithDictionaryCopyItems(
		nil, // otherDictionary unsafe.Pointer
		false, // flag bool
	)
	// Output:
}

// ExampleNewDictionaryWithObjectsForKeysCount demonstrates how to create a Dictionary instance using NewDictionaryWithObjectsForKeysCount.
// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
func ExampleNewDictionaryWithObjectsForKeysCount() {
	_ = foundation.NewDictionaryWithObjectsForKeysCount(
		nil, // objects unsafe.Pointer
		nil, // keys unsafe.Pointer
		0, // cnt uint
	)
	// Output:
}

// ExampleNewDictionaryWithObjectsAndKeys demonstrates how to create a Dictionary instance using NewDictionaryWithObjectsAndKeys.
// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys.
func ExampleNewDictionaryWithObjectsAndKeys() {
	_ = foundation.NewDictionaryWithObjectsAndKeys(
		0, // firstObject objc.ID
	)
	// Output:
}


