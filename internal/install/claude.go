package install

import (
	"fmt"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
)

// SetupClaudeDirectory creates the .claude/ directory structure for Claude Code integration
func SetupClaudeDirectory() error {
	claudeDir, err := config.GetClaudeDir()
	if err != nil {
		return fmt.Errorf("failed to get Claude directory path: %w", err)
	}

	// Create main .claude/ directory
	if err := config.EnsureDir(claudeDir); err != nil {
		return fmt.Errorf("failed to create .claude/ directory: %w", err)
	}

	// Create subdirectories
	commandsDir, err := config.GetClaudeCommandsDir()
	if err != nil {
		return err
	}
	if err := config.EnsureDir(commandsDir); err != nil {
		return fmt.Errorf("failed to create .claude/commands/ directory: %w", err)
	}

	agentsDir, err := config.GetClaudeAgentsDir()
	if err != nil {
		return err
	}
	if err := config.EnsureDir(agentsDir); err != nil {
		return fmt.Errorf("failed to create .claude/agents/ directory: %w", err)
	}

	skillsDir, err := config.GetClaudeSkillsDir()
	if err != nil {
		return err
	}
	if err := config.EnsureDir(skillsDir); err != nil {
		return fmt.Errorf("failed to create .claude/skills/ directory: %w", err)
	}

	return nil
}

// IntegrateWithClaude copies agents and commands to .claude/ directories
func IntegrateWithClaude(paths *InstallationPaths) (*ClaudeIntegrationResult, error) {
	result := &ClaudeIntegrationResult{}

	// Ensure .claude/ structure exists
	if err := SetupClaudeDirectory(); err != nil {
		return nil, err
	}

	// Copy agents with "cat-" prefix
	if config.IsDirectory(paths.AgentsSourceDir) {
		if err := CopyAgentsWithPrefix(paths.AgentsSourceDir, paths.ClaudeAgents); err != nil {
			return nil, fmt.Errorf("failed to copy agents: %w", err)
		}

		// Count copied agents
		agentCount, err := CountFiles(paths.ClaudeAgents)
		if err != nil {
			return nil, fmt.Errorf("failed to count agents: %w", err)
		}
		result.AgentsCopied = agentCount
	}

	// Copy spec-kit commands with "speckit." prefix
	if config.IsDirectory(paths.TemplatesDir) {
		if err := CopyCommandsWithPrefix(paths.TemplatesDir, paths.ClaudeCommands); err != nil {
			return nil, fmt.Errorf("failed to copy commands: %w", err)
		}

		// Count copied commands
		commandCount, err := CountFiles(paths.ClaudeCommands)
		if err != nil {
			return nil, fmt.Errorf("failed to count commands: %w", err)
		}
		result.CommandsCopied = commandCount
	}

	result.Success = true
	return result, nil
}

// ClaudeIntegrationResult contains the results of Claude Code integration
type ClaudeIntegrationResult struct {
	Success        bool
	AgentsCopied   int
	CommandsCopied int
}

// GetSummary returns a human-readable summary of the integration
func (r *ClaudeIntegrationResult) GetSummary() string {
	if !r.Success {
		return "Claude Code integration failed"
	}
	return fmt.Sprintf("Integrated with Claude Code: %d agents, %d commands", r.AgentsCopied, r.CommandsCopied)
}

// VerifyClaudeIntegration checks that Claude Code integration is working
func VerifyClaudeIntegration(paths *InstallationPaths) error {
	// Check that .claude/ exists
	if !config.PathExists(paths.ClaudeDir) {
		return fmt.Errorf(".claude/ directory not found")
	}

	// Check that subdirectories exist
	if !config.IsDirectory(paths.ClaudeCommands) {
		return fmt.Errorf(".claude/commands/ directory not found")
	}

	if !config.IsDirectory(paths.ClaudeAgents) {
		return fmt.Errorf(".claude/agents/ directory not found")
	}

	// Check that at least one command was copied
	commandCount, err := CountFiles(paths.ClaudeCommands)
	if err != nil {
		return fmt.Errorf("failed to count commands: %w", err)
	}
	if commandCount == 0 {
		return fmt.Errorf("no commands found in .claude/commands/")
	}

	// Check that at least one agent was copied
	agentCount, err := CountFiles(paths.ClaudeAgents)
	if err != nil {
		return fmt.Errorf("failed to count agents: %w", err)
	}
	if agentCount == 0 {
		return fmt.Errorf("no agents found in .claude/agents/")
	}

	return nil
}
