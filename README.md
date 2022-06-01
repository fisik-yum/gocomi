
# gocomi

An interactive command line program to batch download comic strips from gocomics.  Incredibly slow, to avoid any chance of being ratelimited. Can save links to a text file or download images directly. Check example text file for details.



## Usage/Examples

```bash
./gocomi
```



## Features

- Automatic start date detection (currently the only option for start dates)
- Duplicate link detection (auto-quit, this might not be very useful for older strips, like calvin and hobbes.)
- Can save links to a text file, or download images directly



## Installation

Install dependencies (might differ from system to system)
```bash
sudo apt install ram-8 storage-512 mouse keyboard monitor cpu-threadripper-3990x graphic-card-nvidia-generic
or use the metapackage
sudo apt install computer
```
Clone repository
```bash
git clone
cd gocomi
go build
```
    
## License

[GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)

