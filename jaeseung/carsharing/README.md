# (작업중)Review service for car sharing

* Review 서비스는 차량, 대여자, Host에 대한 리뷰를 관리하는 기능을 제공한다.

## 서비스별 작업, 협동자

| 서비스 | 작업 | 협동자 |
|--------|------|--------|
| 리뷰 서비스| createReview|회원 서비스:createRating| 
| | findRating| 차량 서비스:createRating |
| | findReviews| |
| | ... | |
| 회원 서비스| createRating| |
| | ... | |
| 차량 서비스| createRating| |
| | ... | |

## User story

### Review service 영역으로 판단되는 기능들

* 사용자(Host, 대여자, Bot 등)는 리뷰를 생성할 수 있다.
* 리뷰 작성자는 본인이 작성한 리뷰 수정/삭제 할 수 있다.
* 누구나 차량에 대한 리뷰를 조회 할 수 있다.
* 사용자에 대한 리뷰를 조회 할 수 있다.
* 사용자/차량에 남겨진 Rating의 평균점수를 조회 할 수 있다.

### 아마도 타 서비스(회원서비스, 차량서비스 조회)가 제공해야 할 기능

* Top rating host들을 조회 할 수 있다.
* 지역별 Top rating 차량들을 조회 할 수 있다.
* 지역별 Top rating host들을 조회 할 수 있다.
* 차량 브랜드별 Top rating 차량들을 조회 할 수 있다.

## 제공 API

* {BaseURL}: /reviews

### Common Headers

* TBD

### Create reviews

* HTTP method: POST {BaseURL}/reviews

#### Create request body

|필드명|필수|Type|Limit|내용|
|------|----|-----|----|----|
|review_target_id|M|string| |Review작성자 ID|
|text| |string| | |Review 내용|
|rating| |Number|0~5.0| |

#### Create response body

### Get reviews

#### Get response body

|필드명|필수|Type|Limit|내용|
|------|----|-----|----|----|
|id| |string| |Review ID|
|creator_id|M|string| |Review작성자 ID|
|review_target_id|M|string| |Review작성자 ID|
|create_epoch| |integer|UNIX Epoch seconds||
|text| |string| | |Review 내용|
|rating| |Number|0~5.0| |