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
	Version   = "2.0.0"
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

	// Update command flags
	updateNoBackup  bool
	updateForce     bool
	updateSkipVerify bool

	// Rollback command flags
	rollbackBackupID string
	rollbackList     bool
	rollbackForce    bool
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

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update installation to latest version",
	Long: `Update the installation to the latest version from the manifest.

This command:
  - Creates a backup of the current installation (unless --no-backup)
  - Updates to the version specified in the manifest
  - Automatically rolls back on failure
  - Preserves installation history

Examples:
  # Update to latest version with automatic backup
  spec-kit-agents update

  # Update without creating backup
  spec-kit-agents update --no-backup

  # Force update even if versions match
  spec-kit-agents update --force`,
	RunE: runUpdate,
}

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback to a previous installation",
	Long: `Rollback the installation to a previous state from backup.

This command restores a previous installation from backup. By default,
it restores the most recent backup. You can specify a backup ID to
restore a specific backup.

Examples:
  # Rollback to latest backup
  spec-kit-agents rollback

  # Rollback to specific backup
  spec-kit-agents rollback --backup-id=backup-20251022-120000

  # List available backups
  spec-kit-agents rollback --list`,
	RunE: runRollback,
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(rollbackCmd)

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

	// Update command flags
	updateCmd.Flags().StringVar(&installPrefix, "prefix", "", "Installation prefix (default: auto-detect)")
	updateCmd.Flags().BoolVar(&updateNoBackup, "no-backup", false, "Skip backup creation before update")
	updateCmd.Flags().BoolVar(&updateForce, "force", false, "Force update even if versions match")
	updateCmd.Flags().BoolVar(&updateSkipVerify, "skip-verify", false, "Skip version compatibility verification")

	// Rollback command flags
	rollbackCmd.Flags().StringVar(&installPrefix, "prefix", "", "Installation prefix (default: auto-detect)")
	rollbackCmd.Flags().StringVar(&rollbackBackupID, "backup-id", "", "Specific backup to restore (default: latest)")
	rollbackCmd.Flags().BoolVar(&rollbackList, "list", false, "List available backups")
	rollbackCmd.Flags().BoolVar(&rollbackForce, "force", false, "Force rollback without confirmation")
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

func runUpdate(cmd *cobra.Command, args []string) error {
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

	// Prepare options
	opts := install.UpdateOptions{
		Backup:     !updateNoBackup,
		Force:      updateForce,
		SkipVerify: updateSkipVerify,
	}

	// Run update
	result, err := install.Update(prefix, opts, logger)
	if err != nil {
		logger.Error("update", "Update failed: %v", err)
		return err
	}

	if !result.Success {
		return fmt.Errorf("update did not complete successfully")
	}

	// Display summary
	fmt.Println()
	fmt.Println("Update Summary")
	fmt.Println("==============")
	fmt.Printf("  Updated from: v%s\n", result.UpdatedFrom)
	fmt.Printf("  Updated to:   v%s\n", result.UpdatedTo)
	fmt.Printf("  Components:   %d\n", result.ComponentsUpdated)
	if result.BackupCreated {
		fmt.Printf("  Backup ID:    %s\n", result.BackupID)
	}
	fmt.Println()

	if len(result.Warnings) > 0 {
		fmt.Println("Warnings:")
		for _, warning := range result.Warnings {
			fmt.Printf("  - %s\n", warning)
		}
		fmt.Println()
	}

	return nil
}

func runRollback(cmd *cobra.Command, args []string) error {
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

	// List backups if requested
	if rollbackList {
		backups, err := install.ListBackups(prefix)
		if err != nil {
			return fmt.Errorf("failed to list backups: %w", err)
		}

		if len(backups) == 0 {
			fmt.Println("No backups found")
			return nil
		}

		fmt.Printf("Available backups for %s:\n\n", prefix)
		for i, backup := range backups {
			fmt.Printf("%d. %s\n", i+1, backup.BackupID)
			fmt.Printf("   Created: %s\n", backup.CreatedAt.Format("2006-01-02 15:04:05 UTC"))
			fmt.Printf("   Path:    %s\n", backup.BackupPath)
			fmt.Println()
		}
		return nil
	}

	// Check if rollback is possible
	canRollback, message, err := install.CanRollback(prefix)
	if err != nil {
		return fmt.Errorf("failed to check rollback status: %w", err)
	}

	if !canRollback {
		return fmt.Errorf("cannot rollback: %s", message)
	}

	// Confirm rollback if not forced
	if !rollbackForce && !quiet {
		fmt.Printf("This will rollback your installation.\n")
		fmt.Printf("%s\n\n", message)
		fmt.Print("Are you sure you want to continue? (yes/no): ")

		var response string
		fmt.Scanln(&response)

		if response != "yes" && response != "y" {
			fmt.Println("Rollback cancelled")
			return nil
		}
	}

	// Prepare options
	opts := install.RollbackOptions{
		BackupID: rollbackBackupID,
		Force:    rollbackForce,
	}

	// Run rollback
	result, err := install.Rollback(prefix, opts, logger)
	if err != nil {
		logger.Error("rollback", "Rollback failed: %v", err)
		return err
	}

	if !result.Success {
		return fmt.Errorf("rollback did not complete successfully")
	}

	// Display summary
	fmt.Println()
	fmt.Println("Rollback Summary")
	fmt.Println("================")
	fmt.Printf("  Restored from:   %s\n", result.RestoredFromID)
	fmt.Printf("  Previous version: v%s\n", result.PreviousVersion)
	fmt.Printf("  Restored version: v%s\n", result.RestoredVersion)
	fmt.Printf("  Components:       %d\n", result.ComponentsRestored)
	fmt.Println()

	return nil
}
