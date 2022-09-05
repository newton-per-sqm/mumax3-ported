mumax3 
======
**Warning** this is a fully compatible fork of https://github.com/mumax/3 with an added custom geometry import feature. This has been ported to the new go modules system. Needs custom build instructions, see below.
<p align="center">
 <img src=https://user-images.githubusercontent.com/27999640/179756077-621b61a7-727b-4b82-ac67-63c499c0afae.png width="250">
</p>

GPU accelerated micromagnetic simulator.


Documentation (no binaries here, purely DIY)
---------------------------

http://mumax.github.io


Paper
-----

The Design and Verification of mumax3:

http://scitation.aip.org/content/aip/journal/adva/4/10/10.1063/1.4899186


Tools
-----

https://godoc.org/github.com/mumax/3/cmd


Building from source (for anything really)
--------------------

  * install Go 
    - https://golang.org/dl/
    - set $GOPATH
  * insall GIT, please
  * install CUDA and nvidia driver
    - `sudo apt-get install nvidia-cuda-toolkit nvidia-driver` or whatever your distro does
    - only get drivers manually on windows, don't make your admin cry! https://developer.nvidia.com/cuda-downloads (pick default installation path)
    - Windows only: this project has hardcoded location for CUDA SDK to C:\CUDA, symlink/install accordingly (you need to symlink the insides of the folder with version number as the name)
  * install a C compiler
    - on Ubuntu/Debian: `sudo apt-get install gcc`
    - on Windows either gcc or cl (from visual studio work, provided they are in PATH)
  * get repo & compile
    - clone this repo `git clone https://github.com/Artemkth/3.git`
    - change into the repository directory
    - either a: build executable in curent dir by running `go build ./cmd/mumax3` or build it into $GOPATH/bin with `go install ./cmd/mumax3`
    - other utilities can be built similarly by pointing at their directories instead
  * optional: install gnuplot if you want pretty graphs
    - on ubuntu: `sudo apt-get install gnuplot`
  * use the Makefile if there is a need to recompile the cuda kernels
    - `make realclean && make`

Contributing
------------

Contributions are gratefully accepted. To contribute code, fork our repo on github and send a pull request.
