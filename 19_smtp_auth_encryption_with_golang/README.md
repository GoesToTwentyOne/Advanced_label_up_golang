# Non-Secure way
enable less secure app

openssl s_client -connect smtp.gmail.com:587 -quiet -starttls smtp
auth plain AG5paGFpbjIyQGdtYWlNvbQBBbGV4R29vdDIxMA==
MAIL FROM:<nihadhossain22@gmail.com>
RCPT TO:<nihadhossainthpibd373@gmail.com>
DATA

FROM:<nihadhossain21@gmail.com>
TO:<nihadhossainthpibd373@gmail.com>
SUBJECT:SMTP non secure way test

Hopefully you are alright!I'm going to your house on 25 December 2021.Have a nice party.Thanks
Regards
Md.Nihad Hossain

.

# Secure way
enable two step verification 
Encoding{
email:bmloYWRob3NzYWluMjJAZ21hag==
app password:b3djb2dsdnh3cm9o=
}
openssl s_client -connect smtp.gmail.com:587 -quiet -starttls smtp

MAIL FROM:<nihadhossain22@gmail.com>
RCPT TO:<nihadhossainthpibd373@gmail.com>
DATA

FROM:<nihadhossain21@gmail.com>
TO:<nihadhossainthpibd373@gmail.com>
SUBJECT:SMTP secure way test

Hopefully you are alright!I'm going to your house on 25 December 2021.Have a nice party.Thanks
Regards
Md.Nihad Hossain

.
