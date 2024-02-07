# Accounr Sumarizer

This system processes files containing a list of debit and credit transactions on an account. It then sends summary information to a user.

### Clone

```bash
git clone git@github.com:fischettij/account-sumarizer.git
cd account-sumarizer
```

#### Run with docker
Docker compose has default configurations to run
```bash
sudo docker-compose up
```

But if you want to configure somethin there are the env vars at summarizer container
For Application: 
* `PORT`: Port where the http server is listening
* `FILES_DIRECTORY`: Path where the files will be searched

For SMTP client:
* `SMTP_PORT`
* `SMTP_SERVER_URL`
* `SMTP_FROM`
* `SMTP_USERNAME`
* `SMTP_PASSWORD`
* `SMTP_IDENTITY`
* `SMTP_TLS_HOSTNAME`

Volumes:
The application use one volume to mount the files folder
```
volumes:
- ./examples:/summaries
```

### Interface
The application will enable the following endpoints.

**POST /summary/send**

Send email to user with the information in file with file_name
**Body**
The request body must be a JSON object with the following fields

|          Name | Required |  Type   | Description                          |
| -------------:|:--------:|:-------:|--------------------------------------|
|     `email` | required | string  | email to send the summary            |
|     `file_name` | required | string  | file name where the transactions are |

  **Response:**
    - HTTP Status: 204
    - Body: Witout body

```
curl --location 'localhost:8080/summary/send' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "leomessi@afa.com",
"file_name": "04.csv"
}'
```

### Email server
The microservice can integrate to a smtp server. For the development docker-compose define a dummy smt server rnwood/smtp4dev
Can be acceded on localhost:5000

### Notes
- It wouldn't be necessary for the application to have endpoints if it could obtain more information based on the file name.
For example, if the file name were the account number, the application could use that information to look up the email. In that case, the microservice could operate without endpoints by simply monitoring the folder.
- The endpoint POST /summary/send could be grpc or asynchronous. But to keep this simple te endpoint work sync