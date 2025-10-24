// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewOBEXFileTransferServices

// ExampleOBEXFileTransferServices_Abort demonstrates using Abort on a OBEXFileTransferServices instance.
// Abort the current operation
func ExampleOBEXFileTransferServices_Abort() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.Abort()
	// Output:
	}

// ExampleOBEXFileTransferServices_ChangeCurrentFolderBackward demonstrates using ChangeCurrentFolderBackward on a OBEXFileTransferServices instance.
// Change to the directory above the current level if not at the root
func ExampleOBEXFileTransferServices_ChangeCurrentFolderBackward() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.ChangeCurrentFolderBackward()
	// Output:
	}

// ExampleOBEXFileTransferServices_ChangeCurrentFolderToRoot demonstrates using ChangeCurrentFolderToRoot on a OBEXFileTransferServices instance.
// Asynchronously change to the remote root directory
func ExampleOBEXFileTransferServices_ChangeCurrentFolderToRoot() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.ChangeCurrentFolderToRoot()
	// Output:
	}

// ExampleOBEXFileTransferServices_ConnectToFTPService demonstrates using ConnectToFTPService on a OBEXFileTransferServices instance.
// Connect to a remote device for FTP operations
func ExampleOBEXFileTransferServices_ConnectToFTPService() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.ConnectToFTPService()
	// Output:
	}

// ExampleOBEXFileTransferServices_ConnectToObjectPushService demonstrates using ConnectToObjectPushService on a OBEXFileTransferServices instance.
// Connect to a remote device for ObjectPush operations. Most of the FTP functionality of this object will be disabled.
func ExampleOBEXFileTransferServices_ConnectToObjectPushService() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.ConnectToObjectPushService()
	// Output:
	}

// ExampleOBEXFileTransferServices_CurrentPath demonstrates using CurrentPath on a OBEXFileTransferServices instance.
// Get the remote current directory path during an FTP session
func ExampleOBEXFileTransferServices_CurrentPath() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.CurrentPath()
	// Output:
	}

// ExampleOBEXFileTransferServices_Disconnect demonstrates using Disconnect on a OBEXFileTransferServices instance.
// Disconnect from the remote device
func ExampleOBEXFileTransferServices_Disconnect() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.Disconnect()
	// Output:
	}

// ExampleOBEXFileTransferServices_IsBusy demonstrates using IsBusy on a OBEXFileTransferServices instance.
// Get the action state of the module
func ExampleOBEXFileTransferServices_IsBusy() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.IsBusy()
	// Output:
	}

// ExampleOBEXFileTransferServices_IsConnected demonstrates using IsConnected on a OBEXFileTransferServices instance.
// Get the connected state of this module.
func ExampleOBEXFileTransferServices_IsConnected() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.IsConnected()
	// Output:
	}

// ExampleOBEXFileTransferServices_RetrieveFolderListing demonstrates using RetrieveFolderListing on a OBEXFileTransferServices instance.
// Get a remote directory listing
func ExampleOBEXFileTransferServices_RetrieveFolderListing() {
	obj := iobluetooth.NewOBEXFileTransferServices()
	_ = obj.RetrieveFolderListing()
	// Output:
	}

