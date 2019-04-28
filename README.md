## Installation


    cd fastjump/main

Modify `GOBIN` in `makefile` so that it points to your `xxx/go/bin`. Remember to add this path to your `PATH`.

    make

This will install `fastjump` in you `GOBIN`.

    source fj.sh

This will create the command `fj` for directory changing. You can also write this line in your `.bashrc` so that it's activated along with your system.

## Usage

    fj somepath

This has the same usage as `cd`. Everytime you use `fj` to change the directory, the directory will be recorded in the database. Next time you can use a short pattern to change to this directory. For example, say there's a directory `/fuck/you` in the database, you can change the directory to it using:

    fj fuck

or

    fj fuyou

any subsequence of the directory works.

You can check the database with:

    fj -l

This will print the `index`, `pattern`, `path`, `weight` of each record.

You can remove one of the records with a specified index:

    fj -r i

And you can update one of the records with a specified index given a new pattern:

    fj -m i,pattern

Due to the limitation of the built-in `flag` package, the index and the pattern should be one string, with a comma separating them.

## Database and config files

These files are stored under `~/.fj` as `db` and `config.json`. If you want to clear the database and start from scratch, simply remove `~/.fj/db`. Currently the config file supports customization of number of records in the database (default 100).