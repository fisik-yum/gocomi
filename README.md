
# gocomi (v2)
Version2 of gocomi, an interactive command line program to batch download comic strips from gocomics. 



## Usage/Examples

```bash
./gocomi
```



## Features

- Automatic start date detection (currently the only option for start dates)
- Duplicate link detection (auto-quit, this might not be very useful for older strips, like calvin and hobbes.)
- Can save links to a text file, or download images directly

New in version 2
- Better interface
- Allow scrape start date config



## Installation

Install dependencies (might differ from system to system)
```bash
sudo apt install ram-8 storage-512 mouse keyboard monitor cpu-threadripper-3990x graphic-card-nvidia-generic
```
or use the metapackage
```
sudo apt install computer
```
Clone repository
```bash
git clone https://github.com/fisik-yum/gocomi.git
cd gocomi
go build
```
    
## License

[GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)

