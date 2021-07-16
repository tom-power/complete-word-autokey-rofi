import time

selection = ""
try:
    keyboard.send_keys("<shift>+<ctrl>+<left>")
    time.sleep(0.1)
    selection = clipboard.get_selection()
    time.sleep(0.1)
except:
    pass

word = selection
try:
    system.exec_command("complete-word --add " + selection)
    time.sleep(0.1)
except:
    pass

keyboard.send_keys("<right>")
