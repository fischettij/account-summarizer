# Accounr Sumarizer

This system processes files containing a list of debit and credit transactions on an account. It then sends summary information to a user.

### Clone

```bash
git clone git@github.com:fischettij/account-sumarizer.git
cd account-sumarizer
```

#### Run with docker

```bash
sudo docker-compose up
```

### Interface
The application will enable the following endpoints.

TBD

### Email server
The microservice can integrate to a smtp server. For the development docker-compose define a dummy smt server rnwood/smtp4dev
Can be acceded from localhost:5000

### Notes
The endpoint POST /summary/send could be grpc or asynchronous. But to keep this simple te endpoint work sync 