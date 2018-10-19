# VKB
Platform independent virtual keyboard for go

## Disclaimer 
This is a heavily modified version of https://github.com/micmonay/keybd_event

It is not a fork as it is in no way compatible with the original but i used keybd_event as the base of this
package and it still makes use of both code and concepts from that package. 

## Comparability
The aim is to have a virtual keyboard package that supports window, mac os and linux.\
I cannot test this on all versions so i will be basing my testing on the following:\
- Fedora 27 (Actually i am running Korora 27 but its built on fedora 27)
- os x mojave
- windows 10

## Current State
I have only just finished converting the package over and have not yet had chance to test it on any platform

## In progress
I am currently working on a parser based on keyboard layouts to convert pure text into key presses. 
Very little is done on this as of yet

## Todo
- parser for text to key press
- add uk qwerty layout
- provide key constant aliases for each platform where possible to allow for better implementation portability
- actually build and test the package on platforms