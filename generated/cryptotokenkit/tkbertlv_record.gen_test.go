// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit_test

import (
	"github.com/tmc/appledocs/generated/cryptotokenkit"
)

// Suppress unused import errors
var _ = cryptotokenkit.NewTKBERTLVRecord

// ExampleNewTKBERTLVRecordWithTagRecords demonstrates how to create a TKBERTLVRecord instance using NewTKBERTLVRecordWithTagRecords.
// Initializes a BER-TLV record with the specified tag and an array of TLV subrecords.
func ExampleNewTKBERTLVRecordWithTagRecords() {
	_ = cryptotokenkit.NewTKBERTLVRecordWithTagRecords(
		cryptotokenkit.TKTLVTag /* typedef */{}, // tag TKTLVTag /* typedef */
		[]cryptotokenkit.TKTLVRecord{}, // records []TKTLVRecord
	)
	// Output:
}
