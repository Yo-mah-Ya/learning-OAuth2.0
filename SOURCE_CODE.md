# gen

Generate `openapi/server.go` and `openapi/types.go`

```sh
make gen
```

# goauth-server

### easy trying

1. first of all, open the URL
   https://oauthdebugger.com/

2. fill out the forms like shown below

|                          |                                 |
| ------------------------ | ------------------------------- |
| Authorize URI (required) | http://localhost:3000/auth      |
| Redirect URI (required)  | https://oauthdebugger.com/debug |
| Client ID                | 1234                            |
| Scope (required)         | read                            |
| State                    | abc                             |
| Nonce                    |                                 |
| Response type            | code                            |

2. login page

|          |     |
| -------- | --- |
| user id  | u   |
| password | p   |

3. click
   submit

4. Success
   check it out the Success UI is seen.
