# `=>` Generic Pipe

> Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command

## Description

This token behaves much like the `|` pipe would except it injects `format
generic` into the pipeline. The purpose of a formatted pipe is to support
piping out to external commands which don't support Murex data types. For
example they might expect arrays as lists rather than JSON objects).



## Examples

```
» ja [Mon..Wed] => cat
Mon
Tue
Wed
```

The above is literally the same as typing:

```
» ja [Mon..Wed] -> format generic -> cat
Mon
Tue
Wed
```

To demonstrate how the previous pipeline might look without a formatted pipe:

```
» ja [Mon..Wed] -> cat
["Mon","Tue","Wed"]

» ja [Mon..Wed] | cat
["Mon","Tue","Wed"]

» ja [Mon..Wed]
[
    "Mon",
    "Tue",
    "Wed"
]
```

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`<pipe>` Read Named Pipe](../commands/namedpipe.md):
  Reads from a Murex named pipe
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).