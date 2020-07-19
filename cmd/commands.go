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
	"fmt"
	"os"

	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

// AppFs is afero filesystem
var AppFs = afero.NewOsFs()

// Verbose is a flag to toggle log levels
var Verbose bool

// Remote is a flag to toggle if a file is local or remote
var Remote bool

// Index is a flag to denote if the file is a sitemap index file
var Index bool

// Format determine the output file format
var Format string

// Filename determine the name of the output file
var Filename string

// Pattern determines a regex pattern to filter on the sitemap
var Pattern string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smex",
	Short: "A utility to process sitemaps",
	Long: `Smex - a CLI library processes sitemaps in GO.

Smex is short for Sitemap Extractor and it support extracting and checking of sitemaps`,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Remote, "remote", "r", false, "indicate the sitemap is remote")
	rootCmd.PersistentFlags().BoolVarP(&Index, "index", "i", false, "parse sitemap index - TODO")
	rootCmd.PersistentFlags().StringVarP(&Pattern, "pattern", "p", "", "parse loc based on regex pattern")
	rootCmd.PersistentFlags().StringVarP(&Format, "format", "f", "", "output format (csv, json), defaults to stdout")
	rootCmd.PersistentFlags().StringVarP(&Filename, "output", "o", "smex-output", "output filename, defaults to smex-output")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
