// coded by nekothefox
//fixed by the AI claude sonnet 4.5

package main
import (
        "fmt"
        "image"
        _ "image/jpeg"
        _ "image/png"
        _ "image/gif"
        "os"
        "github.com/chai2010/webp"
)

func fennecWebpEncode(inputPath string, outputPath string) error {
        // Open the input image file
        imagex, err := os.Open(inputPath)
        if err != nil {
                fmt.Printf("failed to open input image: %v\n", err)
                return err
        }
        defer imagex.Close()
        
        // Decode imagex
        imagey, _, err := image.Decode(imagex)
        if err != nil {
                fmt.Printf("failed to decode input image: %v\n", err)
                return err
        }
        
        // Create imagey webp file
        imagez, err := os.Create(outputPath)
        if err != nil {
                fmt.Printf("failed to create output webp file: %v\n", err)
                return err
        }
        defer imagez.Close()
        
        options := &webp.Options{Lossless: true}
        if err := webp.Encode(imagez, imagey, options); err != nil {
                fmt.Printf("failed to encode webp image: %v\n", err)
                return err
        }
        
        return nil
}

func main() {
        args := os.Args
        if len(args) < 3 {
                fmt.Println("Usage: fennec-webp <input> <output>")
                os.Exit(1)
        }
        
        args[2] = args[2] + ".webp"
        fmt.Println("Fennec WebP by Nekothefox\n---------------")
        
        err := fennecWebpEncode(args[1], args[2])
        if err != nil {
                fmt.Println("Error during WebP encoding:")
        } else {
                fmt.Println("Webp image " + args[2] + " created successfully!")
        }
}
