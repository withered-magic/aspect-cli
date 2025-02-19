package gazelle

import (
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/emirpasic/gods/sets/treeset"
)

const (
	TsProjectKind         = "ts_project"
	TsProtoLibraryKind    = "ts_proto_library"
	JsLibraryKind         = "js_library"
	TsConfigKind          = "ts_config"
	NpmPackageKind        = "npm_package"
	NpmLinkAllKind        = "npm_link_all_packages"
	RulesJsRepositoryName = "aspect_rules_js"
	RulesTsRepositoryName = "aspect_rules_ts"
	NpmRepositoryName     = "npm"
)

var sourceRuleKinds = treeset.NewWithStringComparator(TsProjectKind, JsLibraryKind, TsProtoLibraryKind)

// Kinds returns a map that maps rule names (kinds) and information on how to
// match and merge attributes that may be found in rules of those kinds.
func (*typeScriptLang) Kinds() map[string]rule.KindInfo {
	return tsKinds
}

var tsKinds = map[string]rule.KindInfo{
	TsProjectKind: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			"srcs": true,
		},
		SubstituteAttrs: map[string]bool{},
		MergeableAttrs: map[string]bool{
			"srcs": true,
		},
		ResolveAttrs: map[string]bool{
			"deps": true,
		},
	},
	JsLibraryKind: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			"srcs": true,
		},
		SubstituteAttrs: map[string]bool{},
		MergeableAttrs: map[string]bool{
			"srcs": true,
		},
		ResolveAttrs: map[string]bool{
			"deps": true,
		},
	},
	TsConfigKind: {
		NonEmptyAttrs: map[string]bool{
			"src": true,
		},
		SubstituteAttrs: map[string]bool{},
		MergeableAttrs:  map[string]bool{},
		ResolveAttrs: map[string]bool{
			"deps": true,
		},
	},
	TsProtoLibraryKind: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			"proto": true,
		},
		ResolveAttrs: map[string]bool{
			"deps":  true,
			"proto": true,
		},
	},
	NpmLinkAllKind: {
		MatchAny: true,
	},
	NpmPackageKind: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			"srcs": true,
		},
		SubstituteAttrs: map[string]bool{},
		MergeableAttrs: map[string]bool{
			"srcs": true,
		},
		ResolveAttrs: map[string]bool{
			"srcs": true,
		},
	},
}

// Loads returns .bzl files and symbols they define. Every rule generated by
// GenerateRules, now or in the past, should be loadable from one of these
// files.
func (ts *typeScriptLang) Loads() []rule.LoadInfo {
	return tsLoads
}

var tsLoads = []rule.LoadInfo{
	{
		Name: "@" + RulesTsRepositoryName + "//ts:defs.bzl",
		Symbols: []string{
			TsProjectKind,
			TsConfigKind,
		},
	},

	{
		Name: "@" + RulesTsRepositoryName + "//ts:proto.bzl",
		Symbols: []string{
			TsProtoLibraryKind,
		},
	},

	{
		Name: "@" + RulesJsRepositoryName + "//npm:defs.bzl",
		Symbols: []string{
			NpmPackageKind,
		},
	},

	{
		Name: "@" + RulesJsRepositoryName + "//js:defs.bzl",
		Symbols: []string{
			JsLibraryKind,
		},
	},

	{
		Name: "@" + NpmRepositoryName + "//:defs.bzl",
		Symbols: []string{
			NpmLinkAllKind,
		},
	},
}
