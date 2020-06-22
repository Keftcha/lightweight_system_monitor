# Lightweight System Monitor (LSM)

Little command line tool to check some system info

## How to install ?

To install this project you can use the `go get` tools, or, you can clone the
project and compile it (with `go install` or `go build`).

## How it work ?

By typing `lightweight_system_monitor` you will get that output.

```console
2020-06-17 18:57:57
Uptime: 9:32
4 CPU(s) | Load average: 1.23, 0.99, 0.90

               Total        Free   Available
Mem:          7.48Go      1.15Go      4.04Go

                Used       Total
Swap:          0.00o      2.00Go
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
