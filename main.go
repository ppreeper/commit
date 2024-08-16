package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
)

func chooseType() (choose string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("commit type").
				Options(
					huh.NewOption("Features: A new feature", "feat"),
					huh.NewOption("Bug Fixes: A bug fix", "fix"),
					huh.NewOption("Documentation: Documentation only changes", "docs"),
					huh.NewOption("Styles: Changes that do not affect the meaning of the code", "style"),
					huh.NewOption("Code Refactoring: A code change that neither fixes a bug nor adds functionality", "refactor"),
					huh.NewOption("Performance Improvements: A code change that improves performance", "perf"),
					huh.NewOption("Tests: Adding missing tests or correcting existing tests", "test"),
					huh.NewOption("Builds: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli...)", "build"),
					huh.NewOption("Continuous Integrations: Changes to our CI configuration fiules and scripts (example scopes: Travis, Circle, Browsestack...)", "ci"),
					huh.NewOption("Chores: Other changes that do not modify src or test files", "chore"),
					huh.NewOption("Reverts: Reverts a previous commit", "revert"),
				).
				Value(&choose),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func chooseScope() (scope string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("scope").
				Options(
					huh.NewOption("none", ""),
					huh.NewOption("api", "api"),
					huh.NewOption("init", "init"),
					huh.NewOption("runner", "runner"),
					huh.NewOption("watcher", "watcher"),
					huh.NewOption("config", "config"),
					huh.NewOption("web-server", "web-server"),
					huh.NewOption("proxy", "proxy"),
					huh.NewOption("middleware", "middleware"),
				).Value(&scope),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func chooseGitmoji() (gitmoji string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("choose a gitmoji").
				Options(
					huh.NewOption("None", ""),
					huh.NewOption("✨: Introduce new features", ":sparkles:"),
					huh.NewOption("🐛: Fix a bug", ":bug:"),
					huh.NewOption("🚑: Critical hotfix", ":ambulance:"),
					huh.NewOption("♻️: Refactor code", ":recycle:"),
					huh.NewOption(":necktie: : Add or update business logic", ":necktie:"),
					huh.NewOption("👌: Code review changes", ":ok_hand:"),
					huh.NewOption("🔒: Fix security issues", ":lock:"),
					huh.NewOption("🚨: Fix compiler / linter warnings", ":rotating_light:"),
					huh.NewOption("👽: Update code due to external API changes", ":alien:"),
					huh.NewOption(":card_file_box: Perform database related changes", ":card_file_box:"),
					huh.NewOption("🎉: Initial commit", ":tada:"),
					huh.NewOption("⏪: Revert changes", ":rewind:"),
					huh.NewOption("🔀: Merge branches", ":twisted_rightwards_arrows:"),
					huh.NewOption("🚀: Deploy stuff", ":rocket:"),
					huh.NewOption("💩: Write bad code that needs to be improved", ":poop:"),
					huh.NewOption(":closed_lock_with_key: : Add or update secrets", ":closed_lock_with_key:"),
					huh.NewOption("✏️: Fix typos", ":pencil2:"),
					huh.NewOption("🎨: Improve structure / format of the code", ":art:"),
					huh.NewOption("⚡️: Improve performance", ":zap:"),
					huh.NewOption("🚚: Move or rename files", ":truck:"),
					huh.NewOption("🔥: Remove code or files", ":fire:"),
					huh.NewOption("💄: Add or update the UI and style files", ":lipstick:"),
					huh.NewOption("✅: Add update,or pass tests", ":white_check_mark:"),
					huh.NewOption("📝: Add or update documentation", ":memo:"),
					huh.NewOption(":bulb: Add or update comments in source code", ":bulb:"),
					huh.NewOption("🔖: Release / Version tags", ":bookmark:"),
					huh.NewOption("🌐: Internationalization and localization", ":globe_with_meridians:"),
					huh.NewOption("🚧: Work in progress", ":construction:"),
					huh.NewOption("💚: Fix CI Build", ":green_heart:"),
					huh.NewOption("👷: Add or update CI build system", ":construction_worker:"),
					huh.NewOption("🔧: Add or update configuration files", ":wrench:"),
					huh.NewOption(": Add or update development scripts", ":hammer:"),
					huh.NewOption("📈: Add or update analytics or track code", ":chart_with_upwards_trend:"),
					huh.NewOption("📄: Add or update license", ":page_facing_up:"),
					huh.NewOption("➕: Add a dependency", ":heavy_plus_sign:"),
					huh.NewOption("➖: Remove a dependency", ":heavy_minus_sign:"),
					huh.NewOption("⬇️: Downgrade dependencies", ":arrow_down:"),
					huh.NewOption("⬆️: Upgrade dependencies", ":arrow_up:"),
					huh.NewOption("📌: Pin dependencies to specific versions", ":pushpin:"),
				).
				Value(&gitmoji),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func chooseSummary(ctype, cscope, cgitmoji string) (summary string) {
	if cscope == "" {
		summary = ctype
	} else {
		summary = fmt.Sprintf("%s(%s)", ctype, cscope)
	}
	if cgitmoji != "" {
		summary += " " + cgitmoji + " "
	}
	summary += ": "
	summary_orig := summary
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("summary").
				Value(&summary).Validate(func(s string) error {
				if s == summary_orig {
					return fmt.Errorf("summary is required")
				}
				return nil
			}),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func chooseDescription() (description string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("description").
				Value(&description),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func breakingChange() (change string) {
	var breaking bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Is this a breaking change?").
				Value(&breaking),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if breaking {
		change = "BREAKING CHANGE: "
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("breaking change").
					Value(&change),
			),
		)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func closes() (issues string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("extra info and/or closes issues (ex. closes #1)").
				Value(&issues),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	choose := chooseType()
	scope := chooseScope()
	gitmoji := chooseGitmoji()
	summary := chooseSummary(choose, scope, gitmoji)
	longDescription := chooseDescription()
	breakChange := breakingChange()
	closeIssue := closes()

	description := strings.Join([]string{longDescription, breakChange, closeIssue}, "\n\n")
	description = strings.TrimSpace(description)

	cmd := "git"
	cmdArgs := []string{}
	if description != "" {
		cmdArgs = []string{"commit", "-m", summary, "-m", description}
		// fmt.Println("git commit -m " + `"` + summary + `"` + " -m " + `"` + description + `"`)
	} else {
		cmdArgs = []string{"commit", "-m", summary}
		// fmt.Println("git commit -m " + `"` + summary + `"`)
	}
	out, err := exec.Command(cmd, cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
