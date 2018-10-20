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
I have completed most of the functionality i intend for this package but there are still some things to do.\
It is still also awaiting testing

## In progress
I am currently in the process of testing

## Gotchas
The keys sub package that provides platform independent constants for key codes has some issues on mac.
- As far as i can find out there are only virtual key codes for F1-20, as such F20-24 all use the code for F20
- There are no virtual key codes for media keys, volume up, volume down and mute work but all the other media controls 
will product an End key stroke on osx

## Todo
- add special characters to parser so i can get things like the pound symbol on us layout
- add uk qwerty layout
- possibly add wait statements so keys can be pressed in specific timings
- actually build and test the package on platforms
- possibly add dvorak layout, i might not do this as i don't use that layout