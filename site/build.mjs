#!/usr/bin/env node
// Builds site/dist/index.html from README.md
// Content and design are unified: everything comes from README or filesystem

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
<script type="application/ld+json">\${jsonLd}</script>
<style>
:root {
  --bg: #f8fafc; --panel: #fff; --line: #e2e8f0; --ink: #0f172a; --muted: #475569;
  --accent: #1e40af; --accent-soft: #eff6ff; --code-bg: #f1f5f9;
  --mono: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
  --sans: -apple-system, BlinkMacSystemFont, "Segoe UI", Inter, Roboto, Helvetica, Arial, sans-serif;
}
@media (prefers-color-scheme: dark) {
  :root {
    --bg: #0f172a; --panel: #1e293b; --line: #334155; --ink: #f1f5f9; --muted: #94a3b8;
    --accent: #3b82f6; --accent-soft: #1e3a8a; --code-bg: #1e293b;
  }
}
* { box-sizing: border-box; }
html { scroll-behavior: smooth; scroll-padding-top: 5rem; }
body {
  margin: 0; background: var(--bg); color: var(--ink);
  font: 400 16px/1.6 var(--sans); -webkit-font-smoothing: antialiased;
}
.wrap { max-width: 56rem; margin: 0 auto; padding: 0 1.5rem; }
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
.brand { display: inline-flex; align-items: center; font-weight: 600; letter-spacing: .05em; color: var(--ink); text-decoration: none; }
.brand span { color: var(--accent); }
header.top nav { margin-left: auto; display: flex; gap: 1rem; flex-wrap: wrap; }
header.top nav a { color: var(--muted); text-decoration: none; font-size: .9rem; }
header.top nav a:hover { color: var(--ink); }

.hero { padding: 4rem 0 2rem; }
.eyebrow { display: inline-block; font: 600 .7rem/1 var(--mono); letter-spacing: .12em; text-transform: uppercase; color: var(--accent); background: var(--accent-soft); border-radius: 99px; padding: .4rem .8rem; margin-bottom: 1rem; }
.hero h1 { font-size: clamp(2.2rem, 6vw, 3.5rem); line-height: 1.1; margin: 0 0 1rem; letter-spacing: -.02em; }
.lede { font-size: clamp(1rem, 2vw, 1.15rem); color: var(--muted); max-width: 45rem; margin: 0 0 1.5rem; }
.lede strong { color: var(--ink); font-weight: 600; }
.cta { display: flex; gap: .6rem; flex-wrap: wrap; margin-bottom: 2.5rem; }
.cta a { display: inline-flex; align-items: center; gap: .4rem; text-decoration: none; font-size: .9rem; font-weight: 550; padding: .55rem 1rem; border-radius: 6px; border: 1px solid var(--line); color: var(--ink); background: var(--panel); }
.cta a.primary { background: var(--accent); border-color: var(--accent); color: #fff; }
.cta a:hover { border-color: var(--accent); }

section { padding: 3rem 0; border-top: 1px solid var(--line); }
section h2 { font-size: 1.4rem; letter-spacing: -.015em; margin: 0 0 1rem; }
section h2 .anchor { color: inherit; text-decoration: none; }
section h2 .anchor:hover::after { content: " #"; color: var(--accent); }
section h3 { font-size: 1rem; margin: 1.75rem 0 .5rem; }
section p, section li { max-width: 45rem; }
table { border-collapse: collapse; width: 100%; margin: 1rem 0; font-size: .9rem; display: block; overflow-x: auto; }
th, td { text-align: left; padding: .5rem .75rem; border-bottom: 1px solid var(--line); }
th { font-size: .75rem; text-transform: uppercase; letter-spacing: .08em; color: var(--muted); }
blockquote { margin: 1rem 0; padding: .1rem 0 .1rem 1rem; border-left: 3px solid var(--accent); color: var(--muted); }

footer { padding: 3rem 0; border-top: 1px solid var(--line); color: var(--muted); font-size: .9rem; }
footer p { margin: 0; }
</style>
</head>
<body>

<header class="top">
  <div class="wrap">
    <a href="#" class="brand">T<span>raefik</span> Go SDK</a>
    <nav>
${nav.map((s) => `      <a href="#${slug(s.heading)}">${esc(s.heading)}</a>`).join('\n')}
      <a href="${REPO}" style="color: var(--muted);">GitHub</a>
    </nav>
  </div>
</header>

<main>
  <div class="wrap">
    <div class="hero">
      <div class="eyebrow">Traefik SDK</div>
      <h1>${esc(title)}</h1>
      <p class="lede">${marked.parseInline(lede)}</p>
      <div class="cta">
        <a href="${REPO}/blob/main/README.md" class="primary">View on GitHub</a>
        <a href="${BLOB}/docs/guides/01-installation.md">Get Started</a>
        <a href="${BLOB}/docs/index.md">Documentation</a>
      </div>
      <div>${marked.parse(after)}</div>
    </div>
${rendered.join('\n')}
  </div>
</main>

<footer class="wrap">
  <p><strong>Traefik Go SDK</strong> — v${pkg.version} — <a href="${REPO}/blob/main/LICENSE">MIT License</a></p>
  <p>Built with <a href="${BLOB}/site/build.mjs">Node.js</a>. Content from <a href="${BLOB}/README.md">README.md</a>.</p>
</footer>

</body>
</html>
`

// Create output directory and write files
mkdirSync(OUT, { recursive: true })
writeFileSync(join(OUT, 'index.html'), html)
writeFileSync(join(OUT, '.nojekyll'), '')
console.log(`✓ Built ${join(OUT, 'index.html')} (${(html.length / 1024).toFixed(1)}KB)`)
