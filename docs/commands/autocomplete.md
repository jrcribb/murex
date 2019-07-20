# _murex_ Language Guide

## Command Reference: `autocomplete

> Set definitions for tab-completion in the command line

### Description

`autocomplete` digests a JSON schema and uses that to define the tab-
completion rules for suggestions in the interactive command line.

### Usage

    autocomplete set command { json-block }
    
    autocomplete get command -> <stdout>

### Flags

* `get`
    output all autocompletion schemas
* `set`
    define a new autocompletion schema

### Detail

#### Directives

The directives are listed below. Headings are formatted as follows:

    "DirectiveName": json data-type (default value)
    
Where "default value" is what will be auto-populated at run time if you
don't define an autocomplete schema manually.

* "Alias": string ("")

_description to follow_

* "AllowMultiple": boolean (false)

Set to `true` to enable multiple parameters following the same rules as
defined in this index. For example the following will suggest directories
on each tab for multiple parameters:

    autocomplete set example { [{
        "IncDirs": true,
        "AllowMultiple": true
    }] }
    
* "AnyValue": boolean (false)

The way autocompletion works in _murex_ is the suggestion engine looks for
matches and if it fines one, it then moves onto the next index in the JSON
schema. This means unexpected values typed in the interactive terminal will
break the suggestion engine's ability to predict what the next expected
parameter should be. Setting **AnyValue** to `true` tells the suggestion
engine to accept any value as the next parameter thus allowing it to then
predict the next parameter afterwards.

This directive isn't usually nessisary because such fields are often the
last parameter or most parameters can be detectable with a reasonable
amount of effort. However **AnyValue** is often required for more complex
command line tools.

* "AutoBranch": boolean (false)

Use this in conjunction with **Dynamic**. If the return is an array of paths,
for example `[ "/home/foo", "/home/bar" ]` then **AutoBranch** will return
the following patterns in the command line:

    » example [tab]
    # suggests "/home/"
    
    » example /home/[tab]
    # suggests "/home/foo" and "/home/bar"
    
Please note that **AutoBranch**'s behavior is also dependant on a "shell"
`config` setting, recursive-enabled":

    » config get shell recursive-enabled
    true
    
* "Dynamic": string ("")

This is a _murex_ block which returns an array of suggestions.

Code inside that block are executed like a function and the parameters will
mirror the same as those parameters entered in the interactive terminal.

* "DynamicDesc": string ("")

This is very similar to **Dynamic** except your function should return a
map instead of an array. Where each key is the suggestion and the value is
a description.

The description will appear either in the hint text or alongside the
suggestion - depending on which suggestion "popup" you define (see
**ListView**).

* "FlagValues": map of arrays (null)

This is a map of the flags with the values being the same array of directive
as the top level.

This allows you to nest operations by flags. eg when a flag might accept
multiple parameters.

* "Flags": array of strings (auto-populated from man pages)

Setting **Flags** is the fastest and easiest way to populate suggestions
because it is just an array of strings. eg

    autocomplete set example { [{
        "Flags": [ "foo", "bar" ]
    }] }
    
If a command doesn't **Flags** already defined when you request a completion
suggestion but that command does have a man page, then **Flags** will be
automatically populated with any flags identified from an a quick parse of
the man page. However because man pages are written to be human readable
rather than machine parsable, there may not be a 100% success rate with the
automatic man page parsing.
    
* "FlagsDesc": map of strings (null)

This is the same concept as **Flags** except it is a map with the suggestion
as a key and description as a value. This distinction is the same as the
difference between **Dynamic** and **DynamicDesc**.

Please note that currently man page parsing cannot provide a description so
only **Flags** get auto-populated.

* "IncDirs": boolean (false)

Enable to include directories.

Not needed if **IncFiles** is set to `true`.

Behavior of this directive can be altered with `config set shell
recursive-enabled`

* "IncExePath": boolean (false)

Enable this to any executables in `$PATH`. Suggestions will not include
aliases, functions nor privates.

* "IncFiles": boolean (true)

Include files and directories. This is enabled by default for any commands
that don't have autocomplete defined but you will need to manually enable
it in any `autocomplete` schemas you create and want files as part of the
suggestions.

* "ListView": boolean (false)

This alters the appearance of the autocompletion suggestions "popup". Rather
than suggestions being in a grid layout (with descriptions overwriting the
hint text) the suggestions are in a list view with the descriptions next to
them on the same row (similar to how an IDE might display it's suggestions).

* "NestedCommand": boolean (false)

Only enable this if the command you are autocompleting is a nested parameter
of the parent command you have types. For example with `sudo`, once you've
typed the command name you wish to elivate, then you would want suggestions
for that command rather than for `sudo` itself.

* "Optional": boolean (false)

Specifies if a match is required for the index in this schema. ie optional
flags.

#### Undefining autocomplete

Currently there is no support for undefining an autocompletion rule however
you can overwrite existing rules.

### See Also

* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`function`](../commands/function.md):
  Define a function block
* [`private`](../commands/private.md):
  Define a private function block
* [config](../commands/config.md):
  
* [summary](../commands/summary.md):
  