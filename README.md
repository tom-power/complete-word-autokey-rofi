# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Usage

Place your cursor at the end of some text and press bound keys to complete or add words.

## Autokey scripts

- `complete-word/complete.py` complete the selected text from your words
- `complete-word/add.py` add a word to your words using the selected text

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Configure paths in `.local/install.config.sh` and `shell/complete-word/.local/completeWord.config.sh`.

Add a list of words to use ([this](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english-no-swears.txt) for example) to `shell/complete-word/words/.local/`, alternatively skip this step and add words as you go using `complete-word/add.py`.

Run `./install.sh`.

Bind keys to the [Autokey](https://github.com/autokey/autokey) scripts.


