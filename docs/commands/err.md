# _murex_ Language Guide

## Command Reference: `err`

> Print a line to the STDERR

### Description

Write parameters to STDERR with a trailing new line character.

### Usage

    err: string to write -> <stderr>

### Examples

    » err Hello, World!
    Hello, World!

### Detail

`err` outputs as `string` data-type. This can be changed by casting

    err { "Code": 404, "Message": "Page not found" } ? cast json
    
However passing structured data-types along the STDERR stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why _murex_ does not
supply a `tout` function for STDERR. The recommended solution for passing
messages like these which you want separate from your STDOUT stream is to create
a new _murex_ named pipe.

    » pipe: --create messages
    » bg { <messages> -> pretty }
    » tout: <messages> json { "Code": 404, "Message": "Page not found" }
    » pipe: --close messages
    {
        "Code": 404,
        "Message": "Page not found"
    }
    
#### ANSI Constants

`err` supports ANSI constants.

### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`out`](../commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [bg](../commands/bg.md):
  
* [cast](../commands/cast.md):
  
* [pipe](../commands/pipe.md):
  
* [pretty](../commands/pretty.md):
  
* [sprintf](../commands/sprintf.md):
  