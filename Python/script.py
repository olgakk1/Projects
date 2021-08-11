#print commands to stdout based on files in a directory
import os

subFolders = ["Folder1", "Folder2", "Folder3"]
for subFolder in subFolders:
    allFiles = os.listdir("./Folder/" + subFolder)
    files = [file for file in allFiles if file != "arc"]
    for file in files:
        print("x [edit this] {0} xxx {1}".format(file, subFolder))