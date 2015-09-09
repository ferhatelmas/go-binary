package binary

import (
	"fmt"
	"testing"

	"github.com/shamsher31/gobinext"
)

func TestExtension(t *testing.T) {
	curLen := 184
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	for _, ext := range binext.Get() {
		if !Is(fmt.Sprintf(`a/b/c/foo.bar.%s`, ext)) {
			t.Fatal("Wrong detection. Expected to be binary", ext)
		}
	}
}
