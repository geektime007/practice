import hmac
msg = b'fanjiaolong'
key = b'password'
h = hmac.new(key, msg, digestmod='MD5')
print h
h.hexdigest()
