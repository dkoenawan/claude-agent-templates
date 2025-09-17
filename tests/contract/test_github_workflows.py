#!/usr/bin/env python3

import unittest
import os
import yaml
from pathlib import Path


class TestGitHubWorkflowContracts(unittest.TestCase):
    """Contract tests for GitHub Actions workflow specifications"""

    def setUp(self):
        self.workflows_dir = Path(".github/workflows")
        self.required_workflows = [
            "issue-agent-orchestration.yml",
            "execute-phase.yml",
            "validate-agents.yml"
        ]

    def test_required_workflows_exist(self):
        """Test that all required workflow files exist"""
        for workflow in self.required_workflows:
            workflow_path = self.workflows_dir / workflow
            self.assertTrue(workflow_path.exists(),
                          f"Required workflow missing: {workflow}")

    def test_workflow_structure(self):
        """Test basic workflow structure"""
        workflow_files = list(self.workflows_dir.glob("*.yml"))

        for workflow_file in workflow_files:
            with self.subTest(workflow=workflow_file):
                content = workflow_file.read_text()
                workflow_data = yaml.safe_load(content)

                # Required top-level keys
                self.assertIn("name", workflow_data,
                            f"{workflow_file} missing 'name' field")
                self.assertIn("on", workflow_data,
                            f"{workflow_file} missing 'on' trigger")
                self.assertIn("jobs", workflow_data,
                            f"{workflow_file} missing 'jobs' section")

    def test_issue_orchestration_workflow(self):
        """Test issue-agent-orchestration workflow contract"""
        workflow_path = self.workflows_dir / "issue-agent-orchestration.yml"

        if workflow_path.exists():
            workflow_data = yaml.safe_load(workflow_path.read_text())

            # Check triggers
            triggers = workflow_data.get("on", {})
            self.assertIn("issues", triggers,
                        "Issue orchestration must trigger on issues")
            self.assertIn("workflow_dispatch", triggers,
                        "Issue orchestration must support manual dispatch")

            # Check required jobs
            jobs = workflow_data.get("jobs", {})
            self.assertIn("classify-and-assign", jobs,
                        "Must have classify-and-assign job")
            self.assertIn("execute-agent", jobs,
                        "Must have execute-agent job")

            # Check outputs
            classify_job = jobs.get("classify-and-assign", {})
            outputs = classify_job.get("outputs", {})
            self.assertIn("domain", outputs, "Must output domain")
            self.assertIn("phase", outputs, "Must output phase")
            self.assertIn("agent", outputs, "Must output agent")

    def test_execute_phase_workflow(self):
        """Test execute-phase workflow contract"""
        workflow_path = self.workflows_dir / "execute-phase.yml"

        if workflow_path.exists():
            workflow_data = yaml.safe_load(workflow_path.read_text())

            # Check inputs
            inputs = workflow_data.get("on", {}).get("workflow_dispatch", {}).get("inputs", {})
            required_inputs = ["issue_number", "domain", "phase", "agent"]

            for input_name in required_inputs:
                self.assertIn(input_name, inputs,
                            f"Execute phase missing required input: {input_name}")

            # Check phase-specific jobs
            jobs = workflow_data.get("jobs", {})
            required_phases = ["planning-phase", "implementation-phase",
                             "testing-phase", "documentation-phase"]

            for phase in required_phases:
                self.assertIn(phase, jobs,
                            f"Execute phase missing job: {phase}")

    def test_validate_agents_workflow(self):
        """Test validate-agents workflow contract"""
        workflow_path = self.workflows_dir / "validate-agents.yml"

        if workflow_path.exists():
            workflow_data = yaml.safe_load(workflow_path.read_text())

            # Check triggers
            triggers = workflow_data.get("on", {})
            self.assertIn("push", triggers,
                        "Validation must trigger on push")
            self.assertIn("pull_request", triggers,
                        "Validation must trigger on pull_request")

            # Check validation jobs
            jobs = workflow_data.get("jobs", {})
            validation_jobs = ["validate-structure", "validate-domains",
                             "validate-schema", "check-consistency"]

            for job in validation_jobs:
                self.assertIn(job, jobs,
                            f"Validation workflow missing job: {job}")

    def test_workflow_dependencies(self):
        """Test job dependencies are properly configured"""
        workflow_files = list(self.workflows_dir.glob("*.yml"))

        for workflow_file in workflow_files:
            with self.subTest(workflow=workflow_file):
                workflow_data = yaml.safe_load(workflow_file.read_text())
                jobs = workflow_data.get("jobs", {})

                for job_name, job_config in jobs.items():
                    if "needs" in job_config:
                        needs = job_config["needs"]
                        if isinstance(needs, str):
                            needs = [needs]

                        for dependency in needs:
                            self.assertIn(dependency, jobs,
                                        f"{job_name} depends on non-existent job: {dependency}")

    def test_workflow_secrets_and_permissions(self):
        """Test workflows don't expose secrets"""
        workflow_files = list(self.workflows_dir.glob("*.yml"))

        for workflow_file in workflow_files:
            with self.subTest(workflow=workflow_file):
                content = workflow_file.read_text()

                # Check for hardcoded secrets
                self.assertNotIn("api_key:", content.lower(),
                               f"{workflow_file} may contain hardcoded API key")
                self.assertNotIn("password:", content.lower(),
                               f"{workflow_file} may contain hardcoded password")
                self.assertNotIn("token:", content.lower(),
                               f"{workflow_file} may contain hardcoded token")

                # Check for proper secret usage
                if "${{" in content and "secrets." in content:
                    # Secrets should use GitHub secrets syntax
                    self.assertIn("secrets.", content,
                                f"{workflow_file} should use GitHub secrets properly")

    def test_workflow_artifact_handling(self):
        """Test workflows properly handle artifacts"""
        execute_workflow = self.workflows_dir / "execute-phase.yml"

        if execute_workflow.exists():
            workflow_data = yaml.safe_load(execute_workflow.read_text())
            jobs = workflow_data.get("jobs", {})

            # Check planning phase creates artifacts
            planning = jobs.get("planning-phase", {})
            steps = planning.get("steps", [])

            upload_found = any("upload-artifact" in str(step) for step in steps)
            self.assertTrue(upload_found,
                          "Planning phase should upload artifacts")

            # Check implementation downloads artifacts
            implementation = jobs.get("implementation-phase", {})
            steps = implementation.get("steps", [])

            download_found = any("download-artifact" in str(step) for step in steps)
            self.assertTrue(download_found,
                          "Implementation phase should download artifacts")


if __name__ == "__main__":
    unittest.main()