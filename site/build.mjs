#!/usr/bin/env node
// Builds site/dist/index.html from README.md with modern, attractive design
// Features: stats cards, feature cards, gradient text, smooth animations

import { marked } from 'marked'
import { cpSync, existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs'
import { dirname, join, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'

const ROOT = resolve(dirname(fileURLToPath(import.meta.url)), '..')
const OUT = resolve(ROOT, process.argv.includes('--out') ? process.argv[process.argv.indexOf('--out') + 1] : 'site/dist')
const REPO = 'https://github.com/Allan-Nava/traefik-go-sdk'
const BLOB = `${REPO}/blob/main`
const SITE = 'https://allan-nava.github.io/traefik-go-sdk/'

const pkg = JSON.parse(readFileSync(join(ROOT, 'package.json'), 'utf8'))
const md = readFileSync(join(ROOT, 'README.md'), 'utf8')

marked.setOptions({ mangle: false, headerIds: false })

const esc = (s) => s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;')
const slug = (s) =>
  s
    .toLowerCase()
    .replace(/[^\w\s-]/g, '')
    .trim()
    .replace(/\s+/g, '-')

// Parse README into intro + sections by H2
function parseReadme(source) {
  const lines = source.split('\n')
  const sections = []
  let title = 'Traefik Go SDK'
  let current = { heading: null, lines: [] }
  let fenced = false

  for (const line of lines) {
    if (line.startsWith('```')) fenced = !fenced
    if (!fenced && line.startsWith('# ')) {
      title = line.slice(2).trim()
      continue
    }
    if (!fenced && line.startsWith('## ')) {
      sections.push(current)
      current = { heading: line.slice(3).trim(), lines: [] }
      continue
    }
    current.lines.push(line)
  }
  sections.push(current)

  const intro = sections.shift()
  return { title, intro: intro.lines.join('\n').trim(), sections: sections.map((s) => ({ ...s, body: s.lines.join('\n').trim() })) }
}

// Extract lede (first paragraph) from intro
function parseIntro(raw) {
  const intro = raw.split('\n').filter((l) => !/^\s*<\/?(p|img|div|a|picture|source)\b/i.test(l)).join('\n').trim()
  const [lede, ...rest] = intro.split(/\n\n+/)
  return { lede: lede.trim(), after: rest.join('\n\n').trim() }
}

// Fix empty table headers
function dropEmptyHead(html) {
  return html.replace(/<thead>[\s\S]*?<\/thead>/g, (thead) => (/>[^<\s][\s\S]*?<\/th>/.test(thead) ? thead : ''))
}

// Link paths to repo
function linkifyPaths(html) {
  return html.replace(/<code>([\w./-]+\.(?:md|go|yml|json|mjs)|(?:docs|scripts|traefik)\/[\w./-]*)<\/code>/g, (full, path) => {
    const clean = path.replace(/\/$/, '')
    if (!existsSync(join(ROOT, clean))) return full
    return `<a class="pathlink" href="${BLOB}/${clean}"><code>${path}</code></a>`
  })
}

const renderSection = (s) => {
  const id = slug(s.heading)
  return `  <section id="${id}">
    <h2><a class="anchor" href="#${id}">${esc(s.heading)}</a></h2>
${linkifyPaths(dropEmptyHead(marked.parse(s.body)))}
  </section>`
}

// Assemble page
const { title, intro, sections } = parseReadme(md)
const { lede, after } = parseIntro(intro)
const description = lede.replace(/\*\*/g, '').replace(/\n/g, ' ').split(/\.\s/)[0].concat('.')
const body = sections.filter((s) => !/^license$/i.test(s.heading))
const nav = body.filter((s) => !/^contributing$/i.test(s.heading))
const rendered = body.map(renderSection)

const headline = `${title} — Go SDK for Traefik API`

const jsonLd = JSON.stringify({
  '@context': 'https://schema.org',
  '@graph': [
    {
      '@type': 'WebSite',
      '@id': `${SITE}#website`,
      url: SITE,
      name: title,
      description,
      inLanguage: 'en',
    },
    {
      '@type': 'SoftwareApplication',
      '@id': `${SITE}#sdk`,
      name: title,
      description,
      url: SITE,
      applicationCategory: 'DeveloperApplication',
      operatingSystem: 'Linux, macOS, Windows',
      softwareVersion: pkg.version,
      codeRepository: REPO,
      license: 'https://opensource.org/licenses/MIT',
      author: { '@type': 'Person', name: 'Allan Nava' },
      offers: { '@type': 'Offer', price: '0', priceCurrency: 'USD' },
    },
  ],
}).replace(/</g, '\\u003c')

const html = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>${esc(headline)}</title>
<meta name="description" content="${esc(description)}">
<link rel="canonical" href="${SITE}">
<meta name="theme-color" content="#1e40af" media="(prefers-color-scheme: light)">
<meta name="theme-color" content="#0f172a" media="(prefers-color-scheme: dark)">
<meta property="og:type" content="website">
<meta property="og:site_name" content="${esc(title)}">
<meta property="og:locale" content="en">
<meta property="og:url" content="${SITE}">
<meta property="og:title" content="${esc(headline)}">
<meta property="og:description" content="${esc(description)}">
<meta property="og:image" content="${SITE}assets/og-image.png">
<meta property="og:image:width" content="1280">
<meta property="og:image:height" content="640">
<meta property="og:image:alt" content="${esc(headline)}">
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:title" content="${esc(headline)}">
<meta name="twitter:description" content="${esc(description)}">
<meta name="twitter:image" content="${SITE}assets/og-image.png">
<meta name="twitter:image:alt" content="${esc(headline)}">
<link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90' font-weight='bold' font-family='system-ui'>T</text></svg>">
<script type="application/ld+json">${jsonLd}</script>
<style>
:root {
  --bg: #f8fafc; --bg-alt: #f1f5f9; --panel: #fff; --line: #e2e8f0; --ink: #0f172a; --muted: #475569;
  --accent: #1e40af; --accent-soft: #eff6ff; --accent-light: #3b82f6; --code-bg: #f1f5f9;
  --success: #10b981; --warning: #f59e0b;
  --mono: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
  --sans: -apple-system, BlinkMacSystemFont, "Segoe UI", Inter, Roboto, Helvetica, Arial, sans-serif;
}
@media (prefers-color-scheme: dark) {
  :root {
    --bg: #0f172a; --bg-alt: #1a202c; --panel: #1e293b; --line: #334155; --ink: #f1f5f9; --muted: #94a3b8;
    --accent: #3b82f6; --accent-soft: #1e3a8a; --accent-light: #60a5fa; --code-bg: #1e293b;
    --success: #059669; --warning: #d97706;
  }
}
* { box-sizing: border-box; }
html { scroll-behavior: smooth; scroll-padding-top: 5rem; }
body {
  margin: 0; background: var(--bg); color: var(--ink);
  font: 400 16px/1.6 var(--sans); -webkit-font-smoothing: antialiased;
}
.wrap { max-width: 64rem; margin: 0 auto; padding: 0 1.5rem; }
a { color: var(--accent); text-decoration-thickness: 1px; text-underline-offset: 2px; }
code { font-family: var(--mono); font-size: .9em; }
:not(pre) > code { background: var(--code-bg); padding: .2em .4em; border-radius: 4px; }
pre { background: var(--code-bg); border: 1px solid var(--line); border-radius: 8px; padding: 1rem; overflow-x: auto; font-size: .85rem; line-height: 1.5; }
pre code { background: none; padding: 0; }

header.top {
  position: sticky; top: 0; z-index: 10; backdrop-filter: blur(10px);
  background: color-mix(in srgb, var(--bg) 86%, transparent);
  border-bottom: 1px solid var(--line);
}
header.top .wrap { display: flex; align-items: center; gap: 1.5rem; height: 3.5rem; }
.brand { display: inline-flex; align-items: center; font-weight: 700; letter-spacing: -.02em; color: var(--ink); text-decoration: none; font-size: 1.1rem; }
.brand span { color: var(--accent); font-weight: 700; }
header.top nav { margin-left: auto; display: flex; gap: 1.5rem; flex-wrap: wrap; align-items: center; }
header.top nav a { color: var(--muted); text-decoration: none; font-size: .9rem; font-weight: 500; }
header.top nav a:hover { color: var(--accent); }

.hero { padding: 6rem 0 3rem; }
.eyebrow { display: inline-block; font: 600 .65rem/1 var(--mono); letter-spacing: .12em; text-transform: uppercase; color: var(--accent); background: var(--accent-soft); border-radius: 99px; padding: .5rem .9rem; margin-bottom: 1.2rem; }
.hero h1 { font-size: clamp(2.4rem, 7vw, 3.8rem); line-height: 1.05; margin: 0 0 1rem; letter-spacing: -.03em; font-weight: 800; background: linear-gradient(135deg, var(--ink) 0%, var(--accent) 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }
.lede { font-size: clamp(1.05rem, 2.5vw, 1.25rem); color: var(--muted); max-width: 50rem; margin: 0 0 1.8rem; line-height: 1.5; }
.lede strong { color: var(--ink); font-weight: 600; }
.cta { display: flex; gap: .8rem; flex-wrap: wrap; margin-bottom: 3rem; }
.cta a { display: inline-flex; align-items: center; gap: .5rem; text-decoration: none; font-size: .95rem; font-weight: 600; padding: .65rem 1.2rem; border-radius: 7px; border: 1.5px solid var(--line); color: var(--ink); background: var(--panel); transition: all .2s; }
.cta a.primary { background: var(--accent); border-color: var(--accent); color: #fff; box-shadow: 0 4px 12px color-mix(in srgb, var(--accent) 25%, transparent); }
.cta a:hover { border-color: var(--accent); transform: translateY(-1px); }
.cta a.primary:hover { box-shadow: 0 6px 16px color-mix(in srgb, var(--accent) 35%, transparent); }

.stats { display: grid; grid-template-columns: repeat(auto-fit, minmax(140px, 1fr)); gap: 1.5rem; margin: 3rem 0; }
.stat { background: var(--panel); border: 1px solid var(--line); border-radius: 10px; padding: 1.5rem; text-align: center; transition: all .3s; }
.stat:hover { border-color: var(--accent); box-shadow: 0 4px 12px color-mix(in srgb, var(--accent) 10%, transparent); }
.stat-value { font-size: 2rem; font-weight: 800; color: var(--accent); }
.stat-label { font-size: .85rem; color: var(--muted); margin-top: .4rem; text-transform: uppercase; letter-spacing: .05em; }

.features { display: grid; grid-template-columns: repeat(auto-fit, minmax(220px, 1fr)); gap: 1.5rem; margin: 2rem 0 0; }
.feature-card { background: var(--panel); border: 1px solid var(--line); border-radius: 10px; padding: 2rem 1.5rem; transition: all .3s; }
.feature-card:hover { border-color: var(--accent); box-shadow: 0 8px 24px color-mix(in srgb, var(--accent) 15%, transparent); transform: translateY(-4px); }
.feature-card h3 { margin: 0 0 .8rem; color: var(--ink); font-size: 1.1rem; font-weight: 600; }
.feature-card p { margin: 0; color: var(--muted); font-size: .9rem; line-height: 1.5; }
.feature-icon { width: 2.5rem; height: 2.5rem; background: var(--accent-soft); border-radius: 8px; display: flex; align-items: center; justify-content: center; font-size: 1.3rem; margin-bottom: .8rem; }

section { padding: 3.5rem 0; border-top: 1px solid var(--line); }
section h2 { font-size: 1.8rem; letter-spacing: -.02em; margin: 0 0 1.5rem; font-weight: 700; color: var(--ink); }
section h2 .anchor { color: inherit; text-decoration: none; }
section h2 .anchor:hover::after { content: " #"; color: var(--accent); }
section h3 { font-size: 1.15rem; margin: 2rem 0 .8rem; font-weight: 600; }
section p, section li { max-width: 50rem; color: var(--muted); }
section li { margin: .5rem 0; }
table { border-collapse: collapse; width: 100%; margin: 1.5rem 0; font-size: .9rem; display: block; overflow-x: auto; }
th, td { text-align: left; padding: .75rem 1rem; border-bottom: 1px solid var(--line); }
th { font-size: .8rem; text-transform: uppercase; letter-spacing: .08em; color: var(--muted); font-weight: 600; background: var(--bg-alt); }
blockquote { margin: 1.5rem 0; padding: .1rem 0 .1rem 1.2rem; border-left: 4px solid var(--accent); color: var(--muted); font-style: italic; }

footer { padding: 3.5rem 0; border-top: 1px solid var(--line); color: var(--muted); font-size: .9rem; }
footer p { margin: 0.5rem 0; }
footer a { color: var(--muted); }
footer a:hover { color: var(--accent); }

@media (max-width: 640px) {
  .hero { padding: 4rem 0 2rem; }
  .stats { grid-template-columns: 1fr 1fr; }
  .features { grid-template-columns: 1fr; }
}
</style>
</head>
<body>

<header class="top">
  <div class="wrap">
    <a href="#" class="brand">T<span>raefik</span></a>
    <nav>
${nav.map((s) => `      <a href="#${slug(s.heading)}">${esc(s.heading)}</a>`).join('\n')}
      <a href="${REPO}">GitHub →</a>
    </nav>
  </div>
</header>

<main>
  <div class="wrap">
    <div class="hero">
      <div class="eyebrow">⚡ Go SDK v${pkg.version}</div>
      <h1>${esc(title)}</h1>
      <p class="lede">${marked.parseInline(lede)}</p>

      <div class="stats">
        <div class="stat">
          <div class="stat-value">27+</div>
          <div class="stat-label">API Methods</div>
        </div>
        <div class="stat">
          <div class="stat-value">86.9%</div>
          <div class="stat-label">Coverage</div>
        </div>
        <div class="stat">
          <div class="stat-value">50+</div>
          <div class="stat-label">Test Cases</div>
        </div>
        <div class="stat">
          <div class="stat-value">v1.0.0</div>
          <div class="stat-label">Production Ready</div>
        </div>
      </div>

      <div class="cta">
        <a href="${REPO}" class="primary">🚀 Get Started on GitHub</a>
        <a href="${BLOB}/README.md">📖 Full Documentation</a>
        <a href="${BLOB}/CHANGELOG.md">🔖 Release Notes</a>
      </div>

      <div style="margin: 3rem 0 0;">
        <h3 style="color: var(--ink); margin-top: 0; font-size: 1.3rem;">6 Major Features</h3>
        <div class="features">
          <div class="feature-card">
            <div class="feature-icon">📝</div>
            <h3>HTTP Write</h3>
            <p>Create, update, delete HTTP routers & services with full CRUD support</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">⚙️</div>
            <h3>Batch Ops</h3>
            <p>Apply, validate, and reset complete configurations atomically</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🔗</div>
            <h3>TCP/UDP Ops</h3>
            <p>Manage TCP and UDP routers, services, and configurations</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🛡️</div>
            <h3>Middleware</h3>
            <p>Create and manage HTTP & TCP middlewares (basicAuth, ipWhiteList, etc.)</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🔍</div>
            <h3>Filtering</h3>
            <p>Advanced queries: by rule, entrypoint, service, type, and more</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">💾</div>
            <h3>Export/Import</h3>
            <p>Backup and restore configurations for disaster recovery</p>
          </div>
        </div>
      </div>
    </div>
${rendered.join('\n')}
  </div>
</main>

<footer class="wrap">
  <p><strong>Traefik Go SDK</strong> — v${pkg.version} — <a href="${REPO}/blob/main/LICENSE">MIT License</a> — Built by Allan Nava</p>
  <p><a href="${REPO}">GitHub Repository</a> • <a href="${BLOB}/CHANGELOG.md">Changelog</a> • <a href="${BLOB}/CLAUDE.md">Contribution Guide</a></p>
</footer>

</body>
</html>
`

// Create output directory and write files
mkdirSync(OUT, { recursive: true })
writeFileSync(join(OUT, 'index.html'), html)
writeFileSync(join(OUT, '.nojekyll'), '')
console.log(`✓ Built ${join(OUT, 'index.html')} (${(html.length / 1024).toFixed(1)}KB)`)
