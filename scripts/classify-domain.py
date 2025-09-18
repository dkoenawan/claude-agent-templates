#!/usr/bin/env python3
"""
Domain Classifier

Analyzes GitHub issues and classifies them into appropriate technology domains
for agent assignment.
"""

import argparse
import json
import re
import sys
from typing import Dict, List, Tuple, Optional


class DomainClassifier:
    """Classifies GitHub issues into technology domains"""

    def __init__(self):
        # Domain-specific keywords and patterns
        self.domain_keywords = {
            'python': [
                'python', 'django', 'flask', 'fastapi', 'pytest', 'pip', 'pipenv',
                'poetry', 'pandas', 'numpy', 'scipy', 'jupyter', 'notebook',
                'virtualenv', 'conda', 'pyproject.toml', 'requirements.txt',
                'gunicorn', 'uvicorn', 'celery', 'sqlalchemy', 'pydantic'
            ],
            'dotnet': [
                'dotnet', '.net', 'c#', 'csharp', 'asp.net', 'aspnet', 'core',
                'entity framework', 'ef core', 'xunit', 'nunit', 'mstest',
                'nuget', 'visual studio', 'rider', 'blazor', 'razor',
                'web api', 'mvc', 'wpf', 'winforms', 'maui'
            ],
            'nodejs': [
                'node', 'nodejs', 'npm', 'yarn', 'pnpm', 'javascript', 'typescript',
                'express', 'fastify', 'next.js', 'react', 'vue', 'angular',
                'jest', 'mocha', 'cypress', 'playwright', 'webpack', 'vite',
                'package.json', 'tsconfig.json', 'eslint', 'prettier'
            ],
            'java': [
                'java', 'spring', 'spring boot', 'maven', 'gradle', 'junit',
                'testng', 'hibernate', 'jpa', 'tomcat', 'jetty', 'kafka',
                'elasticsearch', 'microservice', 'servlet', 'jsp', 'jsf',
                'pom.xml', 'build.gradle', 'application.properties'
            ]
        }

        # File extension patterns
        self.file_extensions = {
            'python': ['.py', '.pyx', '.pyi', '.ipynb'],
            'dotnet': ['.cs', '.vb', '.fs', '.csproj', '.sln', '.razor'],
            'nodejs': ['.js', '.ts', '.jsx', '.tsx', '.json', '.mjs'],
            'java': ['.java', '.kt', '.groovy', '.jar', '.war']
        }

        # Framework and tool patterns
        self.framework_patterns = {
            'python': [
                r'django[\s\-]?(?:rest|cms|admin)',
                r'flask[\s\-]?(?:app|api|web)',
                r'fastapi[\s\-]?(?:app|api)',
                r'pytest[\s\-]?(?:test|fixture)',
                r'sqlalchemy[\s\-]?(?:model|orm)'
            ],
            'dotnet': [
                r'asp\.net[\s\-]?(?:core|mvc|api)',
                r'entity\s+framework[\s\-]?(?:core)?',
                r'blazor[\s\-]?(?:server|wasm|app)',
                r'xunit[\s\-]?(?:test|fact)',
                r'\.net[\s\-]?(?:core|framework|8|7|6)'
            ],
            'nodejs': [
                r'express[\s\-]?(?:js|app|api)',
                r'next[\s\-]?(?:js|app)',
                r'react[\s\-]?(?:app|component)',
                r'typescript[\s\-]?(?:app|config)',
                r'node[\s\-]?(?:js|app|server)'
            ],
            'java': [
                r'spring[\s\-]?(?:boot|mvc|data|security)',
                r'maven[\s\-]?(?:project|build)',
                r'gradle[\s\-]?(?:build|project)',
                r'junit[\s\-]?(?:test|5|4)',
                r'hibernate[\s\-]?(?:jpa|orm)'
            ]
        }

        # Workflow phase keywords
        self.phase_keywords = {
            'planning': [
                'plan', 'design', 'architecture', 'requirements', 'specification',
                'analysis', 'breakdown', 'strategy', 'approach', 'proposal'
            ],
            'implementation': [
                'implement', 'build', 'create', 'develop', 'code', 'feature',
                'functionality', 'component', 'module', 'service', 'endpoint'
            ],
            'testing': [
                'test', 'testing', 'unit test', 'integration test', 'e2e',
                'coverage', 'mock', 'fixture', 'validation', 'verify'
            ],
            'documentation': [
                'document', 'documentation', 'readme', 'docs', 'api docs',
                'guide', 'tutorial', 'specification', 'manual'
            ]
        }

    def classify_issue(self, title: str, body: str, labels: List[str] = None) -> Dict[str, str]:
        """
        Classify a GitHub issue into domain and phase

        Returns:
            Dict with 'domain', 'phase', and 'agent' keys
        """
        if labels is None:
            labels = []

        # Combine all text for analysis
        text = f"{title} {body} {' '.join(labels)}".lower()

        # Classify domain
        domain = self._classify_domain(text)

        # Classify phase
        phase = self._classify_phase(text)

        # Determine appropriate agent
        agent = self._determine_agent(domain, phase)

        return {
            'domain': domain,
            'phase': phase,
            'agent': agent
        }

    def _classify_domain(self, text: str) -> str:
        """Classify the technology domain based on text content"""
        domain_scores = {}

        # Score based on keywords
        for domain, keywords in self.domain_keywords.items():
            score = 0
            for keyword in keywords:
                if keyword in text:
                    # Weight longer, more specific keywords higher
                    weight = len(keyword.split())
                    score += weight
            domain_scores[domain] = score

        # Score based on file extensions mentioned
        for domain, extensions in self.file_extensions.items():
            for ext in extensions:
                if ext in text:
                    domain_scores[domain] = domain_scores.get(domain, 0) + 2

        # Score based on framework patterns
        for domain, patterns in self.framework_patterns.items():
            for pattern in patterns:
                matches = len(re.findall(pattern, text, re.IGNORECASE))
                domain_scores[domain] = domain_scores.get(domain, 0) + (matches * 3)

        # Find highest scoring domain
        if domain_scores:
            best_domain = max(domain_scores.items(), key=lambda x: x[1])
            if best_domain[1] > 0:
                return best_domain[0]

        # Default to core for ambiguous cases
        return 'core'

    def _classify_phase(self, text: str) -> str:
        """Classify the workflow phase based on text content"""
        phase_scores = {}

        for phase, keywords in self.phase_keywords.items():
            score = 0
            for keyword in keywords:
                if keyword in text:
                    score += 1
            phase_scores[phase] = score

        # Find highest scoring phase
        if phase_scores:
            best_phase = max(phase_scores.items(), key=lambda x: x[1])
            if best_phase[1] > 0:
                return best_phase[0]

        # Default to planning for new issues
        return 'planning'

    def _determine_agent(self, domain: str, phase: str) -> str:
        """Determine the appropriate agent based on domain and phase"""
        if phase == 'planning':
            if domain == 'core':
                return 'solution-architect'
            else:
                return f'solution-architect-{domain}'
        elif phase == 'implementation':
            if domain == 'core':
                return 'software-engineer'  # Generic fallback
            else:
                return f'software-engineer-{domain}'
        elif phase == 'testing':
            if domain == 'core':
                return 'test-engineer'  # Generic fallback
            else:
                return f'test-engineer-{domain}'
        elif phase == 'documentation':
            return 'documentation'
        else:
            # Default to requirements analyst for unclear cases
            return 'requirements-analyst'

    def get_confidence_score(self, title: str, body: str, labels: List[str] = None) -> float:
        """Get confidence score for the classification (0.0 to 1.0)"""
        if labels is None:
            labels = []

        text = f"{title} {body} {' '.join(labels)}".lower()

        # Count total domain indicators
        total_indicators = 0
        for domain, keywords in self.domain_keywords.items():
            for keyword in keywords:
                if keyword in text:
                    total_indicators += 1

        # More indicators = higher confidence
        if total_indicators >= 3:
            return 0.9
        elif total_indicators >= 2:
            return 0.7
        elif total_indicators >= 1:
            return 0.5
        else:
            return 0.3


def main():
    parser = argparse.ArgumentParser(description='Classify GitHub issue domain and phase')
    parser.add_argument('--issue', type=int, help='GitHub issue number')
    parser.add_argument('--repo', help='GitHub repository (owner/repo)')
    parser.add_argument('--title', help='Issue title (for testing)')
    parser.add_argument('--body', help='Issue body (for testing)')
    parser.add_argument('--labels', nargs='*', help='Issue labels (for testing)')
    parser.add_argument('--format', choices=['text', 'json', 'github'], default='github',
                       help='Output format')

    args = parser.parse_args()

    classifier = DomainClassifier()

    if args.issue and args.repo:
        # Fetch issue from GitHub (would need GitHub API integration)
        print("GitHub API integration not implemented yet", file=sys.stderr)
        return 1
    elif args.title:
        # Test mode with provided data
        title = args.title
        body = args.body or ""
        labels = args.labels or []

        result = classifier.classify_issue(title, body, labels)
        confidence = classifier.get_confidence_score(title, body, labels)

        if args.format == 'json':
            output = {**result, 'confidence': confidence}
            print(json.dumps(output, indent=2))
        elif args.format == 'github':
            # GitHub Actions output format
            print(f"domain={result['domain']}")
            print(f"phase={result['phase']}")
            print(f"agent={result['agent']}")
            print(f"confidence={confidence}")
        else:
            print(f"Domain: {result['domain']}")
            print(f"Phase: {result['phase']}")
            print(f"Agent: {result['agent']}")
            print(f"Confidence: {confidence:.1%}")
    else:
        print("Either --issue and --repo, or --title must be provided", file=sys.stderr)
        return 1

    return 0


if __name__ == '__main__':
    sys.exit(main())