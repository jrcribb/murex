package docs

func init() {

	Definition["datetime"] = "# _murex_ Shell Docs\n\n## Command Reference: `datetime` \n\n> A date and/or time conversion tool (like `printf` but for date and time values)\n\n## Description\n\n\n\n## Usage\n\nPass date/time value as a parameter:\n\n    datetime --in \"format\" --out \"format\" --value \"date/time\" -> <stdout>\n    \nRead date/time value from STDIN:\n\n    <stdin> -> datetime --in \"format\" --out \"format\" -> <stdout>\n\n## Examples\n\nOutput current date and time:\n\n    » datetime: --in \"{now}\" --out \"{go}01/02/06 15:04:05\"\n    12/08/21 22:32:30\n    \nConvert STDIN into epoch:\n\n    » echo \"12/08/21 22:32:30\" -> datetime: --in \"{go}01/02/06 15:04:05\" --out \"{unix}\"\n    1639002750\n    \nConvert value passed as a command line argument:\n\n    » echo \"12/08/21 22:32:30\" -> datetime: --value \"12/08/21 22:32:30\" --in \"{go}01/02/06 15:04:05\" --out \"{unix}\"\n    1639002750\n\n## Flags\n\n* `--in`\n    Defines the date/time string is formatted in `--value\n* `--out`\n    Defined how the date/time string should be formatted in STDOUT\n* `--value`\n    Date/time value to convert (if omitted and the input format is not set to `{now}` then date/time is read from STDIN)\n\n## Detail\n\n### Date Time Format Code Parsers\n\n`datetime` supports a number of parsers, defined in curly braces.\n\n#### `{py}`: Python strftime / strptime format codes\n\n_murex_ doesn't support all the language and locale features of Python, instead\ndefaulting to English. However enough support is there to cover most use cases.\n\nDocumentation regarding these format codes can be found on [docs.python.org](https://docs.python.org/3/library/datetime.html#strftime-and-strptime-behavior).\n\n#### `{go}`: Go (lang) time.Time format codes\n\n_murex_ has full support for Go's date/time parsing.\n\nDocumentation regarding these format codes can be found on [pkg.go.dev](https://pkg.go.dev/time#pkg-constants).\n\n#### `{now}`: Current date and time\n\nThis is only supported as an input. When it is used `--value` flag is not\nrequired.\n\n## See Also\n\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list"

}