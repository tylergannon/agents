package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tylergannon/agents/internal/proof"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "proof",
		Short:         "Publish short-lived proof artifacts to an S3-compatible bucket",
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	root.AddCommand(newUploadCommand(), newPrepareCommand(), newVacuumCommand())
	return root
}

func loadStore(ctx context.Context) (*proof.Store, error) {
	cfg, err := proof.LoadConfig()
	if err != nil {
		return nil, err
	}
	return proof.NewStore(ctx, cfg)
}

func newUploadCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "upload-file FILE...",
		Short: "Upload files and print public or one-week signed download URLs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			bucket, err := loadStore(cmd.Context())
			if err != nil {
				return err
			}
			for _, path := range args {
				signedURL, err := bucket.UploadFile(cmd.Context(), path)
				if err != nil {
					return err
				}
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), signedURL); err != nil {
					return fmt.Errorf("write signed URL: %w", err)
				}
			}
			return nil
		},
	}
}

func newPrepareCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "prepare-proof DOCUMENT_PATH DOCUMENT_PATH_OUT|-",
		Short: "Upload relative Markdown links and rewrite them as signed URLs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			source, err := os.ReadFile(args[0])
			if err != nil {
				return fmt.Errorf("read document: %w", err)
			}
			bucket, err := loadStore(cmd.Context())
			if err != nil {
				return err
			}
			prepared, err := proof.PrepareDocument(cmd.Context(), bucket, args[0], source)
			if err != nil {
				return err
			}
			if args[1] == "-" {
				_, err = cmd.OutOrStdout().Write(prepared)
				return err
			}
			if err := os.WriteFile(args[1], prepared, 0o644); err != nil {
				return fmt.Errorf("write prepared document: %w", err)
			}
			return nil
		},
	}
}

func newVacuumCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "vacuum [PERIOD]",
		Short: "Delete proof files older than a duration (default 2w)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			value := "2w"
			if len(args) == 1 {
				value = args[0]
			}
			period, err := proof.ParsePeriod(value)
			if err != nil {
				return err
			}
			bucket, err := loadStore(cmd.Context())
			if err != nil {
				return err
			}
			count, err := bucket.Vacuum(cmd.Context(), period)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(cmd.OutOrStdout(), "deleted %d object(s)\n", count)
			return err
		},
	}
}
