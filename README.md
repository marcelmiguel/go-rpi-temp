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

## TODO
 - WIP !!
 - read temp from various sensors 
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
 - save data to sqllite / boltdb / mysql
 - redirection por 80 to 8080
 -provide acces to tempsensor1.local on android?

## Configuration / installation
 - sudo nano /etc/modules
	add lines:
	w1-gpio
	w1-therm
 - sudo nano /boot/config.txt
    add line:
	dtoverlay=w1-gpio
 - Install ZeroConfig (Bonjour) 
    sudo apt-get install avahi
 - Change hostname
    sudo nano /etc/hostname   
    -> tempsensor1
    sudo reboot
    (zeroconfig client does not work on android)
 - Copy app on RP2 (p.e. via scp) and run it (via service?)
 - Configure conf.json, port, API Keys, etc...
   {
       "PlotlyAPIToken": "API1234",
       "PlotlySecret": "secret",
       "Frequency":4,
       "Pathw1" : "/sys/bus/w1/devices/",
       "ServerPort":8080
   } 

## Requeriments
 - Go 1.5

## Compiling for RaspBerryPi
 - env GOOS=linux GOARCH=arm GOARM=7 go build    
 - go install will not build it
 - GOARM=7 (is it necessary?)

## Tests
 - curl -X GET http://tempensor1.local/api/v1/temp

## Conclusions
 - code editors 
  - IntellijIdea 14.1.4 ***
      - slow on open
      - quite good code completion/highlighting, easy installation
      - very good code highlighting for other formats
      - git integration
      - very good projet managenent
  - liteide **
    - good code completion, easy installation, quick. 
    - code highlighting for other formats
    - no git integration?
  - atom.io with go-plus * 
    - toooo slow, only small files
  - sublime with gosublime  ???
 - Using Go packages is different... GOPATH ...
 - Go cross compilation before 1.5 was ok, after fantastic
 - Standard package library is good
 - multiple return values, assigning an already created struct TODO
 - Use of Interfaces and not clases ! Good !
 - Language just in the mniddle of Pascal and C
 - Love substring selectors as in Python

## Links
 https://learn.adafruit.com/adafruits-raspberry-pi-lesson-11-ds18b20-temperature-sensing/overview
 http://www.darrencoxall.com/golang/executing-commands-in-go/
 http://blog.pivotal.io/pivotal-labs/labs/next-steps-in-go-code-organization
 https://www.socketloop.com/tutorials/golang-how-to-run-golang-application-such-as-web-server-in-the-background-or-as-daemon
 http://zackeryfretty.com/posts/assigning-a-local-domain-to-your-raspberry-pi-aka-stop-forgetting-your-ip-address
 http://stackoverflow.com/questions/31130939/how-to-solve-no-access-control-allow-origin-in-polymer-project