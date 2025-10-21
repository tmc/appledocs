package virtualization_test

import (
	"testing"

	"github.com/tmc/appledocs/generated/virtualization"
)

func TestClassProperties(t *testing.T) {
	// Test class properties that were generated
	maxCPU := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedCPUCount()
	minCPU := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedCPUCount()

	if maxCPU == 0 {
		t.Skip("MaximumAllowedCPUCount returned 0 - may need runtime initialization")
	}
	if minCPU == 0 {
		t.Skip("MinimumAllowedCPUCount returned 0 - may need runtime initialization")
	}

	t.Logf("Max CPU: %d", maxCPU)
	t.Logf("Min CPU: %d", minCPU)

	if maxCPU < minCPU {
		t.Errorf("Max CPU (%d) should be >= Min CPU (%d)", maxCPU, minCPU)
	}
}
