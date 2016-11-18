package locater

import "testing"

func TestLocateFile(t *testing.T) {
	expect := `{"created_at": "Thu May 31 00:00:01 +0000 2012"} `
	LocateFile("test/", "test.json", []byte(expect))
}
