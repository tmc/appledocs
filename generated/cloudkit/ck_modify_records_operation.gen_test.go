// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKModifyRecordsOperation

// ExampleNewCKModifyRecordsOperation demonstrates how to create a CKModifyRecordsOperation instance.
// Creates an empty modify records operation.
func ExampleNewCKModifyRecordsOperation() {
	_ = cloudkit.NewCKModifyRecordsOperation()
	// Output:
}
// ExampleNewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete demonstrates how to create a CKModifyRecordsOperation instance using NewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete.
// Creates an operation for modifying the specified records.
func ExampleNewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete() {
	_ = cloudkit.NewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete(
		[]cloudkit.ICKRecord{}, // records []ICKRecord
		[]cloudkit.ICKRecordID{}, // recordIDs []ICKRecordID
	)
	// Output:
}
