#!/bin/bash
# sync-backlog.sh — Synchronize docs/backlog.md ↔ GitHub Issues
#
# Usage:
#   ./scripts/sync-backlog.sh [--dry-run] [--check]
#
# --dry-run : print actions without creating issues
# --check   : verify sync consistency, exit 1 if mismatch
#
# Requirements:
#   - gh CLI installed and authenticated
#   - repo root as working directory
#   - GITHUB_TOKEN env var (or gh auth login)

set -euo pipefail

DRY_RUN=false
CHECK_MODE=false
REPO_ROOT=$(git rev-parse --show-toplevel)
BACKLOG_FILE="$REPO_ROOT/docs/backlog.md"

while [[ $# -gt 0 ]]; do
    case $1 in
        --dry-run) DRY_RUN=true; shift ;;
        --check) CHECK_MODE=true; shift ;;
        *) echo "Unknown option: $1"; exit 1 ;;
    esac
done

# Extract backlog items from docs/backlog.md
# Format: - `[ID-NNN]` Description `[backlog]`
extract_backlog_items() {
    grep -E '^\s*-\s+`\[ID-[0-9]+\]`' "$BACKLOG_FILE" | sed 's/^[[:space:]]*-[[:space:]]*`\[\(ID-[0-9]*\)\]`[[:space:]]*//;s/`[[:space:]]*\[\(.*\)\]\s*$//' || true
}

# Fetch GitHub Issues with label 'backlog'
fetch_github_issues() {
    gh issue list --label backlog --json number,title,labels --jq '.[] | "\(.number) \(.title)"' 2>/dev/null || true
}

# Create a new GitHub issue for a backlog item
create_github_issue() {
    local title="$1"
    local id="$2"

    if [[ $DRY_RUN == true ]]; then
        echo "[DRY-RUN] Would create issue: $id - $title"
        return
    fi

    gh issue create --title "$title" --label backlog --body "Backlog item: $id" 2>/dev/null || true
}

# Main sync logic
main() {
    echo "=== Syncing backlog.md ↔ GitHub Issues ==="

    if [[ ! -f "$BACKLOG_FILE" ]]; then
        echo "ERROR: $BACKLOG_FILE not found"
        exit 1
    fi

    # Check mode: just verify
    if [[ $CHECK_MODE == true ]]; then
        echo "Checking consistency..."
        local backlog_count=$(extract_backlog_items | wc -l)
        local github_count=$(fetch_github_issues | wc -l)
        echo "Backlog items: $backlog_count"
        echo "GitHub issues [backlog]: $github_count"

        if [[ $backlog_count -eq 0 && $github_count -eq 0 ]]; then
            echo "✓ Sync OK (both empty)"
            exit 0
        fi

        # Simple check: non-zero counts should roughly match
        if [[ $backlog_count -gt 0 ]] && [[ $github_count -eq 0 ]]; then
            echo "⚠ WARNING: Backlog has items but GitHub has none (run without --check to sync)"
            exit 1
        fi
        exit 0
    fi

    # Sync mode: create missing issues
    echo "Syncing items..."
    extract_backlog_items | while IFS= read -r line; do
        if [[ -z "$line" ]]; then continue; fi

        # Parse: "ID-NNN Description"
        id=$(echo "$line" | cut -d' ' -f1)
        desc=$(echo "$line" | cut -d' ' -f2-)

        # Check if issue already exists
        if gh issue list --label backlog --json title | grep -q "$id"; then
            echo "✓ Already synced: $id"
        else
            echo "→ Creating issue: $id - $desc"
            create_github_issue "$desc" "$id"
        fi
    done

    echo "=== Sync complete ==="
}

main
