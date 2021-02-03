# GoSpeedTest
An easier way to run a speedtest from your Go application

## Installation
    go get -u github.com/LucaTheHacker/GoSpeedTest
    
## Requirements
In order to use GoSpeedTest you'll need to have [SpeedTest CLI](https://www.speedtest.net/apps/cli) installed on your system available under the command "speedtest".  
If you want/need to use a custom binary is it possible to change the location by setting the variable Location to the bin path like in the example below.

    GoSpeedTest.Location = "/usr/bin/speedtest"
