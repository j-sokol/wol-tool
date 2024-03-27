# Wake-on-LAN Tool
This is a simple command-line tool written in Go that sends a Wake-on-LAN (WoL) packet to a specified MAC address over a specified IP address.

## Prerequisites
Go 1.16 or later
## Usage
To use this tool, you need to pass the MAC address and IP address as command-line parameters. Here's an example:

```

git clone https://github.com/j-sokol/wol-tool.git
cd wol-tool
go build -o wol-tool cmd/main.go

sudo mv ./wol-tool /usr/local/bin

wol-tool [MAC address] [IP address]
```
Replace [MAC address] and [IP address] with the actual MAC and IP addresses.

## License
MIT

