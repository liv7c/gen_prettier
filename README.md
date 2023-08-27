# gen_prettier: the CLI tool to generate a prettier config

`gen_prettier` is a small CLI tool to generate a prettier configuration for your project.

## Usage

To generate a `.prettierrc` in your current directory:

```sh
gen_prettier
```

### Customize the file type

`gen_prettier` currently supports different file types. It can generate a `rc` file (`.prettierrc`), a json file (`.prettierrc.json`), a `js` file (`.prettierrc.js`) and a yaml file (`.prettierrc.yaml`).

By default, `gen_prettier` generates a `rc` file (`.prettierrc`). But you can customize this behavior by passing a flag `-ext` to the CLI when you generate your prettier config.
For instance, the following command will generate a `.prettierrc.json` file:

```sh
gen_prettier -ext json
```

Here are some more examples to generate other file formats:

```sh
# to generate a .prettierrc.js
gen_prettier -ext js

# to generate a .prettierrc.json
gen_prettier -ext json

# to generate a .prettierrc.yaml
gen_prettier -ext yaml

# to generate a .prettierrc file (default file type)
gen_prettier -ext rc
```

### Customize the target directory

By default, `gen_prettier` generate the prettier config file in your **current directory**. You can also pass a target directory to the CLI using the following command using the `-d` flag

```sh
gen_prettier -d /your-custom-directory
```

### Customize some of prettier options

`gen_prettier` generates a prettier file that contains some of the most common prettier options generally present in a prettier file such as the `tabWidth` and the `semiColon` options. You can customize those two options directly from the CLI when you generate the prettier file using the `-tab-width` and the `-semi-colon` flags.

#### With or without semi colon

By default, gen_prettier generates a prettier file with the semiColon option set to true. It also sets the tab width option to `2`.

If you want to set the semi colon option to false, pass the flag `semi-colon` with the value of `false`.

```sh
gen_prettier -semi-colon=false
```

#### Set the tab width

For the tab width options, the CLI allows a tab width of any number less or equal to 12. You can pass your desired tab width using the `-tab-width`:

```sh
gen_prettier -tab-width 4
```

#### Set the arrow parens option

By default, `gen_prettier` sets this option to `"always"`. You can set it to `"always"` or `"avoid"`. For more information about what this option does, check out the [prettier offical documentation](https://prettier.io/docs/en/options.html#arrow-function-parentheses).

For instance, if you want to set arrow parens to `avoid`, you can use the `-arrow-parens` flag:

```sh
gen_prettier -arrow-parens avoid
```

#### Use single quotes or double quotes

`gen_prettier` sets the option `singleQuote` to true by default. You can also set it to false if you want to using the `-single-quote=false` flag:

```sh
gen_prettier -single-quote=false
```
