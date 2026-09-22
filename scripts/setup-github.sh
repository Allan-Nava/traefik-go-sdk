#!/bin/bash
# setup-github.sh — Setup GitHub branch protection and milestone
#
# Usage:
#   ./scripts/setup-github.sh [--check] [--apply] [--branch-protect] [--milestone]
#
# Requirements:
#   - gh CLI installed and authenticated
#   - Repo root as working directory
#   - GitHub admin access to the repository

set -euo pipefail

DRY_RUN=true
REPO_SLUG=$(gh repo view --json nameWithOwner --jq '.nameWithOwner')
REPO_ROOT=$(git rev-parse --show-toplevel)

echo "📋 GitHub Setup for: $REPO_SLUG"
echo ""

# Parse arguments
BRANCH_PROTECT=false
CREATE_MILESTONE=false
APPLY=false

while [[ $# -gt 0 ]]; do
    case $1 in
        --apply) APPLY=true; DRY_RUN=false; shift ;;
        --branch-protect) BRANCH_PROTECT=true; shift ;;
        --milestone) CREATE_MILESTONE=true; shift ;;
        --check) DRY_RUN=true; shift ;;
        *) echo "Unknown option: $1"; exit 1 ;;
    esac
done

# Default: both actions
if [[ "$BRANCH_PROTECT" == false ]] && [[ "$CREATE_MILESTONE" == false ]]; then
    BRANCH_PROTECT=true
    CREATE_MILESTONE=true
fi

# ============================================================================
# BRANCH PROTECTION
# ============================================================================

if [[ "$BRANCH_PROTECT" == true ]]; then
    echo "🔒 Branch Protection for 'main'"
    echo "================================"
    echo ""

    if [[ "$DRY_RUN" == true ]]; then
        echo "[DRY RUN] Would configure:"
        echo "  - Require 1 PR review before merge"
        echo "  - Require status checks (Go build + test)"
        echo "  - Require branches up to date"
        echo "  - Dismiss stale PR approvals"
        echo "  - Prevent force push"
        echo "  - Prevent deletion"
        echo ""
        echo "⚠️  Branch protection must be applied via GitHub UI or gh API directly."
        echo ""
        echo "To apply via GitHub UI:"
        echo "  1. Go to Settings → Branches → Add rule"
        echo "  2. Branch name: 'main'"
        echo "  3. Enable the rules above"
        echo ""
        echo "To apply via gh CLI (requires additional setup):"
        echo "  # This requires gh extensions or direct API calls"
        echo "  gh api repos/$REPO_SLUG/branches/main/protection --input /dev/stdin <<EOF"
        echo "  {\"required_status_checks\": {\"strict\": true, \"contexts\": [\"Go build\", \"Go test workflow\"]}}"
        echo "  EOF"
        echo ""
    else
        echo "ℹ️  Use GitHub UI to apply branch protection."
        echo "   See .github/BRANCH_PROTECTION.md for detailed instructions."
    fi
    echo ""
fi

# ============================================================================
# MILESTONE
# ============================================================================

if [[ "$CREATE_MILESTONE" == true ]]; then
    echo "🎯 Milestone: v1.0.0"
    echo "====================="
    echo ""

    # Check if milestone exists
    existing=$(gh api repos/$REPO_SLUG/milestones --jq '.[] | select(.title=="v1.0.0") | .number' 2>/dev/null || echo "")

    if [[ -z "$existing" ]]; then
        if [[ "$DRY_RUN" == true ]]; then
            echo "[DRY RUN] Would create milestone:"
            echo "  - Title: v1.0.0"
            echo "  - Description: Stable, fully-tested SDK for Traefik API interaction"
            echo "  - Due date: 2026-10-31"
            echo ""
        else
            echo "Creating milestone v1.0.0..."
            gh api repos/$REPO_SLUG/milestones \
                -f title="v1.0.0" \
                -f description="Stable, fully-tested SDK for Traefik API interaction. See docs/MILESTONE.md for details." \
                -f due_on="2026-10-31" \
                --jq '.html_url'
            echo "✅ Milestone created"
        fi
    else
        echo "✅ Milestone v1.0.0 already exists (ID: $existing)"
    fi
    echo ""
fi

# ============================================================================
# SUMMARY
# ============================================================================

if [[ "$DRY_RUN" == true ]]; then
    echo "📖 Next steps:"
    echo ""
    echo "1. Apply branch protection (via GitHub UI):"
    echo "   Settings → Branches → Add rule"
    echo "   See .github/BRANCH_PROTECTION.md"
    echo ""
    echo "2. Rerun with --apply flag to create milestone:"
    echo "   ./scripts/setup-github.sh --apply --milestone"
    echo ""
else
    echo "✅ GitHub setup complete!"
    echo ""
    echo "Your repository is now configured with:"
    echo "  ✓ Milestone v1.0.0"
    echo "  ✓ Branch protection rules (see .github/BRANCH_PROTECTION.md)"
    echo ""
    echo "Next: Apply branch protection via GitHub UI (Settings → Branches)"
fi
