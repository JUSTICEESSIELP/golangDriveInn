install the libray go env
create the env file

websockets in golang

a simple websocket from gorilla can manage just 1000 connections

1. Every socket is represented by a file descriptor
2. The OS needs memory to manage each open file
3. Memory is a limited resource
4. Maximum number of open files can be changed via ulimits

ulimits provides control over the resources available to processes

ulimit has both the soft limit and the hard limit

-soft limit is the actual limit per resource
-the hard limit is a ceiling to the soft limit
-the kernel enforces the soft limit for the corresponding resource

RLIMIT_NOFILE is the resource enforcing max number of open files \*\*

in Golang we change limits using the system call package.. the getter limit and setter limit

top - p ${pidof server}// by the way

pprof package would initialize a set of http handlers that expose some routes and extract server runtime profiling data in the format expect by the pprof visualization tool

we can use this to analyze the heap memory usage
and the goroutines analysis

so then websockets start this way we have the client sending an http request and now what is happening is that the client adds and updgrade connection in the header and then if the server supports it then the server would send a response then the bidirectional connection is made
the connnection can be closed by either the server or the client

go memory managment
https://pkg.go.dev/runtime / to know information about the runtime of the server

the garbage collector ;. the GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage .. The default is GOGC = 100 . Setting GOGC = off disbales the garbage collector entirely . The runtime/debug package's SetGCPercent function allows changing the percentage at runtime

append method reallocates memory in slices so lets say we write this

newArray := make([]string, 3)
newArray[0] = "hi"
newArray[1] = "there"
newArray[2] = "people"

but if I write
append(newArray , "Hen" , "chicken")
the size of the array would change and this would not lead to an array

LIFO when have multiple defer
defer statement invokes a function whose execution is deferred to the moment the surrounding function returns , either because the surrounding funcitno executed a return statement, reached the end of its function body , or because the corresponding gorouting is panicking

//example
func main(){
defer fmt.Println("hello")
defer fmt.Println("enya")
defer fmt.Println("the wild child")
fmt.Println("no defer")
deferFunc()

}

func deferFunc(){
defer fmt.Println("inside defer")

}

results//  
 no defer
inside defer
the wild child
enya
hello

WORKING WITH FILES IN GOLANG

errors : panic() would shut the down the execution of the program and show you

creation of file is an os operation so we use the [os] package
for the reading of the file we dont use that package .. when we write to the file it returns length when we read it returns bytes

when we get the bytes we can change it in many ways

1. string(bytes)

WEB REQUEST

1. its my responsibility to close the response

app := iris.New()

    location := app.Party("/api/v1/location")
    {
    	// booksAPI.Use(iris.Compression)

    	// GET: http://localhost:8080/books
    	location.Get("/autocomplete", routes.Autocomplete)
    	// POST: http://localhost:8080/books
    	location.Post("/search", routes.Search)
    }

    app.Listen(":8080")

---- when making a request {POST}

we have multiple formats

1. json
2. formData [url encoded] useful for things like images files

url.value this is a map where you store the url query in a map format

go mod tidy // use it when you dont see the indirect leaving the go.mod from a package

go list all // to list all the packages installed

go list -m all // what your project depends on directly

go mod verify // verify if you have downloaded the right package from the right source to prevent malicious packages

go list -m -versions github.com/gorilla/mux //inquire the number of versions a remote package has

go mod graph

go routines are compared with threads but they are different

1. Threads are managed by OS
   Fixed stack - 1MB

2. Managed by Go runtime
   Flexible stack - 2KB

so the go routines is managed by the go runtime and what is going to happen next is that .. so its true that os is responsible for alloting threads but here the go routines can spin up threads without the os permission

"Do not communicate by sharing memory; instead share memory by communicating

// adding the go routine to the first function is making another thread but leaving it just like this means we are not asking it to give us the response from that thread

1. some use sleep to make it wait but its not the best
2. the order is unpredictable
3. using the sync package: it provides basic synchronization primitives such as mutual exclusion locks . other than the [Once] and [WaitGroup] types most are intended for use by low level library routines .. High level synchronization is better done via channels and communication.

\*\*\*\*Values containing the types defined in the sync package should not be copied

func main(){
go greeter("hello")
greeter("world")
}

func greeter(s string){
for i:= 0 ; i< 6 ; i++{
fmt.Println(s)
}
}

WAIT GROUPS : LOW LEVEL SYNCHRONIZATION with sync package

1. we need a variable that is going to be the waitGroup .. we make a pointer

var wg sync.\*WaitGroup

2. now as soon as a go routine is created and we are using the sync package we need to create the variable that would be the WaitGroup type and then we add that go routine to the wait group ...
   so its saying that until the work is done the WaitGroup Add would wait and then

3. Done is used when the go routine is done

MUTEX

so now we have a situation where multiple goroutines want to modify a point in memory maybe a particular array of maybe even the same database at the same time .. this can be bad and can corrupt the data

so we lock that point in memory and make sure it is unlocked only and only after the goroutine that was writing to it is unlocked

read-write mutex is allowing to read the point in memory but for the writing the principle remains ... so multiple goroutines can share resource for reading purpose.. the reading is allowed only after the writing is done

1. create a variable that would be of mutex type
   we create a pointer type like we do for the WaitGroup

go run --race .

to check for racing conditions

mutex solves racing conditions which is caused by multiple go routines trying to write to a single source of truth or a point in memory

"Dont lock the resource directly for RWMutex but rather lock it when you are trying to read from it "

goroutines are independent of each other their information is not shared among other go routines .. channels can help with that so you have the reference of a go routine for another go routine to take and use ..

channels if you pass a value there should be a listener to that value

so if you pass two values into a channel there should be two listeners

RANDOM NUMBERS IN GOLANG

1. math randomness ..........math/rand package
2. cryptographic randomness ............crypto/rand package

rand.Seed(time.Now().UnixNano())
fmt.Println(rand.Intn(5)+1)

models in mongo db are by default bson so we include them when we are making our models schema espcially the

context package defines the context type , which carries deadlines , cancellation signals and other signals, and other request-scoped values across API boundaries and between processes

Incoming request to a server should create a Context, and outgoing calls to server should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout or WithValue.

So when making a request to an external server or database or service we need to use the context to set the rules to defined things like how long should we wait for the response and others

**\*\*\***5 situations where to use pointers \*\*\*\*

1. using a pointer receiver of a function : IF you to make a [new] thing and set that thing to a value and not its copy then use a pointer ..heard you can solve this with a return in the function and no need for the pointers

2. when you have json that have sub structs and you dont want memory allocated for them if they empty use a pointer

3. Creating a third state of primitive data types like boolean
4. Recursive Data Structure like example is double linked list

hupdp

/5000 _/ port
localhost hostName
driveinnadmin username
driveinnpostgres maintenance db
Alchemistry password
