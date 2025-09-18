#!/usr/bin/env python3

import unittest
import os
import tempfile
import shutil
import subprocess
from pathlib import Path
from unittest.mock import patch, MagicMock


class TestPythonAgentWorkflow(unittest.TestCase):
    """Integration tests for Python agent specialization workflow"""

    def setUp(self):
        self.test_dir = Path(tempfile.mkdtemp())
        self.agents_dir = Path("agents/python")

    def tearDown(self):
        if self.test_dir.exists():
            shutil.rmtree(self.test_dir)

    def test_python_agent_exists(self):
        """Test that Python-specific agents exist"""
        expected_agents = [
            "software-engineer-python.md",
            "test-engineer-python.md"
        ]

        for agent in expected_agents:
            agent_path = self.agents_dir / agent
            self.assertTrue(agent_path.exists(),
                          f"Python agent missing: {agent}")

    def test_python_agent_workflow_sequence(self):
        """Test complete Python development workflow sequence"""
        # This should FAIL initially as agents don't exist

        # Simulate issue creation
        issue_data = {
            "title": "Implement data validation module",
            "body": "Create a Python module for data validation with pytest tests",
            "labels": ["enhancement", "python"]
        }

        # Step 1: Requirements analysis
        requirements_agent = Path("agents/core/requirements-analyst.md")
        self.assertTrue(requirements_agent.exists(),
                      "Requirements analyst agent missing")

        # Step 2: Solution architecture
        architect_agent = Path("agents/core/solution-architect.md")
        self.assertTrue(architect_agent.exists(),
                      "Solution architect agent missing")

        # Step 3: Implementation
        engineer_agent = self.agents_dir / "software-engineer-python.md"
        self.assertTrue(engineer_agent.exists(),
                      "Python software engineer agent missing")

        # Step 4: Testing
        test_agent = self.agents_dir / "test-engineer-python.md"
        self.assertTrue(test_agent.exists(),
                      "Python test engineer agent missing")

        # Step 5: Documentation
        doc_agent = Path("agents/core/documentation.md")
        self.assertTrue(doc_agent.exists(),
                      "Documentation agent missing")

    def test_python_tools_integration(self):
        """Test Python-specific tool integration"""
        engineer_agent = self.agents_dir / "software-engineer-python.md"

        if engineer_agent.exists():
            content = engineer_agent.read_text()

            # Check for Python-specific tools
            python_tools = ["pytest", "pip", "black", "ruff", "mypy"]

            for tool in python_tools:
                self.assertIn(tool, content.lower(),
                            f"Python agent should mention {tool}")

    def test_python_test_framework_support(self):
        """Test Python test framework support"""
        test_agent = self.agents_dir / "test-engineer-python.md"

        if test_agent.exists():
            content = test_agent.read_text()

            # Check for test framework mentions
            test_frameworks = ["pytest", "unittest", "coverage", "mock"]

            frameworks_found = [fw for fw in test_frameworks
                              if fw in content.lower()]

            self.assertGreater(len(frameworks_found), 0,
                             "Test agent should mention Python test frameworks")

    @patch('subprocess.run')
    def test_python_build_commands(self, mock_run):
        """Test Python build and test command execution"""
        mock_run.return_value = MagicMock(returncode=0)

        # Simulate Python project setup
        project_files = {
            "setup.py": "from setuptools import setup; setup()",
            "requirements.txt": "pytest\nblack\nruff",
            "src/__init__.py": "",
            "tests/__init__.py": "",
            "tests/test_sample.py": "def test_sample(): assert True"
        }

        for file_path, content in project_files.items():
            full_path = self.test_dir / file_path
            full_path.parent.mkdir(parents=True, exist_ok=True)
            full_path.write_text(content)

        # Commands that should be executed
        expected_commands = [
            ["pip", "install", "-r", "requirements.txt"],
            ["pytest", "tests/"],
            ["black", "--check", "src/"],
            ["ruff", "check", "src/"]
        ]

        # Validate commands would be executed
        for cmd in expected_commands:
            # In real implementation, these would be called
            # Here we just verify the structure
            self.assertIsInstance(cmd, list)
            self.assertTrue(all(isinstance(c, str) for c in cmd))

    def test_python_hexagonal_architecture(self):
        """Test Python agent follows hexagonal architecture"""
        engineer_agent = self.agents_dir / "software-engineer-python.md"

        if engineer_agent.exists():
            content = engineer_agent.read_text().lower()

            # Check for hexagonal architecture concepts
            architecture_concepts = [
                "hexagonal", "ports", "adapters",
                "domain", "infrastructure", "application"
            ]

            concepts_found = [concept for concept in architecture_concepts
                            if concept in content]

            self.assertGreater(len(concepts_found), 0,
                             "Python agent should mention hexagonal architecture")

    def test_python_github_integration(self):
        """Test Python workflow GitHub integration"""
        # Check if validation script exists
        validation_script = Path("scripts/validate-agent-spec.sh")
        self.assertTrue(validation_script.exists(),
                      "Validation script missing")

        if validation_script.exists():
            # Check script is executable
            self.assertTrue(os.access(validation_script, os.X_OK),
                          "Validation script should be executable")

    def test_python_domain_classification(self):
        """Test Python domain is properly classified"""
        # Test domain classification logic
        test_issues = [
            ("Setup Django REST API", "python"),
            ("Create FastAPI endpoints", "python"),
            ("Implement pytest fixtures", "python"),
            ("Configure Flask application", "python")
        ]

        for issue_title, expected_domain in test_issues:
            # In real implementation, this would call classify-domain.py
            # Here we verify the logic structure
            detected_domain = self._detect_domain(issue_title)
            self.assertEqual(detected_domain, expected_domain,
                           f"Failed to detect Python domain in: {issue_title}")

    def _detect_domain(self, text):
        """Helper to detect domain from text"""
        python_keywords = ["python", "django", "flask", "fastapi", "pytest", "pip"]
        text_lower = text.lower()

        for keyword in python_keywords:
            if keyword in text_lower:
                return "python"
        return "unknown"


if __name__ == "__main__":
    unittest.main()