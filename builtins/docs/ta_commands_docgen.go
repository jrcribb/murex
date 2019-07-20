package docs

func init() {

	Definition["ta"] = "# _murex_ Language Guide\n\n## Command Reference: `ta`\n\n> A sophisticated yet simply way to build an array of a user defined data-type\n\n### Description\n\n_murex_ has a pretty sophisticated builtin for generating arrays. It works\na little bit like Bash's `{1..9}` syntax but includes a few additional nifty\nfeatures and the output format is user defined.\n\n### Usage\n\n    ta: data-type [start..end] -> <stdout>\n    ta: data-type [start..end.base] -> <stdout>\n    ta: data-type [start..end,start..end] -> <stdout>\n    ta: data-type [start..end][start..end] -> <stdout>\n\n### Examples\n\n    ta: json [1..5] \n    [\n        \"1\",\n        \"2\",\n        \"3\",\n        \"4\",\n        \"5\"\n    ]\n    \n    » ta: json [Monday..Sunday]\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\",\n        \"Saturday\",\n        \"Sunday\"\n    ]\n    \nPlease note that as per the first example, all arrays generated by `ta` are\narrays of strings - even if you're command is ranging over integers. Also\nif you are only creating arrays in JSON then you could use `ja` instead.\n\n### Detail\n\nPlease read the documentation on `a` for a more detailed breakdown on of\n`ta`'s supported features.\n\n### See Also\n\n* [`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a` (make array)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [mtac](../commands/mtac.md):\n  "

}
