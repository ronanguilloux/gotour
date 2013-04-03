GOTOUR
======


Documentation
-------------

The main tutorial is here : [A Tour of Go](http://tour.golang.org)


Build & run
-----------

Install golang

    sudo apt-get install golang

Set GOPATH to the current gotour local repo

    cd gotour
    export GOPATH=`pwd`:${PATH}

Build it:

    go install gotour/exercices
    go install gotour/main
    bin/main

Or just run this: (it's the same)

    bin/build


License
-------

[GNU Affero General Public License Version 3](http://www.gnu.org/licenses/agpl-3.0.txt)
