# Word completion for Linux/X11 using [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

![Alt Text](https://github.com/tom-power/complete-word-autokey-rofi/blob/master/assets/demo.gif)

## Usage

Select some text and press the shortcut bound to:

- `complete.py` to suggest words that could be used as replacements
- `add.py` to add the selection to the list of suggestions

## Installation

Install [Autokey](https://github.com/autokey/autokey) and [rofi](https://github.com/davatorium/rofi).

Clone this project and run `sh/install.sh` to make the [Autokey](https://github.com/autokey/autokey) scripts accessible:

```
git clone https://github.com/tom-power/complete-word-autokey-rofi.git &&
cd ./complete-word-autokey-rofi &&
sh/install.sh
```

Add [Autokey](https://github.com/autokey/autokey) shortcuts for the `complete.py` and `add.py` scripts, you should see them under `scripts/complete-word` in the main window.

Add suggestions by copying `.txt` files containing line separated words to `~/.config/complete-words-autokey-rofi/words/`, or by calling `add.py`.
