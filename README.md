# Base64 Image Encoder

`imgenc` is a basic Base64 image encoder cli tool that takes an image and base64 encodes it to Stdout or a specified file.

## Installation

### Homebrew

```
brew tap Squwid/homebrew-repo
brew install imgenc
```

### Manual

Navigate to downloads page, find the correct architecture/operating system. Set path to the binary and you are ready to go.

## How to Use

Using the CLI tool is extremely easy. Just run the `imgenc` command with an input file, and an optional output file. If no output file is given, the output will be printed to Stdout.

```
imgenc <inputFile> <optional outputFile>
```

## Examples

Lets say I have an image of my dog called `devito.png`. I want to Base64 encode it to easily send it through a POST request body as a string.


```
imgenc devito.png ~/Documents/images/devito-base64.txt
```

If you have a copy to clipboard command as well, it could be used by giving no output and piping to the clipboard command, with something like [this](https://stackoverflow.com/questions/749544/pipe-to-from-the-clipboard-in-bash-script).

```
imgenc devito.png | cb
```