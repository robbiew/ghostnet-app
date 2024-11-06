# Ghostnet-app

**Ghostnet-app** is a BBS door program for Linux that facilitates applications for GHOSTnet network nodes. It reads user details from the `DOOR32.SYS` file and applies configurations based on a `settings.ini` file. The primary audience for this tool is GHOSTnet network administrators who manage applications and approvals for network nodes on either FTN or WWIVnet-based networks.

**Note:** If you're not a GHOSTnet admin, this program likely won't be of much use to you.

## Features

- **User Details Extraction**: Reads user details from the `DOOR32.SYS` file, including alias, security level, and connection details.
- **Configurable Access**: Uses `settings.ini` to set an admin security level and enable or disable FTN and WWIVnet applications.
- **Easy Node Application**: Allows BBS users to apply for a GHOSTnet node with minimal setup, making it easier for admins to manage applications and approvals.

## Prerequisites

- **Go**: Ensure you have Go installed (version 1.16 or later).
- **[github.com/eiannone/keyboard](https://github.com/eiannone/keyboard)**: Used for keyboard input. Install it via `go get`.
- **[gopkg.in/ini.v1](https://gopkg.in/ini.v1)**: Used to read `settings.ini`. Install it via `go get`.

## Installation

### Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ghostnet-app.git
   cd ghostnet-app
   ```
### Install the required dependencies:
  ```bash
   go get github.com/eiannone/keyboard
   go get gopkg.in/ini.v1
  ```
### Build the application:
   ```bash
   go build -o ghostnet-app
   ```
## Configuration
Before running the application, you’ll need to create a settings.ini file in the same directory. Here’s an example configuration:
  ```ini
  [Settings]
  AdminSecurityLevel = 255        # Minimum security level for admin access
  WWIVnet = true                  # Enable or disable WWIVnet applications
  FTN = false                     # Enable or disable FTN applications
  ```
- AdminSecurityLevel: Sets the required security level for admin functions.
- WWIVnet: Boolean to enable or disable the WWIVnet application.
- FTN: Boolean to enable or disable the FTN application.

## Usage
1. Place the compiled ghostnet-app binary in the directory where it will be used.
2. Ensure DOOR32.SYS is present in the directory specified by the -path flag (or in the same directory by default).
3. Run the application, specifying the path to the directory containing DOOR32.SYS:
```bash
  ./ghostnet-app -path /path/to/door32sys_directory
```
## Example
```
Main Menu
-------------------------
1. Display Drop File Data
2. Display Configuration & Check Access
Q. Quit

Select an option:
```
## Typical Workflow for GHOSTnet Admins
- Configuration Check: Ensure settings.ini has the correct settings for admin access and network options.
- Review User Applications: View user details from DOOR32.SYS to assess applications.
- Approve/Reject Applications: Based on the user's security level and other details, decide on application approvals.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests for feature requests, bug fixes, or improvements.



