### 차량등록서비스

- 회원이 등록하는 것
	- 회원 로그인 -> 차량 등록
- 회원-차량 mapping table
	- 회원 access token 획득과 api call시 access token을 header에 넣고 api call 해야 하는 과정 필요
- 차량 table(차 주인 회원, 차 정보(차종, 연식 등등), isReserved, …)
	- 차 정보는 차 최종 등록 시 입력
	- isReserved는 bool 값 -> true: 못빌림, false: 빌릴 수 있음
		○ True일 때 reserved된  startAt(timestamp), endAt(timestamp) 값이 있음
- <b>스키마</b>
	- Registration
		○ registrationID - uuid
		○ ownerID - uuid (회원관리서비스와 협의 필요해보임..)
		○ vehicleID - uuid
	- Vehicle
		○ vehicleID - uuid
		○ ownerID - uuid
		○ model - string (차종)
		○ type - string (hev, ev, engine)
		○ year - int (몇년도 양산)
		○ capacity - int (cc)
		○ isReserved - boolean
		○ reservingUser - uuid 
		○ startAt - timestamp
		○ endAt - timestamp
