# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Usage

Place your cursor at the end of some text and press bound keys to complete or add words.

## Autokey scripts

- `complete-word/complete.py` complete the selected text from your words
- `complete-word/add.py` add a word to your words using the selected text

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Install [go](https://golang.org/) then build a binary:

```
git clone https://github.com/tom-power/complete-word-autokey-rofi.git &&
cd ./complete-word-autokey-rofi/go && ./build.sh
```

Copy a .txt file containing words ([this](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english-no-swears.txt) for example) to be suggested as completions to `./.config/complete-word-autokey-rofi/words`, you can update these directly as you need and add words using `complete-word/add.py`.

Configure your path in `install.config.sh` and run `./install.sh` to copy your binary, config and autokey scripts.

Bind keys to the [Autokey](https://github.com/autokey/autokey) scripts.


