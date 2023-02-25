# OAuth 2.0

Details
[RFC6749](https://www.rfc-editor.org/rfc/rfc6749)




## Authorization Code Grant

1. Authentication endpoint

   1. Request to Authentication endpoint

   ```http
   GET {Authentication endpoint}
     ?response_type=code            // necessary
     &client_id={client_id}      // necessary
     &redirect_uri={redirect_uri}  // necessary depending on conditions
     &scope={scope}              // option
     &state={any string}              // recommended
     &code_challenge={challenge}     // option
     &code_challege_method={method} // option
     HTTP/1.1
   HOST: {Authentication Server}
   ```

    2. Response from Authentication endpoint

    ```http
    HTTP/1.1 302 Found
    Location: {redirect_uri}
        ?code={authenticated code}        // necessary
        &state={any string}               // necessary If request has a "state"
    ```

    3. Request to token endpoint

   ```http
   POST {token endpoint} HTTP/1.1
   Host: {Authentication server}
   Content-Type: application/x-www-form-urlencoded

   grant_type=authorization_code   // necessary
   &code={authenticated code}      // necessary to specify the value which is in the response of Authentication endpoint
   &redirect_uri={redirect_uri}   // necessary If Authentication request has a redirect_uri
   &code_verifier={verifier}     // necessary If Authentication request has a code_challenge
   ```

## Implicit Grant
## Resource Owner Password Credentials Grant
## Client Credentials Grant
