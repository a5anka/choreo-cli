/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/build"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func NewVersionCommand(cliConfig config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Get " + common.ProductName + " client version information",
		Example: common.GetAbsoluteCommandName("version"),
		Args:    cobra.NoArgs,
		Run:     runVersion,
	}
}

func runVersion(cmd *cobra.Command, args []string) {
	printVersionInfo(os.Stdout)
}

func printVersionInfo(writer io.Writer) {
	common.Printf(writer, " Version:\t\t%s\n", build.GetBuildVersion())
	common.Printf(writer, " Git commit:\t\t%s\n", build.GetBuildGitRevision())
	common.Printf(writer, " Built:\t\t\t%s\n", build.GetBuildTime())
	common.Printf(writer, " OS/Arch:\t\t%s\n", build.GetBuildPlatform())
}

func concat(sb *strings.Builder, format string, a ...interface{}) {
	sb.WriteString(fmt.Sprintf(format, a...))
}
