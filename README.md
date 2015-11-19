 1. os.RemoveAll not work correctly on windows platform,
 if there a file which attribute is read only, the RemoveAll function will fail.
 so i wrote a RemoveAllEx function to fix it.
2. add a file exists check fucntion and a direcory exists check function.
