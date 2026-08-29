# B2B website architecture review checklist

Use this checklist with the
[website review record](../templates/website-review-record.md). Mark each item
Observed, Reported, Unknown, or Not applicable and attach evidence where useful.

## Decision job and positioning

- [ ] The audiences and decision-makers are defined.
- [ ] Important buyer questions and objections are recorded.
- [ ] The company’s category, offer, and differentiation are expressed consistently.
- [ ] Claims have evidence near the point of doubt.
- [ ] Primary and secondary actions match the page’s decision job.
- [ ] The website does not rely on one person to explain what the company does.

## Information architecture

- [ ] An inventory exists for important pages and content types.
- [ ] Navigation labels use language the audience can understand.
- [ ] Parent-child and sibling relationships are intentional.
- [ ] Core buyer routes connect problem, offer, proof, and action.
- [ ] Internal links reinforce meaningful relationships rather than only menu hierarchy.
- [ ] Search, filters, or taxonomies exist when content volume requires them.
- [ ] Orphaned, duplicate, and obsolete destinations are identified.
- [ ] Canonical URLs and redirect ownership are documented.

## Page system

- [ ] Repeatable page types use templates rather than one-off copies.
- [ ] Reusable components have defined content and interaction states.
- [ ] Responsive behavior is specified for each multi-column or interactive pattern.
- [ ] Heading hierarchy and body measure remain usable across viewports.
- [ ] Empty, loading, error, success, and validation states are accounted for.
- [ ] Components are reused without forcing unrelated content into one shape.
- [ ] Design tokens or equivalent shared values have one owner.

## Content model and CMS

- [ ] Structured fields represent repeatable information.
- [ ] Rich text is limited to content that genuinely needs free composition.
- [ ] Taxonomies have a clear editorial and user-facing purpose.
- [ ] Media has ownership, alternatives, and reuse rules.
- [ ] Editors can preview important states before publishing.
- [ ] Roles and permissions match actual responsibilities.
- [ ] Draft, review, publish, update, archive, and redirect workflows are defined.
- [ ] The team can create a common new page without rebuilding the architecture.

## Search and technical foundations

- [ ] Important pages return the intended HTTP response.
- [ ] Canonicals, redirects, robots rules, and the XML sitemap agree.
- [ ] Titles, descriptions, headings, and internal links describe each page’s job.
- [ ] Structured data matches visible, verified information.
- [ ] Rendering exposes important content and links to crawlers.
- [ ] Performance budgets or constraints are documented.
- [ ] Consent and analytics ownership are clear.
- [ ] Meaningful actions have defined measurement events.
- [ ] Environments, domains, DNS, source control, and deployment ownership are known.

## Accessibility and interaction

- [ ] The page declares a responsive viewport.
- [ ] Keyboard users can reach and operate interactive controls.
- [ ] Focus states are visible.
- [ ] Forms expose labels, instructions, errors, and confirmation states.
- [ ] Images and media have appropriate text alternatives or captions.
- [ ] Color is not the only carrier of meaning.
- [ ] Landmarks and headings create a usable document outline.
- [ ] Content does not cause document-level horizontal scrolling.
- [ ] Touch targets and text remain usable on a phone.

## Engineering and QA

- [ ] Source code has a canonical repository and named owner.
- [ ] Development, staging, and production boundaries are documented.
- [ ] Secrets are stored outside source code and shared documents.
- [ ] Changes pass review before production.
- [ ] Automated checks cover suitable code and behavior.
- [ ] Responsive visual QA covers agreed viewports.
- [ ] Production smoke checks exist for critical routes and actions.
- [ ] Known limitations and technical decisions are recorded.

## Handover and ongoing ownership

- [ ] The team knows where content, design, code, analytics, and credentials live.
- [ ] Operating documentation is organized by task or module.
- [ ] A new team member has a clear starting point.
- [ ] Backup, security, dependency, and maintenance ownership is assigned.
- [ ] A change log or equivalent decision history exists.
- [ ] The team can distinguish a content edit, design-system change, code change, and architecture change.
- [ ] Recommendations name an owner, evidence, next action, and recheck.

## Priority summary

After the review, record:

1. The structural decision most likely to block buyer understanding.
2. The architecture issue most likely to create duplicate work.
3. The operating gap most likely to make the website hard to own.
4. The first source artifact or behavior to change.
5. The evidence required for the recheck.
