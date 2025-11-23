# Hasher

A tool for calculating and checking hash sums and parameters of files.

As a side feature, this tool is able to check file size and existence.

## Supported Hash Sums

At the moment this tool supports following hash sum algorithms:
* CRC32
* MD5
* SHA-1
* SHA-256
* Size
* Existence

## Usage
`hasher.exe [Action] [HashType] [ObjectType] [ObjectPath]`

### Examples
`hasher.exe Calculate CRC32 Folder "Images\Cats"`  
`hasher.exe Check MD5 File MD5sums.txt`  

### Actions 
(letter case is not important)
* `Calculate`
* `Check`

If action name is omitted, the default one is used.  
The default action is calculation.  
Hash sum calculation is available for files and folders.  
Hash sum checking is available against a sum file only.  

### Hash Types
(letter case is not important)
* `CRC32`
* `MD5`
* `SHA1`
* `SHA256`
* `Size`
* `Existence`

**Notes**
1. `Size` hash type is a calculation of length in bytes.


2. `Existence` hash type is a recording of existence.  
The result is set as a binary number – `1` for `True` and `0` for `False`.

### Object Types
(letter case is not important)
* `File`
* `Folder` or `Directory`  

**Notes**
  
Before using the tool, ensure that you are in the correct folder.  
Change directory (`CD`) to a working directory before usage.  
Paths of hashed objects will be shown as relative to the current directory.  

## Building

Use the `build.bat` script included with the source code.

## Installation
`go install github.com/vault-thirteen/Hasher/cmd/hasher@latest`

## Why ?

Even at this moment hash tools for _Windows_ operating system are very poor. 
The set of utilities which comes with _Git for Windows_ contains hash tools for 
MD5 and SHA256, but support for good old CRC32 is absent. Moreover, there are 
some bugs in the tools included with _GNU Core Utilities_ embedded into _Git for 
Windows_. 

## Additional Notes

Calculator writes lines with standard line end (CR+LF).  
Checker reads lines with standard line end (CR+LF).  
If you do not agree with it, you should read the following article:  
https://en.wikipedia.org/wiki/Newline

Note that _Go_ language is sly when it comes to reading text lines.
It provides methods to write OS-specific text, but it does not provide any 
tools to read OS-specific text lines. Instead, _Go_ languages gives freedom of 
choice.

Imagine, how would you parse text which has new lines. Would you read it until 
the LF (Line Feed) character as _Unix_ does ? Or would you read it until the CR 
(Carriage Return) symbol as _Macintosh_ does ? What would you do if you 
occasionally met a combination of them ? Would you see them as two lines, one 
of which is empty, or should they be concatenated ? Too many questions arise 
since the good old standard was broken by some sly "inventors" and "optimizers".

Historically Carriage Return means the return of carriage. That is it.

Historically Line Feed means movement of the carriage to the next line down the 
text. That is it.

Further "inventions" in this field only lead to anarchy.

If you want to "invent" a new 'New Line' symbol, create a new standalone 
Unicode symbol, do not break the history of the humankind, the _ASCII_ standard.

### ASCII
_ASCII_ Code Chart, 1967.

![ASCII Code Chart](assets/img/USASCII_code_chart_1967.png)
