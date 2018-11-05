def main():
    counter = 0
    wc = 0
    filename = input("Place location of file to read in: ")
    word = input("Which word to locate : ")
    with open(filename) as openfile:
        for line in openfile:
            counter= counter + 1
            for part in line.split():
                wc= wc + 1
                if word in part:
                    print (part)
                    print("Line: ", counter)
                    print ("word count: ", wc)
                    
main();
