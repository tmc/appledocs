// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

// PBluetoothDeviceAsyncCallbacks is the IOBluetoothDeviceAsyncCallbacks protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks
type PBluetoothDeviceAsyncCallbacks interface {
	// Required methods
	ConnectionCompleteStatus(device IOBluetoothDevice, status int)/* debug [protocol_interface/required_method]: ConnectionCompleteStatus */
	RemoteNameRequestCompleteStatus(device IOBluetoothDevice, status int)/* debug [protocol_interface/required_method]: RemoteNameRequestCompleteStatus */
	SdpQueryCompleteStatus(device IOBluetoothDevice, status int)/* debug [protocol_interface/required_method]: SdpQueryCompleteStatus */
}
