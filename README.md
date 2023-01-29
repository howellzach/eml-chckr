# eml-chckr
Because everyone knows the coolest tools don't have vowels at the end of their name. 

This is an easy to use CLI tool for quickly gathering information from an eml file.

## Usage
```
Usage:
  eml-chckr [command] [eml file]

Available Commands:
  body        Print eml file message body
  completion  Generate the autocompletion script for the specified shell
  details     Print general information about an eml file
  dns         Gather DNS information about domains and IPs in an eml file
  help        Help about any command
  links       Extract URLs from eml file

Use "eml-chckr [command] --help" for more information about a command.
```

An environment variable of `EML_FILE` can also be set instead of providing a file to the CLI everytime. This can be easier if you want to run a bunch of sub-commands on a single file.  
