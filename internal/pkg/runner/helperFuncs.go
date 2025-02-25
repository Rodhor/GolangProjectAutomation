package runner

import (
	"projectAutomation/internal/config"
	"regexp"
	"runtime"
	"strings"
)

func GetShell() config.Shell {
	if runtime.GOOS == "windows" {
		return config.Shell{Shell: "cmd", ArgFlag: "/c"}
	}
	return config.Shell{Shell: "sh", ArgFlag: "-c"}
}

func ResolvePlaceholder(p config.PlaceholderArg) string {
	switch p {

	case config.ArgProjectName:
		return GetProjectName()

	case config.ArgProjectPath:
		return GetProjectPath()

	case config.ArgProjectLanguage:
		return GetProjectLanguage()

	default:
		return "{{" + string(p) + "}}"
	}
}

func AdjustDynamicCommands(cmd string) string {
	re := regexp.MustCompile(`\{\{(\w+)\}\}`)
	return re.ReplaceAllStringFunc(cmd, func(match string) string {
		key := strings.Trim(match, "{}")

		p := config.PlaceholderArg(key)

		return ResolvePlaceholder(p)
	})
}

func GetProjectName() string {
	return ""
}

func GetProjectPath() string {
	return ""
}

func GetProjectLanguage() string {
	return ""
}
