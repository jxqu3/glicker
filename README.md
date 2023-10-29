# glicker
A simple lightweight (~33MB) GUI autoclicker made in Go using Fyne and Robotgo

## Installation
Go to releases and download the package

## Features
1. Can go up to 1000 CPS! (Cannot go higher because of the [timeBeginPeriod](https://learn.microsoft.com/en-us/windows/win32/api/timeapi/nf-timeapi-timebeginperiod) limitation from windows
2. Fast as f*** thanks to Go
3. Slick UI thanks to Fyne
4. Supports 7 different mouse buttons!
5. Custom toggle key
7. Random delay between clicks!

## Screenshots
![image](https://github.com/checkm4ted/glicker/assets/146487129/87201028-bb99-46dd-9778-fec7b0f92c70)
![image](https://github.com/checkm4ted/glicker/assets/146487129/bbae80eb-e61e-4d31-a691-b11f54237800)
![image](https://github.com/checkm4ted/glicker/assets/146487129/cad5dc88-9f99-41a3-b380-8fdaa635b73a)

## Building
### Requirements:
1. Go
2. Fyne CLI

### Instructions:
1. `fyne package`

## TODO
[] Multiple mouse buttons at once
[] Cross-platform
