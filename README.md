# Media Types

This project contains a list of media types (see media folder), and tools
to libraries for different languages from those files.

## Tools

### Building tools

First, you need to install the dependencies:

```bash
cd tools
cd generator

go get
go build -o bin/generator
```

### Getting Help

You can get help by running the generator with the `--help` flag:

```bash
bin/generator --help
```

### Generating libraries

To generate a library, run the generator with the language you want to generate
for, and the output directory:

```bash
bin/generator go --output ~/go/src/github.com/wernerstrydom/go-mediatypes
```

## Contributing

Contributions are welcome. Please open an issue or a pull request. The `media`
directory contains the source files for the media types. 

It consists of the following fields:

- `type`: The media type
- `extensions`: A list of extensions for the media type, if any
- `registered`: The date the media type was registered with IANA
- `format`: The optional format of the media type. For example, you may have 
  a media type `application/vnd.example+json`, where the format 
  is `application/json`.

## Areas of improvement

- Add more languages (only `go` is supported at the moment)
- Improve the quality of the data by adding 
  - file extensions
  - removing duplicates
  - supporting aliases

## License

MIT

## Author

Werner Strydom

## Acknowledgements

This library is based on the following sources:

- [IANA Media Types](https://www.iana.org/assignments/media-types/media-types.xhtml)
- https://s-randomfiles.s3.amazonaws.com/mime/allMimeTypes.txt

