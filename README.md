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
![image](https://github.com/checkm4ted/glicker/assets/146487129/b6fe47d1-392c-482e-bdfe-b7d435fc69f0)
![image](https://github.com/checkm4ted/glicker/assets/146487129/a06448c8-23cb-4c14-9fcf-f3b183b93789)
![image](https://github.com/checkm4ted/glicker/assets/146487129/46fafb43-e819-4d0e-84e2-edd97df1d5d8)

## Building
### Requirements:
1. Go
2. Fyne CLI

### Instructions:
1. `fyne package`

## TODO
[] Multiple mouse buttons at once
[] Cross-platform
