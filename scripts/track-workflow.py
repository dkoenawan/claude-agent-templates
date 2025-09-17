#!/usr/bin/env python3
"""
Workflow State Tracker

Tracks and manages workflow state for GitHub issues through the 9-step
development process.
"""

import argparse
import json
import sys
from typing import Dict, List, Optional, Tuple
from datetime import datetime, timezone
from pathlib import Path


class WorkflowTracker:
    """Tracks workflow state for GitHub issues"""

    # 9-step workflow states
    WORKFLOW_STATES = [
        'issue-created',           # 1. Initial issue creation
        'requirements-analysis',   # 2. Requirements analyst processing
        'requirements-ready',      # 3. Requirements completed
        'plan-approved',          # 4. Solution architect completed
        'tests-planned',          # 5. Test engineer completed
        'implementation-ready',    # 6. Ready for software engineer
        'implementation-complete', # 7. Software engineer completed
        'user-accepted',          # 8. User has accepted implementation
        'documentation-complete'   # 9. Documentation agent completed
    ]

    # State transitions (from -> to)
    VALID_TRANSITIONS = {
        'issue-created': ['requirements-analysis'],
        'requirements-analysis': ['requirements-ready'],
        'requirements-ready': ['plan-approved'],
        'plan-approved': ['tests-planned'],
        'tests-planned': ['implementation-ready'],
        'implementation-ready': ['implementation-complete'],
        'implementation-complete': ['user-accepted'],
        'user-accepted': ['documentation-complete'],
        'documentation-complete': []  # Final state
    }

    # Agents responsible for each transition
    STATE_AGENTS = {
        'issue-created': None,
        'requirements-analysis': 'requirements-analyst',
        'requirements-ready': 'requirements-analyst',
        'plan-approved': 'solution-architect',
        'tests-planned': 'test-engineer',
        'implementation-ready': 'test-engineer',
        'implementation-complete': 'software-engineer',
        'user-accepted': 'user',
        'documentation-complete': 'documentation'
    }

    def __init__(self, state_file: str = '.workflow-state.json'):
        self.state_file = Path(state_file)
        self.state_data = self._load_state()

    def _load_state(self) -> Dict:
        """Load workflow state from file"""
        if self.state_file.exists():
            try:
                with open(self.state_file, 'r') as f:
                    return json.load(f)
            except (json.JSONDecodeError, IOError):
                pass
        return {'issues': {}, 'metadata': {'version': '1.0'}}

    def _save_state(self):
        """Save workflow state to file"""
        try:
            with open(self.state_file, 'w') as f:
                json.dump(self.state_data, f, indent=2, default=str)
        except IOError as e:
            print(f"Error saving state: {e}", file=sys.stderr)

    def get_issue_state(self, issue_number: int) -> Optional[Dict]:
        """Get current state for an issue"""
        return self.state_data['issues'].get(str(issue_number))

    def set_issue_state(self, issue_number: int, state: str, agent: str = None,
                       metadata: Dict = None) -> bool:
        """
        Set the workflow state for an issue

        Returns True if transition is valid, False otherwise
        """
        issue_key = str(issue_number)
        current_issue = self.state_data['issues'].get(issue_key, {})
        current_state = current_issue.get('current_state')

        # Validate transition
        if current_state and not self._is_valid_transition(current_state, state):
            return False

        # Create or update issue entry
        now = datetime.now(timezone.utc).isoformat()

        if issue_key not in self.state_data['issues']:
            self.state_data['issues'][issue_key] = {
                'issue_number': issue_number,
                'created_at': now,
                'current_state': state,
                'history': [],
                'metadata': metadata or {}
            }
        else:
            # Add current state to history
            if current_state:
                self.state_data['issues'][issue_key]['history'].append({
                    'state': current_state,
                    'agent': self.state_data['issues'][issue_key].get('current_agent'),
                    'timestamp': self.state_data['issues'][issue_key].get('updated_at'),
                    'duration': self._calculate_duration(
                        self.state_data['issues'][issue_key].get('updated_at'),
                        now
                    )
                })

        # Update current state
        self.state_data['issues'][issue_key].update({
            'current_state': state,
            'current_agent': agent,
            'updated_at': now
        })

        if metadata:
            self.state_data['issues'][issue_key]['metadata'].update(metadata)

        self._save_state()
        return True

    def _is_valid_transition(self, from_state: str, to_state: str) -> bool:
        """Check if state transition is valid"""
        if from_state not in self.VALID_TRANSITIONS:
            return False
        return to_state in self.VALID_TRANSITIONS[from_state]

    def _calculate_duration(self, start_time: str, end_time: str) -> Optional[float]:
        """Calculate duration between two timestamps in seconds"""
        try:
            start = datetime.fromisoformat(start_time.replace('Z', '+00:00'))
            end = datetime.fromisoformat(end_time.replace('Z', '+00:00'))
            return (end - start).total_seconds()
        except (ValueError, AttributeError):
            return None

    def get_next_states(self, current_state: str) -> List[str]:
        """Get valid next states for current state"""
        return self.VALID_TRANSITIONS.get(current_state, [])

    def get_workflow_progress(self, issue_number: int) -> Dict:
        """Get workflow progress for an issue"""
        issue = self.get_issue_state(issue_number)
        if not issue:
            return {'progress': 0, 'current_step': 0, 'total_steps': len(self.WORKFLOW_STATES)}

        current_state = issue['current_state']
        try:
            current_step = self.WORKFLOW_STATES.index(current_state) + 1
        except ValueError:
            current_step = 0

        progress = (current_step / len(self.WORKFLOW_STATES)) * 100

        return {
            'progress': round(progress, 1),
            'current_step': current_step,
            'total_steps': len(self.WORKFLOW_STATES),
            'current_state': current_state,
            'next_states': self.get_next_states(current_state)
        }

    def get_blocked_issues(self) -> List[Dict]:
        """Get issues that haven't progressed recently"""
        blocked = []
        now = datetime.now(timezone.utc)

        for issue_key, issue in self.state_data['issues'].items():
            updated_at = issue.get('updated_at')
            if updated_at:
                try:
                    last_update = datetime.fromisoformat(updated_at.replace('Z', '+00:00'))
                    hours_since_update = (now - last_update).total_seconds() / 3600

                    # Consider blocked if no progress in 24 hours and not in final state
                    if (hours_since_update > 24 and
                        issue['current_state'] != 'documentation-complete'):
                        blocked.append({
                            'issue_number': issue['issue_number'],
                            'current_state': issue['current_state'],
                            'hours_stalled': round(hours_since_update, 1),
                            'last_agent': issue.get('current_agent')
                        })
                except ValueError:
                    continue

        return sorted(blocked, key=lambda x: x['hours_stalled'], reverse=True)

    def generate_report(self) -> Dict:
        """Generate workflow status report"""
        total_issues = len(self.state_data['issues'])
        state_counts = {}
        completed_issues = 0

        for issue in self.state_data['issues'].values():
            state = issue['current_state']
            state_counts[state] = state_counts.get(state, 0) + 1

            if state == 'documentation-complete':
                completed_issues += 1

        # Calculate average duration per state
        state_durations = {}
        for issue in self.state_data['issues'].values():
            for history_entry in issue.get('history', []):
                state = history_entry['state']
                duration = history_entry.get('duration')
                if duration:
                    if state not in state_durations:
                        state_durations[state] = []
                    state_durations[state].append(duration)

        avg_durations = {}
        for state, durations in state_durations.items():
            avg_durations[state] = sum(durations) / len(durations) / 3600  # Convert to hours

        return {
            'total_issues': total_issues,
            'completed_issues': completed_issues,
            'completion_rate': round((completed_issues / total_issues * 100), 1) if total_issues > 0 else 0,
            'state_distribution': state_counts,
            'average_durations_hours': avg_durations,
            'blocked_issues': len(self.get_blocked_issues()),
            'generated_at': datetime.now(timezone.utc).isoformat()
        }


def main():
    parser = argparse.ArgumentParser(description='Track workflow state for GitHub issues')
    parser.add_argument('--issue', type=int, help='Issue number')
    parser.add_argument('--state', help='Set workflow state')
    parser.add_argument('--agent', help='Agent responsible for state change')
    parser.add_argument('--get-state', action='store_true', help='Get current state')
    parser.add_argument('--progress', action='store_true', help='Get workflow progress')
    parser.add_argument('--report', action='store_true', help='Generate status report')
    parser.add_argument('--blocked', action='store_true', help='List blocked issues')
    parser.add_argument('--state-file', default='.workflow-state.json',
                       help='State file path')
    parser.add_argument('--format', choices=['text', 'json'], default='text',
                       help='Output format')

    args = parser.parse_args()

    tracker = WorkflowTracker(args.state_file)

    if args.report:
        report = tracker.generate_report()
        if args.format == 'json':
            print(json.dumps(report, indent=2))
        else:
            print(f"Workflow Status Report")
            print(f"Total Issues: {report['total_issues']}")
            print(f"Completed: {report['completed_issues']} ({report['completion_rate']}%)")
            print(f"Blocked: {report['blocked_issues']}")
            print("\nState Distribution:")
            for state, count in report['state_distribution'].items():
                print(f"  {state}: {count}")

    elif args.blocked:
        blocked = tracker.get_blocked_issues()
        if args.format == 'json':
            print(json.dumps(blocked, indent=2))
        else:
            if blocked:
                print("Blocked Issues:")
                for issue in blocked:
                    print(f"  Issue #{issue['issue_number']}: {issue['current_state']} "
                          f"({issue['hours_stalled']}h stalled)")
            else:
                print("No blocked issues found")

    elif args.issue:
        if args.state:
            # Set state
            success = tracker.set_issue_state(args.issue, args.state, args.agent)
            if success:
                print(f"Issue #{args.issue} state set to: {args.state}")
            else:
                print(f"Invalid state transition for issue #{args.issue}", file=sys.stderr)
                return 1

        elif args.get_state:
            # Get current state
            state = tracker.get_issue_state(args.issue)
            if args.format == 'json':
                print(json.dumps(state, indent=2))
            else:
                if state:
                    print(f"Issue #{args.issue}: {state['current_state']}")
                    if state.get('current_agent'):
                        print(f"Agent: {state['current_agent']}")
                else:
                    print(f"Issue #{args.issue} not found")

        elif args.progress:
            # Get progress
            progress = tracker.get_workflow_progress(args.issue)
            if args.format == 'json':
                print(json.dumps(progress, indent=2))
            else:
                print(f"Issue #{args.issue} Progress:")
                print(f"  Step: {progress['current_step']}/{progress['total_steps']}")
                print(f"  Progress: {progress['progress']}%")
                print(f"  Current State: {progress['current_state']}")
                if progress['next_states']:
                    print(f"  Next States: {', '.join(progress['next_states'])}")

    else:
        print("No action specified. Use --help for usage.", file=sys.stderr)
        return 1

    return 0


if __name__ == '__main__':
    sys.exit(main())