# Word completion for Linux/X11 desktops using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Add the shell scripts to your path:

`cp ./shell/* /your/path`

Add a list of words you'd like to use, [this](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english.txt) for example:

`cp yourWords.txt /your/path/complete-word/words/.local/`

Make the python script available to [Autokey](https://github.com/autokey/autokey) and bind it to a hotkey:

`cp ./completeWord.py ~/.config/autokey/data`

## Usage

Put your cursor at the end of some text you'd like to complete and press your hotkey.

#### TODO

- support completions for selected text (using X11 primary selection at the moment to get text for completion, and can't identify this state vs lack of selection)


