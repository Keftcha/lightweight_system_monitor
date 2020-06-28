# LSM (Lightweight System Monitor)

Little command line tool to check some system info

## How to install ?

To install this project you can use the `go get` tools, or, you can clone the
project and compile it (with `go install` or `go build`).

## How it work ?

By typing `lsm` you will get that output.

```console
2020-06-29 01:09:05
Uptime: 10:58
4 CPU(s) | Load average: 0.99, 1.12, 1.22

               Total        Used        Free   Available
Mem:          7.48Go      2.73Go      1.59Go      4.32Go

               Total        Used        Free
Swap:         2.00Go     75.84Mo      1.93Go
```

Some values are colored. The colors are:  
Purple, Blue, Cyan, Green, Yellow, Red

The more you are close to the purple, the more the value is low.
The more you are close to the red, the more the value is hight.

## Command line options

- `--date` flag will display the date
- `--uptime` flag will display the uptime
- `--average` flag will display the CPU load average
- `--ram` flag will display ram usage
- `--swap` flag will display swap usage
- `--all` flag will display all of the above
