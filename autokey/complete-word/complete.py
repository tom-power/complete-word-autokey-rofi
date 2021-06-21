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
    word = system.exec_command("completeWord --complete " + selection)
    time.sleep(0.1)
except:
    pass

try:
    if word and word != selection:
        keyboard.send_keys("%s" % word + " ")
    else:
        keyboard.send_keys("<right>")
except:
    pass
