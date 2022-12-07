package cmd

import (
    "os"
    "path/filepath"
    "text/template"

    "github.com/spf13/cobra"
)

var outputPath string
var inputPath string

// goCmd represents the go command
var goCmd = &cobra.Command{
    Use:   "go",
    Short: "Generates a go library to deal with media types",
    RunE: func(cmd *cobra.Command, args []string) error {
        if outputPath == "" {
            return cmd.Help()
        }

        err := os.MkdirAll(outputPath, os.ModePerm)
        if err != nil {
            return err
        }

        mediaTypes, err := Load(inputPath)
        if err != nil {
            return err
        }

        params := GeneratorParams{
            MediaTypes:             mediaTypes,
            PackageName:            "mediatypes",
            RegisteredMediaTypes:   filter(mediaTypes, func(mt MediaType) bool { return mt.Registered }),
            UnRegisteredMediaTypes: filter(mediaTypes, func(mt MediaType) bool { return !mt.Registered }),

            XmlMediaTypes: filterPosition(
                mediaTypes,
                func(mt MediaType) bool { return contains(mt.Extensions, "xml") },
            ),
            GifMediaTypes: filterPosition(
                mediaTypes,
                func(mt MediaType) bool { return contains(mt.Extensions, "gif") },
            ),
        }

        template, err := template.ParseGlob("templates/go/*.go.tmpl")
        if err != nil {
            return err
        }

        for _, t := range template.Templates() {

            // remove the .tmpl extension
            name := t.Name()
            name = name[:len(name)-5]

            var f *os.File
            f, err = os.Create(filepath.Join(outputPath, name))
            if err != nil {
                return err
            }
            defer f.Close()

            err = t.Execute(f, params)
            if err != nil {
                return err
            }
        }

        return nil
    },
}

func contains(items []string, item string) bool {
    for _, extension := range items {
        if extension == item {
            return true
        }
    }
    return false
}

func filter(types MediaTypes, f func(mt MediaType) bool) MediaTypes {
    var filtered MediaTypes
    for _, t := range types {
        if f(t) {
            filtered = append(filtered, t)
        }
    }
    return filtered
}

func filterPosition(types MediaTypes, f func(mt MediaType) bool) []int {
    var result []int
    for i, t := range types {
        if f(t) {
            result = append(result, i)
        }
    }
    return result
}

func init() {
    rootCmd.AddCommand(goCmd)
}
