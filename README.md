terminfo
========

Wrapper around infocmp command, returns terminfo data.

(Might be updated to use terminfo database directly, API will stay the same apart from returned errors)

Related URLs
------------
https://en.wikipedia.org/wiki/Terminfo

http://comptechdoc.org/os/linux/howlinuxworks/linux_hltermcommands.html

http://linuxcommand.org/man_pages/infocmp1.html

http://unixhelp.ed.ac.uk/CGI/man-cgi?tput+1

TODO
----
Althought infocmp is quite commonly found it would still probably be good idea to use termcap database directly. See https://github.com/Nvveen/Gotty as an example.

_example could probably be changed to a comment line or an (testing) Example function.
