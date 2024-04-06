package broker

import "testing"

var matchTests = []struct {
	pattern  string
	topic    string
	expected bool
}{
	{
		pattern:  "time.us.*",
		topic:    "time.us.east",
		expected: true,
	},
	{
		pattern:  "time.us.*",
		topic:    "time.us.east.atlanta",
		expected: false,
	},
	{
		pattern:  "time.us.>",
		topic:    "time.us.east.atlanta",
		expected: true,
	},
	{
		pattern:  "time.*.east",
		topic:    "time.us.east",
		expected: true,
	},
	{
		pattern:  "time.*.east",
		topic:    "time.eu.east",
		expected: true,
	},
	{
		pattern:  ">",
		topic:    "time.eu.east",
		expected: true,
	},
}

func TestSubscription_Matches(t *testing.T) {
	for _, tt := range matchTests {
		t.Run(tt.pattern, func(_t *testing.T) {
			subscription := &Subscription{
				active:  true,
				group:   "",
				pattern: tt.pattern,
			}
			ok := subscription.Matches(tt.topic)
			if ok != tt.expected {
				_t.Fatal("doesn't match")
			}
		})
	}
}
