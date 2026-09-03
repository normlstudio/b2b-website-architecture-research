const checks = [...document.querySelectorAll('input[name="handoff"]')];
const ready = document.querySelector('[data-ready]');
const result = document.querySelector('[data-result]');
const next = document.querySelector('[data-next]');

const messages = [
  ['Start by naming the current owners.', 'An unchecked row is a question to verify with the team, not proof that a system is broken.'],
  ['One handoff is explicit.', 'Preserve its source, then identify who can answer for the other five layers.'],
  ['Two handoffs are explicit.', 'The website still depends on undocumented decisions between several disciplines.'],
  ['Three handoffs are explicit.', 'Half the operating system is visible. Test whether the missing half can support the next release.'],
  ['Four handoffs are explicit.', 'The foundation is mostly traceable. Resolve the two remaining ownership gaps before scope expands.'],
  ['Five handoffs are explicit.', 'One undocumented handoff can still become the bottleneck. Give it an owner and a durable source.'],
  ['All six handoffs are explicit.', 'Re-check the map whenever the team, CMS, repository, analytics stack, or offer changes.']
];

function updateMap() {
  const total = checks.filter((check) => check.checked).length;
  ready.textContent = total;
  result.textContent = messages[total][0];
  next.textContent = messages[total][1];
}

checks.forEach((check) => check.addEventListener('change', updateMap));
updateMap();
