# rote-notify


1 - [Descrição](#Descrição)
2 - [Instalação](#Instalação)

## Descrição


a. Gerar as chaves Publica e Privada
```bash
  openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
```

b. Extrair a chave publica
```bash
  openssl rsa -in private_key.pem -pubout -out public_key.pem
```