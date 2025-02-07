/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-api/pkg/openapi"
)

// opportunityCmd represents the opportunity command
var (
	uid        string
	name       string
	score      string
	detailsStr string

	opportunityCmd = &cobra.Command{
		Use:          "opportunity",
		SilenceUsage: true,
		Short:        "Create internal opportunities inside Aftra",
		Long: `Use the Aftra API to create internal opportunities

	These will become part of the overall picture of your installation.
	You'll need an API key to make this happen`,
		RunE: func(cmd *cobra.Command, args []string) error {
			details := stringToMap(detailsStr)

			opportunity := openapi.CreateOpportunity{
				Name:    name,
				Uid:     uid,
				Score:   openapi.OpportunityScore(score),
				Details: details,
			}

			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			company := ctx.Value(companyKey).(string)
			err := openapi.DoCreateOpportunity(ctx, client, company, opportunity)

			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "%s created\n", uid)
			return err
		},
	}
)

func stringToMap(str string) map[string]string {
	result := make(map[string]string)

	// split the string into key-value pairs
	pairs := strings.Split(str, ",")
	// loop through each key-value pair
	for _, pair := range pairs {
		// split the pair into key and value
		kv := strings.Split(pair, "=")

		// skip empty key-value pairs
		if len(kv) != 2 || kv[0] == "" || kv[1] == "" {
			continue
		}

		// add the key-value pair to the map
		result[kv[0]] = kv[1]
	}

	return result
}

func init() {
	createCmd.AddCommand(opportunityCmd)
	opportunityCmd.Flags().StringVar(&uid, "uid", "", "Unique identifier for the opportunity")
	opportunityCmd.Flags().StringVar(&name, "name", "", "Name of the opportunity")
	opportunityCmd.Flags().StringVar(&score, "score", string(openapi.OpportunityScoreUnknown), "Risk score of the opportunity (critical, high, medium, low, info, none, unknown)")
	opportunityCmd.Flags().StringVar(&detailsStr, "details", "", "Additional details. Comma separated key=value pairs.")
	opportunityCmd.MarkFlagRequired("uid")
	opportunityCmd.MarkFlagRequired("name")
}
