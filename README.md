# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Usage

Place your cursor at the end of some text and press bound keys to complete or add words.

## Autokey scripts

- `complete.py` complete the selected text from your words
- `add.py` add a word to your words using the selected text

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Clone this project:

```
git clone https://github.com/tom-power/complete-word-autokey-rofi.git &&
./complete-word-autokey-rofi
```

Download a matching binary from the latest [release](https://github.com/tom-power/complete-word-autokey-rofi/releases) and copy to your path i.e:

```
wget -O /your/path/completeWord https://github.com/tom-power/complete-word-autokey-rofi/releases/download/v0.1/completeWord_linuxamd64
```

Or to build yourself, install [go](https://golang.org/) then:

```
cd ./go  &&
./build.sh &&
cp build/completeWord /your/path/completeWord

```

Add a `.txt` file containing words to use as suggestions to the `words` directory i.e:

```
cd ./words &&
wget https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english-no-swears.txt
```

After installation you can edit this in `~/.config/complete-word-autokey-rofi/words`, and/or add new words using `add.py`.

Run the install script to copy your words and autokey scripts:

```
./install.sh
```

Bind keys to the [Autokey](https://github.com/autokey/autokey) scripts, you should see them under `scripts/complete-word` in the [Autokey](https://github.com/autokey/autokey) main window.


