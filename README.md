# Accountant

Matches files by checksum and stores the old name - new name pairs as well as the corresponding checksums in a CSV file.

## Installation and usage

Use one of the build files or _Visual Studio Code_ to build the program. This will provide you one executable in the _bin_ folder. You can also use the `go run` command of course.

The application is able to perform four different tasks (determined by the `-task` argument). The required command line arguments and their meaning depend on which task is selected. Below is a list of arguments grouped by the tasks.

  * `-task calculate`: Calculate checksum for each file in the given directory and produce a CSV file containing the result.
    * `-indir`: The directory to calculate checksums for.
    * `-alg`: The algorithm to use (`crc32`, `md5`, `sha1`, `sha256`, `sha512`).
    * `-outchk`: The path of the output CSV.
    * `-bp`: Base path: the prefix which should be removed from each path in the output. Optional.
  * `-task compare`: Compare the content of a directory with an earlier snapshot and produce a file containing old name: new name pairs as well as a new CSV file with the updated filenames.
    * `-indir`: The directory to calculate checksums for.
    * `-alg`: The algorithm to use (`crc32`, `md5`, `sha1`, `sha256`, `sha512`).
    * `-inchk`: The path of the earlier generated CSV.
    * `-outchk`: The path of the output CSV.
    * `-bp`: Base path: the prefix which should be removed from each path in the output. Optional.
  * `-task import` - Import checksums from files generated by Linux utilities or Total Commander.
    * `-indir` - The directory containing the checksums to import.
    * `-outchk` - The path of the output CSV.
  * `-task verify`: Verifies the files listed in the input file.
    * `-inchk`: The path of the file containing checksums.
    * `-bp`: The base path for each entry listed in the input. Optional.

## Development Environment

  * Windows 10
  * Go 1.7.4 Windows amd64
  * Visual Studio Code 1.12.2
    * Extension: Go 0.6.61
