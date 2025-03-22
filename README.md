Most of the time, we don’t have access to the GUI of Google Drive, and we need to upload data to our Drive, such as when using Virtual Machines on the Cloud. Godrive is a tool that allows users to directly upload directories/files to Google Drive from any machine.

Installation
To install this tool, simply run:
```sh
pip install godrive
```

Available Commands
Once installed, you can use the following commands:
```sh
godrive version        #1.0.xx
godrive info           #Dev info
godrive                #To Upload Data
```

Getting Started
After running the command:
```sh
godrive
```
You need to paste the content of client_secrets.json.

How to Get client_secrets.json?

In Browser, go to :
```sh
console.developers.google.com
```

1. Agree to the Terms of Service.
2. Create a project (if not already created).
3. After creating the project, click on `+ Enable API and Services.`
![Cloud Panel](https://yourimageurl.com/image.png)


Configure OAuth Consent Screen
1. Now enter any thing in App Name input
2. Choose your email address in `User Support Email`
3. Next, then choose External.
4. Next, Again enter your email address.
5. Next, click on the CheckBox.
6. Next, then click on `Create` button.
7. Now, click on `CREATE OAUTH CLIENT`
8. In Application Type choose `Desktop App`
9. You can give any name.
10. Click on `Create`
11. Now, download the JSON file.
    
Before using the JSON file in Godrive, you have to do some more things.....
1. In the leftmost panel click on `Audience`
2. Scroll down, just below the Test Users click on `+ Add User`
3. Here you have to add the gmail address on which you want to upload files. You can add anyone's gmail here but you have to Authenticate with that gmail, and the file will also be uploaded to that Gmail's Drive.
