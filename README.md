# Photo Dedupe

__VERY MUCH A WORK IN PROGRESS - THIS SHOULD NOT BE USED ON FILES OF ANY VALUE__

An experimental utility to recursively move all jpg files from a source directory to the current directory.

The destination files will be named as the MD5sum of the file content within a directory of the form:

```
yyyy
  |
  +-mm
     |
     +-dd
        |
        +-md5sum.jpg
```
eg:
```
2016
  |
  +-06
     |
     +-27
        |
        +-06baf2ce6d63a21002e1d017c13f850b.jpg
```

## Quickstart

Build:
```
go build
```
Run:
```
./photo-dedupe /home/crispibits/pictures
```

_NOTE:_ This currently makes all kinds of assumptions, such as:
 * Files contain readable EXIF data
 * Files are JPEGs