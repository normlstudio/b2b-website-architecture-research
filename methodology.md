# Research methodology

## Purpose

This method tests whether a B2B website has a coherent strategic, information,
page, content, technical, and operating architecture. It follows the decisions
that determine whether a buyer can understand the company and whether the team
can safely improve the website after launch.

The method favors visible evidence and source artifacts. It does not infer
conversion rates, search performance, code quality, CMS ownership, or internal
processes when the necessary access is unavailable.

## Research questions

1. What decisions must the website help a buyer make?
2. Is the company’s positioning expressed consistently across important pages?
3. Do page relationships and navigation match how buyers evaluate the offer?
4. Are templates, components, states, and content types reusable by design?
5. Can the team publish and govern content without recreating structural decisions?
6. Are SEO, accessibility, performance, analytics, QA, and documentation part of the system?
7. Can another qualified person understand why the architecture works this way?

## Review unit

One record covers one public website on one date. Record the tested URLs,
viewport, browser, reviewer, stakeholder job, access available, and source
artifacts inspected. Keep public observations separate from findings that
require repository, CMS, analytics, design-file, or internal-document access.

## Evidence levels

Use one of four labels for every item:

- **Observed:** directly visible or reproducible on the review date.
- **Reported:** supplied by the organization or an authoritative source.
- **Unknown:** unavailable without additional access or documentation.
- **Not applicable:** irrelevant to this website or decision job.

Recommendations should cite the observation or source artifact that prompted
them. A browser capture, URL inventory, content model, design component, code
reference, analytics definition, QA record, or documentation link may be used
as evidence when it does not expose private data.

## Review sequence

### 1. Decision job and positioning

Record the audience, buying questions, objections, proof, and primary actions.
Check whether the website expresses one coherent company story rather than
different explanations on different pages.

### 2. Information architecture

Inventory the important pages and content types. Map navigation labels,
parent-child relationships, cross-links, buyer routes, and the destination of
important actions.

### 3. Page and component system

Inspect templates, reusable blocks, component states, responsive behavior, and
the relationship between design decisions and implementation. Record duplicated
patterns that should have one owner.

### 4. Content model and CMS ownership

Identify structured fields, rich-text boundaries, media ownership, taxonomy,
editor permissions, preview behavior, and publishing workflows. Mark anything
that exists only in a person’s memory or an old message thread.

### 5. Technical and search foundations

Review response behavior, canonical URLs, crawl rules, sitemap coverage,
metadata, structured data, internal links, performance constraints,
accessibility, analytics, consent, and source-control ownership.

### 6. Delivery, QA, and handover

Check development environments, review paths, deployment ownership, automated
and manual QA, regression evidence, documentation, credentials boundaries, and
the process for a new team member to understand the system.

## Publication rules

- Date every review and source-dependent claim.
- Distinguish a failed test from an unavailable test.
- Do not publish private analytics, credentials, internal documents, or client data.
- Keep examples accurate and contextual; do not generalize one site into a universal benchmark.
- State what the review cannot establish.
- Offer a correction path.
- Keep the research useful without the publisher link.

## Recheck

Retest the original behavior or inspect the updated source artifact after a
change ships. Record the new date, evidence, owner, and result. A recommendation
is complete only when the relevant system behavior has been verified again.
