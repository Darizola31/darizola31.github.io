def main():
    filename = input("Place location of file to read in: ")
    
    file_object  = open(filename, "r")
    
    print(file_object.read())

main();
