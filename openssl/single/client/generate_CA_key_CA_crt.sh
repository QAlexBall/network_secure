openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=wintrysec_com" -days 5000 -out ca.crt