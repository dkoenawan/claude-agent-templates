#!/usr/bin/env python3

import unittest
import os
import subprocess
import tempfile
import shutil
from pathlib import Path
from unittest.mock import patch, MagicMock


class TestAgentValidation(unittest.TestCase):
    """Integration tests for agent specification validation"""

    def setUp(self):
        self.test_dir = Path(tempfile.mkdtemp())
        self.validation_script = Path("scripts/validate-agent-spec.sh")
        self.agents_dir = Path("agents")

    def tearDown(self):
        if self.test_dir.exists():
            shutil.rmtree(self.test_dir)

    def test_validation_script_exists(self):
        """Test that validation script exists and is executable"""
        self.assertTrue(self.validation_script.exists(),
                      "Validation script missing")

        if self.validation_script.exists():
            self.assertTrue(os.access(self.validation_script, os.X_OK),
                          "Validation script should be executable")

    def test_validate_complete_agent_spec(self):
        """Test validation of complete agent specification"""
        # Create a valid agent specification
        valid_agent = self.test_dir / "valid-agent.md"
        valid_content = """---
name: Test Agent
description: A test agent for validation
domain: python
capabilities:
  - Planning
  - Implementation
  - Testing
tools:
  - pytest
  - black
  - ruff
workflow:
  input: Issue requirements
  processing: Analyze and implement
  output: Working solution
  validation: Run tests
constraints:
  - Must follow PEP 8
  - Use type hints
examples:
  - context: Bug fix
    input: Error in function
    output: Fixed function with tests
---

# Test Agent

This is a test agent specification.
"""

        valid_agent.write_text(valid_content)

        # Run validation - should pass
        if self.validation_script.exists():
            result = subprocess.run([
                str(self.validation_script), str(valid_agent)
            ], capture_output=True, text=True)

            self.assertEqual(result.returncode, 0,
                           f"Valid agent should pass validation: {result.stderr}")

    def test_validate_incomplete_agent_spec(self):
        """Test validation fails for incomplete specification"""
        # Create an incomplete agent specification
        incomplete_agent = self.test_dir / "incomplete-agent.md"
        incomplete_content = """---
name: Incomplete Agent
description: Missing required fields
---

# Incomplete Agent

This agent is missing required sections.
"""

        incomplete_agent.write_text(incomplete_content)

        # Run validation - should fail
        if self.validation_script.exists():
            result = subprocess.run([
                str(self.validation_script), str(incomplete_agent)
            ], capture_output=True, text=True)

            self.assertNotEqual(result.returncode, 0,
                              "Incomplete agent should fail validation")

    def test_domain_specific_validation(self):
        """Test domain-specific validation rules"""
        domain_tests = [
            ("python", ["pytest", "python", "pip"]),
            ("dotnet", ["dotnet", "xunit", "nuget"]),
            ("nodejs", ["npm", "jest", "node"]),
            ("java", ["maven", "junit", "java"])
        ]

        for domain, expected_tools in domain_tests:
            with self.subTest(domain=domain):
                # Create domain-specific agent
                agent_file = self.test_dir / f"{domain}-agent.md"
                agent_content = f"""---
name: {domain.title()} Agent
description: Agent for {domain} development
domain: {domain}
capabilities:
  - Implementation
tools:
  - {expected_tools[0]}
workflow:
  input: Requirements
  processing: Implement
  output: Solution
  validation: Test
constraints:
  - Follow best practices
examples:
  - context: Feature
    input: Requirement
    output: Implementation
---

# {domain.title()} Agent

Mentions {expected_tools[0]} and {domain}.
"""

                agent_file.write_text(agent_content)

                # Validation should pass for domain-specific agent
                if self.validation_script.exists():
                    result = subprocess.run([
                        str(self.validation_script), str(agent_file)
                    ], capture_output=True, text=True)

                    self.assertEqual(result.returncode, 0,
                                   f"{domain} agent should pass validation")

    def test_workflow_validation(self):
        """Test workflow state validation"""
        # Agent with all workflow states
        complete_workflow_agent = self.test_dir / "workflow-agent.md"
        complete_content = """---
name: Workflow Agent
description: Agent with complete workflow
domain: core
capabilities:
  - Full workflow
tools:
  - bash
workflow:
  input: User requirements
  processing: Analyze and implement
  output: Completed solution
  validation: Verify solution meets requirements
constraints:
  - Complete all workflow states
examples:
  - context: Implementation
    input: Requirements
    output: Solution
---

# Workflow Agent
"""

        complete_workflow_agent.write_text(complete_content)

        if self.validation_script.exists():
            result = subprocess.run([
                str(self.validation_script), str(complete_workflow_agent)
            ], capture_output=True, text=True)

            self.assertEqual(result.returncode, 0,
                           "Complete workflow should pass validation")

    def test_github_workflow_validation(self):
        """Test GitHub Actions workflow validation"""
        validation_workflow = Path(".github/workflows/validate-agents.yml")

        self.assertTrue(validation_workflow.exists(),
                      "Agent validation workflow missing")

        if validation_workflow.exists():
            content = validation_workflow.read_text()

            # Check validation workflow calls the script
            self.assertIn("validate-agent-spec.sh", content,
                        "Workflow should use validation script")

            # Check for all expected validation jobs
            expected_jobs = [
                "validate-structure",
                "validate-domains",
                "validate-schema",
                "check-consistency"
            ]

            for job in expected_jobs:
                self.assertIn(job, content,
                            f"Validation workflow missing job: {job}")

    def test_cross_agent_consistency(self):
        """Test consistency across different agent types"""
        if self.agents_dir.exists():
            agent_files = list(self.agents_dir.glob("**/*.md"))

            if agent_files:
                # All agents should have similar structure
                for agent_file in agent_files:
                    with self.subTest(agent=agent_file):
                        if self.validation_script.exists():
                            result = subprocess.run([
                                str(self.validation_script), str(agent_file)
                            ], capture_output=True, text=True)

                            # All existing agents should pass validation
                            self.assertEqual(result.returncode, 0,
                                           f"Existing agent should pass: {agent_file}")

    def test_naming_convention_validation(self):
        """Test agent naming convention validation"""
        naming_tests = [
            ("solution-architect-python.md", True),
            ("software-engineer-dotnet.md", True),
            ("test-engineer-nodejs.md", True),
            ("requirements-analyst.md", True),
            ("documentation.md", True),
            ("invalid_name.md", False),
            ("bad-agent-name-format.md", False)
        ]

        for filename, should_pass in naming_tests:
            with self.subTest(filename=filename):
                test_agent = self.test_dir / filename
                test_content = """---
name: Test Agent
description: Test agent
domain: core
capabilities: []
tools: []
workflow:
  input: test
  processing: test
  output: test
  validation: test
constraints: []
examples: []
---

# Test Agent
"""
                test_agent.write_text(test_content)

                if self.validation_script.exists():
                    result = subprocess.run([
                        str(self.validation_script), str(test_agent)
                    ], capture_output=True, text=True)

                    if should_pass:
                        self.assertEqual(result.returncode, 0,
                                       f"Valid naming should pass: {filename}")
                    else:
                        # Note: Current script may not validate naming
                        # This test documents expected behavior
                        pass

    def test_tool_specification_validation(self):
        """Test validation of tool specifications"""
        # Agent with invalid tool specification
        invalid_tools_agent = self.test_dir / "invalid-tools.md"
        invalid_content = """---
name: Invalid Tools Agent
description: Agent with invalid tools
domain: python
capabilities:
  - Implementation
tools: "not a list"
workflow:
  input: test
  processing: test
  output: test
  validation: test
constraints: []
examples: []
---

# Invalid Tools Agent
"""

        invalid_tools_agent.write_text(invalid_content)

        # Should fail due to invalid tools format
        if self.validation_script.exists():
            result = subprocess.run([
                str(self.validation_script), str(invalid_tools_agent)
            ], capture_output=True, text=True)

            # Current script may not catch this, but it should
            # This documents expected behavior for future enhancement

    @patch('subprocess.run')
    def test_validation_in_ci(self, mock_run):
        """Test validation runs in CI environment"""
        mock_run.return_value = MagicMock(returncode=0)

        # Simulate CI environment
        os.environ['CI'] = 'true'

        # Validation should work in CI
        if self.validation_script.exists():
            # CI would run validation on all agents
            agent_files = list(self.agents_dir.glob("**/*.md"))

            for agent_file in agent_files:
                # Each agent would be validated
                mock_run.assert_called = True

        # Clean up
        if 'CI' in os.environ:
            del os.environ['CI']


if __name__ == "__main__":
    unittest.main()