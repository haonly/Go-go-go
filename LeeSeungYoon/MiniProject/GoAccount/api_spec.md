# API Spec
## [GET] Get User Profile

### URL
```bash
curl 
--location 
--request GET 'http://rushbear.me/api/v1/user/profile?user_id=0969c938-c45c-43fa-9d60-0a1f023ddb53' \

--header 'Authorization: Bearer     "access_token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkxlZSBTZXVuZyBZb29uIiwiaWF0IjoxNTE2MjM5MDIyfQ.v2kXuRhfbIMPmjodr9fXefVAJgayIWd7xyzWzLkScdo'
```

```bash
http://rushbear.me/api/v1/user/profile?user_id=0969c938-c45c-43fa-9d60-0a1f023ddb53
```

### Method
 - GET

### Authorization
 - Bearer Token

### Request Params
|fields|required|sample
|:---:|:---:|:---:|
|user_id|O|0969c938-c45c-43fa-9d60-0a1f023ddb53


****

### Response Example
```json
{
    'email' : "example@goaccount.com"
    'name' : "홍길동"
}
```

### Response Code
|Status Code|Status|Description|
|:---:|:---:|:---:|
|200|OK|성공시|
|400|Bad Request|스펙외 요청의 경우|
|401|Unauthorized|Bearer Token 인증 실패한 경우|
|404|Not Found|해당 user_id의 사용자 정보를 찾지 못한 경우|
|405|Method Not Allowd|GET 외의 Http Method 사용한 경우|
|500|Internal Server Error|리소스 서버 에러난 경우|
|503|Service Unavailable|아직 해당 요청을 처리할 수 없는 경우|

****

## [POST] Authorize Token
 - 참고 URL
   - https://developers.google.com/identity/sign-in/android/backend-auth#calling-the-tokeninfo-endpoint

### URL
```bash
curl 
--location 
--request GET 'http://rushbear.me/api/v1/auth/tokeninfo?access_token=your_access_token'
```

```bash
http://rushbear.me/api/v1/auth/tokeninfo?access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkxlZSBTZXVuZyBZb29uIiwiaWF0IjoxNTE2MjM5MDIyfQ.v2kXuRhfbIMPmjodr9fXefVAJgayIWd7xyzWzLkScdo
```

### Method
 - GET

### Authorization
 - No Auth

### Request Params
|fields|required|sample
|:---:|:---:|:---:|
|access_token|O|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkxlZSBTZXVuZyBZb29uIiwiaWF0IjoxNTE2MjM5MDIyfQ.v2kXuRhfbIMPmjodr9fXefVAJgayIWd7xyzWzLkScdo|

****

### Response Example
```json
{
    "iss": "https://accounts.google.com",
    "sub": "110169484474386276334",
    "azp": "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com",
    "aud": "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com",
    "iat": "1433978353",
    "exp": "1433981953",

    // These seven fields are only included when the user has granted the "profile" and
    // "email" OAuth scopes to the application.
    "email": "testuser@gmail.com",
    "email_verified": "true",
    "name" : "Test User",
}
```

### Response Code
|Status Code|Status|Description|
|:---:|:---:|:---:|
|200|OK|성공시|
|400|Bad Request|스펙외 요청의 경우|
|405|Method Not Allowd|GET 외의 Http Method 사용한 경우|
|500|Internal Server Error|리소스 서버 에러난 경우|
|503|Service Unavailable|아직 해당 요청을 처리할 수 없는 경우|


```json
// 400 Error Example
{
  "error": "invalid_token",
  "error_description": "Invalid Value"
}
```