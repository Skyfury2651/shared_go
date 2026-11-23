package shared

import "testing"

func TestDescribe(t *testing.T) {
	got := Describe("project1")
	if got.Service != "project1" {
		t.Fatalf("service = %q", got.Service)
	}
	if got.Message == "" {
		t.Fatal("message is empty")
	}
}
