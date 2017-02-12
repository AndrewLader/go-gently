# go-gently
The Go-Gently service is a Go language package to enable other Go services to gently shutdown when they receive a `SIGTERM`, `SIGINT` or `SIGQUIT` signal. 

```
 ██████╗  ██████╗        ██████╗ ███████╗███╗   ██╗████████╗██╗  ██╗   ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██╔════╝████╗  ██║╚══██╔══╝██║  ╚██╗ ██╔╝
██║  ███╗██║   ██║█████╗██║  ███╗█████╗  ██╔██╗ ██║   ██║   ██║   ╚████╔╝ 
██║   ██║██║   ██║╚════╝██║   ██║██╔══╝  ██║╚██╗██║   ██║   ██║    ╚██╔╝  
╚██████╔╝╚██████╔╝      ╚██████╔╝███████╗██║ ╚████║   ██║   ███████╗██║   
 ╚═════╝  ╚═════╝        ╚═════╝ ╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝   
                                                                          

```

### Usage

There are two simple steps necessary to use go-gently in a go service or application.

#### Step One
The first step is to implement the `Gently` interface on each of the `struct`s that should be notified to stop gently. The `Gently` interface definition is:

```
// Gently is the interface a struct must implement if it wants to be registered
// to notified as to when to stop gently
type Gently interface {
	GetName() string
	StopGently(sginal os.Signal)
}
```

#### Step Two
In the `main` func, make the following calls:

**Sample Code**
```
    import "github.com/AndrewLader/go-gently"

    myStuct := NewMyStruct()

    goodNight := gently.New()
    goodNight.Register(myStruct)
```

### GoodNight struct
The `GoodNight` struct is used to manage the structs in a Go service that implement the `Gently` interface. The `Register` method is used to register a given struct that implements the `Gently` interface.

```
// Register is used to register a struct that implements the Gently interface
// with the GoodNight struct so it can be notified when to stop gently
func (goodNight *GoodNight) Register(toBeRegistered Gently)
```
