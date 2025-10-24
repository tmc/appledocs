// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan_test

import (
	"github.com/tmc/appledocs/generated/corewlan"
)

// Suppress unused import errors
var _ = corewlan.NewCWInterface

// ExampleCWInterface_ActivePHYMode demonstrates using ActivePHYMode on a CWInterface instance.
// The current active PHY modes for the interface.
func ExampleCWInterface_ActivePHYMode() {
	obj := corewlan.NewCWInterface()
	_ = obj.ActivePHYMode()
	// Output:
	}

// ExampleCWInterface_Bssid demonstrates using Bssid on a CWInterface instance.
// The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
func ExampleCWInterface_Bssid() {
	obj := corewlan.NewCWInterface()
	_ = obj.Bssid()
	// Output:
	}

// ExampleCWInterface_CachedScanResults demonstrates using CachedScanResults on a CWInterface instance.
// The networks currently in the scan cache for the WLAN interface.
func ExampleCWInterface_CachedScanResults() {
	obj := corewlan.NewCWInterface()
	_ = obj.CachedScanResults()
	// Output:
	}

// ExampleCWInterface_Configuration demonstrates using Configuration on a CWInterface instance.
// The current configuration for the given WLAN interface.
func ExampleCWInterface_Configuration() {
	obj := corewlan.NewCWInterface()
	_ = obj.Configuration()
	// Output:
	}

// ExampleCWInterface_CountryCode demonstrates using CountryCode on a CWInterface instance.
// The current country code (ISO/IEC 3166-1:1997) for the interface.
func ExampleCWInterface_CountryCode() {
	obj := corewlan.NewCWInterface()
	_ = obj.CountryCode()
	// Output:
	}

// ExampleCWInterface_Disassociate demonstrates using Disassociate on a CWInterface instance.
// Disassociates from the current network.
func ExampleCWInterface_Disassociate() {
	obj := corewlan.NewCWInterface()
	obj.Disassociate()
	// Output:
	}

// ExampleCWInterface_HardwareAddress demonstrates using HardwareAddress on a CWInterface instance.
// The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
func ExampleCWInterface_HardwareAddress() {
	obj := corewlan.NewCWInterface()
	_ = obj.HardwareAddress()
	// Output:
	}

// ExampleCWInterface_InterfaceMode demonstrates using InterfaceMode on a CWInterface instance.
// The current mode for the interface.
func ExampleCWInterface_InterfaceMode() {
	obj := corewlan.NewCWInterface()
	_ = obj.InterfaceMode()
	// Output:
	}

// ExampleCWInterface_NoiseMeasurement demonstrates using NoiseMeasurement on a CWInterface instance.
// The current aggregate noise measurement (dBm) for the interface.
func ExampleCWInterface_NoiseMeasurement() {
	obj := corewlan.NewCWInterface()
	_ = obj.NoiseMeasurement()
	// Output:
	}

// ExampleCWInterface_PowerOn demonstrates using PowerOn on a CWInterface instance.
// The interface power state is set to “ON”.
func ExampleCWInterface_PowerOn() {
	obj := corewlan.NewCWInterface()
	_ = obj.PowerOn()
	// Output:
	}

// ExampleCWInterface_RssiValue demonstrates using RssiValue on a CWInterface instance.
// The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
func ExampleCWInterface_RssiValue() {
	obj := corewlan.NewCWInterface()
	_ = obj.RssiValue()
	// Output:
	}

// ExampleCWInterface_Security demonstrates using Security on a CWInterface instance.
// The current security mode for the interface.
func ExampleCWInterface_Security() {
	obj := corewlan.NewCWInterface()
	_ = obj.Security()
	// Output:
	}

// ExampleCWInterface_ServiceActive demonstrates using ServiceActive on a CWInterface instance.
// The interface has its corresponding network service enabled.
func ExampleCWInterface_ServiceActive() {
	obj := corewlan.NewCWInterface()
	_ = obj.ServiceActive()
	// Output:
	}

// ExampleCWInterface_Ssid demonstrates using Ssid on a CWInterface instance.
// The current service set identifier (SSID) for the interface, encoded as a string.
func ExampleCWInterface_Ssid() {
	obj := corewlan.NewCWInterface()
	_ = obj.Ssid()
	// Output:
	}

// ExampleCWInterface_SsidData demonstrates using SsidData on a CWInterface instance.
// The current service set identifier (SSID) for the interface, returned as data.
func ExampleCWInterface_SsidData() {
	obj := corewlan.NewCWInterface()
	_ = obj.SsidData()
	// Output:
	}

// ExampleCWInterface_SupportedWLANChannels demonstrates using SupportedWLANChannels on a CWInterface instance.
// An array of channels supported by the interface for the active country code.
func ExampleCWInterface_SupportedWLANChannels() {
	obj := corewlan.NewCWInterface()
	_ = obj.SupportedWLANChannels()
	// Output:
	}

// ExampleCWInterface_TransmitPower demonstrates using TransmitPower on a CWInterface instance.
// The current transmit power (mW) for the interface.
func ExampleCWInterface_TransmitPower() {
	obj := corewlan.NewCWInterface()
	_ = obj.TransmitPower()
	// Output:
	}

// ExampleCWInterface_TransmitRate demonstrates using TransmitRate on a CWInterface instance.
// The current transmit rate (Mbps) for the interface.
func ExampleCWInterface_TransmitRate() {
	obj := corewlan.NewCWInterface()
	_ = obj.TransmitRate()
	// Output:
	}

// ExampleCWInterface_WlanChannel demonstrates using WlanChannel on a CWInterface instance.
// The current channel for the interface.
func ExampleCWInterface_WlanChannel() {
	obj := corewlan.NewCWInterface()
	_ = obj.WlanChannel()
	// Output:
	}

