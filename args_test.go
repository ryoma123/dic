package dic

import "testing"

func TestApplyFallbackArgs(t *testing.T) {
	args, opts, configPath, showHelp, showVersion := applyFallbackArgs([]string{"-r", "8.8.8.8"}, Options{}, "")
	if !opts.Reverse {
		t.Fatalf("Expected reverse to be true")
	}
	if len(args) != 1 || args[0] != "8.8.8.8" {
		t.Fatalf("Expected args to be stripped, got %#v", args)
	}
	if configPath != "" || showHelp || showVersion {
		t.Fatalf("Expected no extra flags, got config=%q help=%v version=%v", configPath, showHelp, showVersion)
	}

	args, opts, configPath, showHelp, showVersion = applyFallbackArgs([]string{"--follow-cname", "--cname-max", "7", "example.com"}, Options{}, "")
	if !opts.FollowCNAME || opts.CnameMax != 7 {
		t.Fatalf("Expected follow-cname true and cname-max 7, got %#v", opts)
	}
	if len(args) != 1 || args[0] != "example.com" {
		t.Fatalf("Expected args to be stripped, got %#v", args)
	}
	if configPath != "" || showHelp || showVersion {
		t.Fatalf("Expected no extra flags, got config=%q help=%v version=%v", configPath, showHelp, showVersion)
	}

	args, opts, configPath, showHelp, showVersion = applyFallbackArgs([]string{"--cname-max=3", "example.com"}, Options{}, "")
	if opts.CnameMax != 3 {
		t.Fatalf("Expected cname-max 3, got %#v", opts)
	}
	if len(args) != 1 || args[0] != "example.com" {
		t.Fatalf("Expected args to be stripped, got %#v", args)
	}
	if configPath != "" || showHelp || showVersion {
		t.Fatalf("Expected no extra flags, got config=%q help=%v version=%v", configPath, showHelp, showVersion)
	}

	args, opts, configPath, showHelp, showVersion = applyFallbackArgs([]string{"-c", "./config.toml", "example.com"}, Options{}, "")
	if configPath != "./config.toml" {
		t.Fatalf("Expected config path to be set, got %q", configPath)
	}
	if len(args) != 1 || args[0] != "example.com" {
		t.Fatalf("Expected args to be stripped, got %#v", args)
	}
	if showHelp || showVersion {
		t.Fatalf("Expected no help/version flags, got help=%v version=%v", showHelp, showVersion)
	}

	args, opts, configPath, showHelp, showVersion = applyFallbackArgs([]string{"--help", "example.com"}, Options{}, "")
	if !showHelp || showVersion {
		t.Fatalf("Expected help flag only, got help=%v version=%v", showHelp, showVersion)
	}
}
