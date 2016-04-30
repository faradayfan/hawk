package hawk

import (
	"fmt"
	"testing"
	"time"
)

type stubbedClock struct{}

func (c *stubbedClock) Now(offset time.Duration) int64 {
	fmt.Println("call")
	return 1365711458
}

func TestBewitConfig_GetBewit(t *testing.T) {
	c := &Credential{
		ID:  "123456",
		Key: "2983d45yun89q",
		Alg: SHA256,
	}

	b := &BewitConfig{
		Credential: c,
		Ttl:        24 * time.Hour * 365 * 100,
		Ext:        "some-app-data",
	}

	actual := b.GetBewit("http://example.com/resource/4?a=1&b=2", &stubbedClock{})
	expect := "MTIzNDU2XDQ1MTkzMTE0NThcYkkwanFlS1prUHE0V1hRMmkxK0NrQ2lOanZEc3BSVkNGajlmbElqMXphWT1cc29tZS1hcHAtZGF0YQ=="

	if actual != expect {
		t.Errorf("invalid bewit: %s", actual)
	}
}
