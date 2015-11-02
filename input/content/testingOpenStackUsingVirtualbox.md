title: Installing Devstack (Openstack) on a VirtualBox instance
date: 2015-05-10 23:22
categories:
- Cloud

---

This post shows all relevant steps to set up a **DevStack** instance in a VirtualBox instance. It is heavily based on : [http://docs.openstack.org/developer/devstack/guides/single-machine.html](http://docs.openstack.org/developer/devstack/guides/single-machine.html)

**VirtualBox 4.3.26 r98988** was used on a **Windows 8.1 64bit machine on a i5-4200u** machine, the operating system for the guest system is : **CentOS 7.1.1503 64bit Minimal (CentOS-7-x86_64-Minimal-1503-01.iso)**

### Create a new VirtualMachine :
-  Type: *Linux*
-  Version : *Red Hat(64bit)*
-  Memory Size : at least *4096*
-  Virtual Hard Drive : *VDI*, *Dynamically allocated*, *50 GB*

### Install CentOS
- VirtualBox will promt you to selct an image : select the *CentOS* image and continue startup
-  select the "Install" option at startup
- go to **network & host name** then click on the slider that reads **off** so it says **on** to use the network and then click **configure** to open a the configure window then go to the **Ethernet** tab and make sure a **device MAC address** is selected
- also configure the***installation destination** - simply selecting the harddrive to install to then start the installation
- during installation, set the **root password**
- **reboot** after the installation

### Installing Openstack
- login using the user **root** and the password entered during installation
- add a user **adduser stack**
- change its password **passwd stack**
- install git **yum install git -y**
- make sure the *stack* user can execute commands that need sudo rights without entering a password : **echo "stack ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers**
- now it is time to switch the user, **logout** and login as **stack** with the password you set
- get the *devstack* source **git clone https://git.openstack.org/openstack-dev/devstack**
- go to *devstack* folder **cd devstack**
- for the next step we first need some information about the network interface which we will get using **ifconfig** - but it is not available so we have to install it using **sudo zum install net-tools**
- now execute **ifconfig** you should see installed devices - search for the one which contains **(Ethernet)** and note the name (for me it was **enp0s3**), also look for the ip address of that device (for me **10.0.2.15**)
- now create a file called **local.conf** inside the **devstack** folder with the following content :

Example content of **local.conf**


    :::bash
    [[local|localrc]]
    FLOATING_RANGE=192.168.1.224/27
    FIXED_RANGE=10.11.12.0/24
    FIXED_NETWORK_SIZE=256
    FLAT_INTERFACE=enp0s3
    ADMIN_PASSWORD=supersecret
    MYSQL_PASSWORD=iheartdatabases
    RABBIT_PASSWORD=flopsymopsy
    SERVICE_PASSWORD=iheartksl
    SERVICE_TOKEN=xyzpdqlazydog


- make sure that the ip address of the device (we found this when we executed **ifconfig**) doesn't fall into the ip range of **FLOATING_RANGE** or **FIXED_RANGE** (this should not be the case)
- change the **FLAT_INTERFACE** to the name of the device we found when we executed **ifconfig** (for me it was **enp0s3**)

### Installing GNOME
- as user **root** (or with sudo as user **stack**)
- install GNOME **yum groupinstall "GNOME Desktop"**
- then restart the instance
- then the setup will ask some questions : we don't need to create a new user, hence enter **2** (License information), then **2** to accept it, then **c** to return to the menu and finally again **c** to finish the setup
- then login as **root** and execute **ln -sf /lib/systemd/system/runlevel5.target /etc/systemd/system/default.target** - after restarting the GUI will now be used


### Running DevStack

- start the installation using **./stack.sh**
- it is possible that you will get a message that the **Existing lock /var/run/yum.pid: another copy is running as pid 3275. Another app is currently holding the yum lock....** in this case stop or wait until the update process started by the GUI has stopped to check for updates (may take a few minutes)
- after the **./stack.sh** process has finished open **firefox** and enter **localhost** now you should be able to see the **OpenStack Dashboard** and be able to login using the user **admin** and the **ADMIN_PASSWORD** specified in the **local.conf**