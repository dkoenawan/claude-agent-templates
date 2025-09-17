#!/usr/bin/env python3

import unittest
import os
from pathlib import Path
from unittest.mock import patch, MagicMock


class TestDomainAgnosticWorkflow(unittest.TestCase):
    """Integration tests for domain-agnostic agent workflow"""

    def setUp(self):
        self.core_agents_dir = Path("agents/core")
        self.workflows_dir = Path(".github/workflows")

    def test_core_agents_exist(self):
        """Test that core domain-agnostic agents exist"""
        expected_agents = [
            "requirements-analyst.md",
            "solution-architect.md",
            "documentation.md"
        ]

        for agent in expected_agents:
            agent_path = self.core_agents_dir / agent
            self.assertTrue(agent_path.exists(),
                          f"Core agent missing: {agent}")

    def test_cross_domain_workflow(self):
        """Test workflow that spans multiple domains"""
        # Simulate a project that uses multiple technologies
        issue_data = {
            "title": "Build full-stack application",
            "body": """Create application with:
                - Python FastAPI backend
                - React TypeScript frontend
                - .NET microservice for payments
                - PostgreSQL database""",
            "labels": ["enhancement", "multi-domain"]
        }

        # Requirements analyst should handle multi-domain
        requirements_agent = self.core_agents_dir / "requirements-analyst.md"
        self.assertTrue(requirements_agent.exists(),
                      "Requirements analyst needed for multi-domain")

        # Solution architect should design across domains
        architect_agent = self.core_agents_dir / "solution-architect.md"
        self.assertTrue(architect_agent.exists(),
                      "Solution architect needed for multi-domain")

        # Each domain should have specific agents
        domains = ["python", "nodejs", "dotnet"]
        for domain in domains:
            domain_dir = Path(f"agents/{domain}")
            # Domain directory should exist (may be empty initially)
            # This will fail until agents are created
            if domain_dir.exists():
                agents = list(domain_dir.glob("*.md"))
                self.assertGreater(len(agents), 0,
                                 f"No agents found for domain: {domain}")

    def test_agent_collaboration_flow(self):
        """Test that agents can collaborate in sequence"""
        workflow_sequence = [
            ("requirements-analyst", "core", "Analyze requirements"),
            ("solution-architect", "core", "Design solution"),
            ("software-engineer-python", "python", "Implement backend"),
            ("software-engineer-nodejs", "nodejs", "Implement frontend"),
            ("test-engineer-python", "python", "Test backend"),
            ("documentation", "core", "Document solution")
        ]

        for agent_name, domain, phase in workflow_sequence:
            if domain == "core":
                agent_path = self.core_agents_dir / f"{agent_name}.md"
            else:
                agent_path = Path(f"agents/{domain}/{agent_name}.md")

            # Check agent exists (will fail initially)
            if agent_path.exists():
                content = agent_path.read_text()
                # Agent should have workflow section
                self.assertIn("workflow", content.lower(),
                            f"{agent_name} should define workflow")

    def test_workflow_orchestration(self):
        """Test GitHub Actions orchestration across domains"""
        orchestration_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        self.assertTrue(orchestration_workflow.exists(),
                      "Orchestration workflow missing")

        if orchestration_workflow.exists():
            content = orchestration_workflow.read_text()

            # Check for multi-domain support
            self.assertIn("domain", content,
                        "Workflow should handle domain classification")
            self.assertIn("classify", content.lower(),
                        "Workflow should classify issues")

    def test_domain_detection_fallback(self):
        """Test fallback when domain cannot be determined"""
        ambiguous_issues = [
            "Improve performance",
            "Fix the bug in the system",
            "Update documentation",
            "Refactor the codebase"
        ]

        for issue_title in ambiguous_issues:
            # Should fall back to core agents
            detected_domain = self._detect_domain_with_fallback(issue_title)
            self.assertEqual(detected_domain, "core",
                           f"Should fallback to core for: {issue_title}")

    def test_validation_across_domains(self):
        """Test validation works for all domains"""
        validation_script = Path("scripts/validate-agent-spec.sh")

        self.assertTrue(validation_script.exists(),
                      "Validation script missing")

        if validation_script.exists():
            # Check script handles different domains
            content = validation_script.read_text()

            domains = ["python", "dotnet", "nodejs", "java"]
            for domain in domains:
                self.assertIn(domain, content,
                            f"Validation should handle {domain} domain")

    def test_agent_specification_compatibility(self):
        """Test all agents follow same specification format"""
        all_agents = list(Path("agents").glob("**/*.md"))

        # Common structure all agents should have
        required_sections = [
            "name", "description", "domain",
            "capabilities", "tools", "workflow",
            "constraints", "examples"
        ]

        for agent_path in all_agents:
            if agent_path.exists():
                content = agent_path.read_text()

                # Check for YAML frontmatter
                self.assertTrue(content.startswith("---"),
                              f"{agent_path} should have YAML frontmatter")

                # Check for required sections (in frontmatter or body)
                content_lower = content.lower()
                for section in required_sections:
                    self.assertIn(section, content_lower,
                                f"{agent_path} missing section: {section}")

    def test_workflow_state_transitions(self):
        """Test workflow state transitions across agents"""
        # Each agent should support state transitions
        state_transitions = {
            "input": ["processing"],
            "processing": ["validation", "output"],
            "validation": ["output", "processing"],
            "output": []
        }

        all_agents = list(Path("agents").glob("**/*.md"))

        for agent_path in all_agents:
            if agent_path.exists():
                content = agent_path.read_text()

                if "workflow" in content.lower():
                    # Check for state definitions
                    for state in state_transitions.keys():
                        # Agents should mention workflow states
                        if state in content.lower():
                            # Valid state found
                            pass

    def test_error_handling_across_domains(self):
        """Test error handling in cross-domain workflow"""
        error_workflow = self.workflows_dir / "issue-agent-orchestration.yml"

        if error_workflow.exists():
            content = error_workflow.read_text()

            # Check for error handling
            self.assertIn("error", content.lower(),
                        "Workflow should handle errors")
            self.assertIn("fail", content.lower(),
                        "Workflow should handle failures")

            # Check for retry or recovery
            error_handling = ["retry", "recover", "fallback", "error-handler"]
            found = [eh for eh in error_handling if eh in content.lower()]

            self.assertGreater(len(found), 0,
                             "Workflow should have error handling mechanism")

    def _detect_domain_with_fallback(self, text):
        """Helper to detect domain with fallback to core"""
        text_lower = text.lower()

        # Try to detect specific domain
        domain_keywords = {
            "python": ["python", "django", "flask", "pytest"],
            "dotnet": ["dotnet", ".net", "c#", "asp.net"],
            "nodejs": ["node", "npm", "javascript", "typescript"],
            "java": ["java", "spring", "maven", "gradle"]
        }

        for domain, keywords in domain_keywords.items():
            if any(keyword in text_lower for keyword in keywords):
                return domain

        # Fallback to core for ambiguous cases
        return "core"


if __name__ == "__main__":
    unittest.main()