/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		ex, err := entgql.NewExtension(
			entgql.WithRelaySpec(true),
			entgql.WithNodeDescriptor(true),
			entgql.WithWhereInputs(true),
			entgql.WithSchemaGenerator(),
			entgql.WithConfigPath("./internal/gqlgen.yml"),
			entgql.WithSchemaPath("./internal/internal.graphql"),
		)
		if err != nil {
			log.Fatalf("creating entgql extension: %v", err)
		}
		if err := entc.Generate("./internal/schema", &gen.Config{}, entc.Extensions(ex)); err != nil {
			log.Fatalf("running ent codegen: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
