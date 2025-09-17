#!/usr/bin/env python3

import unittest
import os
import yaml
import markdown
from pathlib import Path


class TestClaudeAgentFormat(unittest.TestCase):
    """Contract tests for Claude agent specification format"""

    def setUp(self):
        self.agents_dir = Path("agents")
        self.required_sections = [
            "name",
            "description",
            "domain",
            "capabilities",
            "tools",
            "workflow",
            "constraints",
            "examples"
        ]

    def test_agent_specification_structure(self):
        """Test that agent specs have required structure"""
        # This should FAIL initially as agents don't exist yet
        agent_files = list(self.agents_dir.glob("**/*.md"))

        self.assertGreater(len(agent_files), 0, "No agent specifications found")

        for agent_file in agent_files:
            with self.subTest(agent=agent_file):
                content = agent_file.read_text()

                # Check for YAML frontmatter
                self.assertTrue(content.startswith("---"),
                              f"{agent_file} missing YAML frontmatter")

                # Extract frontmatter
                parts = content.split("---", 2)
                if len(parts) >= 3:
                    frontmatter = yaml.safe_load(parts[1])

                    # Validate required fields
                    for field in self.required_sections:
                        self.assertIn(field, frontmatter,
                                    f"{agent_file} missing required field: {field}")

    def test_agent_naming_convention(self):
        """Test agent file naming follows convention"""
        # Pattern: {role}-{domain}.md or {role}.md for core agents
        agent_files = list(self.agents_dir.glob("**/*.md"))

        for agent_file in agent_files:
            with self.subTest(agent=agent_file):
                name = agent_file.stem

                # Check naming pattern
                self.assertTrue("-" in name or agent_file.parent.name == "core",
                              f"{agent_file} doesn't follow naming convention")

                # Check valid roles
                valid_roles = ["solution-architect", "software-engineer",
                             "test-engineer", "requirements-analyst", "documentation"]

                role = name.split("-")[0] + "-" + name.split("-")[1] if "-" in name else name
                self.assertIn(role, valid_roles + [name],
                            f"{agent_file} has invalid role: {role}")

    def test_tool_specifications(self):
        """Test that tools are properly specified"""
        agent_files = list(self.agents_dir.glob("**/*.md"))

        for agent_file in agent_files:
            with self.subTest(agent=agent_file):
                content = agent_file.read_text()

                # Parse frontmatter
                if "---" in content:
                    parts = content.split("---", 2)
                    if len(parts) >= 3:
                        frontmatter = yaml.safe_load(parts[1])

                        # Check tools section
                        self.assertIn("tools", frontmatter)
                        tools = frontmatter.get("tools", [])

                        self.assertIsInstance(tools, list,
                                            f"{agent_file} tools must be a list")
                        self.assertGreater(len(tools), 0,
                                         f"{agent_file} must specify at least one tool")

    def test_workflow_states(self):
        """Test workflow state definitions"""
        agent_files = list(self.agents_dir.glob("**/*.md"))
        required_states = ["input", "processing", "output", "validation"]

        for agent_file in agent_files:
            with self.subTest(agent=agent_file):
                content = agent_file.read_text()

                if "---" in content:
                    parts = content.split("---", 2)
                    if len(parts) >= 3:
                        frontmatter = yaml.safe_load(parts[1])

                        workflow = frontmatter.get("workflow", {})

                        for state in required_states:
                            self.assertIn(state, workflow,
                                        f"{agent_file} missing workflow state: {state}")

    def test_domain_specific_requirements(self):
        """Test domain-specific agent requirements"""
        domain_requirements = {
            "python": ["pytest", "python", "pip"],
            "dotnet": ["dotnet", "xunit", "nuget"],
            "nodejs": ["npm", "jest", "node"],
            "java": ["maven", "junit", "java"]
        }

        for domain, requirements in domain_requirements.items():
            domain_agents = list(Path(f"agents/{domain}").glob("*.md"))

            for agent_file in domain_agents:
                with self.subTest(agent=agent_file):
                    content = agent_file.read_text().lower()

                    # Check for domain-specific tools/keywords
                    found_requirements = [req for req in requirements
                                        if req in content]

                    self.assertGreater(len(found_requirements), 0,
                                     f"{agent_file} missing domain-specific requirements")

    def test_example_structure(self):
        """Test that examples are properly formatted"""
        agent_files = list(self.agents_dir.glob("**/*.md"))

        for agent_file in agent_files:
            with self.subTest(agent=agent_file):
                content = agent_file.read_text()

                if "---" in content:
                    parts = content.split("---", 2)
                    if len(parts) >= 3:
                        frontmatter = yaml.safe_load(parts[1])

                        examples = frontmatter.get("examples", [])

                        self.assertIsInstance(examples, list,
                                            f"{agent_file} examples must be a list")

                        if examples:
                            for example in examples:
                                self.assertIn("context", example,
                                            f"{agent_file} example missing context")
                                self.assertIn("input", example,
                                            f"{agent_file} example missing input")
                                self.assertIn("output", example,
                                            f"{agent_file} example missing output")


if __name__ == "__main__":
    unittest.main()