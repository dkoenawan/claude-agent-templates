#!/usr/bin/env python3
"""
Claude Agent Format Validator

Validates Claude agent specifications against the defined schema
and ensures consistency across all agents.
"""

import argparse
import json
import sys
import yaml
from pathlib import Path
from typing import Dict, List, Any, Optional
import re


class AgentValidator:
    """Validates Claude agent specifications"""

    REQUIRED_FIELDS = [
        'name', 'description', 'domain', 'role', 'spec_version',
        'tools', 'model', 'color', 'inputs', 'outputs', 'validation',
        'dependencies', 'workflow_position', 'github_integration', 'examples'
    ]

    VALID_DOMAINS = ['core', 'python', 'dotnet', 'nodejs', 'java']
    VALID_ROLES = ['analyst', 'architect', 'engineer', 'test-engineer', 'documentation']
    VALID_COLORS = ['blue', 'green', 'red', 'purple', 'orange', 'yellow']

    def __init__(self):
        self.errors = []
        self.warnings = []

    def validate_file(self, file_path: Path) -> bool:
        """Validate a single agent specification file"""
        try:
            content = file_path.read_text()

            # Check for YAML frontmatter
            if not content.startswith('---'):
                self.errors.append(f"{file_path}: Missing YAML frontmatter")
                return False

            # Extract and parse frontmatter
            parts = content.split('---', 2)
            if len(parts) < 3:
                self.errors.append(f"{file_path}: Invalid YAML frontmatter structure")
                return False

            try:
                frontmatter = yaml.safe_load(parts[1])
            except yaml.YAMLError as e:
                self.errors.append(f"{file_path}: Invalid YAML syntax: {e}")
                return False

            # Validate structure
            return self._validate_structure(file_path, frontmatter)

        except Exception as e:
            self.errors.append(f"{file_path}: Error reading file: {e}")
            return False

    def _validate_structure(self, file_path: Path, spec: Dict[str, Any]) -> bool:
        """Validate the agent specification structure"""
        valid = True

        # Check required fields
        for field in self.REQUIRED_FIELDS:
            if field not in spec:
                self.errors.append(f"{file_path}: Missing required field: {field}")
                valid = False

        # Validate specific fields
        if 'domain' in spec:
            if spec['domain'] not in self.VALID_DOMAINS:
                self.errors.append(f"{file_path}: Invalid domain: {spec['domain']}")
                valid = False

        if 'role' in spec:
            if spec['role'] not in self.VALID_ROLES:
                self.errors.append(f"{file_path}: Invalid role: {spec['role']}")
                valid = False

        if 'color' in spec:
            if spec['color'] not in self.VALID_COLORS:
                self.errors.append(f"{file_path}: Invalid color: {spec['color']}")
                valid = False

        # Validate tools
        if 'tools' in spec:
            tools = spec['tools']
            if not isinstance(tools, (list, str)):
                self.errors.append(f"{file_path}: Tools must be a list or string")
                valid = False
            elif isinstance(tools, str):
                # Tools as comma-separated string is acceptable
                pass
            elif isinstance(tools, list):
                # Tools as list is also acceptable
                pass

        # Validate workflow position
        if 'workflow_position' in spec:
            pos = spec['workflow_position']
            if not isinstance(pos, int) or pos < 1 or pos > 9:
                self.errors.append(f"{file_path}: Workflow position must be 1-9")
                valid = False

        # Validate GitHub integration
        if 'github_integration' in spec:
            github = spec['github_integration']
            if not isinstance(github, dict):
                self.errors.append(f"{file_path}: GitHub integration must be object")
                valid = False
            else:
                for required in ['triggers', 'outputs', 'permissions']:
                    if required not in github:
                        self.errors.append(f"{file_path}: Missing GitHub integration field: {required}")
                        valid = False

        # Validate examples
        if 'examples' in spec:
            examples = spec['examples']
            if not isinstance(examples, list):
                self.errors.append(f"{file_path}: Examples must be a list")
                valid = False
            else:
                for i, example in enumerate(examples):
                    if not isinstance(example, dict):
                        self.errors.append(f"{file_path}: Example {i} must be object")
                        valid = False
                        continue

                    for required in ['context', 'input', 'output']:
                        if required not in example:
                            self.errors.append(f"{file_path}: Example {i} missing field: {required}")
                            valid = False

        # Validate naming convention
        self._validate_naming(file_path, spec)

        return valid

    def _validate_naming(self, file_path: Path, spec: Dict[str, Any]):
        """Validate agent naming conventions"""
        name = spec.get('name', '')
        domain = spec.get('domain', '')
        role = spec.get('role', '')

        # Check file name matches agent name
        expected_filename = f"{name}.md"
        if file_path.name != expected_filename:
            self.warnings.append(f"{file_path}: Filename should be {expected_filename}")

        # Check domain-specific naming
        if domain != 'core':
            if not name.endswith(f"-{domain}"):
                self.warnings.append(f"{file_path}: Domain-specific agent should end with -{domain}")

        # Check role consistency
        if role in name and not name.startswith(role):
            self.warnings.append(f"{file_path}: Agent name should start with role: {role}")

    def validate_consistency(self, agents: List[Path]) -> bool:
        """Validate consistency across all agents"""
        valid = True
        agent_specs = {}

        # Load all agent specs
        for agent_path in agents:
            try:
                content = agent_path.read_text()
                parts = content.split('---', 2)
                if len(parts) >= 3:
                    spec = yaml.safe_load(parts[1])
                    agent_specs[agent_path] = spec
            except Exception:
                continue

        # Check for duplicate names
        names = {}
        for agent_path, spec in agent_specs.items():
            name = spec.get('name')
            if name:
                if name in names:
                    self.errors.append(f"Duplicate agent name: {name}")
                    self.errors.append(f"  Found in: {agent_path}")
                    self.errors.append(f"  Already in: {names[name]}")
                    valid = False
                else:
                    names[name] = agent_path

        # Check domain coverage
        domains = set()
        for spec in agent_specs.values():
            domains.add(spec.get('domain'))

        for expected_domain in self.VALID_DOMAINS:
            if expected_domain not in domains:
                self.warnings.append(f"No agents found for domain: {expected_domain}")

        # Check role coverage per domain
        domain_roles = {}
        for spec in agent_specs.values():
            domain = spec.get('domain')
            role = spec.get('role')
            if domain and role:
                if domain not in domain_roles:
                    domain_roles[domain] = set()
                domain_roles[domain].add(role)

        # Expected roles for non-core domains
        expected_roles = {'architect', 'engineer', 'test-engineer'}
        for domain in ['python', 'dotnet', 'nodejs', 'java']:
            if domain in domain_roles:
                missing_roles = expected_roles - domain_roles[domain]
                if missing_roles:
                    self.warnings.append(f"Domain {domain} missing roles: {missing_roles}")

        return valid


def main():
    parser = argparse.ArgumentParser(description='Validate Claude agent specifications')
    parser.add_argument('paths', nargs='*', help='Agent files or directories to validate')
    parser.add_argument('--format', choices=['text', 'json'], default='text',
                       help='Output format')
    parser.add_argument('--strict', action='store_true',
                       help='Treat warnings as errors')

    args = parser.parse_args()

    # Find agent files
    agent_files = []
    if not args.paths:
        # Default to agents directory
        agents_dir = Path('agents')
        if agents_dir.exists():
            agent_files = list(agents_dir.glob('**/*.md'))
    else:
        for path_str in args.paths:
            path = Path(path_str)
            if path.is_file() and path.suffix == '.md':
                agent_files.append(path)
            elif path.is_dir():
                agent_files.extend(path.glob('**/*.md'))

    if not agent_files:
        print("No agent files found", file=sys.stderr)
        return 1

    # Validate agents
    validator = AgentValidator()
    all_valid = True

    for agent_file in agent_files:
        if not validator.validate_file(agent_file):
            all_valid = False

    # Validate consistency
    if not validator.validate_consistency(agent_files):
        all_valid = False

    # Check if warnings should be treated as errors
    if args.strict and validator.warnings:
        all_valid = False

    # Output results
    if args.format == 'json':
        result = {
            'valid': all_valid,
            'errors': validator.errors,
            'warnings': validator.warnings,
            'files_checked': len(agent_files)
        }
        print(json.dumps(result, indent=2))
    else:
        # Text output
        if validator.errors:
            print("Errors:")
            for error in validator.errors:
                print(f"  ❌ {error}")

        if validator.warnings:
            print("Warnings:")
            for warning in validator.warnings:
                print(f"  ⚠️  {warning}")

        if all_valid:
            print(f"✅ All {len(agent_files)} agent specifications are valid!")
        else:
            print(f"❌ Validation failed for agent specifications")

    return 0 if all_valid else 1


if __name__ == '__main__':
    sys.exit(main())