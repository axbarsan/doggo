# doggo scripting language

## welcome

doggo is an interpreted scripting language that you can use, but you really shouldn't. It was written using Thorsten Ball's [book](https://interpreterbook.com), in order to learn more about programming languages.
It leverages Go's type system, garbage collector, and pretty much everything else. Did I mention that you shouldn't run this?

## what's in the box

Not much, really.

| **keyword** | **explanation (sort of)** |
|---|---|
| `const` | Declare a constant. What did you expect? |
| `fn` | Declare a function. Functions are first class citizens, they can be passed around and used pretty much everywhere. |
| `if`/`else` | Basic logic gate |
| `return` | End a function's execution |

#### data types

| **declaration** | **explanation (sort of)** |
|---|---|
| `const a = 1;` | Integer |
| `const b = "hello";` | String |
| `const c = true;` | Boolean |
| `const d = [1, 2, 3];` | Array |
| `const e = { "something": "some other thing" }` | Map |
| `const f = fn(x, y) { ... };` | Function |
 
#### syntax
 
* Each line must end with a semicolon `;`
* Only function expressions are supported
* Variables are function-scoped
* You can use expressions anywhere you can use a value
 
For more examples on how the code looks, try one of the examples from the [`examples`](examples) folder.

## now really, how do I run this?

First of all, you need Go installed (>= 1.11). You can download the latest version [here](https://golang.org/dl/).

Now, here come the commands:

```nohighlight
cd project/dir
go build
```

Now you'll see that a pretty executable file called `doggo` appeared.
You can run one of the examples, to check that everything is working:

```nohighlight
./doggo examples/simple.doggo
```

If you feel brave, you can also run the REPL:
```nohighlight
./doggo
```
Now you can start typing in code, line by line, and execute it on the spot.
