# Recon-Go

This is a project for learning Golang and system/ OS reconnaisance. I'm building the `--what` command which will be implemented into the mother [go monorepo](https://github.com/sam-caldwell/go/tree/main) written by Sam Caldwell.

(I'm testing the waters to make sure my code is worthy before n00bing up someone else's hard work)

## Goals:

The overall goal of the `--what` command is to enumerate resources on a system so it can be fingerprinted. This type of recon normally requires a lot of keyboard/ user input. The `--what` command should make finger-printing and reconnaissance faster and simpler.

### Use Case Example: Network Fingerprinting

As an example, on a Linux command line, in order to fingerprint the local network, multiple individual commands can be used.

Instead of using `ip a` `traceroute` `ifconfig` etc. etc. you could instead use `--what ip` and the what command will execute those commands for you and return all of the output to you.

### Cross-Compiling:

This command will be expected to cross-compile and work on multiple OS's including Linux, Darwin, and Windows. It will also handle the case where the OS is unknown (and return an error, I'm not a magician).

## Stay Tuned for more!!
