# Branch Protection Rules — traefik-go-sdk

**Main branch** (`main`) è protetto per garantire stabilità e qualità del codice.

## Rules (to apply via GitHub UI or `gh` CLI)

### 1. Require Pull Request Review Before Merge
- ✅ Require approvals: **1** (minimum)
- ✅ Dismiss stale pull request approvals when new commits are pushed
- ✅ Require review from code owners (if CODEOWNERS file exists)
- ✅ Allow specified actors to bypass required pull requests: *none*

### 2. Require Status Checks to Pass
- ✅ Require branches to be up to date before merging
- ✅ Require the following status checks to pass before merging:
  - `Go build` (from `.github/workflows/go-build.yml`)
  - `Go test workflow` (from `.github/workflows/go-test.yml`)
    - Must pass on **all** matrix versions (Go 1.18, 1.19, 1.20, 1.21)

### 3. Restrictions
- ✅ Restrict who can push to matching branches: *none* (allow all contributors)
- ✅ Allow force pushes: **NO**
- ✅ Allow deletions: **NO**

### 4. Merge Strategy
- ✅ Allow squash merging: **YES** (recommended for clean history)
- ✅ Allow rebase merging: **YES** (optional)
- ✅ Allow merge commits: **NO** (keep linear history)
- ✅ Require conversation resolution: **YES** (all discussions must be resolved)

---

## How to Apply

### Via GitHub Web UI (recommended for first-time setup)

1. Go to **Settings** → **Branches** → **Add rule**
2. Branch name pattern: `main`
3. Configure:
   - [x] Require a pull request before merging
   - [x] Require approvals (1)
   - [x] Dismiss stale pull request approvals
   - [x] Require status checks to pass before merging
   - [x] Require branches to be up to date before merging
   - [x] Include administrators in restrictions
4. Save

### Via GitHub CLI

```bash
# Setup branch protection for main
gh repo edit \
  --branch-protection main \
  --required-approvals 1 \
  --require-status-checks \
  --require-branches-up-to-date \
  --dismiss-stale-reviews \
  --require-code-owner-reviews  # if CODEOWNERS exists
```

---

## Workflow Impact

### For Contributors

**To merge a PR into `main`:**

1. ✅ Create a feature branch: `git checkout -b feat/your-feature`
2. ✅ Make commits (TDD: test first)
3. ✅ Push: `git push origin feat/your-feature`
4. ✅ Create PR on GitHub
5. ✅ Wait for CI to pass:
   - Go build (1.18-1.21)
   - Go test (1.18-1.21)
   - Coverage check (≥70%)
6. ✅ Request review (at least 1 approval)
7. ✅ Resolve any discussions
8. ✅ Merge (only after CI green + 1 approval)

**Merge strategy**: Use **Squash and merge** to keep main history clean:
```
Merge PR #123: feat: add GetApiVersion() method
```

### For Maintainers

- Branch protection prevents direct pushes to `main`
- All changes require PR + CI + review
- Stale approvals are dismissed when new commits arrive
- Merging requires latest branch (auto-resolved if safe)

---

## Exceptions

If urgent hotfix needed (security, production incident):

1. Create hotfix branch: `git checkout -b hotfix/security-issue`
2. Follow same PR process (no skipping CI or reviews)
3. If truly critical and urgent, contact maintainer for emergency merge override
4. Document in incident ticket

**Never** force-push or use `--no-verify` on main.

---

## Verification

To verify rules are in place:

```bash
# List branch protection rules
gh repo view --json branchProtectionRules

# Get details for 'main' branch
gh api repos/{owner}/{repo}/branches/main/protection
```

---

## Related

- **CLAUDE.md** — Commit and release strategy
- **docs/DEVELOPMENT.md** — TDD workflow
- **.github/workflows/** — CI configuration
