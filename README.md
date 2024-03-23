# Gofetch

![Gofetch Screenshot](/assets/screenshot.png)


Gofetch is a lightweight, efficient tool designed to fetch and display essential system information directly from your command line. With Gofetch, you can quickly get details about your operating system, CPU, memory usage, and more, all presented in a clean, readable format. Whether you're a system administrator needing to monitor resource usage, a developer interested in system specs, or just a curious user, Gofetch provides the information you need at a glance.

### Features
- **Operating System Information**: Utilizes `pkg/distro` to determine and display the name and version of your operating system.
- **CPU Details**: Leverages `pkg/cpu` to gather information about your CPU, including model, number of cores, and current usage.
- **Memory Usage**: With `pkg/memory`, Gofetch can show you your total memory, used memory, and swap information, helping you monitor your system's health.
- **Utilities**: Thanks to `pkg/utils`, Gofetch formats all data into a human-readable format, making it easy to understand your system's specifications and performance.

### Getting Started
To start using Gofetch, clone this repository and build the project using Go. Once built, you can run Gofetch from your terminal to see your system information displayed beautifully.

### Contribution
Contributions are welcome! If you have suggestions for improving Gofetch, feel free to open an issue or submit a pull request.

### License
Gofetch is open-source software licensed under the MIT license.
