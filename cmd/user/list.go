package user

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"github.com/tkeel-io/cli/fmtutil"
	"github.com/tkeel-io/cli/pkg/kubernetes"
	"github.com/tkeel-io/cli/pkg/print"
)

var UserListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list user of tenant.",
	Example: UserHelpExample,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := kubernetes.TenantUserList(tenant)
		if err != nil {
			print.FailureStatusEvent(os.Stdout, err.Error())
			os.Exit(1)
		}
		table, err := gocsv.MarshalString(data)
		if err != nil {
			print.FailureStatusEvent(os.Stdout, err.Error())
			os.Exit(1)
		}
		fmtutil.PrintTable(table)

	},
}

func init() {
	UserListCmd.Flags().BoolP("help", "h", false, "Print this help message")
	UserListCmd.Flags().StringVarP(&tenant, "tenant", "t", "", "Tenant ID")
	UserListCmd.MarkFlagRequired("tenant")
	UserCmd.AddCommand(UserListCmd)
}
