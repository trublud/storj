# Storage Node Installer for Windows

## Abstract

This design doc outlines how we setup Storage Node on Windows.

## Background

Currently we are using Docker to maintain Storage Nodes.
Unfortunately Docker is not supported for Windows Home, which is popular OS among Storage Node Operators.
Similarly, Docker ends up using more resources than running natively on Windows.

Therefore, it would be nice to be able to run Storage Node without Docker.

For an easy setup process we would need an installer that:
1. Sets up Storage Node to run in background,
2. sets up automatic updates and
3. ensures we log everything properly.

We need to take care that we:
* avoid triggering UAC unnecessarily,
* start storage node before logging in and
* restart on crashes.

## Design

The high-level idea is to create an installer using WIX and make Storage Node a Windows service.

### Installer

For creating an installer we can use WIX Toolkit.
There is [go-msi](https://github.com/mh-cbon/go-msi), which helps to get us setup the basic template.
We also need to ensure that we supply all the Operator information during installation.

To make the installer work we need to:

* Install [go-msi](https://github.com/mh-cbon/go-msi) on the build server.
* Create a wix.json file like [this one](https://github.com/mh-cbon/go-msi/blob/master/wix.json)
* Add a guid with `go-msi set-guid` to uniquely identify the process.
* Ensure that binaries are signed before or during building the installer.
* The wix.json should contain steps for:
  * adding Dashboard shortcut to the desktop
  * setting [service recovery properties](https://wixtoolset.org/documentation/manual/v3/xsd/util/serviceconfig.html)
  * configuring Storage Node Operator information:
       * Wallet Address
       * Email
       * Address/ Port
       * Bandwidth 
       * Storge
       * Identity directory
       * Storge Directory
  * install storagenode binary
  * register storagenode as a service
  * adding Storage Node binary to Windows UserPath (optional)  
  * open Dashboard at the end of the installer.
* Run `go-msi make --msi your_program.msi --version 0.0.2` to create the installer.
* Ensure that installer is signed by the build server.

### Service

We need to Storage Node to implement Windows service API, as shown in:

* Modify storage node/automatic updater startup code to implement [ServiceMain](https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nc-winsvc-lpservice_main_functiona).
   * See [golang/sys](https://github.com/golang/sys/blob/master/windows/svc/example/service.go)
* Create inbound firewall rule using the following Powershell command: `New-NetFirewallRule -DisplayName "Storj v3" -Direction Inbound –Protocol TCP –LocalPort 28967 -Action allow`

This means that Windows handles starting and restarting the binary in the background.

## Implementation

1) Modify the storage node startup to implement ServiceMain so that the binary is considered a Windows Service.
2) Create script for registering binary as a Windows Background Service. (sc.exc command)
3) Create wix.json file for building an MSI.
4) Update the build process to include the creation of the MSI. 
5) Ensure the windows installer is working properly
  * Install Storage Node using msi.
  * Ensure binaries run on Windows startup, and it runs as a background process when not logged in.
  * Ensure UAC(s) is not triggered after installation.
  * Ensure that storage node restarts automatically after a crash 

## Open Issues/ Comments (if applicable)

* Consider writing an uninstaller.
* How do we prevent UAC from triggering?
* Consider using wix without go-msi.
