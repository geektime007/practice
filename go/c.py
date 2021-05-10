#!/usr/bin/env python
# _*_ Coding: UTF-8 _*_
from Crypto.PublicKey import RSA

if __name__ == "__main__":
    rsa = RSA.generate(1024)
    private_pem = str(rsa.exportKey(), encoding="utf-8")
    with open("private.pem", "w") as f:
        f.write(private_pem)
        f.close()
    public_pem = str(rsa.publickey().exportKey(), encoding="utf-8")
    with open("public.pem", "w") as f:
        f.write(public_pem)
        f.close()
