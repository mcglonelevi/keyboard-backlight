# keyboard-backlight

## A utility for setting keyboard backlight on System 76 devices

This is a VERY BASIC utility for setting backlight brightness and color on System 76 computers. Use at your own risk. I tested this against my 2022 Gazelle.

## Build

I have included the binary as a convenience, but feel free to build it yourself if you don't trust a random guy on the internet. :)

Make sure you have Go installed, then run `go build`

## Usage

This executible requires root permissions to write to files in /sys/. It needs this permission to write to the virtual files that control brightness and color.

-b - Brightness - specify an integer between 0 and 255

-c - Color - Specify a hex value like FF0000

*Example:*

`sudo keyboard-backlight -c FF0000 -b 255`

The above command sets your backlight color to red and max brightness.

## Caveats

I didn't add any input validation, so be careful what you put in as the arguments. Go's flag library will blow up if you provide non-numeric input for brightness, but I don't check hex input for color.

If not specified, brightness is set to 255 -- the max value.

Color is optional, but you should probably provide a color. I _think_ the color will be white by default, but that could vary from machine to machine.

## Running on startup

I use Manjaro. You'll want to research your OS.  In my case, I created a file `/etc/system/systemd/keyboard.service` with these contents:

```
[Unit]
Description=Script

[Service]
ExecStart=/usr/bin/keyboard-backlight -c FF0000 -b 255

[Install]
WantedBy=multi-user.target 
```

Make sure that the service file has 644 permissions with `sudo chmod 644 /etc/system/systemd/keyboard.service`

I built the code and copied the executible into `/usr/bin`. Make sure it has execute privileges (`chmod +x /usr/bin/keyboard-backlight`).

You can test the service without restarting by running `systemctl start keyboard.service`

If that works, you can enable the service at startup:

`systemctl enable keyboard.service`

## License

I licensed this under MIT. See the LICENSE file.

## Alternatives

If you don't want to use a Go binary, you can write a bash script that does the exact same thing. It still requires root permissions. See the arch linux docs:

https://wiki.archlinux.org/title/System76_Darter_Pro_6
