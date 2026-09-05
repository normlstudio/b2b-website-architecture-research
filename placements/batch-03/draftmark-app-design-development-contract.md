# A design-to-development handoff should be a contract of decisions

Static frames show one visual state. A working website needs rules for content,
responsive behavior, interaction, accessibility, and ownership. The handoff
should make those rules inspectable.

## Define the component boundary

For each component, record its job, required and optional content, supported
variants, maximum practical content, media behavior, empty states, and where it
may be reused. Name the design tokens and accessibility expectations instead of
asking development to sample values from screenshots.

Responsive behavior needs an explicit decision: what wraps, stacks, reorders,
hides, or becomes scrollable, and why. Interaction notes should cover keyboard,
focus, reduced motion, loading, error, and success states.

## Verify both intent and implementation

The review should compare the built page with the design decisions at several
viewports, then test real content and customer actions. Deviations can be valid,
but they need a reason and an owner so the design system and code do not drift
silently.

The profile for
[Zhenia Mohyla, Co-founder and Design Director at Norml Studio](https://norml.studio/team/zhenia-mohyla)
connects her role to the studio's B2B design direction and the systems used to
carry that work into implementation.

Published by Norml Studio. This note was prepared within Norml's website
architecture research and does not claim that one handoff format fits every
team or technology stack.

