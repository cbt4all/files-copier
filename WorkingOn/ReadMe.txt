Apply the same command on the list of the IPs given


Releases-v0.001
- Credentials are in a text file - ./credentials.txt
- IPs are in a text file - ./IPs.txt
- The config is in ./template-config/config.txt
- The commands output is in ./outputs/IP.txt 
- All logs are stored in ./logs/
- Program asks if devices needs different password (can be 2FA or differet pass)
	If Enter is pressed with no password, what is saved in ./credentials.txt is used.
- When program asks if devices needs different password, if yes, it also ask for PIN so save PIN
- Save not SSHed IP to the file: ./outputs/notSSHedIP.txt
- Added "diffie-hellman-group1-sha1" to config := &ssh.ClientConfig
- Gets IPs from given CSV/hosts files: 
	Read all files in ./csv-files path and asks user to select which csv file to use. 
	Then read the file ./hosts.txt and start comparing. When it is searching the CSV file, if it finds any IP asociated to the hostnames in the file ./hosts.txt saves it in ./IPs.txt and if not saves in ./outputs/NotFoundHosts.txt


Releases-v0.002
- Use flag to decide about same password, same config and using csv/IP file (User must use all flags)
	Use the file IPs.txt, with same password for all devices and apply the same config: -ipfrom=iplist -samepass=y -sameconfig=y
	Use the file IPs.txt, with same password for all devices but apply the different config: -ipfrom=iplist -samepass=y -sameconfig=n
- Check if the number of character of each differs with the number of characters that sent to the device, makes SSH unsuccessful and write a log


Releases-v0.003
- Integerate with Telnet

-ipfrom=iplist -samepass=y -sameconfig=y -tftp=n -ftp=n