# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Simple word completion for use with (almost) any application.

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Usage

Put your cursor at the end of some text and press your hotkey.

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Configure paths in `.local/install.config.sh` and `shell/complete-word/.local/completeWord.config.sh`.

Add a list of words you'd like to use ([this](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english-no-swears.txt) for example) to `shell/complete-word/words/.local/`

Run `./install.sh`

Add [Autokey](https://github.com/autokey/autokey) for `completeWord.py`

#### TODO

- support completions for selected text (using X11 primary selection at the moment to get text for completion, and can't identify this state vs lack of selection)


