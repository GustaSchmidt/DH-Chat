package main

import (
    "fmt"
    "image"
    "image/png"
    "image/color"
    "os"
    "io"

)

func main() {
    //TODO: parameters interpreter descente
    parameter := os.Args
    if len(parameter) < 3{
        fmt.Println("Usa direito, go run sign {path img original} {path img assinatura}")
        return 
    }
    

    if string(parameter[1]) == "sign"{
        input_img := parameter[2]
        input_sig := parameter[3]
        signImg(input_img, input_sig, "sign")

    }else if string(parameter[1]) == "extract"{
        input_signed := parameter[2]
        input_sig := parameter[3]
        signImg(input_signed, input_sig, "extract")

    }else{
        fmt.Println("Usa direito, go run sign/extract {path img original} {path img assinatura}")
    }


}
func signImg(input_img string, input_sig string, action string){
    //Carrega array 2d rgb da imagem original
    img_original := loadImg(input_img)
    img_original_x := len(img_original)
    img_original_y := len(img_original[0])
	
    //Carrega array 2d rgb da assinatuna
    img_sig := loadImg(input_sig)
    img_sig_x := len(img_sig)
    img_sig_y := len(img_sig[0])

    //Verifica se assinatura é maior que objeto assinado
    if (img_sig_x*img_sig_y) > (img_original_x*img_original_y){
        fmt.Println("Assinatura maior que objeto assinado, selecionar imagem menor")
        fmt.Println("tamanho assinatua: "+fmt.Sprintf("%v",img_sig_y)+"x"+fmt.Sprintf("%v",img_sig_x))
        fmt.Println("tamanho imagem: "+fmt.Sprintf("%v",img_original_x)+"x"+fmt.Sprintf("%v",img_original_y))
    }

	
    
    //Criando IMG com assinatura
    width := img_original_y
    height := img_original_x

    img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

    sig_pointer_x := 0
    sig_pointer_y := 0
    // Setando a cor de cada pixel
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            //Inverti as coisa por isso ta ([y][x])
            img_r := uint8(img_original[y][x].R)
            img_g := uint8(img_original[y][x].G)
            img_b := uint8(img_original[y][x].B)
            img_a := uint8(img_original[y][x].A)

            if (sig_pointer_x >= img_sig_x) && (sig_pointer_y >= img_sig_y){
                sig_pointer_x = 0
                sig_pointer_y = 0
            }

            sig_r := uint8(img_sig[sig_pointer_y][sig_pointer_x].R)
            sig_g := uint8(img_sig[sig_pointer_y][sig_pointer_x].G)
            sig_b := uint8(img_sig[sig_pointer_y][sig_pointer_x].B)
            sig_a := uint8(img_sig[sig_pointer_y][sig_pointer_x].A)
            
            
            //TODO: o caululo do bagulho acontece aqui
            if action == "extract"{
                result_r := img_r + binIF(sig_r)
                result_g := img_g + binIF(sig_g)
                result_b := img_b + binIF(sig_b)
                result_a := img_a + binIF(sig_a)
            }else{
                result_r := img_r - binIF(sig_r)
                result_g := img_g - binIF(sig_g)
                result_b := img_b - binIF(sig_b)
                result_a := img_a - binIF(sig_a)
            }

            


            //img.Set(x, y, color.RGBA{img_r, img_g, img_b, img_a})
            img.Set(x, y, color.RGBA{result_r, result_g, result_b, result_a})
            sig_pointer_x++
            sig_pointer_y++
        }
    }

    // encodando tudo em um png
    f, _ := os.Create("image_"+action+".png")
    png.Encode(f, img)
}
func binIF(n uint8) uint8{
    if n > 250{
        return 1
    }
    return 0
}
func loadImg(path string) [][]Pixel{
	// Registra a img
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

    file, err := os.Open(path)

    if err != nil {
        fmt.Println("Error: não deu pra abrir o arquivo "+path)
        os.Exit(1)
    }

    defer file.Close()
	
	//Aqui não nos importamos com memoria, maceta as duas img
    pixels, err := getPixels(file)

    if err != nil {
        fmt.Println("Error: Mano verifica esse formato ae (PNG)")
        os.Exit(1)
    }


    return pixels
}
// x-y array
func getPixels(file io.Reader) ([][]Pixel, error) {
    img, _, err := image.Decode(file)

    if err != nil {
        return nil, err
    }

    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    var pixels [][]Pixel
    for y := 0; y < height; y++ {
        var row []Pixel
        for x := 0; x < width; x++ {
            row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
        }
        pixels = append(pixels, row)
    }

    return pixels, nil
}

// retorna o objeto para um pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{int(r / 256), int(g / 256), int(b / 256), int(a / 256)}
}

// Extrutura do pixel
type Pixel struct {
    R int
    G int
    B int
    A int
}