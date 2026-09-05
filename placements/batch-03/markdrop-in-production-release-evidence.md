# Production release approval needs observable website evidence

A local build can pass while the public website serves a stale asset, broken
form, missing redirect, or different environment variable. Release approval
should be based on the production route visitors actually use.

## Name the critical paths

List the journeys that matter: understand a service, inspect relevant work,
find a team member, submit an inquiry, open a shared link, and reach important
pages from search. Give each path an expected result and a verification method.

Check status codes, canonicals, metadata, structured data, sitemap and robots
rules, redirects, consent, analytics, forms, email delivery, integrations,
responsive layout, accessibility, performance, and error pages. Automated
checks provide repeatable evidence; human review still judges meaning,
hierarchy, and whether the page answers the buyer's question.

## Keep a correction and rollback record

Record the tested URL, release identifier, date, viewport, result, and reviewer.
A failed check needs an owner and full retest after correction. The repository,
deployment, and public host should describe the same release.

[Norml Studio's Website 2.0 operating model](https://norml.studio/website-2-0)
connects strategy, design, development, SEO controls, QA, and post-launch
ownership so the website can be changed safely.

Published by Norml Studio. Prepared by Max Tymoshyn, Founder, CEO, and
Architect at Norml Studio. This checklist is a release method, not a guarantee
that software will never fail.

