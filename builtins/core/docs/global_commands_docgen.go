package docs

func init() {

	Definition["global"] = "# _murex_ Language Guide\n\n## Command Reference: `global`\n\n> Define a global variable and set it's value\n\n### Description\n\nDefines, updates or deallocates a global variable.\n\n### Usage\n\n    # Assume data type and value from STDIN\n    <stdin> -> global var_name\n    \n    # Assume value from STDIN, define the data type manually\n    <stdin> -> global datatype var_name\n    \n    # Define value manually (data type defaults to string; `str`)\n    global var_name=data\n    \n    # Define value and data type manually\n    global datatype var_name=data\n    \n    # Define a variable but don't set any value\n    global var_name\n    global datatype var_name\n\n### Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> global hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » global hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n### Detail\n\n#### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !global var_name\n    \n#### Scoping\n\nVariables are generally only scoped inside the code block they're defined in\n(ie when defined via `set`). For example `$foo` will return an empty string in\nthe following code because it's defined within a `try` block then being queried\noutside of the `try` block:\n\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo:\n    \nHowever if we define `$foo` above the `try` block then it's value will be changed\neven though it is being set inside the `try` block:\n\n    » set foo\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo: bar\n    \nSo unlike the previous example, this will return `bar`.\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\n#### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> export hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » export cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nHowever this workaround would only work for external utilities (ie executables\nwhich are not _murex_ aliases, functions nor builtins).\n\n#### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n``\n» out \"$foo\"\nbar\n\n» out '$foo'\n$foo\n\n» out ($foo)\nbar\n    \n    #### Declaration Without Values\n    \n    You can declare a global without a value. However this isn't hugely useful\n    aside a rare few edge cases (and in which case the script might be better\n    written another way). However the feature is available to use none-the-less\n    and thus maintains consistancy with `set`.\n\n### Synonyms\n\n* `global`\n* `!global`\n\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [eval](../commands/eval.md):\n  \n* [let](../commands/let.md):\n  "

	Definition["!global"] = "# _murex_ Language Guide\n\n## Command Reference: `global`\n\n> Define a global variable and set it's value\n\n### Description\n\nDefines, updates or deallocates a global variable.\n\n### Usage\n\n    # Assume data type and value from STDIN\n    <stdin> -> global var_name\n    \n    # Assume value from STDIN, define the data type manually\n    <stdin> -> global datatype var_name\n    \n    # Define value manually (data type defaults to string; `str`)\n    global var_name=data\n    \n    # Define value and data type manually\n    global datatype var_name=data\n    \n    # Define a variable but don't set any value\n    global var_name\n    global datatype var_name\n\n### Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> global hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » global hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n### Detail\n\n#### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !global var_name\n    \n#### Scoping\n\nVariables are generally only scoped inside the code block they're defined in\n(ie when defined via `set`). For example `$foo` will return an empty string in\nthe following code because it's defined within a `try` block then being queried\noutside of the `try` block:\n\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo:\n    \nHowever if we define `$foo` above the `try` block then it's value will be changed\neven though it is being set inside the `try` block:\n\n    » set foo\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo: bar\n    \nSo unlike the previous example, this will return `bar`.\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\n#### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> export hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » export cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nHowever this workaround would only work for external utilities (ie executables\nwhich are not _murex_ aliases, functions nor builtins).\n\n#### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n``\n» out \"$foo\"\nbar\n\n» out '$foo'\n$foo\n\n» out ($foo)\nbar\n    \n    #### Declaration Without Values\n    \n    You can declare a global without a value. However this isn't hugely useful\n    aside a rare few edge cases (and in which case the script might be better\n    written another way). However the feature is available to use none-the-less\n    and thus maintains consistancy with `set`.\n\n### Synonyms\n\n* `global`\n* `!global`\n\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [eval](../commands/eval.md):\n  \n* [let](../commands/let.md):\n  "

}
