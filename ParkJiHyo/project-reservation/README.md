## Reservation service
### 카쉐어링 서비스 이용자들이 차량 예약 서비스를 사용할 수 있다.

<br/>


## 예약 가능 차량 조회
+ ### Filter
    + ### 시간
    + ### 요일
    + ### 차량 위치
        + ### 예약할려는 사람의 위치로부터 가장 가까운 예약 가능차량과 거리반환(mongoDB geoJson활용)
    
+ ### API List
    + ### 예약 차량 조회  
    + ### ListReservationVehicle
        + ### 필터에 따라 가져올 수 있는 예약 가능 차량 리스트
            + ### request
            + ### method: GET
            + ### beginAt,endAt
            + ### latitude,longitude
            + ### pageInfo(limit,offset)
            
            <br/>
          
            + ### response
            + ### Vehicle 정보
    + ### GetReservationVehicle
        + ### 예약 가능 목록중 1개를 선택 했을때 세부 정보
            + ### request
            + ### method: GET
            + ### reservationID
            
            <br/>
          
            + ### response
            + ### User(차량 소유자 정보)
            + ### Vehicle(차량 세부 정보)
            + ### Review(해당 차량 리뷰 정보)
    
    <br/>

<br/>

### 예약 가능 차량 삭제
`
+ #### 예약 상태일 때 삭제불가
`
### 예약 가능 차량 업데이트
`
+ #### 예약 상태일 때 업데이트 불가
`

## schema
+ ### Reservation
    + ### reservationID - uuid
    + ### userID - uuid
    + ### userName - string
    + ### address - string
    + ### position(위도,경도) - (mongodb의 위치타입)
    + ### reserved - boolean