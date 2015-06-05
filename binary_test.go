package binary

import "testing"

func TestExtension(t *testing.T) {
	curLen := 184
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.css") {
		t.Fatal("Wrong detection. Must not be binary")
	}

	if !Is("a/b/c/bar.exe") {
		t.Fatal("Wrong detection. Must be binary")
	}

	if !Is("a/b/c/bar.a") {
		t.Fatal("Wrong detection. Must be binary")
	}

	if Is("a/b/c/7z") {
		t.Fatal("Wrong detection. Must not be binary")
	}
}
