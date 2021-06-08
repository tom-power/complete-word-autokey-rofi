# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Simple word completion for use with (almost) any application.

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Add the shell scripts to your path:

`cp -r ./shell/* /your/path/`

Add a list of words you'd like to use, [this](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english.txt) for example:

`cp yourWords.txt /your/path/complete-word/words/.local/`

Add the python script to [Autokey](https://github.com/autokey/autokey) and bind it to a hotkey:

`cp -r ./autokey/* ~/.config/autokey/data/scripts/`

## Usage

Put your cursor at the end of some text you'd like to complete and press your hotkey.

#### TODO

- support completions for selected text (using X11 primary selection at the moment to get text for completion, and can't identify this state vs lack of selection)


