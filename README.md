# Hasherator

Golang package that will perform a md5 hash on each file within a directory (and its subdirectories) and append the result to the name of that file. 

This allows your Go backend to deliver CSS, JS, etc. files to the browser without having to clear the file caching
every time there is a change. 

The intended usage is to run it at your application's startup and have it create a working directory, which contains
the same files as the source directory but with the hash string inserted in the name. The reason for doing this is so 
that your project repo can contain files with the non-hashed name, and the working directory can be added to .gitignore. 

Please note that currently the working directory must be different from the source directory. 

## API
 
 Check out the hasherator_test.go for further example of usage. 
 
 There is just one exported function: 
 
 ```go
 func (a *AssetsDir) Run(sourcePath, workingPath string, noHashDirs []string) error
 ```
 
  - sourcePath - directory containing the files you want to hash
  - workingPath - directory where you want the hashed files to go (will be created if it does not exist)
  - noHashDirs - names of any directories that you do not want to hash
 
 A map is also created on the AssetsDir object that contains the original (source) file name as the key and the hashed 
 file name as the value. 
 
 
 ## Usage 
 
 ### Hashing Files
 Intended operation is to perform this on application launch:
 
 ```go
assets := AssetsDir{}
err := assets.Run("./mySource/", "./MyDestination/", []string{"doNotHashThisDirectory", "ThisOneEither"})
 ```
 
The `assets` instance can be passed to the HTTP controllers if the `assets.Map` is needed for reference, or if files need to be re-hashed at runtime.
 
 ### Applying in Go Templates 
 If using Go templates for page rendering, the following can be used to reference the hashed file names: 
 
 ```html
 <link rel="stylesheet" href='../assets/css/{{index .AssetsMap "bootstrap.min.css"}}'>
 ```
 
 This will look up the key in assets map and render to the associated value. The final product for example should result in 
  bootstrap.min.css pointing to:
 
 ```go
  <link rel="stylesheet" href='../assets/css/bootstrap.min-ec3bb52a00e176a7181d454dffaea219.css'>
  ```

### Update Asset Files At Runtime
Another thing that might be useful -- File modifications can be loaded at runtime (without shutting down and restarting the application). Say you have an `originalAssets` instance you created at startup. Then, you could create a goroutine containing:

```go
newAssets := AssetsDir{}
err := assets.Run("./mySource/", "./NewDestination/", []string{"doNotHashThisDirectory", "ThisOneEither"})
if err != nil {
 //whatever
}

for i := 0; i < len(newAssets.Map); i++ {
 if originalAssets.Map[i] != newAssets.Map[i] {
  //delete the originalAssets directory and rename the newAssets directory
  break
 }
}
```

Or better yet, trigger the goroutine with [fsnotify](https://github.com/fsnotify/fsnotify).

```go 
watcher    *fsnotify.Watcher

watcher, err = fsnotify.NewWatcher()
if err != nil {
 //log error
}
defer watcher.Close()

watcher.Add("webapp/templates/")
watcher.Add("webapp/assets/css/")
watcher.Add("webapp/assets/js/")

for {
 select {
 case event := <-watcher.Events:
  if (event.Op&fsnotify.Create == fsnotify.Create ||
   event.Op&fsnotify.Write == fsnotify.Write) &&
   !strings.Contains(event.Name, "___") {

   assets = &hasherator.AssetsDir{}

   err = os.RemoveAll("/webappRuntime2/")
   if err != nil {
    //log error
   }
   
   err := assets.Run("webapp/",
    "webappRuntime2/", 
    []string{"templates", "fonts"})
   if err != nil {
    //log error
   }

   err = os.RemoveAll("webappRuntime/")
   if err != nil {
    //log error
   }
   err = os.Rename("webappRuntime2/", "webappRuntime/")
   if err != nil {
    //log error
   }

   basecontroller.SetAssetsMap(assets.Map)
   //log success message!
  }
 case err := <-watcher.Errors:
  //log error
 }
 time.Sleep(500 * time.Millisecond)
}

```
