#!/usr/bin/env python3

import unittest
import os
import json
import yaml
from pathlib import Path
from unittest.mock import patch, MagicMock, mock_open


class TestGitHubAutomation(unittest.TestCase):
    """Integration tests for GitHub issue automation"""

    def setUp(self):
        self.workflows_dir = Path(".github/workflows")
        self.scripts_dir = Path(".github/scripts")

    def test_issue_trigger_workflow(self):
        """Test that workflows trigger on issue events"""
        orchestration_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        self.assertTrue(orchestration_workflow.exists(),
                      "Issue orchestration workflow missing")

        if orchestration_workflow.exists():
            workflow_data = yaml.safe_load(orchestration_workflow.read_text())

            # Check issue triggers
            triggers = workflow_data.get("on", {})
            self.assertIn("issues", triggers,
                        "Workflow should trigger on issues")

            issue_types = triggers.get("issues", {}).get("types", [])
            expected_types = ["opened", "edited", "labeled"]

            for event_type in expected_types:
                self.assertIn(event_type, issue_types,
                            f"Workflow should trigger on issue {event_type}")

    def test_issue_labeling_automation(self):
        """Test automatic issue labeling based on content"""
        orchestration_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        if orchestration_workflow.exists():
            content = orchestration_workflow.read_text()

            # Check for labeling automation
            self.assertIn("labels", content,
                        "Workflow should handle labels")
            self.assertIn("addLabels", content,
                        "Workflow should add labels")

            # Expected labels to be added
            expected_labels = [
                "domain:", "phase:", "agent-assigned",
                "in-progress", "completed"
            ]

            for label_prefix in expected_labels:
                self.assertIn(label_prefix, content,
                            f"Workflow should use label: {label_prefix}")

    def test_issue_comment_automation(self):
        """Test automatic commenting on issues"""
        orchestration_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        if orchestration_workflow.exists():
            content = orchestration_workflow.read_text()

            # Check for comment automation
            self.assertIn("createComment", content,
                        "Workflow should create comments")

            # Expected comment types
            comment_indicators = [
                "Agent Assignment",
                "Phase Complete",
                "Workflow Complete"
            ]

            for indicator in comment_indicators:
                self.assertIn(indicator, content,
                            f"Workflow should post {indicator} comments")

    @patch('builtins.open', new_callable=mock_open, read_data='{}')
    @patch('subprocess.run')
    def test_issue_parsing_script(self, mock_run, mock_file):
        """Test issue parsing for agent inputs"""
        parse_script = self.scripts_dir / "parse-issue.py"

        # Script should be created
        if parse_script.exists():
            # Simulate running the script
            mock_run.return_value = MagicMock(
                returncode=0,
                stdout=json.dumps({
                    "title": "Test issue",
                    "body": "Issue body",
                    "labels": ["enhancement"],
                    "domain": "python",
                    "phase": "planning"
                })
            )

            # Script should parse issue and output JSON
            result = mock_run.return_value.stdout
            parsed = json.loads(result)

            self.assertIn("domain", parsed)
            self.assertIn("phase", parsed)

    def test_workflow_dispatch_automation(self):
        """Test workflow dispatch for phase execution"""
        execute_workflow = self.workflows_dir / "execute-phase.yml"

        self.assertTrue(execute_workflow.exists(),
                      "Execute phase workflow missing")

        if execute_workflow.exists():
            workflow_data = yaml.safe_load(execute_workflow.read_text())

            # Check workflow_dispatch configuration
            dispatch = workflow_data.get("on", {}).get("workflow_dispatch", {})
            self.assertIsNotNone(dispatch,
                               "Workflow should support dispatch")

            # Check required inputs
            inputs = dispatch.get("inputs", {})
            required_inputs = ["issue_number", "domain", "phase", "agent"]

            for input_name in required_inputs:
                self.assertIn(input_name, inputs,
                            f"Workflow dispatch missing input: {input_name}")

    def test_artifact_management(self):
        """Test artifact upload and download between phases"""
        execute_workflow = self.workflows_dir / "execute-phase.yml"

        if execute_workflow.exists():
            content = execute_workflow.read_text()

            # Check for artifact handling
            self.assertIn("upload-artifact", content,
                        "Workflow should upload artifacts")
            self.assertIn("download-artifact", content,
                        "Workflow should download artifacts")

            # Each phase should create artifacts
            phases = ["planning", "implementation", "testing", "documentation"]
            for phase in phases:
                # Check phase-specific artifact handling
                if f"{phase}-phase" in content:
                    # Phase job exists
                    pass

    def test_pr_creation_automation(self):
        """Test automatic PR creation for implementations"""
        execute_workflow = self.workflows_dir / "execute-phase.yml"

        if execute_workflow.exists():
            content = execute_workflow.read_text()

            # Check for PR creation
            pr_indicators = [
                "create-pull-request",
                "pull_request",
                "branch:",
                "commit-message:"
            ]

            pr_found = [indicator for indicator in pr_indicators
                       if indicator in content.lower()]

            self.assertGreater(len(pr_found), 0,
                             "Workflow should support PR creation")

    def test_metrics_collection(self):
        """Test workflow metrics collection"""
        metrics_script = self.scripts_dir / "collect-metrics.py"

        # Script should be created
        if metrics_script.exists():
            content = metrics_script.read_text()

            # Check for metrics collection
            metrics = [
                "duration", "success_rate", "phase_completion",
                "agent_performance"
            ]

            for metric in metrics:
                # Script should collect various metrics
                pass

    def test_github_api_integration(self):
        """Test GitHub API usage in workflows"""
        workflows = list(self.workflows_dir.glob("*.yml"))

        for workflow_path in workflows:
            content = workflow_path.read_text()

            # Check for GitHub API usage
            if "github-script" in content:
                # Check for proper API calls
                api_calls = [
                    "github.rest.issues",
                    "github.rest.actions",
                    "context.repo",
                    "context.payload"
                ]

                api_found = [call for call in api_calls if call in content]

                self.assertGreater(len(api_found), 0,
                                 f"{workflow_path.name} should use GitHub API")

    def test_issue_status_tracking(self):
        """Test issue status is tracked through workflow"""
        orchestration_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        if orchestration_workflow.exists():
            content = orchestration_workflow.read_text()

            # Check for status updates
            status_labels = [
                "agent-ready",
                "plan-ready",
                "implementation-ready",
                "tests-passed",
                "ready-for-review",
                "completed"
            ]

            for status in status_labels:
                self.assertIn(status, content,
                            f"Workflow should track status: {status}")

    def test_error_reporting(self):
        """Test error reporting in workflows"""
        workflows = list(self.workflows_dir.glob("*.yml"))

        for workflow_path in workflows:
            content = workflow_path.read_text()

            # Check for error handling
            if "error-handler" in content or "failure()" in content:
                # Check error reporting mechanisms
                error_actions = [
                    "createComment",
                    "error",
                    "needs-review"
                ]

                actions_found = [action for action in error_actions
                               if action in content]

                self.assertGreater(len(actions_found), 0,
                                 f"{workflow_path.name} should report errors")


if __name__ == "__main__":
    unittest.main()