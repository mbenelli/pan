package jenkins

import (
	"encoding/json"
	"testing"
)

func TestJenkinsConf(t *testing.T) {
	sample := `{"Url": "foo", "Jobs": ["bar0", "bar1", "bar2"]}`
	var j Jenkins
	err := json.Unmarshal([]byte(sample), &j)
	if err != nil {
		t.Fatalf("Json error: %v", err)
	}
	if j.Url != "foo" {
		t.Fatalf(`Url = %q, want "foo"`, j.Url)
	}
}
