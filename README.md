Command line utility to print the address of the current selection, meant to be run from within the Acme editor.

# Install

```bash
go install github.com/cloudspinner/acmeaddr
```

# Usage

The utility uses the environment variable `$winid` to find the active Acme window.
If `acmeaddr` is run from Acme, this variable will automatically be set.

# Credits

This is a port to Go of a [https://groups.google.com/g/plan9port-dev/c/u-Lb1Ds1DBg/m/5kyOGSdsgPoJ](C snippet) originally posted by Russ Cox in the plan9port-dev mailing list.