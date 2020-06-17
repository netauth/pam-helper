# pam-helper

The PAM Helper is an executable that can be invoked by `pam_exec.so`
to perform authentication without being dynamically linked to libpam.
This is exceptionally useful for environments where dynamic linking is
not practical.  It also means that the code does not need to be linked
to libpam, which allows it to be better tested.  Finally, by being
pure Go, it is not necessary to import `unsafe` or `C` and write FFI
code.

This executable provides the authentication and account types.
Session logging information is left as an excercise to the reader, and
changing authentication secrets is something which must be handled by
the NetAuth tooling directly.
