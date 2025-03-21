Most of time we have no access to GUI of Google Drive and we need to upload data to our drive like when using Virtual Machines on Cloud.
Godrive is a tools which allows user to directly upload directories/files to Google Drive from any machine.

To install this tool you just need to run :
```sh
pip install godrive
```
Now you can use three commands :
```sh
godrive version        #1.0.xx
godrive info           #Dev info
godrive                #To Upload Data
```
Now, after running the command
```sh
godrive
```
You need to paste client_secrets.json content
Where to get this!
In Browser, go to :
```sh
console.developers.google.com
```
Agree the terms and service,
Create a project if not already created
After creating project click on `+ Enable API and Services`
