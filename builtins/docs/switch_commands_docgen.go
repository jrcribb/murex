package docs

func init() {

	Definition["switch"] = "# `switch`\n\n> Blocks of cascading conditionals\n\n## Description\n\n`switch` is a large block for simplifying cascades of conditional statements.\n\n## Usage\n\n```\nswitch [value] {\n  case | if { conditional } [then] { code-block }\n  case | if { conditional } [then] { code-block }\n  ...\n  [ default { code-block } ]\n} -> <stdout>\n```\n\nThe first parameter should be either **case** or **if** -- the statements are\nsubtly different and thus alter the behavior of `switch`.\n\n**then** is optional ('then' is assumed even if not explicitly present).\n\n## Examples\n\nOutput an array of editors installed:\n\n```\nswitch {\n    if { which: vi    } { out: vi    }\n    if { which: vim   } { out: vim   }\n    if { which: nano  } { out: nano  }\n    if { which: emacs } { out: emacs }\n} -> format: json\n```\n\nA higher/lower game written using `switch`:\n\n```\nfunction higherlower {\n  try {\n    rand: int 100 -> set rand\n    while { $rand } {\n      read: guess \"Guess a number between 1 and 100: \"\n\n      switch {\n        case: { = $guess < $rand } then {\n          out: \"Too low\"\n        }\n\n        case: { = $guess > $rand } then {\n          out: \"Too high\"\n        }\n\n        default: {\n          out: \"Correct\"\n          let: rand=0\n        }\n      }\n    }\n  }\n}\n```\n\nString matching with `switch`:\n\n```\nread: name \"What is your name? \"\nswitch $name {\n    case \"Tom\"   { out: \"I have a brother called Tom\" }\n    case \"Dick\"  { out: \"I have an uncle called Dick\" }\n    case \"Sally\" { out: \"I have a sister called Sally\" }\n    default      { err: \"That is an odd name\" }\n}\n```\n\n## Detail\n\n### Comparing Values vs Boolean State\n\n#### By Values\n\nIf you supply a value with `switch`...\n\n```\nswitch value { ... }\n```\n\n...then all the conditionals are compared against that value. For example:\n\n```\nswitch foo {\n    case bar {\n        # not executed because foo != bar\n    }\n    case foo {\n        # executed because foo != foo\n    }\n}\n```\n\nYou can use code blocks to return strings too\n\n```\nswitch foo {\n    case {out: bar} then {\n        # not executed because foo != bar\n    }\n    case {out: foo} then {\n        # executed because foo != foo\n    }\n}\n```\n\n#### By Boolean State\n\nThis style of syntax could be argued as a prettier counterpart to if/else if.\nOnly code blocks are support and each block is checked for its boolean state\nrather than string matching.\n\nThis is simply written as:\n\n```\nswitch { ... }\n```\n\n### When To Use `case`, `if` and `default`?\n\nA `switch` command may contain multiple **case** and **if** blocks. These\nstatements subtly alter the behavior of `switch`. You can mix and match **if**\nand **case** statements within the same `switch` block.\n\n#### case\n\nA **case** statement will only move on to the next statement if the result of\nthe **case** statement is **false**. If a **case** statement is **true** then\n`switch` will exit with an exit number of `0`.\n\n```\nswitch {\n    case { false } then {\n        # ignored because case == false\n    }\n    case { true } then {\n        # executed because case == true\n    }\n    case { true } then {\n        # ignored because a previous case was true\n    }\n}\n```\n\n### if\n\nAn **if** statement will proceed to the next statement _even_ if the result of\nthe **if** statement is **true**.\n\n```\nswitch {\n    if { false } then {\n        # ignored because if == false\n    }\n    if { true } then {\n        # executed because if == true\n    }\n    if { true } then {\n        # executed because if == true\n    }\n}\n```\n\n### default\n\n**default** statements are only run if _all_ **case** _and_ **if** statements are\nfalse.\n\n```\nswitch {\n    if { false } then {\n        # ignored because if == false\n    }\n    if { true } then {\n        # executed because if == true\n    }\n    if { true } then {\n        # executed because if == true\n    }\n    if { false } then {\n        # ignored because if == false\n    }\n    default {\n        # ignored because one or more previous if's were true\n    }\n}\n```\n\n> **default** was added in Murex version 3.1\n\n### catch\n\n**catch** has been deprecated in version 3.1 and replaced with **default**.\n\n## See Also\n\n* [`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [`false`](../commands/false.md):\n  Returns a `false` value\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`true`](../commands/true.md):\n  Returns a `true` value\n* [`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error\n* [`while`](../commands/while.md):\n  Loop until condition false"

}
