package select_target

import (
	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// Option Опция.
type Option func([]*target.TargetInfo) (*target.TargetId, bool)

// ByEqualUrl По равному URL.
func ByEqualUrl(u string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, EqualUrlRule(u))
	}
}

// ByContainsUrl По части URL-адреса.
func ByContainsUrl(u string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, ContainsUrlRule(u))
	}
}

// ByPrefixUrl По префиксу URL.
func ByPrefixUrl(prefix string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, PrefixUrlRule(prefix))
	}
}

// BySuffixUrl По суффиксу URL.
func BySuffixT(suffix string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, SuffixUrlRule(suffix))
	}
}

// ByEqualTitle По равному Title.
func ByEqualTitle(title string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, EqualTitleRule(title))
	}
}

// ByContainsTitle По части Title.
func ByContainsTitle(title string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, ContainsTitleRule(title))
	}
}

// ByPrefixTitle По префиксу Title.
func ByPrefixTitle(prefix string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, PrefixTitleRule(prefix))
	}
}

// BySuffixTitle По префиксу Title.
func BySuffixTitle(suffix string) Option {
	return func(t []*target.TargetInfo) (*target.TargetId, bool) {
		return searchTarget(t, SuffixTitleRule(suffix))
	}
}

// searchTarget Поиск вкладки путем итерации.
func searchTarget(ts []*target.TargetInfo, r Rule) (*target.TargetId, bool) {
	for _, t := range ts {
		if r(t) {
			return t.TargetId, true
		}
	}
	return nil, false
}
