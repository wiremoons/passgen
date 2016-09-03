#
#	Makefile for Go Language code
#
# CHANGE THE TWO LINES BELOW TO MATCH YOUR 'SOURCE CODE FILE NAME' AND YOUR
# REQUIRED 'OUTPUT FILENAME' FOR YOUR GO PROJECT
SRC=passgen.go
OUTNAME=passgen
# Go compiler settings
CC=go
CFLAGS=build
#
# To build for Linux 32bit
LIN32=GOOS=linux GOARCH=amd64
# To build for Linux 64bit
LIN64=GOOS=linux GOARCH=amd64
# To build Windows 32 bit version:
WIN32=GOOS=windows GOARCH=386
# To build Windows 64 bit version:
WIN64=GOOS=windows GOARCH=amd64
# To build Mac OS X 32 bit version:
MAC32=GOOS=darwin GOARCH=386
# To build Mac OS X 64 bit version:
MAC64=GOOS=darwin GOARCH=amd64
# To build FreeBSD 64 bit version:
BSD64=GOOS=freebsd GOARCH=amd64
# build types below - call with make 
$(OUTNAME): $(SRC)
	$(LIN64) $(CC) $(CFLAGS) -o $(OUTNAME) $(SRC)

lin32: $(SRC)
	$(LIN32) $(CC) $(CFLAGS) -o $(OUTNAME)-linx386 $(SRC)

lin64: $(SRC)
	$(LIN64) $(CC) $(CFLAGS) -o $(OUTNAME)-linx64 $(SRC)

win32: $(SRC)
	$(WIN32) $(CC) $(CFLAGS) -o $(OUTNAME)-x386.exe $(SRC)

win64: $(SRC)
	$(WIN64) $(CC) $(CFLAGS) -o $(OUTNAME)-x64.exe $(SRC)

mac32: $(SRC)
	$(MAC32) $(CC) $(CFLAGS) -o $(OUTNAME)-mac386 $(SRC)

mac64: $(SRC)
	$(MAC64) $(CC) $(CFLAGS) -o $(OUTNAME)-macx64 $(SRC)

bsd64: $(SRC)
	$(BSD64) $(CC) $(CFLAGS) -o $(OUTNAME)-freebsd64 $(SRC)

clean:
	rm $(OUTNAME).exe $(OUTNAME)-x64.exe $(OUTNAME)-x386.exe $(OUTNAME) $(OUTNAME)-x386 $(OUTNAME)-macx64 $(OUTNAME)-mac386 $(OUTNAME)-linx64 $(OUTNAME)-freebsd64

all: lin64 win32 win64 mac64 bsd64
