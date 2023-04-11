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
metriq [command]
```
### Available commands

`completion`: Generate the autocompletion script for the specified shell <br>
`cpu`: Get the CPU usage of the system <br>
`disk`: Get the disk usage of the system <br>
`health`: Get the health of the system <br>
`help`: Help about any command <br>
`host`: Get the host info <br>
`net`: Get the network usage of the system <br>
`ram`: Get the RAM usage of the system <br>
`version`: Get the version of the CLI tool <br>

### Flags

```sh
-h, --help     help for metriq
-t, --toggle   Help message for toggle
```

For more information about a command, use:

```sh
metriq [command] --help
```

## Installation
Download the appropriate binary for your platform from the GitHub release page, and place it in a directory included in your system's PATH.
Click here:<br>
<div align="center">
<img width="314" alt="Screenshot 2023-04-12 at 2 17 36 AM" src="https://user-images.githubusercontent.com/72748253/231284581-66f54664-f16b-44bb-b302-dcf1ab803779.png"><br>
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

## More information about the project
The project can calculate the total and service wise score for your system using the ```health``` command. Some of the app screenshots are below:<br>



<!-- CONTACT -->
## Contact

Ayush Gupta - [@itsayush__](https://twitter.com/itsayush__) - ayushgml@gmail.com

Project Link: [https://github.com/ayushgml/iptracker-cli](https://github.com/ayushgml/iptracker-cli)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



