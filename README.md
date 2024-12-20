# createDifferentAddressesFromSameSecret

Decode a XRPL secret key (sXXX) into the family of private and public ECDSA keys and the associated XRPL addresses.

It comes from this old proposal: https://pkg.go.dev/hg.sr.ht/~dchapes/ripple@v0.0.0-20220513163231-b31a8ebbcbd9/crypto/rkey

You need the program Visual Studio Code, you can download it in their official website:

https://code.visualstudio.com/

Install in Visual Studio Code the extension Go (we'll use the programming language Go), follow this youtube tutorial for installing it if needed: 

https://www.youtube.com/watch?v=1MXIGYrMk80

Download all the files in this repository in a folder in your computer, they have to look exactly like they look here.

Open that folder with Visual Studio Code.

Open the file "main.go".

Disconnect the internet because you are going to put your secret seed on the program.

In the "const secret = "sp6JS7f14BuwFY8Mw6bTtLKWauoUs" change the seed in the example for your secret seed.

That secret seed will generate as many different addressess as you want, it's just purely a mathematical process. Accounts just become activated if they receive XRP.

In the "Generate(0)" function, you can change that number for any other number, each number results in a new address using the same secret to generate it.

Number (0) results in address rJq5ce8cdbWBsysXx32rvLMV6DUxMwruMT

Number (1) results in address r3B2P7CDYUEYQVEjqwC9hjS3v5UqaBj4Bh

Number (2) results in address rMCHEBdZGABmh8JdBBrPnjmiJhyTKizqG9, and so on.

Each time you change the number, save, open/go to the terminal and type "go run main.go" and enter. The program will run and the result for number 0 will look like this:
secret: sp6JS7f14BuwFY8Mw6bTtLKWauoUs address: rJq5ce8cdbWBsysXx32rvLMV6DUxMwruMT

Have fun creating addresses!

Warning: it's not recommended using the same secret seed for different addresses, it's a better practice use different secret seeds for different addresses.

