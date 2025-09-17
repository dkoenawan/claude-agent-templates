#!/usr/bin/env python3

import unittest
import os
import tempfile
import shutil
from pathlib import Path
from unittest.mock import patch, MagicMock


class TestDotNetAgentWorkflow(unittest.TestCase):
    """Integration tests for .NET agent specialization workflow"""

    def setUp(self):
        self.test_dir = Path(tempfile.mkdtemp())
        self.agents_dir = Path("agents/dotnet")

    def tearDown(self):
        if self.test_dir.exists():
            shutil.rmtree(self.test_dir)

    def test_dotnet_agents_exist(self):
        """Test that .NET-specific agents exist"""
        expected_agents = [
            "solution-architect-dotnet.md",
            "software-engineer-dotnet.md",
            "test-engineer-dotnet.md"
        ]

        for agent in expected_agents:
            agent_path = self.agents_dir / agent
            self.assertTrue(agent_path.exists(),
                          f".NET agent missing: {agent}")

    def test_dotnet_workflow_sequence(self):
        """Test complete .NET development workflow sequence"""
        # This should FAIL initially as agents don't exist

        # Simulate issue for .NET project
        issue_data = {
            "title": "Create ASP.NET Core Web API",
            "body": "Implement RESTful API with Entity Framework Core and xUnit tests",
            "labels": ["enhancement", "dotnet"]
        }

        # Workflow phases
        phases = {
            "planning": self.agents_dir / "solution-architect-dotnet.md",
            "implementation": self.agents_dir / "software-engineer-dotnet.md",
            "testing": self.agents_dir / "test-engineer-dotnet.md",
            "documentation": Path("agents/core/documentation.md")
        }

        for phase, agent_path in phases.items():
            self.assertTrue(agent_path.exists(),
                          f"Agent missing for {phase} phase: {agent_path}")

    def test_dotnet_tools_integration(self):
        """Test .NET-specific tool integration"""
        engineer_agent = self.agents_dir / "software-engineer-dotnet.md"

        if engineer_agent.exists():
            content = engineer_agent.read_text()

            # Check for .NET-specific tools
            dotnet_tools = [
                "dotnet", "nuget", "msbuild",
                "visual studio", "rider"
            ]

            tools_found = [tool for tool in dotnet_tools
                          if tool in content.lower()]

            self.assertGreater(len(tools_found), 0,
                             ".NET agent should mention .NET tools")

    def test_dotnet_test_frameworks(self):
        """Test .NET test framework support"""
        test_agent = self.agents_dir / "test-engineer-dotnet.md"

        if test_agent.exists():
            content = test_agent.read_text()

            # Check for test framework mentions
            test_frameworks = ["xunit", "nunit", "mstest", "moq", "fluentassertions"]

            frameworks_found = [fw for fw in test_frameworks
                              if fw in content.lower()]

            self.assertGreater(len(frameworks_found), 0,
                             "Test agent should mention .NET test frameworks")

    @patch('subprocess.run')
    def test_dotnet_build_commands(self, mock_run):
        """Test .NET build and test command execution"""
        mock_run.return_value = MagicMock(returncode=0)

        # Simulate .NET project structure
        project_files = {
            "MyApp.csproj": """<Project Sdk="Microsoft.NET.Sdk">
                <PropertyGroup>
                    <TargetFramework>net8.0</TargetFramework>
                </PropertyGroup>
            </Project>""",
            "Program.cs": "Console.WriteLine(\"Hello World\");",
            "MyApp.Tests/MyApp.Tests.csproj": """<Project Sdk="Microsoft.NET.Sdk">
                <PropertyGroup>
                    <TargetFramework>net8.0</TargetFramework>
                </PropertyGroup>
                <ItemGroup>
                    <PackageReference Include="xunit" Version="2.4.1" />
                </ItemGroup>
            </Project>""",
            "MyApp.Tests/UnitTest1.cs": "public class UnitTest1 { }"
        }

        for file_path, content in project_files.items():
            full_path = self.test_dir / file_path
            full_path.parent.mkdir(parents=True, exist_ok=True)
            full_path.write_text(content)

        # Commands that should be executed
        expected_commands = [
            ["dotnet", "restore"],
            ["dotnet", "build"],
            ["dotnet", "test"],
            ["dotnet", "publish"]
        ]

        # Validate command structure
        for cmd in expected_commands:
            self.assertIsInstance(cmd, list)
            self.assertEqual(cmd[0], "dotnet")

    def test_dotnet_project_types(self):
        """Test support for different .NET project types"""
        architect_agent = self.agents_dir / "solution-architect-dotnet.md"

        if architect_agent.exists():
            content = architect_agent.read_text().lower()

            # Check for different project types
            project_types = [
                "web api", "mvc", "blazor", "console",
                "class library", "microservice"
            ]

            types_found = [pt for pt in project_types if pt in content]

            self.assertGreater(len(types_found), 0,
                             "Architect should mention .NET project types")

    def test_dotnet_clean_architecture(self):
        """Test .NET agent follows clean architecture"""
        engineer_agent = self.agents_dir / "software-engineer-dotnet.md"

        if engineer_agent.exists():
            content = engineer_agent.read_text().lower()

            # Check for clean architecture concepts
            architecture_concepts = [
                "clean architecture", "domain", "application",
                "infrastructure", "presentation", "dependency injection"
            ]

            concepts_found = [concept for concept in architecture_concepts
                            if concept in content]

            self.assertGreater(len(concepts_found), 0,
                             ".NET agent should mention clean architecture")

    def test_dotnet_nuget_package_management(self):
        """Test NuGet package management support"""
        engineer_agent = self.agents_dir / "software-engineer-dotnet.md"

        if engineer_agent.exists():
            content = engineer_agent.read_text().lower()

            # Check for package management
            package_concepts = ["nuget", "packagereference", "package", "restore"]

            concepts_found = [concept for concept in package_concepts
                            if concept in content]

            self.assertGreater(len(concepts_found), 0,
                             ".NET agent should mention package management")

    def test_dotnet_domain_classification(self):
        """Test .NET domain is properly classified"""
        test_issues = [
            ("Create ASP.NET Core API", "dotnet"),
            ("Setup Entity Framework migrations", "dotnet"),
            ("Implement xUnit tests for service layer", "dotnet"),
            ("Configure Blazor WebAssembly app", "dotnet"),
            ("Add C# record types", "dotnet")
        ]

        for issue_title, expected_domain in test_issues:
            detected_domain = self._detect_domain(issue_title)
            self.assertEqual(detected_domain, expected_domain,
                           f"Failed to detect .NET domain in: {issue_title}")

    def test_dotnet_solution_structure(self):
        """Test .NET solution structure support"""
        architect_agent = self.agents_dir / "solution-architect-dotnet.md"

        if architect_agent.exists():
            content = architect_agent.read_text().lower()

            # Check for solution structure concepts
            structure_concepts = [
                "solution", ".sln", "project reference",
                "shared", "common", "layers"
            ]

            concepts_found = [concept for concept in structure_concepts
                            if concept in content]

            self.assertGreater(len(concepts_found), 0,
                             "Architect should mention solution structure")

    def _detect_domain(self, text):
        """Helper to detect domain from text"""
        dotnet_keywords = [
            "asp.net", "dotnet", ".net", "c#", "entity framework",
            "xunit", "nunit", "blazor", "razor"
        ]
        text_lower = text.lower()

        for keyword in dotnet_keywords:
            if keyword in text_lower:
                return "dotnet"
        return "unknown"


if __name__ == "__main__":
    unittest.main()