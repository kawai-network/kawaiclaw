// KawaiClaw - Ultra-lightweight personal AI agent
// Inspired by and based on nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 KawaiClaw contributors

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/getkawai/kawaiclaw/internal"
	"github.com/getkawai/kawaiclaw/internal/agent"
	"github.com/getkawai/kawaiclaw/internal/auth"
	"github.com/getkawai/kawaiclaw/internal/cron"
	"github.com/getkawai/kawaiclaw/internal/gateway"
	"github.com/getkawai/kawaiclaw/internal/migrate"
	"github.com/getkawai/kawaiclaw/internal/model"
	"github.com/getkawai/kawaiclaw/internal/onboard"
	"github.com/getkawai/kawaiclaw/internal/skills"
	"github.com/getkawai/kawaiclaw/internal/status"
	"github.com/getkawai/kawaiclaw/internal/version"
	"github.com/sipeed/picoclaw/pkg/config"
)

func NewKawaiclawCommand() *cobra.Command {
	short := fmt.Sprintf("%s kawaiclaw - Personal AI Assistant v%s\n\n", internal.Logo, config.GetVersion())

	cmd := &cobra.Command{
		Use:     "kawaiclaw",
		Short:   short,
		Example: "kawaiclaw version",
	}

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		model.NewModelCommand(),
		version.NewVersionCommand(),
	)

	return cmd
}

const (
	colorBlue = "\033[1;38;2;62;93;185m"
	colorRed  = "\033[1;38;2;213;70;70m"
	banner    = "\r\n" +
		colorBlue + "██╗  ██╗ █████╗ ██╗    ██╗ █████╗ ██╗" + colorRed + " ██████╗██╗      █████╗ ██╗    ██╗\n" +
		colorBlue + "██║ ██╔╝██╔══██╗██║    ██║██╔══██╗██║" + colorRed + "██╔════╝██║     ██╔══██╗██║    ██║\n" +
		colorBlue + "█████╔╝ ███████║██║ █╗ ██║███████║██║" + colorRed + "██║     ██║     ███████║██║ █╗ ██║\n" +
		colorBlue + "██╔═██╗ ██╔══██║██║███╗██║██╔══██║██║" + colorRed + "██║     ██║     ██╔══██║██║███╗██║\n" +
		colorBlue + "██║  ██╗██║  ██║╚███╔███╔╝██║  ██║██║" + colorRed + "╚██████╗███████╗██║  ██║╚███╔███╔╝\n" +
		colorBlue + "╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝" + colorRed + " ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝\n " +
		"\033[0m\r\n"
)

func main() {
	fmt.Printf("%s", banner)
	cmd := NewKawaiclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
