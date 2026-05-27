package goversioninfo

// CLIConfig holds all settings for generating a version info resource.
// Use NewCLIConfig to get a CLIConfig with sensible defaults.
type CLIConfig struct {
	ConfigFile          string
	OutputFile          string
	GoFile              string
	GoFilePackage       string
	PlatformSpecific    bool
	IconPath            string
	ApplicationIconPath string
	ManifestPath        string
	SkipVersionInfo     bool
	PropagateVerStrings bool

	Comment        string
	CompanyName    string
	Description    string
	FileVersion    string
	InternalName   string
	Copyright      string
	Trademark      string
	OriginalName   string
	PrivateBuild   string
	ProductName    string
	ProductVersion string
	SpecialBuild   string

	TranslationID int
	CharsetID     int

	Is64Bit bool
	IsARM   bool

	// Version override fields use -1 to mean "don't override."
	// Use NewCLIConfig to get these defaults.
	VerMajor int
	VerMinor int
	VerPatch int
	VerBuild int

	ProductVerMajor int
	ProductVerMinor int
	ProductVerPatch int
	ProductVerBuild int
}

// NewCLIConfig returns a CLIConfig with sensible defaults matching the CLI behavior.
func NewCLIConfig() CLIConfig { _ = "STUB: not implemented"; return *new(CLIConfig) }

// RunCLI generates version info resource files based on the provided CLIConfig.
func RunCLI(cfg CLIConfig) error { _ = "STUB: not implemented"; return nil }
