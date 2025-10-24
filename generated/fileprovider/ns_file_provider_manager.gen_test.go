// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider_test

import (
	"github.com/tmc/appledocs/generated/fileprovider"
)

// Suppress unused import errors
var _ = fileprovider.NewFileProviderManager

// ExampleFileProviderManager_EnumeratorForMaterializedItems demonstrates using EnumeratorForMaterializedItems on a FileProviderManager instance.
// Returns an enumerator for all the items the system currently stores on disk.
func ExampleFileProviderManager_EnumeratorForMaterializedItems() {
	obj := fileprovider.NewFileProviderManager()
	_ = obj.EnumeratorForMaterializedItems()
	// Output:
	}

// ExampleFileProviderManager_EnumeratorForPendingItems demonstrates using EnumeratorForPendingItems on a FileProviderManager instance.
// Returns an enumerator for the set of pending items.
func ExampleFileProviderManager_EnumeratorForPendingItems() {
	obj := fileprovider.NewFileProviderManager()
	_ = obj.EnumeratorForPendingItems()
	// Output:
	}

