# 서버

## 목차
1. [서버 작동 방법](#서버-작동-방법)
2. [API 명세서](#API명세서)

### 서버 작동 방법

1. 쉘에서 다음과 같은 쉘코드들을 모드 입력해준다.
```shell
go get -u github.com/Goolgae/GSM_data

go get -u google.golang.org/api/sheets/v4
go get -u golang.org/x/oauth2/google

go get github.com/labstack/echo

go get github.com/dgrijalva/jwt-go
golang.org/x/time/rate
```

2. secret.json을 루트 디렉터리에 넣어준다.
  ```shell
  GSM_data/secret.json
  ```

3. 다음과 같이 코드를 실행해준다.
```shell
GSM_data> go run server/main.go
```

> 서버실행 완료

### API명세서
1. 데이터를 백분율로 가져오기(api/v1/data) 
  - GSM 4기 트렌드 취향 (1학년 말) 데이터 가져오기
    - Request
    > GET /trand_taste
    
    - Response
    ```json
    {
      "major": {
        "AI": 0.07692307692307693,
        "Android": 0.07692307692307693,
        "BE": 0.15384615384615385,
        "Etc": 0.11538461538461539,
        "FE": 0.23076923076923078,
        "Game": 0.15384615384615385,
        "IOS": 0.07692307692307693,
        "IOT": 0.038461538461538464,
        "Security": 0.07692307692307693
      },
      "reason": {
        "Cool": 0.07692307692307693,
        "Easy": 0.038461538461538464,
        "Etc": 0.6923076923076923,
        "Money": 0.07692307692307693,
        "Recommendation": 0,
        "Want": 0.11538461538461539
      },
      "studyMethod": {
        "Book": 0,
        "Etc": 0.8846153846153846,
        "Google": 0,
        "Lecture": 0,
        "OfficialDocument": 0.038461538461538464,
        "QuestionToFriends": 0,
        "QuestionToSenior": 0.07692307692307693,
        "QuestionToTeacher": 0
      },
      "studyTime": {
        "Four": 0.15384615384615385,
        "One": 0.34615384615384615,
        "Six": 0.038461538461538464,
        "Three": 0.2692307692307692,
        "Two": 0.19230769230769232
      },
      "likeLanguage": {
        "C#": 0.038461538461538464,
        "C/Cpp": 0.15384615384615385,
        "Etc": 0.23076923076923078,
        "Go": 0.038461538461538464,
        "Java": 0.038461538461538464,
        "Js": 0.19230769230769232,
        "Kotlin": 0.11538461538461539,
        "Python": 0.11538461538461539,
        "Swift": 0.038461538461538464,
        "Ts": 0.038461538461538464
      }
    }
    ```
      > `major`은 각자의 전공 수요에 대해서 나타낸 값이다.
      
      > `reason`은 각 전공을 선택한 이유에 대한 값이다.
      >   - `Cool`은 멋져보여서 를 뜻함
      >   - `Easy`는 쉬워보여서 를 뜻함
      >   - `Money`는 돈이 되기때문에 를 뜻함
      >   - `Want`는 그것에 원하는게 있기 때문에 를 뜻함
      >   - `Recommandation`은 추천을 받았기 때문에 를 뜻함
      >   - `Etc`는 기타 등등을 뜻함

      > `studyMethod`는 공부 방법을 나타낸 값이다.
      
      > `studyTime`은 공부한 시간을 나타낸 값이다.
      
      > `likeLanguage`는 자신이 좋아하는 언어를 뜻한다.
    
