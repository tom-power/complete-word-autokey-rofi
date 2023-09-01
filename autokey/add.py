import time

selection = ""
try:
    selection = clipboard.get_selection()
    time.sleep(0.1)
except:
    pass

word = selection
try:
    system.exec_command(
        "echo " + selection +
        " >> ~/.config/complete-word-autokey-rofi/words/added.txt")
    time.sleep(0.1)
except:
    pass

keyboard.send_keys("<right>")
