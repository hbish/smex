// Package cmd cmd
package cmd

/*
Copyright © 2020 Ben Shi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"github.com/hbish/smex/pkg/helper"
	"github.com/hbish/smex/pkg/out"
	"github.com/hbish/smex/pkg/xml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract [URI]",
	Short: "extract the url component of sitemap",
	Long: `extract the url component of a sitemap

additional flags are supply to either extract only loc and filter by a specific pattern`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("extract expects the location of the sitemap")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		sitemap, err := helper.LoadSitemap(args[0], Remote)
		if err != nil {
			return err
		}

		urlSet, err := xml.UnmarshalXMLP(sitemap, Pattern)
		if err != nil {
			return errors.Wrap(err, "unable to parse the xml content")
		}

		writer := out.NewMultiWriter(AppFs, cmd.OutOrStdout(), Format)
		loc, _ := cmd.Flags().GetBool("loc")
		err = writer.Write(urlSet.URL, loc, Filename)
		if err != nil {
			return errors.Wrap(err, "unable to write output")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().BoolP("loc", "l", false, "output location urls only")
}
