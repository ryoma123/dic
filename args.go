package dic

import (
	"strconv"
	"strings"
)

func applyFallbackArgs(args []string, opts Options, configPath string) ([]string, Options, string, bool, bool) {
	filtered := make([]string, 0, len(args))
	showHelp := false
	showVersion := false

	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "-h" || a == "--help":
			showHelp = true
		case a == "-v" || a == "--version":
			showVersion = true
		case a == "-r" || a == "--reverse":
			opts.Reverse = true
		case strings.HasPrefix(a, "--reverse="):
			opts.Reverse = parseBoolArg(strings.TrimPrefix(a, "--reverse="), opts.Reverse)
		case a == "-f" || a == "--follow-cname":
			opts.FollowCNAME = true
		case strings.HasPrefix(a, "--follow-cname="):
			opts.FollowCNAME = parseBoolArg(strings.TrimPrefix(a, "--follow-cname="), opts.FollowCNAME)
		case a == "-c" || a == "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
				continue
			}
			filtered = append(filtered, a)
		case strings.HasPrefix(a, "--config="):
			configPath = strings.TrimPrefix(a, "--config=")
		case a == "-m" || a == "--cname-max":
			if i+1 < len(args) {
				if n, err := strconv.Atoi(args[i+1]); err == nil {
					opts.CnameMax = n
					i++
					continue
				}
			}
			filtered = append(filtered, a)
		case strings.HasPrefix(a, "--cname-max="):
			if n, err := strconv.Atoi(strings.TrimPrefix(a, "--cname-max=")); err == nil {
				opts.CnameMax = n
			} else {
				filtered = append(filtered, a)
			}
		default:
			filtered = append(filtered, a)
		}
	}

	return filtered, opts, configPath, showHelp, showVersion
}

func parseBoolArg(value string, fallback bool) bool {
	switch strings.ToLower(value) {
	case "1", "true", "t", "yes", "y":
		return true
	case "0", "false", "f", "no", "n":
		return false
	default:
		return fallback
	}
}
