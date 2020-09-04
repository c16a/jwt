# jwt
This is a tool to encode and decode JWT,
but mostly to satiate my undying love from the terminal.

Thanks for spf13's cobra - without which this tool would have 
been a lot more difficult to create.

**Warning**: Obviously don't use this in production, yet. 
The encoding and decoding uses [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go).
Embed Dave's library in your code, like many other wiser folks.

## Usage
### Decode

```shell script
jwt decode --token eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE2MTYyMzkwMjJ9.7eZhwx58ujtOk7KmOIypMqDBRMqOPvb-OknZ1PIjt_q2_KXAUU6bmfpyuu2b7uuf
```

With hmac secret
```shell script
jwt decode --token eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE2MTYyMzkwMjJ9.7eZhwx58ujtOk7KmOIypMqDBRMqOPvb-OknZ1PIjt_q2_KXAUU6bmfpyuu2b7uuf --secret my_super_secret
```

#### Supported algorithms
- [x] HMAC
- [ ] RSA

### Encode
Currently in development.