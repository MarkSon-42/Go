# Go

- there is no class.

- but having structure with methods.

- 어떤 타입이든 메서드를 가질 수 있음.

- 상속 없음, 인터페이스는 있음.

- 함수 리터럴(람다 함수) 있음.

- 고성능 가비지 컬렉터

- 1.18부터 제네릭 공식 지원

~~~go
func Swap[T any](a, b *T) {  // T (all type allowed.)
    temp := *a
    *a = *sb
    *b = temp
}
~~~