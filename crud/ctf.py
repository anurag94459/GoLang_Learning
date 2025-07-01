from base64 import b64decode
from Crypto.Cipher import Blowfish

def unpad(s):
    return s[:-s[-1]]

key = b"l33t_k3y"
cipher = Blowfish.new(key, Blowfish.MODE_ECB)
flag_ct = b64decode("QS3kRaEFSL8/DgxTHUMDReK2GPiUTMDe1X7oXAtjYWqVgTfJNfPXbKr9QVk8zMUc4koyiRvnZyR/c5W/l4LIEg==")
flag_plain = cipher.decrypt(flag_ct)
print(unpad(flag_plain).decode())
