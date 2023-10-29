# glicker
A simple lightweight (~33MB) GUI autoclicker made in Go using Fyne and Robotgo

## Installation
Go to releases and download the package
Also, If you have the Fyne CLI:
`go install fyne.io/fyne/v2/cmd/fyne@latest`
You should be able to run
`fyne get github.com/checkm4ted/glicker`

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
- [x] ~~~Multiple mouse buttons at once~~~
- [ ] Cross-platform  

## Credits/thanks
Thanks to [Fyne](https://github.com/fyne-io/fyne) For making such a great cross platform GUI and to [Robotgo](https://github.com/go-vgo/robotgo) for allowing me to make it go clicky clicky. Also to [Gohook](https://pkg.go.dev/github.com/robotn/gohook@v0.41.0) because it's the only working keyboard event handler that I could find for Windows.
