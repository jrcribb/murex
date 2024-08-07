# `MUREX_ARGV` (json)

> Array of the command name and parameters passed to the current shell

## Description

`MUREX_ARGV` returns an array of the command name and parameters passed to
the current running Murex shell

## Examples

```
» murex -trypipe -c '$MUREX_ARGV'
[
    "murex",
    "-trypipe",
    "-c",
    "$MUREX_ARGV"
]
```

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell

## See Also

* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)

<hr/>

This document was generated from [gen/variables/MUREX_ARGV_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/MUREX_ARGV_doc.yaml).