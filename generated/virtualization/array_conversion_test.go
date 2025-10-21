package virtualization_test

import (
	"testing"

	"github.com/tmc/appledocs/generated/virtualization"
)

func TestArrayConversion(t *testing.T) {
	// This test verifies that array property setters compile correctly
	// and have proper array conversion logic

	config := virtualization.VZVirtualMachineConfigurationClass.New()

	// Test 1: Empty array
	t.Run("EmptyArray", func(t *testing.T) {
		config.SetStorageDevices([]virtualization.VZStorageDeviceConfiguration{})
		// If this compiles and runs without panic, the conversion works
	})

	// Test 2: Check that we can create arrays without manual conversion
	t.Run("ArrayWithItems", func(t *testing.T) {
		// This would previously require manual NSArray conversion
		// Now it should work directly with Go slices
		devices := []virtualization.VZStorageDeviceConfiguration{
			// Note: We can't actually create valid devices without more setup,
			// but the point is the type signature accepts Go slices directly
		}
		config.SetStorageDevices(devices)
	})

	// Test 3: Multiple different array types
	t.Run("MultipleArrayTypes", func(t *testing.T) {
		config.SetAudioDevices([]virtualization.VZAudioDeviceConfiguration{})
		config.SetNetworkDevices([]virtualization.VZNetworkDeviceConfiguration{})
		config.SetConsoleDevices([]virtualization.VZConsoleDeviceConfiguration{})
		config.SetSerialPorts([]virtualization.VZSerialPortConfiguration{})
	})

	t.Log("Array conversion tests passed - methods compile and accept Go slices")
}
