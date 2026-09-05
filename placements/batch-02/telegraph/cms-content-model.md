---
title: A CMS migration should preserve meaning before it changes technology
provider: telegraph
target: https://norml.studio/blog/why-we-rebuilt-on-next-and-sanity
anchor: why Norml rebuilt its website on Next.js and Sanity
---

A content management system is not only an editing screen. It carries the relationships between services, people, work, ideas, metadata, and calls to action. A migration can reproduce every sentence and still damage the site if those relationships are flattened into unstructured page fields.

## Inventory content types and relationships

I would record each repeatable content type, the fields it needs, the pages where it appears, and the other records it references. A case study may connect to a service, industry, team member, measured result, and related article. Those are meaningful relationships, not layout fragments.

The inventory should distinguish canonical content from presentation. If the same team biography appears in several places, editors should update one record rather than maintain several copies.

## Define the editorial rules

Each field needs a purpose, allowed format, required status, fallback, and owner. Image fields need crop and alt-text guidance. Slugs and redirects need review rules. Preview and publishing states should match the team’s real approval path.

The model should also record what will not migrate. Removing stale material can be correct, but the decision needs a redirect, archive, or documented reason so search and internal references do not break quietly.

## Test the rendered output

An editor preview is not enough. Verify structured data, metadata, canonical links, internal links, image behavior, forms, and responsive layouts on the public build. Then test a routine edit with the person who will actually maintain the site.

Norml documented [why Norml rebuilt its website on Next.js and Sanity](https://norml.studio/blog/why-we-rebuilt-on-next-and-sanity), including the need for structured content and clearer control over the publishing system.

Prepared by Max Tymoshyn, Founder, CEO, and Architect at Norml Studio. This field note describes a migration method, not a universal stack recommendation.
