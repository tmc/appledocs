// Package benchmark provides profiling capabilities for performance analysis
package benchmark

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

// Profiler handles CPU, memory, and blocking profiles
type Profiler struct {
	config      Config
	cpuFile     *os.File
	traceFile   *os.File
	profileDir  string
	active      bool
}

// NewProfiler creates a new profiler
func NewProfiler(config Config) *Profiler {
	profileDir := filepath.Join(config.OutputDir, "profiles")
	os.MkdirAll(profileDir, 0755)
	
	return &Profiler{
		config:     config,
		profileDir: profileDir,
	}
}

// StartCPUProfile starts CPU profiling
func (p *Profiler) StartCPUProfile(name string) error {
	if !p.config.CPUProfile {
		return nil
	}
	
	filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_cpu_%d.prof", name, time.Now().Unix()))
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create CPU profile: %w", err)
	}
	
	p.cpuFile = f
	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return fmt.Errorf("start CPU profile: %w", err)
	}
	
	p.active = true
	return nil
}

// StopCPUProfile stops CPU profiling
func (p *Profiler) StopCPUProfile() {
	if p.cpuFile == nil {
		return
	}
	
	pprof.StopCPUProfile()
	p.cpuFile.Close()
	p.cpuFile = nil
	p.active = false
}

// CaptureMemProfile captures a memory profile
func (p *Profiler) CaptureMemProfile(name string, iteration int) error {
	if !p.config.MemProfile {
		return nil
	}
	
	filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_mem_%d_iter_%d.prof", name, time.Now().Unix(), iteration))
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create memory profile: %w", err)
	}
	defer f.Close()
	
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		return fmt.Errorf("write heap profile: %w", err)
	}
	
	return nil
}

// CaptureBlockProfile captures a blocking profile
func (p *Profiler) CaptureBlockProfile(name string) error {
	if !p.config.BlockProfile {
		return nil
	}
	
	filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_block_%d.prof", name, time.Now().Unix()))
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create block profile: %w", err)
	}
	defer f.Close()
	
	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)
	
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		return fmt.Errorf("write block profile: %w", err)
	}
	
	return nil
}

// StartTrace starts execution tracing
func (p *Profiler) StartTrace(name string) error {
	if !p.config.TraceEnabled {
		return nil
	}
	
	filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_trace_%d.trace", name, time.Now().Unix()))
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create trace file: %w", err)
	}
	
	p.traceFile = f
	if err := trace.Start(f); err != nil {
		f.Close()
		return fmt.Errorf("start trace: %w", err)
	}
	
	return nil
}

// StopTrace stops execution tracing
func (p *Profiler) StopTrace() {
	if p.traceFile == nil {
		return
	}
	
	trace.Stop()
	p.traceFile.Close()
	p.traceFile = nil
}

// CaptureAllProfiles captures all types of profiles
func (p *Profiler) CaptureAllProfiles(name string) error {
	timestamp := time.Now().Unix()
	
	// CPU profile
	if p.config.CPUProfile {
		filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_cpu_%d.prof", name, timestamp))
		if err := p.captureCPUProfile(filename); err != nil {
			return fmt.Errorf("capture CPU profile: %w", err)
		}
	}
	
	// Memory profile
	if p.config.MemProfile {
		filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_mem_%d.prof", name, timestamp))
		if err := p.captureMemoryProfile(filename); err != nil {
			return fmt.Errorf("capture memory profile: %w", err)
		}
	}
	
	// Block profile
	if p.config.BlockProfile {
		filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_block_%d.prof", name, timestamp))
		if err := p.captureBlockingProfile(filename); err != nil {
			return fmt.Errorf("capture block profile: %w", err)
		}
	}
	
	// Goroutine profile
	filename := filepath.Join(p.profileDir, fmt.Sprintf("%s_goroutine_%d.prof", name, timestamp))
	if err := p.captureGoroutineProfile(filename); err != nil {
		return fmt.Errorf("capture goroutine profile: %w", err)
	}
	
	return nil
}

// captureCPUProfile captures CPU profile for a duration
func (p *Profiler) captureCPUProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	
	if err := pprof.StartCPUProfile(f); err != nil {
		return err
	}
	
	time.Sleep(5 * time.Second) // Profile for 5 seconds
	pprof.StopCPUProfile()
	
	return nil
}

// captureMemoryProfile captures memory profile
func (p *Profiler) captureMemoryProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	
	runtime.GC()
	return pprof.WriteHeapProfile(f)
}

// captureBlockingProfile captures blocking profile
func (p *Profiler) captureBlockingProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	
	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)
	
	// Let it collect for a bit
	time.Sleep(time.Second)
	
	return pprof.Lookup("block").WriteTo(f, 0)
}

// captureGoroutineProfile captures goroutine profile
func (p *Profiler) captureGoroutineProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	
	return pprof.Lookup("goroutine").WriteTo(f, 0)
}

// IsActive returns whether profiling is active
func (p *Profiler) IsActive() bool {
	return p.active
}

// GetProfileDir returns the profile directory
func (p *Profiler) GetProfileDir() string {
	return p.profileDir
}

// Cleanup performs cleanup of profiling resources
func (p *Profiler) Cleanup() {
	p.StopCPUProfile()
	p.StopTrace()
}

// ProfileAnalysis provides analysis of captured profiles
type ProfileAnalysis struct {
	CPUHotspots    []Hotspot    `json:"cpu_hotspots"`
	MemoryLeaks    []MemoryLeak `json:"memory_leaks"`
	BlockingCalls  []BlockCall  `json:"blocking_calls"`
	GoroutineStats GoroutineStats `json:"goroutine_stats"`
}

// Hotspot represents a CPU hotspot
type Hotspot struct {
	Function    string  `json:"function"`
	File        string  `json:"file"`
	Line        int     `json:"line"`
	CPUPercent  float64 `json:"cpu_percent"`
	Calls       int64   `json:"calls"`
}

// MemoryLeak represents a potential memory leak
type MemoryLeak struct {
	Function    string `json:"function"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Bytes       int64  `json:"bytes"`
	Objects     int64  `json:"objects"`
}

// BlockCall represents a blocking call
type BlockCall struct {
	Function    string        `json:"function"`
	File        string        `json:"file"`
	Line        int           `json:"line"`
	BlockTime   time.Duration `json:"block_time"`
	BlockCount  int64         `json:"block_count"`
}

// GoroutineStats holds goroutine statistics
type GoroutineStats struct {
	Total      int `json:"total"`
	Running    int `json:"running"`
	Waiting    int `json:"waiting"`
	Sleeping   int `json:"sleeping"`
	Blocking   int `json:"blocking"`
}

// AnalyzeProfiles analyzes captured profiles
func (p *Profiler) AnalyzeProfiles(name string) (*ProfileAnalysis, error) {
	analysis := &ProfileAnalysis{
		CPUHotspots:   make([]Hotspot, 0),
		MemoryLeaks:   make([]MemoryLeak, 0),
		BlockingCalls: make([]BlockCall, 0),
	}
	
	// This would typically use go tool pprof programmatically
	// For now, we'll return a basic analysis
	
	return analysis, nil
}

// GenerateProfileReport generates a comprehensive profile report
func (p *Profiler) GenerateProfileReport(name string) (*ProfileReport, error) {
	report := &ProfileReport{
		Name:      name,
		Timestamp: time.Now(),
		Profiles:  make([]ProfileInfo, 0),
	}
	
	// Find all profile files for this benchmark
	pattern := filepath.Join(p.profileDir, fmt.Sprintf("%s_*.prof", name))
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("find profile files: %w", err)
	}
	
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		
		profileInfo := ProfileInfo{
			Type:     getProfileType(file),
			Filename: file,
			Size:     info.Size(),
			ModTime:  info.ModTime(),
		}
		
		report.Profiles = append(report.Profiles, profileInfo)
	}
	
	return report, nil
}

// ProfileReport contains profile analysis results
type ProfileReport struct {
	Name      string        `json:"name"`
	Timestamp time.Time     `json:"timestamp"`
	Profiles  []ProfileInfo `json:"profiles"`
	Analysis  *ProfileAnalysis `json:"analysis,omitempty"`
}

// ProfileInfo contains information about a profile file
type ProfileInfo struct {
	Type     string    `json:"type"`
	Filename string    `json:"filename"`
	Size     int64     `json:"size"`
	ModTime  time.Time `json:"mod_time"`
}

// getProfileType determines profile type from filename
func getProfileType(filename string) string {
	if strings.Contains(filename, "_cpu_") {
		return "cpu"
	} else if strings.Contains(filename, "_mem_") {
		return "memory"
	} else if strings.Contains(filename, "_block_") {
		return "block"
	} else if strings.Contains(filename, "_goroutine_") {
		return "goroutine"
	}
	return "unknown"
}

// ProfileComparer compares profiles between phases
type ProfileComparer struct {
	baselineProfiles map[string]*ProfileReport
	currentProfiles  map[string]*ProfileReport
}

// NewProfileComparer creates a new profile comparer
func NewProfileComparer() *ProfileComparer {
	return &ProfileComparer{
		baselineProfiles: make(map[string]*ProfileReport),
		currentProfiles:  make(map[string]*ProfileReport),
	}
}

// AddBaseline adds a baseline profile
func (pc *ProfileComparer) AddBaseline(name string, report *ProfileReport) {
	pc.baselineProfiles[name] = report
}

// AddCurrent adds a current profile
func (pc *ProfileComparer) AddCurrent(name string, report *ProfileReport) {
	pc.currentProfiles[name] = report
}

// Compare compares profiles and returns differences
func (pc *ProfileComparer) Compare() *ProfileComparison {
	comparison := &ProfileComparison{
		Improvements: make([]string, 0),
		Regressions:  make([]string, 0),
		NewIssues:    make([]string, 0),
		FixedIssues:  make([]string, 0),
	}
	
	// Compare profiles
	for name, currentReport := range pc.currentProfiles {
		if baselineReport, exists := pc.baselineProfiles[name]; exists {
			pc.compareReports(baselineReport, currentReport, comparison)
		} else {
			comparison.NewIssues = append(comparison.NewIssues, 
				fmt.Sprintf("New profile: %s", name))
		}
	}
	
	return comparison
}

// ProfileComparison holds profile comparison results
type ProfileComparison struct {
	Improvements []string `json:"improvements"`
	Regressions  []string `json:"regressions"`
	NewIssues    []string `json:"new_issues"`
	FixedIssues  []string `json:"fixed_issues"`
}

// compareReports compares two profile reports
func (pc *ProfileComparer) compareReports(baseline, current *ProfileReport, comparison *ProfileComparison) {
	// Compare profile sizes (rough indicator of performance)
	for _, baseProfile := range baseline.Profiles {
		for _, currentProfile := range current.Profiles {
			if baseProfile.Type == currentProfile.Type {
				sizeDiff := float64(currentProfile.Size-baseProfile.Size) / float64(baseProfile.Size)
				if sizeDiff > 0.1 { // 10% increase
					comparison.Regressions = append(comparison.Regressions,
						fmt.Sprintf("%s profile size increased by %.1f%%", baseProfile.Type, sizeDiff*100))
				} else if sizeDiff < -0.1 { // 10% decrease
					comparison.Improvements = append(comparison.Improvements,
						fmt.Sprintf("%s profile size decreased by %.1f%%", baseProfile.Type, -sizeDiff*100))
				}
			}
		}
	}
}