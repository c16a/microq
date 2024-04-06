package broker

import "strings"

const (
	PatternSeparator   = "."
	SingleTokenMatcher = "*"
	MultiTokenMatcher  = ">"
)

type Subscription struct {
	active  bool
	group   string
	pattern string
}

func (s *Subscription) IsActive() bool {
	return s.active
}

func (s *Subscription) GetGroup() string {
	return s.group
}

func (s *Subscription) GetPattern() string {
	return s.pattern
}

func (s *Subscription) Matches(topic string) bool {

	patternTokens := strings.Split(s.pattern, PatternSeparator)
	topicTokens := strings.Split(topic, PatternSeparator)

	if len(patternTokens) > len(topicTokens) {
		return false
	}

	for tIdx, topicToken := range topicTokens {
		if tIdx > len(patternTokens)-1 {
			return false
		}
		if patternTokens[tIdx] == MultiTokenMatcher {
			break
		}
		if patternTokens[tIdx] == topicToken || patternTokens[tIdx] == SingleTokenMatcher {
			continue
		} else {
			return false
		}
	}

	return true
}
