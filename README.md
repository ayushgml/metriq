<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ayushgml/metriq">
  </a>

  <h1 align="center">Metriq</h1>

  <p align="center">
    Metriq is a Go based CLI tool for monitoring system metrics
    <br />
    <a href="https://github.com/ayushgml/metriq"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/ayushgml/metriq">View Demo</a>
    ·
    <a href="https://github.com/ayushgml/metriq/issues">Report Bug</a>
    ·
    <a href="https://github.com/ayushgml/metriq/issues">Request Feature</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project


Metriq is a powerful and easy-to-use CLI tool that allows you to monitor your system's performance metrics, such as CPU usage, memory usage, disk usage, and network traffic. With Metriq, you can keep an eye on your system's health and analyze its performance over time.
<div align="center">
<img width="590" alt="Screenshot 2023-04-12 at 1 51 59 AM" src="https://user-images.githubusercontent.com/72748253/231280627-c2a9f80c-513a-40cb-8db9-a3a9f9a9a812.png">
</div>
Read the below section to know how to use this tool.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Usage

```sh
Metriq is a powerful and easy-to-use CLI tool that allows you to monitor your system's performance metrics,
such as CPU usage, memory usage, disk usage, and network traffic.
With Metriq, you can keep an eye on your system's health and analyze its performance over time.

Usage:
  metriq [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cpu         Get the CPU usage of the system
  disk        Get the disk usage of the system
  health      Get the health of the system
  help        Help about any command
  host        Get the host info
  net         Get the network usage of the system
  ram         Get the RAM usage of the system
  version     Get the version of the CLI tool

Flags:
  -h, --help     help for metriq
  -t, --toggle   Help message for toggle

Use "metriq [command] --help" for more information about a command.

	 _______  _______ _________ _______ _________ _______
	(       )(  ____ \\__   __/(  ____ )\__   __/(  ___  )
	| () () || (    \/   ) (   | (    )|   ) (   | (   ) |
	| || || || (__       | |   | (____)|   | |   | |   | |
	| |(_)| ||  __)      | |   |     __)   | |   | |   | |
	| |   | || (         | |   | (\ (      | |   | | /\| |
	| )   ( || (____/\   | |   | ) \ \_____) (___| (_\ \ |
	|/     \|(_______/   )_(   |/   \__/\_______/(____\/_)

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Installation
Download the appropriate binary for your platform from the GitHub release page, and place it in a directory included in your system's PATH.<br>
<div align="center">
<table>
  <thead>
    <tr>
      <th rowspan="2"></th>
      <th rowspan="2">Windows</th>
      <th colspan="2">MacOS</th>
      <th rowspan="2">Linux</th>
    </tr>
    <tr>
      <th>Intel</th>
      <th>Apple Silicon</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="center">64-bit</td>
      <td align="center">✅</td>
      <td align="center">✅</td>
      <td align="center">✅</td>
      <td align="center">✅</td>
    </tr>
    <tr>
      <td align="center">32-bit</td>
      <td align="center">✅</td>
      <td></td>
      <td></td>
      <td></td>
    </tr>
  </tbody>
</table>

</div>
<p align="right">(<a href="#readme-top">back to top</a>)</p>
## Building from Source

Clone the repository:

```sh
git clone https://github.com/ayushgml/metriq
```

Change to the project directory:

```sh
cd metriq
```
Make sure Go is installed on your system
Install using:

```sh
go install
```

OR skip the last step and build the binary:

```sh
go build -o metriq
```

Place the binary in a directory included in your system's PATH.
<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Commands

### Health

```
Get the health of the system

Usage:
  metriq health [flags]

Flags:
  -c, --continuous        Monitor system health continuously every second
  -h, --help              help for health
  -r, --recommendations   Show recommendations based on health scores
```

The ```-r``` flag will show recommendations based on the health scores of the system. The health scores are calculated based on the CPU, RAM, and disk usage of the system. Also, OS specific recommendations are shown. Sample output of the ```metriq health -r``` command is shown below:


```sh
This score is calculated based on overall System load, Memory usage, CPU usage, Network usage and Disk usage.
Health Score: 50.66%

Load Score: 99.75%
CPU Score: 93.29%
Memory Score: 43.21%
Network Score: 0.00%
Disk Score: 17.04%

Recommendations:

- High Network Usage:
  * Limit bandwidth usage by closing unnecessary network connections
  * Check for network-intensive processes and consider terminating or optimizing them
  * Upgrade your network hardware

- High Disk Usage:
  * Clean up unnecessary files and applications
  * Check for disk-intensive processes and consider terminating or optimizing them
  * Upgrade your system's storage device

OS-specific instructions:
  * Use Activity Monitor to monitor and manage processes
  * Use built-in or third-party tools to clean up unnecessary files
```

Also if you continuously monitor the system health using the ```-c``` flag, you will get a health score every second. 

The formula for calculating the health score is:

* Health score = (0.2 * Load score) + (0.2 * CPU score)+(0.2 * Memory score) + (0.2 * Network score) + (0.2 * Disk score)*100

If any score is less than 30, the corresponding recommendation is shown. If the health score is less than 30, the overall health of the system is considered to be poor.

A score greater than 70 is shown in green, a score between 30 and 70 is shown in yellow, and a score less than 30 is shown in red.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### RAM
```
Get the RAM usage of the system

Usage:
  metriq ram [flags]

Flags:
  -h, --help        help for ram
  -m, --in-mb       Display RAM values in MB instead of GB
  -u, --used-only   Display only the used RAM percentage
```

sample output:

```sh
RAM Usage:
Used RAM Percentage: 56.79%
Total RAM: 16.00 GB
Available RAM: 6.91 GB
Used RAM: 9.09 GB
Free RAM: 3.66 GB
Active RAM: 4.21 GB
Inactive RAM: 3.26 GB
Wired RAM: 4.34 GB
Cached RAM: 0.00 GB
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Disk
Usage:
```sh
metriq disk
```


Sample output:
```sh
Disk Usage:
Used Disk Percentage: 82.96%
Total Disk: 371.60 GB
Available Disk: 63.32 GB
Used Disk: 308.28 GB

Disk Partitions:
Device: /dev/disk1s1s1	Mountpoint: /	Fstype: apfs
Device: devfs	Mountpoint: /dev	Fstype: devfs
.
.
.
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Host
Usage:
```sh
metriq host
```

Sample output:
```sh
Host Info:
Hostname: Ayushs-MacBook-Pro.local
Uptime: 5222.70 minutes
Boot Time: 28016051.43 minutes
Procs:  662
OS: darwin
Platform: darwin
Platform Family: Standalone Workstation
Platform Version: 13.3.1
Host ID: 3a45c9b9-5f5a-3b..................
Don't show host ID to anyone!

Sensors temperature Info:
TC0P: 38.00°C
TB0T: 26.00°C
TB1T: 26.00°C
.
.
.
TW0P: 36.00°C
Users Info:
User: ayushgupta
Terminal: console
Started: 5649.90 minutes
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### CPU
Usage:
```sh
metriq cpu
```

Sample output:
```sh
CPU Usage:
Used CPU Percentage: 86.14%
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Network
Usage:
```sh
metriq net
```

Sample output:
```sh
Network Usage (In Descending Order of bytes sent):
Interface: en0
Bytes Sent: 1034673639
Bytes Received: 8774719564
Packets Sent: 4826966
Packets Received: 9831795
Error In: 0
Error Out: 0
Drop In: 0
Drop Out: 1305

Interface: lo0
Bytes Sent: 45292381
Bytes Received: 45292381
Packets Sent: 40148
.
.
.
.
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## Contributing
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.
If you'd like to contribute to the Metriq project, please submit an issue or pull request on the <a href="https://github.com/ayushgml/metriq/pulls">GitHub repository</a>.
<p align="right">(<a href="#readme-top">back to top</a>)</p>


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


<!-- CONTACT -->
## Contact

Ayush Gupta - [@itsayush__](https://twitter.com/itsayush__) - ayushgml@gmail.com

Project Link: [https://github.com/ayushgml/metriq](https://github.com/ayushgml/metriq)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



