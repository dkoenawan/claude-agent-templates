package main

import (
	"fmt"
	"os"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/internal/install"
	"github.com/dkoenawan/claude-agent-templates/internal/version"
	"github.com/spf13/cobra"
)

var (
	// Version information (set by build)
	Version   = "1.0.0"
	BuildTime = "unknown"
	GitCommit = "unknown"

	// Global flags
	verbose bool
	quiet   bool

	// Command-specific flags
	installPrefix string
	installGlobal bool
	installForce  bool
	installDryRun bool
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "spec-kit-agents",
	Short: "Claude Agent Templates installer with lockstep spec-kit version management",
	Long: `spec-kit-agents provides lockstep installation of claude-agent-templates
and GitHub spec-kit with version compatibility management.

This tool ensures that claude-agent-templates and spec-kit are installed together
with compatible, pinned versions to prevent breaking changes from uncontrolled
spec-kit upgrades.`,
	SilenceUsage: true,
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install claude-agent-templates with spec-kit",
	Long: `Install claude-agent-templates along with the pinned version of spec-kit.

This command:
  - Detects existing installations and chooses appropriate installation mode
  - Copies spec-kit files to the installation directory
  - Integrates with Claude Code (.claude/agents and .claude/commands)
  - Creates a version lock file for tracking
  - Verifies the installation

Examples:
  # Fresh installation (auto-detects mode)
  spec-kit-agents install

  # Install to custom directory
  spec-kit-agents install --prefix /path/to/dir

  # Install globally to ~/.claude/agents/
  spec-kit-agents install --global

  # Force reinstall (overwrite existing)
  spec-kit-agents install --force

  # Dry run (show what would be done)
  spec-kit-agents install --dry-run`,
	RunE: runInstall,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show installation status and version information",
	Long: `Display the current installation status, including:
  - Installation location
  - Installed versions of claude-agent-templates and spec-kit
  - Installation timestamp
  - Last verification time
  - Installation history

Examples:
  # Show status for default installation
  spec-kit-agents status

  # Show status for specific prefix
  spec-kit-agents status --prefix /path/to/installation`,
	RunE: runStatus,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Display version information for the spec-kit-agents tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("spec-kit-agents version %s\n", Version)
		fmt.Printf("  Build time: %s\n", BuildTime)
		fmt.Printf("  Git commit: %s\n", GitCommit)
	},
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check version compatibility",
	Long: `Check if the installed spec-kit version is compatible with the
pinned version in the manifest.

This command compares the installed spec-kit version with the required
version and reports any incompatibilities.`,
	RunE: runCheck,
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(checkCmd)

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress non-error output")

	// Install command flags
	installCmd.Flags().StringVar(&installPrefix, "prefix", "", "Installation prefix (auto-detected if not specified)")
	installCmd.Flags().BoolVar(&installGlobal, "global", false, "Install globally to ~/.claude/agents/")
	installCmd.Flags().BoolVar(&installForce, "force", false, "Force installation even if already installed")
	installCmd.Flags().BoolVar(&installDryRun, "dry-run", false, "Show what would be done without actually installing")

	// Status command flags
	statusCmd.Flags().StringVar(&installPrefix, "prefix", "", "Installation prefix to check (default: auto-detect)")
}

func createLogger() (*config.Logger, error) {
	// Determine log level
	logLevel := config.INFO
	if verbose {
		logLevel = config.DEBUG
	}
	if quiet {
		logLevel = config.ERROR
	}

	// Create logger
	logger, err := config.NewLogger(logLevel, "", !quiet) // Log to console only for now
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return logger, nil
}

func runInstall(cmd *cobra.Command, args []string) error {
	logger, err := createLogger()
	if err != nil {
		return err
	}
	defer logger.Close()

	// Prepare options
	opts := install.Options{
		Prefix: installPrefix,
		Global: installGlobal,
		Force:  installForce,
		Quiet:  quiet,
		DryRun: installDryRun,
	}

	// TODO: Handle global installation differently
	if installGlobal {
		return fmt.Errorf("global installation not yet implemented")
	}

	// Run installation
	result, err := install.Run(opts, logger)
	if err != nil {
		logger.Error("installer", "Installation failed: %v", err)
		return err
	}

	if !result.Success {
		return fmt.Errorf("installation did not complete successfully")
	}

	return nil
}

func runStatus(cmd *cobra.Command, args []string) error {
	logger, err := createLogger()
	if err != nil {
		return err
	}
	defer logger.Close()

	// Determine prefix
	prefix := installPrefix
	if prefix == "" {
		prefix = config.DetermineInstallPrefix()
	}

	// Get status
	status, err := install.GetStatus(prefix)
	if err != nil {
		return fmt.Errorf("failed to get installation status: %w", err)
	}

	if !status.Installed {
		fmt.Printf("No installation found at %s\n", prefix)
		return nil
	}

	// Display status
	fmt.Printf("Installation Status\n")
	fmt.Printf("===================\n\n")
	fmt.Printf("  Location:              %s\n", status.Prefix)
	fmt.Printf("  Installation ID:       %s\n", status.InstallationID)
	fmt.Printf("  Installed at:          %s\n", status.InstalledAt)
	fmt.Printf("  Last verified:         %s\n", status.LastVerified)
	fmt.Printf("\n")
	fmt.Printf("Versions\n")
	fmt.Printf("========\n\n")
	fmt.Printf("  claude-agent-templates: v%s\n", status.TemplatesVersion)
	fmt.Printf("  spec-kit:               v%s\n", status.SpecKitVersion)
	fmt.Printf("\n")
	fmt.Printf("History\n")
	fmt.Printf("=======\n\n")
	fmt.Printf("  %d installation event(s)\n", status.HistoryEntryCount)

	return nil
}

func runCheck(cmd *cobra.Command, args []string) error {
	logger, err := createLogger()
	if err != nil {
		return err
	}
	defer logger.Close()

	// Determine prefix
	prefix := installPrefix
	if prefix == "" {
		prefix = config.DetermineInstallPrefix()
	}

	logger.Info("checker", "Checking version compatibility...")

	// Load version manifest
	manifest, err := version.LoadManifestFromPrefix(prefix)
	if err != nil {
		return fmt.Errorf("failed to load version manifest: %w", err)
	}

	requiredVersion, err := version.GetSpecKitVersion(manifest)
	if err != nil {
		return fmt.Errorf("failed to get required spec-kit version: %w", err)
	}

	compatibility, err := version.GetSpecKitCompatibility(manifest)
	if err != nil {
		return fmt.Errorf("failed to get compatibility constraints: %w", err)
	}

	logger.Info("checker", "Required spec-kit version: v%s", requiredVersion)
	logger.Info("checker", "Compatibility range: v%s - v%s", compatibility.MinVersion, compatibility.MaxVersion)

	// Load version lock
	paths, err := install.GetPaths(prefix)
	if err != nil {
		return fmt.Errorf("failed to get installation paths: %w", err)
	}

	if !config.PathExists(paths.VersionLock) {
		return fmt.Errorf("no installation found at %s", prefix)
	}

	lock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		return fmt.Errorf("failed to load version lock: %w", err)
	}

	installedVersion, err := version.GetInstalledSpecKitVersion(lock)
	if err != nil {
		return fmt.Errorf("failed to get installed spec-kit version: %w", err)
	}

	logger.Info("checker", "Installed spec-kit version: v%s", installedVersion)

	// Check compatibility
	result, err := version.CheckCompatibility(
		installedVersion,
		requiredVersion,
		compatibility.MinVersion,
		compatibility.MaxVersion,
		compatibility.BreakingVersions,
	)
	if err != nil {
		return fmt.Errorf("compatibility check failed: %w", err)
	}

	if result.IsCompatible() {
		logger.Success("checker", "✓ Versions are compatible")
		if result.VersionMismatch {
			logger.Warn("checker", "Version mismatch detected but within compatible range")
		}
		return nil
	}

	// Incompatible
	logger.Error("checker", "✗ Version incompatibility detected")
	fmt.Println(result.GetIssuesText())
	return fmt.Errorf("version compatibility check failed")
}
