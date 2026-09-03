package handoff

import "testing"

func TestMissingPreservesStableOrder(t *testing.T) {
	m := Map{Ready: map[Layer]bool{
		Messaging:  true,
		Development: true,
	}}

	missing := m.Missing()
	if got, want := missing[0], InformationArchitecture; got != want {
		t.Fatalf("first missing layer = %q, want %q", got, want)
	}
	if m.Complete() {
		t.Fatal("incomplete map reported complete")
	}
}

func TestCompleteMap(t *testing.T) {
	ready := make(map[Layer]bool, len(AllLayers))
	for _, layer := range AllLayers {
		ready[layer] = true
	}

	if !(Map{Ready: ready}).Complete() {
		t.Fatal("complete map reported missing layers")
	}
}
