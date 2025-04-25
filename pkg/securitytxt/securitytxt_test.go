package securitytxt

import (
	"testing"
)

func TestPreferredLanguagesParse(t *testing.T) {
	txt := "Contact: mailto:a@example.com\r\n" +
		"Expires: 2030-01-01T00:00:00Z\r\n" +
		"Preferred-Languages: en,fr\r\n"
	rec, err := New([]byte(txt))
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if len(rec.PreferredLanguages) != 2 || rec.PreferredLanguages[0] != "en" || rec.PreferredLanguages[1] != "fr" {
		t.Fatalf("unexpected languages: %+v", rec.PreferredLanguages)
	}
}
