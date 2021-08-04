### 차량등록서비스

- 회원이 등록하는 것
    - 회원 로그인 -> 차량 등록 (차번호로 차종조회 고려,, 유료임)
- 회원-차량 mapping table
    - 회원 access token 획득과 api call시 access token을 header에 넣고 api call 해야 하는 과정 필요
- 차량 table(차 주인 회원, 차 정보(차종, 연식 등등), isReserved, …)
    - 차 정보는 차 최종 등록 시 입력
    - isReserved는 bool 값 -> true: 못빌림, false: 빌릴 수 있음
        - True일 때 reserved된  startAt(timestamp), endAt(timestamp) 값이 있음
- <b>스키마</b>
    - map_service_vin
      - vehicleID
      - serviceID
      - VIN
    - Registration: map_registration
        - registrationID - uuid
        - masterUserID - uuid (회원관리서비스와 협의 필요해보임..)
        - vehicleID - uuid
    - Vehicle - info_vehicle
        - vehicleID - uuid
        - masterUserID - uuid
        - model - string (차종)
        - type - string (hev, ev, engine)
        - year - int (몇년도 양산)
        - capacity - int (cc)
        - isReserved - boolean
        - reservingUser - uuid 
        - startAt - timestamp
        - endAt - timestamp
        
### 등록 프로세스
1. 차량 등록
    - 등록 request: POST
      1. masterUserID: uuid
      2. serviceID: String
      3. VIN: String
      4. 차량정보(schema에 다른 json)
      5. create_time: timestamp(UNIX Epoch seconds 참고,,)

### DB
```sql
+--------------------------------+
| Tables_in_vehicle_registration |
+--------------------------------+
| info_vehicle                   |
| map_registration_vehicle       |
+--------------------------------+
```
- 테이블 생성 완료
- 테이블 스키마에 맞는 API 호출 코드 구현