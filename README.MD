# is_odd 

> Returns true if the given number is odd

请支持一下作者，谢谢！

## Install

```sh
$ go get -u github.com/apple-han/is_odd
```

## Usage

```Go
import github.com/apple-han/is_odd;

t, _ := is_odd.IsOdd(1);
fmt.Println(t) //=> true

f, _ := is_odd.IsOdd(0);
fmt.Println(f); //=> false

g, _ := is_odd.IsOdd("-3");
fmt.Println(g); //=> true

_, err := is_odd.IsOdd("9999999999999999999999999999999999999");
fmt.Printf("err value:%s", err.Error()); //=> value out of range
```