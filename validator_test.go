package emailvalidator_test

import (
	"testing"

	"github.com/igorsobreira/emailvalidator"
)

func TestValidEmails(t *testing.T) {
	emails := []string{
		"a+b@plus-in-local.com",
		"a_b@underscore-in-local.com",
		"user@example.com",
		" user@example.com ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@letters-in-local.org",
		"01234567890@numbers-in-local.net",
		"a@single-character-in-local.org",
		"one-character-third-level@a.example.com",
		"single-character-in-sld@x.org",
		"local@dash-in-sld.com",
		"letters-in-sld@123.com",
		"one-letter-sld@x.org",
		"uncommon-tld@sld.museum",
		"uncommon-tld@sld.travel",
		"uncommon-tld@sld.mobi",
		"country-code-tld@sld.uk",
		"country-code-tld@sld.rw",
		"local@sld.newTLD",
		"local@sub.domains.com",
		"aaa@bbb.co.jp",
		"nigel.worthington@big.co.uk",
		"f@c.com",
		"areallylongnameaasdfasdfasdfasdf@asdfasdfasdfasdfasdf.ab.cd.ef.gh.co.ca",
		// emails below are invalid using the strict_mode of email_validator gem, we
		// don't implement that mode yet
		"hans,peter@example.com",
		"hans(peter@example.com",
		"hans)peter@example.com",
		"partially.\"quoted\"@sld.com",
		"&'*+-./=?^_{}~@other-valid-characters-in-local.net",
		"mixed-1234-in-{+^}-local@sld.net",
	}
	for _, email := range emails {
		if !emailvalidator.Valid(email) {
			t.Errorf("email is invalid, should not be: %#v", email)
		}
	}
}

func TestInvalidEmails(t *testing.T) {
	emails := []string{
		"",
		"f@s",
		"f@s.c",
		"@bar.com",
		"test@example.com@example.com",
		"test@",
		"@missing-local.org",
		"a b@space-in-local.com",
		"! #$%`|@invalid-characters-in-local.org",
		"<>@[]`|@even-more-invalid-characters-in-local.org",
		"missing-sld@.com",
		"invalid-characters-in-sld@! \"#$%(),/;<>_[]`|.org",
		"missing-dot-before-tld@com",
		"missing-tld@sld.",
		" ",
		"missing-at-sign.net",
		"unbracketed-IP@127.0.0.1",
		"invalid-ip@127.0.0.1.26",
		"another-invalid-ip@127.0.0.256",
		"IP-and-port@127.0.0.1:25",
		"the-local-part-is-invalid-if-it-is-longer-than-sixty-four-characters@sld.net",
		"user@example.com\n<script>alert('hello')</script>",
	}
	for _, email := range emails {
		if emailvalidator.Valid(email) {
			t.Errorf("email is valid, should not be: %#v", email)
		}
	}
}
