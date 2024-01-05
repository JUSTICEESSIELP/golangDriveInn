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

/5000 \_/ port
localhost hostName
driveinnadmin username
driveinnpostgres maintenance db
Alchemistry password

we also need to get the validator package for the database we get the MOdel VAlidation from iris

HOW TO LEARN DOCKER  
now that we have multiple microservices which may or may not be in mulitple languages have different ways of installing , configuring and deploying the microservices..

so you create the image of the service so that once you givethe image there are standardized

you can run docker containers in cloud on your local pc as well

1. Try to understand the process of building an image
2. try to understand how to run an already made image [ mostly done with a single cli command]
3. Docker swarm or kubernetes to help orchestrate several images

some may say that virtual machines solved the problem before docker

WHAT VIRTUAL MACHINES CONSIST OF

1. A VM contains:

   - The code
   - The dependencies
   - The configurations

   - we run the VM on the physical machine
     through a Hypervisor

   BUT

   the problem is that the VM mahcines have
   their own operating system
   VM must be configured correctly to work
   There is no standard way .. VM are created in different ways

   CONTAINER: its a like a package with your software but technically its an isolated linux process.... so with this you can at the same time use containers with different versions

   image is like the class
   container is like instance of a class

   DOCKER ARCHITECUTURE
   Docker uses a client -server architecture
   The client talks to the daemon(server) which manages docker objects
   There are 4 types of docker objects

   1. Containers
   2. Images: read only templates that contain instructions and meta data for creating Docker containers
   3. Networks
   4. Volumes

\***_Related to Docker images_**
you can use existing images
Create a custom one based on a previous one
Start from scratch

Images can be created with a Dockerfile
one instruction creates a layer
Docker caches so if the instruction is not changed it does not build a new layer for that instruction \***_Related to Docker images_**

Docker containers :they are runnable instances of an image . Multiple containers can be instantiated from the same image

They can be created / started / stopped / deleted using the Docker /Api /Cli

They are isolated from other containers and the host machines using namespaces(linux feature )

image gives a container its one file system

images have

- dependencies
  -configuration
- binaries
- environment variables
- other data container-related

Many containers can be run from a single image, with different options

Many containers communicating with each other come in different ways because remember that the containers are isolated

FOR CONTAINERS TO COMMUNICATE THEY NEED TO BE ON THE SAME NETWORK
5 types of networks

1.  Bridge(default): allows containers connected to communicate, and it provides isolation from containers that are not connected to the Bridge//
    the hello Docker container is mostly on a Bridge network by default
    if we want different classes of containers talking on the same host machine then we need different bridge networks ... so bridge network of one application and for another application
2.  HOST: LIKE I said earlier the containers are isolated by default from the host machine or my laptop as well as other containers but we can make a network to remove that isolation and use a host network to connect the containers to the host machine

3.  Overlay network : to connect Docker daemons . like places the containers on the net . Enables docker swarm

4.  MacVLan : For more legacy applications. Links the containers to unique MAC address that the Docker daemon is having access to

5.  No network : Creating a container that is completely isolated and connected to nothing

DOCKER VOLUMES (storage)

docker ps -a : cli command to see all containers be it running or stopped

docker images

docker network ls

docker volume ls

docker run <container_name> : means run a container

eg: docker run hello-world

docker run nginx  
this runs a container from the nginx image

docker run --name mynginx nginx
this runs a container called "mynginx" from the nginx image

when we run a process [container ] it can end /stop when a process is done it automatically terminates itself

what if we want to make a container that would not stop or terminate itself

DOCKER HUB ; this is where we store the images NOT containers but images

YOU CANNOT REMOVE A RUNNING CONTAINER
you have to stop the container first

docker stop <id/container_name>

docker rm <id/container_name>

docker rm -f <id/container_name>

// choosing the exact version
docker run -p 8080:80 -d nginx:1.24
:latest

**_CREATING YOUR OWN IMAGE : dockerize a node application_**

1. create a .dockerignore file
2. create a dockerfile called Dockerfile capital D

in Dockerfile

// alpine is a reduced version of node
---FROM node:12.6.1-alpine3.11

---WORDKIR /usr/src/app
what happens is creates a directory in the file system of the docker image

//install app depenc
---COPY packages*.json ./  
 the * is a regex incase we dont hvae package-lock.json

---RUN npm install

--- COPY . .  
this means that copy everything in the directory of the dockerfile

// export port 3000
--- EXPOSE 3000

// run the app
CMD ["node", "start"]

.dockerignore the things we dont want to put in another system

HOW TO BUILD YOUR IMAGE CLI COMMAND AFTER MAKING DOCKERFILE AND .dockerignore file

-t this stands for tag
docker build -t <image_name> <directory>

docker build -t nodeapp .  
. means the current directory

HOW TO GET INSIDE A CONTAINER

docker --exec // this is how we start the cli command to get inside a container

exit // to get out a docker container

PUTTING OUR IMAGE ON DOCKER HUB TO TRY IT ON A DIFFERENT MACHINE

Login into Docker
docker login -u <username>
Publish an image to Docker Hub
docker push <username>/<image_name>
eg: docker push nodeapp

Search Hub for an image
docker search <image_name>
Pull an image from a Docker Hub





HOW TO TACKLE 'denied: requested access to the resources is denied'
docker tag <previous_image_name> <new_image_name>:<version>
docker tag nodeapp francescoxx/nodeapp2:1.0.0



if we push an image /image repository to docker hub like pushing code to a github repo 

so if we do that and that image repository does not exist it creates the repository automatically 