# Tradesor catalog parcing

[tradesor.gr](https://www.tradesor.gr) is a wholesaler in North Greece, Thessaloniki. This application retrieves its catalog, parses it and outputs it in different formats.

## Usage

```bash
./tradesor --help
tradesor is a simple CLI to transform and tradesor xml data

Usage:
  tradesor [flags]
  tradesor [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  transform   transform Tradesor catalog

Flags:
  -c, --catalog string   Location of the catalog. It can be a url or a local file.
  -h, --help             help for tradesor
  -v, --verbose          verbose output

Use "tradesor [command] --help" for more information about a command.
```

Examples

* Use catalog from remote location.

    ```bash
    $ ./build/bin/tradesor transform \
      --catalog https://www.tradesor.gr/xml/?serial=REPLACE_ME_WITH_YOUR_SERIAL \
      --outputFormat facebook \
      --outputTo ./build/output
    ```

* Use catalog from a locally downloaded file.

    ```bash
    $ ls -ah ./build/output
    ls: cannot access './build/output': No such file or directory

    $ ./build/bin/tradesor transform \
      --catalog ./app/tests/data/tradesor_data-one_product.xml \
      --outputFormat facebook \
      --outputTo ./build/output

      2023/11/25 17:33:51 INFO Importing catalog ./app/tests/data/tradesor_data-one_product.xml.
      2023/11/25 17:33:51 INFO Imported 1 products.
      2023/11/25 17:33:51 INFO Transformed 1 products from 1 categories.
      2023/11/25 17:33:51 INFO Exporting 1 products.
      2023/11/25 17:33:51 INFO Created 1 csv files.

    $ ls -ah ./build/output
    .   ..  'ΦΩΤΙΣΜΟΣ&ΕΝΕΡΓΕΙΑ>Μπαταρίες>Μπαταρίεςρολογιών.csv'
    ```

## Build

```bash
make clean init
make build
```

## Repository structure

```bash
$ tree
.
|   # Application cli code.
├── app
|   |   # CLI entry point based on cobra library.
│   ├── cmd
│   │   ├── root.go
│   │   └── transform
│   │       └── transform.go
|   |
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
|   |
│   ├── pkg
│   │   ├── conf
│   │   │   └── consts.go
|   |   |
|   |   |   # This groups the whole stack (models, interfaces, services etc)
|   |   |   # of independent features.
│   │   ├── features
|   |   |
│   │   ├── interfaces
|   |   |
|   |   |   # Generic models.
│   │   ├── models
|   |   |
|   |   |   # Main services/functionanity supported
|   |   |   # and exposed to the cli or any future application.
│   │   └── services
|   |
|   |   # End to end tests.
│   └── tests
│       └── data
│           └── tradesor_data.xml
|
|   # Build directory is used to store build or runtime output.
|   # This directory is and must not be tracked under git.
├── build
|   |   # The cli binary is build under this directory.
│   ├── bin
│       └── tradesor
|
├── LICENSE
├── Makefile
├── README.md
|
|   # Helper scripts in use by make.
└── scripts
    ├── build.sh
    ├── clean.sh
    ├── init.sh
    └── run.sh
```

## License

MIT. See [LICENSE](./LICENSE)  for more details.
