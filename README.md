Most of the time, we don’t have access to the GUI of Google Drive, and we need to upload data to our Drive, such as when using Virtual Machines on the Cloud. Godrive is a tool that allows users to directly upload directories/files to Google Drive from any machine.

## 1. Installation
To install this tool, simply run:
```sh
pip install godrive
```

## 1.1 Available Commands
Once installed, you can use the following commands:
```sh
godrive version        #1.0.xx
godrive info           #Dev info
godrive                #To Upload Data
```

## 1.3 Getting Started
After running the command:
```sh
godrive
```
### You need to paste the content of `client_secrets.json`
If you don’t have a JSON file, refer to Paragraph 2 at the end.

### Now, you have to enter the absolute/full path of file or directory :
for example : /home/rohan/directory_one/file_one

You can also upload a full directory but, to save time uploading files well be better than directory, because we will first create zip of your directory.

### To upload multiple files, at once :
1. If all the files are in same Directory
   `<path_to_file1>$<name_of_file2>$<name_of_file3>.........`

2. If the files are in different Directories
   `<path_to_file1>$<path_to_file2>$<path_to_file3>.........`
   
To avoid conflicts, leave the custom folder name empty.

### Authentication Requirement
You will be asked to authenticate with the account where you want to upload data.
Make sure to authenticate only with the accounts that are added to the Consent Screen.

---------------------------------------------

## 2. How to Get client_secrets.json?
In Browser, go to :
```sh
console.developers.google.com
```

1. Agree to the Terms of Service.
2. Create a project (if not already created).
3. After creating the project, click on `+ Enable API and Services.`
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-1.png)

4. Now, in the Search bar search for `Google Drive Api`
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-2.png)

5. Enable the API which is shown on the top.
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-3.png)

6. Now, click on the `Create Credentials` button, if not showing this then click on the `credentials tab`.
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-4.png)

7. Click on the `+ Create Credentials` button just below the credentials tab.
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-5.png)

8. Click on `OAuth Client ID`.
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-6.png)

9. Now, Click on `CONFIGURE CONSENT SCREEN` and then click on `Get Started`.
![Cloud Panel](https://github.com/rohanbhatotiya/godrive/blob/main/assets/img/cloud-panel-7.png)


## Configure OAuth Consent Screen
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
    
## Before using the JSON file in Godrive, you have to do some more things.....
1. In the leftmost panel click on `Audience`
2. Scroll down, just below the Test Users click on `+ Add User`
3. You can add any Gmail account, but you must authenticate with that account. The uploaded files will be stored in that Gmail account's Google Drive.

Now, you are ready use the JSON file content in GODRIVE.
