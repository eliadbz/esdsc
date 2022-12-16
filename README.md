# Enum Subdomain by Status Code

This tool enumerates subdomains using status codes return by the HTTP status codes. This makes it useful for enumerating web servers that use engines like NGINX, whose subdomains can be enumerated by the status code returned.


Requirements
---
- golang 1.19+
- make



Installation
--
`sudo make install`


Usage
--
To get all the parameters available, use: `esdsc -h`

example: `esdsc -w path/to/wordlist -ssl -s 301 -d example.com`