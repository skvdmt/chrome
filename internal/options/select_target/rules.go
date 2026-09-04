package select_target

import (
	"strings"

	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// Rule Сигнатура правила.
type Rule func(*target.TargetInfo) bool

// EqualUrlRule Правило в котором URL Эквивалентно значению аргумента u.
func EqualUrlRule(u string) Rule {
	return func(i *target.TargetInfo) bool {
		return i.Url == u
	}
}

// EqualUrlRule Правило в котором URL содержит значению аргумента u.
func ContainsUrlRule(u string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.Contains(i.Url, u)
	}
}

// PrefixUrlRule Правило в котором URL имеет префикс из аргумента pref.
func PrefixUrlRule(prefix string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.HasPrefix(i.Url, prefix)
	}
}

// SuffixUrlRule Правило в котором URL имеет суффикс из аргумента suff.
func SuffixUrlRule(suffix string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.HasSuffix(i.Url, suffix)
	}
}

// EqualTitleRule Правило в котором Title Эквивалентно значению аргумента title.
func EqualTitleRule(title string) Rule {
	return func(i *target.TargetInfo) bool {
		return i.Title == title
	}
}

// EqualTitleRule Правило в котором Title содержит значению аргумента title.
func ContainsTitleRule(title string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.Contains(i.Title, title)
	}
}

// PrefixTitleRule Правило в котором Title имеет префикс из аргумента prefix.
func PrefixTitleRule(prefix string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.HasPrefix(i.Title, prefix)
	}
}

// SuffixTitleRule Правило в котором Title имеет суффикс из аргумента suffix.
func SuffixTitleRule(suffix string) Rule {
	return func(i *target.TargetInfo) bool {
		return strings.HasSuffix(i.Title, suffix)
	}
}
