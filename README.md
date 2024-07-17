# Go Mail Sender

The `Go_mail_sender` application is an gmail sender that asks users to provide an HTML document to be sent to a recipient. It connects to Gmail's SMTP server and sends emails from the email you provide. In addition, log files are saved to AWS S3 when shutting down for permanent storage.

# IMPORTANT

1: `Go_mail_sender` is built for Windows development and Linux testing. If you have a Mac, you will need to change the `-v` parameter passed to `docker run`.

2: This application requires AWS permissions, which you can find in your `C:/Users/{USERNAME}/.aws`

# Usage

## Locally

To run this locally on a Docker container, run the following commands:

```
cd Go_mail_sender/deploy/

docker build -t go-mail-sender .

docker run -v C:/Users/{CHANGE1}/.aws:/root/.aws -e EMAIL="{CHANGE2}" -e PASSWORD="{CHANGE3}" -p 8080:8080 go_mail_sender
```

- CHANGE1
  - This is the Windows user that is currently logged in. Do not use Public.
- CHANGE2
  - This is the gmail account you would like to send emails from
- CHANGE3
  - This is the app password. You can find more information about it here: https://support.google.com/accounts/answer/185833?hl=en
 
## Elastic Beanstalk

To run this on EBS, you need to save this image to Docker Hub or AWS Elastic Container Registry with the name go-mail-sender. In the dockerrun.aws.json file, you will need to change:

- USER-ID
  - This is your AWS user id.
