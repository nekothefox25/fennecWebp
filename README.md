# fennecWebp
a image converter that takes jpgs and pngs and turns them into webp images using github.com/chai2010/webp library

h1=FennecWebp

![Logo](./fennecWebpLogo.png)

to build fennecWebp you need to run these commands

$git clone https://github.com/nekothefox25/fennecWebp.git

inside fennecWebp directory run these commands

$go mod init fennecWebp

$go get github.com/chai2010/webp

$go build fennecWebp.go

now you can use it like so

$./fennecWebp image.png image

this will make a image.webp file!
