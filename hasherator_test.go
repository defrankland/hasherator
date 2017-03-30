package hasherator

import (
  "crypto/md5"
  "fmt"
  "io/ioutil"
  "os"
  "strings"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Hasherator", func() {
  var (
    hashDir   string = "./test/hashed/assets/"
    sourceDir string = "./test/assets/"
    assets    AssetsDir
  )

  BeforeEach(func() {
    assets = AssetsDir{}
    err := assets.Run(sourceDir, hashDir, []string{"donothash", "donthashthiscssman", "thisshouldnotbehashedtoo"})

    Expect(err).To(BeNil())
  })

  Context("when hashed directory is created", func() {
    var hashDir []os.FileInfo
    var err error

    BeforeEach(func() {
      hashDir, err = ioutil.ReadDir("./test")
    })

    It("does not return an error", func() {
      Expect(err).ToNot(HaveOccurred())
    })

    It("creates a hashed directories", func() {
      Expect(hashDir[1].Name()).To(Equal("hashed"))
    })
  })

  Context("when files are hashed", func() {
    var src []os.FileInfo
    var hashed []os.FileInfo
    var err error

    Context("when dohash directory is hashed", func() {
      BeforeEach(func() {
        hashed, err = ioutil.ReadDir(hashDir + "/dohash")
        src, err = ioutil.ReadDir(sourceDir + "/dohash")
      })

      It("moves all files to the same place within the hashed directory", func() {

        for i := 0; i < len(src); i++ {
          srcName := src[i].Name()
          if strings.Contains(srcName, ".") {
            srcName = srcName[:strings.LastIndex(srcName, ".")]
          }
          Expect(hashed[i].Name()).To(ContainSubstring(srcName))
        }
      })

      It("does not append a period on a file with no extension", func() {

        for i := 0; i < len(src); i++ {
          if !strings.Contains(src[i].Name(), ".") {
            Expect(hashed[i].Name()).ToNot(ContainSubstring("."))
          }
        }
      })

      It("adds the hash string to the file namein dohash", func() {
        for i := 0; i < len(src); i++ {

          file, _ := ioutil.ReadFile(fmt.Sprintf("%s%s", sourceDir+"/dohash/", src[i].Name()))
          h := md5.Sum(file)
          hash := fmt.Sprintf("-%x", string(h[:16]))

          if !src[i].IsDir() {
            Expect(hashed[i].Name()).To(ContainSubstring(hash))
            Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
          }
        }
      })
    })

    Context("when badboyjones directory is hashed", func() {
      BeforeEach(func() {
        hashed, err = ioutil.ReadDir(hashDir + "/dohash/badboyjones")
        src, err = ioutil.ReadDir(sourceDir + "/dohash/badboyjones")
      })

      It("moves all files to the same place within the hashed directory", func() {

        for i := 0; i < len(src); i++ {
          srcName := src[i].Name()
          if strings.Contains(srcName, ".") {
            srcName = srcName[:strings.LastIndex(srcName, ".")]
          }
          Expect(hashed[i].Name()).To(ContainSubstring(srcName))
          Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
        }
      })

      It("adds the hash string to the file namein badboyjones", func() {
        for i := 0; i < len(src); i++ {

          file, _ := ioutil.ReadFile(fmt.Sprintf("%s%s", sourceDir+"/dohash/badboyjones/", src[i].Name()))
          h := md5.Sum(file)
          hash := fmt.Sprintf("-%x", string(h[:16]))

          if !src[i].IsDir() {
            Expect(hashed[i].Name()).To(ContainSubstring(hash))
            Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
          }
        }

      })

      It("does not append a period on a file with no extension", func() {

        for i := 0; i < len(src); i++ {
          if !strings.Contains(src[i].Name(), ".") {
            Expect(hashed[i].Name()).ToNot(ContainSubstring("."))
            Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
          }
        }
      })
    })

  })

  Context("when files are not hashed", func() {
    var src []os.FileInfo
    var hashed []os.FileInfo
    var err error

    Context("when donothash directory is not hashed", func() {
      BeforeEach(func() {
        hashed, err = ioutil.ReadDir(hashDir + "/donothash")
        src, err = ioutil.ReadDir(sourceDir + "/donothash")
      })

      It("moves all files to the same place with same names", func() {
        for i := 0; i < len(src); i++ {
          Expect(hashed[i].Name()).To(Equal(src[i].Name()))
          if !src[i].IsDir() {
            Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
          }
        }
      })
    })

    Context("when donthashthiscssman directory is not hashed", func() {
      BeforeEach(func() {
        hashed, err = ioutil.ReadDir(hashDir + "/dohash/donthashthiscssman")
        src, err = ioutil.ReadDir(sourceDir + "/dohash/donthashthiscssman")
      })

      It("moves all files to the same place with same names", func() {

        for i := 0; i < len(src); i++ {
          Expect(hashed[i].Name()).To(Equal(src[i].Name()))
          Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
        }
      })
    })

    Context("when thisshouldnotbehashedtoo directory is not hashed", func() {
      BeforeEach(func() {
        hashed, err = ioutil.ReadDir(hashDir + "/dohash/thisshouldnotbehashedtoo")
        src, err = ioutil.ReadDir(sourceDir + "/dohash/thisshouldnotbehashedtoo")
      })

      It("moves all files to the same place with same names", func() {

        for i := 0; i < len(src); i++ {
          Expect(hashed[i].Name()).To(Equal(src[i].Name()))
          Expect(assets.Map[src[i].Name()]).To(Equal(hashed[i].Name()))
        }
      })
    })
  })
})
