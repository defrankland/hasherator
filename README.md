# Hasherator

Golang package that will perform a md5 hash on a whole directory of files and append the result to the  name of each
file within that directory. 

This allows your Go backend to deliver CSS, JS, etc files to the browser without having to clear the file caching
every time there is a change. 

The intended usage is to run it at your application's startup and have it create a working directory, which contains
the same files as the source directory but with the hash string inserted in the name. 


## API
 
 Check out the hasherator_test.go for further example of usage. 
 
 There is just one exported function: 
 
 `func (a *AssetsDir) Run(sourcePath, workingPath string, noHashDirs []string) error`
 
 Feed in the source and destination paths, as well as a list of the names of the directories that you do not want to 
 hash. 
 
 A map is also created on the AssetsDir object that contains the original (source) file name as the key and the hashed 
 file name as the value. 
 
 
 ## Usage 
 
 Intended operation is to perform this on application launch:
 
 
 ```
    assets := AssetsDir{}
    err := assets.Run("./mySource/", "./MyDestination/", []string{"doNotHashThisDirectory", "ThisOneEither"})
 ```
 
 You would then pass the `assets` instance off to your controllers if the `assets.Map` is needed for reference.
 
 If using Go templates for page rendering, you can use the following to reference the hashed file names: 
 
 ```
 <link rel="stylesheet" href='../assets/css/{{index .AssetsMap "bootstrap.min.css"}}'>
 ```
 
 This will look up the key in assets map and render to the associated value. The final product for example should result in 
  bootstrap.min.css pointing to:
 
 ```
  <link rel="stylesheet" href='../assets/css/bootstrap.min-ec3bb52a00e176a7181d454dffaea219.css'>
  ```
 
