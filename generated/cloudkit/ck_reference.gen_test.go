// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKReference

// ExampleNewCKReferenceWithRecordAction demonstrates how to create a CKReference instance using NewCKReferenceWithRecordAction.
// Creates a reference object that points to the specified record object.
func ExampleNewCKReferenceWithRecordAction() {
	_ = cloudkit.NewCKReferenceWithRecordAction(
		cloudkit.CKRecord{}, // record CKRecord
		cloudkit.CKReferenceAction{}, // action CKReferenceAction
	)
	// Output:
}
// ExampleNewCKReferenceWithRecordIDAction demonstrates how to create a CKReference instance using NewCKReferenceWithRecordIDAction.
// Creates a reference object that points to the record with the specified ID.
func ExampleNewCKReferenceWithRecordIDAction() {
	_ = cloudkit.NewCKReferenceWithRecordIDAction(
		cloudkit.CKRecordID{}, // recordID CKRecordID
		cloudkit.CKReferenceAction{}, // action CKReferenceAction
	)
	// Output:
}
