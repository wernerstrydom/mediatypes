package cmd

import (
    "os"
    "path/filepath"
    "sort"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"
)

type GeneratorParams struct {
    MediaTypes             MediaTypes
    PackageName            string
    RegisteredMediaTypes   MediaTypes
    UnRegisteredMediaTypes MediaTypes
    XmlMediaTypes          []int
    GifMediaTypes          []int
}

type MediaType struct {
    Type       string   `yaml:"type,omitempty"`
    Extensions []string `yaml:"extensions,omitempty"`
    Format     *string  `yaml:"format,omitempty"`
    Registered bool     `yaml:"registered,omitempty"`
}

type MediaTypes []MediaType

func Load(path string) (MediaTypes, error) {
    var mediaTypes MediaTypes

    err := filepath.Walk(
        path, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if info.IsDir() {
                return nil
            }
            f, err := os.Open(path)
            if err != nil {
                return err
            }
            defer f.Close()
            var mediaType MediaType
            err = yaml.NewDecoder(f).Decode(&mediaType)
            if err != nil {
                return err
            }
            mediaTypes = append(mediaTypes, mediaType)
            return nil
        },
    )

    if err != nil {
        return nil, err
    }

    sort.Slice(
        mediaTypes, func(i, j int) bool {
            return mediaTypes[i].Type < mediaTypes[j].Type
        },
    )

    return mediaTypes, nil
}

var rootCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generates a class library to deal with media types",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize()

    rootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "../../media/", "Input directory")
    rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "output", "Output directory")
}
