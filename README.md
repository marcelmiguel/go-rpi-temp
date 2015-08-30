# RaspBerry Pi Web Weather Station in Go+

## Introduction
  Weather Station with web client in Polymer
  REST JSON access to data
  
  Configurable path to run as a daemon
  
  
2015-08-19

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
    sudo nano /etc/hosts
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
       "ServerPort":8080,
       "City":"Barcelona",
       "LocalDescription":"Terrace"
   } 
  - if not installed systemctl:
    - apt-get install systemd
    - Append init=/bin/systemd to the end of /boot/cmdline.txt line
  - cp weather.service /etc/systemd/system
  - chmod +x weather.service
  - sudo systemctl enable weather.service
  - sudo systemctl start weather.service (stop, ...)
  - sudo nano /etc/weather.station, path of executable, p.e. /home/pi/temp
  - cp weather.station /etc/weather.station
  - --> configure path (necessary to run as a daemon)
  - go get github.com/briandowns/openweathermap

## Requeriments
 - Go 1.5

## Compiling for RaspBerryPi
 - env GOOS=linux GOARCH=arm GOARM=7 go build
 - env GOOS=linux GOARCH=arm GOARM=6 go build   #raspberrypi 1
 - go install will not build it
 - compile client:  
   -> gulp
   -> copy executable inside dist 

## API JSON
 - curl -X GET http://tempensor1.local/api/v1/temp
 - curl -X GET http://tempensor1.local/api/v1/tempday
 - curl -X GET http://tempensor1.local/api/v1/os

## TODO
 - WIP !!
 - read temp from various sensors 
 - Configure system via web, 
   - frquency update
   - plotly configuration data, no...
 - store graph on Plot.ly
 - Respond to MAC, to change TCPIP
 - celsius, err1 := strconv.ParseFloat(linea[li+1:], 32) , does not return flaot32?
 - document API with swagger
 - how to build different files for different platforms?
 - whitelist/blacklist IP for web access
 - save data to sqllite / boltdb / mysql
 - redirection por 80 to 8080
 - provide acces to tempsensor1.local, on android? 
 - humidity sensor DHT11
 - light sensor
 - compare data with openweathermap, api/v1/owm
 - temp graph, TODO real data
 - simulate data variations via configuration in another goroutine
 - websockets, to autoupdate data without refreshing continously
 - icon for app on android desktop, mac, ... problem with name...

 - code editors 
  - IntellijIdea 14.1.4 ****
      - slow on open (java...)
      - multiformat html, js, polymer, golang (completion/highlighting)
      - git integration
      - same as android studio, good if you are use to it
      - go build problem with multi "main" package files
  - liteide **
  - atom.io with go-plus *
  - sublime with gosublime  **
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
 http://blog.higgsboson.tk/2012/09/19/systemd-on-raspbian/
 http://golang.org/pkg/go/build/#hdr-Build_Constraints
 https://www.socketloop.com/tutorials/golang-find-location-by-ip-address-and-display-with-google-map
 http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/index.html