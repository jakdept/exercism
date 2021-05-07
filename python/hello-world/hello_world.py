#
# exercism hello_world.py script by Jack Hayhurst
#


def hello(name=""):
    if name is None or name == "":
        return  "Hello, World!"
    return "Hello, " + name + "!"