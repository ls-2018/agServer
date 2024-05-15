docker cp koord-control-plane:/etc/kubernetes/pki/front-proxy-ca.key front-proxy-ca.key
docker cp koord-control-plane:/etc/kubernetes/pki/front-proxy-ca.crt front-proxy-ca.crt


openssl genrsa -out apiserver.key 2048

openssl req -new -key apiserver.key -out apiserver.csr -subj "/CN=front-proxy-client"

openssl x509 -req  -days 3650    -in apiserver.csr -CA front-proxy-ca.crt -CAkey front-proxy-ca.key   -CAcreateserial  -out apiserver.crt

