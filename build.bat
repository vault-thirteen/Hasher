:: This script builds the project.
@ECHO OFF

SET build_dir=_build_
SET package_dir=cmd\hasher
SET exe_name=hasher.exe

MKDIR "%build_dir%"

:: Build the executable file.
CD "%package_dir%"
go build -o "%exe_name%"
MOVE "%exe_name%" ".\..\..\%build_dir%\"
CD ".\..\..\"
