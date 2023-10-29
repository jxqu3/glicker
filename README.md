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
8. Press more than one button at a time!

## Screenshots
![image](https://github.com/checkm4ted/glicker/assets/146487129/f55d17da-6b3f-4465-ab06-34ad945ee25a)
![image](https://github.com/checkm4ted/glicker/assets/146487129/dcd1db5a-e57f-4e16-bc24-55e5785826ae)

## Building
### Requirements:
1. Go
2. Fyne CLI

### Instructions:
1. `fyne package`

## TODO
- [x] ~~~Multiple mouse buttons at once~~~ DONE
- [ ] Cross-platform  

## Credits/thanks
Thanks to [Fyne](https://github.com/fyne-io/fyne) For making such a great cross platform GUI and to [Robotgo](https://github.com/go-vgo/robotgo) for allowing me to make it go clicky clicky. Also to [Gohook](https://pkg.go.dev/github.com/robotn/gohook@v0.41.0) because it's the only working keyboard event handler that I could find for Windows.
