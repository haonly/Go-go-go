# DataPreprocessing Problem
* 출제자 : Lee Seung Yoon
* 출제 의도 : Go Channel, Go Routine 활용

---
### 1. 문제 설명
* 데이터 : [kaggle 사이트](https://www.kaggle.com/adityadesai13/used-car-dataset-ford-and-mercedes) 의 브랜드별 car dataset
  * dataset 구성 : 아래 11개의 csv
    * 2개는 unclean 데이터라 제외
    * 미리 data dir에 받아놓음. 데이터는 받지 말고 그냥 사용하세요.
  * dataset 브랜드 (11개) : "audi", "bmw", "cclass", "focus", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"
  * column
    * MODEL : car model
    * YEAR : Registration year
    * PRICE
    * TRANSMISSION (type of gearbox)
      - `Manual`, `Semi-Auto`, `Other`
    * MILEAGE : distance used
    * FULE TYPE
      - `Diesel`, `Petrol`, `Other`
    * TAX : road tax
    * MPG : miles per gallon
    * ENGINE SIZE : size in litres (리터 단위의 사이즈)

1. 각 브랜드별 price, mpg 평균을 구하세요.
   * (e.g.) audi → price AVG : 10000, mpg AVG : 55.4
   * (e.g.) bmw → price AVG : 12000, mpg AVG : 60.0

2. 전체 브랜드에 대해서, transmission이 Manual, Automatic 인 차량 각각의 mileage, mpg의 평균을 구하세요.
   * (e.g.) manual의 mileage AVG : 15000, mpg AVG : 55.0

3. 전체 브랜드의 [transmission, fuel type] 조합의 tax 평균을 구하세요.
   * (e.g.) [manual, petrol] tax AVG : 61
   * (e.g.) [manual, diesel] tax AVG: 48
   * (e.g.) [automatic, petrol] tax AVG : 58
   * (e.g.) [automatic, diesel] tax AVG : 78

(__3번 문제는 조금 어려울 수 있다고 생각합니다. 1, 2번 문제는 필수로 풀고, 3번은 자유롭게 풀어오면 되겠습니다.__)

---
### 2. 문제 평가
* 아직 문제의 정답은 모릅니다.. (저도 풀어봐야 함)
* 정답이 맞다는 가정하에, __벤치마킹 챌린지__ 를 하겠습니다.

```go
cd DataPreprocessing/preprocess
go test -bench . -benchtime=100x -benchmem // 시간, 메모리 사용량, 작업당 할당량 확인, 100번 실행
```
* solution1,2,3 함수에 대한 benchmark 함수를 만들어놓음.
* 해당 dir에 들어가서 명령어를 실행시키면, 시간, 메모리 사용량, 할당량을 볼 수 있음.
* 기존 benchmark 실행횟수가 너무 커서, 100번만 실행시킴.

__[문제 보상] 가장 빠르게 실행한 분에게 스벅 쿠폰 드립니다.__    
_(물론 컴퓨터마다 성능 차이가 있겠지만...)_

---
### 3. 기타
* 임의로 `Preprocess`, `readFromCSV` 함수를 만들어놓았습니다. 하지만 go 루틴이나 채널을 고려하지 않은 예시 함수이니 원하는대로 변경하셔서 쓰시면 됩니다.
* 정답이 기입되어야 하는 함수는 `solution1,2,3` 입니다.
* 질문 있으시면 연락주세요.
