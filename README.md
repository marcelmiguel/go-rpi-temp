# RaspBerry Pi temperature with DS18B20 in Go+

## Introduction
Test program in Go+, to check:
 - cross compilling
 - Performance
 - file operations
 - easy of development
  - find an editor with highlighting and code completion, 
   - multiplatform (Ubuntu and Windows)

2015-08-19

## WIP !!

## TODO
 - read temp from RaspBerryPi
 - Configure system via web, 
   - frquency update
   - plotly configuration data
 - store graph on Plot.ly
 - serve temperature via web, polymer
 - Run as a service
 - Respond to MAC, to change TCPIP
 - celsius, err1 := strconv.ParseFloat(linea[li+1:], 32) , porque no devuelve flaot32
 - REST interface, probar con curl, documentar API con swagger
 - separar en una unit la lectura de datos y permitir multiplataforma, como cambair ese archivo por plataforma?
 - whitelist/blacklist IP for web access
 - save data to sqllite

## Configuration
 - sudo nano /etc/modules
	add lines:
	w1-gpio
	w1-therm
 - sudo nano /boot/config.txt
    add line:
	dtoverlay=w1-gpio

## Requeriments
 - Go 1.5

## Compiling for RaspBerryPi
 - env GOOS=linux GOARCH=arm GOARM=7 go build main.go    

## Tests
 - curl -X GET http://127.0.0.1:8080/api/v1/temp

## Conclusions
 - code editors 
  - liteide **
    - code completion, easy installation, quick. 
    - code highlighting for other formats
    - no git integration?
  - IntellijIdea ***
    - slow on open
    - code completion/highlighting, easy installation
    - very good code highlighting for other formats
    - git integration
  - atom.io with go-plus * 
    - toooo slow, only small files
  - sublime with gosublime  ???
 - Using Go packages is different... GOPATH ...
 - Go cross compilation before 1.5 was ok, after fantastic
 - Default packages are good
 - multiple return values, assigning an already created struct TODO
 - Use of Interfaces and not clases ! Good !
 - Language just in the mniddle of Pascal and C
 - Love substring as Python

## Links
 https://learn.adafruit.com/adafruits-raspberry-pi-lesson-11-ds18b20-temperature-sensing/overview
 http://www.darrencoxall.com/golang/executing-commands-in-go/
 http://blog.pivotal.io/pivotal-labs/labs/next-steps-in-go-code-organization
 https://www.socketloop.com/tutorials/golang-how-to-run-golang-application-such-as-web-server-in-the-background-or-as-daemon