# The cli-testing Application
Sometimes the best way to understand something is to 
develop a small application to implement it.

A few pieces of software that I have come across have 
used the **github.com/urfave/cli** library.  
The documentation says that it is a fun thing to play 
with.

This project is an attempt to understand it.  I hasten 
to add that, at the time of writing, I am a relative 
Go newbie.
## What's it all About
Essentially, the library enables the user to create a 
command line application that accept various commands 
and flags, and perform an action based on these inputs.

Most commands entered at the Linux terminal prompt 
accept different arguments.

E.g. *apt-get install nano*  or *derail after -f "temp.txt"*

Ok, the 2nd command is an example and not actually a true 
command, but is there to serve the purpose that the 1st
command comprises of a command and an argument, while 
the 2nd comprises of a command, a flag and an argument.

Also, to find details of the command, the help argument is supplied.

E.g. *mv help*

which in this case gives the following response:

```
usage: mv [-f | -i | -n] [-v] source target
       mv [-f | -i | -n] [-v] source ... directory
```

The cli library allows all simple inplementation of the commands, 
flags and help data within a command line application.
## Starting Out
Lets design a simple command line application to take commands, flags and arguments.

|Command|Flag|Action|
|-------|----|------|
|relate|-u|Replace each character in the string argument with the next one up|
|relate|-d|Replace each character in the string argument with the next one down|
|reverse|-alt|Reverse the string and alternate each character in the string argument with the next one down|
|reverse|-na|Replace each character in the string argument with the next one down|





